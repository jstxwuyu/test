package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var a [10]int
	n := len(a)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100)
		fmt.Printf("%d,", a[i])

	}
	fmt.Printf("\n")
	//冒泡排序，挨着的两个数比较，大于则交换
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	//打印排序结果
	println("排序之后的结果为：")
	for i := 0; i < n; i++ {
		fmt.Printf("%d,", a[i])
	}
	fmt.Printf("\n")
}
