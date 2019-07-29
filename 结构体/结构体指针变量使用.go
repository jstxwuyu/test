package main

import "fmt"

type Student3 struct {
	id   int
	name string
	sex  byte //字符类型
	age  int
	addr string
}

func main() {
	//指针有合法指向后，才操作成员
	//先定义一个普通结构体变量
	var s Student3
	//再定义一个指针变量，保存s的地址
	var p1 *Student3
	p1 = &s
	p1.id = 12
	(*p1).name = "mike"
	p1.age = 22
	fmt.Println("p1=", *p1)

	//
	p2 := new(Student3)
	p2 = &s
	p2.id = 12
	(*p2).name = "mike"
	p2.age = 22
	fmt.Println("p2=", *p2)
}
