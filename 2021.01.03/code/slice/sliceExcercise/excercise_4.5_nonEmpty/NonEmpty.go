package main

import "fmt"

func main() {
	s := []string{"abc", "abc", "", "123", "12", "12", "123"}
	fmt.Println(nonEmpty(s))
}

func nonEmpty(s []string) []string {
	var sameStr = ""
	fmt.Println(sameStr)
	//进行判断
	for i,str := range  s {
		if i!=len(s)-1&&str==sameStr {
			sameStr = str
			fmt.Print(i)
			fmt.Println(str)
			s[i] = ""
			continue
		}else {
			sameStr = str
		}
	}
	//查看循环后遗漏的最后两个数有无重复
	if len(s)>0 {
		if s[len(s)-1]==s[len(s)-2] {
			s[(len(s)-2)]=""
		}
	}
	fmt.Println(s)
	return s
}
