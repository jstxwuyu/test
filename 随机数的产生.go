package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//设置种子//只需要一次

	//如果种子参数一样，每次运行程序一样
	//以当前时间作为种子参数
	rand.Seed(time.Now().UnixNano())
	println("time=", time.Now().Unix()) //unix时间戳
	println("time=", time.Now().UnixNano())
	//产生随机数
	for i := 0; i < 5; i++ {
		//fmt.Println("rand=", rand.Int())     //随机很大的数
		fmt.Println("rand=", rand.Intn(100)) //限制在100内的数
	}

}
