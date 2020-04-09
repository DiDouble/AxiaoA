package main

import "fmt"

type persion struct {
	name string
	age  int
}

//构造函数
// 返回的是结构体还是结构体指针
//当结构体比较大的时候,尽量使用结构体指针,减少程序的内存开销.
func newPersion(name string, age1 int) persion {
	return persion{
		name: name, //前面的是结构体的变量，后面的是方法参数传入的值
		age:  age1,
	}
}

type dog struct {
	name string
	age  int
}

func newdog(name string, age int) dog {
	return dog{
		name: name,
		age:  age,
	}
}
func (d dog) wang() {
	fmt.Printf("%s :wangwangwang ,%v", d.name ,d.age )
}
func main() {
	// p1 := newPersion("112", 11)
	// fmt.Println(p1)
	d1 := newdog("zhizhi", 11)
	d1.wang()

}
