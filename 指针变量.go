package main

import "fmt"

func main() {
	var a int = 10
	fmt.Printf("&a = %v\n", &a) //地址
	fmt.Printf("a= %d \n", a)   //内存
	var p *int
	p = &a //指针变量指向谁，就把谁的地址赋值给指针变量
	fmt.Printf("p=%v , &a =%v\n", p, &a)
	*p = 666 //*p不是控制p的内存，而是控制p所指向的内存，即是a
	fmt.Printf("*p = %v , a = %v\n", *p, a)
}
