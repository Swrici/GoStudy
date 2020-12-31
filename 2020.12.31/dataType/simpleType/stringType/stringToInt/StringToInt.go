//string 转 int
package main

import (
	"fmt"
	"strconv"
)

func main()  {
	str := "10086"
	//转换过程返回形式是 int err
	fmt.Println(strconv.Atoi(str))

	fmt.Println(strconv.ParseInt(str,10,32))
}
