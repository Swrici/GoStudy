// 打印XML文档中选定元素的文本。
//练习 7.17： 扩展xml select程序以便让元素不仅仅可以通过名称选择，也可以通过它们CSS样式上属性进行选择；例如一个像这样
package main

import (
"encoding/xml"
"fmt"
"io"
"os"
"strings"
)

func main() {
	//var b []byte
	//os.Stdin.Write(b)
	//for _,item :=range b {
	//	if item=='&' || item=='?'{
	//		item = ' '
	//	}
	//}
	//writer := io.Writer(writb)
	dec := xml.NewDecoder(os.Stdin)
	var stack []string // 名为stack的string切片
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
			case xml.StartElement:
				stack = append(stack, tok.Name.Local) // push

			case xml.EndElement:
				stack = stack[:len(stack)-1] // pop
			case xml.CharData:
				if containsAll(stack, os.Args[1:]) {
					fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
				}
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			if tok.Name.Local == "id"  {
				stack = append(stack, tok.Name.Local) // push
			}

		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
