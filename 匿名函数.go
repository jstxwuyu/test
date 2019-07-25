package main

import "fmt"

func main() {
	a := 10
	b := "mike"
	f1 := func() { //匿名函数的定义，没有名字，还没调用
		println("a=", a)
		println("b=", b)
	}
	f1() //匿名函数的调用

	type FuncType func() //别名定义函数类型
	var f2 FuncType
	f2 = f1 //赋值，等价
	f2()    //调用函数

	func() {
		fmt.Printf("a=%d b=%s\n", a, b) //可以捕获到外面的变量
	}() //代表调用此匿名函数//传参

	//带参数的匿名函数
	func(i, j int) {
		fmt.Printf("i=%d ,j=%d\n", i, j)
	}(10, 20) //有参数的匿名函数直接调用
	//等价于

	f4 := func(i, j int) {
		fmt.Printf("i=%d ,j=%d\n", i, j)
	}

	f4(10, 20)
	//*****************************************
	//有参数有返回值
	x, y := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = i
			min = j
		}
		return
	}(10, 20)
	fmt.Printf("x=%d,y=%d\n", x, y)
}
