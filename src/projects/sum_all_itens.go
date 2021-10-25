package main

import (
	"fmt"
)

func main() {
	x := [5]int{1, 2, 3, 4, 5}
	sum := 0
	for _, i := range x {
		sum += i
	}

	fmt.Println(sum)
}