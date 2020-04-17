
package main

import (
	"fmt"
)

func main()  {
	//打印菜单
	fmt.Println("欢迎光临学生管理系统")
	fmt.Println(
		`
		1. 查看所有学生
		2. 新增学生
		3. 删除学生
		4. 退出
		`
	)
	fmt.Print("请输入你要干啥:")
	//等待用户做选择
	var choice int
	fmt.Scanln(&choice)
	fmt.Printf("你选择的选项是!\n",choice)

	//3. 执行对应的函数
	switch choice {
	case 1:
		showAllStudent()
	case 2:
		addStudent()
	case 3:
		deleteStudent()
	case 4:
		os.Exit(1) //退出,返回状态码是1
	default:
		fmt.Println("滚~~")

	}
	
}

