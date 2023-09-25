// Разработать программу нахождения расстояния между двумя точками, 
// которые представлены в виде структуры Point с инкапсулированными 
// параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

// Создаем структуру Point для точек
type Point struct {
	x, y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// вычисление расстояния между двумя точками
func DistanceTo(p1 Point, p2 Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаем две точки
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	// Вычисляем расстояние между ними
	distance := DistanceTo(point1, point2)

	fmt.Printf("Расстояние между точкой 1 и точкой 2: %.2f\n", distance)
}
