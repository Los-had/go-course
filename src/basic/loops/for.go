package main

import "fmt"

func main() {
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}
	fmt.Println("-------------------------------------------")
	// other way to use for loops
	y := 0
	for y <= 15 {
		fmt.Println(y)
		y++
	}
	fmt.Println("-------------------------------------------")
	// prints the divisible by 3 numbers in range of 0 to 1000
	for z := 1; z <= 1000; z++ {
		if z % 3 == 0 {
			fmt.Println(z)
		} else {
			continue
		}
	} 
}