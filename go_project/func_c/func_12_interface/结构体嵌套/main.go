package main

import "fmt"

type animous interface {
	mover
	eater
}
type mover interface{ move() }
type eater interface{ eat() }
type cat struct {
	name string
	age  int
}

//使用值接受者实现接口的所有方法
func (c cat) move() {
	fmt.Println("走路", c.name)
}
func (c cat) eat() {
	fmt.Println("吃饭", c.age)
}

func main() {

}
