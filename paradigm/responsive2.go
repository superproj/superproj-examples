package main

import (
	"fmt"
	"time"

	"github.com/reactivex/rxgo/v2"
)

type ScheduledTask struct {
	RecordId        int
	HandleStartTime string
	Status          bool
}

func main() {
	ch := make(chan rxgo.Item)
	go producer(ch)

	time.Sleep(time.Second * 3)
	observable := rxgo.FromChannel(ch)
	observable = observable.Filter(func(i interface{}) bool {
		st := i.(*ScheduledTask)
		return st.Status
	}, rxgo.WithBufferedChannel(1))

	// 消费可观测量
	for customer := range observable.Observe() {
		st := customer.V.(*ScheduledTask)
		fmt.Printf("result: %+v\n", st)
	}
}

func producer(ch chan<- rxgo.Item) {
	for i := 0; i < 10; i++ {
		status := false
		if i%2 == 0 {
			status = true
		}
		st := &ScheduledTask{
			RecordId:        i,
			HandleStartTime: time.Now().Format("2006-01-02 13:04:05"),
			Status:          status,
		}
		ch <- rxgo.Of(st)
	}

	// 这里千万不要忘记了
	close(ch)
}
