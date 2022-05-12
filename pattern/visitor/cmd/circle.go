package cmd

type Circle struct {
	R int
}

func (r Circle) Accept(visitor Visitor) {
	visitor.VisitForCircle(r)
}
