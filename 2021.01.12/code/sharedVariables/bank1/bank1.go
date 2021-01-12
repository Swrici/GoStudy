// Package bank provides a concurrency-safe bank with one account.
package main

var deposits = make(chan int) // 负责给 deposit函数发送金额
var balances = make(chan int) // 负责接受balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int // 余额仅限于出teller goroutines
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // 启动监控goroutine
}
