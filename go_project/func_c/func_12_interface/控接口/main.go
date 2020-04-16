package main

import "fmt"

func show(a interface{}) {
	fmt.Printf("%v\n", a)
}

func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "ll"
	m1["age"] = 900
	m1["merried"] = true
	m1["hobby"] = [...]string{"唱", "跳", "rap"}
	fmt.Println(m1)

	show(m1)
	show(true)
}
