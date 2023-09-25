// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. 
// По завершению программа должна выводить итоговое значение счетчика.


package main

import (
    "fmt"
    "sync"
)

// Мьютекс в структуре счетчика используется для предотвращения одновременного 
// доступа к инкрементации и чтению (получению значения) счетчика 
type Counter struct {
    mu      sync.Mutex
    value   int
}

// Метод для инкремнтации счетчика
func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

// Getter для значения счетчика
func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}

func main() {
    // Создаем экземпляр счетчика
    counter := &Counter{}

    // Создаем WaitGroup для синхронизации завершения горутин
    var wg sync.WaitGroup

    // Запускаем 10 горутин для инкрементации счетчика
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
			// Каждая горутина увеличит значение счетчика на 1000
            for j := 0; j < 1000; j++ {
                counter.Increment()
            }
        }()
    }

    // Ожидаем завершения всех горутин
    wg.Wait()

    // Выводим итоговое значение счетчика
    fmt.Printf("Итоговое значение счетчика: %d\n", counter.Value())
}
