package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name      string
	Age       int64
	BirthDate time.Time
	Height float64
}

func main() {
	var x Person = Person{"Miguel", 16, time.Now(),  1.74}
	fmt.Println(x.Name)
	fmt.Println(x.Age)
	fmt.Println(x.BirthDate)
	fmt.Println(x.Height)
}