package main

import (
	// "context"
	// "fmt"
	// "time"
)

// 1. Остановка через канал уведомления:
// func worker(stop <-chan bool) {
// 	for {
// 		select {
// 		case <- stop:
// 			fmt.Println("Горутина остановлена")
// 			return
// 		default:
// 			fmt.Println("Work")
// 			time.Sleep(1 * time.Second)
// 		}
// 	}
// }

// 2. Остановка через контекст
// func worker(ctx context.Context) {
// 	for {
// 		select {
// 		case <- ctx.Done():
// 			fmt.Println("Горутина остановлена")
// 			return
// 		default:
// 			fmt.Println("Work")
// 			time.Sleep(1 * time.Second)
// 		}
// 	}
// }

// 3. Остановка через runtime.Goexit()
// func worker() {
// 	fmt.Println("work")
// 	time.Sleep(2 * time.Second)
// 	runtime.Goexit()
// }

// 4. Выход по условию
// func worker() {
// 	var counter int

// 	for {
// 		if counter >= 5 {
// 			fmt.Println("Горутина завершена")
// 			return
// 		}

// 		fmt.Println("Work")
// 		counter++
// 		time.Sleep(1 * time.Second)
// 	}
// }

func main() {
	// 1
	// stop := make(chan bool)
	// go worker(stop)

	// time.Sleep(5 * time.Second)
	// stop <- true

	// 2
	// ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	// defer cancel()

	// go worker(ctx)

	// 3
	// go worker()

	// 4
	// go worker()

	// time.Sleep(7 * time.Second)
	// fmt.Println("Программа завершена")
}
