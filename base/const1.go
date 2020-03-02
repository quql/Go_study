package main

import "fmt"

func main() {
	//定义常量
	const PATH string = "www.baidu.com"
	const PI = 3.14

	fmt.Println(PATH)

	//定义一组常量
	const C1, C2, C3 = 100, 500, "hah"

	const (
		NAME   = 1
		FNAME  = 2
		GNAMAE = 3
	)

	//一组常量中,没有初始值,默认和上一行一致
	const (
		a int = 1
		b
	)
	fmt.Printf("%T,%d\n", a, a)
	fmt.Printf("%T,%d\n", b, b)

	const (
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 3
	)
}
