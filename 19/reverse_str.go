// Разработать программу, которая переворачивает подаваемую на ход строку 
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import (
    "fmt"
)

func reverseString(input string) string {
    // Преобразуем строку в срез рун
    runes := []rune(input)
    
    // Получим длину среза
    length := len(runes)

	left, right := 0, length-1;
    
    // Переворачиваем руны в срезе
    for left < right {
        runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
    }
    
    // Преобразуем руны обратно в строку
    return string(runes)
}

func main() {
    input1 := "главрыба"
    res1 := reverseString(input1)
    fmt.Printf("%s --> %s\n", input1, res1 )

	input2 := "Hello"
    res2 := reverseString(input2)
    fmt.Printf("%s --> %s\n", input2, res2 )
}

