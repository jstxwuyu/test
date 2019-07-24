package main

import "fmt"

func funca(a, i int) {
	if a == 1 { //函数递归结束条件很重要
		println("a=", a)
		return
	}

	funca(a-1, i+1)
	println("a=", a, "i=", i)
}
func main() {
	var i int
	println("请输入i")
	fmt.Scan(&i)
	funca(i, 1)

}
