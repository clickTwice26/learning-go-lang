package main
import "time"
import "fmt"
func main(){

	// simple switch
	i := 1

	switch i {
		case 1:
			fmt.Println("i is 1")

		case 2:
			fmt.Println("i is 2")
		default:
			fmt.Println("default 0")
	}

	// multiple condition switch
	switch time.Now().Weekday(){
		case time.Saturday, time.Sunday:
			{fmt.Println("It's weekend")}
		
		default:
			fmt.Println("It's workday")
	}

	// type switch
	whoami := func(i interface{}){
		switch t := i.(type){
			case int:
				fmt.Println("int")
			case string:
				fmt.Println("string")
			case bool:
				fmt.Println("bool")
			default:
				fmt.Println("unknown", t)
		}
	}
	whoami("golang")	
	whoami(1)	
	whoami(true)	
	// no single quote in go
}