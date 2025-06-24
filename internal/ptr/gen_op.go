// This file is part of the original coven project:
// https://github.com/petersunbag/coven
//
// Copyright © 2018 petersunbag
// Licensed under the MIT License
// https://opensource.org/licenses/MIT
// Modified by smgrushb in 2025

package ptr

import (
	"fmt"
	"github.com/smgrushb/conv/internal/generics/gvalue"
	"reflect"
	"strconv"
	"unsafe"
)

func init() {
	cvtOps[convertKind{reflect.Bool, reflect.Bool}] = cvtBoolBool
	cvtOps[convertKind{reflect.Bool, reflect.String}] = cvtBoolString
	cvtOps[convertKind{reflect.Bool, reflect.Int}] = cvtBoolInt
	cvtOps[convertKind{reflect.Bool, reflect.Uint}] = cvtBoolUint
	cvtOps[convertKind{reflect.Bool, reflect.Int8}] = cvtBoolInt8
	cvtOps[convertKind{reflect.Bool, reflect.Uint8}] = cvtBoolUint8
	cvtOps[convertKind{reflect.Bool, reflect.Int16}] = cvtBoolInt16
	cvtOps[convertKind{reflect.Bool, reflect.Uint16}] = cvtBoolUint16
	cvtOps[convertKind{reflect.Bool, reflect.Int32}] = cvtBoolInt32
	cvtOps[convertKind{reflect.Bool, reflect.Uint32}] = cvtBoolUint32
	cvtOps[convertKind{reflect.Bool, reflect.Int64}] = cvtBoolInt64
	cvtOps[convertKind{reflect.Bool, reflect.Uint64}] = cvtBoolUint64
	cvtOps[convertKind{reflect.Bool, reflect.Float32}] = cvtBoolFloat32
	cvtOps[convertKind{reflect.Bool, reflect.Float64}] = cvtBoolFloat64
	cvtOps[convertKind{reflect.Int, reflect.Int}] = cvtIntInt
	cvtOps[convertKind{reflect.Int, reflect.Uint}] = cvtIntUint
	cvtOps[convertKind{reflect.Int, reflect.Int8}] = cvtIntInt8
	cvtOps[convertKind{reflect.Int, reflect.Uint8}] = cvtIntUint8
	cvtOps[convertKind{reflect.Int, reflect.Int16}] = cvtIntInt16
	cvtOps[convertKind{reflect.Int, reflect.Uint16}] = cvtIntUint16
	cvtOps[convertKind{reflect.Int, reflect.Int32}] = cvtIntInt32
	cvtOps[convertKind{reflect.Int, reflect.Uint32}] = cvtIntUint32
	cvtOps[convertKind{reflect.Int, reflect.Int64}] = cvtIntInt64
	cvtOps[convertKind{reflect.Int, reflect.Uint64}] = cvtIntUint64
	cvtOps[convertKind{reflect.Int, reflect.Float32}] = cvtIntFloat32
	cvtOps[convertKind{reflect.Int, reflect.Float64}] = cvtIntFloat64
	cvtOps[convertKind{reflect.Int, reflect.Uintptr}] = cvtIntUintptr
	cvtOps[convertKind{reflect.Int, reflect.String}] = cvtIntString
	cvtOps[convertKind{reflect.Int, reflect.Bool}] = cvtIntBool
	cvtOps[convertKind{reflect.Uint, reflect.Int}] = cvtUintInt
	cvtOps[convertKind{reflect.Uint, reflect.Uint}] = cvtUintUint
	cvtOps[convertKind{reflect.Uint, reflect.Int8}] = cvtUintInt8
	cvtOps[convertKind{reflect.Uint, reflect.Uint8}] = cvtUintUint8
	cvtOps[convertKind{reflect.Uint, reflect.Int16}] = cvtUintInt16
	cvtOps[convertKind{reflect.Uint, reflect.Uint16}] = cvtUintUint16
	cvtOps[convertKind{reflect.Uint, reflect.Int32}] = cvtUintInt32
	cvtOps[convertKind{reflect.Uint, reflect.Uint32}] = cvtUintUint32
	cvtOps[convertKind{reflect.Uint, reflect.Int64}] = cvtUintInt64
	cvtOps[convertKind{reflect.Uint, reflect.Uint64}] = cvtUintUint64
	cvtOps[convertKind{reflect.Uint, reflect.Float32}] = cvtUintFloat32
	cvtOps[convertKind{reflect.Uint, reflect.Float64}] = cvtUintFloat64
	cvtOps[convertKind{reflect.Uint, reflect.Uintptr}] = cvtUintUintptr
	cvtOps[convertKind{reflect.Uint, reflect.String}] = cvtUintString
	cvtOps[convertKind{reflect.Uint, reflect.Bool}] = cvtUintBool
	cvtOps[convertKind{reflect.Int8, reflect.Int}] = cvtInt8Int
	cvtOps[convertKind{reflect.Int8, reflect.Uint}] = cvtInt8Uint
	cvtOps[convertKind{reflect.Int8, reflect.Int8}] = cvtInt8Int8
	cvtOps[convertKind{reflect.Int8, reflect.Uint8}] = cvtInt8Uint8
	cvtOps[convertKind{reflect.Int8, reflect.Int16}] = cvtInt8Int16
	cvtOps[convertKind{reflect.Int8, reflect.Uint16}] = cvtInt8Uint16
	cvtOps[convertKind{reflect.Int8, reflect.Int32}] = cvtInt8Int32
	cvtOps[convertKind{reflect.Int8, reflect.Uint32}] = cvtInt8Uint32
	cvtOps[convertKind{reflect.Int8, reflect.Int64}] = cvtInt8Int64
	cvtOps[convertKind{reflect.Int8, reflect.Uint64}] = cvtInt8Uint64
	cvtOps[convertKind{reflect.Int8, reflect.Float32}] = cvtInt8Float32
	cvtOps[convertKind{reflect.Int8, reflect.Float64}] = cvtInt8Float64
	cvtOps[convertKind{reflect.Int8, reflect.Uintptr}] = cvtInt8Uintptr
	cvtOps[convertKind{reflect.Int8, reflect.String}] = cvtInt8String
	cvtOps[convertKind{reflect.Int8, reflect.Bool}] = cvtInt8Bool
	cvtOps[convertKind{reflect.Uint8, reflect.Int}] = cvtUint8Int
	cvtOps[convertKind{reflect.Uint8, reflect.Uint}] = cvtUint8Uint
	cvtOps[convertKind{reflect.Uint8, reflect.Int8}] = cvtUint8Int8
	cvtOps[convertKind{reflect.Uint8, reflect.Uint8}] = cvtUint8Uint8
	cvtOps[convertKind{reflect.Uint8, reflect.Int16}] = cvtUint8Int16
	cvtOps[convertKind{reflect.Uint8, reflect.Uint16}] = cvtUint8Uint16
	cvtOps[convertKind{reflect.Uint8, reflect.Int32}] = cvtUint8Int32
	cvtOps[convertKind{reflect.Uint8, reflect.Uint32}] = cvtUint8Uint32
	cvtOps[convertKind{reflect.Uint8, reflect.Int64}] = cvtUint8Int64
	cvtOps[convertKind{reflect.Uint8, reflect.Uint64}] = cvtUint8Uint64
	cvtOps[convertKind{reflect.Uint8, reflect.Float32}] = cvtUint8Float32
	cvtOps[convertKind{reflect.Uint8, reflect.Float64}] = cvtUint8Float64
	cvtOps[convertKind{reflect.Uint8, reflect.Uintptr}] = cvtUint8Uintptr
	cvtOps[convertKind{reflect.Uint8, reflect.String}] = cvtUint8String
	cvtOps[convertKind{reflect.Uint8, reflect.Bool}] = cvtUint8Bool
	cvtOps[convertKind{reflect.Int16, reflect.Int}] = cvtInt16Int
	cvtOps[convertKind{reflect.Int16, reflect.Uint}] = cvtInt16Uint
	cvtOps[convertKind{reflect.Int16, reflect.Int8}] = cvtInt16Int8
	cvtOps[convertKind{reflect.Int16, reflect.Uint8}] = cvtInt16Uint8
	cvtOps[convertKind{reflect.Int16, reflect.Int16}] = cvtInt16Int16
	cvtOps[convertKind{reflect.Int16, reflect.Uint16}] = cvtInt16Uint16
	cvtOps[convertKind{reflect.Int16, reflect.Int32}] = cvtInt16Int32
	cvtOps[convertKind{reflect.Int16, reflect.Uint32}] = cvtInt16Uint32
	cvtOps[convertKind{reflect.Int16, reflect.Int64}] = cvtInt16Int64
	cvtOps[convertKind{reflect.Int16, reflect.Uint64}] = cvtInt16Uint64
	cvtOps[convertKind{reflect.Int16, reflect.Float32}] = cvtInt16Float32
	cvtOps[convertKind{reflect.Int16, reflect.Float64}] = cvtInt16Float64
	cvtOps[convertKind{reflect.Int16, reflect.Uintptr}] = cvtInt16Uintptr
	cvtOps[convertKind{reflect.Int16, reflect.String}] = cvtInt16String
	cvtOps[convertKind{reflect.Int16, reflect.Bool}] = cvtInt16Bool
	cvtOps[convertKind{reflect.Uint16, reflect.Int}] = cvtUint16Int
	cvtOps[convertKind{reflect.Uint16, reflect.Uint}] = cvtUint16Uint
	cvtOps[convertKind{reflect.Uint16, reflect.Int8}] = cvtUint16Int8
	cvtOps[convertKind{reflect.Uint16, reflect.Uint8}] = cvtUint16Uint8
	cvtOps[convertKind{reflect.Uint16, reflect.Int16}] = cvtUint16Int16
	cvtOps[convertKind{reflect.Uint16, reflect.Uint16}] = cvtUint16Uint16
	cvtOps[convertKind{reflect.Uint16, reflect.Int32}] = cvtUint16Int32
	cvtOps[convertKind{reflect.Uint16, reflect.Uint32}] = cvtUint16Uint32
	cvtOps[convertKind{reflect.Uint16, reflect.Int64}] = cvtUint16Int64
	cvtOps[convertKind{reflect.Uint16, reflect.Uint64}] = cvtUint16Uint64
	cvtOps[convertKind{reflect.Uint16, reflect.Float32}] = cvtUint16Float32
	cvtOps[convertKind{reflect.Uint16, reflect.Float64}] = cvtUint16Float64
	cvtOps[convertKind{reflect.Uint16, reflect.Uintptr}] = cvtUint16Uintptr
	cvtOps[convertKind{reflect.Uint16, reflect.String}] = cvtUint16String
	cvtOps[convertKind{reflect.Uint16, reflect.Bool}] = cvtUint16Bool
	cvtOps[convertKind{reflect.Int32, reflect.Int}] = cvtInt32Int
	cvtOps[convertKind{reflect.Int32, reflect.Uint}] = cvtInt32Uint
	cvtOps[convertKind{reflect.Int32, reflect.Int8}] = cvtInt32Int8
	cvtOps[convertKind{reflect.Int32, reflect.Uint8}] = cvtInt32Uint8
	cvtOps[convertKind{reflect.Int32, reflect.Int16}] = cvtInt32Int16
	cvtOps[convertKind{reflect.Int32, reflect.Uint16}] = cvtInt32Uint16
	cvtOps[convertKind{reflect.Int32, reflect.Int32}] = cvtInt32Int32
	cvtOps[convertKind{reflect.Int32, reflect.Uint32}] = cvtInt32Uint32
	cvtOps[convertKind{reflect.Int32, reflect.Int64}] = cvtInt32Int64
	cvtOps[convertKind{reflect.Int32, reflect.Uint64}] = cvtInt32Uint64
	cvtOps[convertKind{reflect.Int32, reflect.Float32}] = cvtInt32Float32
	cvtOps[convertKind{reflect.Int32, reflect.Float64}] = cvtInt32Float64
	cvtOps[convertKind{reflect.Int32, reflect.Uintptr}] = cvtInt32Uintptr
	cvtOps[convertKind{reflect.Int32, reflect.String}] = cvtInt32String
	cvtOps[convertKind{reflect.Int32, reflect.Bool}] = cvtInt32Bool
	cvtOps[convertKind{reflect.Uint32, reflect.Int}] = cvtUint32Int
	cvtOps[convertKind{reflect.Uint32, reflect.Uint}] = cvtUint32Uint
	cvtOps[convertKind{reflect.Uint32, reflect.Int8}] = cvtUint32Int8
	cvtOps[convertKind{reflect.Uint32, reflect.Uint8}] = cvtUint32Uint8
	cvtOps[convertKind{reflect.Uint32, reflect.Int16}] = cvtUint32Int16
	cvtOps[convertKind{reflect.Uint32, reflect.Uint16}] = cvtUint32Uint16
	cvtOps[convertKind{reflect.Uint32, reflect.Int32}] = cvtUint32Int32
	cvtOps[convertKind{reflect.Uint32, reflect.Uint32}] = cvtUint32Uint32
	cvtOps[convertKind{reflect.Uint32, reflect.Int64}] = cvtUint32Int64
	cvtOps[convertKind{reflect.Uint32, reflect.Uint64}] = cvtUint32Uint64
	cvtOps[convertKind{reflect.Uint32, reflect.Float32}] = cvtUint32Float32
	cvtOps[convertKind{reflect.Uint32, reflect.Float64}] = cvtUint32Float64
	cvtOps[convertKind{reflect.Uint32, reflect.Uintptr}] = cvtUint32Uintptr
	cvtOps[convertKind{reflect.Uint32, reflect.String}] = cvtUint32String
	cvtOps[convertKind{reflect.Uint32, reflect.Bool}] = cvtUint32Bool
	cvtOps[convertKind{reflect.Int64, reflect.Int}] = cvtInt64Int
	cvtOps[convertKind{reflect.Int64, reflect.Uint}] = cvtInt64Uint
	cvtOps[convertKind{reflect.Int64, reflect.Int8}] = cvtInt64Int8
	cvtOps[convertKind{reflect.Int64, reflect.Uint8}] = cvtInt64Uint8
	cvtOps[convertKind{reflect.Int64, reflect.Int16}] = cvtInt64Int16
	cvtOps[convertKind{reflect.Int64, reflect.Uint16}] = cvtInt64Uint16
	cvtOps[convertKind{reflect.Int64, reflect.Int32}] = cvtInt64Int32
	cvtOps[convertKind{reflect.Int64, reflect.Uint32}] = cvtInt64Uint32
	cvtOps[convertKind{reflect.Int64, reflect.Int64}] = cvtInt64Int64
	cvtOps[convertKind{reflect.Int64, reflect.Uint64}] = cvtInt64Uint64
	cvtOps[convertKind{reflect.Int64, reflect.Float32}] = cvtInt64Float32
	cvtOps[convertKind{reflect.Int64, reflect.Float64}] = cvtInt64Float64
	cvtOps[convertKind{reflect.Int64, reflect.Uintptr}] = cvtInt64Uintptr
	cvtOps[convertKind{reflect.Int64, reflect.String}] = cvtInt64String
	cvtOps[convertKind{reflect.Int64, reflect.Bool}] = cvtInt64Bool
	cvtOps[convertKind{reflect.Uint64, reflect.Int}] = cvtUint64Int
	cvtOps[convertKind{reflect.Uint64, reflect.Uint}] = cvtUint64Uint
	cvtOps[convertKind{reflect.Uint64, reflect.Int8}] = cvtUint64Int8
	cvtOps[convertKind{reflect.Uint64, reflect.Uint8}] = cvtUint64Uint8
	cvtOps[convertKind{reflect.Uint64, reflect.Int16}] = cvtUint64Int16
	cvtOps[convertKind{reflect.Uint64, reflect.Uint16}] = cvtUint64Uint16
	cvtOps[convertKind{reflect.Uint64, reflect.Int32}] = cvtUint64Int32
	cvtOps[convertKind{reflect.Uint64, reflect.Uint32}] = cvtUint64Uint32
	cvtOps[convertKind{reflect.Uint64, reflect.Int64}] = cvtUint64Int64
	cvtOps[convertKind{reflect.Uint64, reflect.Uint64}] = cvtUint64Uint64
	cvtOps[convertKind{reflect.Uint64, reflect.Float32}] = cvtUint64Float32
	cvtOps[convertKind{reflect.Uint64, reflect.Float64}] = cvtUint64Float64
	cvtOps[convertKind{reflect.Uint64, reflect.Uintptr}] = cvtUint64Uintptr
	cvtOps[convertKind{reflect.Uint64, reflect.String}] = cvtUint64String
	cvtOps[convertKind{reflect.Uint64, reflect.Bool}] = cvtUint64Bool
	cvtOps[convertKind{reflect.Float32, reflect.Int}] = cvtFloat32Int
	cvtOps[convertKind{reflect.Float32, reflect.Uint}] = cvtFloat32Uint
	cvtOps[convertKind{reflect.Float32, reflect.Int8}] = cvtFloat32Int8
	cvtOps[convertKind{reflect.Float32, reflect.Uint8}] = cvtFloat32Uint8
	cvtOps[convertKind{reflect.Float32, reflect.Int16}] = cvtFloat32Int16
	cvtOps[convertKind{reflect.Float32, reflect.Uint16}] = cvtFloat32Uint16
	cvtOps[convertKind{reflect.Float32, reflect.Int32}] = cvtFloat32Int32
	cvtOps[convertKind{reflect.Float32, reflect.Uint32}] = cvtFloat32Uint32
	cvtOps[convertKind{reflect.Float32, reflect.Int64}] = cvtFloat32Int64
	cvtOps[convertKind{reflect.Float32, reflect.Uint64}] = cvtFloat32Uint64
	cvtOps[convertKind{reflect.Float32, reflect.Float32}] = cvtFloat32Float32
	cvtOps[convertKind{reflect.Float32, reflect.Float64}] = cvtFloat32Float64
	cvtOps[convertKind{reflect.Float32, reflect.Uintptr}] = cvtFloat32Uintptr
	cvtOps[convertKind{reflect.Float32, reflect.String}] = cvtFloat32String
	cvtOps[convertKind{reflect.Float32, reflect.Bool}] = cvtFloat32Bool
	cvtOps[convertKind{reflect.Float64, reflect.Int}] = cvtFloat64Int
	cvtOps[convertKind{reflect.Float64, reflect.Uint}] = cvtFloat64Uint
	cvtOps[convertKind{reflect.Float64, reflect.Int8}] = cvtFloat64Int8
	cvtOps[convertKind{reflect.Float64, reflect.Uint8}] = cvtFloat64Uint8
	cvtOps[convertKind{reflect.Float64, reflect.Int16}] = cvtFloat64Int16
	cvtOps[convertKind{reflect.Float64, reflect.Uint16}] = cvtFloat64Uint16
	cvtOps[convertKind{reflect.Float64, reflect.Int32}] = cvtFloat64Int32
	cvtOps[convertKind{reflect.Float64, reflect.Uint32}] = cvtFloat64Uint32
	cvtOps[convertKind{reflect.Float64, reflect.Int64}] = cvtFloat64Int64
	cvtOps[convertKind{reflect.Float64, reflect.Uint64}] = cvtFloat64Uint64
	cvtOps[convertKind{reflect.Float64, reflect.Float32}] = cvtFloat64Float32
	cvtOps[convertKind{reflect.Float64, reflect.Float64}] = cvtFloat64Float64
	cvtOps[convertKind{reflect.Float64, reflect.Uintptr}] = cvtFloat64Uintptr
	cvtOps[convertKind{reflect.Float64, reflect.String}] = cvtFloat64String
	cvtOps[convertKind{reflect.Float64, reflect.Bool}] = cvtFloat64Bool
	cvtOps[convertKind{reflect.Complex64, reflect.Complex64}] = cvtComplex64Complex64
	cvtOps[convertKind{reflect.Complex64, reflect.Complex128}] = cvtComplex64Complex128
	cvtOps[convertKind{reflect.Complex64, reflect.String}] = cvtComplex64String
	cvtOps[convertKind{reflect.Complex128, reflect.Complex64}] = cvtComplex128Complex64
	cvtOps[convertKind{reflect.Complex128, reflect.Complex128}] = cvtComplex128Complex128
	cvtOps[convertKind{reflect.Complex128, reflect.String}] = cvtComplex128String
	cvtOps[convertKind{reflect.Uintptr, reflect.Int}] = cvtUintptrInt
	cvtOps[convertKind{reflect.Uintptr, reflect.Uint}] = cvtUintptrUint
	cvtOps[convertKind{reflect.Uintptr, reflect.Int8}] = cvtUintptrInt8
	cvtOps[convertKind{reflect.Uintptr, reflect.Uint8}] = cvtUintptrUint8
	cvtOps[convertKind{reflect.Uintptr, reflect.Int16}] = cvtUintptrInt16
	cvtOps[convertKind{reflect.Uintptr, reflect.Uint16}] = cvtUintptrUint16
	cvtOps[convertKind{reflect.Uintptr, reflect.Int32}] = cvtUintptrInt32
	cvtOps[convertKind{reflect.Uintptr, reflect.Uint32}] = cvtUintptrUint32
	cvtOps[convertKind{reflect.Uintptr, reflect.Int64}] = cvtUintptrInt64
	cvtOps[convertKind{reflect.Uintptr, reflect.Uint64}] = cvtUintptrUint64
	cvtOps[convertKind{reflect.Uintptr, reflect.Float32}] = cvtUintptrFloat32
	cvtOps[convertKind{reflect.Uintptr, reflect.Float64}] = cvtUintptrFloat64
	cvtOps[convertKind{reflect.Uintptr, reflect.Uintptr}] = cvtUintptrUintptr
	cvtOps[convertKind{reflect.Uintptr, reflect.String}] = cvtUintptrString
	cvtOps[convertKind{reflect.String, reflect.Int}] = cvtStringInt
	cvtOps[convertKind{reflect.String, reflect.Uint}] = cvtStringUint
	cvtOps[convertKind{reflect.String, reflect.Int8}] = cvtStringInt8
	cvtOps[convertKind{reflect.String, reflect.Uint8}] = cvtStringUint8
	cvtOps[convertKind{reflect.String, reflect.Int16}] = cvtStringInt16
	cvtOps[convertKind{reflect.String, reflect.Uint16}] = cvtStringUint16
	cvtOps[convertKind{reflect.String, reflect.Int32}] = cvtStringInt32
	cvtOps[convertKind{reflect.String, reflect.Uint32}] = cvtStringUint32
	cvtOps[convertKind{reflect.String, reflect.Int64}] = cvtStringInt64
	cvtOps[convertKind{reflect.String, reflect.Uint64}] = cvtStringUint64
	cvtOps[convertKind{reflect.String, reflect.Float32}] = cvtStringFloat32
	cvtOps[convertKind{reflect.String, reflect.Float64}] = cvtStringFloat64
	cvtOps[convertKind{reflect.String, reflect.String}] = cvtStringString
	cvtOps[convertKind{reflect.String, reflect.Bool}] = cvtStringBool

	anyCvtOps[reflect.Bool] = cvtAnyBool
	anyCvtOps[reflect.Int] = cvtAnyInt
	anyCvtOps[reflect.Uint] = cvtAnyUint
	anyCvtOps[reflect.Int8] = cvtAnyInt8
	anyCvtOps[reflect.Uint8] = cvtAnyUint8
	anyCvtOps[reflect.Int16] = cvtAnyInt16
	anyCvtOps[reflect.Uint16] = cvtAnyUint16
	anyCvtOps[reflect.Int32] = cvtAnyInt32
	anyCvtOps[reflect.Uint32] = cvtAnyUint32
	anyCvtOps[reflect.Int64] = cvtAnyInt64
	anyCvtOps[reflect.Uint64] = cvtAnyUint64
	anyCvtOps[reflect.Float32] = cvtAnyFloat32
	anyCvtOps[reflect.Float64] = cvtAnyFloat64
	anyCvtOps[reflect.String] = cvtAnyString
}

