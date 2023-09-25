// Реализовать пересечение двух неупорядоченных множеств.


// При пересечении двух множеств формируем третье с элементами, 
// которые принадлежат и первому и второму множествам

package main

import (
	"fmt"
)

// Функция для вычисления пересечения двух множеств.
func intersection(set1, set2 map[int]bool) map[int]bool {
	result := make(map[int]bool)

	// Итерируемся по элементам первого множества.
	for element := range set1 {
		// Если элемент также присутствует во втором множестве, добавляем его в результат.
		if set2[element] {
			result[element] = true
		}
	}

	return result
}

func main() {
	// Создаем два неупорядоченных множества как мапы.

	set1 := map[int]bool{1: true, 2: true, 3: true, 4: true}
	set2 := map[int]bool{3: true, 4: true, 5: true, 6: true}

	// Вычисляем пересечение множеств.
	result := intersection(set1, set2)

	// Выводим результат.
	fmt.Println("Пересечение множеств:", result)
}
