package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileobj, err := os.Open("./main.go") //打开文件
	if err != nil {
		fmt.Printf("%v", err)
		return

	}
	defer fileobj.Close() //这里为了防止忘记关,所以讲这个提前写在这里

	//读文件
	//var tmp = make([]byte,128) //指定长度.
	var tmp [128]byte
	for {
		n, err := fileobj.Read(tmp[:])
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("%v\n", err)
			return

		}
		fmt.Println(n)
		fmt.Println(string(tmp[:n]))

		if n < 128 {
			return
		}
	}

}
