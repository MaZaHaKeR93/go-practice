package main

import "fmt"

func main() {
	var num int
	fmt.Print("Input number: ")
	fmt.Scanln(&num)

	i := 1
	total := 0
	for i <= num {
		total += i
		i++
	}

	fmt.Printf("Total amount: %d\n", total)
}
