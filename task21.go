package main

import (
    "fmt"
)

// Printer - интерфейс, определяющий метод Show()
type Printer interface {
    Show()
}

// LegacyPrinter - структура, представляющая старый принтер
type LegacyPrinter struct{}

// Display - метод для отображения
func (lp *LegacyPrinter) Display() {
    fmt.Println("Legacy Printer is displaying")
}

// PrinterImpl - структура, реализующая интерфейс Printer и использующая LegacyPrinter
type PrinterImpl struct {
    legacyPrinter *LegacyPrinter
}

// Show - метод для реализации интерфейса Printer
func (pi *PrinterImpl) Show() {
    pi.legacyPrinter.Display()
}

func main() {
    // Создание экземпляра LegacyPrinter
    legacyPrinter := &LegacyPrinter{}

    // Создание адаптера PrinterImpl
    adapter := &PrinterImpl{
        legacyPrinter: legacyPrinter,
    }

    // Использование адаптера через интерфейс Printer
    adapter.Show()
}
