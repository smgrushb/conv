// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0
// Inspired by coven (MIT License) by petersunbag

package convextend

import (
	"fmt"
	"github.com/smgrushb/conv/internal"
	"github.com/smgrushb/conv/internal/generics/gvalue"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reflect"
	"time"
	"unsafe"
)

func init() {
	ProtoConverter = append(ProtoConverter,
		Time2Timestamp(),
		Timestamp2Time(),
	)
}

type TimeWrapper[T any] struct{}

func (t *TimeWrapper[T]) Is(value reflect.Type) bool {
	_, ok := reflect.New(value).Interface().(*T)
	return ok
}

func (t *TimeWrapper[T]) As(p unsafe.Pointer) *time.Time {
	return (*time.Time)(p)
}

type asTime func(unsafe.Pointer) *time.Time

func cvtTimePbTimestamp(as asTime, dPtr, sPtr unsafe.Pointer) {
	*(*timestamppb.Timestamp)(dPtr) = *timestamppb.New(*as(sPtr))
}

func cvtPbTimestampTime(as asTime, dPtr, sPtr unsafe.Pointer) {
	*as(dPtr) = (*timestamppb.Timestamp)(sPtr).AsTime().In(time.Local)
}

type time2Timestamp[T any] struct {
	wrapper TimeWrapper[T]
}

func CustomTime2Timestamp[T any]() internal.CustomConverter {
	return &time2Timestamp[T]{}
}

func Time2Timestamp() internal.CustomConverter {
	return CustomTime2Timestamp[time.Time]()
}

func (s *time2Timestamp[T]) Is(dstTyp, srcTyp reflect.Type) bool {
	if !s.wrapper.Is(srcTyp) {
		return false
	}
	_, ok := reflect.New(dstTyp).Interface().(*timestamppb.Timestamp)
	return ok
}

func (s *time2Timestamp[T]) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		cvtTimePbTimestamp(s.wrapper.As, dPtr, sPtr)
	}
}

func (s *time2Timestamp[T]) Key() string {
	return fmt.Sprintf("[time2Timestamp:%s]", gvalue.ReflectPathType[T]())
}

type timestamp2Time[T any] struct {
	wrapper TimeWrapper[T]
}

func Timestamp2CustomTime[T any]() internal.CustomConverter {
	return &timestamp2Time[T]{}
}

func Timestamp2Time() internal.CustomConverter {
	return Timestamp2CustomTime[time.Time]()
}

func (s *timestamp2Time[T]) Is(dstTyp, srcTyp reflect.Type) bool {
	if !s.wrapper.Is(dstTyp) {
		return false
	}
	_, ok := reflect.New(srcTyp).Interface().(*timestamppb.Timestamp)
	return ok
}

func (s *timestamp2Time[T]) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		cvtPbTimestampTime(s.wrapper.As, dPtr, sPtr)
	}
}

func (s *timestamp2Time[T]) Key() string {
	return fmt.Sprintf("[timestamp2Time:%s]", gvalue.ReflectPathType[T]())
}
