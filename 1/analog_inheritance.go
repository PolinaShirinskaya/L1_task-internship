// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской
// структуры Human (аналог наследования).

package main

import (
	"fmt"
)

// Родительская структура Human
type Human struct {
	Name string
}

func (h *Human) Mood() {
	fmt.Printf("%s is happy.\n", h.Name)
}

func (h *Human) Work() {
	fmt.Printf("%s works at office.\n", h.Name)
}

// Определим структуру Action
// и встроим методы от родительской структуры Human
type Action struct {
	*Human
	Location string
}

func (a *Action) MoveTo(location string) {
	a.Location = location
	fmt.Printf("%s is moving %s.\n", a.Name, location)
}

func main() {
	// Создаем экземпляр Human.
	h := &Human{
		Name: "John",
	}

	// Создаем экземпляр Action, который также имеет доступ к методам Human.
	action := Action{
		Human:    h,
		Location: "Home",
	}

	// Теперь можем вызывать методы Eat, Sleep и MoveTo через экземпляр action.
	fmt.Println("__Методы от родительской структуры Human:")
	action.Mood()
	action.Work()

	// Мы можем также использовать собственный метод структуры action
	fmt.Println("__Собственный метод структуры Action:")
	action.MoveTo("home")
}
