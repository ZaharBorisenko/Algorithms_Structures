package main

import "fmt"

func main() {
	fmt.Println(SelectionSort([]int{8, 76, 4, 98, 3, 7, 2, 9, 866, 1}))
}

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		lowest := i
		for j := i + 1; j < len(arr); j++ {
			if arr[lowest] > arr[j] {
				lowest = j
			}
		}
		arr[i], arr[lowest] = arr[lowest], arr[i]
	}
	return arr
}
