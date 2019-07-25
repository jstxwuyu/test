package main

import "fmt"

func main() {
	a := 10
	str := "mice"

	func() {
		a = 333
		str = "nol"
		fmt.Printf("内  a = %d ,str = %s \n", a, str)
	}() //调用
	fmt.Printf("外  a = %d ,str = %s \n", a, str)
	//闭包引用方式引用外部变量
	//外部变量一起变化

}
