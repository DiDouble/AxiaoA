package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := int32(2316)
	ret2 := string(i)
	fmt.Println(ret2)

	str2 := "1000000"
	ret1, err := strconv.ParseInt(str2, 10, 64)
	if err != nil {
		fmt.Println(111)
		return
	}
	fmt.Println(ret1)

}
