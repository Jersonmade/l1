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

		switch t.Kind() {
		case reflect.Chan:
			elemType := t.Elem()
			fmt.Printf("Канальный тип: chan %s", elemType)
		case reflect.Slice:
			elemType := t.Elem()
			fmt.Printf("Слайс: []%s", elemType)
		case reflect.Map:
			keyType := t.Key()
			elemType := t.Elem()
			fmt.Printf("Хеш-таблица: map[%s]%s", keyType, elemType)
		default:
			fmt.Printf("Неизвестный тип %s", t)
		}
	}

	fmt.Println()
}

func main() {
	a := 10
	str := "hkcbakcb"
	flag := true
	ch := make(chan bool)
	var slice []int
	m := make(map[string]int)

	getType(a)
	getType(str)
	getType(flag)
	getType(ch)
	getType(slice)
	getType(m)
}