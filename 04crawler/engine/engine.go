/*
  author='du'
  date='2020/4/18 21:45'
*/
package con_engine

import (
	"fmt"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run() //这里就维护了两个队列。

	for i := 0; i < e.WorkerCount; i++ {
		//每一个channel of request都会对应一个worker
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		//把request送进我们scheduler维护的request的队列中
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <-out

		//将拿到的值的数量和其值打印出来
		for _, item := range result.Items {
			fmt.Printf("这是拿到的第#%d个数据。拿到item的值是%s:\n", itemCount, item)
			fmt.Printf("这是拿到的第#%d个数据。拿到item的值是%s:\n", itemCount, item)
			itemCount++
		}

		//将拿到的seed给到scheduler去提交
		for _, item := range result.Requests {
			e.Scheduler.Submit(item)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			//告诉scheduler我准备好了。
			ready.WorkerReady(in) //这里就是把worker的channel放到了scheduler维护的worker队列里。
			request := <-in
			res, err := Worker(request) //在这里把request给到worker.
			if err != nil {
				continue
			}
			out <- res
		}
	}()
}
