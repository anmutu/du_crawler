/*
  author='du'
  date='2020/1/28 13:42'
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

/*
 这里就是最简单的拉取网页，然后按照一定规则匹配，然后将结果打印出来。
*/

func main() {
	blogTest()
}

func blogTest() {
	resp, err := http.Get("https://www.cnblogs.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("错误，状态号", resp.StatusCode)
		return
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	printBlogList(contents)
}

//传入拉取下来的网页内容，进行正则匹配，并打印结果。
func printBlogList(contents []byte) {
	//<h3><a class="titlelnk" href="https://www.cnblogs.com/ITnoteforlsy/p/12228149.html" target="_blank">B-Tree 和 B+Tree 结构及应用，InnoDB 引擎， MyISAM 引擎</a></h3>
	//用"[a-zA-Z0-9]","[0-9]","[^<]"匹配。
	re := regexp.MustCompile(`<h3><a class="titlelnk" href="https://www.cnblogs.com/[a-zA-Z0-9]+/p/[0-9]+.html" target="_blank">[^<]+</a></h3>`)
	mathes := re.FindAll(contents, -1)
	for _, m := range mathes {
		fmt.Printf("%s\n", m)
	}
	fmt.Println(len(mathes))
}

