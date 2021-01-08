//它在主goroutine中（译注：就是执行main函数的goroutine）将标准输入复制到server，
//因此当客户端程序关闭标准输入时，后台goroutine可能依然在工作。
//我们需要让主goroutine等待后台goroutine完成工作后再退出，我们使用了一个channel来同步两个goroutine：
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	//发起一个链接
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	//并发进行
	go func() {
		//输出conn接受到的数据
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		//被阻塞
		log.Println("done")
		fmt.Println("goroutine中的值")
		fmt.Println("goroutine中的值")
		fmt.Println("goroutine中的值")
		fmt.Println("goroutine中的值")
		fmt.Println("goroutine中的值")
		done <- struct{}{} // 作为信号当顺利传输接受后即可不足阻塞
	}()
	mustCopy(conn, os.Stdin)
	fmt.Println("main 中执行快")
	conn.Close()
	<-done // wait for background goroutine to finish
	//等待goroutine执行完再执行~~
	//被阻塞了
	fmt.Println("等到chan 接受值 没有阻塞再执行")
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
