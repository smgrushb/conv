// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0

package set

import (
	"github.com/smgrushb/conv/internal/generics/gmap"
	"github.com/smgrushb/conv/internal/generics/gvalue"
)

const initSize = 1 << 5

// Set 集合
// 注意: 非线程安全
type Set[T comparable] struct {
	m map[T]struct{}
}

// New 初始化集合，可传入初始化元素
func New[T comparable](v ...T) *Set[T] {
	m := make(map[T]struct{}, gvalue.Valid(len(v), initSize))
	for _, v := range v {
		m[v] = struct{}{}
	}
	return newSet(m)
}

func newSet[T comparable](m map[T]struct{}) *Set[T] {
	return &Set[T]{m: m}
}

// Clone 复制当前集合
func (s *Set[T]) Clone() *Set[T] {
	if s == nil {
		return nil
	}
	ns := New[T]()
	for v := range s.m {
		ns.m[v] = struct{}{}
	}
	return ns
}

// ForEach 对集合中的所有元素执行指定方法
func (s *Set[T]) ForEach(f func(T)) {
	for v := range s.m {
		f(v)
	}
}

// Add 添加元素，返回是否添加成功
func (s *Set[T]) Add(v T) bool {
	_, ok := s.m[v]
	if !ok {
		s.m[v] = struct{}{}
	}
	return !ok
}

// AddN 批量添加元素
func (s *Set[T]) AddN(v ...T) *Set[T] {
	for _, vv := range v {
		s.m[vv] = struct{}{}
	}
	return s
}

// Len 获取集合长度，为nil时返回0
func (s *Set[T]) Len() int {
	if s == nil {
		return 0
	}
	return len(s.m)
}

// Empty 判断集合是否为空，为nil时返回true
func (s *Set[T]) Empty() bool { return s.Len() == 0 }

// Contains 是否包含元素
func (s *Set[T]) Contains(v T) bool {
	return gmap.Contains(s.m, v)
}
