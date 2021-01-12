package main

var (
	//容量为1 ，阻塞后续操作
	sema = make(chan struct{}, 1) // 保护平衡的二进制信号量
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // acquire token
	balance = balance + amount
	<-sema // release token
}

func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
	return b
}
