package main


import (
	"fmt"
)
type paymenter interface {
	pay(amount float32) 
}
type payment struct{


	gateway paymenter
}

func (p payment) makePayment(amount float32){
	// razorpayPaymentGw := razorpay{}
	// stripePaymentGw := stripe{}
	
	// razorpayPaymentGw.pay(amount)
	// stripePaymentGw.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct {
	
}
func (r razorpay) pay(amount float32){
	// logic to make payment
	fmt.Println("making payment using razor pay", amount)
}

type stripe struct{}
func (s stripe) pay(amount float32){
	fmt.Println("Making payment using stripe", amount)
}

type fakePayment struct{}

func (f fakePayment) pay (amount float32){
	fmt.Println("making payment using fake gateway for testing purpose", amount)
}
func main(){

	fakeGw := fakePayment{}
	newPayment := payment{
		gateway: fakeGw,
	}
	newPayment.makePayment(100)

}