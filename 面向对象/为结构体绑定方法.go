package main

import "fmt"

type S1 struct {
	id   int
	name string
	sex  byte //字符类型
	age  int
	addr string
}

//带接受者的函数叫方法
func (tmp S1) Print() {
	fmt.Println("tmp=", tmp)
}

//通过一个函数，给成员赋值
func (p *S1) set(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
}
func main() {
	p := S1{1, "pl", 'm', 55, "bj"}
	p.Print()

	//定义一个结构体变量
	var p2 S1
	(&p2).set("yp", 'g', 88)
	p2.Print()
}
