//int 转 String
package main

import (
	"fmt"
	"strconv"
)

func main()  {
	//整数转字符串的两种方法
	x := 10086
	y := fmt.Sprintf("%d",x)
	fmt.Println(y,strconv.Itoa(x)) //10086 10086

	//转换数字进制 strconv.FormatInt 和 FormatUint
	fmt.Println(strconv.FormatInt(int64(x),2))
	fmt.Println(strconv.FormatInt(int64(x),8))


}
