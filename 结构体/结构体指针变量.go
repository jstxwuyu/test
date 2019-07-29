package main

import "fmt"

type Student1 struct {
	id   int
	name string
	sex  byte //字符类型
	age  int
	addr string
}

func main() {
	var p1 *Student1 = &Student1{1, "mike", 'm', 18, "bj"}
	//每个都必须初始化//别忘了取地址
	fmt.Println("s1=", *p1)

	//指定成员初始化，没有初始化的自动赋值为0
	p2 := &Student1{name: "mike", addr: "bjjj"}
	fmt.Println("s2=", *p2)
}
