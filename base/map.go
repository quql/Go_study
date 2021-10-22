package main

import (
	"fmt"
	"sort"
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

	//rand.Seed(time.Now().UnixNano())
	//
	//var scoreMap = make(map[string]int,100)
	//for i:=0;i<100;i++ {
	//	key := fmt.Sprintf("stu%02d",i)
	//	value := rand.Intn(99)
	//	scoreMap[key] = value
	//}
	//var keys = make([]string,0,100)
	//for key:=range scoreMap {
	//	fmt.Println(key)
	//	keys = append(keys,key)
	//}
	//fmt.Println(keys)
	//sort.Strings(keys)
	//fmt.Println(keys)
	//for _,v := range keys {
	//	fmt.Println(v,scoreMap[v])
	//}

	//var mapSlice = make([]map[string]string, 10)
	//for index,value :=range mapSlice {
	//	fmt.Printf("index:%d value:%v\n", index, value)
	//}
	//mapSlice[0] = make(map[string]string,3)
	//mapSlice[0]["name"] = "张三"
	//mapSlice[0]["pass"] = "123456"
	//for index,value:=range mapSlice {
	//	fmt.Printf("index:%d value:%v\n",index,value)
	//}
	//
	//var sliceMap = make(map[string][]string,6)
	//fmt.Println("init")
	//key:="中国"
	//value,ok:=sliceMap[key]
	//if !ok{
	//	value=make([]string,3)
	//}
	//value = append(value,"北京","上海")
	//sliceMap[key] = value
	//fmt.Println(sliceMap)
	//
	//var mapInit = map[string]string {"xiaoli":"湖南", "xiaoliu":"天津"}
	//fmt.Print(mapInit)

	map1 := make(map[int]string, 5)
	map1[1] = "www.topgoer.com"
	map1[2] = "rpc.topgoer.com"
	map1[5] = "ceshi"
	map1[3] = "xiaohong"
	map1[4] = "xiaohuang"
	sli := []int{}
	for k, _ := range map1 {
		sli = append(sli, k)
	}
	sort.Ints(sli)
	for i := 0; i < len(map1); i++ {
		fmt.Println(map1[sli[i]])
	}
}
