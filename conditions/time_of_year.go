package main

import "fmt"

func main() {
	fmt.Println("Time of year by month")

	var monthNum int
	fmt.Printf("Input number of month: ")
	fmt.Scanln(&monthNum)

	switch monthNum {
	case 3, 4, 5:
		fmt.Println("It's spring")
	case 6, 7, 8:
		fmt.Println("It's summer")
	case 9, 10, 11:
		fmt.Println("It's autumn")
	case 12, 1, 2:
		fmt.Println("It's winter")
	}
}
