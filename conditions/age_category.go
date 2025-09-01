package main

import "fmt"

func main() {
	fmt.Println("Age category by age")

	var age int
	fmt.Print("Please, input age: ")
	fmt.Scanln(&age)

	if age < 12 {
		fmt.Println("It's a child")
	} else if age < 18 {
		fmt.Println("It's a teenager")
	} else if age < 60 {
		fmt.Println("It's a adult")
	} else {
		fmt.Println("It's a pensioner")
	}
}
