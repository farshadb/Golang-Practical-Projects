package maptool

import (
	"fmt"
	"sort"
)

func SortMap(m map[string]int) {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}