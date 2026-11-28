package main

import (
	"fmt"
)
type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed 
	Prepared 
	Delivered
)
func changeTheOrder(){
	fmt.Println("Change the data of the website")
}
func changeOrderStatus(status OrderStatus){
	fmt.Println("Chaning order status to", status)
}
func main(){

	changeOrderStatus(Prepared)
	for true {
		fmt.Println("Testing")
		
	}
}

