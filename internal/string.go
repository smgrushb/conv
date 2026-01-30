// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0
// Inspired by coven (MIT License) by petersunbag

package internal

import (
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/smgrushb/conv/internal/generics/gvalue"
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

func (s *stringsConverter) convert(dPtr, sPtr unsafe.Pointer) bool {
	if strings, ok := reflect.NewAt(s.srcTyp, sPtr).Interface().(fmt.Stringer); ok {
		*(*string)(dPtr) = strings.String()
		return true
	}
	return false
}

type marshalJsonConverter struct {
	*convertType
}

func (m *marshalJsonConverter) convert(dPtr, sPtr unsafe.Pointer) bool {
	if marshaler, ok := reflect.NewAt(m.srcTyp, sPtr).Interface().(json.Marshaler); ok {
		if bs, err := marshaler.MarshalJSON(); err == nil {
			*(*string)(dPtr) = string(bs)
			return true
		}
	}
	return false
}

type serializeConverter struct {
	*convertType
	nilValuePolicy NilValuePolicy
}

func (s *serializeConverter) convert(dPtr, sPtr unsafe.Pointer) bool {
	p := reflect.NewAt(s.srcTyp, sPtr)
	if p.IsNil() {
		if s.nilValuePolicy == NilValuePolicyIgnore {
			return false
		}
		p = reflect.New(s.srcTyp)
	}
	e := p.Elem()
	if k := e.Kind(); gvalue.In(k, reflect.Interface, reflect.Pointer, reflect.Slice, reflect.Map) && e.IsNil() {
		if s.nilValuePolicy == NilValuePolicyIgnore {
			return false
		}
		switch k {
		case reflect.Slice:
			e = reflect.MakeSlice(s.srcTyp, 0, 0)
		case reflect.Map:
			e = reflect.MakeMap(s.srcTyp)
		default:
			e = reflect.New(s.srcTyp).Elem()
		}
	}
	if str, err := sonic.MarshalString(e.Interface()); err == nil {
		*(*string)(dPtr) = str
		return true
	}
	return false
}

func newStringsConverter(typ *convertType) converter {
	return &stringsConverter{convertType: typ}
}

func newMarshalJsonConverter(typ *convertType) converter {
	return &marshalJsonConverter{convertType: typ}
}

func newSerializeConverter(typ *convertType, nilValuePolicy NilValuePolicy) converter {
	return &serializeConverter{convertType: typ, nilValuePolicy: nilValuePolicy}
}
