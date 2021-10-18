package main

import (
	"strconv"
	"os"
	"fmt"
	"bufio"
)

func main() {
	fmt.Println("Welcome!")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type your age: ")
	scanner.Scan()
	input, err := strconv.ParseInt(scanner.Text(), 10, 64)

	if err != nil {
		fmt.Println(err)
		
		return
	}

	if input < 18 && input > 14 {
		fmt.Println("Can ride with a parent.")
	} else if input < 14 {
		fmt.Println("Can't ride")
	} else {
		fmt.Println("You can ride alone")
	}
}