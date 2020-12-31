package main

import (
	"fmt"
	"os"
	"testing"
)

func BenchWhoIsQuick(b *testing.B)  {
	var s,step string
	for i:=1;i<len(os.Args);i++{
		s += step + os.Args[i]
		step = " "
	}
}
func main()  {
	//echo1.1
	fmt.Println("即将输出os.Args[0]：")
	fmt.Println(os.Args[0])
	//echo1.2
	fmt.Println("即将依行输出：")
	for _,args:=range os.Args{
		fmt.Println(1,args)
		fmt.Println("我是换行符")
	}

}
