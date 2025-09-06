package main

import "fmt"

func main() {
	fmt.Println(IsValidWalk([]rune{'n', 's', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}))
	fmt.Println(IsValidWalk([]rune{'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e'}))
	fmt.Println(IsValidWalk([]rune{'w'}))
}

func IsValidWalk(walk []rune) bool {
	if len(walk) != 10 {
		return false
	}

	x, y := 0, 0

	for _, direction := range walk {
		switch direction {
		case 'n':
			x++
		case 's':
			x--
		case 'e':
			y++
		case 'w':
			y--
		}
	}

	return x == 0 && y == 0
}