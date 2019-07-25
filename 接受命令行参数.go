package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args
	n := len(list)
	println("n=", n)

	for i := 0; i < n; i++ {
		fmt.Printf("list[%d]=%s\n", i, list[i])
	}

	//等价于
	for i, data := range list {
		fmt.Printf("list[%d]=%s\n", i, data)
	}
}
