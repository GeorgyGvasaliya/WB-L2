package main

import "L2/pattern/chain_of_resp/pkg"

func main() {

	cashier := &pkg.Cashier{}

	//Set next for medical department
	medical := &pkg.Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &pkg.Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &pkg.Reception{}
	reception.SetNext(doctor)

	patient := &pkg.Patient{Name: "abc"}
	//Patient visiting
	reception.Execute(patient)
}
