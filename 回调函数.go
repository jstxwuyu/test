package main

type FuncType func(int, int) int

func minu(a, b int) int { //减法
	return a - b
}

func add(a, b int) int { //加法
	return a + b
}
func Calc(a, b int, fTest FuncType) (result int) { //回调函数的优势p52
	println("calc")      //拓展性强，多态
	result = fTest(a, b) //函数还没有实现，先有了想法
	return
}
func main() {
	a := Calc(1, 1, add)  //d调用加法
	b := Calc(1, 1, minu) //减法
	println("a=", a)
	println("b=", b) //灵活性高，可以只修改主函数，其他什么都不要改
}
