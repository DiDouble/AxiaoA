package main

//goroutine
import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() //goroutine 结束就登记-1
	fmt.Printf("Hello  %v\n", i)
}

func main() {
	for i := 0; i < 10; i++ {
		// go hello(i)
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	fmt.Println("main")
}
