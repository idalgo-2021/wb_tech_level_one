// Конкурентная запись в map
//
// * Реализовать безопасную для конкуренции запись данных в структуру map.
// * Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).
// * Проверьте работу кода на гонки (util go run -race).
///////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := make(map[int]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				mu.Lock()
				m[id*100+j] = id
				mu.Unlock()
				time.Sleep(time.Microsecond)
			}
		}(i)
	}

	wg.Wait()

	mu.Lock()
	fmt.Println("Всего записей в map:", len(m))
	mu.Unlock()
}
