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
