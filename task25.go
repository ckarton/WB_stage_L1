package main

import (
    "fmt"
    "time"
)

// sleep приостанавливает выполнение программы на указанное количество секунд
func sleep(seconds int) {
    time.Sleep(time.Duration(seconds) * time.Second)
}

func main() {
    fmt.Println("Начало программы")
    sleep(5) // Приостановка выполнения на 5 секунд
    fmt.Println("Программа завершена")
}
