// This file is part of the original coven project:
// https://github.com/petersunbag/coven
//
// Copyright © 2018 petersunbag
// Licensed under the MIT License
// https://opensource.org/licenses/MIT
// Modified by smgrushb in 2025

package internal

import (
	"github.com/smgrushb/conv/internal/generics/collection/set"
	"github.com/smgrushb/conv/internal/generics/gptr"
	"github.com/smgrushb/conv/internal/generics/gslice"
	"github.com/smgrushb/conv/internal/generics/gvalue"
	"github.com/smgrushb/conv/internal/ptr"
	"github.com/smgrushb/conv/internal/unsafeheader"
	"google.golang.org/protobuf/proto"
	"reflect"
	"strings"
	"unicode"
	"unsafe"
)

var (
	TagName         = "json"
	PriorityTagName = "conv"
)

var (
	errorRT     = ReflectType[error]()
	intDoubleRT = ReflectType[**int]()
)

type structConverter struct {
	*convertType
	fieldConverters []converter
	size            uintptr
	convMap         bool
	enable          bool // 兜底
}

func newStructConverter(typ *convertType) converter {
	if typ.srcTyp == typ.dstTyp {
		return &structConverter{convertType: typ, size: typ.srcTyp.Size(), enable: true}
	}
	c := &structConverter{convertType: typ}
	key := typ.key()
	// 先预注册进去，不然循环依赖下会循环解析
	createdConverters[key] = &Converter{convertType: typ, converter: c}
	sFields := make(map[string]*structItem)
	extractFields(typ.srcTyp, typ.option, sFields, nil)
	if typ.option == nil || !typ.option.IgnoreFunc {
		extractMethods(typ.srcTyp, typ.option, sFields)
	}
	dFieldIndex := extractFields(typ.dstTyp, typ.option, nil, nil)
	if typ.option != nil {
		dFieldIndex = filterField(dFieldIndex, typ.option.BannedFields)
		dFieldIndex = aliasField(dFieldIndex, typ.option.AliasFields)
	}
	fieldConverters := make([]converter, 0, len(dFieldIndex))
	for _, df := range dFieldIndex {
		if sf, ok := sFields[df.name]; ok {
			if typ.option != nil && !typ.option.WhiteListFields.Empty() && !typ.option.WhiteListFields.Contains(sf.name) {
				continue
			}
			var nestOption *StructOption
			if typ.option != nil && typ.option.NestedOption != nil {
				nestOption = typ.option.NestedOption[df.name]
			}
			if nestOption == nil {
				nestOption = typ.option
			}
			if fc := newFieldConverter(*df, *sf, nestOption); fc != nil {
				fieldConverters = append(fieldConverters, fc)
			}
		}
	}
	if len(fieldConverters) == 0 {
		// 把预注册的内容删了
		delete(createdConverters, key)
		return nil
	}
	c.fieldConverters = fieldConverters
	c.enable = true
	return c
}

func newStructMapConverter(typ *convertType, valueType reflect.Type) converter {
	c := &structConverter{convertType: typ, convMap: true}
	key := typ.key()
	// 先预注册进去，不然循环依赖下会循环解析
	createdConverters[key] = &Converter{convertType: typ, converter: c}
	sFields := make(map[string]*structItem)
	// 映射到map时源和目标无区别，直接banned和alias直接作用到源上即可
	sFieldIndex := extractFieldsOnMap(typ.srcTyp, typ.option, sFields, nil)
	if typ.option != nil {
		sFieldIndex = filterField(sFieldIndex, typ.option.BannedFields)
		sFieldIndex = aliasField(sFieldIndex, typ.option.AliasFields)
	}
	fieldConverters := make([]converter, 0, len(sFieldIndex))
	for _, sf := range sFieldIndex {
		if typ.option != nil && !typ.option.WhiteListFields.Empty() && !typ.option.WhiteListFields.Contains(sf.name) {
			continue
		}
		var nestOption *StructOption
		if typ.option != nil && typ.option.NestedOption != nil {
			nestOption = typ.option.NestedOption[sf.name]
		}
		if nestOption == nil {
			nestOption = typ.option
		}
		if fc := newFieldMapConverter(valueType, *sf, nestOption); fc != nil {
			fieldConverters = append(fieldConverters, fc)
		}
	}
	if len(fieldConverters) == 0 {
		// 把预注册的内容删了
		delete(createdConverters, key)
		return nil
	}
	c.fieldConverters = fieldConverters
	c.enable = true
	return c
}

