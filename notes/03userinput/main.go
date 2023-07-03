package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Howdy"
	fmt.Println(welcome)

	// to take input from user we need to import "bufio" and "os" package
	// bufio is used to read the input from the user
	// os is used to read the input from the user

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the ratings for my service:")

	// In GO we don't have any try & catch function/type instead we have "comma ok" and "comma error" syntax
	// comma ok syntax is used to check if the input is correct or not
	// comma error syntax is used to check if there is any error in the input

	input, _ := reader.ReadString('\n') // ReadString reads until space or newline character

	// _ is used to ignore the value of the variable here the value of error
	// input is used to store the value of the variable

	fmt.Println("Thanks for rating: ", input)
	fmt.Printf("Type of this rating is: %T", input)

}