func cvtBoolBool(sPtr, dPtr unsafe.Pointer) {
	*(*bool)(dPtr) = (bool)(*(*bool)(sPtr))
}

func cvtBoolString(sPtr, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = fmt.Sprintf("%v", *(*bool)(sPtr))
}

func cvtBoolInt(sPtr, dPtr unsafe.Pointer) {
	if *(*bool)(sPtr) {
		*(*int)(dPtr) = 1
	}
}

func cvtBoolUint(sPtr, dPtr unsafe.Pointer) {
	if *(*bool)(sPtr) {
		*(*uint)(dPtr) = 1
	}
}

func cvtBoolInt8(sPtr, dPtr unsafe.Pointer) {
	if *(*bool)(sPtr) {
		*(*int8)(dPtr) = 1
	}
}

func cvtBoolUint8(sPtr, dPtr unsafe.Pointer) {
	if *(*bool)(sPtr) {
		*(*uint8)(dPtr) = 1
	}
}

func cvtBoolInt16(sPtr, dPtr unsafe.Pointer) {
	if *(*bool)(sPtr) {
		*(*int16)(dPtr) = 1
	}
}

func cvtBoolUint16(sPtr, dPtr unsafe.Pointer) {
	if *(*bool)(sPtr) {
		*(*uint16)(dPtr) = 1
	}
}

