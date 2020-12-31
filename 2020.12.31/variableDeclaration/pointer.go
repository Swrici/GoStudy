// Echo4 prints its command-line arguments.
package main

import (
"flag"
"fmt"
"strings"
)

var n = flag.Bool("n", false, "omit trailing newline 忽略尾部换行符")
var sep = flag.String("s", " ", "separator 分离器")

func main() {
	flag.Parse() //这个函数的主要是把用户传递的命令行参数解析为对应变量的值
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}