func filterField(fields []*structItem, bannedFields *set.Set[string]) []*structItem {
	if bannedFields == nil || bannedFields.Empty() {
		return fields
	}
	res := make([]*structItem, 0, len(fields)/2)
	for _, f := range fields {
		if !bannedFields.Contains(f.name) {
			res = append(res, f)
		}
	}
	return res
}

func aliasField(fields []*structItem, aliasFields map[string]string) []*structItem {
	if len(aliasFields) == 0 {
		return fields
	}
	for _, f := range fields {
		if alias, ok := aliasFields[f.name]; ok {
			f.name = alias
		}
	}
	return fields
}

func (s *structConverter) convert(dPtr, sPtr unsafe.Pointer) {
	if !s.enable {
		return
	}
	if s.convMap {
		s.mapConvert(dPtr, sPtr)
		return
	}
	if s.dstTyp == s.srcTyp {
		ptr.Copy(dPtr, sPtr, s.size)
		return
	}
fcLoop:
	for _, v := range s.fieldConverters {
		fc, ok := v.(*fieldConverter)
		if !ok {
			continue
		}
		dAnonymousPtr := gslice.Or(fc.dAnonymousPtr)
		if !dAnonymousPtr && !gslice.Or(fc.sAnonymousPtr) {
			fc.convert(unsafe.Pointer(uintptr(dPtr)+gslice.Sum(fc.dOffset)), unsafe.Pointer(uintptr(sPtr)+gslice.Sum(fc.sOffset)))
		} else {
			fsPtr, fdPtr := unsafe.Pointer(uintptr(sPtr)+fc.sOffset[0]), unsafe.Pointer(uintptr(dPtr)+fc.dOffset[0])
			sOffset, dOffset := fc.sOffset[1:], fc.dOffset[1:]
			for i, isPtr := range fc.sAnonymousPtr {
				if isPtr {
					fsPtr = unsafe.Pointer(*((**int)(fsPtr)))
					if fsPtr == nil {
						if dAnonymousPtr {
							*(**int)(fdPtr) = nil
							continue fcLoop
						}
						fsPtr = fc.converter.sEmptyDereferValPtr
						break
					}
				}
				fsPtr = unsafe.Pointer(uintptr(fsPtr) + sOffset[i])
			}
			var i int
			var dNil bool
			for ; i < len(fc.dAnonymousPtr); i++ {
				if fc.dAnonymousPtr[i] {
					oldPtr := fdPtr
					fdPtr = unsafe.Pointer(*((**int)(fdPtr)))
					if dNil = fdPtr == nil; dNil {
						fdPtr = oldPtr
						break
					}
				}
				fdPtr = unsafe.Pointer(uintptr(fdPtr) + dOffset[i])
			}
			if dNil {
				v := unsafe.Pointer(uintptr(newValuePtr(fc.dStructType)) + dOffset[len(fc.dAnonymousPtr)-1])
				fc.convert(v, fsPtr)
				for j := len(fc.dAnonymousPtr) - 1; j >= i; j-- {
					v = unsafe.Pointer(uintptr(v) - dOffset[j])
					if fc.dAnonymousPtr[j] {
						v = unsafe.Pointer(gptr.Of(v))
					}
				}
				*(**int)(fdPtr) = *(**int)(v)
			} else {
				fc.convert(fdPtr, fsPtr)
			}
		}
	}
}

