package main

import "fmt"

func main() {
	fmt.Println()
}

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		isSwapped := false
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				isSwapped = true
				//старый подход с созданием доп.переменной в некоторых языках
				//tmp := arr[j]
				//arr[j] = arr[j+1]
				//arr[j+1] = tmp
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if !isSwapped {
			break
		}
	}
	return arr
}
