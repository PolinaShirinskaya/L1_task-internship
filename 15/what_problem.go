/*К каким негативным последствиям может привести данный фрагмент кода, 
  и как это исправить? Приведите корректный пример реализации.

	var justString string
	func someFunc() {
		v := createHugeString(1 << 10)
		justString = v[:100]
	}

	func main() {
		someFunc()
	}
*/


/* 1) в приведенном в задании коде  подстрока justString ссылается на те же байты данных, что и v, пока она существует. 
 и если в дальнейшем изменить исходную строку v, это также отразится на подстроке justString.
 Лучше создать justString, как новый срез и скопировать в енго первые 100 байт из v

   2) так же использование рун может быть наиболее безопасным способом работы с символами
   
   3) могут возникнуть проблемы с использованием глобальных переменных,
к глобальным переменным можно обращаться из любой функции, что не очень безопасно для данных.
Так же ГП создаются на куче, что может, а ЛП на стеке, работа на стеке быстрее */

package main

import (
	"fmt"
)

var justString string

func createHugeString(size int) string {
	// Создаем строку заданного размера
	runes := make([]rune, size)
	// Заполняем строку данными
	for i := 0; i < size; i++ {
		runes[i] = 'a'
	}
	return string(runes)
}

func someFunc() {
	v := createHugeString(1 << 10)

	// Выделение памяти для нового среза размером 100 рун
	justStringRunes := make([]rune, 100)

	// Копирование первых 100 рун из огромной строки в новый срез
	copy(justStringRunes, []rune(v[:100]))

	// Преобразование среза рун обратно в строку
	justString = string(justStringRunes)

	// Проверка
	if justString == v[:100] {
		fmt.Println("justString содержит первые 100 символов из v.")
	} else {
		fmt.Println("justString не содержит первые 100 символов из v.")
	}
}

func main() {
	someFunc()
	fmt.Println(justString)
}
