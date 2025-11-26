package main

import "fmt"

func reverseString(s string) string {
	runes := []rune(s)

	start := 0
	end := len(runes) - 1

	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}

	return string(runes)
}

func main() {
	str := "главрыба"
	reversedStr := reverseString(str)
	fmt.Println("строка:", str)
	fmt.Println("перевернутая строка:", reversedStr)
}