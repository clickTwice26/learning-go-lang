package main
import "fmt"
func add(a int, b int) int{
	return a + b
}
// multiple value return 

func getLanguages() (string, string, string) {
	return "go", "javascript", "c"
}

func main(){
	lang1, lang2, lang3 := getLanguages()
	fmt.Println(add(1, 2))
	fmt.Println(getLanguages())
	fmt.Println(lang1, lang2, lang3)

	// remaining :https://youtu.be/tG5qvZIUJx4?list=PLXQpH_kZIxTWUe-Ee-DZEX5gfeoo4tHV6&t=481
}