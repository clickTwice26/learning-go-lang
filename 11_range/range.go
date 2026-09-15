package main
// iterating over data structure
import "fmt"
func main(){
	// num := make([]int,5)
	num := []int{1,2,3,4}

	for i:=0; i < len(num); i++{
		fmt.Println(num[i])
	}
	print("*\n")
	sum := 0
	for _, n := range(num){ // here `_` is the index
		sum+=n
	}
	fmt.Println("the sum is" , sum)


	user_data := map[string]string{
		"name" : "Shagato Chowdhury",
		"age" : "30",
	}
	for index, value := range user_data{
		fmt.Println("Index", index, "Value", value)
	}
	// Unicode point rune
	// here i is not index, starting byte of the rune
	// ASCII - 255
	for i, v := range "go-lang"{
		fmt.Println(i, v)
		fmt.Println(i, string(v))
	}
	
	
}