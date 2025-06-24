// This file is part of the original coven project:
// https://github.com/petersunbag/coven
//
// Copyright © 2018 petersunbag
// Licensed under the MIT License
// https://opensource.org/licenses/MIT
// Modified by smgrushb in 2025

package ptr

import (
	"github.com/smgrushb/conv/internal/generics/gptr"
	"github.com/smgrushb/conv/internal/unsafeheader"
	"reflect"
	"unsafe"
)

var (
	intAlign  = unsafe.Alignof(int(1))
	cvtOps    = make(map[convertKind]CvtOp)  // 基础类型互转
	anyCvtOps = make(map[reflect.Kind]CvtOp) // any转基础类型
	AnyType   = reflect.TypeOf((*any)(nil)).Elem()
)

type convertKind struct {
	srcTyp reflect.Kind
	dstTyp reflect.Kind
}

type CvtOp func(unsafe.Pointer, unsafe.Pointer)

func GetCvtOp(st, dt reflect.Type, strBytesZeroCopy bool) CvtOp {
	sk := st.Kind()
	dk := dt.Kind()
	if st == AnyType {
		if op := anyCvtOps[dk]; op != nil {
			return op
		}
	}
	if op := cvtOps[convertKind{sk, dk}]; op != nil {
		return op
	}
	switch sk {
	case reflect.Slice:
		if dk == reflect.String && st.Elem().PkgPath() == "" {
			switch st.Elem().Kind() {
			case reflect.Uint8:
				if strBytesZeroCopy {
					return cvtBytesStringZeroCopy
				}
				return cvtBytesString
			case reflect.Int32:
				return cvtRunesString
			}
		}
	case reflect.String:
		if dk == reflect.Slice && dt.Elem().PkgPath() == "" {
			switch dt.Elem().Kind() {
			case reflect.Uint8:
				if strBytesZeroCopy {
					return cvtStringBytesZeroCopy
				}
				return cvtStringBytes
			case reflect.Int32:
				return cvtStringRunes
			}
		}
	}
	return nil
}

func cvtRunesString(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = (string)(*(*[]rune)(sPtr))
}

func cvtBytesString(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = (string)(*(*[]byte)(sPtr))
}

func cvtBytesStringZeroCopy(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	byteSlice := (*unsafeheader.SliceHeader)(sPtr)
	strHeader := (*unsafeheader.StringHeader)(dPtr)
	strHeader.Data = byteSlice.Data
	strHeader.Len = byteSlice.Len
}

func cvtStringRunes(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*[]rune)(dPtr) = ([]rune)(*(*string)(sPtr))
}

func cvtStringBytes(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	*(*[]byte)(dPtr) = ([]byte)(*(*string)(sPtr))
}

func cvtStringBytesZeroCopy(sPtr unsafe.Pointer, dPtr unsafe.Pointer) {
	strHeader := (*unsafeheader.StringHeader)(sPtr)
	byteSlice := (*unsafeheader.SliceHeader)(dPtr)
	byteSlice.Data = strHeader.Data
	byteSlice.Len = strHeader.Len
	byteSlice.Cap = strHeader.Len
}

func Copy(dPtr, sPtr unsafe.Pointer, size uintptr) {
	align := uintptr(0)
	if size >= intAlign {
		for ; align+intAlign <= size; align += intAlign {
			*(*int)(unsafe.Pointer(uintptr(dPtr) + align)) = *(*int)(unsafe.Pointer(uintptr(sPtr) + align))
		}
	}
	for ; align < size; align++ {
		*(*byte)(unsafe.Pointer(uintptr(dPtr) + align)) = *(*byte)(unsafe.Pointer(uintptr(sPtr) + align))
	}
}

func NewValuePtr(k reflect.Kind) unsafe.Pointer {
	switch k {
	case reflect.Bool:
		return newPtr[bool]()
	case reflect.Int:
		return newPtr[int]()
	case reflect.Uint:
		return newPtr[uint]()
	case reflect.Int8:
		return newPtr[int8]()
	case reflect.Uint8:
		return newPtr[uint8]()
	case reflect.Int16:
		return newPtr[int16]()
	case reflect.Uint16:
		return newPtr[uint16]()
	case reflect.Int32:
		return newPtr[int32]()
	case reflect.Uint32:
		return newPtr[uint32]()
	case reflect.Int64:
		return newPtr[int64]()
	case reflect.Uint64:
		return newPtr[uint64]()
	case reflect.Float32:
		return newPtr[float32]()
	case reflect.Float64:
		return newPtr[float64]()
	case reflect.Complex64:
		return newPtr[complex64]()
	case reflect.Complex128:
		return newPtr[complex128]()
	case reflect.Uintptr:
		return newPtr[uintptr]()
	case reflect.String:
		return newPtr[string]()
	default:
		return nil
	}
}

func newPtr[T any]() unsafe.Pointer {
	return unsafe.Pointer(gptr.Zero[T]())
}
