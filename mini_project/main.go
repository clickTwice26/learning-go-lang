package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main(){

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your age: ")
	scanner.Scan()
	text := strings.TrimSpace(scanner.Text())

	age, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("Please enter a valid integer!")
		return
	}
	fmt.Printf("You are %d years old.\n", age)
	
	

	
}