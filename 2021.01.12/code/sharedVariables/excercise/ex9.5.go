//练习 9.5: 写一个有两个goroutine的程序，两个goroutine会向两个无buffer channel反复地发送ping-pong消息。这样的程序每秒可以支持多少次通信？
package main

import (
	"fmt"
	"time"
)

func main() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	var i = 0
	var j = 0

	//ch2 <- j
	go func() {
		for  {
				i++
				fmt.Println(<-ch1)
		}
	}()
	go func() {
		for  {
			ch1 <- (i)
		}
	}()

	go func() {
		for  {
			j++
			fmt.Println(<-ch2)
		}
	}()
	go func() {
		for  {
			ch2 <- (j)
		}
	}()
	time.Sleep(time.Duration(1)*time.Second)
	fmt.Println("i:",i," j:",j)
}
