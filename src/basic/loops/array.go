package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	fmt.Println(sum)
}