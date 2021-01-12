package main

import "sync"

var (
	//mutex会保护共享变量
	mu   sync.Mutex // 类似互斥锁，Lock方法拿token(这里叫锁)，并且Unlock放token：
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}
