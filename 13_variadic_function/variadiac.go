package main

import (
	"fmt"
	"reflect"
)
//  function with unlimited type of unlimited paramenter

func sum(nums ...int) int{
	var sum int = 0
	fmt.Println("Type : ", reflect.TypeOf(nums))
	for index, value := range (nums){
		fmt.Println(index, value)
		sum+=value
	}
	return sum
	
}

func sum_any(nums ...interface{}) int{
	var sum int = 0
	fmt.Println("Type : ", reflect.TypeOf(nums))
	for index, value := range (nums){
		fmt.Println(index, value)
		// sum+=value
	}
	return sum
	
}

func main(){
	numbers := []int{3,4,5,6}
	fmt.Println(1,2,3,4,5,"hello")
	fmt.Println("Sum is ", sum(1,2,3,4))
	fmt.Println("Sum is ", sum(numbers...)) // spreaded
	
}