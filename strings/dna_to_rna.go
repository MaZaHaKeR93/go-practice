package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(DNAtoRNA("GCAT"))
	fmt.Println(DNAtoRNA("TTTT"))
	fmt.Println(DNAtoRNA("GACCGCCGCC"))
}

func DNAtoRNA(dna string) string {
	return strings.ReplaceAll(dna, "T", "U")
}
