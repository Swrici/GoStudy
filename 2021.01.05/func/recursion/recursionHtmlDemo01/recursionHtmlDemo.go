//Findlinks1打印从标准输入读取的HTML文档中的链接。
//Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	//Parse返回给定读取器中HTML的解析树,以及一个错误信息。
	var (
		doc, err = html.Parse(os.Stdin)
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	//无报错，遍历链接
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result.
//visit将在n中找到的每个链接附加到链接并返回结果。
func visit(links []string, n *html.Node) []string {
	//在html结点里找到a标签
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				//获取a标签里的链接放入links中
				links = append(links, a.Val)
			}
		}
	}
	//对每个兄弟节点进行遍历，循环调用
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
