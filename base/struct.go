package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

type student struct {
	name string
	age  int
}
func main() {
	//var user struct {name string;age int}
	//user.name="张三"
	//user.age=20
	//fmt.Print(user)
	//
	//person:=person{
	//	city:"22",
	//	age:18,
	//	name:"张三",
	//}
	//fmt.Print(person)
	//stus:= []student{
	//	{name: "张三", age:  10},
	//	{name: "张三01", age:  10},
	//	{name: "张三02", age:  10},
	//}
	//m:=make(map[string]*student)
	//for _,v:=range stus{
	//	m[v.name] = &v
	//}
	//
	//for key,value:=range m {
	//	fmt.Println(key,value.name)
	//}

	var ce []student
	ce = []student{
		student{"张三01",10},
		student{"张三02",20},
		student{"张三03",30},
	}
	demo(ce)
	fmt.Println(ce)
}

func demo(ce[]student) {
	ce[1].age=999
}
