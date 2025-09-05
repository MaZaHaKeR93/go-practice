package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ReverseWords("The quick brown fox jumps over the lazy dog."))
	fmt.Println(ReverseWords("apple"))
	fmt.Println(ReverseWords("a b c d"))
	fmt.Println(ReverseWords("double  spaced  words"))
	fmt.Println(ReverseWords("stressed desserts"))
	fmt.Println(ReverseWords("hello hello"))
}

func ReverseWords(str string) string {
	words := strings.Split(str, " ")
	for i, word := range words {
		bytes := []byte(word)
		for l, r := 0, len(bytes)-1; l < r; l, r = l+1, r-1 {
			bytes[l], bytes[r] = bytes[r], bytes[l]
		}
		words[i] = string(bytes)
	}
	return strings.Join(words, " ")
}
