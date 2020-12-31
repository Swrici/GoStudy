package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main()  {
	counts := make(map[string]int)
	for _,filename := range os.Args[1:] {
		//ReadFile函数（来自于io/ioutil包），其读取指定文件的全部内容
		data,err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr,"dup3:%v\n",err)
			continue
		}
		//ReadFile函数返回一个字节切片（byte slice），必须把它转换为string，才能用strings.Split分割。
		for _,line := range strings.Split(string(data),"\n") {
			counts[line]++
		}
		for lines,n := range counts {
			if n>=0 {
				fmt.Printf("%d\t%s\n", n, lines)
			}
		}
	}
}
