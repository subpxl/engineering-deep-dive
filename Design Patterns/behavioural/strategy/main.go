package main

import "fmt"

func main() {

	upi := &UPI{}
	creditCard := &CreditCard{}
	bankTransfer := &BankTransfer{}

	shop := &Shop{}
	shop.ProcessPayment(100.0, upi)
	shop.ProcessPayment(100.0, creditCard)
	shop.ProcessPayment(100.0, bankTransfer)
}

type Shop struct {
	payment Payment
}

func (s *Shop) ProcessPayment(amount float64, paymentType Payment) {
	s.payment =paymentType
	s.payment.Pay(amount)
}

type Payment interface {
	Pay(amount float64)
}

type CreditCard struct{}

func (c *CreditCard) Pay(amount float64) {
	fmt.Println("credit card processing payment of ", amount)
}

type UPI struct{}

func (c *UPI) Pay(amount float64) {
	fmt.Println("upi processing payment  of ", amount)
}

type BankTransfer struct{}

func (c *BankTransfer) Pay(amount float64) {
	fmt.Println("processing bank tranfer   of ", amount)
}
