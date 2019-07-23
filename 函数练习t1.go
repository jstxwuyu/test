package main

import "fmt"

func A1(){
	var a int
	println("a=",a)
}
func A2(a int){
	println("x=",a)
}
func A3(a int , b bool)  {
	fmt.Printf("x=%d,n=%v\n",a,b)
}
func A4(a,b int)  {
	fmt.Printf("x=%d v=%d\n",a,b)
}
func main() {
	A1()
	A2(666)
	A3(4,true)
	A4(1,3)
}
