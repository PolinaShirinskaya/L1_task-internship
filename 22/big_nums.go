// Разработать программу, которая перемножает, делит, складывает, вычитает 
// две числовых переменных a,b, значение которых > 2^20

package main

import (
	"fmt"
	"math/big"
)

// a + b
func add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

// a - b
func sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

// a * b
func multiply(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

// a / b
func divide(a, b *big.Int) *big.Int {
	return new(big.Int).Div(b, a)
}

func main() {
	// создадим два больших числа, NewInt принимает int64 и возвращает big.Int
	bigIntA := big.NewInt(int64(2 << 21))
	// так как число может быть настолько большим, что не поместиться в int64 - попробуем еще один вариант
	bigIntB := new(big.Int)
	bigIntB.SetString("60000000000000000", 10)

	fmt.Println(bigIntA)
	fmt.Println(bigIntB)

	fmt.Println("Сложение: ", add(bigIntA, bigIntB))

	fmt.Println("Вычитание: ", sub(bigIntA, bigIntB))

	fmt.Println("Умножение: ", multiply(bigIntA, bigIntB))

	fmt.Println("Деление: ", divide(bigIntA, bigIntB))

}