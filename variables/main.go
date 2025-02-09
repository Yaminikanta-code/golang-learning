package main

import "fmt"

const LoginToken string = "asdfghjkl" // constant

// First letter should be capital for public variables

func main() {
	var username string = "John"
	fmt.Println("Hello", username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isTrue bool = true
	fmt.Println(isTrue)
	fmt.Printf("Variable is of type: %T \n", isTrue)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.1234
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// Default values and aliases
	// Go puts default values to the variables
	var anotherString string
	fmt.Println(anotherString)

	var anotherBool bool
	fmt.Println(anotherBool)

	var anotherInt int
	fmt.Println(anotherInt)

	var anotherUint uint
	fmt.Println(anotherUint)

	var anotherFloat float32
	fmt.Println(anotherFloat)

	//Implicit type
	var website = "google.com" // automatically infer the type of the variable based on the value
	fmt.Println(website)       // website is of type string. If we try website=3000 it will throw an error

	//No var style
	numberOfUsers := 3000000 // we can use this style of variable declaration only inside functions and methods
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}

//Read more about ranges here: https://golang.org/ref/spec#Numeric_types
