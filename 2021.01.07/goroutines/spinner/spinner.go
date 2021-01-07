package main

import (
	"fmt"
	"time"
)

func main() {
	//开启函数spinner打印字符，避免无聊充斥等待时间
	go spinner(100 * time.Millisecond)
	const n = 45
	//进行斐波那契数列的计数，输入需要计算到的位置
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		//循环打印，知道主线程结束，goroutines自动退群
		for _, r := range `-\|/` {
			fmt.Printf("%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	//前两位返回本身
	if x < 2 {
		return x
	}
	//计算思路
	return fib(x-1) + fib(x-2)
}
