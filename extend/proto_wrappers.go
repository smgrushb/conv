// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0
// Inspired by coven (MIT License) by petersunbag

package convextend

import (
	"github.com/smgrushb/conv/internal"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"reflect"
	"unsafe"
)

func init() {
	ProtoConverter = append(ProtoConverter,
		Bool2BoolValue(),
		BoolValue2Bool(),
		Bytes2BytesValue(),
		BytesValue2Bytes(),
		Float642DoubleValue(),
		DoubleValue2Float64(),
		Float322FloatValue(),
		FloatValue2Float32(),
		Int322Int32Value(),
		Int32Value2Int32(),
		Int642Int64Value(),
		Int64Value2Int64(),
		String2StringValue(),
		StringValue2String(),
		UInt322UInt32Value(),
		UInt32Value2UInt32(),
		UInt642UInt64Value(),
		UInt64Value2UInt64(),
	)
}

// bool start

type bool2BoolValue struct{}

func Bool2BoolValue() internal.CustomConverter {
	return &bool2BoolValue{}
}

func (s *bool2BoolValue) Is(dstTyp, srcTyp reflect.Type) bool {
	if srcTyp.Kind() != reflect.Bool {
		return false
	}
	_, ok := reflect.New(dstTyp).Interface().(*wrapperspb.BoolValue)
	return ok
}

func (s *bool2BoolValue) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		*(*wrapperspb.BoolValue)(dPtr) = *wrapperspb.Bool(*(*bool)(sPtr))
	}
}

func (s *bool2BoolValue) Key() string {
	return "[bool2BoolValue]"
}

type boolValue2Bool struct{}

func BoolValue2Bool() internal.CustomConverter {
	return &boolValue2Bool{}
}

func (s *boolValue2Bool) Is(dstTyp, srcTyp reflect.Type) bool {
	if dstTyp.Kind() != reflect.Bool {
		return false
	}
	_, ok := reflect.New(srcTyp).Interface().(*wrapperspb.BoolValue)
	return ok
}

func (s *boolValue2Bool) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		*(*bool)(dPtr) = (*wrapperspb.BoolValue)(sPtr).GetValue()
	}
}

func (s *boolValue2Bool) Key() string {
	return "[boolValue2Bool]"
}

// bool end

// bytes start

type bytes2BytesValue struct{}

func Bytes2BytesValue() internal.CustomConverter {
	return &bytes2BytesValue{}
}

func (s *bytes2BytesValue) Is(dstTyp, srcTyp reflect.Type) bool {
	if srcTyp.Kind() != reflect.Slice || srcTyp.Elem().Kind() != reflect.Uint8 {
		return false
	}
	_, ok := reflect.New(dstTyp).Interface().(*wrapperspb.BytesValue)
	return ok
}

func (s *bytes2BytesValue) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		*(*wrapperspb.BytesValue)(dPtr) = *wrapperspb.Bytes(*(*[]byte)(sPtr))
	}
}

func (s *bytes2BytesValue) Key() string {
	return "[bytes2BytesValue]"
}

type bytesValue2Bytes struct{}

func BytesValue2Bytes() internal.CustomConverter {
	return &bytesValue2Bytes{}
}

func (s *bytesValue2Bytes) Is(dstTyp, srcTyp reflect.Type) bool {
	if dstTyp.Kind() != reflect.Slice || dstTyp.Elem().Kind() != reflect.Uint8 {
		return false
	}
	_, ok := reflect.New(srcTyp).Interface().(*wrapperspb.BytesValue)
	return ok
}

func (s *bytesValue2Bytes) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		*(*[]byte)(dPtr) = (*wrapperspb.BytesValue)(sPtr).GetValue()
	}
}

func (s *bytesValue2Bytes) Key() string {
	return "[bytesValue2Bytes]"
}

// bytes end

// float64 start

type float642DoubleValue struct{}

func Float642DoubleValue() internal.CustomConverter {
	return &float642DoubleValue{}
}

func (s *float642DoubleValue) Is(dstTyp, srcTyp reflect.Type) bool {
	if srcTyp.Kind() != reflect.Float64 {
		return false
	}
	_, ok := reflect.New(dstTyp).Interface().(*wrapperspb.DoubleValue)
	return ok
}

func (s *float642DoubleValue) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		*(*wrapperspb.DoubleValue)(dPtr) = *wrapperspb.Double(*(*float64)(sPtr))
	}
}

func (s *float642DoubleValue) Key() string {
	return "[float642DoubleValue]"
}

