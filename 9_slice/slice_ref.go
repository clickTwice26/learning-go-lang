package main

import (
	"fmt"
	"slices"
)

func main(){
	fmt.Println("Hello")

	names := [4]string{
		"john",
		"paul",
		"george",
		"Ringo"}
	a := names[0:2]
	b := names[1:3]
	fmt.Println(names)
	fmt.Println(a, b)

	b[0] = "shagato"

	fmt.Println(b)
	fmt.Println(names)


	// Copy a slice
	nums2 := make([]int, len(names))

	fmt.Println(names, nums2)
	source := []int{1,2,3}
	copy(source, nums2)
	fmt.Println(nums2)


	// slice operator
	var nums = []int{1, 2, 3}

	fmt.Println("Slice Operator: ", nums[0:1])
	
	// slice 
	var num1 = []int{1, 2}
	var num2 = []int{1, 2}

	fmt.Println(slices.Equal(num1, num2))
	
}