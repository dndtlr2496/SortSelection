package main

import "fmt"

func SelectionSort(arr []int, size int) []int {
	minIndex := 0
	for i := 0; i < size-1; i++ {
		minIndex = i
		for j := i + 1; j < size; j++ {
			if arr[j] < arr[minIndex] {
				temp := arr[minIndex]
				temp2 := arr[j]
				arr[minIndex] = temp2
				arr[j] = temp
			}
		}
	}
	return arr
}

func main() {
	arr := []int{3, 6, 2, 1}
	result := SelectionSort(arr, 4)
	fmt.Println(result)
}
