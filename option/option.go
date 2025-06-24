// This file is part of the original coven project:
// https://github.com/petersunbag/coven
//
// Copyright © 2018 petersunbag
// Licensed under the MIT License
// https://opensource.org/licenses/MIT
// Modified by smgrushb in 2025

package option

import (
	"github.com/smgrushb/conv/constant"
	convextend "github.com/smgrushb/conv/extend"
	"github.com/smgrushb/conv/internal"
	"github.com/smgrushb/conv/internal/generics/gptr"
	"github.com/smgrushb/conv/internal/generics/gslice"
	"time"
)

type Option = func(*internal.StructOption)

// WhiteList 字段白名单，结构体类型需要a.b方式描述
// 生效于源类型
// 注意: 结构体类型映射的目标类型是any时对字段的限制失效，整个结构体字段的值都会被映射
func WhiteList(fieldNames ...string) Option {
	return func(o *internal.StructOption) {
		o.WhiteListFields.AddN(fieldNames...)
	}
}

// Banned 屏蔽字段，结构体类型需要a.b方式描述
// 生效于目标类型,生效先于Alias
// 注意: 结构体类型映射的目标类型是any时对字段的限制失效，整个结构体字段的值都会被映射
func Banned(fieldNames ...string) Option {
	return func(o *internal.StructOption) {
		o.BannedFields.AddN(fieldNames...)
	}
}

// Alias 字段别名，结构体类型需要a.b方式描述
// 生效于目标类型,生效后于Banned
// 注意: 结构体类型映射的目标类型是any时对字段的别名失效
func Alias(k, v string) Option {
	return func(o *internal.StructOption) {
		o.AliasFields[k] = v
	}
}

// AliasMap 字段别名，结构体类型需要a.b方式描述
// 生效于目标类型,生效后于Banned
// 注意: 结构体类型映射的目标类型是any时对字段的别名失效
func AliasMap(m map[string]string) Option {
	return func(o *internal.StructOption) {
		for k, v := range m {
			o.AliasFields[k] = v
		}
	}
}

// TimeFormat 指定time.Time转string的format格式
func TimeFormat(format string) Option {
	return func(o *internal.StructOption) {
		o.TimeFormat = format
	}
}

// MinUnix 指定time.Time(包含time.Time各别名类型)转换到非time.Time(包含time.Time各别名类型)时的最低有效时间，低于此时间则不转换
// 目前场景：time/string互转(默认生效)，time转any(默认不生效)
// 可以传入scene来自订生效规则
// 注意：别名类型仅检查被RegisterTimeWrapper方法注册的
func MinUnix(unix int64, scene ...internal.MinUnixSceneType) Option {
	return func(o *internal.StructOption) {
		o.MinUnix = &unix
		if len(scene) > 0 {
			o.MinUnixScene = constant.MinUnixScene(scene...)
		} else {
			o.MinUnixScene = constant.DefaultMinUnixScene
		}
	}
}

// MinUnixBy 指定time.Time(包含time.Time各别名类型)转换到非time.Time(包含time.Time各别名类型)时的最低有效时间，低于此时间则不转换
// 目前场景：time/string互转(默认生效)，time转any(默认不生效)
// 可以传入scene来自订生效规则
// 注意：别名类型仅检查被RegisterTimeWrapper方法注册的
func MinUnixBy(t time.Time, scene ...internal.MinUnixSceneType) Option {
	return func(o *internal.StructOption) {
		o.MinUnix = gptr.Of(t.Unix())
		if len(scene) > 0 {
			o.MinUnixScene = constant.MinUnixScene(scene...)
		} else {
			o.MinUnixScene = constant.DefaultMinUnixScene
		}
	}
}

// TagName 指定解析字段名称时使用的tag
func TagName(name string) Option {
	return func(o *internal.StructOption) {
		o.TagName = name
	}
}

// PriorityTagName 指定解析字段名称时使用的高优先级tag
func PriorityTagName(name string) Option {
	return func(o *internal.StructOption) {
		o.PriorityTagName = name
	}
}

// IgnorePrivateFields 屏蔽私有字段(结构体转map外场景生效)
func IgnorePrivateFields() Option {
	return func(o *internal.StructOption) {
		o.IgnorePrivateFields = true
	}
}

// IncludePrivateFields 包含私有字段(仅结构体转map场景生效)
func IncludePrivateFields() Option {
	return func(o *internal.StructOption) {
		o.IncludePrivateFields = true
	}
}

// IgnoreEmptyFields 忽略零值字段(仅结构体转map场景生效)
func IgnoreEmptyFields() Option {
	return func(o *internal.StructOption) {
		o.IgnoreEmptyFields = true
	}
}

// IgnoreTag 屏蔽标签，使用字段名
func IgnoreTag() Option {
	return func(o *internal.StructOption) {
		o.IgnoreTag = true
	}
}

// IgnoreFunc 屏蔽方法，不再进行方法调用获取返回值进行转换
func IgnoreFunc() Option {
	return func(o *internal.StructOption) {
		o.IgnoreFunc = true
	}
}

// UseStrings 如果类型实现了fmt.Stringer, 优先调用String()方法来转换成string，优先级高于UseMarshal
func UseStrings() Option {
	return func(o *internal.StructOption) {
		o.UseStrings = true
	}
}

// UseMarshal 如果类型实现了json.Marshaler, 优先调用MarshalJSON()方法来转换成string，优先级低于UseStrings
func UseMarshal() Option {
	return func(o *internal.StructOption) {
		o.UseMarshal = true
	}
}

// StrBytesZeroCopy string和[]byte互转时是否零拷贝, 默认零拷贝
func StrBytesZeroCopy(zeroCopy ...bool) Option {
	return func(o *internal.StructOption) {
		o.StrBytesZeroCopy = gslice.FirstOr(zeroCopy, true)
	}
}

// CustomConverter 自定义转换器
func CustomConverter(custom ...internal.CustomConverter) Option {
	return func(o *internal.StructOption) {
		o.CustomConv = append(o.CustomConv, custom...)
	}
}

// ConvProto 自定义Proto转换器
// 自动携带所有基础转换器
func ConvProto(custom ...internal.CustomConverter) Option {
	return CustomConverter(append(custom, convextend.ProtoConverter...)...)
}

// Phase 两阶段转换时分开指定每个阶段的option时使用
func Phase(opts ...Option) Option {
	return func(o *internal.StructOption) {
		for _, opt := range opts {
			opt(o)
		}
		o.Phase++
	}
}
