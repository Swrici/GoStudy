package main

import (
	"fmt"
	"reflect"
)

func main() {
	//一个 Type 表示一个Go类型. 它是一个接口,
	t1 := reflect.TypeOf(999)
	fmt.Println("t1:",t1,",t1.String():",t1.String())

}
