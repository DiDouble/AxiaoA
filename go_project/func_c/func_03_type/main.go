package main

import (
	"fmt"
)

func f1(x, y int) int {
	return x + y
}

func f2() int {
	return 10
}

func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func f5(x func() int) func(int, int) int {
	return f1
}

func main() {
	// f3(f2)

	f7 := f5(f2)
	// fmt.Println(f5(f2))
	fmt.Printf("%T \n", f7)

}
