package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	Gender string
	Name   string
}

type Class struct {
	Title string
	Students []*Student
}
func main() {
	c := Class{
		Title:    "101",
		Students: make([]*Student,0,100),
	}
	for i:=0; i<10; i++  {
		stu:=&Student{
			ID:     i,
			Gender: "男",
			Name:   fmt.Sprintf("stu%02d",i),
		}
		c.Students = append(c.Students,stu)
	}
	data,err:=json.Marshal(c)
	if err!=nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)

	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err1:=json.Unmarshal([]byte(str),c1)
	if err1 != nil {
		fmt.Print(err1)
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}
