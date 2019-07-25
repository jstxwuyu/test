package main

func test111() int { //函数被调用时，x才分配空间，才初始化为0
	var x int //没有初始化，值为0
	x++
	return x * x //调用完毕，x自动释放
}

//func main() {
//	println(test111())  //1
//	println(test111())   //1
//	println(test111())   //1
//	println(test111())   //1
//}

//函数的返回值是一个匿名函数，返回一个函数类型
func test2() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}

}
func main() {
	//返回值为一个匿名函数，返回一个函数类型
	f12 := test2() //f12来调用返回的匿名函数，f12来调用匿名函数
	println(f12()) //1
	println(f12()) //4
	println(f12()) //9
	println(f12()) //16
	//闭包调用时并不会被释放，只有闭包还在使用，它们就一直存在

}
