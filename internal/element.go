// This file is part of the original coven project:
// https://github.com/petersunbag/coven
//
// Copyright © 2018 petersunbag
// Licensed under the MIT License
// https://opensource.org/licenses/MIT
// Modified by smgrushb in 2025

package internal

import (
	"github.com/smgrushb/conv/internal/generics/gptr"
	"reflect"
	"unsafe"
)

type elemConverter struct {
	dType               reflect.Type
	sType               reflect.Type
	dDereferType        reflect.Type
	sDereferType        reflect.Type
	dReferDeep          int
	sReferDeep          int
	sEmptyDereferValPtr unsafe.Pointer
	nilValuePolicy      NilValuePolicy
	converter           converter
}

func newElemConverter(dType, sType reflect.Type, option *StructOption) (*elemConverter, bool) {
	ec := &elemConverter{dType: dType, sType: sType}
	ec.dDereferType, ec.dReferDeep = referDeep(dType)
	ec.sDereferType, ec.sReferDeep = referDeep(sType)
	if c := newConverter(ec.dDereferType, ec.sDereferType, option); c != nil {
		if ac, ok := IsAnyConverter(c); ok {
			ac.SetSrcReferDeep(ec.sReferDeep)
		}
		ec.converter = c
		ec.sEmptyDereferValPtr = newValuePtr(ec.sDereferType)
		ec.nilValuePolicy = option.NilValuePolicy
		return ec, true
	}
	return nil, false
}

func (e *elemConverter) convert(dPtr, sPtr unsafe.Pointer) bool {
	for i := 0; i < e.sReferDeep; i++ {
		sPtr = unsafe.Pointer(*((**int)(sPtr)))
		if sPtr == nil {
			if e.dReferDeep > 0 {
				*(**int)(dPtr) = nil
				return true
			}
			if e.nilValuePolicy == NilValuePolicyIgnore {
				return false
			}
			sPtr = e.sEmptyDereferValPtr
			break
		}
	}
	var deep int
	for ; deep < e.dReferDeep; deep++ {
		oldPtr := dPtr
		dPtr = unsafe.Pointer(*((**int)(dPtr)))
		if dPtr == nil {
			dPtr = oldPtr
			break
		}
	}
	if deep := e.dReferDeep - deep; deep > 0 {
		v := newValuePtr(e.dDereferType)
		if !e.converter.convert(v, sPtr) {
			return false
		}
		for i := 0; i < deep; i++ {
			v = unsafe.Pointer(gptr.Of(v))
		}
		*(**int)(dPtr) = *(**int)(v)
	} else {
		return e.converter.convert(dPtr, sPtr)
	}
	return true
}
