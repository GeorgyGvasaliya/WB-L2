package pkg

import "fmt"

type Doctor struct {
	next department
}

func (d *Doctor) execute(p *Patient) {
	if p.DoctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.DoctorCheckUpDone = true
	d.next.execute(p)
}

func (d *Doctor) SetNext(next department) {
	d.next = next
}
