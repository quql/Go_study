package main

import "fmt"

func main() {
	a:=10

	fmt.Println("a的数值为:",a)
	fmt.Printf("a的数据类型为:%T\n",a)
	fmt.Printf("a的地址:%p\n",&a)

	//声明一个变量指针,存储a的地址
	var p1 *int;
	fmt.Println(p1) //<nil>空指针
	p1 = &a
	fmt.Println(p1)
	fmt.Println("p1的数值为:",*p1)
}