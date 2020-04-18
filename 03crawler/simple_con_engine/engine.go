/*
  author='du'
  date='2020/1/28 16:37'
*/
package simple_con_engine

import (
	"fmt"
)

//简单调度器
type SimpleConcurrentEngine struct {
	Scheduler   Scheduler //调度器
	WorkerCount int       //worker的数量
}

//Scheduler
type Scheduler interface {
	Submit(Request)                        //用于将Request送给worker的channel
	ConfigerMasterWorkerChan(chan Request) //配置workChan,实现也就是个赋值操作。
}

//这里是所有的worker共用一个输入，但是会有多个去处理送过来的输入。
func (e *SimpleConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)

	//把in给到workerChan
	e.Scheduler.ConfigerMasterWorkerChan(in)
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	//现在是由scheduler去把request提交
	for _, r := range seeds {
		//把request送进workerChan
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <-out

		//将拿到的值的数量和其值打印出来
		for _, item := range result.Items {
			fmt.Printf("这是拿到的第#%d个数据。拿到item的值是%s:\n", itemCount, item)
			itemCount++
		}

		//将拿到的seed给到scheduler去提交
		for _, item := range result.Requests {
			e.Scheduler.Submit(item)
		}
	}
}

//创建worker，传入Request的channel
func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			res, err := Worker(request)
			if err != nil {
				continue
			}
			out <- res
		}
	}()
}