func cvtBoolInt32(sPtr, dPtr unsafe.Pointer) {
	if *(*bool)(sPtr) {
		*(*int32)(dPtr) = 1
	}
}

func cvtBoolUint32(sPtr, dPtr unsafe.Pointer) {
	if *(*bool)(sPtr) {
		*(*uint32)(dPtr) = 1
	}
}

func cvtBoolInt64(sPtr, dPtr unsafe.Pointer) {
	if *(*bool)(sPtr) {
		*(*int64)(dPtr) = 1
	}
}

func cvtBoolUint64(sPtr, dPtr unsafe.Pointer) {
	if *(*bool)(sPtr) {
		*(*uint64)(dPtr) = 1
	}
}

func cvtBoolFloat32(sPtr, dPtr unsafe.Pointer) {
	if *(*bool)(sPtr) {
		*(*float32)(dPtr) = 1
	}
}

func cvtBoolFloat64(sPtr, dPtr unsafe.Pointer) {
	if *(*bool)(sPtr) {
		*(*float64)(dPtr) = 1
	}
}

func cvtIntInt(sPtr, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*int)(sPtr))
}

func cvtIntUint(sPtr, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*int)(sPtr))
}

func cvtIntInt8(sPtr, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*int)(sPtr))
}

func cvtIntUint8(sPtr, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*int)(sPtr))
}

