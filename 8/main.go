package main

import "fmt"

func setBitValue(num int64, i int, value int) int64 {
	if value == 1 {
		return num | (1 << i - 1)
	} else {
		return num &^ (1 << i - 1)
	}
}

func main() {
	var num int64 = 5

	result := setBitValue(num, 1, 0)
	fmt.Printf("5 (0101) -> %d (%04b)\n", result, result)
}