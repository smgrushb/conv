// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0
// Inspired by coven (MIT License) by petersunbag

package internal

import (
	"unsafe"
)

type customConverter struct {
	version int8
	cvtOp   func(unsafe.Pointer, unsafe.Pointer)
	cvtOpV2 func(unsafe.Pointer, unsafe.Pointer) bool
}

func Custom(cvtOp func(unsafe.Pointer, unsafe.Pointer)) converter {
	return &customConverter{version: 1, cvtOp: cvtOp}
}

func CustomV2(cvtOp func(unsafe.Pointer, unsafe.Pointer) bool) converter {
	return &customConverter{version: 2, cvtOpV2: cvtOp}
}

func (c *customConverter) convert(dPtr, sPtr unsafe.Pointer) bool {
	if c.version == 1 {
		c.cvtOp(dPtr, sPtr)
		return true
	}
	return c.cvtOpV2(dPtr, sPtr)
}
