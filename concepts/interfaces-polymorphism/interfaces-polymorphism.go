package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (c Circle) Area() float64 {

	return 3.14 * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

type PaymentGateway interface {
	Pay(amount float64) string
}

type PayPal struct{}
type Stripe struct{}

func (PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f via PayPal", amount)
}

func (Stripe) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f via Stripe", amount)
}

func processPayment(pg PaymentGateway, amt float64) {
	fmt.Println(pg.Pay(amt))
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 2, Height: 2}

	// Go uses dynamic dispatch here — it calls the correct method based on the actual type at runtime.
	printArea(c)
	printArea(r)

	processPayment(PayPal{}, 200)
	processPayment(Stripe{}, 5)
}
