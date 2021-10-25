package main

import "fmt"

func ReturnFunc(x string) func() {
	return func() {fmt.Println(x)}
}

func main() {
	x := test
	x()
	// creating functions

	func() {
		fmt.Println("test2")
	}()

	//test2()

	ReturnFunc("Hello, world!")()
}

func test() {
	fmt.Println("teste")
}