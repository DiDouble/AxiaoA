//defer
package main

import (
	"fmt"
)

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("1111") //defer把后面的语句延迟到函数即将返回的时候再执行
	defer fmt.Println("2222")
	defer fmt.Println("3333")

}

func f1() int {
	x := 5
	defer func() { //函数内不能嵌套带有函数名的函数,但是可以声明一个没有函数名的匿名函数
		x++
		fmt.Println(x)
	}() //匿名函数后面必须有()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() { x++ }()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func main() {
	// deferDemo()
	// fmt.Println(f1())
	// fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	// f1()

}
