package main

import (
	"fmt"
)

func main() {
	var x []int = []int{1, 2, 3, 4, 5, 4}

	for i, e1 := range x {
		for j, e2 := range x {
			if e1 == e2 && i != j {
				fmt.Println(e1)
			}
		}
	}
}