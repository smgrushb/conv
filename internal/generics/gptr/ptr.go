// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0

package gptr

// Zero 返回指定类型零值的指针
func Zero[T any]() *T {
	var t T
	return &t
}

// Of 取址
// 注意: 此方法主要用于对未定义变量的值进行一次取址。
// 如果已定义变量，调用方法时会进行一次值拷贝，返回地址为拷贝后的数据的地址。
// 请根据实际需求判断是否符合预期，如不符合预期，请使用取址符 &
func Of[T any](v T) *T { return &v }
