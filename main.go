package main
import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your full name: ")
	if scanner.Scan(){
		full_name := scanner.Text()
		fmt.Println("Your full name is ", full_name)
	}
}