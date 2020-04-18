/*
  author='du'
  date='2020/1/28 17:32'
*/
package scheduler

import (
	"du_crawler/03crawler/simple_con_engine"
)

//调度器里有一个worker的channel
type SimpleScheduler struct {
	workChan chan simple_con_engine.Request
}

//这里的workChan我们自己来配置，也就是将定义好的channel赋值给我们的workChan就完了。
func (s *SimpleScheduler) ConfigerMasterWorkerChan(c chan simple_con_engine.Request) {
	s.workChan = c
}

//就做把request送进workchan就这么一件事情。
//也就是scheduler里的channel里的数据送到worker里去。
//为每一个request建立一个gorutine.
func (s *SimpleScheduler) Submit(r simple_con_engine.Request) {
	go func() { s.workChan <- r }()
}
