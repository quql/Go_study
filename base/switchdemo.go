package main

import "fmt"

func main() {
	//num:=0
	//fmt.Println("请输入数字:")
	//fmt.Scanln(&num)
	//
	//switch num {
	//	case 1:
	//		fmt.Println("这是第一季度")
	//	case 2:
	//		fmt.Println("这是第二季度")
	//	case 3:
	//		fmt.Println("这是第三季度")
	//	case 4:
	//		fmt.Println("这是第四季度")
	//	default:
	//		fmt.Println("这是第10季度")
	//}
	//
	//fmt.Println("over.....")

	/*
		num1:=0
		num2:=0
		oper:=""

		fmt.Println("请输入num1:")
		fmt.Scanln(&num1)

		fmt.Println("请输入num2:")
		fmt.Scanln(&num2)

		fmt.Println("请输入oper + - * /:")
		fmt.Scanln(&oper)

		switch oper {
			case "+":
				fmt.Printf("num1 + num2 = %d\n",num1+num2)
			case "-":
				fmt.Printf("num1 - num2 = %d\n",num1-num2)
			case "*":
				fmt.Printf("num1 * num2 = %d\n",num1*num2)
			case "/":
				fmt.Printf("num1 / num2 = %d\n",num1/num2)
			default:
				fmt.Println("错误")
		}
	*/

	score := 88
	switch {
	case score >= 80 && score <= 90:
		fmt.Println("良好")
	case score >= 60 && score <= 70:
		fmt.Println("及格")
	case score >= 70 && score <= 80:
		fmt.Println("一般")
	case score >= 80 && score <= 90:
		fmt.Println("良好")
	}
	fmt.Println("=====================")

	letter:="A"

	switch letter{
	case "A","E","I","O","U":
		fmt.Println("元音")
	case "M","N":
		fmt.Println("m或n")

	}
}
