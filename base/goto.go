package main

import "fmt"

func main() {
	/**
	goto语句
	 */
	a:=10;
	LOOP:
		for a<20  {
			if a==15{
				a+=1
				goto LOOP
			}
			fmt.Println(a)
			a++
		}
}