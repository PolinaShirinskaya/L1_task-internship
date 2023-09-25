// Использование sync.WaitGroup
// sync.WaitGroup для ожидания завершения выполнения одной или нескольких горутин. 

package main

import (
    "fmt"
    "sync"
    "time"
)

func myGoroutine(wg *sync.WaitGroup) {
    // уменьшает внутренний счетчик активных горутин на один
	defer wg.Done()
    for i := 0; i < 5; i++ {
        fmt.Println("MyGorutine работает")
        time.Sleep(time.Second)
    }
	fmt.Println("Уменьшение счетчика актиных горутин")
}

func main() {
	// создание группы
    var wg sync.WaitGroup
	// добавдение горутины(элемента) в группу
    wg.Add(1)

    go myGoroutine(&wg)

    // Ожидание завершения горутины
	// Метод wg.Wait() деблокирует функцию main, 
	// когда внутренний счетчик элементов в группе wg станет равным 0
    wg.Wait()
	fmt.Println("wg.Wait деблокирует основную горутину main")
    fmt.Println("Основная горутина main завершает выполнение.")
}
