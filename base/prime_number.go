package main

import (
	"fmt"
	"math"
)

func main() {
	/**
	打印出2-100内的素数,只能被1和自己整除的数
	 */

	for i:=2;i<=100;i++  {
		flag := true
		for j:=2;j<int(math.Sqrt(float64(i)));j++  {
			if i%j==0{
				flag = false
			}
		}
		if flag==true{
			fmt.Println(i)
		}
	}
}