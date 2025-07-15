package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	target := 7
	index := BinarySearch(arr, target)
	fmt.Printf("BinarySearch: %d found at index %d\n", target, index)
}

func BinarySearch(arr []int, target int) int {
	//two pointers
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		}

	}

	return -1
}
