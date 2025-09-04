package main

import "fmt"

func main() {
	arr := []int{3, 5, 8, 13}
	fmt.Println(EachCons(arr, 1))
	fmt.Println(EachCons(arr, 2))
	fmt.Println(EachCons(arr, 3))
	fmt.Println(EachCons([]int{}, 3))
}

func EachCons(arr []int, n int) (newArr [][]int) {
	for i := 0; i <= len(arr)-n; i++ {
		newArr = append(newArr, arr[i:i+n])
	}
	return newArr
}
