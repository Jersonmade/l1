package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	numsCh := make(chan int)
	squareNumsCh := make(chan int)
	
	go func() {
		for _, v := range arr {
			numsCh <- v
		}
		close(numsCh)
	}()

	go func() {
		for v := range numsCh {
			squareNumsCh <- v * 2
		}
		close(squareNumsCh)
	}()

	for v := range squareNumsCh {
		fmt.Println(v)
	}
}