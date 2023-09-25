package main

import (
	"fmt"
	"sync"
)

func main() {
	// канал для записи данных в map
	dataCh := make(chan map[string]int)

	// WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	numWorkers := 5

	// Запускаем горутины для записи данных в канал
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			// Выполняем запись данных в канал
			data := make(map[string]int)
			key := fmt.Sprintf("Key%d", workerID)
			value := workerID * 10
			data[key] = value
			fmt.Printf("Worker %d wrote %s: %d\n", workerID, key, value)
			dataCh <- data
		}(i)
	}

	// Ожидаем завершения всех горутин и закрываем канал
	go func() {
		wg.Wait()
		close(dataCh)
	}()

	// Создаем map для хранения данных и собираем результаты из канала
	data := make(map[string]int)
	for workerData := range dataCh {
		for key, value := range workerData {
			data[key] = value
		}
	}

	// Выводим данные из map
	fmt.Println("Data in map:")
	for key, value := range data {
		fmt.Printf("%s: %d\n", key, value)
	}
}
