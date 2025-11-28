package main

import "fmt"

func removeElemFromSlice(slice []int, index int) []int {
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice) - 1]
}

func main() {
	sl := []int{10, 20, 30, 40, 50, 60}
	fmt.Println(removeElemFromSlice(sl, 3))
}