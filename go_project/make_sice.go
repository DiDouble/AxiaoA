package main

import (
	"fmt"
)

func main() {
	// 	a1 := make([]int, 2, 10)
	// 	fmt.Printf("%d ,len(a1)=%d .cap(a1)=%d", a1, len(a1), cap(a1))
	//
	s3 := []int{1, 2, 3, 4}
	s4 := s3
	fmt.Println(s3, s4)
	s3[2] = 100
	fmt.Println(s3, s4)
}
