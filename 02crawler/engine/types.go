/*
author='du'
date='2020/1/23 20:49'
*/

package engine

// Request 请求的结构体，包含一个url和解析这个url的一个的函数
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult //解析的函数,传入content，返回ParseResult的结构体
}

// ParseResult 返回的结构体，其中"interface{}"表示任何表示任何类型，有点类似c#里的泛型T
type ParseResult struct {
	Requests []Request     //返回的Request切片可以继续用于解析
	Items    []interface{} //这里的就可以打印或者存起来了
}

// NilParser 为了让blog list里的编译通过
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
