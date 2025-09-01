package main

import "fmt"

func main() {
	fmt.Println("Is it leap year?")

	var year int
	fmt.Print("Input year: ")
	fmt.Scanln(&year)

	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println("It's a leap year")
	} else {
		fmt.Println("It's not a leap year")
	}
}
