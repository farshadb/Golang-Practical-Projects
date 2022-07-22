package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(median([]int{1, 2, 3, 4, 5}))
	fmt.Println(median([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

func median(num []int) float64 {
	sort.Ints(num)
	if len(num)%2 == 0 {
		return float64(num[len(num)/2-1]+num[len(num)/2]) / 2
	} else {
		return float64(num[len(num)/2])
	}
}