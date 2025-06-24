// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0
// Inspired by coven (MIT License) by petersunbag

package convextend

import (
	"github.com/smgrushb/conv/internal"
	"sync"
)

var (
	convProtoOnce  sync.Once
	ProtoConverter []internal.CustomConverter
)

func ConvProto(custom ...internal.CustomConverter) {
	convProtoOnce.Do(func() {
		internal.ConvProto = true
		internal.ProtoConverter = append(custom, ProtoConverter...)
	})
}
