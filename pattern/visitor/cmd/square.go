package cmd

type Square struct {
	A int
	B int
}

// единственное изменение square
func (s Square) Accept(visitor Visitor) {
	visitor.VisitForSquare(s)
}
