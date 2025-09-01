package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is calculator")

	var firstValue, secondValue float64
	var operationSign string
	fmt.Print("Input first value: ")
	fmt.Scanln(&firstValue)
	fmt.Print("Input second value: ")
	fmt.Scanln(&secondValue)
	fmt.Print("Input operation sign (+, -, *, /): ")
	fmt.Scanln(&operationSign)

	switch operationSign {
	case "+":
		fmt.Printf("%g %s %g = %g\n", firstValue, operationSign, secondValue, firstValue+secondValue)
	case "-":
		fmt.Printf("%g %s %g = %g\n", firstValue, operationSign, secondValue, firstValue-secondValue)
	case "*":
		fmt.Printf("%g %s %g = %g\n", firstValue, operationSign, secondValue, firstValue*secondValue)
	case "/":
		if secondValue == 0 {
			fmt.Println("You can't divide by zero!")
		} else {
			fmt.Printf("%g %s %g = %g\n", firstValue, operationSign, secondValue, firstValue/secondValue)
		}
	}
}
