// Разработать программу, которая в рантайме способна определить 
// тип переменной: int, string, bool, channel из переменной типа interface{}.

// утверждениями типа (type assertions):
// switch v.(type) конструкция для обработки типов переменных интерфейса
// используется только с interface{} (пустым интерфейсом), так как
// именно в таких переменных может хранится значение произвольного типа
package main

import "fmt"

func typeAssertion(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Printf("%s - тип: %T\n", v, v)
	case int:
		fmt.Printf("%d - тип: %T\n", v, v)
	case bool:
		fmt.Printf("%v - тип: %T\n", v, v)
	case chan int:
		fmt.Printf("%v - тип: %T\n", v, v)
	}
}

func main() {
	str := "Hello"
	num := 42
	boolean := true
	ch := make(chan int)

	typeAssertion(str)
	typeAssertion(num)
	typeAssertion(boolean)
	typeAssertion(ch)
}