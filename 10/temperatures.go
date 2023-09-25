/* Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. 
   Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

   Пример: -20:{-25.4, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc. */

   package main

   import (
	   "fmt"
	   "math"
	   "strings"
	   "sort"
   )
   
   func groupTemperatures(temperatures []float64) map[int][]float64 {
	   groupedTemperatures := make(map[int][]float64)
   
	   for _, temp := range temperatures {
			var group int

			// Отличается способ округления + и - чисел
			if temp > 0 {
				// Определяем группу с шагом 10 градусов и добавляем в определенную группу
				group = int(math.Floor(temp/10.0)) * 10
			} else  {
				// Определяем группу с шагом 10 градусов и добавляыем в группу
				group = int(math.Ceil(temp/10.0)) * 10
			}
			groupedTemperatures[group] = append(groupedTemperatures[group], temp)
	   }
   
	   return groupedTemperatures
   }
   
   func formatTemperature(temp float64) string {
	   // Форматируем температуру для вывода с одним десятичным знаком
	   return fmt.Sprintf("%.1f", temp)
   }
   
   func main() {
	   temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
   
	   result := groupTemperatures(temperatures)

		// Сортируем ключи групп
		var sortedKeys []int
		for key := range result {
			sortedKeys = append(sortedKeys, key)
		}
		sort.Ints(sortedKeys)

		// Выводим результаты в порядке возрастания ключей групп
		for _, group := range sortedKeys {
			// Преобразуем значения в строки и объединяем их с запятой
			formattedValues := make([]string, len(result[group]))
			for i, val := range result[group] {
				formattedValues[i] = formatTemperature(val)
			}
			valuesStr := strings.Join(formattedValues, ", ")

			fmt.Printf("%d: {%s}\n", group, valuesStr)
		}

		//fmt.Printf("resul of splitting by groups %v\n", result)

}
   