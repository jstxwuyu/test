package main

import (
	"fmt"
)

func myadc() int {
	return 123
}

//给返回值起个名字
func mycc() (result int) {
	//return 123
	result = 123
	return
}
func myacc() (x int, y int, z int) {
	x, y, z = 11, 22, 33
	return
}
func main() {
	var a int
	a = myadc()
	b := mycc()
	x, y, z := myacc()
	println("a=", a)
	println("b=", b)
	fmt.Printf("x=%d , y=%d ,z=%d", x, y, z)
}
