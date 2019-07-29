package main

type Hum interface {
	sayhi()
}
type pp interface {
	Humaner //匿名字段，继承了sayhi
	sing(lrc string)
}

func main() {

}
