package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//输入
	//var x int
	//var y float64
	//fmt.Println("请输入一个整数,一个浮点类型")
	//fmt.Scanln(&x,&y)//通过操作地址复制
	//fmt.Printf("x的数值:%d,y的数值:%f\n",x,y)
	//
	//fmt.Scanf("%d,%f",&x,&y) //通过指定格式输入
	//fmt.Printf("x的数值:%d,y的数值:%f\n",x,y)
	fmt.Println("请输入一个字符串: ")
	reader:=bufio.NewReader(os.Stdin)
	s1,_:=reader.ReadString('\n')
	fmt.Println("读到的数据: ",s1)
}