func (s *structConverter) mapConvert(dPtr, sPtr unsafe.Pointer) {
	dEmptyMapInterface := (*unsafeheader.EmptyInterface)(unsafe.Pointer(gptr.Of(reflect.New(s.convertType.dstTyp).Interface())))
	dv := ptrToMapValue(dEmptyMapInterface, dPtr)
	if dv.IsNil() {
		dv.Set(reflect.MakeMapWithSize(s.convertType.dstTyp, len(s.fieldConverters)))
	}
fcLoop:
	for _, v := range s.fieldConverters {
		fc, ok := v.(*fieldMapConverter)
		if !ok {
			continue
		}
		fsPtr := unsafe.Pointer(uintptr(sPtr) + fc.sOffset[0])
		sOffset := fc.sOffset[1:]
		for i, isPtr := range fc.sAnonymousPtr {
			if isPtr {
				fsPtr = unsafe.Pointer(*((**int)(fsPtr)))
				if fsPtr == nil {
					if s.option.IgnoreEmptyFields {
						continue fcLoop
					}
					fsPtr = fc.converter.sEmptyDereferValPtr
					break
				}
			}
			fsPtr = unsafe.Pointer(uintptr(fsPtr) + sOffset[i])
		}
		if s.option.IgnoreEmptyFields && reflect.DeepEqual(reflect.NewAt(fc.converter.sDereferType, fsPtr).Elem().Interface(), reflect.New(fc.converter.sDereferType).Elem().Interface()) {
			continue
		}
		dKey := reflect.ValueOf(fc.dName)
		dVal := reflect.New(fc.dType).Elem()
		fc.convert(unsafe.Pointer(dVal.UnsafeAddr()), fsPtr)
		dv.SetMapIndex(dKey, dVal)
	}
}

type fieldConverter struct {
	converter     *elemConverter
	sType         structItemType
	sOutType      funcOutType
	sFieldType    reflect.Type
	sStructType   reflect.Type
	dStructType   reflect.Type
	dAnonymousPtr []bool
	sAnonymousPtr []bool
	dOffset       []uintptr
	sOffset       []uintptr
	dName         string
	sName         string
	sFieldName    string
}

func (f *fieldConverter) convert(dPtr, sPtr unsafe.Pointer) {
	switch f.sType {
	case typeField:
		f.converter.convert(dPtr, sPtr)
	case typeFieldMethod, typeMethod:
		var method reflect.Value
		if f.sType == typeFieldMethod {
			sPtr = unsafe.Pointer(uintptr(sPtr) - gslice.Sum(f.sOffset))
			method = reflect.NewAt(f.sStructType, sPtr).Elem().FieldByName(f.sFieldName)
		} else {
			method = reflect.NewAt(f.sStructType, sPtr).MethodByName(f.sName)
		}
		if f.sType == typeMethod || !method.IsNil() {
			callback := method.Call(nil)
			value := callback[0]
			var vPtr unsafe.Pointer
			if value.CanAddr() {
				vPtr = unsafe.Pointer(value.UnsafeAddr())
			} else {
				vPtr = PtrOfAny(value)
			}
			if gvalue.In(value.Type().Kind(), reflect.Map, reflect.Pointer) {
				p := unsafe.Pointer(reflect.New(intDoubleRT).Pointer())
				*(**int)(p) = (*int)(vPtr)
				vPtr = p
			}
			switch f.sOutType {
			case boolOut:
				if !callback[1].Bool() {
					return
				}
			case errorOut:
				if !callback[1].IsNil() {
					return
				}
			}
			f.converter.convert(dPtr, vPtr)
		}
	}
}

func newFieldConverter(df, sf structItem, option *StructOption) *fieldConverter {
	format := gvalue.Valid(df.format, sf.format)
	if len(format) > 0 {
		option = gvalue.Safe(option.Clone())
		option.TimeFormat = format
	}
	ec, ok := newElemConverter(df.typ, sf.typ, option)
	if !ok {
		return nil
	}
	return &fieldConverter{
		converter:     ec,
		sType:         sf.itemType,
		sOutType:      sf.outType,
		sStructType:   sf.structType,
		dStructType:   df.structType,
		dAnonymousPtr: df.anonymousPtr,
		sAnonymousPtr: sf.anonymousPtr,
		dOffset:       df.offset,
		sOffset:       sf.offset,
		dName:         df.name,
		sName:         sf.name,
		sFieldName:    sf.filedName,
	}
}

