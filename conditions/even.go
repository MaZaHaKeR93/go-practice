package main

import "fmt"

func main() {
	fmt.Println("Ever or not?")
	fmt.Print("Input some number: ")

	var num int
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Println("Is even number")
	} else {
		fmt.Println("Is odd number")
	}
}
