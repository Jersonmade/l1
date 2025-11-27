package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Неправильно переданы параметры")
		return
	}

	n, err := strconv.Atoi(os.Args[1])

	if err != nil || n <= 0 {
		fmt.Println("Колво воркеров должно быть положительным числом")
		return
	}

	ch := make(chan int)
	var wg sync.WaitGroup

	go func() {
		data := 0

		for {
			ch <- data
			data++
			time.Sleep(time.Second * 1)
		}
	}()

	for i := 1; i <= n; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			for data := range ch {
				fmt.Printf("Worker %d прочитал: %d\n", i, data)
			}
		}(i)
	}

	wg.Wait()
}