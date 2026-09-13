package main

import "fmt"

func main(){

	

	fmt.Println("Learning constants")
	// constant grouping 
	const (
		port = 5000
		host = "localhost"
	)
	const name string = "Go-Lang"
	// name = "javascript"
	fmt.Println(port, host, name)

	
}

// constant can be declared outside of the functiona also
// := can't use this shorthand assignment outside of the function
// constant value can't be changed and must be initialized with a value when declared