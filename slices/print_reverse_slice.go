package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Original slice: ", slice)
	fmt.Printf("Reverse slice: ")
	for i := len(slice) - 1; i >= 0; i-- {
		fmt.Print(" ", slice[i])
	}
	fmt.Println()
}
