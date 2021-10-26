package main

import (
	"fmt"
)

type Student struct {
	Name   string
	Age    int64
	Grades []float64
}

func (s Student) GetGrades() []float64 {
	return s.Grades
}

func (s Student) GetBirthDate(age int64) int64 {
	return 2021 - age
}

func main() {
	student := Student{"Miguel", 16, []float64{9.3, 8.8, 7.8}}
	fmt.Println(student.GetGrades())
	fmt.Println(student.GetBirthDate(student.Age))
}