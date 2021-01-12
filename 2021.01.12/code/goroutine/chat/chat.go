package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//一个只发送string的channel结构体
type client chan<- string

var (
	//刚进入客户端chan
	entering = make(chan client)
	//离开客户端的chan
	leaving  = make(chan client)
	//检录string
	messages = make(chan string) // all incoming client messages
)
//broadcaster函数，转发信息到每个客户端，然后负责检索开新增和离开的用户
func broadcaster() {
	//一个set client为key bool为值
	clients := make(map[client]bool)

	for {
		//select能够执行的case时去执行
		select {

		case msg := <-messages:
			// 负责向所有人广播传入消息
			// 遍历clients中的值，写入相关信息
			for cli := range clients {
				//将消息输送到每个客户端
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func main() {
	//监听tcp里的8000端口
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		//日志报错
		log.Fatal(err)
	}
	//运行broadcaster函数
	go broadcaster()
	//进入循环
	for {
		//Accept等待并返回到listener的下一个连接。
		conn, err := listener.Accept()
		//日志报错
		if err != nil {
			log.Print(err)
			continue
		}
		//运行handleConn函数
		go handleConn(conn)
	}
}

//handleConn 有新的客户端登入显示信息，输出显示其他客户端输出信息
func handleConn(conn net.Conn) {
	//传出客户端消息
	ch := make(chan string)
	//客户端写入数据
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch
	//扫描conn中连接中有无接受到新的信息
	input := bufio.NewScanner(conn)
	for input.Scan() {
		//有的话，就把信息输入到messages中，由广播进行传递
		messages <- who + ": " + input.Text()
	}

	//当跳出了检查输入循环，证明已经离开，将ch中保存的地址信息传输给Leaving，再显示已离开信息，然后关闭连接
	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

//向conn写入信息
func clientWriter(conn net.Conn, ch <-chan string) {
	//向conn写入信息
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}



