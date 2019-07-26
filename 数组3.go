package main

import (
	"fmt"
)

func main() {
	//数组的初始化==定义同时赋值
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	for i, data := range a {
		fmt.Printf("第%d个 ,数据%d\n", i, data)

	}
	for i, data := range b {
		fmt.Printf("第%d个 ,数据%d\n", i, data)

	}
}
