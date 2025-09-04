package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(NearestSquare(20))
	fmt.Println(NearestSquare(22))
	fmt.Println(NearestSquare(9))
	fmt.Println(NearestSquare(1))
	fmt.Println(NearestSquare(0))
}

func NearestSquare(n int) int {
	if n < 0 {
		return 0
	}

	root := math.Sqrt(float64(n))
	a := int(math.Floor(root))
	b := int(math.Ceil(root))
	
	sqA := a * a
	sqB := b * b
	
	if n-sqA <= sqB-n {
		return sqA
	}

	return sqB
}
