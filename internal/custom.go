// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0
// Inspired by coven (MIT License) by petersunbag

package internal

import (
	"unsafe"
)

type customConverter struct {
	cvtOp func(unsafe.Pointer, unsafe.Pointer)
}

func Custom(cvtOp func(unsafe.Pointer, unsafe.Pointer)) converter {
	return &customConverter{cvtOp: cvtOp}
}

func (c *customConverter) convert(dPtr, sPtr unsafe.Pointer) {
	c.cvtOp(dPtr, sPtr)
}
