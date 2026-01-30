// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0
// Inspired by coven (MIT License) by petersunbag

package convextend

import (
	"github.com/smgrushb/conv/internal"
	"google.golang.org/protobuf/types/known/durationpb"
	"reflect"
	"time"
	"unsafe"
)

func init() {
	ProtoConverter = append(ProtoConverter,
		TimeDuration2PbDuration(),
		PbDuration2TimeDuration(),
	)
}

type timeDuration2PbDuration struct{}

func TimeDuration2PbDuration() internal.CustomConverterV2 {
	return &timeDuration2PbDuration{}
}

func (s *timeDuration2PbDuration) Is(dstTyp, srcTyp reflect.Type) bool {
	if srcTyp.Kind() != reflect.Int64 {
		return false
	}
	_, ok := reflect.New(dstTyp).Interface().(*durationpb.Duration)
	return ok
}

func (s *timeDuration2PbDuration) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) bool {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) bool {
		*(*durationpb.Duration)(dPtr) = *durationpb.New(*(*time.Duration)(sPtr))
		return true
	}
}

func (s *timeDuration2PbDuration) Key() string {
	return "[timeDuration2PbDuration]"
}

type pbDuration2TimeDuration struct{}

func PbDuration2TimeDuration() internal.CustomConverterV2 {
	return &pbDuration2TimeDuration{}
}

func (s *pbDuration2TimeDuration) Is(dstTyp, srcTyp reflect.Type) bool {
	if dstTyp.Kind() != reflect.Int64 {
		return false
	}
	_, ok := reflect.New(srcTyp).Interface().(*durationpb.Duration)
	return ok
}

func (s *pbDuration2TimeDuration) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) bool {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) bool {
		*(*time.Duration)(dPtr) = (*durationpb.Duration)(sPtr).AsDuration()
		return true
	}
}

func (s *pbDuration2TimeDuration) Key() string {
	return "[pbDuration2TimeDuration]"
}
