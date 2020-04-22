package main

import (
	"fmt"
	"os"
)

func f1() {
	fileObj, err := os.OpenFile("./sb.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v ", err)
		return
	}
	defer fileObj.Close()
	fileObj.Seek(1, 0) //光标向后移动一个字节
	var ret [1]byte

	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(ret[:n]))
}

func main() {
	f1()
}
