// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0
// Inspired by coven (MIT License) by petersunbag

package internal

import (
	"github.com/smgrushb/conv/internal/generics/gslice"
	"github.com/smgrushb/conv/internal/ptr"
	"reflect"
	"unsafe"
)

func IsAnyType[T any](_ ...T) bool {
	return ReflectType[T]() == ptr.AnyType
}

type AnyConverter interface {
	converter
	SetSrcReferDeep(int)
}

type anyConverter struct {
	*convertType
	sReferDeep int
	isTimeType bool
	minUnix    *int64
	asTime     asTime
}

func IsAnyConverter(c converter) (AnyConverter, bool) {
	switch cc := c.(type) {
	case *Converter:
		return cc.isAnyConverter()
	case *anyConverter:
		return cc, true
	}
	return nil, false
}

func newAnyConverter(typ *convertType, srcReferDeep ...int) converter {
	c := &anyConverter{convertType: typ, sReferDeep: gslice.FirstOrZero(srcReferDeep)}
	if c.minUnix = getMinUnix(typ.option, MinUnixTimeAny); c.minUnix != nil {
		for _, v := range TimeWrappers {
			if v.Is(typ.srcTyp) {
				c.isTimeType = true
				c.asTime = v.As
				break
			}
		}
	}
	return c
}

func (a *anyConverter) convert(dPtr, sPtr unsafe.Pointer) bool {
	if a.isTimeType {
		t := a.asTime(sPtr)
		if t.Unix() < *a.minUnix {
			return false
		}
	}
	var val reflect.Value
	val = reflect.NewAt(a.srcTyp, sPtr).Elem()
	if a.sReferDeep > 0 {
		var ptrVal reflect.Value
		valType := a.srcTyp
		for i := 0; i < a.sReferDeep; i++ {
			ptrVal = reflect.New(valType)
			ptrVal.Elem().Set(val)
			val = ptrVal
			valType = ptrVal.Type()
		}
	}
	*(*any)(dPtr) = val.Interface()
	return true
}

func (a *anyConverter) SetSrcReferDeep(deep int) {
	a.sReferDeep = deep
}
