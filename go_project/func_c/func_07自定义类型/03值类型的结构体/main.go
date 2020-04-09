package main

import (
	"fmt"
)

type persion struct {
	name string
	age  int
}

func fileName(p persion) {
	p.age = 24
	fmt.Println("fileName", &p.age)
}
func fileName2(p *persion) {
	(*p).age = 24
	fmt.Println("fileName", &p.age)
}

type x struct {
	a int8
	b int8
	c int8
}

func main() {
	var p persion
	p.name = "ll"
	p.age = 22
	// fileName(p)
	// fmt.Println(&p.age)
	// fileName2(&p)
	// fmt.Println(p.age)
	// fmt.Println(&p.age)
	// var p2 = new(persion)
	// p2.name = "112"
	// fmt.Printf("%T\n", p2)
	// fmt.Printf("%v", p2.name)
	// var p3 = persion{
	// 	name: "uuu",
	// 	age:  11,
	// }
	// fmt.Printf("%v \n", p3)
	// p4 := persion{
	// 	"xiao",
	// 	11,
	// }
	// fmt.Printf("%#v \n", p4)

	m := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}
	fmt.Printf("%p\n", &(m.a))
	fmt.Printf("%p\n", &(m.b))
	fmt.Printf("%p\n", &(m.c))

}
