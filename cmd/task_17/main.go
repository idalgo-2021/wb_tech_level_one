// Бинарный поиск
//
// Реализовать алгоритм бинарного поиска встроенными методами языка.
// Функция должна принимать отсортированный слайс и искомый элемент, возвращать индекс элемента или -1, если элемент не найден.
//
// Подсказка: можно реализовать рекурсивно или итеративно, используя цикл for.
/////////////////////////////////////////////////////////////

package main

import "fmt"

// На входе ожидается отсортированный массив)
func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2 // текущий средний индекс

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	sortedSlice := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	target1 := 23
	target2 := 50
	target3 := 2
	target4 := 91

	fmt.Printf("Поиск %d: Индекс = %d\n", target1, BinarySearch(sortedSlice, target1)) // Ожидается: 5
	fmt.Printf("Поиск %d: Индекс = %d\n", target2, BinarySearch(sortedSlice, target2)) // Ожидается: -1
	fmt.Printf("Поиск %d: Индекс = %d\n", target3, BinarySearch(sortedSlice, target3)) // Ожидается: 0
	fmt.Printf("Поиск %d: Индекс = %d\n", target4, BinarySearch(sortedSlice, target4)) // Ожидается: 9
}
