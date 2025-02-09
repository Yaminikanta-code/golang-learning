package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")

	// comma ok or comma error syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
	fmt.Printf("Type of this rating is %T", input)
}

// Notice bufio and os starts with a small case letter while Stdin and NewReader starts with a capital letter
// This is important to understand what is public
// input, _ := reader.ReadString('\n') we are either going to get an input or an error
// input, _ := reader.ReadString('\n') keep reading until we get a new line
// Type of input is string
// input, err := reader.ReadString('\n') here we catch the error in err
// input, _ := reader.ReadString('\n') here we are ignoring the error
// For anything we ignore in Go we use _
// comma ok or comma error syntax works like try catch but Go expect us to hold the error in a variable
