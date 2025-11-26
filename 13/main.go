package main

import "fmt"

func main() {
	a := 121
	b := 365
	
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	a = a + b
	b = a - b
	a = a - b
	
	fmt.Println("Переменные после обмена значениями")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}