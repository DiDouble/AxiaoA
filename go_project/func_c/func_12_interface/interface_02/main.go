package main

import "fmt"

type animal interface {
	move()
	eat()
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("cat move")
}
func (c cat) eat() {
	fmt.Println("cat eat", c.name)
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("激动")
}
func (c chicken) eat() {
	fmt.Println("吃鸡屎")
}

func main() {

	var a1 animal //定义一个接口类型的变量

	bc := cat{
		name: "tapqo",
		feet: 4,
	}
	a1 = bc
	a1.eat()

}