type fieldMapConverter struct {
	converter     *elemConverter
	sAnonymousPtr []bool
	sOffset       []uintptr
	dName         string
	dType         reflect.Type
}

func (f *fieldMapConverter) convert(dPtr, sPtr unsafe.Pointer) {
	f.converter.convert(dPtr, sPtr)
}

func newFieldMapConverter(valueType reflect.Type, sf structItem, option *StructOption) *fieldMapConverter {
	if len(sf.format) > 0 {
		option = gvalue.Safe(option.Clone())
		option.TimeFormat = sf.format
	}
	ec, ok := newElemConverter(valueType, sf.typ, option)
	if !ok {
		return nil
	}
	return &fieldMapConverter{
		converter:     ec,
		sAnonymousPtr: sf.anonymousPtr,
		sOffset:       sf.offset,
		dName:         sf.name,
		dType:         valueType,
	}
}

func getFieldName(f reflect.StructField, opt *StructOption) string {
	name, tag := f.Tag.Get(opt.PriorityTagName), opt.PriorityTagName
	if len(name) == 0 {
		name, tag = f.Tag.Get(opt.TagName), opt.TagName
	}
	if len(name) == 0 {
		return f.Name
	}
	if tag != "json" {
		return name
	}
	if i := strings.Index(name, ","); i != -1 {
		return name[:i]
	}
	return name
}

type structItemType int64

const (
	typeField = iota + 1
	typeFieldMethod
	typeMethod
)

type funcOutType int64

const (
	singleOut = iota + 1
	boolOut
	errorOut
)

type structItem struct {
	itemType     structItemType
	outType      funcOutType
	name         string
	filedName    string
	format       string
	typ          reflect.Type
	structType   reflect.Type
	anonymousPtr []bool
	offset       []uintptr
}

func newStructItem(structType reflect.Type) *structItem {
	return &structItem{structType: structType}
}

func (s *structItem) setField(field reflect.StructField, anonymousPtr []bool, offset []uintptr) *structItem {
	s.itemType = typeField
	s.name = field.Name
	s.typ = field.Type
	s.anonymousPtr = anonymousPtr
	s.offset = offset
	return s
}

func (s *structItem) setFieldMethod(outType funcOutType, name, fieldName string, outRfType reflect.Type, anonymousPtr []bool, offset []uintptr) *structItem {
	s.itemType = typeFieldMethod
	s.outType = outType
	s.name = name
	s.filedName = fieldName
	s.typ = outRfType
	s.anonymousPtr = anonymousPtr
	s.offset = offset
	return s
}

func (s *structItem) setMethod(outType funcOutType, name string, outRfType reflect.Type) *structItem {
	s.itemType = typeMethod
	s.outType = outType
	s.name = name
	s.typ = outRfType
	return s
}

var protoPrivateField = set.New("state", "sizeCache", "unknownFields")

func extractFields(t reflect.Type, opt *StructOption, fieldMap map[string]*structItem, anonymousPtr []bool, offset ...uintptr) (fieldSlice []*structItem) {
	if opt == nil {
		opt = defaultStructOption()
	}
	isProto := isProtoMessage(t)
	if fieldMap == nil {
		fieldMap = make(map[string]*structItem)
	}
	anonymous := make([]*structItem, 0, t.NumField())
	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		if isProto && protoPrivateField.Contains(f.Name) {
			continue
		}
		sf := newStructItem(t)
		fieldName := f.Name
		if !opt.IgnoreTag {
			if f.Name = getFieldName(f, opt); f.Name == "-" {
				continue
			}
			sf.format = f.Tag.Get("format")
		}
		if !opt.IgnoreFunc && f.Type.Kind() == reflect.Func && f.Type.NumIn() == 0 {
			if outSize := f.Type.NumOut(); outSize == 1 {
				sf.setFieldMethod(singleOut, f.Name, fieldName, f.Type.Out(0), anonymousPtr, append(gslice.Clone(offset), f.Offset))
			} else if outSize == 2 {
				second := f.Type.Out(1)
				if second.Kind() == reflect.Bool {
					sf.setFieldMethod(boolOut, f.Name, fieldName, f.Type.Out(0), anonymousPtr, append(gslice.Clone(offset), f.Offset))
				} else if second.AssignableTo(errorRT) {
					sf.setFieldMethod(errorOut, f.Name, fieldName, f.Type.Out(0), anonymousPtr, append(gslice.Clone(offset), f.Offset))
				}
			}
		} else {
			sf.setField(f, anonymousPtr, append(gslice.Clone(offset), f.Offset))
		}
		if !opt.IgnorePrivateFields || unicode.IsUpper(rune(fieldName[0])) {
			if _, ok := fieldMap[f.Name]; ok {
				continue
			}
			fieldMap[f.Name] = sf
			fieldSlice = append(fieldSlice, sf)
		}
		if f.Anonymous {
			anonymous = append(anonymous, sf)
		}
	}
	for _, af := range anonymous {
		afTyp, isPtr := dereferencedTypeDeep(af.typ)
		s := extractFields(afTyp, opt, fieldMap, append(gslice.Clone(anonymousPtr), isPtr), af.offset...)
		for _, f := range s {
			fieldSlice = append(fieldSlice, f)
			fieldMap[f.name] = f
		}
	}
	return
}

