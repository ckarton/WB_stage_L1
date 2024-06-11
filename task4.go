package main

import (
    "fmt"
    "math/rand"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "time"
)

// worker - это функция, которая будет выполняться в горутине для обработки заданий
func worker(id int, jobs <-chan int, wg sync.WaitGroup) {
    defer wg.Done() // Уменьшаем счетчик WaitGroup после завершения работы воркера
    for job := range jobs { // Читаем задания из канала jobs
        fmt.Printf("Worker %d received job: %d\n", id, job) // Обрабатываем задание (в данном случае просто выводим на экран)
    }
}

func main() {
    numWorkers := 3 // Количество воркеров по умолчанию

    // Если передан аргумент командной строки, переопределяем количество воркеров
    if len(os.Args) > 1 {
        fmt.Sscanf(os.Args[1], "%d", &numWorkers)
    }

    jobs := make(chan int) // Канал для передачи заданий воркерам
    var wg sync.WaitGroup // WaitGroup для ожидания завершения всех воркеров

    // Запускаем воркеров
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1) // Увеличиваем счетчик WaitGroup
        go worker(i, jobs, wg) // Запускаем воркера в отдельной горутине
    }

    stop := make(chan os.Signal, 1) // Канал для получения сигналов прерывания
    signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM) // Настраиваем канал на прием сигналов SIGINT и SIGTERM

    // Горутина для отправки случайных заданий воркерам
    go func() {
        for {
            select {
            case <-stop: // Если получен сигнал прерывания
                close(jobs) // Закрываем канал jobs
                return
            default:
                jobs <- rand.Intn(100) // Отправляем случайное задание в канал
                time.Sleep(500 * time.Millisecond) // Задержка между заданиями
            }
        }
    }()

    <-stop // Ожидаем получения сигнала прерывания
    close(jobs)  // Убедимся, что канал jobs закрыт
    wg.Wait() // Ожидаем завершения всех воркеров
}
