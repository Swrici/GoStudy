//练习 7.4：
//strings.NewReader函数通过读取一个string参数返回一个满足io.Reader接口类型的值（和其它值）。
//实现一个简单版本的NewReader，并用它来构造一个接收字符串输入的HTML解析器（§5.2）
package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
)
type par interface {
	 NewReader() io.Reader
}
type str string

func (s *str)NewReader(str string)  {
	//
	//将str 转成一个io.Reader赋值给read
	//
	doc, err := html.Parse(read)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	sliceMap := make(map[string]int)
	outline(sliceMap, doc)
	for tag,count := range sliceMap{
		fmt.Printf("%s元素的次数为:%d次\n",tag,count)
	}
}

func main() {
	fmt.Println(1)
	var s str = "https://baidu.com"
	s.NewReader(s)
}

//输出结果  fetch https://golang.org | recursionExcercise5.2
func outline(stack map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		stack[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

