package main

import "fmt"

func main() {
	var s string
	print("你是谁（是王思聪还是屌丝）")
	fmt.Scan(&s)
	if s== "王思聪"{
		println("左手一个妹子，右手一个大妈")
	}
	if s=="屌丝" {
		println("洗洗睡吧")
	}
	if s!="王思聪"  &&   s!="屌丝" && s!="吴煜"{
		println("请输入 王思聪 或 屌丝")
	}
	if s=="吴煜" {
		println("是帅哥嗷")
	}
}
