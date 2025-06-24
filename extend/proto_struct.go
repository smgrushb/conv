// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0
// Inspired by coven (MIT License) by petersunbag

package convextend

import (
	"github.com/smgrushb/conv/internal"
	"github.com/smgrushb/conv/internal/ptr"
	"google.golang.org/protobuf/types/known/structpb"
	"reflect"
	"unsafe"
)

func init() {
	ProtoConverter = append(ProtoConverter,
		Anys2ListValue(),
		ListValue2Anys(),
		H2Struct(),
		Struct2H(),
		Any2Value(),
		Value2Any(),
	)
}

// []any start

type anys2ListValue struct{}

func Anys2ListValue() internal.CustomConverter {
	return &anys2ListValue{}
}

func (s *anys2ListValue) Is(dstTyp, srcTyp reflect.Type) bool {
	if srcTyp.Kind() != reflect.Slice || srcTyp.Elem() != ptr.AnyType {
		return false
	}
	_, ok := reflect.New(dstTyp).Interface().(*structpb.ListValue)
	return ok
}

func (s *anys2ListValue) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		if list, err := structpb.NewList(*(*[]any)(sPtr)); err == nil {
			*(*structpb.ListValue)(dPtr) = *list
		}
	}
}

func (s *anys2ListValue) Key() string {
	return "[anys2ListValue]"
}

type listValue2Anys struct{}

func ListValue2Anys() internal.CustomConverter {
	return &listValue2Anys{}
}

func (s *listValue2Anys) Is(dstTyp, srcTyp reflect.Type) bool {
	if dstTyp.Kind() != reflect.Slice || dstTyp.Elem() != ptr.AnyType {
		return false
	}
	_, ok := reflect.New(srcTyp).Interface().(*structpb.ListValue)
	return ok
}

func (s *listValue2Anys) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		list := (*structpb.ListValue)(sPtr).GetValues()
		anys := make([]any, len(list))
		for i, v := range list {
			anys[i] = v.AsInterface()
		}
		*(*[]any)(dPtr) = anys
	}
}

func (s *listValue2Anys) Key() string {
	return "[listValue2Anys]"
}

// []any end

// map[string]any start

type h2Struct struct{}

func H2Struct() internal.CustomConverter {
	return &h2Struct{}
}

func (s *h2Struct) Is(dstTyp, srcTyp reflect.Type) bool {
	if srcTyp.Kind() != reflect.Map || srcTyp.Key().Kind() != reflect.String || srcTyp.Elem() != ptr.AnyType {
		return false
	}
	_, ok := reflect.New(dstTyp).Interface().(*structpb.Struct)
	return ok
}

func (s *h2Struct) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		if v, err := structpb.NewStruct(*(*map[string]any)(sPtr)); err == nil {
			*(*structpb.Struct)(dPtr) = *v
		}
	}
}

func (s *h2Struct) Key() string {
	return "[h2Struct]"
}

type struct2H struct{}

func Struct2H() internal.CustomConverter {
	return &struct2H{}
}

func (s *struct2H) Is(dstTyp, srcTyp reflect.Type) bool {
	if dstTyp.Kind() != reflect.Map || dstTyp.Key().Kind() != reflect.String || dstTyp.Elem() != ptr.AnyType {
		return false
	}
	_, ok := reflect.New(srcTyp).Interface().(*structpb.Struct)
	return ok
}

func (s *struct2H) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		*(*map[string]any)(dPtr) = (*structpb.Struct)(sPtr).AsMap()
	}
}

func (s *struct2H) Key() string {
	return "[struct2H]"
}

// map[string]any end

// any start

type any2Value struct{}

func Any2Value() internal.CustomConverter {
	return &any2Value{}
}

func (s *any2Value) Is(dstTyp, srcTyp reflect.Type) bool {
	if srcTyp != ptr.AnyType {
		return false
	}
	_, ok := reflect.New(dstTyp).Interface().(*structpb.Value)
	return ok
}

func (s *any2Value) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		if v, err := structpb.NewValue(*(*any)(sPtr)); err == nil {
			*(*structpb.Value)(dPtr) = *v
		}
	}
}

func (s *any2Value) Key() string {
	return "[any2Value]"
}

type value2Any struct{}

func Value2Any() internal.CustomConverter {
	return &value2Any{}
}

func (s *value2Any) Is(dstTyp, srcTyp reflect.Type) bool {
	if dstTyp != ptr.AnyType {
		return false
	}
	_, ok := reflect.New(srcTyp).Interface().(*structpb.Value)
	return ok
}

func (s *value2Any) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		*(*any)(dPtr) = (*structpb.Value)(sPtr).AsInterface()
	}
}

func (s *value2Any) Key() string {
	return "[value2Any]"
}

// any end
