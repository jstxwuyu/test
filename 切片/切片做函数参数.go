package main

import (
	"fmt"
	"math/rand"
	"time"
)

//数组是值传递，切片是引用
func BULL(s []int) {
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	//打印排序结果
	//println("排序之后的结果为：")
	//for i := 0; i < n; i++ {
	//	fmt.Printf("%d,", s[i])
	//}
	//fmt.Printf("\n")
}
func Date(s []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100)
	}
}
func main() {
	var n int = 10
	s := make([]int, n)
	Date(s)
	fmt.Println("排序前：", s)
	BULL(s)
	fmt.Println("排序之后的结果为：", s)
}
