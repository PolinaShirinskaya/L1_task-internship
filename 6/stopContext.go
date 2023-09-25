// Использование Context
// Мы можем передать контекст в горутину и отменить его, что бы завершить выполнение горутины

package main

import (
    "context"
    "fmt"
    "time"
)

func myGoroutine(ctx context.Context) {
    for {
        select {
		// Отслеживание сигнала об отмене (выполнении) контекста
        case <-ctx.Done():
			fmt.Println("Сигнал об отмене контекста")
            fmt.Println("Завершение...")
            return
        default:
            fmt.Println("MyGorutine работает")
            time.Sleep(time.Second)
        }
    }
}

func main() {
	// Создаем контекст и посылаем его в горутину для отслеживания
    ctx, cancel := context.WithCancel(context.Background())
    go myGoroutine(ctx)

    // Время для работы горутины перед завершением
    time.Sleep(5 * time.Second)

    // Отменяем контекст, чтобы завершить горутину
    cancel()

    // Время для отслеживания завершения горутины
    time.Sleep(2 * time.Second)
    fmt.Println("Основная горутина main завершает выполнение.")
}
