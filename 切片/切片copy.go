package main

import "fmt"

func main() {
	srcSlice := []int{1, 2}
	otdslice := []int{6, 6, 6, 6}
	copy(otdslice, srcSlice) //src拷到otd
	fmt.Printf("otd=", otdslice)
}
