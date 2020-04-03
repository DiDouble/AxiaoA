package main

import (
	"fmt"
	// "unicode"
	"strings"
)

func main() {
	// 	s1 := "存卡啊啊 额;啊123shaha ana"

	// 	var count int

	// 	for _, c := range s1 {
	// 		if unicode.Is(unicode.Han, c) {
	// 			count++
	// 		}
	// 	}
	// 	fmt.Println(count)
	// }
	//判断字符串的有多少个

	s2 := "how do you do"
	// s3 := strings.Split(s2 , " ")
	s3 := strings.Split(s2, " ")
	//     s2 := "how do you do"
	// s3 := strings.Split(s2 , " ")

	m1 := make(map[string]int, 10)
	for _, w := range s3 {
		// fmt.Println(string(w1))
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++
		}
		fmt.Println(w)
	}
	

}
