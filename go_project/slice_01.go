package main

import (
	"fmt"
)

func main() {
	// 定义一个切片
	var s1 []int
	var s2 []string
	//初始化切片
	s1 = []int{1, 2, 3, 4}
	s2 = []string{"name", "age", "addr"}
	fmt.Println(s1, s2)

	//长度和容量
	fmt.Printf("len(s1): %d cap(s1):%d \n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	s3 := a1[0:4]
	fmt.Println("\n", s3)

	a1[6] = 1300
	fmt.Printf("s6 = %d", s3[:7])
	fmt.Println(a1)

}
