package main

import (
	"L2/pattern/visitor/cmd"
	"fmt"
)

func main() {
	s := cmd.Square{A: 2, B: 5}
	c := cmd.Circle{R: 14}

	calc := &cmd.AreaCalculator{}

	s.Accept(calc)
	fmt.Printf("Square area: %.2f\n", calc.Area)
	c.Accept(calc)
	fmt.Printf("Circle area: %.2f\n", calc.Area)
}
