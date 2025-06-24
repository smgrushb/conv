// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0
// Inspired by coven (MIT License) by petersunbag

package constant

import (
	"github.com/smgrushb/conv/internal"
)

const (
	MinUnixTimeString = internal.MinUnixTimeString
	MinUnixStringTime = internal.MinUnixStringTime
	MinUnixTimeAny    = internal.MinUnixTimeAny
)

const DefaultMinUnixScene = internal.DefaultMinUnixScene

func MinUnixScene(scene ...internal.MinUnixSceneType) (res internal.MinUnixSceneType) {
	for _, v := range scene {
		res |= v
	}
	return
}
