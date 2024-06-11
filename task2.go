package main

import (
    "fmt"
    "sync"
)

func main() {

    // Инициализируем срез чисел
    numbers := []int{2, 4, 6, 8, 10}

    // Создаем канал для передачи результатов
    results := make(chan int)

    // Создаем экземпляр синхронизатора WaitGroup
    var wg sync.WaitGroup

    // Запускаем горутины для вычисления квадратов чисел
    for _, num := range numbers {
        wg.Add(1) // Увеличиваем счетчик WaitGroup
        go func(n int) {
            defer wg.Done() // Уменьшаем счетчик WaitGroup после завершения горутины
            results <- n * n // Отправляем результат в канал
        }(num) // Передаем значение num в горутину
    }

    // Запускаем горутину для закрытия канала после завершения всех вычислений
    go func() {
        wg.Wait() // Ожидаем завершения всех горутин
        close(results) // Закрываем канал
    }()

    // Читаем и выводим результаты из канала
    for square := range results {
        fmt.Println(square)
    }
}
