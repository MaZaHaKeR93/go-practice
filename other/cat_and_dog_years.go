package main

import "fmt"

func main() {
	fmt.Println(CalculateYears(10))
	fmt.Println(CalculateYears(1))
	fmt.Println(CalculateYears(5))
	fmt.Println(CalculateYears(2))
}

func CalculateYears(years int) (result [3]int) {
	catYears, dogYears := 15, 15

	if years > 1 {
		catYears += 9
		dogYears += 9
	}

	if years > 2 {
		catYears += (years - 2) * 4
		dogYears += (years - 2) * 5
	}

	return [3]int{years, catYears, dogYears}
}
