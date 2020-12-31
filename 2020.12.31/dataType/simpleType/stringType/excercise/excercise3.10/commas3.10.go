//练习 3.10： 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作。
package main

import (
	"bytes"
	"fmt"
)

func main()  {
	fmt.Println(commas("helloGolang"))
}
//从前面字符往后数隔3个数加个逗号
func commas(s string) string {
	var buf bytes.Buffer
	b := []byte(s)
	//j负责协助计数
	j := 0
	//遍历一遍，
	//fmt.Println(len(b))
	for i:=0;i+j<len(b);i++ {
		//i+1是因为i从0开始，
		if (i)%3==0&&i!=0 {
			buf.WriteByte(',')
			buf.WriteByte(b[i+j])
			j = j+i
			i=0
			continue
		}
		buf.WriteByte(b[i+j])
	}
	return buf.String()
}
