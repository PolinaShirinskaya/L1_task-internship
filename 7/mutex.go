// Реализовать конкурентную запись данных в map

// Реализация конкурентной записи с помощью мьютексов:

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Мьютекс для защиты доступа к map
	var mute sync.Mutex

	// map для хранения данных
	data := make(map[string]int)

	// WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	numWorkers := 5

	// Запускаем горутины для записи данных в map
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			// Защищаем доступ к map с помощью мьютекса
			mute.Lock()
			defer mute.Unlock()

			// Выполняем запись данных в map
			key := fmt.Sprintf("Key%d", workerID)
			value := workerID * 10
			data[key] = value
			fmt.Printf("Worker %d wrote %s: %d\n", workerID, key, value)
		}(i)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим данные из map
	fmt.Println("Data in map:")
	for key, value := range data {
		fmt.Printf("%s: %d\n", key, value)
	}
}
