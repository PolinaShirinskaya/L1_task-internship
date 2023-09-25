// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, 
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout

package main

import (
	"fmt"
)

func main() {
	// Создаем каналы для передачи чисел и их умножения
	dataCh := make(chan int)
	doubledCh := make(chan int)

	// Массив чисел для отправки в канал
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Исходный массив чисел:", numbers)

	// Отправляем числа в первый канал
	go func() {
		for _, num := range numbers {
			dataCh <- num
		}
		close(dataCh)
	}()

	// Запускаем горутину для чтения чисел и умножения
	go func() {
		for x := range dataCh {
			doubledCh <- x * 2
		}
		close(doubledCh)
	}()

	// Читаем результаты из второго канала и выводим их
	for doubled := range doubledCh {
		fmt.Println("Результат умножения:", doubled)
	}
}
