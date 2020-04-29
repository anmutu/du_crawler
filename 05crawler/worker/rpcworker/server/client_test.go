/*
  author='du'
  date='2020/4/27 16:19'
*/
package main

import (
	"du_crawler/05crawler/config"
	"du_crawler/05crawler/rpchelper"
	worker "du_crawler/05crawler/worker/rpcworker"
	"fmt"
	"testing"
	"time"
)

func TestCrawlerService(t *testing.T) {
	//创建client端
	const host = "9000"
	go rpchelper.ServeRpc(host, worker.CrawlerService{})
	time.Sleep(time.Second)
	client, err := rpchelper.NewClient(host)
	if err != nil {
		panic(err)
	}

	//参数组装
	req := worker.Request{
		Url: "https://www.cnblogs.com",
		Parser: worker.SerializedParser{
			Name: config.ParseBlogList,
			Args: nil,
		},
	}

	//调用rpc服务
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}

}
