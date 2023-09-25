// Реализовать бинарный поиск встроенными методами языка.

package main
import "fmt"

// Функция binarySearch реализует алгоритм бинарного поиска
// и возвращает true, если искомый элемент найден в массиве haystack,
// и false в противном случае.
func binarySearch(needle int, haystack []int) bool {

	low := 0
	high := len(haystack) - 1

	// Цикл выполняется пока low меньше или равно high
	for low <= high {
		median := (low + high) / 2

		// Если элемент в середине массива меньше needle,
		// обновляем значение low на median + 1
		if haystack[median] < needle {
			low = median + 1
		} else {
			// В противном случае обновляем значение high на median - 1
			high = median - 1
		}
	}

	// Проверяем, если low равно длине массива или значение haystack[low] 
	// не равно needle, то возвращаем false
	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	// Возвращаем true, если искомый элемент найден
	return true
}


func main() {
	sortedArr := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(63, sortedArr))
}
