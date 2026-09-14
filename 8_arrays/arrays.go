package main

import "fmt"

func main(){

	//numbered sequence of specific length
	var nums[4]int // if int type then all the elements will be zero

	nums[3] = 1
	fmt.Println(len(nums))
	fmt.Println(nums)

	var vals[4]bool //by default fills with false
	fmt.Println(vals)


	var strs[4]string
	fmt.Println(strs)

	
	// remaining class : https://youtu.be/ZQmZmoziWj0?list=PLXQpH_kZIxTWUe-Ee-DZEX5gfeoo4tHV6&t=486
	numbers := [3]int{1, 2, 3}
	fmt.Println(numbers)

	//2-D Class
	dimen := [2][2]int{{3,4}, {5, 6}}
	fmt.Println(dimen)


	// When to use array
	// 1. When we know the size and it is fixed , memory optmized
	// 2. Constant time access O(1)
}