type doubleValue2Float64 struct{}

func DoubleValue2Float64() internal.CustomConverter {
	return &doubleValue2Float64{}
}

func (s *doubleValue2Float64) Is(dstTyp, srcTyp reflect.Type) bool {
	if dstTyp.Kind() != reflect.Float64 {
		return false
	}
	_, ok := reflect.New(srcTyp).Interface().(*wrapperspb.DoubleValue)
	return ok
}

func (s *doubleValue2Float64) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		*(*float64)(dPtr) = (*wrapperspb.DoubleValue)(sPtr).GetValue()
	}
}

func (s *doubleValue2Float64) Key() string {
	return "[doubleValue2Float64]"
}

// float64 end

// float32 start

type float322FloatValue struct{}

func Float322FloatValue() internal.CustomConverter {
	return &float322FloatValue{}
}

func (s *float322FloatValue) Is(dstTyp, srcTyp reflect.Type) bool {
	if srcTyp.Kind() != reflect.Float32 {
		return false
	}
	_, ok := reflect.New(dstTyp).Interface().(*wrapperspb.FloatValue)
	return ok
}

func (s *float322FloatValue) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		*(*wrapperspb.FloatValue)(dPtr) = *wrapperspb.Float(*(*float32)(sPtr))
	}
}

func (s *float322FloatValue) Key() string {
	return "[float322FloatValue]"
}

type floatValue2Float32 struct{}

func FloatValue2Float32() internal.CustomConverter {
	return &floatValue2Float32{}
}

func (s *floatValue2Float32) Is(dstTyp, srcTyp reflect.Type) bool {
	if dstTyp.Kind() != reflect.Float32 {
		return false
	}
	_, ok := reflect.New(srcTyp).Interface().(*wrapperspb.FloatValue)
	return ok
}

func (s *floatValue2Float32) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		*(*float32)(dPtr) = (*wrapperspb.FloatValue)(sPtr).GetValue()
	}
}

func (s *floatValue2Float32) Key() string {
	return "[floatValue2Float32]"
}

// float32 end

// int32 start

type int322Int32Value struct{}

func Int322Int32Value() internal.CustomConverter {
	return &int322Int32Value{}
}

func (s *int322Int32Value) Is(dstTyp, srcTyp reflect.Type) bool {
	if srcTyp.Kind() != reflect.Int32 {
		return false
	}
	_, ok := reflect.New(dstTyp).Interface().(*wrapperspb.Int32Value)
	return ok
}

func (s *int322Int32Value) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		*(*wrapperspb.Int32Value)(dPtr) = *wrapperspb.Int32(*(*int32)(sPtr))
	}
}

func (s *int322Int32Value) Key() string {
	return "[int322Int32Value]"
}

type int32Value2Int32 struct{}

func Int32Value2Int32() internal.CustomConverter {
	return &int32Value2Int32{}
}

func (s *int32Value2Int32) Is(dstTyp, srcTyp reflect.Type) bool {
	if dstTyp.Kind() != reflect.Int32 {
		return false
	}
	_, ok := reflect.New(srcTyp).Interface().(*wrapperspb.Int32Value)
	return ok
}

func (s *int32Value2Int32) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		*(*int32)(dPtr) = (*wrapperspb.Int32Value)(sPtr).GetValue()
	}
}

func (s *int32Value2Int32) Key() string {
	return "[int32Value2Int32]"
}

// int32 end

// int64 start

type int642Int64Value struct{}

func Int642Int64Value() internal.CustomConverter {
	return &int642Int64Value{}
}

func (s *int642Int64Value) Is(dstTyp, srcTyp reflect.Type) bool {
	if srcTyp.Kind() != reflect.Int64 {
		return false
	}
	_, ok := reflect.New(dstTyp).Interface().(*wrapperspb.Int64Value)
	return ok
}

func (s *int642Int64Value) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		*(*wrapperspb.Int64Value)(dPtr) = *wrapperspb.Int64(*(*int64)(sPtr))
	}
}

func (s *int642Int64Value) Key() string {
	return "[int642Int64Value]"
}

type int64Value2Int64 struct{}

func Int64Value2Int64() internal.CustomConverter {
	return &int64Value2Int64{}
}

func (s *int64Value2Int64) Is(dstTyp, srcTyp reflect.Type) bool {
	if dstTyp.Kind() != reflect.Int64 {
		return false
	}
	_, ok := reflect.New(srcTyp).Interface().(*wrapperspb.Int64Value)
	return ok
}

