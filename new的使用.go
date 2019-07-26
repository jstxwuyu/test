package main

func main() {
	//老办法
	//a:=10
	//    var p *int
	//p=&a

	//new方法
	var p *int
	p = new(int) //p是*int 指向int
	*p = 89
	println("*p = ", *p)

	q := new(int)
	*q = 98 //自动推导类型
	println("*p = ", *q)

}
