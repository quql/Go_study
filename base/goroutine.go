package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
func hello(i int) {
	time.Sleep(time.Second)
	fmt.Println("hello Goroutine",i,time.Now().UnixNano())
}

func main () {
	//go hello()
	//fmt.Println("main goroutine done!")
	//time.Sleep(time.Second)
	for i:=0;i<10000;i++ {
		//wg.Add(1)//启动一个goroutine就登记+1
		go hello(i)
	}
	time.Sleep(time.Second*10)
    //wg.Wait()
}