func (s *int64Value2Int64) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		*(*int64)(dPtr) = (*wrapperspb.Int64Value)(sPtr).GetValue()
	}
}

func (s *int64Value2Int64) Key() string {
	return "[int64Value2Int64]"
}

// int64 end

// string start

type string2StringValue struct{}

func String2StringValue() internal.CustomConverter {
	return &string2StringValue{}
}

func (s *string2StringValue) Is(dstTyp, srcTyp reflect.Type) bool {
	if srcTyp.Kind() != reflect.String {
		return false
	}
	_, ok := reflect.New(dstTyp).Interface().(*wrapperspb.StringValue)
	return ok
}

func (s *string2StringValue) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		*(*wrapperspb.StringValue)(dPtr) = *wrapperspb.String(*(*string)(sPtr))
	}
}

func (s *string2StringValue) Key() string {
	return "[string2StringValue]"
}

type stringValue2String struct{}

func StringValue2String() internal.CustomConverter {
	return &stringValue2String{}
}

func (s *stringValue2String) Is(dstTyp, srcTyp reflect.Type) bool {
	if dstTyp.Kind() != reflect.String {
		return false
	}
	_, ok := reflect.New(srcTyp).Interface().(*wrapperspb.StringValue)
	return ok
}

func (s *stringValue2String) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		*(*string)(dPtr) = (*wrapperspb.StringValue)(sPtr).GetValue()
	}
}

func (s *stringValue2String) Key() string {
	return "[stringValue2String]"
}

// string end

// uint32 start

type uint322UInt32Value struct{}

func UInt322UInt32Value() internal.CustomConverter {
	return &uint322UInt32Value{}
}

func (s *uint322UInt32Value) Is(dstTyp, srcTyp reflect.Type) bool {
	if srcTyp.Kind() != reflect.Uint32 {
		return false
	}
	_, ok := reflect.New(dstTyp).Interface().(*wrapperspb.UInt32Value)
	return ok
}

func (s *uint322UInt32Value) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		*(*wrapperspb.UInt32Value)(dPtr) = *wrapperspb.UInt32(*(*uint32)(sPtr))
	}
}

func (s *uint322UInt32Value) Key() string {
	return "[uint322UInt32Value]"
}

type uint32Value2UInt32 struct{}

func UInt32Value2UInt32() internal.CustomConverter {
	return &uint32Value2UInt32{}
}

func (s *uint32Value2UInt32) Is(dstTyp, srcTyp reflect.Type) bool {
	if dstTyp.Kind() != reflect.Uint32 {
		return false
	}
	_, ok := reflect.New(srcTyp).Interface().(*wrapperspb.UInt32Value)
	return ok
}

func (s *uint32Value2UInt32) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		*(*uint32)(dPtr) = (*wrapperspb.UInt32Value)(sPtr).GetValue()
	}
}

func (s *uint32Value2UInt32) Key() string {
	return "[uint32Value2UInt32]"
}

// uint32 end

// uint64 start

type uint642UInt64Value struct{}

func UInt642UInt64Value() internal.CustomConverter {
	return &uint642UInt64Value{}
}

func (s *uint642UInt64Value) Is(dstTyp, srcTyp reflect.Type) bool {
	if srcTyp.Kind() != reflect.Uint64 {
		return false
	}
	_, ok := reflect.New(dstTyp).Interface().(*wrapperspb.UInt64Value)
	return ok
}

func (s *uint642UInt64Value) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		*(*wrapperspb.UInt64Value)(dPtr) = *wrapperspb.UInt64(*(*uint64)(sPtr))
	}
}

func (s *uint642UInt64Value) Key() string {
	return "[uint642UInt64Value]"
}

type uint64Value2UInt64 struct{}

func UInt64Value2UInt64() internal.CustomConverter {
	return &uint64Value2UInt64{}
}

func (s *uint64Value2UInt64) Is(dstTyp, srcTyp reflect.Type) bool {
	if dstTyp.Kind() != reflect.Uint64 {
		return false
	}
	_, ok := reflect.New(srcTyp).Interface().(*wrapperspb.UInt64Value)
	return ok
}

func (s *uint64Value2UInt64) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		*(*uint64)(dPtr) = (*wrapperspb.UInt64Value)(sPtr).GetValue()
	}
}

func (s *uint64Value2UInt64) Key() string {
	return "[uint64Value2UInt64]"
}

// uint64 end
