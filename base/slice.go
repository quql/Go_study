package main

import "fmt"

func main()  {
	var s1 []int
	if s1==nil {
		fmt.Println("是空")
	} else {
		fmt.Println("不是空")
	}
	arr:=[5]int{1,2,3,4,5}
	var s2 []int
	s2 = arr[1:4]
	fmt.Println(s2)
	s3:=arr[:]
	fmt.Println(s3)

	var slice0 []int = make([]int, 10)
	var slice1 []int = make([]int,10,10)
	fmt.Println(slice0)
	fmt.Println(slice1)

	data:=[...]int{0,1,2,3,4,5,6,7,8}
	s:=data[2:5]
	fmt.Println(s)
	s[0]+=100
	s[1]+=200
	fmt.Println(s)
	fmt.Println(data)
}
