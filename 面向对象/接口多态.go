package main

import "fmt"

type Humaner interface {
	//方法，只有声明，灭有实现
	sayhi()
}
type St struct {
	name string
	id   int
}

//St实现了此方法
func (tmp *St) sayhi() {
	fmt.Printf("St[%s,%d] sayhi\n", tmp.name, tmp.id)
}

type Teacher struct {
	name string
	id   int
}

//Teacher 实现 了此方法
func (tmp *Teacher) sayhi() {
	fmt.Printf("teacher[%s,%d] sayhi\n", tmp.name, tmp.id)
}

type Mystr string

//Mystr 实现了此方法
func (tmp *Mystr) sayhi() {
	fmt.Printf("Mysqz[%s] sayhi\n", *tmp)
}

func whosay(i Humaner) {
	i.sayhi()
}

//只有一个函数，可以有多种表示
//定义一个普通函数，函数的参数为接口类型
func main() {
	s := &St{"MIEK", 8}
	t := &Teacher{"bj", 88}
	var str Mystr = "hello go"
	whosay(s)
	whosay(t)
	whosay(&str)

	//创建一个切片
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str
	for _, i := range x {
		i.sayhi()
	}
}
func main01() {
	//定义接口类型的变量
	var i Humaner
	//只要实现了此接口方法的类型，那么这个类型的变量（接收者类型）就可以给i赋值
	s := &St{"mike", 666}
	i = s
	i.sayhi()
	t := &Teacher{"mike", 666}
	i = t
	i.sayhi()

	var str Mystr = "hello go"
	i = &str
	i.sayhi()
}
