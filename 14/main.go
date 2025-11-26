package main

import (
	"fmt"
	"reflect"
)

func getType(param any) {
	switch v := param.(type) {
	case int:
		fmt.Printf("Целочисленный тип %T", v)
	case string:
		fmt.Printf("Строковый тип %T", v)
	case bool:
		fmt.Printf("Булевый тип %T", v)
	default:
		t := reflect.TypeOf(param)
		fmt.Printf("Канальный тип: %v", t)
	}

	fmt.Println()
}

func main() {
	a := 10
	str := "hkcbakcb"
	flag := true
	ch := make(chan bool)

	getType(a)
	getType(str)
	getType(flag)
	getType(ch)
}