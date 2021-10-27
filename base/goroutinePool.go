package main
//线程池
import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type job struct {
	id int
	randnumber int
}

type result struct {
	job *job
	numbersum int
}

func main()  {
	chJob :=make(chan *job,128)
	chResult :=make(chan *result,128)
	createPool(5,chJob,chResult)
	go func(chResult chan *result) {
		for result:=range chResult {
			fmt.Printf("id:%d,randnum:%d,numbersum:%d \n",result.job.id,result.job.randnumber,result.numbersum)
		}
	}(chResult)
	var id int
	var rand_number int
	for i:=0;i<10;i++  {
		id++
		rand_number,_ = strconv.Atoi(fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000)))
		j:= &job{
			id:id,
			randnumber:rand_number,
		}
		chJob <- j
	}
	time.Sleep(time.Second*10)
}

func createPool(num int,chJob chan *job,chResult chan *result) {
	for i:=0;i<num;i++ {
		go func(chJob chan *job,chResult chan *result) {
			//遍历管道数据
			for job:=range chJob{
				r_num:=job.randnumber
				var sum int
				for r_num!=0 {
					temp:=r_num % 10
					sum+=temp
					r_num /=10
				}
				r := &result {
					job:job,
					numbersum: sum,
				}
				chResult <- r
			}

		}(chJob,chResult)
	}
}