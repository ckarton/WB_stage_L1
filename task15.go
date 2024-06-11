package main

import "fmt"

// createHugeString создает строку большого размера, заполненную символами 'a'
func createHugeString(size int) string {
    hugeString := make([]byte, size)
    for i := range hugeString {
        hugeString[i] = 'a'
    }
    return string(hugeString)
}

// someFunc вызывает createHugeString и возвращает либо всю строку, либо первые 100 символов
func someFunc() string {
    v := createHugeString(1 << 10) // 1 << 10 - это 1024
    if len(v) < 100 {
        return v
    }
    return string(v[:100])
}

func main() {
    // Вызываем someFunc, чтобы получить строку, и выводим ее
    justString := someFunc()
    fmt.Println(justString)
}
