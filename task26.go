package main

import (
    "fmt"
    "strings"
)

// isUnique проверяет, содержат ли все символы в строке уникальные значения
func isUnique(str string) bool {
    str = strings.ToLower(str) // Приводим строку к нижнему регистру
    charMap := make(map[rune]bool) // Создаем карту для отслеживания встреченных символов

    // Проходим по всем символам строки
    for _, char := range str {
        // Если символ уже встречался, возвращаем false
        if charMap[char] {
            return false
        }
        // Иначе отмечаем символ как встреченный
        charMap[char] = true
    }
    return true // Если все символы уникальны, возвращаем true
}

func main() {
    stringsToCheck := []string{"qwer", "asdf", "ghjk"}

    // Проверяем каждую строку на уникальность и выводим результат
    for _, str := range stringsToCheck {
        fmt.Printf("Строка \"%s\": %t\n", str, isUnique(str))
    }
}
