package main

import "fmt"

func main() {
	for i:=1;i<10 ;i++  {
		for j:=i; j>=1 ; j-- {
			fmt.Printf("%dX%d=%d\t",i,j,i*j)
		}
		fmt.Println()
	}
}