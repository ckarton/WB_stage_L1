package main

import "fmt"

// Определяем структуру Human, которая имеет одно поле - Name
type Human struct {
    Name string
}

// Определяем метод Speak для структуры Human, который выводит на экран имя
func (h Human) Speak() {
    fmt.Println("Hi, my name is", h.Name)
}

// Определяем структуру Action, которая встраивает структуру Human
type Action struct {
    Human
}

func main() {
    // Создаем экземпляр структуры Action и инициализируем встроенную структуру Human с именем "Bob"
    action := Action{
        Human: Human{
            Name: "Bob",
        },
    }

    // Вызываем метод Speak, который был унаследован от структуры Human
    action.Speak()
}
