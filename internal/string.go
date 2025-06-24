// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0
// Inspired by coven (MIT License) by petersunbag

package internal

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unsafe"
)

var (
	stringerType  = ReflectType[fmt.Stringer]()
	marshalerType = ReflectType[json.Marshaler]()
)

type stringsConverter struct {
	*convertType
}

func (s *stringsConverter) convert(dPtr, sPtr unsafe.Pointer) {
	if strings, ok := reflect.NewAt(s.srcTyp, sPtr).Interface().(fmt.Stringer); ok {
		*(*string)(dPtr) = strings.String()
	}
}

type marshalJsonConverter struct {
	*convertType
}

func (m *marshalJsonConverter) convert(dPtr, sPtr unsafe.Pointer) {
	if marshaler, ok := reflect.NewAt(m.srcTyp, sPtr).Interface().(json.Marshaler); ok {
		if bs, err := marshaler.MarshalJSON(); err == nil {
			*(*string)(dPtr) = string(bs)
		}
	}
}

func newStringsConverter(typ *convertType) converter {
	return &stringsConverter{convertType: typ}
}

func newMarshalJsonConverter(typ *convertType) converter {
	return &marshalJsonConverter{convertType: typ}
}
