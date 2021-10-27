package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
    go counter(ch1)
	go squarer(ch2,ch1)
	printer(ch2)
}

//单向传递，只能传值不能取值
func counter(out chan <- int) {
	for i:=0;i<10 ;i++  {
		out<-i
	}
	close(out)
}
//out只能传值不能取值，in只能取值不能传值
func squarer(out chan <- int,in <-chan int) {
	for i:= range in{
		out<- i*i
	}
	close(out)
}

func printer(in <-chan int)  {
	for i:=range in {
		fmt.Println(i)
	}
}