package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	target := 7
	index := LinearSearch(arr, target)
	fmt.Printf("LinearSearch: %d found at index %d\n", target, index)
}

func LinearSearch(arr []int, target int) int {
	for i := 0; i < len(arr)-1; i++ {
		el := arr[i]
		if el == target {
			return i
		}
	}
	return -1
}
