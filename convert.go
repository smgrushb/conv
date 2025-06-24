// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0
// Inspired by coven (MIT License) by petersunbag

package conv

import (
	"errors"
	"fmt"
	"github.com/smgrushb/conv/internal"
	"github.com/smgrushb/conv/internal/generics/gvalue"
	"github.com/smgrushb/conv/option"
	"reflect"
)

type Converter interface {
	Convert(dPtr, sPtr any) error
}

func NewConverter[To, From any](opts ...option.Option) (Converter, error) {
	return NewConverterOf(gvalue.Zero[To](), gvalue.Zero[From](), opts...)
}

func NewConverterOf[To, From any](dst To, src From, opts ...option.Option) (Converter, error) {
	return newConverterOf(dst, src, 0, opts...)
}

func newConverterOf[To, From any](dst To, src From, phase int, opts ...option.Option) (Converter, error) {
	dstTyp, srcTyp := reflect.TypeOf(dst), reflect.TypeOf(src)
	// 接口类型且不是any
	if dstTyp == nil && !internal.IsAnyType[To]() {
		return nil, fmt.Errorf("bad destination type:%s", gvalue.ReflectPathType(dst))
	}
	// 接口类型
	if srcTyp == nil {
		return nil, fmt.Errorf("bad source type:%s", gvalue.ReflectPathType(src))
	}
	if c := internal.NewConverter(dstTyp, srcTyp, internal.GetOption(phase, opts...)); c == nil {
		return nil, fmt.Errorf("can't convert source type %s to destination type %s", srcTyp, dstTyp)
	} else {
		return c, nil
	}
}

func ConvertTo[To, From any](from From, to *To, opts ...option.Option) error {
	return convertTo(from, to, 0, opts...)
}

func convertTo[To, From any](from From, to *To, phase int, opts ...option.Option) error {
	if gvalue.IsNil(to) {
		return errors.New("[conv]destination should be a pointer")
	}
	c, err := newConverterOf(to, from, phase, opts...)
	if err != nil {
		return err
	}
	// 如果是接口类型，判断接口的实现类型是否是指针，是则直接往下传，不是转成指针类型
	if fromType := internal.ReflectType[From](); fromType.Kind() == reflect.Interface {
		fromValue := reflect.ValueOf(from)
		if fromValue.Kind() == reflect.Pointer {
			return c.Convert(to, from)
		}
		if !fromValue.IsValid() {
			return c.Convert(to, &from)
		}
		ptr := reflect.New(fromValue.Type())
		ptr.Elem().Set(fromValue)
		return c.Convert(to, ptr.Interface())
	}
	return c.Convert(to, &from)
}

func Convert[To, From any](from From, opts ...option.Option) (t To, err error) {
	res := gvalue.Safe(gvalue.Zero[To]())
	if reflect.TypeOf(res) == nil && !internal.IsAnyType(res) {
		err = fmt.Errorf("bad dst type:%s", gvalue.ReflectPathType[To]())
		return
	}
	if err = ConvertTo(from, &res, opts...); err != nil {
		return
	}
	return res, nil
}

func OstrichConvert[To, From any](from From, opts ...option.Option) To {
	t, _ := Convert[To](from, opts...)
	return t
}

func TwoPhaseConvertTo[To, Temp, From any](from From, to *To, opts ...option.Option) error {
	temp := gvalue.Safe(gvalue.Zero[Temp]())
	if reflect.TypeOf(temp) == nil && !internal.IsAnyType(temp) {
		return fmt.Errorf("bad temp type:%s", gvalue.ReflectPathType[To]())
	}
	if err := convertTo(from, &temp, 1, opts...); err != nil {
		return err
	}
	return convertTo(temp, to, 2, opts...)
}

func TwoPhaseConv[To, Temp, From any](from From, opts ...option.Option) (t To, err error) {
	temp := gvalue.Safe(gvalue.Zero[Temp]())
	if reflect.TypeOf(temp) == nil && !internal.IsAnyType(temp) {
		err = fmt.Errorf("bad temp type:%s", gvalue.ReflectPathType[To]())
		return
	}
	res := gvalue.Safe(gvalue.Zero[To]())
	if reflect.TypeOf(res) == nil && !internal.IsAnyType(res) {
		err = fmt.Errorf("bad dst type:%s", gvalue.ReflectPathType[To]())
		return
	}
	if err = convertTo(from, &temp, 1, opts...); err != nil {
		return
	}
	if err = convertTo(temp, &res, 2, opts...); err != nil {
		return
	}
	return res, nil
}

func OstrichTwoPhaseConvert[To, Temp, From any](from From, opts ...option.Option) To {
	t, _ := TwoPhaseConv[To, Temp](from, opts...)
	return t
}
