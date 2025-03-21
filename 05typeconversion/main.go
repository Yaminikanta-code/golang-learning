package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1-5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)

	// numRating, err := strconv.ParseFloat(input, 64) this gives error strconv.ParseFloat: parsing "3\n": invalid syntax
	// So we need to trim the string

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating:", numRating+1)
	}
}

// input := reader.ReadString('\n') this is an error. Can't initialize one variable with two values
