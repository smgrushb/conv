// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0

package gmap

// Contains 判断map中是否存在指定key
func Contains[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}

// MapValues 将指定map内的value进行一次映射
func MapValues[K comparable, V1, V2 any](m map[K]V1, f func(V1) V2) map[K]V2 {
	if m == nil {
		return nil
	}
	res := make(map[K]V2, len(m))
	for k, v := range m {
		res[k] = f(v)
	}
	return res
}

// Clone 复制map(浅复制)
func Clone[K comparable, V any](m map[K]V) map[K]V {
	if m == nil {
		return nil
	}
	return cloneWithoutNilCheck(m)
}

// CloneBy 复制map, value采用自定义方法复制（是否是深复制取决去方法内部逻辑）
func CloneBy[K comparable, V any](m map[K]V, f func(V) V) map[K]V {
	if m == nil {
		return nil
	}
	return MapValues(m, f)
}

func cloneWithoutNilCheck[K comparable, V any](m map[K]V) map[K]V {
	res := make(map[K]V, len(m))
	for k, v := range m {
		res[k] = v
	}
	return res
}
