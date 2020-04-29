/*
  author='du'
  date='2020/1/25 14:30'
*/
package parser

import (
	dis_engine "du_crawler/05crawler/engine"
	"regexp"
	"strconv"
	"strings"
)

const blogListRe = `<h3><a class="titlelnk" href="(https://www.cnblogs.com/[a-zA-Z0-9]+/p/[0-9]+.html)" target="_blank">([^<]+)</a></h3>`

var index = 0

//传入contents，输出是一个request的一个item的集合
func ParseBlogList(contents []byte, _ string) dis_engine.ParseResult {
	re := regexp.MustCompile(blogListRe)
	result := dis_engine.ParseResult{}
	matches := re.FindAllSubmatch(contents, -1)
	pre := "https://www.cnblogs.com/sitehome/p/"
	for _, m := range matches {
		index++
		var build strings.Builder
		build.WriteString(pre)
		build.WriteString(strconv.Itoa(index))
		url := build.String()
		result.Requests = append(result.Requests, dis_engine.Request{Url: url, Parser: dis_engine.NewFuncParser(
			ParseBlogList, "ParseBlogList")})
		result.Items = append(result.Items, string(m[1])+string(m[2]))
	}
	return result
}
