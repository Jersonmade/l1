package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := []int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	m := make(map[int]struct{})

	var res []int

	for _, v := range slice1 {
		m[v] = struct{}{}
	}

	for _, v := range slice2 {
		if _, ok := m[v]; ok {
			res = append(res, v)
		}
	}

	fmt.Println(res)
}