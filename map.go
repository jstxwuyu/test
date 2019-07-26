package main

import "fmt"

func main() {
	var m1 map[int]string
	fmt.Println("m1=", m1)
	//对于map只有len 没有cap
	//通过make创建
	m2 := make(map[int]string, 10) //指定容量
	m2[1] = "mike"
	m2[2] = "go"
	fmt.Println("m2=", m2)
	fmt.Println("len", len(m2))

	//定义初值
	//键值唯一
	m4 := map[int]string{1: "mike", 2: "go", 3: "c33"}
	fmt.Println("m4=", m4)
}
