package main

import (
	"fmt"
	"os"
)

func selectStudent() {

}
func addStudent() {

}
func deleteStudent() {

}

func main() {
	fmt.Println("欢迎来到学生管理系统")
	fmt.Println(`
		1. 查询学生
		2. 增加学生
		3. 删除学生
		4. 退出
		`)
	fmt.Print("你要干啥:")
	var casina int
	fmt.Scanln(&casina)
	fmt.Printf("你输入的是%d", casina)
	switch casina {
	case 1:
		selectStudent()
	case 2:
		addStudent()
	case 3:
		deleteStudent()
	case 4:
		os.Exit(1)
	default:
		fmt.Println("滚~~")
	}
}
