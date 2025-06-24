// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0
// Inspired by coven (MIT License) by petersunbag

package unsafeheader

import (
	"unsafe"
)

type StringHeader struct {
	Data unsafe.Pointer
	Len  int
}

type SliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

type EmptyInterface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
}

func NewEmptyInterface(typ, word unsafe.Pointer) EmptyInterface {
	return EmptyInterface{typ: typ, word: word}
}

func (e *EmptyInterface) GetTyp() unsafe.Pointer {
	return e.typ
}

func (e *EmptyInterface) GetWord() unsafe.Pointer {
	return e.word
}

func (e *EmptyInterface) SetWord(p unsafe.Pointer) *EmptyInterface {
	e.word = p
	return e
}
