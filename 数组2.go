package main

import "fmt"

func main() {
	//定义一个数组，[10]int 和 [5]int是不同类型
	var a [10]int
	var b [5]int
	fmt.Printf("%d  %d  \n", len(a), len(b))
	a[0] = 1
	i := 2
	a[i] = 2 //a[2]=2

	//赋值
	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}
	//打印
	for i, data := range a {
		fmt.Printf("a[%d]=%d\n", i, data)
	}
}
