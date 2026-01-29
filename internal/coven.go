// This file is part of the original coven project:
// https://github.com/petersunbag/coven
//
// Copyright © 2018 petersunbag
// Licensed under the MIT License
// https://opensource.org/licenses/MIT
// Modified by smgrushb in 2025

package internal

import (
	"fmt"
	"github.com/smgrushb/conv/internal/generics/gvalue"
	"github.com/smgrushb/conv/internal/ptr"
	"reflect"
	"sync"
	"unsafe"
)

var (
	createdConvertersMu sync.Mutex
	createdConverters   = make(map[convertTypeKey]*Converter)

	zeroReflectValue reflect.Value
)

type convertType struct {
	dstTyp reflect.Type
	srcTyp reflect.Type
	option *StructOption
}

// 因为StructOption不可比，增加此类型来判断是否为同一个convertType
type convertTypeKey struct {
	dstTyp    reflect.Type
	srcTyp    reflect.Type
	optionKey string
}

func (c *convertType) key() convertTypeKey {
	return convertTypeKey{
		dstTyp:    c.dstTyp,
		srcTyp:    c.srcTyp,
		optionKey: c.option.key(),
	}
}

type converter interface {
	convert(dPtr, sPtr unsafe.Pointer)
}

type Converter struct {
	*convertType
	converter
}

func (c *Converter) Convert(dst, src any) error {
	if gvalue.IsNil(dst) || gvalue.IsNil(src) {
		return nil
	}
	dv := dereferencedValue(dst)
	if dv == zeroReflectValue {
		return nil
	}
	if !dv.CanSet() {
		return fmt.Errorf("[conv]destination should be a pointer. [actual:%v]", dv.Type())
	}
	if dv.Type() != c.dstTyp {
		return fmt.Errorf("[conv]invalid destination type. [expected:%v] [actual:%v]", c.dstTyp, dv.Type())
	}
	sv := dereferencedValue(src)
	if sv == zeroReflectValue {
		return nil
	}
	if !sv.CanAddr() {
		return fmt.Errorf("[conv]source should be a pointer. [actual:%v]", sv.Type())
	}
	if sv.Type() != c.srcTyp {
		return fmt.Errorf("[conv]invalid source type. [expected:%v] [actual:%v]", c.srcTyp, sv.Type())
	}
	c.converter.convert(unsafe.Pointer(dv.UnsafeAddr()), unsafe.Pointer(sv.UnsafeAddr()))
	return nil
}

func (c *Converter) isAnyConverter() (AnyConverter, bool) {
	return IsAnyConverter(c.converter)
}

func NewConverter(dstTyp, srcTyp reflect.Type, option *StructOption) *Converter {
	createdConvertersMu.Lock()
	defer createdConvertersMu.Unlock()
	return newConverter(dstTyp, srcTyp, option)
}

func newConverter(dstTyp, srcTyp reflect.Type, option *StructOption) *Converter {
	var sReferDeep int
	dstTyp, _ = dereferencedType(dstTyp)
	srcTyp, sReferDeep = dereferencedType(srcTyp)
	cTyp := &convertType{dstTyp: dstTyp, srcTyp: srcTyp, option: option}
	key := cTyp.key()
	if dc, ok := createdConverters[key]; ok {
		return dc
	}
	var c converter
	if option != nil {
		for _, v := range option.CustomConv {
			if v.Is(dstTyp, srcTyp) {
				c = Custom(v.Converter())
				break
			}
		}
	}
	if c == nil {
		if option != nil && dstTyp.Kind() == reflect.String {
			if k := srcTyp.Kind(); option.SerializeToString &&
				gvalue.NotIn(k, reflect.Invalid, reflect.Complex64, reflect.Complex128, reflect.Chan, reflect.Func, reflect.UnsafePointer) {
				c = newSerializeConverter(cTyp)
			} else if srcPtr := reflect.PointerTo(srcTyp); option.UseStrings && srcPtr.Implements(stringerType) {
				c = newStringsConverter(cTyp)
			} else if option.UseMarshal && srcPtr.Implements(marshalerType) {
				c = newMarshalJsonConverter(cTyp)
			}
		}
	}
	if c == nil {
		c = newBasicConverter(cTyp)
	}
	if c == nil {
		if dstTyp == ptr.AnyType {
			c = newAnyConverter(cTyp, sReferDeep)
		} else {
			switch sk, dk := srcTyp.Kind(), dstTyp.Kind(); {
			// todo: 数组转换
			case sk == reflect.Struct && dk == reflect.Struct:
				c = newStructConverter(cTyp)
			case sk == reflect.Slice && dk == reflect.Slice:
				c = newSliceConverter(cTyp)
			case sk == reflect.Map && dk == reflect.Map:
				if dstTyp.Elem() == ptr.AnyType || dstTyp.Elem().Kind() != reflect.Interface {
					c = newMapConverter(cTyp)
				}
			case sk == reflect.Struct && dk == reflect.Map:
				if dstTyp.Key().Kind() == reflect.String && (dstTyp.Elem() == ptr.AnyType || dstTyp.Elem().Kind() == reflect.String) {
					c = newStructMapConverter(cTyp, dstTyp.Elem())
				}
			case sk == reflect.Map && dk == reflect.Struct:
				if srcTyp.Key().Kind() == reflect.String && (srcTyp.Elem() == ptr.AnyType || srcTyp.Elem().Kind() == reflect.String) {
					// TODO: map[string]any/map[string]string转结构体
				}
			default:
				c = newTimeConverter(cTyp)
			}
		}
	}
	if c != nil {
		// 可能预注册进去了，那就不要再注册
		if _, ok := createdConverters[key]; !ok {
			dc := &Converter{convertType: cTyp, converter: c}
			createdConverters[key] = dc
			return dc
		}
		return createdConverters[key]
	}
	return nil
}
