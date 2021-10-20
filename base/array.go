package main

import "fmt"

func main()  {
	a := [3] int {1,2}
	b := [...]int{1,2,3,4,5,6,7}
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10}, // 可省略元素类型。
		{"user2", 20}, // 别忘了最后一行的逗号。
	}
	fmt.Print(a,b)
	fmt.Print(d)

	e := [2][3]int{{1,2,3},{4,5,6}}
	fmt.Println(e)

	for k1,v1:=range e {
		for k2,v2 :=range v1 {
			fmt.Printf("(%d,%d)=%d",k1,k2,v2)
		}
		fmt.Println()
	}
	aa:=[5]int{1,3,5,8,7}
	myTest(aa,8)
}

func myTest(a [5]int,target int)  {
	for i:=0;i<len(a);i++  {
		for j:=i+1;j<len(a);j++  {
			if (a[i]+a[j])==target {
				fmt.Printf("(%d,%d)\n",i,j);
			}
		}
	}
}
