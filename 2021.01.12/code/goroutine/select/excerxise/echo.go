//练习 8.8： 使用select来改造8.3节中的echo服务器，为其增加超时，
//这样服务器可以在客户端10秒中没有任何喊话时自动断开连接。
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	//调用 net.Listen方法监听8000端口
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		//当监听器接受到链接时，并发运行handleConn函数
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}

//echo函数将输入的string按大写、原型、小写由相同间隔输出
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//该函数的作用是向conn输入信息，进行调用echo方法
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		// ...create abort channel...
		abort := make(chan struct{})
		go func() {
			//os.Stdin.Read(make([]byte, 1)) // read a single byte
			abort <- struct{}{}
		}()
		select {
		case <-time.After(10 * time.Second):
			// Do nothing.
		default:
		}
		echo(c, input.Text(), 1*time.Second)
	}

	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}
