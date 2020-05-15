package main

import (
	"fmt"
	"runtime"
	"sync"
)

// func fff() {
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < 5; i++ {
// 		r1 := rand.Int()
// 		r2 := rand.Intn(10)
// 		fmt.Println(r1, r2)
// 	}
// }
var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("A:%d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("B:%d\n", i)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
