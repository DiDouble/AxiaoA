package main

import (
	"encoding/json"
	"fmt"
)

type persion struct {
	Name string
	Age  int
}

func main() {
	p1 := persion{
		Name: "zj",
		Age:  90000,
	}
	k, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("have err is %v\n", err)
		return
	}
	fmt.Println(k)
	fmt.Println(string(k))

	//反序列话
	str1 := `{"name": "理想","age":19}`
	var p2 persion
	json.Unmarshal([]byte(str1), &p2)
	fmt.Printf(" %v\n", p2)
}
