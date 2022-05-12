package pkg

import "fmt"

type Medical struct {
	next department
}

func (m *Medical) execute(p *Patient) {
	if p.MedicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.MedicineDone = true
	m.next.execute(p)
}

func (m *Medical) SetNext(next department) {
	m.next = next
}
