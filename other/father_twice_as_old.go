package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(TwiceAsOld(36, 7))
	fmt.Println(TwiceAsOld(55, 30))
	fmt.Println(TwiceAsOld(42, 21))
	fmt.Println(TwiceAsOld(22, 1))
	fmt.Println(TwiceAsOld(29, 0))
}

func TwiceAsOld(dadYearsOld int, sonYearsOld int) int {
	return int(math.Abs(float64(dadYearsOld - (sonYearsOld * 2))))
}
