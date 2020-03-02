package main

import (
	"fmt"
	"math"
)

/**
水仙花数 [100,999]
三位数 ,每个位数上的三次方的和与这个数相等,为水仙花数
*/

func main() {
	for i := 100; i < 1000; i++ {
		x := i / 100     //百位
		y := i / 10 % 10 //十位
		z := i % 10      //个位
		if math.Pow(float64(x), 3)+math.Pow(float64(y), 3)+math.Pow(float64(z), 3) == float64(i) {
			fmt.Println(i)
		}
	}
}
