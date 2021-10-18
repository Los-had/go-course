package main

import "fmt"

func main() {
	answer := 3

	switch answer {
	case 1, -1:
		fmt.Println("1")
	case 2, -2:
		fmt.Println("2")
	case 3, -3:
		fmt.Println("3")
	default:
		fmt.Println("You don't find the right answer")
	}
}