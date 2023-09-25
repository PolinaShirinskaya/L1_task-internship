// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные 
// из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.



//----------CONTEXT--------------
// Используется контекст для управления жизненным циклом горутин (воркеров). 
// При получении сигнала о прерывании выполнения или SIGTERM, 
// контекст отменяется, что приводит к прекращению работы и главной горутины и каждого воркера. 
// В предоставленном коде, ctx.Done() используется в блоке select в каждом воркере и при записи данных в канал. 

// Когда вызывается функция cancel() (которая возвращается из context.WithCancel(parentContext)),
// она отправляет сигнал всем горутинам, которые "слушают" этот контекст, то есть проверяют ctx.Done().
// Если горутина видит, что ctx.Done() - это закрытый канал, она знает, что был получен сигнал о выполнении, 
// и она должна остановить свою работу.

package main

import (
    "context"
    "fmt"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "time"
)

func main() {
    const N = 3 // Варьирование количества воркеров

    // Создаем контекст и канал для сигнала завершения
    ctx, cancel := context.WithCancel(context.Background())
    signalCh := make(chan os.Signal, 1)

    // Обрабатываем сигналы интеррупции и SIGTERM для отмены контекста
    signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-signalCh
        cancel()
    }()

    // Создаем канал для передачи данных
    dataCh := make(chan string)

    // Создаем waitgroup для воркеров
    var wg sync.WaitGroup
    wg.Add(N)

    // Создаем и запускаем воркеров в цикле
    for i := 0; i < N; i++ {
        go worker(ctx, i, dataCh, &wg)
    }

    // Главный поток записывает данные в канал
    go writeDataToChannel(ctx, dataCh)

    // Ожидаем завершение работы воркеров
    wg.Wait()
    fmt.Println("Программа завершена")
}

func worker(ctx context.Context, id int, dataCh <-chan string, wg *sync.WaitGroup) {
    defer wg.Done()

    // Читаем данные из канала и выводим в stdout пока не получим сигнал о выполнении контекста
    for {
        select {
        case <-ctx.Done():
            return
        case data := <-dataCh:
            fmt.Printf("Worker %d: %s\n", id, data)
        }
    }
}

func writeDataToChannel(ctx context.Context, dataCh chan<- string) {
    // Устанавливаем таймер для записи данных каждые 500 миллисекунд
    ticker := time.Tick(500 * time.Millisecond)
    i := 0

    // Бесконечно записываем данные в канал пока не получим сигнал о выполнении контекста
    for {
        select {
        case <-ctx.Done():
            return
        case <-ticker:
            dataCh <- fmt.Sprintf("Data %d", i)
            i++
        }
    }
}
