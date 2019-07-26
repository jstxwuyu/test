package main

import "fmt"

func swap(a, b int) {
	a, b = b, a
	fmt.Printf("swap:a=%d,b=%d\n", a, b) //20,10交换成功
}
func main() {
	a, b := 10, 20 //通过函数交换a b 的内容
	//变量本身传递，值传递
	swap(a, b)
	fmt.Printf("main:a=%d,b=%d\n", a, b) //10,20，没有交换
}
