package main

import "fmt"

var rpsMap = map[string]string{
	"rock":     "scissors",
	"scissors": "paper",
	"paper":    "rock",
}

func main() {
	fmt.Println(Rps("rock", "scissors"))
	fmt.Println(Rps("scissors", "rock"))
	fmt.Println(Rps("rock", "rock"))
}

func Rps(p1, p2 string) string {
	if p1 == p2 {
		return "Draw!"
	}

	if rpsMap[p1] == p2 {
		return "Player 1 won!"
	}

	return "Player 2 won!"
}
