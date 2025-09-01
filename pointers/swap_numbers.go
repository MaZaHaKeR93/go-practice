package main

import "fmt"

func main() {
	var firstNum, secondNum int
	fmt.Print("Input first number: ")
	fmt.Scanln(&firstNum)
	fmt.Print("Input second number: ")
	fmt.Scanln(&secondNum)

	swapNumbers(&firstNum, &secondNum)
	fmt.Printf("After swap: first number - %d, second number - %d", firstNum, secondNum)
}

func swapNumbers(firtsNum, secondNum *int) {
	tmpVar := *firtsNum
	*firtsNum = *secondNum
	*secondNum = tmpVar
}
