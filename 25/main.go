package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	done := make(chan struct{})

	go func() {
		<-time.After(duration)
		close(done)
	}()

	<-done
}

func main() {
	fmt.Println("Begin")

	sleep(5 * time.Second)
	fmt.Println("End")
}
