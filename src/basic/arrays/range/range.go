package main

import (
	"fmt"
)

func main() {
	var x []int = []int{1, 2, 3, 4, 5}

	for i, e := range x {
		fmt.Println(i, ":", e)
	}
}