func cvtIntInt16(sPtr, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*int)(sPtr))
}

func cvtIntUint16(sPtr, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*int)(sPtr))
}

func cvtIntInt32(sPtr, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*int)(sPtr))
}

func cvtIntUint32(sPtr, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*int)(sPtr))
}

func cvtIntInt64(sPtr, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*int)(sPtr))
}

func cvtIntUint64(sPtr, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*int)(sPtr))
}

func cvtIntFloat32(sPtr, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*int)(sPtr))
}

func cvtIntFloat64(sPtr, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*int)(sPtr))
}

func cvtIntUintptr(sPtr, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*int)(sPtr))
}

func cvtIntString(sPtr, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = fmt.Sprintf("%v", *(*int)(sPtr))
}

func cvtIntBool(sPtr, dPtr unsafe.Pointer) {
	*(*bool)(dPtr) = *(*int)(sPtr) != 0
}

func cvtUintInt(sPtr, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*uint)(sPtr))
}

func cvtUintUint(sPtr, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*uint)(sPtr))
}

func cvtUintInt8(sPtr, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*uint)(sPtr))
}

func cvtUintUint8(sPtr, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*uint)(sPtr))
}

func cvtUintInt16(sPtr, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*uint)(sPtr))
}

func cvtUintUint16(sPtr, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*uint)(sPtr))
}

