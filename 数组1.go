package main

import "fmt"

func main() {
	var id [50]int
	for i := 0; i < len(id); i++ {
		id[i] = i + 1
		fmt.Printf("第%d个数是%d\n", i+1, id[i])
	}
}
