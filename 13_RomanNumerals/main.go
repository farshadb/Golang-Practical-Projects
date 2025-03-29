package main

import (
	"errors"
	"fmt"
	"strings"
)

var romanDic = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

var ErrInvalidNumber = errors.New("number out of range")

func ToRomanNumeral(i int) (string, error) {
	if i <= 0 || i >= 4000 {
		return "", ErrInvalidNumber
	}

	keys := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	var roman strings.Builder
	for _, k := range keys {
		for i >= k {
			roman.WriteString(romanDic[k])
			i -= k
		}
	}

	return roman.String(), nil
}

func main() {
	num := 3999
	RomanNum, err := ToRomanNumeral(num)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("%d in Roman numeric is: %s\n", num, RomanNum)
}