func cvtUintInt32(sPtr, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*uint)(sPtr))
}

func cvtUintUint32(sPtr, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*uint)(sPtr))
}

func cvtUintInt64(sPtr, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*uint)(sPtr))
}

func cvtUintUint64(sPtr, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*uint)(sPtr))
}

func cvtUintFloat32(sPtr, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*uint)(sPtr))
}

func cvtUintFloat64(sPtr, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*uint)(sPtr))
}

func cvtUintUintptr(sPtr, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*uint)(sPtr))
}

func cvtUintString(sPtr, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = fmt.Sprintf("%v", *(*uint)(sPtr))
}

func cvtUintBool(sPtr, dPtr unsafe.Pointer) {
	*(*bool)(dPtr) = *(*uint)(sPtr) != 0
}

func cvtInt8Int(sPtr, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*int8)(sPtr))
}

func cvtInt8Uint(sPtr, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*int8)(sPtr))
}

func cvtInt8Int8(sPtr, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*int8)(sPtr))
}

func cvtInt8Uint8(sPtr, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*int8)(sPtr))
}

func cvtInt8Int16(sPtr, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*int8)(sPtr))
}

func cvtInt8Uint16(sPtr, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*int8)(sPtr))
}

func cvtInt8Int32(sPtr, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*int8)(sPtr))
}

func cvtInt8Uint32(sPtr, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*int8)(sPtr))
}

func cvtInt8Int64(sPtr, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*int8)(sPtr))
}

func cvtInt8Uint64(sPtr, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*int8)(sPtr))
}

func cvtInt8Float32(sPtr, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*int8)(sPtr))
}

func cvtInt8Float64(sPtr, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*int8)(sPtr))
}

func cvtInt8Uintptr(sPtr, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*int8)(sPtr))
}

func cvtInt8String(sPtr, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = fmt.Sprintf("%v", *(*int8)(sPtr))
}

func cvtInt8Bool(sPtr, dPtr unsafe.Pointer) {
	*(*bool)(dPtr) = *(*int8)(sPtr) != 0
}

func cvtUint8Int(sPtr, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*uint8)(sPtr))
}

func cvtUint8Uint(sPtr, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*uint8)(sPtr))
}

func cvtUint8Int8(sPtr, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*uint8)(sPtr))
}

func cvtUint8Uint8(sPtr, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*uint8)(sPtr))
}

func cvtUint8Int16(sPtr, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*uint8)(sPtr))
}

func cvtUint8Uint16(sPtr, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*uint8)(sPtr))
}

func cvtUint8Int32(sPtr, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*uint8)(sPtr))
}

func cvtUint8Uint32(sPtr, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*uint8)(sPtr))
}

func cvtUint8Int64(sPtr, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*uint8)(sPtr))
}

func cvtUint8Uint64(sPtr, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*uint8)(sPtr))
}

