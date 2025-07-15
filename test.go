package main

import (
	"fmt"
)

func main() {
	fmt.Println(Splitter("кот"))
}

func Splitter(str string) []string {
	arr := []string{}

	runes := []rune(str)

	for i := 0; i <= len(runes)-1; i++ {
		arr = append(arr, string(runes[i]), string(runes[i+1]))
	}

	return arr
}
