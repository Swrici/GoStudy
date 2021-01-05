// 编写函数，记录在HTML树中出现的同名元素的次数。
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	//解析输入html,返回解析树
	doc, err := html.Parse(os.Stdin)
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
//输出结果  fetch https://golang.org | recursionExcercise5.2
func outline(stack map[string]int, n *html.Node) {

	if n.Type == html.ElementNode {
		stack[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}

}
