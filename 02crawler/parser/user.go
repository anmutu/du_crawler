/*
  author='du'
  date='2020/4/14 14:35'
*/
package parser

import (
	"du_crawler/02crawler/engine"
	"regexp"
)

//<a href="https://home.cnblogs.com/u/NanoDragon/">柠檬橙1024</a>
//const userRe = `(<a href="https://home.cnblogs.com/u/[a-zA-Z0-9]+/)">([^<]*)</a>`

//<a href="https://www.cnblogs.com/lbhym/">耶low</a>
const userRe = `(<a href="https://www.cnblogs.com/([a-zA-Z0-9]+)/">([^<]*)</a>)`

//传入contents，输出是一个request的一个item的集合
func ParseUser(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(userRe)
	result := engine.ParseResult{}
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Items = append(result.Items, string(m[1]))
		url := "https://home.cnblogs.com/u/" + string(m[2])
		result.Requests = append(result.Requests, engine.Request{Url: url, ParserFunc: engine.NilParser})
	}
	return result
}
