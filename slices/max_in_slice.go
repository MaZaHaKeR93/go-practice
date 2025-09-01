package main

import "fmt"

func main() {
	slice := []int{4, 2, 8, 10, 4, 6, 44, 15, 22}
	maxElement := slice[0]
	for _, v := range slice {
		if v > maxElement {
			maxElement = v
		}
	}
	fmt.Printf("Max element: %d\n", maxElement)

}