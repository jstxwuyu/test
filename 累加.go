package main

//func leijia() (sum int) {
//	for i := 1; i <= 100; i++ {
//		sum += i
//	}
//	return
//}
func leijia(i int) int {
	if i == 1 {
		return 1
	}

	return i + leijia(i-1)
}

//100+99+98+97+++++++leijia(1)
//i+i+i+i+i+i+i++++++leijia(i-1)//leijia(1)
func main() {
	var sum int
	//sum = leijia()
	sum = leijia(100)
	println("sum= ", sum)
}
