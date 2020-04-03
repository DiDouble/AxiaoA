package main

import "fmt"

func sum(x int, y int) (ret int) {
	// return x * y
	return
}

func f3() int {
	ff3 := 3
	return ff3
}

func f4(x, y, z int) int {
	return x * y * z
}

func f5(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}
func main() {
	fmt.Println(sum(2, 3))
	ff5 := make([]int, 5, 5)
	fmt.Printf("%T  ,%d \n", ff5, ff5)
	for _, value := range ff5 {
		f5("id", value)
	}
	// fmt.Println(f5("name", ff5))

}
