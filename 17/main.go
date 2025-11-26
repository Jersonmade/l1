package main

import "fmt"

func binarySearch(arr []int, target int) int {
	if len(arr) < 2 {
		return 0
	}

	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high - low) / 2

		if target == arr[mid] {
			return mid
		} else if target < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	arr := []int{12, 45, 46, 78, 79, 85, 112, 115}
	target := 78

	fmt.Println(binarySearch(arr, target))
}