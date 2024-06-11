package main

import "fmt"

// Функция intersection находит пересечение двух множеств
func intersection(set1, set2 map[int]bool) map[int]bool {
    result := make(map[int]bool) // Создаем карту для результата пересечения

    // Проходим по элементам первого множества
    for element := range set1 {
        // Если текущий элемент также содержится во втором множестве,
        // добавляем его в результат
        if set2[element] {
            result[element] = true
        }
    }

    return result
}

func main() {
    // Определяем два множества
    set1 := map[int]bool{1: true, 2: true, 3: true, 4: true}
    set2 := map[int]bool{3: true, 4: true, 5: true, 6: true}

    // Находим пересечение множеств set1 и set2
    intersectionSet := intersection(set1, set2)

    // Выводим результат на экран
    fmt.Println("Пересечение множеств:", intersectionSet)
}
