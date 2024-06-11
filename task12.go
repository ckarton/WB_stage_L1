package main

import "fmt"

func main() {
    strings := []string{"cat", "cat", "dog", "cat", "tree"} // Заданный срез строк
    set := make(map[string]bool) // Создаем пустое множество строк

    // Проходим по всем строкам в срезе
    for _, str := range strings {
        set[str] = true // Добавляем каждую строку в множество
    }

    fmt.Println("Множество строк:", set) // Выводим созданное множество строк
}
