## 第八章　Goroutines和Channels
### 8.7. 基于select的多路复用
select的语法：
```go
select {
case <-ch1:
    // ...
case x := <-ch2:
    // ...use x...
case ch3 <- y:
    // ...
default:
    // ...
}
```

每一个case都代表一个通信操作。

select会等待case中有能够执行的case时去执行。当条件满足时，select才会去通信并执行case之后的语句；这时候其它通信是不会执行的。一个没有任何case的select语句写作select{}，会永远地等待下去。

下面的select语句会在abort channel中有值时，从其中接收值；无值时什么都不做。这是一个非阻塞的接收操作；反复地做这样的操作叫做 ==“轮询channel”== 。
```go
select {
case <-abort:
    fmt.Printf("Launch aborted!\n")
    return
default:
    // do nothing
}
```

channel的零值是nil。因为对一个nil的channel发送和接收操作会永远阻塞，在select语句中操作nil的channel永远都不会被select到。这使得我们可以用nil来激活或者禁用case，来达成处理其它输入或输出事件时超时和取消的逻辑。

### 8.9 并发的退出
Go语言并没有提供在一个goroutine中终止另一个goroutine的方法，由于这样会导致goroutine之间的共享变量落在未定义的状态上。

为了能够达到我们退出goroutine的目的，我们需要更靠谱的策略，来通过一个频道把消息广播出去，这样goroutine们就能看到这条事件消息，并且在事件完成之后，可以知道这件事已经发生过了。
