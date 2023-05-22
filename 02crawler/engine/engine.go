/*
author='du'
date='2020/1/24 11:16'
*/

package engine

import (
	"du_crawler/02crawler/fetcher"
	"log"
)

type SimpleEngine struct{}

// Run 传入Request的种子。
func (e SimpleEngine) Run(seeds ...Request) {

	//把seeds里的放到requests里。
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	//只要种子里有就一直去
	itemCount := 0
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		//这里就是每次来都会有一个worker去做事情。
		parseResult, err := e.worker(r)
		if err != nil {
			log.Printf("遇到error:%s,此时parseResult的request是%v,items是%v",
				err, parseResult.Requests, parseResult.Items)
			continue
		}
		requests = append(requests, parseResult.Requests...) //注意这三个.的语法。

		for _, item := range parseResult.Items {
			log.Printf("这是取到的第#%d条数据。其对应的item值为:%s", itemCount, item)
			itemCount++
		}
	}
}

// 传入Request结构体，返回ParseResult。
func (SimpleEngine) worker(r Request) (ParseResult, error) {
	log.Printf("fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		//如果requests里有就一直请求，如果其中有错，注意不要panic,要纪录日志。
		log.Printf("Fetcher失败，fetch的url是%s:,错误信息是:%v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil //这里的返回就是parser的一个过程了。
}
