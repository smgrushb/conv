// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0
// Inspired by coven (MIT License) by petersunbag

package conv

import (
	"fmt"
	"github.com/smgrushb/conv/constant"
	"github.com/smgrushb/conv/internal"
	"github.com/smgrushb/conv/internal/generics/gptr"
	"time"
)

var timeType = internal.ReflectType[time.Time]()

func RegisterTimeWrapper[T any]() {
	if rt := internal.ReflectType[T](); !rt.ConvertibleTo(timeType) {
		panic(fmt.Errorf("generic type [%s] should be convertible to time.Time", rt.String()))
	}
	internal.TimeWrappers = append(internal.TimeWrappers, &internal.TimeWrapper[T]{})
}

func SetTimeFormat(format string) {
	internal.TimeFormat = format
}

func SetMinUnix(unix int64) {
	internal.MinUnix = &unix
}

func SetMinUnixBy(t time.Time) {
	internal.MinUnix = gptr.Of(t.Unix())
}

func SetMinUnixScene(scene ...internal.MinUnixSceneType) {
	internal.MinUnixScene = constant.MinUnixScene(scene...)
}
