package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(median([]float64{1, 2, 3, 4, 5}))
	fmt.Println(median([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

func median(num []float64) float64 {
	//work on a copy of the array
	vals := make([]float64, len(num))
	copy(vals, num)
	sort.Float64s(vals)

	i := len(vals) / 2
	if len(vals)%2 == 1 {
		return vals[i]
	}	
	return (vals[i-1] + vals[i]) / 2
}

