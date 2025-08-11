package main

type User struct {
	name          string
	email         string
	paymentMethod string
}

type PaymentProcess interface {
	processPayment()
}

type BankPayment struct{}

func (b BankPayment) PaymentProcess(u *User) {}
