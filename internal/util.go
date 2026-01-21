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
	"reflect"
	"unsafe"
)

func dereferencedType(t reflect.Type) (reflect.Type, int) {
	var d int
	for k := t.Kind(); k == reflect.Ptr || k == reflect.Interface; k = t.Kind() {
		if t == ptr.AnyType {
			break
		}
		t = t.Elem()
		d++
	}
	return t, d
}

func dereferencedTypeDeep(t reflect.Type) (reflect.Type, bool) {
	var isPtr bool
	for k := t.Kind(); k == reflect.Ptr || k == reflect.Interface; k = t.Kind() {
		if t == ptr.AnyType {
			break
		}
		t = t.Elem()
		isPtr = isPtr || k == reflect.Ptr
	}
	return t, isPtr
}

func dereferencedValue(value any) reflect.Value {
	v := reflect.ValueOf(value)
	for k := v.Kind(); k == reflect.Ptr || k == reflect.Interface; k = v.Kind() {
		if v.Type() == ptr.AnyType {
			break
		}
		v = v.Elem()
	}
	return v
}

func referDeep(t reflect.Type) (reflect.Type, int) {
	var d int
	for k := t.Kind(); k == reflect.Ptr; k = t.Kind() {
		t = t.Elem()
		d++
	}
	return t, d
}

func newValuePtr(t reflect.Type) unsafe.Pointer {
	var v unsafe.Pointer
	if v = ptr.NewValuePtr(t.Kind()); v == nil {
		v = unsafe.Pointer(reflect.New(t).Elem().UnsafeAddr())
	}
	return v
}

func PtrOfAny(v reflect.Value) unsafe.Pointer {
	if v.CanAddr() {
		return unsafe.Pointer(v.UnsafeAddr())
	}
	newVal := reflect.New(v.Type())
	newVal.Elem().Set(v)
	return unsafe.Pointer(newVal.Pointer())
}

func ReflectType[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}
