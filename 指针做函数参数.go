package main

import "fmt"

func swap1(p1, p2 *int) {
	*p1, *p2 = *p2, *p1
	fmt.Printf("swap:a=%d,b=%d\n", *p1, *p2)
}
func main() {
	a, b := 10, 20 //通过函数交换a b 的内容
	//与变量本身传递，值传递不同
	swap1(&a, &b)                        //地址传递
	fmt.Printf("main:a=%d,b=%d\n", a, b) //也随之一起变换
}
