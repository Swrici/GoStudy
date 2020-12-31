//常数 iota练习
package main

import "fmt"

//
const (
	_ = 1 << (10 *iota) //左移10位
	kib //左移 20 位
	mib	//以此类推
	gib
	tib
	pib
	eib
	zib
	yib
)

func main()  {
	fmt.Println(yib) //打印不出来，越界了
}

