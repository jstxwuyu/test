package main

//面向过程
func ad01(a, b int) int {
	return a + b
}

//面向对象
//方法：给一个类型绑定一个函数
type long int

func (tmp long) ad2(other long) long { //tmp  叫接受者,接收者就是传递的一会参数
	return tmp + other
}
func main() {
	var res int
	res = ad01(1, 1) //普通函数调用
	println("res=", res)

	var a long = 2
	r := a.ad2(3)
	println("r=", r)
	//调用方法：变量名.函数（所需参数）
	//面向对象只是换了一种表现形式
}
