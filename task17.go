package main
import (
    "fmt"
    "sort"
)

func main() {
    // Инициализируем отсортированный массив
    arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

    // Искомое значение
    target := 13

    // Используем встроенную функцию sort.Search для бинарного поиска
    index := sort.Search(len(arr), func(i int) bool {
        return arr[i] >= target
    })

    // Проверяем, найдено ли значение
    if index < len(arr) && arr[index] == target {
        fmt.Printf("Значение %d найдено на позиции %d\n", target, index)
    } else {
        fmt.Printf("Значение %d не найдено\n", target)
    }
}