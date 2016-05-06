// Copyright 2016 orivil Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package sorter for sort data with data priority.
package sorter

import "sort"

type IntMapSorter struct {
	intMap map[int]int
	keys []int
}

func (s *IntMapSorter) Len() int {
	return len(s.intMap)
}

func (s *IntMapSorter) Swap(i, j int) {
	s.keys[i], s.keys[j] = s.keys[j], s.keys[i]
}

func (s *IntMapSorter) Less(i, j int) bool {
	return s.intMap[s.keys[i]] < s.intMap[s.keys[j]]
}

func NewIntMapSorter(keyValue map[int]int) *IntMapSorter {
	_keys := make([]int, len(keyValue))
	index := 0
	for key, _ := range keyValue {
		_keys[index] = key
		index++
	}
	return &IntMapSorter{
		intMap: keyValue,
		keys: _keys,
	}
}

func SortReverseIntMap(data map[int]int) (keys []int) {
	sorter := NewIntMapSorter(data)
	return sorter.SortReverse()
}

func SortIntMap(data map[int]int) (keys []int) {
	sorter := NewIntMapSorter(data)
	return sorter.Sort()
}

func (s *IntMapSorter) Sort() (keys []int) {
	sort.Sort(s)
	return s.keys
}

func (s *IntMapSorter) SortReverse() (keys []int) {
	sort.Sort(sort.Reverse(s))
	return s.keys
}
