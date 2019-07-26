package main

import "fmt"

//append切片容量变化
func main() {
	s := make([]int, 0, 1)
	oldCap := cap(s)
	for i := 0; i < 8; i++ {
		s = append(s, i)
		if newCap := cap(s); oldCap < newCap {
			fmt.Printf("cap:%d===>%d\n", oldCap, newCap)
			oldCap = newCap
		}
	}
}
