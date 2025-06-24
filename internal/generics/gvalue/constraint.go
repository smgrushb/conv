// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0

package gvalue

// Signed 有符号整数
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned 无符号整数
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer 全体整数
type Integer interface {
	Signed | Unsigned
}

// Float 全体浮点数
type Float interface {
	~float32 | ~float64
}

// Number 全体数字
type Number interface {
	Integer | Float
}

// Ordered 可排序类型
type Ordered interface {
	Number | ~string
}

// Complex 全体复数
type Complex interface {
	~complex64 | ~complex128
}
