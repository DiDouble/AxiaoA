package main

import (
	"fmt"
)

type person struct { //创建一个自定义的类型，定义了一个person类型
	name   string
	age    int
	gender string
	hoddy  []string
}

func main() {

	// var p person
	// p.name = "what"
	// p.age = 11
	// p.gender = "9000"
	// p.hoddy = []string{"皮球", "气球"}
	// fmt.Println(p)
	// fmt.Println(p.name)

	var s struct {
		name string
		age  int
	}
	s.name = "aa"
	s.age = 11
	fmt.Println(s.name)
}
