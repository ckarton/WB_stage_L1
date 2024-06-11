package main
import (
    "fmt"
)

func main() {
    a := 2<<20 + 1 // Пример числа, большего, чем 2^20
    b := 2<<20 + 2 // Еще один пример числа, большего, чем 2^20

    // Умножение
    multiplication := a * b
    fmt.Printf("%d * %d = %d\n", a, b, multiplication)

    // Деление
    division := a / b
    fmt.Printf("%d / %d = %d\n", a, b, division)

    // Сложение
    addition := a + b
    fmt.Printf("%d + %d = %d\n", a, b, addition)

    // Вычитание
    subtraction := a - b
    fmt.Printf("%d - %d = %d\n", a, b, subtraction)
}