package main

import "fmt"

func main() {
	str :="abcd"
	//for i:=0;i<len(str) ;i++  {
	//	fmt.Printf("str[%d]=%c",i,str[i])
	for i ,data :=range str{
		fmt.Printf("str[%d]=%c\n",i,data)
	}
}
