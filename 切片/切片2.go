package main

import "fmt"

func main() {
	//切片和数组的区别
	//数组方括号里的长度是固定的
	a := [5]int{}
	fmt.Printf("len=%d,cap=%d\n", len(a), cap(a))
	//数组是不能修改长度的，len和cap永远是5

	s := []int{}
	fmt.Printf("1len=%d,1cap=%d\n", len(s), cap(s))

	s = append(s, 11) //给切片追加一个成员
	fmt.Printf("len2=%d,cap2=%d\n", len(s), cap(s))
}
