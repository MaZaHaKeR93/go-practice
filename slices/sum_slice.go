package main

import "fmt"

func main() {
	var sliceCount int
	fmt.Print("Input slice length: ")
	_, err := fmt.Scanln(&sliceCount)
	if err != nil {
		fmt.Println("You neet to input number :)")
		return
	}

	slice := make([]int, 0, sliceCount)

	for i := 0; i < sliceCount; i++ {
		var sliceVal int
		fmt.Printf("Input value %d: ", i+1)
		_, err := fmt.Scanln(&sliceVal)
		if err != nil {
			fmt.Println("You neet to input number :)")
			return
		}
		slice = append(slice, sliceVal)
	}

	totalAmount := 0
	for _, v := range slice {
		totalAmount += v
	}

	fmt.Printf("Total slice amount: %d\n", totalAmount)
}
