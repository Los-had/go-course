package main

import "fmt"

func main() {
	x := 5
	// fmt.Println(&x)
	y := &x
	fmt.Println(x, y)
	*y = 7
	fmt.Println(x, y)

}