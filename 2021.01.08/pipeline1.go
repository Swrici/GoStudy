//对chan进行编写测试，了解
package main

import "fmt"

func main()  {
	//1. 无缓存的chan
	//chan赋值，要通过make
	num1 := make(chan int)
	num2 := make(chan int)
	//两个goroutine是用num1进行沟通的
	go func() {
		num1 <- 9
	}()

	go func() {
		x := <- num1
		x++
		num2 <- x
	}()

	//main goroutine 进行编写
	fmt.Println(<-num2)
	//当这里chan被关闭的时候，会返回一个零值
	close(num1)
	fmt.Println(<-num1)

	//为了避免被关闭的chan被使用可以
	if y,ok := <- num1;ok{
		fmt.Println(y)
	}else {
		fmt.Println("chan已被关闭")
	}
	//2. 有缓存的chan
	num3 := make(chan int,3)
	//输入chan
	num3 <- 1
	num3 <- 13
	num3 <- 1314
	//获取chan的容量
	fmt.Println(cap(num3))
	//输出chan
	fmt.Println(<- num3)
	fmt.Println(<- num3)
	//获取chan的有效元素长度
	fmt.Println(len(num3))

}
