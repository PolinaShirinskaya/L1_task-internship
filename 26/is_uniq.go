/* Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). 
   Функция проверки должна быть регистронезависимой.

	Например: 
	abcd — true
	abCdefAaf — false
	aabcd — false*/


package main

import (
	"fmt"
)

// IsUnique проверяет, что все символы в строке уникальные (регистронезависимо).
// использование типа rune позволяет корректно обрабатывать символы, которые занимают разное количество байтов
func IsUnique(s string) bool {
	charSet := make(map[rune]bool)

	for _, char := range s {
		if charSet[char] {
			return false
		}
		charSet[char] = true
	}

	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Printf("'%s': %v\n", str1, IsUnique(str1))
	fmt.Printf("'%s': %v\n", str2, IsUnique(str2))
	fmt.Printf("'%s': %v\n", str3, IsUnique(str3))
}
