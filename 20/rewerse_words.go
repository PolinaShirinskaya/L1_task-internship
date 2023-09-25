// Разработать программу, которая переворачивает слова в строке. 
// Пример: «snow dog sun — sun dog snow».


package main

import (
	"fmt"
	"strings"
)

func reverseWordsOrder(input string) string {
	// Разбиваем строку на слова по пробелам --> срез строк
	words := strings.Fields(input)

	// Переворачиваем порядок слов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Собираем слова обратно в строку + пробелы
	reversed := strings.Join(words, " ")

	return reversed
}

func main() {
	input1 := "snow dog sun"
    res1 := reverseWordsOrder(input1)
    fmt.Printf("%s --> %s\n", input1, res1)

	input2 := "Hello world !"
    res2 := reverseWordsOrder(input2)
    fmt.Printf("%s --> %s\n", input2, res2)

}

