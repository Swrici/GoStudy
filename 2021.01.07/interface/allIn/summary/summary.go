package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main()  {
	//7.5接口值
	//接口值由动态类型和动态值组成
	//io.Writer是一个具体的接口
	var w io.Writer
	//赋予具体的类型
	w = new(bytes.Buffer)
	//调用Writer接口中Write方法在w写入内容
	w.Write([]byte("hello"))
	fmt.Println(w)
	w = os.Stdout
	//这里调用了
	w.Write([] byte("helloWorld"))
	w= nil
	//会报错，空接口值中的所有方法调用都会导致报错
	//w.Write([] byte("helloWorld"))
	//7.10 接口断言 x.(T)
	var wr io.Writer
	if wr,ok := wr.(*os.File);ok{
		//根据T(具体类型，指针类型)进行判断,
		//具体类型，前者和后者的类型进行判段,失败报panic，成功无事发生
		//指针类型,前者和后者的类型进行判段，成功有了后者的类型，大于等于原类型
		fmt.Println("具体类型判断值%T",wr)
	}


}
