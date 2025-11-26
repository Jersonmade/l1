package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr) / 2]
	var left []int
	var middle []int
	var right []int

	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		} else {
			middle = append(middle, v)
		}
	}

	resultSlice := make([]int, 0, len(arr))
	resultSlice = append(resultSlice, quickSort(left)...)
	resultSlice = append(resultSlice, middle...)
	resultSlice = append(resultSlice, quickSort(right)...)

	return resultSlice
}

func main() {
	arr := []int{10, 4, 5, 7, 1, 8, 6, 2, 3, 9, 6}
	fmt.Println(quickSort(arr))
}