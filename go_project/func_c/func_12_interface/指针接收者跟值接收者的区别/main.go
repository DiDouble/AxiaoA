package main

import "fmt"

type animous interface {
	move()
	eat()
}

type cat struct {
	name string
	age  int
}

//使用值接受者实现接口的所有方法
// func (c cat) move() {
// 	fmt.Println("走路", c.name)
// }
// func (c cat) eat() {
// 	fmt.Println("吃饭", c.age)
// }
func (c *cat) move() {
	fmt.Println("走路", c.name)
}
func (c *cat) eat() {
	fmt.Println("吃饭", c.age)
}

func main() {
	var a1 animous
	c1 := cat{"tom", 4}
	c2 := cat{"tommmmm", 4}
	a1 = &c1
	fmt.Println(a1)
	a1 = &c2
	fmt.Println(a1)
}
