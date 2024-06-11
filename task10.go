package main

import (
    "fmt"
    "math"
)

// Функция для группировки температур по заданному шагу
func groupTemperatures(temperatures []float64, step float64) map[float64][]float64 {
    groups := make(map[float64][]float64) // Создаем карту для хранения групп температур

    // Проходим по всем температурам
    for _, temp := range temperatures {
        // Вычисляем ключ для текущей температуры на основе шага
        groupKey := math.Floor(temp/step) * step

        // Добавляем текущую температуру в соответствующую группу
        groups[groupKey] = append(groups[groupKey], temp)
    }

    return groups
}

func main() {
    temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} // Заданные температуры
    groupStep := 10.0 // Шаг группировки

    // Группируем температуры
    groups := groupTemperatures(temperatures, groupStep)

    // Выводим информацию о группах температур
    for key, group := range groups {
        fmt.Printf("Диапазон: %.1f - %.1f\n", key, key+groupStep)
        fmt.Println("Температуры:", group)
        fmt.Println()
    }
}
