// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0
// Inspired by coven (MIT License) by petersunbag

package conv

import (
	"github.com/smgrushb/conv/internal"
)

func SetStructTageName(name string) {
	internal.TagName = name
}

func SetStructPriorityTagName(name string) {
	internal.PriorityTagName = name
}
