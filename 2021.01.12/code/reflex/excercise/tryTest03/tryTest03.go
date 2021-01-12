//通过反射获取方法的相关信息
package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type tv struct {
}

func (* tv)stringS(s string)  {
	fmt.Println(s)
}

func main() {
	Print(new(strings.Replacer))
	fmt.Println("==============================")
	Print(time.Hour)
}

func Print(x interface{}) {
	//获得x的值
	v := reflect.ValueOf(x)
	//拿到x的类型
	t := v.Type()
	//打印类型
	fmt.Printf("type %s\n", t)
	//遍历该类型的方法 v.NumMethod()
	for i := 0; i < v.NumMethod(); i++ {
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name,
			strings.TrimPrefix(methType.String(), "func"))
	}
}