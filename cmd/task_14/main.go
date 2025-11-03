// Определение типа переменной в runtime
//
// Разработать программу, которая в runtime способна определить тип переменной, переданной в неё (на вход подаётся interface{}).
// Типы, которые нужно распознавать: int, string, bool, chan (канал).
//
// Подсказка: оператор типа switch v.(type) поможет в решении.
/////////////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

func main() {
	determineType(42)                // int
	determineType("hello")           // string
	determineType(true)              // bool
	ch := make(chan interface{})     // Создаём канал
	determineType(ch)                // chan
	determineType(make(chan string)) // chan

	determineType(3.14) // float64 — Неизвестный тип

}

func determineType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Тип: int")
	case string:
		fmt.Println("Тип: string")
	case bool:
		fmt.Println("Тип: bool")
	case chan int, chan string, chan bool, chan interface{}:
		fmt.Println("Тип: chan")
	default:
		fmt.Println("Неизвестный тип")
	}
}
