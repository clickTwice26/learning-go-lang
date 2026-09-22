package main

import "fmt"

// slice -> dynamic array (most used constructor in go)
// many useful method
func main(){
	var nums []int
	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	var number = make([]int, 5, 10)
	// capacity -> maximum number of elements can fit
	fmt.Println(number)
	fmt.Println(cap(number))
	number = append(number, 1)
	number = append(number, 1)
	number = append(number, 1)
	number = append(number, 1)
	number = append(number, 1)
	number = append(number, 1)
	number = append(number, 1)
	fmt.Println(number)
	fmt.Println(cap(number))


	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] 
	fmt.Println("Output",s)

}