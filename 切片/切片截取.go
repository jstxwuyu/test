package main

import "fmt"

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//[low:high:max],len=high-low,cap=max-low
	s1 := array[:] //[0:len(array):len(array)]
	fmt.Println("s1=", s1)
	fmt.Println("len=", len(s1))
	fmt.Println("cap=", cap(s1))

	//操作元素方式和数组一样
	da := array[1]
	fmt.Println("da=", da)

	//截取
	s2 := array[3:6:7] //a3 a4 a5   len = 6-3=3   cap = 7-3=4
	fmt.Println("s2=", s2)
	fmt.Println("s2len=", len(s2))
	fmt.Println("s2cap=", cap(s2))

	//tip2
	s3 := array[:6] //0到6，长度6，容量和原切片一样10-0
	fmt.Println("s3=", s3)
	fmt.Println("s3len=", len(s3))
	fmt.Println("s3cap=", cap(s3))
}
