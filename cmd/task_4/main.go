// Завершение по Ctrl+C
//
// Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).
//
// Выберите и обоснуйте способ завершения работы всех горутин-воркеров при получении сигнала прерывания.
//
// Подсказка: можно использовать контекст (context.Context) или канал для оповещения о завершении.
/////////////////////////////////////////////////////////////////////////////////

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

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
	for {
		select {
		case <-ctx.Done(): // Сигнал завершения
			fmt.Printf("Worker %d stopping...\n", id)
			return
		default:
			// Эмуляция работы
			fmt.Printf("Worker %d is working...\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {

	// Контекст, через который будет передваться отмена
	ctx, cancel := context.WithCancel(context.Background())

	// Канал для перехвата сигналов
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM) // "подписка" на сигналы

	var wg sync.WaitGroup

	// Воркеры
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	// Ждём сигнал в главной горутине
	<-sigChan
	fmt.Println("\nGot interrupt signal, shutting down...")

	// Стгнал получен: оповещение воркеров через cancel() - хакрывает канал ctx.Done() и
	// каждый воркер обработает это событие(есть соответствующий обработчик)
	cancel()

	// Ожидаем завершения всех горутин
	wg.Wait()
	fmt.Println("All workers stopped, program finished.")
}
