// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0

package set

import (
	"github.com/bytedance/sonic"

	"github.com/smgrushb/conv/internal/generics/gmap"
	"github.com/smgrushb/conv/internal/generics/gvalue"
)

const initSize = 1 << 5

// Set 集合
// 注意: 非线程安全
type Set[T comparable] struct {
	m map[T]struct{}
	s []T
}

// New 初始化集合，可传入初始化元素
func New[T comparable](v ...T) *Set[T] {
	m := make(map[T]struct{}, gvalue.Valid(len(v), initSize))
	s := make([]T, 0, len(m))
	for _, vv := range v {
		if _, ok := m[vv]; !ok {
			m[vv] = struct{}{}
			s = append(s, vv)
		}
	}
	return &Set[T]{m: m, s: s}
}

// Clone 复制当前集合
func (s *Set[T]) Clone() *Set[T] {
	if s == nil {
		return nil
	}
	ns := &Set[T]{
		m: make(map[T]struct{}, len(s.m)),
		s: make([]T, len(s.s)),
	}
	copy(ns.s, s.s)
	for _, v := range s.s {
		ns.m[v] = struct{}{}
	}
	return ns
}

// ForEach 对集合中的所有元素执行指定方法
func (s *Set[T]) ForEach(f func(T)) {
	if s == nil {
		return
	}
	for _, v := range s.s {
		f(v)
	}
}

// Add 添加元素，返回是否添加成功
func (s *Set[T]) Add(v T) bool {
	if s.m == nil {
		s.m = make(map[T]struct{})
	}
	return s.add(v)
}

func (s *Set[T]) add(v T) bool {
	_, ok := s.m[v]
	if !ok {
		s.m[v] = struct{}{}
		s.s = append(s.s, v)
	}
	return !ok
}

// AddN 批量添加元素
func (s *Set[T]) AddN(v ...T) *Set[T] {
	if s.m == nil {
		s.m = make(map[T]struct{})
	}
	for _, vv := range v {
		s.add(vv)
	}
	return s
}

// Len 获取集合长度，为nil时返回0
func (s *Set[T]) Len() int {
	if s == nil {
		return 0
	}
	return len(s.s)
}

// Empty 判断集合是否为空，为nil时返回true
func (s *Set[T]) Empty() bool { return s.Len() == 0 }

// Contains 是否包含元素
func (s *Set[T]) Contains(v T) bool {
	return gmap.Contains(s.m, v)
}

func (s *Set[T]) MarshalJSON() ([]byte, error) {
	if s == nil {
		return []byte("[]"), nil
	}
	return sonic.Marshal(s.s)
}
