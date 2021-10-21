package main

import (
	"fmt"
)

func main() {
	var x [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(x)
	var y []int = x[1:3]
	fmt.Println(y)

	// creating slices
	var z []string = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y"}

	z = append(z, "z")

	b := make([]int, 5)

	fmt.Println(b)
}