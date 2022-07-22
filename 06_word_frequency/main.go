package main

import (
	"Golang-Practical-Project/06_word_frequency/maptool"
	"fmt"
	//"github.com/F4r5h4d/Golang-Practical-Projects/06_word_frequency/maptool"
)

func main() {
	
	words := []string{ 
		"Recent", "analyses", "from", "FiveThirtyEight", "Cook", "Political", "House",
		"from", "FiveThirtyEight", "Cook", "Political", "House", "Democrats", "Republicans",
		"Democrats", "Republicans", "Democrats", "Republicans", "Democrats", "Republicans",
		"Democrats", "Republicans", "Democrats", "Republicans", "Democrats", "Republicans",
	}
	fmt.Println(frequency(words))
}

func frequency(word []string) map[string]int {
	freq := make(map[string]int)
	for _, w := range word {
		freq[w]++
	}
	maptool.SortMap(freq)
	return freq
}

