package main

import (
	"fmt"
	"os"
	"time"
)

func test1(ch1 chan string) {
	time.Sleep(time.Second)
	ch1 <- "test1"
}

func test2(ch2 chan string) {
	time.Sleep(time.Second)
	ch2 <- "test2"
}

func main() {
	//ch1 :=make(chan string)
	//ch2 :=make(chan string)
	//go test1(ch1)
	//go test2(ch2)
	//select {
	//case s1:=<-ch1:
	//	fmt.Println("ch1=",s1)
	//case s2:=<-ch2:
	//	fmt.Println("ch2=",s2)
	//}
	//for {
	//
	//}
	// 创建管道
	output1 := make(chan string, 10)
	// 子协程写数据
	go write(output1)
	// 取数据
	for s := range output1 {
		fmt.Println("res:", s)
		time.Sleep(time.Second)
	}
}

func write(ch chan string) {
	for {
		select {
		// 写数据
		case ch <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel full")
			os.Exit(1)
		}
		time.Sleep(time.Millisecond * 500)
	}
}
