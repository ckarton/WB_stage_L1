package main

import (
    "fmt"
    "sync"
)

func main() {
    // Создаем карту для хранения данных
    data := make(map[string]int)
    // Создаем мьютекс для синхронизации доступа к карте
    var mutex sync.Mutex

    // Запускаем 5 горутин
    for i := 0; i < 5; i++ {
        go func(index int) {
            // Блокируем мьютекс перед записью в карту
            mutex.Lock()
            // Записываем данные в карту
            data[fmt.Sprintf("key%d", index)] = index
            // Разблокируем мьютекс после записи
            mutex.Unlock()
        }(i)
    }

    // Сообщаем, что ждем завершения всех горутин
    fmt.Println("Waiting for goroutines to finish...")

    // Блокируем мьютекс перед чтением карты, чтобы дождаться завершения всех горутин
    mutex.Lock()
    // Используем defer для разблокировки мьютекса после завершения работы с картой
    defer mutex.Unlock()

    // Выводим данные из карты
    fmt.Println("Data in map:")
    for key, value := range data {
        fmt.Printf("%s: %d\n", key, value)
    }
}
