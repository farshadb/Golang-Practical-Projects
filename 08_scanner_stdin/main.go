package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

)

func main () {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type number of games:")
	// ensoure that new line or /n is read and it type is bool
	scanner.Scan()
	fmt.Println(scanner.Text())

	fmt.Printf("%T - %s", scanner.Text(), scanner.Text())
	
	input,err := strconv.ParseInt(scanner.Text(), 10, 64)

	if err != nil {
		fmt.Println(err)
		log.Fatal("Input is not a number")
	}

	fmt.Println("Type the result of matches for each game:", input)
}