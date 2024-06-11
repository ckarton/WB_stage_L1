package main

import "fmt"

// removeIndex удаляет элемент из среза по указанному индексу
func removeIndex(slice []int, index int) []int {
    return append(slice[:index], slice[index+1:]...)
}

func main() {
    slice := []int{1, 2, 3, 4, 5}
    index := 0
    fmt.Println("Исходный срез:", slice)
    
    // Удаление элемента с указанным индексом
    slice = removeIndex(slice, index)
    fmt.Println("Срез после удаления элемента с индексом", index, ":", slice)
}
