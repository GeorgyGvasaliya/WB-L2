package cmd

import "math"

// Shape раз работаем с фигурами то будет интерфейс фигуры
type Shape interface {
	Accept(visitor Visitor)
}

// Visitor интерфейс для добавления недостающей логики в основной интерфейс взаимодействия с фигурами Shape
type Visitor interface {
	VisitForSquare(square Square)
	VisitForCircle(circle Circle)
}

type AreaCalculator struct {
	Area float64
}

func (c *AreaCalculator) VisitForSquare(square Square) {
	c.Area = float64(square.A * square.B)
}

func (c *AreaCalculator) VisitForCircle(circle Circle) {
	c.Area = math.Pi * math.Pow(float64(circle.R), 2)
}
