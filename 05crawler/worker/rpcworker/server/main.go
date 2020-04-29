/*
  author='du'
  date='2020/4/27 15:40'
*/
package main

import (
	"du_crawler/05crawler/config"
	"du_crawler/05crawler/rpchelper"
	worker "du_crawler/05crawler/worker/rpcworker"
)

func main() {
	rpchelper.ServeRpc(string(config.WorkerPort0), worker.CrawlerService{})
}
