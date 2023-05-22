/*
  author='du'
  date='2020/4/18 21:45'
*/

package scheduler

import (
	con_engine "du_crawler/04crawler/engine"
)

// QueuedScheduler 队列调度器 里有request相关的channel和worker相关的channel
type QueuedScheduler struct {
	requestChan chan con_engine.Request
	workerChan  chan chan con_engine.Request
}

// Submit 把request放到我们维护的request的队列里。
func (s *QueuedScheduler) Submit(r con_engine.Request) {
	s.requestChan <- r
}

// WorkerReady 把request放到我们维护的worker的队列里。
// engine里维护的worker队列里有数据了，可以交给Worker函数去fetch和parser了。就这么个意思。
func (s *QueuedScheduler) WorkerReady(w chan con_engine.Request) {
	s.workerChan <- w
}

// WorkerChan 创建worker的channel
func (s *QueuedScheduler) WorkerChan() chan con_engine.Request {
	return make(chan con_engine.Request)
}

// Run 这里是Scheduler的Run函数。
// scheduler就是维护了两个队列。一旦两个队列里都有值就尝试把request发给worker。
func (s *QueuedScheduler) Run() {
	s.requestChan = make(chan con_engine.Request) //指针接收者才能改变里面的内容。
	s.workerChan = make(chan chan con_engine.Request)
	go func() {
		var requestQ []con_engine.Request
		var workerQ []chan con_engine.Request
		for {
			var activeRequest con_engine.Request
			var activeWorker chan con_engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			//收到一个request就让request排队，收到一个worker就让worker排队。所有的channel操作都放到select里。
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
