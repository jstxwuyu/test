package main

import (
	"fmt"
)

func MAXmin(a, b int) (max, min int) {
	if a > b {
		max = a
		min = b
		return
	} else {
		max = b
		min = a
		return
	}

}
func main() {
	var i, j int
	println("请输入：i=")
	fmt.Scan(&i)
	println("j=")
	fmt.Scan(&j)
	max, min := MAXmin(i, j)
	fmt.Printf("max=%d\nmin=%d\n", max, min)
}
