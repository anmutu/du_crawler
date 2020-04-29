/*
  author='du'
  date='2020/4/26 19:00'
*/
package worker

import (
	"du_crawler/05crawler/config"
	dis_engine "du_crawler/05crawler/engine"
	"du_crawler/05crawler/worker/parser"
	"errors"
	"log"
)

//request的结构体
type Request struct {
	Url    string
	Parser SerializedParser
}

//request结构体里SerializedParser的结构体
type SerializedParser struct {
	Name string
	Args interface{}
}

//返回的ParseResult结构体
type ParseResult struct {
	Items    interface{}
	Requests []Request
}

//序列化request。把dis_engine里的Request转换成这里的Request。
func SerializeRequest(r dis_engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

//序列化返回的result。把dis_engine里的ParseResult转换成这里的ParseResult。
func SerializeResult(r dis_engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

//反序列化request
func DeserializeRequest(r Request) (dis_engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return dis_engine.Request{}, err
	}
	return dis_engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

//反序列化result
func DeserializeResult(r ParseResult) dis_engine.ParseResult {
	result := dis_engine.ParseResult{
		//Items: r.Items,
	}

	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("反序列化request失败: %v", err)
			continue
		}
		result.Requests = append(result.Requests,
			engineReq)
	}
	return result
}

func deserializeParser(p SerializedParser) (dis_engine.Parser, error) {
	switch p.Name {
	case config.ParseBlogList:
		return dis_engine.NewFuncParser(
			parser.ParseBlogList,
			config.ParseBlogList), nil
	default:
		return nil, errors.New(
			"小杜同学表示这个parser")
	}
}
