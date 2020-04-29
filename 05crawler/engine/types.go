/*
  author='du'
  date='2020/1/23 20:49'
*/
package dis_engine

//请求的结构体
type Request struct {
	Url    string
	Parser Parser //这里不在跟并发版本的一样了。这里不在是一个函数了。
}

//Parser的接口
type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

//一个函数类型的parser和其对interface接口的实现
type FuncParser struct {
	parser ParseFunc
	name   string
}

//parse函数的结构体
type ParseFunc func(contents []byte, url string) ParseResult

//实现接口
func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.Parse(contents, url)
}

//实现接口，返回name和参数。r.Parser.Serialize()
func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParseFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}

//返回的结构体，其中"interface{}"表示任何表示任何类型，有点类似c#里的泛型T
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}
