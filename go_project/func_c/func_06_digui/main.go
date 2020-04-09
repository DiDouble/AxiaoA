package main

import (
	"fmt"
)

func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * (n - 1)
}

func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}

func main() {
	ret := taijie(4)
	fmt.Println(ret)
	// 	ret := f(7)
	// 	fmt.Println(ret)
}
