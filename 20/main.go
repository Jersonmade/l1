package main

import (
	"fmt"
)

func reverseWordsInSentence(sentence string) string {
	runes := []rune(sentence)

	for i, j := 0, len(runes) - 1; i < j; {
		runes[i], runes[j] = runes[j], runes[i]
		i++ 
		j--
	}

	start := 0
	for k := 0; k <= len(runes) - 1; k++ {
		if k == len(runes) - 1 || runes[k+1] == ' ' {
			end := k

			for start < end {
				runes[start], runes[end] = runes[end], runes[start]
				start++
				end--
			}

			start = k + 2
		} 
	}

	return string(runes)
}

func main() {
	sentence := "snow dog sun"
	fmt.Println(sentence)
	fmt.Println(reverseWordsInSentence(sentence))
}