// Реализовать собственную функцию sleep.

// https://tuzov.su/posts/как-устроена-функция-time-sleep/
// https://github.com/golang/go/blob/release-branch.go1.19/src/runtime/time.go#L178

package main

import (
	"fmt"
	"time"
)

func Sleep(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	start := time.Now()
	fmt.Println("Начало выполнения программы")

	fmt.Println("Программа будет ожидать 3 секунды")
	Sleep(3)

	fmt.Println("Программа завершилась за: ", time.Since(start))
}