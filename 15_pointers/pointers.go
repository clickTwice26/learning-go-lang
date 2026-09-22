package main

import (
	"fmt"
)
// By value
func changeNum(num int){
	num = 5
	fmt.Println("Inside changeNum : ", num)
} 
func changeNumByReference(num *int){
	*num = 5
	fmt.Println("Inside changeNumByReference", *num)
	
}
func main(){
	num := 1

	// changeNum(num)
	changeNumByReference(&num)
	fmt.Println("Memory address of `num`", &num)
	
	fmt.Println("After changeNum : ", num)

	
}