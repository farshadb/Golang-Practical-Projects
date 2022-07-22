package maptool

import (
	"fmt"
	"sort"
)

func SortMap(m map[string]int) {
	var keys []string
	for k := range m {
		keys = append(keys, k)
		//fmt.Println("key:", k, "value:", m[k])
	}
	
	for k, v := range m {
		fmt.Println("key:", k, "value:", v)
		
	}
	sort.Strings(keys)
}