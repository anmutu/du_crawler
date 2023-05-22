/*
author='du'
date='2020/4/18 21:34'
*/
package main

import (
	con_engine "du_crawler/04crawler/engine"
	"du_crawler/04crawler/scheduler"
	"du_crawler/04crawler/worker/parser"
)

func main() {
	e := con_engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 50,
	}
	e.Run(con_engine.Request{
		Url:        "https://www.cnblogs.com",
		ParserFunc: parser.ParseBlogList,
	})
}
