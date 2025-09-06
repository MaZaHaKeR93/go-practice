package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(SpinWords("Welcome"))
	fmt.Println(SpinWords("to"))
	fmt.Println(SpinWords("CodeWars"))
	fmt.Println(SpinWords("Hey fellow warriors"))
	fmt.Println(SpinWords("Burgers are my favorite fruit"))
	fmt.Println(SpinWords("Pizza is the best vegetable"))
}

func SpinWords(str string) string {
	words := strings.Split(str, " ")

	for i, word := range words {
		if len(word) < 5 {
			continue
		}

		bytes := []byte(word)
		for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
			bytes[i], bytes[j] = bytes[j], bytes[i]
		}
		words[i] = string(bytes)
	}

	return strings.Join(words, " ")
}
