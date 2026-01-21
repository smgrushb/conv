// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0
// Inspired by coven (MIT License) by petersunbag

package convextend

import (
	"fmt"
	"github.com/smgrushb/conv/internal"
	"github.com/smgrushb/conv/internal/generics/gslice"
	"github.com/smgrushb/conv/internal/generics/gvalue"
	"github.com/smgrushb/conv/internal/unsafeheader"
	"github.com/smgrushb/conv/model"
	"reflect"
	"unsafe"
)

type KVP[K comparable, V any] interface {
	GetKey() K
	SetKey(K)
	GetValue() V
	SetValue(V)
}

func newKeyValuePair[K comparable, V any]() KVP[K, V] {
	return &model.KeyValuePair[K, V]{}
}

type map2KVP[K comparable, V any] struct {
	kvp   KVP[K, V]
	rtype reflect.Type
	size  uintptr
}

func Map2KVP[K comparable, V any](kvp ...KVP[K, V]) internal.CustomConverter {
	v := gslice.FirstOr(kvp, newKeyValuePair[K, V]())
	rtype := reflect.TypeOf(v).Elem()
	return &map2KVP[K, V]{kvp: v, rtype: rtype, size: rtype.Size()}
}

func (m *map2KVP[K, V]) Is(dstTyp, srcTyp reflect.Type) bool {
	if srcTyp.Kind() != reflect.Map || srcTyp.Key() != internal.ReflectType[K]() || srcTyp.Elem() != internal.ReflectType[V]() {
		return false
	}
	return dstTyp.Kind() == reflect.Slice && dstTyp.Elem() == m.rtype
}

func (m *map2KVP[K, V]) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		s := *(*map[K]V)(sPtr)
		length := len(s)
		dSlice := (*unsafeheader.SliceHeader)(dPtr)
		dSlice.Len = length
		if dSlice.Cap < length || dSlice.Data == nil {
			sliceType := reflect.SliceOf(m.rtype)
			dv := reflect.NewAt(sliceType, dPtr).Elem()
			newVal := reflect.MakeSlice(sliceType, length, length)
			dv.Set(newVal)
			dSlice.Data = unsafe.Pointer(newVal.Pointer())
			dSlice.Cap = length
		}
		var offset uintptr
		for k, v := range s {
			elemPtr := unsafe.Pointer(uintptr(dSlice.Data) + offset)
			elemVal := reflect.NewAt(m.rtype, elemPtr)
			kvp := elemVal.Interface().(KVP[K, V])
			kvp.SetKey(k)
			kvp.SetValue(v)
			offset += m.size
		}
	}
}

func (m *map2KVP[K, V]) Key() string {
	return fmt.Sprintf("[map2KVP-%s]", gvalue.ReflectPathType(m.kvp))
}

type kvp2Map[K comparable, V any] struct {
	kvp   KVP[K, V]
	rtype reflect.Type
	size  uintptr
}

func KVP2Map[K comparable, V any](kvp ...KVP[K, V]) internal.CustomConverter {
	v := gslice.FirstOr(kvp, newKeyValuePair[K, V]())
	rtype := reflect.TypeOf(v).Elem()
	return &kvp2Map[K, V]{kvp: v, rtype: rtype, size: rtype.Size()}
}

func (m *kvp2Map[K, V]) Is(dstTyp, srcTyp reflect.Type) bool {
	if dstTyp.Kind() != reflect.Map || dstTyp.Key() != internal.ReflectType[K]() || dstTyp.Elem() != internal.ReflectType[V]() {
		return false
	}
	return srcTyp.Kind() == reflect.Slice && srcTyp.Elem() == m.rtype
}

func (m *kvp2Map[K, V]) Converter() func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
	return func(dPtr unsafe.Pointer, sPtr unsafe.Pointer) {
		sSlice := (*unsafeheader.SliceHeader)(sPtr)
		length := sSlice.Len
		data := make(map[K]V, length)
		var offset uintptr
		for i := 0; i < length; i++ {
			elemPtr := unsafe.Pointer(uintptr(sSlice.Data) + offset)
			elemVal := reflect.NewAt(m.rtype, elemPtr)
			kvp := elemVal.Interface().(KVP[K, V])
			data[kvp.GetKey()] = kvp.GetValue()
			offset += m.size
		}
		*(*map[K]V)(dPtr) = data
	}
}

func (m *kvp2Map[K, V]) Key() string {
	return fmt.Sprintf("[kvp2Map-%s]", gvalue.ReflectPathType(m.kvp))
}
