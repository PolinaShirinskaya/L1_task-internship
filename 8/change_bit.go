// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0

package main

import (
    "fmt"
)

func setBit(value int64, bitIndex uint) int64 {
    // Маска с установленным битом в позиции bitIndex
    bitIndex--
    mask := int64(1) << bitIndex

	res := value ^ mask
    return res
}

func main() {
    var num int64 = -42
    fmt.Printf("Певоначальное число: %d\n Битовое представление: %08b\n", num, num)


    // Изменяем i-й бит
    changeBit := setBit(num, 3)
    fmt.Printf("Измененный бит: %d\n Битовое представление: %08b\n", changeBit, changeBit)

}