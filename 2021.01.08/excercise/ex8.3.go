//练习 8.3： 在netcat3例子中，conn虽然是一个interface类型的值，
//但是其底层真实类型是*net.TCPConn，代表一个TCP连接。
//一个TCP连接有读和写两个部分，可以使用CloseRead和CloseWrite方法分别关闭它们。
//修改netcat3的主goroutine代码，只关闭网络连接中写的部分，
//这样的话后台goroutine可以在标准输入被关闭后继续打印从reverb1服务器传回的数据。
//（要在reverb2服务器也完成同样的功能是比较困难的；参考练习 8.4。）
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
		//goroutine可以在标准输入被关闭后继续打印从reverb1服务器传回的数据。
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors

		//被阻塞
		log.Println("done")
		fmt.Println("goroutine中的值")
		fmt.Println("goroutine中的值")
		fmt.Println("goroutine中的值")
		done <- struct{}{} // 作为信号当顺利传输接受后即可不足阻塞
	}()
	mustCopy(conn, os.Stdin)
	fmt.Println("main 中执行快")

	conn.Close()
	<-done // 作为信号当顺利传输接受后即可不足阻塞
	//等待goroutine执行完再执行~~
	//下面语句被阻塞了
	fmt.Println("等到chan 接受值 没有阻塞再执行")
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
