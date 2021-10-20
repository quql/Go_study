package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main()  {
	//sorceMap :=make(map[string]int,8)
	//sorceMap["张三"]=10
	//sorceMap["李四"]=20
	//sorceMap["王五"]=30
	//fmt.Println(sorceMap)
	//delete(sorceMap,"王五")//删除key
	//fmt.Println(sorceMap)
	//fmt.Println(sorceMap["李四"])
	//
	//v,ok:=sorceMap["张三"]
	//fmt.Println(v,ok)
	//
	//for k,v:=range sorceMap {
	//	fmt.Println(k,v)
	//}

	rand.Seed(time.Now().UnixNano())

	var scoreMap = make(map[string]int,100)
	for i:=0;i<100;i++ {
		key := fmt.Sprintf("stu%02d",i)
		value := rand.Intn(99)
		scoreMap[key] = value
	}
	var keys = make([]string,0,100)
	for key:=range scoreMap {
		fmt.Println(key)
		keys = append(keys,key)
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)
	for _,v := range keys {
		fmt.Println(v,scoreMap[v])
	}
}
