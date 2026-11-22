package main

import (
	"fmt"
	"time"
)
type Customer struct{
	name string
	phone string
}
type Order struct{
	id string
	amount float32
	status string
	createAt time.Time
	Customer
}
func newOrder(id string, amount float32, status string, customerName string, customerPhone string) *Order{
	myOrder := Order{
		id : id,
		amount : amount,
		status : status,
		createAt: time.Now(),
	 	
	}
	return &myOrder
}

// receiver type
func (O *Order) changeStatus(status string){
	O.status = status
	
}

func main(){
	// var order Order = 
	// if you dont set any field the default value is 0 value
	order := Order{
		id : "test123",
		amount : 123.2,
		status : "pending",
	}
	order.changeStatus("changed")
	order.amount = 132.23
	order.createAt = time.Now()
	order.Customer.name = "Shagato"
	order.Customer.phone = "1234"
	fmt.Println(order)
	fmt.Println("Customer info : ", order.Customer)
	myOrder2 := Order {
		id : "2",
		amount : 15,
		status : "delivered",
		createAt : time.Now(),
	}
	fmt.Println(myOrder2)
}