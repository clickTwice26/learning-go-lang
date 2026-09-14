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
}