package main

import "fmt"

func main() {
	//ch := make(chan int)
	//go recv(ch)
	//ch<-10
	//fmt.Println("发送成功")
	//close(ch)
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i:=0;i<100;i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for  {
			i,ok := <-ch1
			if (!ok) {
				break
			}
			ch2 <- i*i
		}
		close(ch2)
	}()
	for j:=range ch2 {
		println(j)
	}
}

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}