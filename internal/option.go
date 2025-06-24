// This file is part of the original coven project:
// https://github.com/petersunbag/coven
//
// Copyright © 2018 petersunbag
// Licensed under the MIT License
// https://opensource.org/licenses/MIT
// Modified by smgrushb in 2025

package internal

import (
	"fmt"
	"github.com/bytedance/sonic/encoder"
	"github.com/smgrushb/conv/internal/generics/collection/set"
	"github.com/smgrushb/conv/internal/generics/gmap"
	"github.com/smgrushb/conv/internal/generics/gslice"
	"github.com/smgrushb/conv/internal/generics/gvalue"
	"reflect"
	"strings"
	"unsafe"
)

var (
	ConvProto           bool
	ProtoConverter      []CustomConverter
	defaultStructOption = func() *StructOption {
		// TimeFormat的优先级是option > timeWrapper > global，所以这里不做处理
		return &StructOption{
			TagName:         TagName,
			PriorityTagName: PriorityTagName,
			CustomConv:      ProtoConverter,
		}
	}
)

type Option = func(*StructOption)

type CustomConverter interface {
	Is(dstTyp, srcTyp reflect.Type) bool
	Converter() func(dPtr, sPtr unsafe.Pointer)
	Key() string
}

type StructOption struct {
	Phase                int `json:"-"`
	IgnorePrivateFields  bool
	IncludePrivateFields bool
	IgnoreEmptyFields    bool
	IgnoreTag            bool
	IgnoreFunc           bool
	UseStrings           bool
	UseMarshal           bool
	StrBytesZeroCopy     bool
	TagName              string
	PriorityTagName      string
	TimeFormat           string
	MinUnix              *int64
	MinUnixScene         MinUnixSceneType
	BannedFields         *set.Set[string]
	WhiteListFields      *set.Set[string]
	AliasFields          map[string]string
	NestedOption         map[string]*StructOption
	CustomConv           []CustomConverter `json:"-"`
}

func newOption() *StructOption {
	return &StructOption{
		StrBytesZeroCopy: true,
		BannedFields:     set.New[string](),
		WhiteListFields:  set.New[string](),
		AliasFields:      make(map[string]string),
		NestedOption:     make(map[string]*StructOption),
	}
}

func GetOption(phase int, opts ...Option) *StructOption {
	if len(opts) == 0 {
		return defaultStructOption()
	}
	opt := newOption()
	for _, v := range opts {
		if phase == 0 {
			v(opt)
		} else {
			old := opt.Clone()
			v(opt)
			if opt.Phase < phase && opt.Phase != 0 {
				opt = newOption()
			} else if opt.Phase > phase {
				opt = old
				break
			}
		}
	}
	if ConvProto {
		opt.CustomConv = append(opt.CustomConv, ProtoConverter...)
	}
	// TimeFormat的优先级是option > timeWrapper > global，所以这里不做处理
	opt.TagName = gvalue.Valid(opt.TagName, TagName)
	opt.PriorityTagName = gvalue.Valid(opt.PriorityTagName, PriorityTagName)
	return opt.parse()
}

func (o *StructOption) Clone() *StructOption {
	if o == nil {
		return nil
	}
	return &StructOption{
		Phase:                o.Phase,
		IgnorePrivateFields:  o.IgnorePrivateFields,
		IncludePrivateFields: o.IncludePrivateFields,
		IgnoreEmptyFields:    o.IgnoreEmptyFields,
		IgnoreTag:            o.IgnoreTag,
		IgnoreFunc:           o.IgnoreFunc,
		UseStrings:           o.UseStrings,
		UseMarshal:           o.UseMarshal,
		StrBytesZeroCopy:     o.StrBytesZeroCopy,
		TagName:              o.TagName,
		PriorityTagName:      o.PriorityTagName,
		TimeFormat:           o.TimeFormat,
		MinUnix:              o.MinUnix,
		MinUnixScene:         o.MinUnixScene,
		BannedFields:         o.BannedFields.Clone(),
		WhiteListFields:      o.WhiteListFields.Clone(),
		AliasFields:          gmap.Clone(o.AliasFields),
		NestedOption:         gmap.CloneBy(o.NestedOption, (*StructOption).Clone),
		CustomConv:           o.CustomConv,
	}
}

func (o *StructOption) inherit(parent *StructOption) *StructOption {
	o.IgnorePrivateFields = parent.IgnorePrivateFields
	o.IncludePrivateFields = parent.IncludePrivateFields
	o.IgnoreEmptyFields = parent.IgnoreEmptyFields
	o.IgnoreTag = parent.IgnoreTag
	o.IgnoreFunc = parent.IgnoreFunc
	o.TagName = parent.TagName
	o.PriorityTagName = parent.PriorityTagName
	o.TimeFormat = parent.TimeFormat
	o.MinUnix = parent.MinUnix
	o.MinUnixScene = parent.MinUnixScene
	o.CustomConv = parent.CustomConv
	return o
}

func (o *StructOption) parse() *StructOption {
	o.BannedFields.ForEach(func(s string) {
		if first, second, ok := split(s); ok {
			nest, ok := o.NestedOption[first]
			if !ok {
				nest = newOption().inherit(o)
				o.NestedOption[first] = nest
			}
			nest.BannedFields.Add(second)
		}
	})
	o.WhiteListFields.ForEach(func(s string) {
		if first, second, ok := split(s); ok {
			o.WhiteListFields.Add(first)
			nest, ok := o.NestedOption[first]
			if !ok {
				nest = newOption().inherit(o)
				o.NestedOption[first] = nest
			}
			nest.WhiteListFields.Add(second)
		}
	})
	for f, a := range o.AliasFields {
		if first, second, ok := split(f); ok {
			nest, ok := o.NestedOption[first]
			if !ok {
				nest = newOption().inherit(o)
				o.NestedOption[first] = nest
			}
			nest.AliasFields[second] = a
		}
	}
	for _, nest := range o.NestedOption {
		nest.parse()
	}
	return o
}

func (o *StructOption) key() string {
	if o == nil {
		return ""
	}
	convKey := strings.Join(gslice.Sort(gslice.Map(o.CustomConv, CustomConverter.Key)), ";")
	bs, _ := encoder.Encode(o, encoder.SortMapKeys)
	return fmt.Sprintf("%s%s", string(bs), convKey)
}

func split(s string) (first, second string, ok bool) {
	const sep = "."
	index := strings.Index(s, sep)
	if ok = index != -1; !ok {
		return
	}
	first, second = s[:index], s[index+1:]
	return
}
