package chain_of_responsibility

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment already done")
	}
	fmt.Println("Cashier getting payment from patient")
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}
