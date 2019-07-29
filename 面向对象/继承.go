package main

import "fmt"

type Person struct {
	id   int
	name string
	sex  byte //字符类型
}
type Stu struct {
	Person //只有类型，没有名字，继承了person的字段
	uid    int
	goal   int
}

func main() {
	var s1 Stu = Stu{Person{1, "ikun", 'm'}, 15, 88}
	fmt.Println("s1=", s1)
	//s1= {{1 ikun 109} 15 88}

	s2 := Stu{Person{1, "ikun", 'm'}, 15, 88}
	fmt.Printf("s2=%+v\n", s2)
	//s2={Person:{id:1 name:ikun sex:109} uid:15 goal:88}
	//%+v,显示更详细
	s1.name = "golang"
	fmt.Println(s1.id, s1.name, s1.goal)
}
