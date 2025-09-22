// Остановка горутины
//
// * Реализовать все возможные способы остановки выполнения горутины.
// * Классические подходы: выход по условию, через канал уведомления, через контекст, прекращение работы runtime.Goexit() и др.
// * Продемонстрируйте каждый способ в отдельном фрагменте кода.
//
///////////////////////////////////////////////////////////////////////////////////

package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// 1. Остановка по условию
func stopByCondition() {
	fmt.Println("\n--- Остановка по условию ---")

	done := make(chan struct{})

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Работаю:", i)
			time.Sleep(300 * time.Millisecond)
		}
		fmt.Println("Горутина завершила работу по условию")
		close(done)
	}()

	<-done
}

// 2. Остановка через канал уведомления
func stopByChannel() {
	fmt.Println("\n--- Остановка через канал ---")

	stop := make(chan struct{})
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Получен сигнал остановки")
				close(done)
				return
			default:
				fmt.Println("Работаю...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	close(stop) // сигнал остановки
	<-done
}

// 3.Остановка через контекст
func stopByContext() {
	fmt.Println("\n--- Остановка через context ---")

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context отменён")
				return
			default:
				fmt.Println("Работаю...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	cancel() // отменяем контекст
	<-done
}

// 4. Остановка через runtime.Goexit
func stopByGoexit() {
	fmt.Println("\n--- Остановка через runtime.Goexit ---")

	done := make(chan struct{})

	go func() {
		defer close(done)
		for i := 0; i < 3; i++ {
			fmt.Println("Работаю:", i)
			time.Sleep(300 * time.Millisecond)
		}
		fmt.Println("Вызываем runtime.Goexit()")
		runtime.Goexit()
	}()

	<-done
}

// 5. Остановка при панике (с recover)
func stopByPanic() {
	fmt.Println("\n--- Остановка при панике (с recover) ---")

	done := make(chan struct{})

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Паника перехвачена:", r)
			}
			close(done)
		}()

		for i := 0; i < 3; i++ {
			fmt.Println("Работаю:", i)
			time.Sleep(300 * time.Millisecond)
		}
		panic("что-то пошло не так") // горутина аварийно завершается
	}()

	<-done
}

// 6. Остановка по таймеру
func stopByTimer() {
	fmt.Println("\n--- Остановка по таймеру ---")

	done := make(chan struct{})

	go func() {
		defer close(done)
		timer := time.After(1 * time.Second) // через 1 сек. остановка

		for {
			select {
			case <-timer:
				fmt.Println("Таймер сработал, завершаем работу(не дожидаемся завершения)")
				return
			default:
				fmt.Println("Работаю...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	<-done
}

func main() {
	stopByCondition()
	stopByChannel()
	stopByContext()
	stopByGoexit()
	stopByPanic()
	stopByTimer()
}
