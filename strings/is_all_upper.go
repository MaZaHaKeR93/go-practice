package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(isUpperCase("abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(isUpperCase("ABCDEFGHIJKLMNOPQRSTUVWXYz"))
	fmt.Println(isUpperCase("WHAT DOES THE FOX SAY"))
}

func isUpperCase(str string) bool {
	if str == "" {
		return true
	}

	for _, v := range str {
		if unicode.IsLetter(v) && !unicode.IsUpper(v) {
			return false
		}
	}

	return true
}
