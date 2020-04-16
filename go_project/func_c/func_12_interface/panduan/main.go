package main

import (
	cara "AxiaoA/go_project/func_c/func_12_interface/func_package"
	"fmt"
)

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str1, ok := a.(string)
	if !ok {
		fmt.Println("不是string类型")
	} else {
		fmt.Printf("%s是string类型", str1)
	}
}

func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	// str1, ok := a.(string)
	// if !ok {
	// 	fmt.Println("不是string类型")
	// } else {
	// 	fmt.Printf("%s是string类型", str1)
	// }
	switch a.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	}
}

func main() {
	// assign(100)
	// assign2(100)
	str := cara.Add(100, 200)
	fmt.Println(str)
}
