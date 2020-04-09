package main

import (
	"fmt"
)

type myint int

type yourint = int

func main() {
	var n myint
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	var m yourint
	m = 200

	fmt.Println(m)
	fmt.Printf("%T", m)
}
