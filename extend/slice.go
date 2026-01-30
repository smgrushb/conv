// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0
// Inspired by coven (MIT License) by petersunbag

package convextend

import (
	"fmt"
	"github.com/smgrushb/conv/internal"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// EmptyStringSplit 切割空字符串
type EmptyStringSplit int64

const (
	DefaultSplit EmptyStringSplit = iota // 默认		"" => [""]
	EmptySplit                           // 空切片	"" => []
	NilSplit                             // nil切片	"" => nil
)

var _ internal.CustomConverterV2 = (*string2Strings)(nil)

type string2Strings struct {
	sep   string
	split EmptyStringSplit
}

func String2Strings() *string2Strings {
	const sep = ","
	return &string2Strings{sep: sep}
}

func (s *string2Strings) Sep(sep string) *string2Strings {
	s.sep = sep
	return s
}

func (s *string2Strings) SplitStrategy(split EmptyStringSplit) *string2Strings {
	s.split = split
	return s
}

func (s *string2Strings) Is(dstTyp, srcTyp reflect.Type) bool {
	return srcTyp.Kind() == reflect.String &&
		dstTyp.Kind() == reflect.Slice && dstTyp.Elem().Kind() == reflect.String
}

func (s *string2Strings) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) bool {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) bool {
		if str := *(*string)(sPtr); s.split == 0 || len(str) > 0 {
			*(*[]string)(dPtr) = strings.Split(str, s.sep)
			return true
		} else if s.split == EmptySplit && len(str) == 0 {
			*(*[]string)(dPtr) = make([]string, 0)
			return true
		}
		return false
	}
}

func (s *string2Strings) Key() string {
	return fmt.Sprintf("[string2Strings::sep:%s,split:%d]", s.sep, s.split)
}

var _ internal.CustomConverterV2 = (*strings2String)(nil)

type strings2String struct {
	sep string
}

func Strings2String() *strings2String {
	const sep = ","
	return &strings2String{sep: sep}
}

func (s *strings2String) Sep(sep string) *strings2String {
	s.sep = sep
	return s
}

func (s *strings2String) Is(dstTyp, srcTyp reflect.Type) bool {
	return dstTyp.Kind() == reflect.String &&
		srcTyp.Kind() == reflect.Slice && srcTyp.Elem().Kind() == reflect.String
}

func (s *strings2String) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) bool {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) bool {
		*(*string)(dPtr) = strings.Join(*(*[]string)(sPtr), s.sep)
		return true
	}
}

func (s *strings2String) Key() string {
	return fmt.Sprintf("[strings2String::sep:%s]", s.sep)
}

var _ internal.CustomConverterV2 = (*string2Int64s)(nil)

type string2Int64s struct {
	sep string
}

func String2Int64s() *string2Int64s {
	const sep = ","
	return &string2Int64s{sep: sep}
}

func (s *string2Int64s) Sep(sep string) *string2Int64s {
	s.sep = sep
	return s
}

func (s *string2Int64s) Is(dstTyp, srcTyp reflect.Type) bool {
	return srcTyp.Kind() == reflect.String &&
		dstTyp.Kind() == reflect.Slice && dstTyp.Elem().Kind() == reflect.Int64
}

func (s *string2Int64s) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) bool {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) bool {
		strSlice := strings.Split(*(*string)(sPtr), s.sep)
		int64Slice := make([]int64, 0, len(strSlice))
		for _, v := range strSlice {
			if i64, err := strconv.ParseInt(v, 10, 64); err == nil {
				int64Slice = append(int64Slice, i64)
			}
		}
		*(*[]int64)(dPtr) = int64Slice
		return true
	}
}

func (s *string2Int64s) Key() string {
	return fmt.Sprintf("[string2Int64s::sep:%s]", s.sep)
}

var _ internal.CustomConverterV2 = (*strings2String)(nil)

type int64s2String struct {
	sep string
}

func Int64s2String() *int64s2String {
	const sep = ","
	return &int64s2String{sep: sep}
}

func (s *int64s2String) Sep(sep string) *int64s2String {
	s.sep = sep
	return s
}

func (s *int64s2String) Is(dstTyp, srcTyp reflect.Type) bool {
	return dstTyp.Kind() == reflect.String &&
		srcTyp.Kind() == reflect.Slice && srcTyp.Elem().Kind() == reflect.Int64
}

func (s *int64s2String) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) bool {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) bool {
		int64Slice := *(*[]int64)(sPtr)
		strSlice := make([]string, len(int64Slice))
		for i, v := range int64Slice {
			strSlice[i] = strconv.FormatInt(v, 10)
		}
		*(*string)(dPtr) = strings.Join(strSlice, s.sep)
		return true
	}
}

func (s *int64s2String) Key() string {
	return fmt.Sprintf("[int64s2String::sep:%s]", s.sep)
}

var _ internal.CustomConverterV2 = (*string2Ints)(nil)

type string2Ints struct {
	sep string
}

func String2Ints() *string2Ints {
	const sep = ","
	return &string2Ints{sep: sep}
}

func (s *string2Ints) Sep(sep string) *string2Ints {
	s.sep = sep
	return s
}

func (s *string2Ints) Is(dstTyp, srcTyp reflect.Type) bool {
	return srcTyp.Kind() == reflect.String &&
		dstTyp.Kind() == reflect.Slice && dstTyp.Elem().Kind() == reflect.Int
}

func (s *string2Ints) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) bool {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) bool {
		strSlice := strings.Split(*(*string)(sPtr), s.sep)
		intSlice := make([]int, 0, len(strSlice))
		for _, v := range strSlice {
			if i, err := strconv.Atoi(v); err == nil {
				intSlice = append(intSlice, i)
			}
		}
		*(*[]int)(dPtr) = intSlice
		return true
	}
}

func (s *string2Ints) Key() string {
	return fmt.Sprintf("[string2Ints::sep:%s]", s.sep)
}

var _ internal.CustomConverterV2 = (*strings2String)(nil)

type ints2String struct {
	sep string
}

func Ints2String() *ints2String {
	const sep = ","
	return &ints2String{sep: sep}
}

func (s *ints2String) Sep(sep string) *ints2String {
	s.sep = sep
	return s
}

func (s *ints2String) Is(dstTyp, srcTyp reflect.Type) bool {
	return dstTyp.Kind() == reflect.String &&
		srcTyp.Kind() == reflect.Slice && srcTyp.Elem().Kind() == reflect.Int
}

func (s *ints2String) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) bool {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) bool {
		intSlice := *(*[]int)(sPtr)
		strSlice := make([]string, len(intSlice))
		for i, v := range intSlice {
			strSlice[i] = strconv.Itoa(v)
		}
		*(*string)(dPtr) = strings.Join(strSlice, s.sep)
		return true
	}
}

func (s *ints2String) Key() string {
	return fmt.Sprintf("[ints2String::sep:%s]", s.sep)
}
