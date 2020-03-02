package main

import "fmt"

func main() {
	//type1:定义变量,然后赋值
	var num1 int
	num1 = 20
	fmt.Printf("num1的数值是:%d\n",num1)
	//写在一行
	var num2 int = 50
	fmt.Printf("num2的数值是:%d\n",num2)
	//类型推断

	var name = "李四"
	fmt.Printf("类型是: %T,数值是:%s\n",name,name)

	//第三种,简短定义
	sum :=10
	fmt.Println(sum)

	//多个变量通知定义
	var a, b, c int
	a = 1
	b = 2
	c = 3
	fmt.Println(a,b,c)

	var q, w, e = 100,3.14,"GO"
	fmt.Println(q,w,e)
	//批量声明
	var (
		stuname = "张三"
		age = 18
		sex = "女"
	)
	fmt.Printf("学生姓名:%s,年龄:%d,性别:%s\n",stuname,age,sex)
}