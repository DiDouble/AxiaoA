package main

import (
	"fmt"
)

type persion2 struct {
	string
	int
}

type student struct {
	name string
	age  int
	addr address
}
type address struct {
	province string
	city     string
}
type persion struct {
	name string
	age  int
}
type teacher struct {
	persion
	address
}

type company struct {
	name string
	addr address
}

func main() {
	p1 := persion2{
		"周坤",
		18,
	}
	fmt.Println(p1.string)

	p2 := student{
		name: "lili",
		age:  19,
		addr: address{
			province: "shandong",
			city:     "威海",
		},
	}

	fmt.Println(p2, p2.name, p2.addr.city)

	p3 := teacher{
		persion: persion{
			name: "lll",
			age:  30,
		},
		address: address{
			province: "shanghai",
			city:     "chaoshan",
		},
	}

	fmt.Println(p3.name)

}
