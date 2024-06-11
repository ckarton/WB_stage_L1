package main
import (
    "fmt"
    "math"
)

type Point struct {
    x, y float64
}

// Конструктор для структуры Point.
func NewPoint(x, y float64) Point {
    return Point{x, y}
}

// Метод для вычисления расстояния между текущей точкой и другой точкой.
func (p Point) DistanceTo(other Point) float64 {
    dx := p.x - other.x
    dy := p.y - other.y
    return math.Sqrt(dx*dx + dy*dy)
}

func main() {
    point1 := NewPoint(1, 2)
    point2 := NewPoint(4, 6)

    distance := point1.DistanceTo(point2)

    fmt.Printf("Расстояние между точками (%.2f, %.2f) и (%.2f, %.2f) составляет %.2f\n",
        point1.x, point1.y, point2.x, point2.y, distance)
}