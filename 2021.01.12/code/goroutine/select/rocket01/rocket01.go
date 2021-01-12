//暖暖手
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("倒计时：")
	tick := time.Tick(1 * time.Second)
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()

}

func launch() {
	fmt.Println("Ohhhhhh火箭发射了！！！")
}
