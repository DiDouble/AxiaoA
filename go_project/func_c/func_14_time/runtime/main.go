package main

import (
	"fmt"
	"runtime"
)

func main() {
	pc, file, line, ok := runtime.Caller(0) //file指的是文件的地址,line是多少行被调用.
	if !ok {
		fmt.Println("failed")
		return
	}
	pc1 := runtime.FuncForPC(pc).Name() //这里返回调用的函数是啥
	fmt.Println(pc1)
	fmt.Println(file)
	fmt.Println(line)
}
