package main

import (
    "fmt"
    "time"
)

func main() {
    N := 5 // Количество секунд, через которое программа завершится
    ch := make(chan int) // Канал для передачи значений
    done := make(chan struct{}) // Канал для сигнала завершения

    // Горутина для приема значений из канала ch
    go func() {
        for {
            select {
            case val := <-ch: // Если получили значение из канала ch
                fmt.Println("Received:", val)
            case <-done: // Если получили сигнал завершения
                return
            }
        }
    }()

    // Горутина для отправки значений в канал ch
    go func() {
        for i := 0; ; i++ {
            select {
            case ch <- i: // Отправляем значение в канал ch
                fmt.Println("Sent:", i)
            case <-done: // Если получили сигнал завершения
                return
            }
            time.Sleep(time.Second) // Ждем 1 секунду перед отправкой следующего значения
        }
    }()

    time.Sleep(time.Duration(N) * time.Second) // Основная горутина ждет N секунд

    close(done) // Отправляем сигнал завершения всем горутинам
}
