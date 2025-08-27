// Встраивание структур
//
// Дана структура Human (с произвольным набором полей и методов).
//
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
//
// Подсказка: используйте композицию (embedded struct), чтобы Action имел все методы Human.
/////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

type Human struct {
	Name string
	Age  int
}

func (h *Human) Speak(msg string) {
	fmt.Printf("%s: %s\n", h.Name, msg)
}
func (h *Human) Birthday() {
	h.Age++
}
func (h Human) Info() string {
	return fmt.Sprintf("Human{name=%q, age=%d}", h.Name, h.Age)
}

///////////

type Action struct {
	Human // embedded: методы Human наследуются в Action
	Role  string
}

func (a *Action) Do(task string) {
	fmt.Printf("%s (%s) что делает: %s\n", a.Name, a.Role, task)
}

// Можем переопределить
// func (a *Action) Speak(msg string) {
// 	fmt.Printf("[ACTION] %s (%s): %s\n", a.Name, a.Role, msg)
// }

func main() {
	h := Human{Name: "Миша", Age: 29}
	a := Action{
		Human: h,
		Role:  "Водитель",
	}

	fmt.Println(h.Info()) // Human{name="Миша", age=29}
	fmt.Println(a.Info()) // Human{name="Миша", age=29}
	fmt.Println("")

	h.Birthday()
	fmt.Println(h.Info()) // Human{name="Миша", age=30}
	a.Birthday()
	fmt.Println(a.Info()) // Human{name="Миша", age=30}
	fmt.Println("")

	a.Birthday()
	fmt.Println(h.Info()) // Human{name="Миша", age=30}
	fmt.Println(a.Info()) // Human{name="Миша", age=31}
	fmt.Println("")

	a.Do("рулит")
	fmt.Println("") // Миша (Водитель) что делает: рулит

	a.Speak("Привет from Action!")      // Миша: Привет from Action!
	a.Human.Speak("Привет from Human!") // Миша: Привет from Human!
	fmt.Println("")

	a.Human.Name = "Michael"
	fmt.Println(h.Info()) // Human{name="Миша", age=30}
	fmt.Println(a.Info()) // Human{name="Michael", age=31}
	fmt.Println("")

	a.Name = "Михаил"
	fmt.Println(h.Info()) // Human{name="Миша", age=30}
	fmt.Println(a.Info()) // Human{name="Михаил", age=31}

}
