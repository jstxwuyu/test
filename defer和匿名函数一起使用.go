package main

import "fmt"

func main() {
	a := 100
	b := 1020
	defer func() {
		fmt.Printf("a=%d b=%d\n", a, b)
	}() //最后才会执行

	a = 140
	b = 120
	fmt.Printf("wai  a=%d b=%d\n", a, b)
}
