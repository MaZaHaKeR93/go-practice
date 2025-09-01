package main

import "fmt"

func main() {
	var num int
	fmt.Print("Input number: ")
	fmt.Scanln(&num)

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Print(i, " ")
		}
	}
}
