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
	fmt.Printf("%s :wangwangwang ,%v \n", d.name, d.age)
}

func (p persion) guonian() {
	p.age++
}

func (p *persion) zhenguonian() {
	p.age++
}
func main() {
	// p1 := newPersion("112", 11)
	// fmt.Println(p1)
	d1 := newdog("zhizhi", 11)
	d1.wang()
	p1 := newPersion("weilai", 18)
	p1.guonian()
	fmt.Println("\n", p1.age)
	p1.zhenguonian()
	fmt.Println("\n", p1.age)

}
