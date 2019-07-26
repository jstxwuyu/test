package main

import "fmt"

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := array[2:5] //新切片//容量还是10
	s1[1] = 666
	fmt.Println("s1=", s1)
	fmt.Println("array=", array) //会一起变成666的

	s2 := s1[2:7] //另外s1 的新切片/容量还是10
	s2[2] = 777
	fmt.Println("s2=", s2)
	fmt.Println("array=", array)

}
