// Пересечение множеств
//
//
// Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов) — т.е. вывести элементы, присутствующие и в первом, и во втором.
//
// Пример:
// A = {1,2,3}
// B = {2,3,4}
// Пересечение = {2,3}
/////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

func main() {

	// Учтем вариант, что в множествах могут быть дублирующиеся элементы(хотя это не совсем корректно)

	var set1 = []int{11, 7, 4, 8, 5, 7, 1}
	var set2 = []int{10, 3, 7, 5, 6, 3, 12, 33}

	set_res := intersection(set1, set2)
	fmt.Println("Пересечение множеств: ", set_res)
}

func intersection(a, b []int) []int {
	result := []int{}
	m := make(map[int]bool)

	for _, v := range a {
		m[v] = true
	}

	seen := make(map[int]bool)
	for _, v := range b {
		if m[v] && !seen[v] {
			result = append(result, v)
			seen[v] = true
		}
	}

	return result
}
