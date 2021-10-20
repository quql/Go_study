package main

import "fmt"

func main() {
	var num int
	num = 100
	fmt.Printf("num的数值:%d,num的地址:%p\n",num,&num)
	num=200
	fmt.Printf("num的数值:%d,num的地址:%p\n",num,&num)
	var num1 int
	num1 = num
	fmt.Printf("num1的数值:%d,num1的地址:%p\n",num1,&num1)

	var a,b,c int
	a=1
	b=2
	c=3
	fmt.Println(a+b+c)
}