func extractFieldsOnMap(t reflect.Type, opt *StructOption, fieldMap map[string]*structItem, anonymousPtr []bool, offset ...uintptr) (fieldSlice []*structItem) {
	if opt == nil {
		opt = defaultStructOption()
	}
	proto := isProtoMessage(t)
	anonymous := make([]*structItem, 0, t.NumField())
	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		if proto && protoPrivateField.Contains(f.Name) {
			continue
		}
		if f.Type.Kind() == reflect.Func {
			continue
		}
		sf := newStructItem(t)
		fieldName := f.Name
		if !opt.IgnoreTag {
			if f.Name = getFieldName(f, opt); f.Name == "-" {
				continue
			}
			sf.format = f.Tag.Get("format")
		}
		sf.setField(f, anonymousPtr, append(gslice.Clone(offset), f.Offset))
		if opt.IncludePrivateFields || unicode.IsUpper(rune(fieldName[0])) {
			if _, ok := fieldMap[f.Name]; ok {
				continue
			}
			fieldMap[f.Name] = sf
			fieldSlice = append(fieldSlice, sf)
		}
		if f.Anonymous {
			anonymous = append(anonymous, sf)
		}
	}
	for _, af := range anonymous {
		afTyp, isPtr := dereferencedTypeDeep(af.typ)
		s := extractFieldsOnMap(afTyp, opt, fieldMap, append(gslice.Clone(anonymousPtr), isPtr), af.offset...)
		for _, f := range s {
			fieldSlice = append(fieldSlice, f)
			fieldMap[f.name] = f
		}
	}
	return
}

func extractMethods(t reflect.Type, opt *StructOption, fieldMap map[string]*structItem) {
	if isProtoMessage(t) {
		return
	}
	if opt == nil {
		opt = defaultStructOption()
	}
	pv := reflect.New(t) // 要用指针类型，不然取不到挂在指针接收器下面的方法
	pt := pv.Type()
	for i, n := 0, pt.NumMethod(); i < n; i++ {
		name := pt.Method(i).Name
		if _, ok := fieldMap[name]; ok {
			continue
		}
		m := pv.MethodByName(name) // 用pv来获取Method，pt获取的会有1个入参（接收器变量），还需要额外判断逻辑
		if mt := m.Type(); mt.NumIn() == 0 {
			if outSize := mt.NumOut(); outSize == 1 {
				fieldMap[name] = newStructItem(t).setMethod(singleOut, name, mt.Out(0))
			} else if outSize == 2 {
				second := mt.Out(1)
				if second.Kind() == reflect.Bool {
					fieldMap[name] = newStructItem(t).setMethod(boolOut, name, mt.Out(0))
				} else if second.AssignableTo(errorRT) {
					fieldMap[name] = newStructItem(t).setMethod(errorOut, name, mt.Out(0))
				}
			}
		}
	}
}

var protoMessageType = ReflectType[proto.Message]()

func isProtoMessage(rt reflect.Type) bool {
	return rt.Implements(protoMessageType)
}
