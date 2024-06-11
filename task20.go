package main

import (
    "fmt"
    "strings"
)

// Функция reverseWords переворачивает слова во входной строке
func reverseWords(input string) string {
    words := strings.Fields(input) // Разбиваем строку на слова

    n := len(words)
    for i := 0; i < n/2; i++ {
        words[i], words[n-1-i] = words[n-1-i], words[i] // Инвертируем массив слов
    }
    return strings.Join(words, " ") // Объединяем слова в строку с пробелами и возвращаем результат
}

func main() {
    input := "snow dog sun"
    fmt.Println("Исходная строка:", input)
    
    // Переворачиваем слова в строке и выводим результат
    reversed := reverseWords(input)
    fmt.Println("Строка с перевёрнутыми словами:", reversed)
}
