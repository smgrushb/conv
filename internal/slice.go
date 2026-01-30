// This file is part of the original coven project:
// https://github.com/petersunbag/coven
//
// Copyright © 2018 petersunbag
// Licensed under the MIT License
// https://opensource.org/licenses/MIT
// Modified by smgrushb in 2025

package internal

import (
	"github.com/smgrushb/conv/internal/ptr"
	"github.com/smgrushb/conv/internal/unsafeheader"
	"reflect"
	"unsafe"
)

type sliceConverter struct {
	*convertType
	*elemConverter
	dElemSize uintptr
	sElemSize uintptr
	enable    bool // 兜底
}

func newSliceConverter(typ *convertType) converter {
	c := &sliceConverter{
		convertType: typ,
		dElemSize:   typ.dstTyp.Elem().Size(),
		sElemSize:   typ.srcTyp.Elem().Size(),
	}
	if c.enable = typ.srcTyp == typ.dstTyp; c.enable {
		return c
	}
	key := typ.key()
	// 先预注册进去，不然循环依赖下会循环解析
	createdConverters[key] = &Converter{convertType: typ, converter: c}
	if ec, ok := newElemConverter(typ.dstTyp.Elem(), typ.srcTyp.Elem(), typ.option); ok {
		c.elemConverter = ec
		c.enable = true
		return c
	}
	// 把预注册的内容删了
	delete(createdConverters, key)
	return nil
}

func (s *sliceConverter) convert(dPtr, sPtr unsafe.Pointer) bool {
	if !s.enable {
		return false
	}
	dSlice, sSlice := (*unsafeheader.SliceHeader)(dPtr), (*unsafeheader.SliceHeader)(sPtr)
	length := sSlice.Len
	dSlice.Len = length
	if dSlice.Cap < length || dSlice.Data == nil {
		dv := reflect.NewAt(s.dstTyp, dPtr).Elem()
		newVal := reflect.MakeSlice(s.dstTyp, length, length)
		dv.Set(newVal)
	}
	if s.srcTyp == s.dstTyp {
		ptr.Copy(dSlice.Data, sSlice.Data, uintptr(length)*s.sElemSize)
		return true
	}
	for dOffset, sOffset, i := uintptr(0), uintptr(0), 0; i < length; i++ {
		dElemPtr := unsafe.Pointer(uintptr(dSlice.Data) + dOffset)
		sElemPtr := unsafe.Pointer(uintptr(sSlice.Data) + sOffset)
		s.elemConverter.convert(dElemPtr, sElemPtr)
		dOffset += s.dElemSize
		sOffset += s.sElemSize
	}
	return true
}
