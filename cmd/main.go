// Задание 18.2.1
// Создайте программу, которая запускает 5 рутин,
// каждая из которых печатает свой порядковый номер 10 раз.
// Добиться такого результата за один вызов wg.Add.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	const routineNumber = 5
	const printNumber = 10

	for i := 1; i <= routineNumber; i++ {
		wg.Add(1)
		go func(routine int) {
			defer wg.Done()

			for i := 0; i < printNumber; i++ {
				fmt.Println(routine)
			}
		}(i)
	}
	wg.Wait()
}
