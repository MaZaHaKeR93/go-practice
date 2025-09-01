package main

import "fmt"

func main() {
	var num int
	fmt.Print("Input some number: ")
	fmt.Scanln(&num)
	increase_by_10(&num)
	fmt.Printf("Result is: %d\n", num)
}

func increase_by_10(num *int) {
	*num = *num + 10
}
