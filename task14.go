package main

import (
    "fmt"
)

// Функция identifyType определяет тип значения v и выводит соответствующее сообщение
func identifyType(v interface{}) {
    switch v.(type) {
    case int:
        fmt.Println("Это int")
    case string:
        fmt.Println("Это string")
    case bool:
        fmt.Println("Это bool")
    case chan int:
        fmt.Println("Это channel int")
    case chan string:
        fmt.Println("Это channel string")
    case chan bool:
        fmt.Println("Это channel bool")
    default:
        fmt.Println("Неизвестный тип")
    }
}

func main() {
    // Определение переменных различных типов
    var a int = 10
    var b string = "hello"
    var c bool = true
    var d chan int = make(chan int)
    var e chan string = make(chan string)
    var f chan bool = make(chan bool)

    // Определение типов переменных и вывод сообщений
    identifyType(a)
    identifyType(b)
    identifyType(c)
    identifyType(d)
    identifyType(e)
    identifyType(f)
}
