// Собственное множество строк
//
//
// Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.
//
// Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.
//
///////////////////////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

func main() {
	set_words := []string{"cat", "cat", "dog", "cat", "tree"}
	set_res := []string{}

	m := make(map[string]bool)
	for _, v := range set_words {

		if _, ok := m[v]; ok {
			continue
		} else {
			m[v] = true
			set_res = append(set_res, v)
		}
	}

	fmt.Println(set_res)

}
