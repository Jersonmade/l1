package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	to := time.After(time.Duration(5) * time.Second)

	go func () {
		i := 1
		for {
			ch <- i
			i++
			time.Sleep(500 * time.Millisecond)
		}
	}()

	for {
		select {
		case <- to:
			fmt.Println("Программа завершилась")
			return
		case num := <- ch:
			fmt.Println(num)
		}
	}
}