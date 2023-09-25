// Использования runtime.Goexit()
// Выполняет завершение только текущей горутины

package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()

    fmt.Printf("Воркер %d начал\n", id)
    defer fmt.Printf("Воркер %d закончил\n", id)

    for i := 0; i < 3; i++ {
        fmt.Printf("Воркер %d: Таска %d\n", id, i)
        time.Sleep(time.Second)
    }

    // Завершает только текущую горутину
    runtime.Goexit()
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    wg.Wait()

    fmt.Println("Все воркеры завершили работу.")
}
