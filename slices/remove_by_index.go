package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Index for removing: ")
	var index int
	fmt.Scanln(&index)
	fmt.Println("Original slice: ", slice)
	if index == 0 {
		slice = slice[index+1:]
	} else {
		slice = append(slice[:index], slice[index+1:]...)
	}
	fmt.Println("New slice:      ", slice)
}
