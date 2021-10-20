package main

import "fmt"

func main() {
	var str string
	str = "小明"
	fmt.Printf("%T,%s\n",str,str)

	v1 :='A' //ASCII编码
	v2 :="A"
	v3 :="B"
	v4 :="cccc"

	fmt.Printf("%T,%d\n",v1,v1)
	fmt.Printf("%T,%s\n",v2,v2)
	fmt.Println(v3+v4)
}
