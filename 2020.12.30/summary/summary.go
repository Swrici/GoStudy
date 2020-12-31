//流程控制:switch
package main

import (
	"fmt"
)

func main()  {

	var x int
	fmt.Scanln(&x)
	switchChoose(x)

}

func switchChoose(x int)  {
	switch  {
	case x <60:
		fmt.Println("不及格了啊")
		return
	case x<80:
		fmt.Println("可以可以，下次要争取优秀哦")
	case x<100:
		fmt.Println("优秀的人还要更进一步才行~")
	case x==100:
		fmt.Println("哇，你不知道满分是正无穷吗~")
	case x>100:
		fmt.Println("可以可以，你又要开始学习了。")
	}
}
