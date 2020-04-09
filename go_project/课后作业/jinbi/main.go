package main

import (
	"fmt"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emolie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() (left int) {
	for _, name := range users {
		for _, c := range name {
			switch c {
			case 'e', 'E':
				distribution[name]++
				coins--
			case 'i', 'I':
				distribution[name] += 2
				coins -= 2
			case 'o', 'O':
				distribution[name] += 3
				coins -= 3
			case 'u', 'U':
				distribution[name] += 4
				coins -= 4
			}
		}
	}
	left = coins
	return left
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下", left)
	for k, num := range distribution {
		fmt.Printf("%v,%d\n", k, num)
	}
}
