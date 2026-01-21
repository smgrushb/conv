// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0
// Inspired by coven (MIT License) by petersunbag

package internal

import (
	"github.com/smgrushb/conv/internal/generics/gvalue"
	"reflect"
	"time"
	"unsafe"
)

type timestampAccuracy int64

const (
	UnixTimestamp timestampAccuracy = iota
	UnixMilliTimestamp
	UnixMicroTimestamp
)

var (
	TimeWrappers      []timeWrapper
	TimeFormat        = "2006-01-02 15:04:05"
	TimestampAccuracy = UnixTimestamp
	MinUnix           *int64
	MinUnixScene      = DefaultMinUnixScene
)

func init() {
	TimeWrappers = append(TimeWrappers, &TimeWrapper[time.Time]{})
}

type timeWrapper interface {
	Is(reflect.Type) bool
	As(unsafe.Pointer) *time.Time
	GetFormat() (string, bool)
}

type TimeWrapper[T any] struct{}

func (t *TimeWrapper[T]) Is(value reflect.Type) bool {
	_, ok := reflect.New(value).Interface().(*T)
	return ok
}

func (t *TimeWrapper[T]) As(p unsafe.Pointer) *time.Time {
	return (*time.Time)(p)
}

func (t *TimeWrapper[T]) GetFormat() (string, bool) {
	if f, ok := gvalue.TypeAs[interface{ GetFormat() string }, T](); ok {
		return f.GetFormat(), true
	}
	return "", false
}

type asTime func(unsafe.Pointer) *time.Time

type timeConverter struct {
	format  string
	minUnix *int64
	as      asTime
	cvtOp   func(string, *int64, asTime, unsafe.Pointer, unsafe.Pointer)
}

func (t *timeConverter) convert(dPtr, sPtr unsafe.Pointer) {
	t.cvtOp(t.format, t.minUnix, t.as, dPtr, sPtr)
}

func cvtTimeString(format string, minUnix *int64, as asTime, dPtr, sPtr unsafe.Pointer) {
	t := as(sPtr)
	if minUnix != nil && t.Unix() < *minUnix {
		return
	}
	*(*string)(dPtr) = t.Format(format)
}

var (
	localZeroUnix = time.Date(1970, 1, 1, 0, 0, 0, 0, time.Local)
	localZeroTime = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	utcZeroUnix   = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	utcZeroTime   = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
)

func cvtStringTime(format string, minUnix *int64, as asTime, dPtr, sPtr unsafe.Pointer) {
	t, err := time.ParseInLocation(format, *(*string)(sPtr), time.Local)
	if err != nil {
		return
	}
	if minUnix != nil && t.Unix() < *minUnix {
		return
	}
	if t.Equal(localZeroUnix) {
		t = utcZeroUnix
	} else if t.Equal(localZeroTime) {
		t = utcZeroTime
	}
	*as(dPtr) = t
}

func getTimeFormat(tw timeWrapper, option *StructOption) string {
	if option != nil && len(option.TimeFormat) > 0 {
		return option.TimeFormat
	}
	if format, ok := tw.GetFormat(); ok {
		return format
	}
	return TimeFormat
}

func getMinUnixScene(option *StructOption) MinUnixSceneType {
	if option != nil {
		return option.MinUnixScene
	}
	return MinUnixScene
}

func getMinUnix(option *StructOption, scene MinUnixSceneType) *int64 {
	if getMinUnixScene(option)&scene != scene {
		return nil
	}
	if option != nil && option.MinUnix != nil {
		return option.MinUnix
	}
	return MinUnix
}

func newTimeConverter(typ *convertType) converter {
	// todo: time.Time和时间戳互转
	if typ.dstTyp.Kind() == reflect.String {
		for _, v := range TimeWrappers {
			if v.Is(typ.srcTyp) {
				return &timeConverter{format: getTimeFormat(v, typ.option), minUnix: getMinUnix(typ.option, MinUnixTimeString), as: v.As, cvtOp: cvtTimeString}
			}
		}
	} else if typ.srcTyp.Kind() == reflect.String {
		for _, v := range TimeWrappers {
			if v.Is(typ.dstTyp) {
				return &timeConverter{format: getTimeFormat(v, typ.option), minUnix: getMinUnix(typ.option, MinUnixStringTime), as: v.As, cvtOp: cvtStringTime}
			}
		}
	}
	return nil
}
