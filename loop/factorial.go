package main

import "fmt"

func main() {
	var num int
	fmt.Print("Input number: ")
	fmt.Scanln(&num)
	total := 1
	for i := 1; i <= num; i++ {
		total *= i
	}
	fmt.Printf("Factorial %d = %d\n", num, total)
}
