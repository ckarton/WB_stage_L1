package main

import (
    "fmt"
)

// Функция reverseString переворачивает входную строку
func reverseString(input string) string {
    runes := []rune(input) // Преобразуем строку в массив рун
    n := len(runes)

    // Инвертируем массив рун
    for i := 0; i < n/2; i++ {
        runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
    }

    return string(runes) // Преобразуем массив рун обратно в строку и возвращаем ее
}

func main() {
    input := "qwertyuio"
    fmt.Println("Исходная строка:", input)
    
    // Переворачиваем строку и выводим результат
    reversed := reverseString(input)
    fmt.Println("Перевернутая строка:", reversed)
}
