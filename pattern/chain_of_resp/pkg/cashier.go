package pkg

import "fmt"

type Cashier struct {
	next department
}

func (c *Cashier) execute(p *Patient) {
	if p.PaymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient")
}

func (c *Cashier) SetNext(next department) {
	c.next = next
}
