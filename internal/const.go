// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0
// Inspired by coven (MIT License) by petersunbag

package internal

type MinUnixSceneType int64

const (
	MinUnixTimeString MinUnixSceneType = 1 << iota
	MinUnixStringTime
	MinUnixTimeAny
)

const (
	DefaultMinUnixScene = MinUnixTimeString | MinUnixStringTime
)

// NilValuePolicy 定义了当转换 nil 指针到值类型时的行为。
type NilValuePolicy int64

const (
	NilValuePolicyIgnore NilValuePolicy = iota // 忽略字段（跳过赋值）
	NilValuePolicyZero                         // 使用源类型的零值（例如：nil *Struct -> Struct{}）
)
