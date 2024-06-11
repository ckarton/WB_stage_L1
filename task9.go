package main

import "fmt"

func main() {
    input := make(chan int)
    output := make(chan int)

    // Горутина для отправки чисел в канал input
    go func() {
        defer close(input) // Закрываем канал input после завершения работы горутины
        numbers := []int{1, 2, 3, 4, 5}
        for _, num := range numbers {
            input <- num // Отправляем число в канал input
        }
    }()

    // Горутина для чтения чисел из канала input, удваивания и отправки в канал output
    go func() {
        defer close(output) // Закрываем канал output после завершения работы горутины
        for num := range input {
            output <- num * 2 // Удваиваем число и отправляем в канал output
        }
    }()

    // Чтение удвоенных чисел из канала output и вывод на экран
    for result := range output {
        fmt.Println(result)
    }
}
