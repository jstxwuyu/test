package main

import (
	"fmt"
)

func main() {
	var a [3][4]int
	k := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			k++
			a[i][j] = k
			fmt.Printf("a[%d][%d]=%d\n", i, j, a[i][j])
		}
		fmt.Printf("\n")
	}
	b := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	//初始化
	fmt.Println("b=", b)

	c := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	//F部分初始化
	fmt.Println("c=", c)

	d := [3][4]int{2: {5, 6, 7, 8}}
	//指定部分初始化
	fmt.Println("d=", d)
}
