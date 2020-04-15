package main

import (
	"fmt"
)

type animous struct {
	name string
}

func (a animous) move() {
	fmt.Printf("%s,会移动", a.name)
}

type dog struct {
	feet uint8
	animous
}

func (d dog) wangwang() {
	fmt.Printf("%s 在汪汪汪~\n", d.name)
}

func main() {
	d1 := dog{
		animous: animous{name: "周林"},
		feet:    4,
	}
	d1.move()
	d1.wangwang()
}
