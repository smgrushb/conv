// Copyright 2025 smgrushb
// Licensed under the Apache License, Version 2.0
// https://www.apache.org/licenses/LICENSE-2.0
// Inspired by coven (MIT License) by petersunbag

package model

type KeyValuePair[K comparable, V any] struct {
	Key   K `bson:"key" gorm:"column:key" json:"key" form:"key" query:"key"`
	Value V `bson:"value" gorm:"column:value" json:"value" form:"value" query:"value"`
}

func NewKeyValuePair[K comparable, V any]() *KeyValuePair[K, V] {
	return &KeyValuePair[K, V]{}
}

func (k *KeyValuePair[K, V]) GetKey() K {
	return k.Key
}

func (k *KeyValuePair[K, V]) SetKey(key K) {
	k.Key = key
}

func (k *KeyValuePair[K, V]) GetValue() V {
	return k.Value
}

func (k *KeyValuePair[K, V]) SetValue(value V) {
	k.Value = value
}
