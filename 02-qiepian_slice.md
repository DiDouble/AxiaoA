# 切片

1. 自定义切片

   ```go
   package main
   
   import (
   	"fmt"
   )
   
   func main() {
   	// 定义一个切片
   	var s1 []int
   	var s2 []string
   	//初始化切片
   	s1 = []int{1, 2, 3, 4}
   	s2 = []string{"name", "age", "addr"}
   	fmt.Println(s1, s2)
   
   	//长度和容量
   	fmt.Printf("len(s1): %d cap(s1):%d \n", len(s1), cap(s1))
   	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))
   }
   
   //结果
   [1 2 3 4] [name age addr]
   len(s1): 4 cap(s1):4 
   len(s2):3 cap(s2):3
   ```

   

2. 由数组得到切片

   ```go
   	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
   	s3 := a1[0:4]  //基于一个数组的切割，左包含，右不包含
   	fmt.Println("\n", s3)
   //结果
    [1 2 3 4]
   ```

   注: 错误

   ```shell
   # command-line-arguments
   AxiaoA\go_project\slice_01.go:20:2: undefined: a1
   AxiaoA\go_project\slice_01.go:21:8: undefined: a1
   #错误原因是因为没有定义a1
   ```

   

3. 切片再切片

   1. 就是拿已经切片了的来切片.
   2. 切片的容量是底层数组的容量

![image-20200330152402974](E:\AllProject\src\AxiaoA\images\image-20200330152402974.png)

修改切片内的数字

