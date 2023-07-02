package main

import "fmt"

// we can use constant outside the method and can declare them as
// const + constant name + constant type = constant value

const LoginToken string = "ammuah" //capital 'L' represents that it can be exported anywhere in the file meaning it is a public variable

func main() {

	// Explicit type i.e, declaring a variable with assigning its type
	// var + variable name + variable type = variable value

	var username string = "Manish"
	fmt.Println(username)                              // To print the username i.e, Manish
	fmt.Printf("Variable is of type: %T \n", username) // To print it's type i.e, String

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 123.456987454597568
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var smallFloat1 float64 = 123.456987454597568
	fmt.Println(smallFloat1)
	fmt.Printf("Variable is of type: %T \n", smallFloat1)

	// Default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var anotherVariable1 float32
	fmt.Println(anotherVariable1)
	fmt.Printf("Variable is of type: %T \n", anotherVariable1)

	var anotherVariable2 string
	fmt.Println(anotherVariable2)
	fmt.Printf("Variable is of type: %T \n", anotherVariable2)

	var isLoggedIn1 bool
	fmt.Println(isLoggedIn1)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn1)

	// implicit type i.e, declaring a variable without assigning its type
	// var variable name = value of the variable, no need for variable type here

	var username1 = "Manish" // variable type "string" is not written even though the code is working fine
	fmt.Println(username1)
	fmt.Printf("Variable is of type: %T \n", username1)

	var isLoggedIn2 = true
	fmt.Println(isLoggedIn2)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn2)

	var website = "https://github.com/money8203"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	// No var style i.e, no neet to type var, variable type just variable name and it's value
	// := '' helps in accomplising this
	// := is for declaration + assignment, whereas = is for assignment only. For example, var foo int = 10 is the same as foo := 10 .
	// we can use := operator only inside the method, not outside the method, we can only use proper variable declaration type outside the method

	username3 := "Manish"
	fmt.Println(username3)

	numberOfUsers := 30000.0
	fmt.Println(numberOfUsers)

	//to print the constant
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
