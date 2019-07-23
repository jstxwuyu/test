package main

import "fmt"

func main() {
	var s string
	print("你是谁\n")
	fmt.Scan(&s)
	if s== "王思聪"{
		println("左手一个妹子，右手一个大妈")
	}else if s=="屌丝" {
		println("洗洗睡吧")
	}else if s=="吴煜" {
		println("是帅哥嗷")
	}else {
		println("请输入 王思聪 或 屌丝")
	}
}
