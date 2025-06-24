// This file is part of the original coven project:
// https://github.com/petersunbag/coven
//
// Copyright © 2018 petersunbag
// Licensed under the MIT License
// https://opensource.org/licenses/MIT
// Modified by smgrushb in 2025

package internal

import (
	"github.com/smgrushb/conv/internal/generics/gptr"
	"github.com/smgrushb/conv/internal/ptr"
	"github.com/smgrushb/conv/internal/unsafeheader"
	"reflect"
	"unsafe"
)

type mapConverter struct {
	*convertType
	dKeyType           reflect.Type
	dValType           reflect.Type
	keyConverter       *elemConverter
	valConverter       *elemConverter
	dEmptyMapInterface *unsafeheader.EmptyInterface
	sEmptyMapInterface *unsafeheader.EmptyInterface
	enable             bool // 兜底
}

func newMapConverter(typ *convertType) converter {
	dKeyTyp, sKeyTyp := typ.dstTyp.Key(), typ.srcTyp.Key()
	dValTyp, sValTyp := typ.dstTyp.Elem(), typ.srcTyp.Elem()
	c := &mapConverter{
		convertType:        typ,
		dKeyType:           dKeyTyp,
		dValType:           dValTyp,
		dEmptyMapInterface: (*unsafeheader.EmptyInterface)(unsafe.Pointer(gptr.Of(reflect.New(typ.dstTyp).Interface()))),
		sEmptyMapInterface: (*unsafeheader.EmptyInterface)(unsafe.Pointer(gptr.Of(reflect.New(typ.srcTyp).Interface()))),
	}
	key := typ.key()
	// 先预注册进去，不然循环依赖下会循环解析
	createdConverters[key] = &Converter{convertType: typ, converter: c}
	kc, ok := newElemConverter(dKeyTyp, sKeyTyp, typ.option)
	if !ok {
		// 把预注册的内容删了
		delete(createdConverters, key)
		return nil
	}
	c.keyConverter = kc
	vc, ok := newElemConverter(dValTyp, sValTyp, typ.option)
	if !ok {
		// 把预注册的内容删了
		delete(createdConverters, key)
		return nil
	}
	c.valConverter = vc
	c.enable = true
	return c
}

func (m *mapConverter) convert(dPtr, sPtr unsafe.Pointer) {
	if !m.enable {
		return
	}
	sv := ptrToMapValue(m.sEmptyMapInterface, sPtr)
	dv := ptrToMapValue(m.dEmptyMapInterface, dPtr)
	keys := sv.MapKeys()
	if dv.IsNil() {
		dv.Set(reflect.MakeMapWithSize(m.dstTyp, len(keys)))
	}
	for _, sKey := range keys {
		val := sv.MapIndex(sKey)
		var sValPtr unsafe.Pointer
		if val.Type() == ptr.AnyType {
			sValPtr = PtrOfAny(val)
		} else {
			sValPtr = valuePtr(val)
		}
		sKeyPtr := valuePtr(sKey)
		dKey := reflect.New(m.dKeyType).Elem()
		dVal := reflect.New(m.dValType).Elem()
		m.keyConverter.convert(unsafe.Pointer(dKey.UnsafeAddr()), sKeyPtr)
		m.valConverter.convert(unsafe.Pointer(dVal.UnsafeAddr()), sValPtr)
		dv.SetMapIndex(dKey, dVal)
	}
}

func ptrToMapValue(emptyMapInterface *unsafeheader.EmptyInterface, ptr unsafe.Pointer) reflect.Value {
	realInterface := *(*any)(unsafe.Pointer(gptr.Of(unsafeheader.NewEmptyInterface(emptyMapInterface.GetTyp(), ptr))))
	return reflect.ValueOf(realInterface).Elem()
}

func valuePtr(v reflect.Value) unsafe.Pointer {
	inter := v.Interface()
	return (*unsafeheader.EmptyInterface)(unsafe.Pointer(&inter)).GetWord()
}
