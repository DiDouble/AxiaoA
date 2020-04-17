package main

import (
	"fmt"
)

type people struct {
	name string 
	age int 
}

//方法一：
var p persion  //声明一个person类型的变量p
p.name = "11"
p.age = "18"
fmt.Println(p)
//方法二
persion {
	name: "22",
	age: 15,
}

// 给自定义类型方法
// 不能给别的包里面的类型添加方法，只能给自己包添加
type myInt int

func (m myInt) myintd() {

	fmt.Println("i am a int ")
}

func main() {
	m := myInt(100)
	var x  int 100
	var x = 100    //以上三个是一样的。
	m.myintd()

}
