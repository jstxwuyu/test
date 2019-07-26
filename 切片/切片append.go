package main

import "fmt"

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("array=", array)

	//append在原切片末尾添加元素，容量不够会扩容
	s := append(array, 11, 10)
	fmt.Println("s=", s)

	//append通常以两倍容量扩容
}
