//字符串的常用方法
package main

import (
	"fmt"
	"strings"
)

func main()  {
	s := "hello,Golang~"
	//输出s的长度
	fmt.Println(len(s))
	fmt.Println("%q",s[0:2])
	//fmt.Println(s[len(s)]) //超出长度，弹出panic异常，数组越界
	//== >
	str1,str2 := "hello2","hello2" //"hello2","hello2" == "hello1",
	// "hello1","hello2" <
	// "hello3","hello2" >
	// "hello11","hello2" <
	//比较是通过 逐个字节比较 完成的
	if str1 == str2 {
		fmt.Println("str1==str2")
	}else if str1 > str2 {
		fmt.Println("str1 > str2")
	}else if str1 < str2 {
		fmt.Println("str1 < str2")
	}
}

//strings包下的HasPrefix和HasSuffix函数分别判断有无相关的前后缀
func hasPrefix(string2 string,prefix string) bool {
	return strings.HasPrefix(string2,prefix)
}

func hasSuffix(string2 string,suffix string) bool {
	return strings.HasSuffix(string2,suffix)
}