func cvtUint8Float32(sPtr, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*uint8)(sPtr))
}

func cvtUint8Float64(sPtr, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*uint8)(sPtr))
}

func cvtUint8Uintptr(sPtr, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*uint8)(sPtr))
}

func cvtUint8String(sPtr, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = fmt.Sprintf("%v", *(*uint8)(sPtr))
}

func cvtUint8Bool(sPtr, dPtr unsafe.Pointer) {
	*(*bool)(dPtr) = *(*uint8)(sPtr) != 0
}

func cvtInt16Int(sPtr, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*int16)(sPtr))
}

func cvtInt16Uint(sPtr, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*int16)(sPtr))
}

func cvtInt16Int8(sPtr, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*int16)(sPtr))
}

func cvtInt16Uint8(sPtr, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*int16)(sPtr))
}

func cvtInt16Int16(sPtr, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*int16)(sPtr))
}

func cvtInt16Uint16(sPtr, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*int16)(sPtr))
}

func cvtInt16Int32(sPtr, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*int16)(sPtr))
}

func cvtInt16Uint32(sPtr, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*int16)(sPtr))
}

func cvtInt16Int64(sPtr, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*int16)(sPtr))
}

func cvtInt16Uint64(sPtr, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*int16)(sPtr))
}

func cvtInt16Float32(sPtr, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*int16)(sPtr))
}

func cvtInt16Float64(sPtr, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*int16)(sPtr))
}

func cvtInt16Uintptr(sPtr, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*int16)(sPtr))
}

func cvtInt16String(sPtr, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = fmt.Sprintf("%v", *(*int16)(sPtr))
}

func cvtInt16Bool(sPtr, dPtr unsafe.Pointer) {
	*(*bool)(dPtr) = *(*int16)(sPtr) != 0
}

func cvtUint16Int(sPtr, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*uint16)(sPtr))
}

func cvtUint16Uint(sPtr, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*uint16)(sPtr))
}

func cvtUint16Int8(sPtr, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*uint16)(sPtr))
}

func cvtUint16Uint8(sPtr, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*uint16)(sPtr))
}

func cvtUint16Int16(sPtr, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*uint16)(sPtr))
}

func cvtUint16Uint16(sPtr, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*uint16)(sPtr))
}

func cvtUint16Int32(sPtr, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*uint16)(sPtr))
}

func cvtUint16Uint32(sPtr, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*uint16)(sPtr))
}

func cvtUint16Int64(sPtr, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*uint16)(sPtr))
}

func cvtUint16Uint64(sPtr, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*uint16)(sPtr))
}

func cvtUint16Float32(sPtr, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*uint16)(sPtr))
}

func cvtUint16Float64(sPtr, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*uint16)(sPtr))
}

func cvtUint16Uintptr(sPtr, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*uint16)(sPtr))
}

func cvtUint16String(sPtr, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = fmt.Sprintf("%v", *(*uint16)(sPtr))
}

func cvtUint16Bool(sPtr, dPtr unsafe.Pointer) {
	*(*bool)(dPtr) = *(*uint16)(sPtr) != 0
}

func cvtInt32Int(sPtr, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*int32)(sPtr))
}

func cvtInt32Uint(sPtr, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*int32)(sPtr))
}

func cvtInt32Int8(sPtr, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*int32)(sPtr))
}

func cvtInt32Uint8(sPtr, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*int32)(sPtr))
}

func cvtInt32Int16(sPtr, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*int32)(sPtr))
}

func cvtInt32Uint16(sPtr, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*int32)(sPtr))
}

func cvtInt32Int32(sPtr, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*int32)(sPtr))
}

func cvtInt32Uint32(sPtr, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*int32)(sPtr))
}

func cvtInt32Int64(sPtr, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*int32)(sPtr))
}

func cvtInt32Uint64(sPtr, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*int32)(sPtr))
}

func cvtInt32Float32(sPtr, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*int32)(sPtr))
}

func cvtInt32Float64(sPtr, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*int32)(sPtr))
}

func cvtInt32Uintptr(sPtr, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*int32)(sPtr))
}

func cvtInt32String(sPtr, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = fmt.Sprintf("%v", *(*int32)(sPtr))
}

func cvtInt32Bool(sPtr, dPtr unsafe.Pointer) {
	*(*bool)(dPtr) = *(*int32)(sPtr) != 0
}

func cvtUint32Int(sPtr, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*uint32)(sPtr))
}

func cvtUint32Uint(sPtr, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*uint32)(sPtr))
}

func cvtUint32Int8(sPtr, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*uint32)(sPtr))
}

func cvtUint32Uint8(sPtr, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*uint32)(sPtr))
}

func cvtUint32Int16(sPtr, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*uint32)(sPtr))
}

func cvtUint32Uint16(sPtr, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*uint32)(sPtr))
}

func cvtUint32Int32(sPtr, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*uint32)(sPtr))
}

func cvtUint32Uint32(sPtr, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*uint32)(sPtr))
}

func cvtUint32Int64(sPtr, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*uint32)(sPtr))
}

func cvtUint32Uint64(sPtr, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*uint32)(sPtr))
}

func cvtUint32Float32(sPtr, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*uint32)(sPtr))
}

func cvtUint32Float64(sPtr, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*uint32)(sPtr))
}

func cvtUint32Uintptr(sPtr, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*uint32)(sPtr))
}

func cvtUint32String(sPtr, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = fmt.Sprintf("%v", *(*uint32)(sPtr))
}

func cvtUint32Bool(sPtr, dPtr unsafe.Pointer) {
	*(*bool)(dPtr) = *(*uint32)(sPtr) != 0
}

func cvtInt64Int(sPtr, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*int64)(sPtr))
}

func cvtInt64Uint(sPtr, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*int64)(sPtr))
}

func cvtInt64Int8(sPtr, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*int64)(sPtr))
}

func cvtInt64Uint8(sPtr, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*int64)(sPtr))
}

func cvtInt64Int16(sPtr, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*int64)(sPtr))
}

func cvtInt64Uint16(sPtr, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*int64)(sPtr))
}

func cvtInt64Int32(sPtr, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*int64)(sPtr))
}

func cvtInt64Uint32(sPtr, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*int64)(sPtr))
}

