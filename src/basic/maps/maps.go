package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"apple":     5,
		"pear":      4,
		"pineapple": 10,
		"orange":    7,
	}
	mp["apple"] = 6
	mp["grapes"] = 3
	fmt.Println(mp)
	delete(mp, "pineapple")
	fmt.Println(mp)
	// create maps using the make function
	mp2 := make(map[string]int)
	fmt.Println(mp2)
}