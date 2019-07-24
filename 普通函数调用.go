package main

import "fmt"

func a(a int) {
	b(a - 1)
	println("a=", a)
}
func b(b int) {
	c(b - 1)
	println("b=", b)
}
func c(c int) {
	println("c=", c)
}
func main() {
	var x int
	fmt.Scan(&x)
	a(x)
}
