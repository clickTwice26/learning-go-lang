package main

// no while or any other looping feature only with the for loop
// for -> only construct in go for looping
import "fmt"
func main(){

	// while loop fashioned loop in go
	i := 1
	for i <= 10{
		fmt.Println(i)
		i++
	}

	// infinite loop
	// for {
	// 	fmt.Println("Running forever")
	// }

	// classic for loop (c style)

	for i := 0; i < 3; i++{
		fmt.Println("This is the standard classic loop")
	}

	// `continue` and `break`
	for i := 0; i < 3; i++{
		if i == 2{
			continue
		}
		fmt.Println(i, "This is the standard classic loop")
	}


	// go new feature : range , starts from 0 by default to n-1
	for i := range(3){
		fmt.Println(i)
	}
}
