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
	"unsafe"
)

type basicConverter struct {
	*convertType
	cvtOp ptr.CvtOp
}

func newBasicConverter(typ *convertType) converter {
	if cvtOp := ptr.GetCvtOp(typ.srcTyp, typ.dstTyp, typ.option.StrBytesZeroCopy); cvtOp != nil {
		return &basicConverter{convertType: typ, cvtOp: cvtOp}
	}
	return nil
}

func (g *basicConverter) convert(dPtr, sPtr unsafe.Pointer) bool {
	g.cvtOp(sPtr, dPtr)
	return true
}
