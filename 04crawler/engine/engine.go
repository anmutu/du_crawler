/*
  author='du'
  date='2020/4/18 21:45'
*/

package con_engine

import (
	"log"
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
		//每一个channel of worker(chan Request)都会对应一个Worker去做事情
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
			log.Printf("这是取到的第#%d条数据。其对应的item值为:%s", itemCount, item)
			itemCount++
		}

		//把request放到scheduler维护的request队列里
		for _, item := range result.Requests {
			e.Scheduler.Submit(item)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			//告诉scheduler我准备好了。把chan Request放进scheduler维护的worker队列里就说明准备好了。
			ready.WorkerReady(in) //这里就是把worker的channel放到了scheduler维护的worker队列里。
			request := <-in
			res, err := Worker(request) //在这里把放到workChan里的request给到worker.
			if err != nil {
				continue
			}
			out <- res
		}
	}()
}
