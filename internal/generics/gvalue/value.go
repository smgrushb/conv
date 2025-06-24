// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0

package gvalue

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"unsafe"
)

type p struct {
	p    uintptr
	data unsafe.Pointer
}

func IsNil[T any](v T) bool {
	i := any(v)
	return (*p)(unsafe.Pointer(&i)).data == nil
}

func Zero[T any]() (v T) { return }

func IsZero[T comparable](v T) bool { return v == Zero[T]() }

// Valid 当v为零值时，返回指定值，否则返回v
func Valid[T comparable](v, or T) T {
	// 不能用choose.If 会循环依赖
	if IsZero(v) {
		return or
	}
	return v
}

func In[T comparable](v T, s ...T) bool {
	for _, vv := range s {
		if v == vv {
			return true
		}
	}
	return false
}

// Safe 当v为nil时，基于反射初始化v并返回，反之返回v
// 注意: 当v为nil且类型为接口并且无附带实现类型时，无法初始化，将直接返回
func Safe[T any](v T) T {
	vt := reflect.TypeOf(v)
	if vt == nil {
		return v
	}

	result, _ := safeInitItem(reflect.ValueOf(v), vt)
	return result.Interface().(T)
}

func safeInitItem(v reflect.Value, t reflect.Type) (reflect.Value, bool) {
	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			newPtr := reflect.New(t.Elem())
			elemValue, _ := safeInitItem(reflect.New(t.Elem()).Elem(), t.Elem())
			newPtr.Elem().Set(elemValue)
			return newPtr, true
		}
		elemValue, nilElem := safeInitItem(v.Elem(), t.Elem())
		v.Elem().Set(elemValue)
		return v, nilElem
	case reflect.Map:
		if v.IsNil() {
			return reflect.MakeMap(t), true
		}
		return v, false
	case reflect.Slice:
		if v.IsNil() {
			return reflect.MakeSlice(t, 0, 0), true
		}
		return v, false
	case reflect.Chan:
		if v.IsNil() {
			return reflect.MakeChan(t, 0), true
		}
		return v, false
	case reflect.Interface:
		if v.IsNil() {
			return reflect.Zero(t), true
		}
		elemValue := v.Elem()
		initializedElem, nilItem := safeInitItem(elemValue, elemValue.Type())
		return initializedElem.Convert(t), nilItem
	}
	return v, false
}

// ReflectPathType 在ReflectType的输入上增加pkg路径前缀
// 主要用于区分不同路径下相同包名和类型名称的问题
func ReflectPathType[T any](v ...T) string {
	var t T
	// 如果v是接口类型且指定实现类型，不对t赋值的话会打印接口类型而非实现类型
	if len(v) > 0 {
		t = v[0]
	}
	rt := reflect.TypeOf(t)
	if rt == nil || rt.Kind() == reflect.Pointer {
		rt = reflect.TypeOf(&t).Elem()
	}
	return fmt.Sprintf("%s.%s", rt.PkgPath(), rt.String())
}

func AsInt64[T any](t T) int64 {
	if IsNil(t) {
		return 0
	}
	rv := reflect.ValueOf(t)
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return 0
		}
		rv = rv.Elem()
	}
	switch v := rv.Interface().(type) {
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case uint:
		return int64(v)
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		return int64(v)
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	case string:
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			return i
		}
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return int64(f)
		}
	case bool:
		if v {
			return 1
		}
		return 0
	}
	if rv.CanInt() {
		return rv.Int()
	}
	if rv.CanUint() {
		return int64(rv.Uint())
	}
	return 0
}

func AsFloat64[T any](t T) float64 {
	if IsNil(t) {
		return 0
	}
	rv := reflect.ValueOf(t)
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return 0.0
		}
		rv = rv.Elem()
	}
	switch v := rv.Interface().(type) {
	case int:
		return float64(v)
	case int8:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint:
		return float64(v)
	case uint8:
		return float64(v)
	case uint16:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case float32:
		return float64(v)
	case float64:
		return v
	case string:
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f
		}
	case bool:
		if v {
			return 1.0
		}
		return 0.0
	}
	if rv.CanFloat() {
		return rv.Float()
	}
	if rv.CanInt() {
		return float64(rv.Int())
	}
	if rv.CanUint() {
		return float64(rv.Uint())
	}
	return 0.0
}

func AsString[T any](t T) string {
	if IsNil(t) {
		return ""
	}
	rv := reflect.ValueOf(t)
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return ""
		}
		rv = rv.Elem()
	}
	switch v := rv.Interface().(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	case []byte:
		return string(v)
	case []rune:
		return string(v)
	case fmt.Stringer:
		return v.String()
	case io.Reader:
		bs, _ := io.ReadAll(v)
		return string(bs)
	default:
		if rv.Kind() == reflect.String {
			return rv.String()
		}
		return fmt.Sprint(t)
	}
}

func TypeConv[To, From any](v From) (t To, ok bool) {
	t, ok = any(v).(To)
	return
}

func TypeAs[To, From any]() (To, bool) {
	return TypeConv[To](Zero[From]())
}
