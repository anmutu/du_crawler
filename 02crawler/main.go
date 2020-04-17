/*
  author='du'
  date='2020/4/14 13:43'
*/
package main

import (
	"du_crawler/02crawler/engine"
	"du_crawler/02crawler/parser"
)

func main() {
	e := engine.SimpleEngine{}
	e.Run(engine.Request{
		Url:        "https://www.cnblogs.com",
		ParserFunc: parser.ParseBlogList,
	})

}
