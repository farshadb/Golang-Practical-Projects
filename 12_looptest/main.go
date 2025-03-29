package main

import "fmt"

func main() {
	var sum int64
	var i int64
	for i = 0; i <= 10000000000000; i++ {
		sum += i

	}
	fmt.Println(sum)
}
