package main

import (
	"fmt"
	"reflect"
)

func main() {
	//尝试指针更新
	 x := 66
	 //获取任意变量x对应的可取地址的Value。
	 y := reflect.ValueOf(&x).Elem()
	 //通过调用Interface()方法，也就是返回一个interface{}，里面包含指向变量的指针，再通过断言将得到的interface{}类型的接口强制转为普通的类型指针
	 z := y.Addr().Interface().(*int)
	 *z = 99
	 fmt.Println(*z)

	 //进行调用可取地址的reflect.Value的reflect.Value.Set方法
	y.Set(reflect.ValueOf(88))
	 fmt.Println(x)


}
