package main

import (
	"fmt"
	"sync"
)

type counter struct {
	count int
	mu sync.Mutex
}

func (c *counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count++
}

func (c *counter) getCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.count
}

func main() {
	counter := &counter{}
	var wg sync.WaitGroup

	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.increment()
		}()
	}

	wg.Wait()

	fmt.Println(counter.getCount())
}