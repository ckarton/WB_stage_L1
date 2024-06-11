package main
import "fmt"

// Устанавливает i-й бит в 1
func setBit(num *int64, i uint) {
    *num |= 1 << i
}

// Устанавливает i-й бит в 0
func clearBit(num *int64, i uint) {
    *num &^= 1 << i
}

func main() {
    var num int64 = 42
    fmt.Printf("Исходное число: %d\n", num)

    // Установка i-го бита в 1
    i := uint(2)
    setBit(&num, i)
    fmt.Printf("Число после установки %d-го бита в 1: %d\n", i, num)

    // Установка i-го бита в 0
    clearBit(&num, i)
    fmt.Printf("Число после установки %d-го бита в 0: %d\n", i, num)
}