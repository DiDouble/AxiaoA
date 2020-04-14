package main

import (
	"fmt"
)

// 给自定义类型方法
// 不能给别的包里面的类型添加方法，只能给自己包添加
type myInt int

func (m myInt) myintd() {

	fmt.Println("i am a int ")
}

func main() {
	m := myInt(100)
	m.myintd()

}
