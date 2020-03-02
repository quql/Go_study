package main

import "fmt"

func main() {
	//for i :=58 ;i > 28;i--  {
	//	fmt.Println(i)
	//}
	//1-100的和
	//sum := 0;
	//for i := 1;i <= 100; i++  {
	//	sum+=i
	//}
	//fmt.Printf("数据的和为%d",sum)

	//打印1-100能被3整除不能被5整除的数字,统计打印数字的个数,每行打印五个
	count := 0
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			count++
			if count%5 == 0 {
				fmt.Println()
			} else {
				fmt.Print(i, "\t")
			}
		}
	}
	fmt.Printf("\n共有%d个数字", count)
}
