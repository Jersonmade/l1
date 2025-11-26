package main

import "fmt"

type Human struct {
	name         string
	surname      string
	age          int
	email        string
	phone_number string
}

func (h *Human) setName(name string) {
	h.name = name
}

func (h *Human) getName() string {
	return h.name
}

type Action struct {
	Human
	something string
}

func main() {
	h1 := Human{
		name:         "Maxim",
		surname:      "Soldatov",
		age:          21,
		email:        "mail@inbox.ru",
		phone_number: "89291234567",
	}

	action1 := Action{h1, "something"}
	action1Name := action1.getName()
	fmt.Println(action1)
	fmt.Println()
	fmt.Println(action1Name)
	fmt.Println()

	action1.setName("Max")
	action1NameUpdated := action1.getName()
	fmt.Println(action1)
	fmt.Println()
	fmt.Println(action1NameUpdated)
}
