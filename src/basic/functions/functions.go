package main

import "fmt"

func main() {
	printx(5)
}

func printx(x int) {
	fmt.Println(x)
	ans := add(5, 5)
	fmt.Println(ans)
}

func add(x, y int) int {
	return x + y
}