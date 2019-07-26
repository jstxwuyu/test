package main

import "fmt"

func main() {
	//自动推导类型
	a := []int{1, 2, 3}
	fmt.Println("a=", a)

	//借助make函数
	//make（切片类型,长度,容量）
	b := make([]int, 5, 10)
	fmt.Println("b切片数据")
	fmt.Println("b=", b)
	fmt.Println("len=", len(b))
	fmt.Println("cap=", cap(b))

	c := make([]int, 5) //不指定容量，容量和长度一样
	fmt.Println("c切片数据")
	fmt.Println("b=", c)
	fmt.Println("len=", len(c))
	fmt.Println("cap=", cap(c))
}
