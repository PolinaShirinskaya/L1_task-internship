// Написать программу, которая конкурентно рассчитает значение квадратов 
// чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
    "fmt"
    "sync"
)

func main() {
    numbers := []int{2, 4, 6, 8, 10}
	// тип позволяет определить группу горутин
    var wg sync.WaitGroup

    for _, num := range numbers {
		// добавление горутины(элемента) в группу
        wg.Add(1)
        go func(n int) {
			// уменьшает внутренний счетчик активных элементов на один
            defer wg.Done()
            square := n * n
            fmt.Printf("%d squared = %d\n", n, square)
        }(num)
    }

    // Т. к. выполнение main и square выполняются фактически одновременно, нам необходимо сделать корректное завершение main,
    // во избежании неопределенного поведения нашей программы.
	// Метод wg.Wait() деблокирует функцию main, когда внутренний счетчик элементов в группе wg станет равным 0
    wg.Wait()
}
