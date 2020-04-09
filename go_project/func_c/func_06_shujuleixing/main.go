package main

import (
	"fmt"
)

type myint int

func main() {
	var n myint
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)
}
