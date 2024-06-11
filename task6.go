package main

import (
    "fmt"
    "time"
)

// worker - функция, выполняющаяся в горутине, которая ждет сигнал завершения
func worker(done chan struct{}) {
    fmt.Println("Worker started")
    defer fmt.Println("Worker stopped") // Сообщение будет выведено при завершении функции
    <-done // Ожидаем сигнал завершения из канала done
}

func main() {
    done := make(chan struct{}) // Создаем канал для сигнала завершения
    go worker(done) // Запускаем worker в отдельной горутине
    time.Sleep(1 * time.Second) // Имитируем выполнение другой работы в главной горутине
    close(done) // Отправляем сигнал завершения, закрывая канал done
    time.Sleep(1 * time.Second) // Ждем, чтобы дать время worker завершить работу
}