func cvtInt64Int64(sPtr, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*int64)(sPtr))
}

func cvtInt64Uint64(sPtr, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*int64)(sPtr))
}

func cvtInt64Float32(sPtr, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*int64)(sPtr))
}

func cvtInt64Float64(sPtr, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*int64)(sPtr))
}

func cvtInt64Uintptr(sPtr, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*int64)(sPtr))
}

func cvtInt64String(sPtr, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = fmt.Sprintf("%v", *(*int64)(sPtr))
}

func cvtInt64Bool(sPtr, dPtr unsafe.Pointer) {
	*(*bool)(dPtr) = *(*int64)(sPtr) != 0
}

func cvtUint64Int(sPtr, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*uint64)(sPtr))
}

func cvtUint64Uint(sPtr, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*uint64)(sPtr))
}

func cvtUint64Int8(sPtr, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*uint64)(sPtr))
}

func cvtUint64Uint8(sPtr, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*uint64)(sPtr))
}

func cvtUint64Int16(sPtr, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*uint64)(sPtr))
}

func cvtUint64Uint16(sPtr, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*uint64)(sPtr))
}

func cvtUint64Int32(sPtr, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*uint64)(sPtr))
}

func cvtUint64Uint32(sPtr, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*uint64)(sPtr))
}

func cvtUint64Int64(sPtr, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*uint64)(sPtr))
}

func cvtUint64Uint64(sPtr, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*uint64)(sPtr))
}

func cvtUint64Float32(sPtr, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*uint64)(sPtr))
}

func cvtUint64Float64(sPtr, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*uint64)(sPtr))
}

func cvtUint64Uintptr(sPtr, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*uint64)(sPtr))
}

func cvtUint64String(sPtr, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = fmt.Sprintf("%v", *(*uint64)(sPtr))
}

func cvtUint64Bool(sPtr, dPtr unsafe.Pointer) {
	*(*bool)(dPtr) = *(*uint64)(sPtr) != 0
}

func cvtFloat32Int(sPtr, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*float32)(sPtr))
}

func cvtFloat32Uint(sPtr, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*float32)(sPtr))
}

func cvtFloat32Int8(sPtr, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*float32)(sPtr))
}

func cvtFloat32Uint8(sPtr, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*float32)(sPtr))
}

func cvtFloat32Int16(sPtr, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*float32)(sPtr))
}

func cvtFloat32Uint16(sPtr, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*float32)(sPtr))
}

func cvtFloat32Int32(sPtr, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*float32)(sPtr))
}

func cvtFloat32Uint32(sPtr, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*float32)(sPtr))
}

func cvtFloat32Int64(sPtr, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*float32)(sPtr))
}

func cvtFloat32Uint64(sPtr, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*float32)(sPtr))
}

func cvtFloat32Float32(sPtr, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*float32)(sPtr))
}

func cvtFloat32Float64(sPtr, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*float32)(sPtr))
}

func cvtFloat32Uintptr(sPtr, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*float32)(sPtr))
}

func cvtFloat32String(sPtr, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = fmt.Sprintf("%v", *(*float32)(sPtr))
}

func cvtFloat32Bool(sPtr, dPtr unsafe.Pointer) {
	*(*bool)(dPtr) = *(*float32)(sPtr) != 0
}

func cvtFloat64Int(sPtr, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*float64)(sPtr))
}

func cvtFloat64Uint(sPtr, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*float64)(sPtr))
}

func cvtFloat64Int8(sPtr, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*float64)(sPtr))
}

func cvtFloat64Uint8(sPtr, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*float64)(sPtr))
}

func cvtFloat64Int16(sPtr, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*float64)(sPtr))
}

func cvtFloat64Uint16(sPtr, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*float64)(sPtr))
}

func cvtFloat64Int32(sPtr, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*float64)(sPtr))
}

func cvtFloat64Uint32(sPtr, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*float64)(sPtr))
}

func cvtFloat64Int64(sPtr, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*float64)(sPtr))
}

func cvtFloat64Uint64(sPtr, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*float64)(sPtr))
}

func cvtFloat64Float32(sPtr, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*float64)(sPtr))
}

func cvtFloat64Float64(sPtr, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*float64)(sPtr))
}

func cvtFloat64Uintptr(sPtr, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*float64)(sPtr))
}

func cvtFloat64String(sPtr, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = fmt.Sprintf("%v", *(*float64)(sPtr))
}

func cvtFloat64Bool(sPtr, dPtr unsafe.Pointer) {
	*(*bool)(dPtr) = *(*float64)(sPtr) != 0
}

func cvtComplex64Complex64(sPtr, dPtr unsafe.Pointer) {
	*(*complex64)(dPtr) = (complex64)(*(*complex64)(sPtr))
}

func cvtComplex64Complex128(sPtr, dPtr unsafe.Pointer) {
	*(*complex128)(dPtr) = (complex128)(*(*complex64)(sPtr))
}

func cvtComplex64String(sPtr, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = fmt.Sprintf("%v", *(*complex64)(sPtr))
}

func cvtComplex128Complex64(sPtr, dPtr unsafe.Pointer) {
	*(*complex64)(dPtr) = (complex64)(*(*complex128)(sPtr))
}

func cvtComplex128Complex128(sPtr, dPtr unsafe.Pointer) {
	*(*complex128)(dPtr) = (complex128)(*(*complex128)(sPtr))
}

func cvtComplex128String(sPtr, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = fmt.Sprintf("%v", *(*complex128)(sPtr))
}

func cvtUintptrInt(sPtr, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = (int)(*(*uintptr)(sPtr))
}

func cvtUintptrUint(sPtr, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = (uint)(*(*uintptr)(sPtr))
}

func cvtUintptrInt8(sPtr, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = (int8)(*(*uintptr)(sPtr))
}

func cvtUintptrUint8(sPtr, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = (uint8)(*(*uintptr)(sPtr))
}

func cvtUintptrInt16(sPtr, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = (int16)(*(*uintptr)(sPtr))
}

func cvtUintptrUint16(sPtr, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = (uint16)(*(*uintptr)(sPtr))
}

func cvtUintptrInt32(sPtr, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = (int32)(*(*uintptr)(sPtr))
}

func cvtUintptrUint32(sPtr, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = (uint32)(*(*uintptr)(sPtr))
}

