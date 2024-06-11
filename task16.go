package main

import (
    "fmt"
)

// Функция quicksort реализует алгоритм быстрой сортировки для среза целых чисел
func quicksort(arr []int) []int {
    // Если длина среза меньше или равна 1, он уже отсортирован
    if len(arr) <= 1 {
        return arr
    }

    pivot := arr[0] // Опорный элемент
    left := []int{} // Срез для элементов меньше опорного
    right := []int{} // Срез для элементов больше или равных опорному

    // Разбиваем срез на две части относительно опорного элемента
    for _, v := range arr[1:] {
        if v < pivot {
            left = append(left, v) // Элемент меньше опорного, добавляем в left
        } else {
            right = append(right, v) // Элемент больше или равен опорному, добавляем в right
        }
    }

    // Рекурсивно сортируем левую и правую части
    left = quicksort(left)
    right = quicksort(right)

    // Соединяем отсортированные части вместе с опорным элементом
    return append(append(left, pivot), right...)
}

func main() {
    // Создаем исходный массив
    arr := []int{10, 7, 8, 9, 1, 5}
    fmt.Println("Исходный массив:", arr)

    // Сортируем массив с помощью quicksort и выводим отсортированный массив
    sortedArr := quicksort(arr)
    fmt.Println("Отсортированный массив:", sortedArr)
}
