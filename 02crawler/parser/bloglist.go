/*
  author='du'
  date='2020/1/25 14:30'
*/

package parser

import (
	"du_crawler/02crawler/engine"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// const blogListRe = `<h3><a class="titlelnk" href="(https://www.cnblogs.com/[a-zA-Z0-9]+/p/[0-9]+.html)" target="_blank">([^<]+)</a></h3>`

const blogListRe = `<a class="post-item-title" href="(https://www.cnblogs.com/[a-zA-Z0-9]+/p/[0-9]+.html)" target="_blank">([^<]+)</a>`

var index = 0

// ParseBlogList 传入contents，输出是一个request的一个item的集合
func ParseBlogList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(blogListRe)
	result := engine.ParseResult{}
	matches := re.FindAllSubmatch(contents, -1)
	pre := "https://www.cnblogs.com/sitehome/p/"
	for _, m := range matches {
		index++
		var build strings.Builder
		build.WriteString(pre)
		build.WriteString(strconv.Itoa(index))
		url := build.String()
		fmt.Println("url=>", url)
		result.Requests = append(result.Requests, engine.Request{Url: url, ParserFunc: ParseBlogList})
		result.Items = append(result.Items, string(m[1])+string(m[2]))
	}
	return result
}
