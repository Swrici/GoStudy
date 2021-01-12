//添加一个Withdraw(amount int)取款函数。
//其返回结果应该要表明事务是成功了还是因为没有足够资金失败了。
//这条消息会被发送给monitor的goroutine，
//且消息需要包含取款的额度和一个新的channel，
//这个新channel会被monitor goroutine来把boolean结果发回给Withdraw。
package main

var deposits = make(chan int) // 负责给 deposit函数发送金额
var balances = make(chan int) // 负责接受balance
var withdrowBool = make(chan bool)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func Withdraw(amount int) bool {
	Deposit(-amount)
	if Balance() < 0  {
		Deposit(amount)
		return false
	}

	withdrowBool <- false
	return true
}
func teller() {
	var balance int // 余额仅限于出teller goroutines
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			Withdraw(balance)
		case <-withdrowBool:
			Withdraw(balance)
		}
	}
}

func init() {
	go teller() // 启动监控goroutine
}
