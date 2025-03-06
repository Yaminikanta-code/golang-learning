package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers in golang")

	var ptr *int                                 // pointer to an int
	fmt.Printf("Type of pointer is: %T \n", ptr) // type is *int
	fmt.Println("Value of pointer is: ", ptr)    // default value is <nil>

	myNum := 23
	var ptr2 = &myNum                             // pointer to myNum & is for reference
	fmt.Printf("Type of pointer is: %T \n", ptr2) // type is *int
	fmt.Println("Value of pointer is: ", ptr2)    // printing memory address
	fmt.Println("Value of pointer is: ", *ptr2)   // dereferencing prints the value inside pointer 23
	fmt.Println("Value of variable is: ", myNum)  // prints value 23

	*ptr2 = *ptr2 * 2
	fmt.Println("New Value of variable is: ", myNum) // prints value 46
	fmt.Println("New Value of pointer is: ", *ptr2)  // prints value 46
}

/* There is a problem in programming languages sometimes when we pass on some variables
   into variety of functions, classes, objects etc, sometimes these variables are not
   actually passed on, but rather a copy of the variable is passed on. This is sometimes
   create some of the irregularities in the program. To avoid this problem we use pointers.
   A pointer is a variable that holds the memory address of another variable. */
