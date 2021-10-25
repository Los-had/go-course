package main

type Book struct {
	Name string
	Author *Author
	Description string
	Price float64
}

type Author struct {
	FirstName string
	LastName string
	Books []string
}

func main() {
	return
}