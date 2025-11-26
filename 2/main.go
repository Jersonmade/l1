package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	var resArr []int

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < len(nums); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			res := nums[i] * nums[i]
			
			mu.Lock()
			resArr = append(resArr, res)
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	for _, v := range resArr {
		fmt.Println(v)
	}
}