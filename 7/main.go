package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	data := make(map[string]int)

	var wg sync.WaitGroup
	var mtx sync.Mutex

	initTime := time.Now()

	for i := 0; i < 2_000; i++ {
		wg.Add(1)
		go func(key string, value int) {
			defer wg.Done()
			mtx.Lock()
			data[key] = value
			mtx.Unlock()
		}(fmt.Sprintf("Key%d", i), i)
	}

	wg.Wait()

	fmt.Println("Время прошло:", time.Since(initTime))

	fmt.Println(len(data))
}