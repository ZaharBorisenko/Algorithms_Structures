package main

import "fmt"

func main() {
	fmt.Println(insertionSort([]int{8, 76, 4, 98, 3, 7, 2, 9, 866, 1, 7, 2, 9, 866, 1}))
}

func insertionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}

	return arr
}
