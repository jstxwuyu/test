package main

import "fmt"

type Student2 struct {
	id   int
	name string
	sex  byte //字符类型
	age  int
	addr string
}

func main() {
	var s Student2
	//操作成员，需要用.运算符
	s.id = 2
	s.name = "jek"
	s.sex = 'm'
	s.age = 18
	fmt.Println("s=", s)
}
