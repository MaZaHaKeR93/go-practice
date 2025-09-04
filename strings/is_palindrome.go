package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("a"))
	fmt.Println(isPalindrome("aba"))
	fmt.Println(isPalindrome("Abba"))
	fmt.Println(isPalindrome("hello"))
}

func isPalindrome(str string) bool {
	str = strings.ToLower(str)

	runes := []rune(str)
	left := 0
	right := len(str) - 1
	for left < right {
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}
	return true
}
