// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0

package gslice

import (
	"github.com/smgrushb/conv/internal/generics/gvalue"
)

// Map 将指定切片进行一次映射
func Map[S ~[]F, T, F any](s S, f func(F) T) []T {
	res := make([]T, 0, len(s))
	for _, v := range s {
		res = append(res, f(v))
	}
	return res
}

// First 获取第一个符合过滤条件的元素
// filter 过滤条件，不传时尝试返回第一项
func First[S ~[]T, T any](s S, f ...func(T) bool) (v T, ok bool) {
	if len(f) == 0 {
		if len(s) == 0 {
			return
		}
		return s[0], true
	}
	fn := f[0]
	for _, v = range s {
		if fn(v) {
			return v, true
		}
	}
	return
}

// FirstOr 获取第一个符合过滤条件的元素，没有符合的元素时返回指定值
// filter 过滤条件，不传时尝试返回第一项
func FirstOr[S ~[]T, T any](s S, or T, f ...func(T) bool) T {
	if v, ok := First(s, f...); ok {
		return v
	}
	return or
}

// FirstOrZero 获取第一个符合过滤条件的元素，没有符合的元素时返回对应类型的零值
// filter 过滤条件，不传时尝试返回第一项
func FirstOrZero[S ~[]T, T any](s S, f ...func(T) bool) T {
	v, _ := First(s, f...)
	return v
}

// Or 对bool切片取或
func Or[S ~[]T, T ~bool](s S) bool {
	for _, v := range s {
		if v {
			return true
		}
	}
	return false
}

// Sum 返回和
func Sum[S ~[]T, T gvalue.Number](s S) (sum T) {
	for _, v := range s {
		sum += v
	}
	return
}

// Sort 对切片排序
func Sort[S ~[]T, T gvalue.Ordered](s S) S {
	for i := (len(s) - 1) / 2; i >= 0; i-- {
		siftDown(s, i)
	}
	for i := len(s) - 1; i >= 1; i-- {
		s[0], s[i] = s[i], s[0]
		siftDown(s[:i], 0)
	}
	return s
}

func siftDown[S ~[]T, T gvalue.Ordered](s S, node int) {
	for {
		child := 2*node + 1
		if child >= len(s) {
			return
		}
		if child+1 < len(s) && s[child] < s[child+1] {
			child++
		}
		if s[node] >= s[child] {
			return
		}
		s[node], s[child] = s[child], s[node]
		node = child
	}
}

// Clone 复制切片
func Clone[S ~[]T, T any](s S) S {
	if s == nil {
		return nil
	}
	res := make(S, len(s))
	copy(res, s)
	return res
}
