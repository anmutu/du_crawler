/*
  author='du'
  date='2020/4/17 8:09'
*/
package main

import (
	"du_crawler/03crawler/parser"
	"du_crawler/03crawler/scheduler"
	"du_crawler/03crawler/simple_con_engine"
)

func main() {
	e := simple_con_engine.SimpleConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	e.Run(simple_con_engine.Request{
		Url:        "https://www.cnblogs.com",
		ParserFunc: parser.ParseBlogList,
	})
}
