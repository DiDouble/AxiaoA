package main

import (
	"fmt"
)

func funcA() {
	fmt.Println("A")
}
func funcB() {
	defer func() {
		err := recover()
		fmt.Println("解决错误")
		fmt.Println(err)
		fmt.Println("解决错误")
	}()
	panic("严重错误")
	fmt.Println("B")
}
func funcC() {
	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()
}
