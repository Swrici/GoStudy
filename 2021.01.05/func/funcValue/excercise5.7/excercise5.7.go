//练习 5.7： 完善startElement和endElement函数，使其成为通用的HTML输出器。
//要求：输出注释结点，文本结点以及每个元素的属性（< a href='...'>）。
//编写测试，验证程序输出的格式正确。（详见11章）
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

var depth int

func main() {
	//解析输入html,返回解析树
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc,startElement,endElement)
}
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	//if n.Type == html.TextNode || n.Type == html.CommentNode{
	//	//stack = append(stack, n.Data) // push tag
	//
	//}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
func startElement(n *html.Node) {
	//html.TextNode文本节点    html.CommentNode 注释节点
	if n.Type == html.TextNode  {
		fmt.Printf("%*s<%s %s>\n", depth*2, "", n.Data,n.Attr)
		depth++
	}
	if n.Type == html.CommentNode  {
	fmt.Printf("%*s<--！%s %s>\n", depth*2, "", n.Data,n.Attr)
	depth++
	}
}
func endElement(n *html.Node) {
	if n.Type == html.TextNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
	if n.Type == html.CommentNode {
		depth--
		fmt.Printf("%*s</%s -->\n", depth*2, "", n.Data)
	}
}