package main

import "fmt"

type Pers struct {
	id   int
	name string
	sex  byte //字符类型
}
type Stus struct {
	*Pers //只有类型，没有名字，继承了person的字段
	uid   int
	goal  int
}

func main() {
	var s1 Stus = Stus{&Pers{1, "ikun", 'm'}, 15, 88}
	fmt.Println(s1.id, s1.name)

	var s2 Stus
	s2.Pers = new(Pers)
}
