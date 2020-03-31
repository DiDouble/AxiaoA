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

#### make函数创建切片

切片的本质就是一个框,框柱一个连续的内存.

```go
package main

import (
	"fmt"
)

func main() {
	a1 := make([]int, 2, 10)  //创建一个长度为2 ,容量为10 的一个切片.
	fmt.Printf("%d ,len(a1)=%d .cap(a1)=%d", a1, len(a1), cap(a1))
}

//运行结果:
[0 0] ,len(a1)=2 .cap(a1)=10
```

切片的复制拷贝

```go
	s3 := []int{1, 2, 3, 4}
	s4 := s3
	fmt.Println(s3, s4)
	s3[2] = 100
	fmt.Println(s3, s4)
}
//结果是:  当切片所引用的底层数组发生改变的时候,其他所有引用该数组的切片都会改变.
[1 2 3 4] [1 2 3 4]
[1 2 100 4] [1 2 100 4]
```

切片遍历