func cvtUintptrInt64(sPtr, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = (int64)(*(*uintptr)(sPtr))
}

func cvtUintptrUint64(sPtr, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = (uint64)(*(*uintptr)(sPtr))
}

func cvtUintptrFloat32(sPtr, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = (float32)(*(*uintptr)(sPtr))
}

func cvtUintptrFloat64(sPtr, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = (float64)(*(*uintptr)(sPtr))
}

func cvtUintptrUintptr(sPtr, dPtr unsafe.Pointer) {
	*(*uintptr)(dPtr) = (uintptr)(*(*uintptr)(sPtr))
}

func cvtUintptrString(sPtr, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = fmt.Sprintf("%v", *(*uintptr)(sPtr))
}

func cvtStringInt(sPtr, dPtr unsafe.Pointer) {
	i64, _ := strconv.ParseInt(*(*string)(sPtr), 10, 64)
	*(*int)(dPtr) = int(i64)
}

func cvtStringUint(sPtr, dPtr unsafe.Pointer) {
	ui64, _ := strconv.ParseUint(*(*string)(sPtr), 10, 64)
	*(*uint)(dPtr) = uint(ui64)
}

func cvtStringInt8(sPtr, dPtr unsafe.Pointer) {
	i64, _ := strconv.ParseInt(*(*string)(sPtr), 10, 8)
	*(*int8)(dPtr) = int8(i64)
}

func cvtStringUint8(sPtr, dPtr unsafe.Pointer) {
	ui64, _ := strconv.ParseUint(*(*string)(sPtr), 10, 8)
	*(*uint8)(dPtr) = uint8(ui64)
}

func cvtStringInt16(sPtr, dPtr unsafe.Pointer) {
	i64, _ := strconv.ParseInt(*(*string)(sPtr), 10, 16)
	*(*int16)(dPtr) = int16(i64)
}

func cvtStringUint16(sPtr, dPtr unsafe.Pointer) {
	ui64, _ := strconv.ParseUint(*(*string)(sPtr), 10, 16)
	*(*uint16)(dPtr) = uint16(ui64)
}

func cvtStringInt32(sPtr, dPtr unsafe.Pointer) {
	i64, _ := strconv.ParseInt(*(*string)(sPtr), 10, 32)
	*(*int32)(dPtr) = int32(i64)
}

func cvtStringUint32(sPtr, dPtr unsafe.Pointer) {
	ui64, _ := strconv.ParseUint(*(*string)(sPtr), 10, 32)
	*(*uint32)(dPtr) = uint32(ui64)
}

func cvtStringInt64(sPtr, dPtr unsafe.Pointer) {
	*(*int64)(dPtr), _ = strconv.ParseInt(*(*string)(sPtr), 10, 64)
}

func cvtStringUint64(sPtr, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr), _ = strconv.ParseUint(*(*string)(sPtr), 10, 64)
}

func cvtStringFloat32(sPtr, dPtr unsafe.Pointer) {
	f64, _ := strconv.ParseFloat(*(*string)(sPtr), 32)
	*(*float32)(dPtr) = float32(f64)
}

func cvtStringFloat64(sPtr, dPtr unsafe.Pointer) {
	*(*float64)(dPtr), _ = strconv.ParseFloat(*(*string)(sPtr), 64)
}

func cvtStringString(sPtr, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = (string)(*(*string)(sPtr))
}

func cvtStringBool(sPtr, dPtr unsafe.Pointer) {
	b, _ := strconv.ParseBool(*(*string)(sPtr))
	*(*bool)(dPtr) = bool(b)
}

func cvtAnyBool(sPtr, dPtr unsafe.Pointer) {
	*(*bool)(dPtr) = gvalue.AsInt64((any)(*(*any)(sPtr))) == 1
}

func cvtAnyInt(sPtr, dPtr unsafe.Pointer) {
	*(*int)(dPtr) = int(gvalue.AsInt64((any)(*(*any)(sPtr))))
}

func cvtAnyUint(sPtr, dPtr unsafe.Pointer) {
	*(*uint)(dPtr) = uint(gvalue.AsInt64((any)(*(*any)(sPtr))))
}

func cvtAnyInt8(sPtr, dPtr unsafe.Pointer) {
	*(*int8)(dPtr) = int8(gvalue.AsInt64((any)(*(*any)(sPtr))))
}

func cvtAnyUint8(sPtr, dPtr unsafe.Pointer) {
	*(*uint8)(dPtr) = uint8(gvalue.AsInt64((any)(*(*any)(sPtr))))
}

func cvtAnyInt16(sPtr, dPtr unsafe.Pointer) {
	*(*int16)(dPtr) = int16(gvalue.AsInt64((any)(*(*any)(sPtr))))
}

func cvtAnyUint16(sPtr, dPtr unsafe.Pointer) {
	*(*uint16)(dPtr) = uint16(gvalue.AsInt64((any)(*(*any)(sPtr))))
}

func cvtAnyInt32(sPtr, dPtr unsafe.Pointer) {
	*(*int32)(dPtr) = int32(gvalue.AsInt64((any)(*(*any)(sPtr))))
}

func cvtAnyUint32(sPtr, dPtr unsafe.Pointer) {
	*(*uint32)(dPtr) = uint32(gvalue.AsInt64((any)(*(*any)(sPtr))))
}

func cvtAnyInt64(sPtr, dPtr unsafe.Pointer) {
	*(*int64)(dPtr) = int64(gvalue.AsInt64((any)(*(*any)(sPtr))))
}

func cvtAnyUint64(sPtr, dPtr unsafe.Pointer) {
	*(*uint64)(dPtr) = uint64(gvalue.AsInt64((any)(*(*any)(sPtr))))
}

func cvtAnyFloat32(sPtr, dPtr unsafe.Pointer) {
	*(*float32)(dPtr) = float32(gvalue.AsFloat64((any)(*(*any)(sPtr))))
}

func cvtAnyFloat64(sPtr, dPtr unsafe.Pointer) {
	*(*float64)(dPtr) = float64(gvalue.AsFloat64((any)(*(*any)(sPtr))))
}

func cvtAnyString(sPtr, dPtr unsafe.Pointer) {
	*(*string)(dPtr) = string(gvalue.AsString((any)(*(*any)(sPtr))))
}
