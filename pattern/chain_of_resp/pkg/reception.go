package pkg

import "fmt"

type Reception struct {
	next department
}

func (r *Reception) Execute(p *Patient) {
	if p.RegistrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.RegistrationDone = true
	r.next.execute(p)
}

func (r *Reception) SetNext(next department) {
	r.next = next
}
