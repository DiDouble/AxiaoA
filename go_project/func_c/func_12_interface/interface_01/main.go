package main

import "fmt"

//定义一个类型

type speaker interface {
	speak() //只要实现了speaker方法的都是spaker类型
}

type cat struct{}
type dog struct{}
type persion struct{}

func (c cat) speak() {
	fmt.Println("miaomiaomiao")
}
func (d dog) speak() {
	fmt.Println("miaomiaomiao")
}
func (p persion) speak() {
	fmt.Println("miaomiaomiao")
}

func da(x speaker) {
	x.speaker()
}

func main() {
	var c1 cat
	var d1 dog
	var p1 persion

	da(c1)
	da(d1)
	da(p1)

}
