// Разработать программу, которая в рантайме способна определить 
// тип переменной: int, string, bool, channel из переменной типа interface{}.


/* В данном варианте решения задачи используется пакет 'reflect', а именно
  метод reflect.TypeOf()
  Переменная типа интерфейса хранит пару: конкретное значение, присвоенное переменной, и тип этого элемента.
  На базовом уровне reflect является всего лишь механизмом для изучения пары тип и значение, 
  хранящейся внутри переменной интерфейса. */
package main

import (
  "fmt"
  "reflect"
)

func main() {
    var f float64 = 3.4
    fmt.Println("Переменная:", 3.4)
    // Когда мы вызываем reflect.TypeOf(x), x сначала сохраняется в пустом интерфейсе, 
    // который затем передается в качестве аргумента; reflect.TypeOf() распаковывает этот пустой интерфейс 
    // для восстановления информации о типе.
    fmt.Printf("Тип: %s\n\n", reflect.TypeOf(f))

    var s string = "Hello!"
    fmt.Println("Переменная:", s)
    fmt.Printf("Тип: %s\n\n", reflect.TypeOf(s))

    var b bool = true
    fmt.Println("Переменная:", b)
    fmt.Printf("Тип: %s\n\n", reflect.TypeOf(b))

    var ch chan int
    ch = make(chan int)

    fmt.Println("Переменная:", ch)
    fmt.Println("Тип:", reflect.TypeOf(ch))
    
}


// https://habr.com/ru/articles/415171/