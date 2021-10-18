package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the year that you was born: ")
	scanner.Scan()
	input, err := strconv.ParseInt(scanner.Text(), 10, 64)
	
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println("You will be", 2022 - input, "years in 2022")
}