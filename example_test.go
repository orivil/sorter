package sorter_test

import (
	"fmt"
	"gopkg.in/orivil/sorter.v0"
)

var data = map[string]int{
	"home": 0,
	"login": 1,
	"logout": 2,
	"register": 3,
}

func ExamplePrioritySorter() {
	// 1. new sorter
	sorter := sorter.NewPrioritySorter(data)

	// 2. sort
	keys := sorter.Sort()
	fmt.Println(keys)

	// 3. reverse sort
	keys = sorter.SortReverse()
	fmt.Println(keys)

	// Output:
	// [home login logout register]
	// [register logout login home]
}

func ExampleIntMapSorter() {
	m := map[int]int{
		1: 33,
		2: 22,
		3: 11,
	}

	// 1. new sorter
	sorter := sorter.NewIntMapSorter(m)

	// 2. sort
	keys := sorter.Sort()
	fmt.Println(keys)

	// 3. reverse sort
	keys = sorter.SortReverse()
	fmt.Println(keys)

	// Output:
	// [3 2 1]
	// [1 2 3]
}
