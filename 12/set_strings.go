// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество

package main

import "fmt"

// Создадим множутсво, при этом удалив все дубликаты значений
func makeSet(arr []string) []string {
	var result []string
	setOfString := make(map[string]bool, len(arr))

	// добавляем элементы в множество по ключу, т.о. значения будут уникальны
	for _, val := range arr {
		setOfString[val] = true
	}

	for key := range setOfString {
		result = append(result, key)
	}

	return result
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	setArr := makeSet(arr)

	fmt.Printf("Последовательность: %v\nМножество: %v\n", arr, setArr)
}