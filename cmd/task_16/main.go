// Быстрая сортировка (quicksort)
//
// Реализовать алгоритм быстрой сортировки массива встроенными средствами языка. Можно использовать рекурсию.
//
// Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел. Для выбора опорного элемента можно взять середину или первый элемент.
////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]

	var left []int
	var right []int
	var middle []int

	for _, v := range arr {
		switch {
		case v < pivot:
			left = append(left, v)
		case v > pivot:
			right = append(right, v)
		default:
			middle = append(middle, v)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, middle...), right...)
}

func main() {
	arr := []int{5, 3, 8, 4, 2, 7, 1, 10, 6}
	fmt.Println("До сортировки:", arr)

	sorted := quickSort(arr)
	fmt.Println("После сортировки:", sorted)
}
