package main

import "fmt"

func main() {
	fmt.Println(SelectionSort([]int{8, 76, 4, 98, 3, 7, 2, 9, 866, 1}))
}

func SelectionSort(arr []int) []int {

	for i := len(arr) - 1; i > 0; i-- {
		maxI := i
		for j := 0; j < i; j++ {
			if arr[j] > arr[maxI] {
				maxI = j
			}
		}

		arr[i], arr[maxI] = arr[maxI], arr[i]
	}

	return arr
}
