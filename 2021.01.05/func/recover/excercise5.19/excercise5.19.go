//使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。
package main

import (
	"fmt"
	"runtime"
)

func main()  {
	rec(1)


}

type bailout struct {

}

func rec(val int)  {
	type errorType struct {
		msg string
	}
	defer func() {
		//recover没有return但是有一个非零值的返回值
		 err := recover()
		switch err.(type)  {
			case runtime.Error:
				fmt.Println("runtime error:", err)		// no panic
			default: // "expected" panic
				fmt.Println(fmt.Errorf("error：%q", err))
		}
	}()
	panic(errorType{msg:"报错"})
}

