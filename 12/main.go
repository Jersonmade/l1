package main

import "fmt"

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree", "dog", "dog", "horse", "tree"}
	m := make(map[string]struct{})

	for _, v := range slice {
		m[v] = struct{}{}
	}

	var set []string

	for k := range m {
		set = append(set, k)
	}

	fmt.Println(set)
}