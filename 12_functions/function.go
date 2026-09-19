package main
import "fmt"
func add(a int, b int) int{
	return a + b
}
// multiple value return 

func getLanguages() (string, string, string) {
	return "go", "javascript", "c"
}

func processIt(fn func(a int) int){
	fn(1)
}
func returnIt() func(a int) int{
	return func(a int) int{
		return 1	
	}
}
func main(){
	lang1, lang2, lang3 := getLanguages()
	fmt.Println(add(1, 2))
	fmt.Println(getLanguages())
	fmt.Println(lang1, lang2, lang3)

	fn := func(a int) int {
		return 2
	}
	processIt(fn)

	
}