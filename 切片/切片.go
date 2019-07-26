package main

import "fmt"

func main() {
	//切片比于不定长的数组，但是不是数组
	a := []int{1, 2, 3, 4, 5}
	s := a[0:3:5]
	fmt.Println("s=", s)
	fmt.Println("len(s)=", len(s)) //长度3
	fmt.Println("cap(s)=", cap(s)) //容量5

	x := a[1:4:5]
	fmt.Println("x=", x)
	fmt.Println("len(x)=", len(x)) //长度4-1
	fmt.Println("cap(x)=", cap(x)) //容量5-1
}
