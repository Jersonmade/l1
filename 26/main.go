package main

import (
	"fmt"
	"strings"
)

func checkDuplicate(str string) bool {
	m := make(map[rune]bool)

	for _, val := range str {
		if _, ok := m[val]; ok {
			return true
		}

		m[val] = true
	}

	return false
}

func main() {
	str := "fasS"
	lowerCaseStr := strings.ToLower(str)

	fmt.Println(lowerCaseStr)

	if isDuplicate := checkDuplicate(lowerCaseStr); isDuplicate {
		fmt.Println("Строка содержит дубликаты")
	} else {
		fmt.Println("Строка не содержит дубликаты")
	}
}