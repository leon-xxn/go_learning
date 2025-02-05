package chain_of_responsibility

import "testing"

func TestChain(t *testing.T) {
	cashier := &Cashier{}

	//set next for medical department
	medical := &Medical{}
	medical.setNext(cashier)

	//set next for doctor department
	doctor := &Doctor{}
	doctor.setNext(medical)

	//set next for reception department
	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "abc"}
	reception.execute(patient)

}
