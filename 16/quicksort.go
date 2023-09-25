// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import (
	"fmt"
	"strconv"
)

// quickSort выполняет сортировку массива arr с помощью алгоритма быстрой сортировки
func quickSort(arr []int, low int, high int) {
	if low < high {
		pi := partion(arr, low, high) // Вызываем функцию partion, чтобы разделить массив и вернуть индекс опорного элемента

		quickSort(arr, low, pi-1)  // Рекурсивно сортируем левую половину массива
		quickSort(arr, pi+1, high) // Рекурсивно сортируем правую половину массива
	}
}

// partion выполняет разделение массива и возвращает индекс опорного элемента
// меняет местами элементы таким образом, чтобы все элементы, меньшие опорного, находились слева от него.
func partion(arr []int, low int, high int) int {
	pivot := arr[high] // Опорный элемент - последний элемент массива
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			// Меняем элементы местами, чтобы все элементы, меньшие опорного, находились слева от него
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Помещаем опорный элемент на правильную позицию
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// printArray выводит элементы массива arr
func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(strconv.Itoa(arr[i]) + " ")
	}
	fmt.Println("")
}

func main() {
	var arr = []int{26, 3, 1, 6, -8, 9, 0}

	fmt.Print("До сортировки: ")
	printArray(arr)

	quickSort(arr, 0, len(arr)-1)

	fmt.Print("После сортировки: ")
	printArray(arr)
}
