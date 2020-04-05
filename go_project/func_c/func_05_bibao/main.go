package main

import (
	"fmt"
)

func adder() func(int) int {
	var x = 100
	return func(y int) int {
		x += y
		return x
	}

}

func f1(f func()) {
	fmt.Println("This is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("This is f2")
	fmt.Println(x + y)
}

func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		f(x, y)
	}
	return tmp
}

func main() {
	// ret := adder()

	// ret1 := ret(200) //
	// fmt.Println(ret1)
	var x = 100
	var y = 200
	ret := f3(f2, x, y)
	f1(ret)
}
