package main

import "fmt"

func main() {
	fmt.Println(mean([]int{1, 2, 3, 4, 5}))
	fmt.Println(mean([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

func mean(num []int) float64 {
	s := sum(num)
	return float64(s) / float64(len(num))
}

func sum(num []int) int {
	total := 0
	for _, i := range num {
		total += i
	}
	return total
}


