package main

import (
	"fmt"
	"os"
)

// funcopenFile(name string ,flag int ,perm FileMode)(*File, error)  {

// }
func main() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fileObj.Write([]byte("1111111"))
	fileObj.WriteString("222222")
	fileObj.Close()
}
