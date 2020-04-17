/*
  author='du'
  date='2020/1/25 14:30'
*/
package parser

import (
	"du_crawler/02crawler/engine"
	"regexp"
)

const blogListRe = `<h3><a class="titlelnk" href="(https://www.cnblogs.com/[a-zA-Z0-9]+/p/[0-9]+.html)" target="_blank">([^<]+)</a></h3>`

//传入contents，输出是一个request的一个item的集合
func ParseBlogList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(blogListRe)
	result := engine.ParseResult{}
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Items = append(result.Items, string(m[1])+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]), ParserFunc: ParseUser})
	}
	return result
}
