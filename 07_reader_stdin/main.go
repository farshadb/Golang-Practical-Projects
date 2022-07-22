package main

import (
	"fmt"
	"strconv"
	"strings"
	"bufio"
	"os"
	"log"
)

func main() {
	fmt.Print("Enter number of games: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		log.Fatal("error while reading input")
	}
	
	numInput := strings.TrimSpace(input)
	n, err := strconv.ParseInt(numInput, 10, 64)

	if err != nil {
		fmt.Println(err)
		log.Fatal("Input is not a number")
	}

	fmt.Println("Enter the result of matches for each game:")
	//game1 := GetInputSlice()

	fmt.Println(n)
}