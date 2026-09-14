package main

import "fmt"

func main(){
	age := 5
	if age >= 18 {
		fmt.Println("You are an adult")
		
	}else if age <=5 {
		fmt.Println("You are a super minor")
	}else{
		fmt.Println("You are a minor")
	}
	// var role string = "admin"
	var role string = "user"
	var hasPermission bool = true
	// logical operator like c ||( or ) && ( and )
	if role == "admin" || hasPermission{
		hasPermission = true
	}else{
		hasPermission = false
	}
	fmt.Println(hasPermission)

	if age:=15; age >= 18{
		fmt.Println("You are an adult")
	}else if age <=12{
		fmt.Println("You are an teenager")
		
	}
	// go does not have ternery operator , you have to use normal version
}