## 第八章　Goroutines和Channels
### 8.2 示例: 并发的Clock服务
使用Goroutines创建服务端和客户端，实现并发功能。例子是让服务端为多个客户端不断返回当前时间。

### 8.3 示例: 并发的Echo服务

### 8.4 Channels
如果说**goroutine是Go语言程序的并发体**的话，那么**channels则是它们之间的通信机制**。

一个channel是一个通信机制，它可以让一个goroutine通过它给另一个goroutine发送值信息。

每个channel都有一个特殊的类型，也就是channels可发送数据的类型。一个可以发送int类型数据的channel一般写为chan int。

和map类似，**channel也对应一个make创建的底层数据结构的引用**。当我们**复制一个channel或用于函数参数传递时，我们只是拷贝了一个channel引用**，因此调用者和被调用者将引用同一个channel对象。和其它的引用类型一样，**channel的零值也是nil**。

两个**相同类型的channel**可以使用==运算符比较。如果两个channel引用的是**相同的对象**，那么比较的**结果为真**。一个channel也可以和nil进行比较。

==channel的操作行为==

一个channel有**发送和接受**两个主要操作，都是通信行为。

 发送和接收两个操作都使用<-运算符。
- 在发送语句中，<-运算符分割channel和要发送的值。
- 在接收语句中，<-运算符写在channel对象之前。
- 一个不使用接收结果的接收操作也是合法的。
```go
ch <- x //x是要发送的值
x  =<- ch // 接收表达式
<-ch //接受表达式，结果被丢弃
```

除发送和接受两种行为之外，还有==关闭==操作。
```go
close(ch)
```

#### 8.4.1 不带缓存的Channels
特点：
- 一个基于无缓存Channels的发送操作将导致发送者goroutine阻塞，直到另一个goroutine在相同的Channels上执行接收操作，
- 反之，如果接收操作先发生，那么接收者goroutine也将阻塞，直到有另一个goroutine在相同的Channels上执行发送操作。

不带缓存的Channels有时候也被称为同步Channels。因为无缓存Channels的发送和接收操作将导致两个goroutine做一次同步操作。

#### 8.4.2 串联的Channels（Pipeline）
有时Channels也被叫成的管道（pipeline）

当Channels将多个goroutine连接在一起，一个Channel的输出作为下一个Channel的输入。这种串联的Channels就是所谓的管道（pipeline）。

当一个channel被关闭后，再向该channel发送数据将导致panic异常。当一个被关闭的channel中已经发送的数据都被成功接收后，后续的接收操作将不再阻塞，它们会立即返回一个零值。

试图重复关闭一个channel将导致panic异常，试图关闭一个nil值的channel也将导致panic异常。关闭一个channels还会触发一个广播机制。

#### 8.4.3 单方向的Channel
Go语言的类型系统提供了单方向的channel类型，分别用于只发送或只接收的channel。

- 类型chan<- int表示一个只发送int的channel，只能发送不能接收。
- 相反，类型<-chan int表示一个只接收int的channel，只能接收不能发送。（箭头<-和关键字chan的相对位置表明了channel的方向。）这种限制将在编译期检测。

因为关闭操作只用于断言不再向channel发送新的数据，所以只有在发送者所在的goroutine才会调用close函数，因此对一个只接收的channel调用close将是一个编译错误。

`注：双向chan可以隐式转换成单向chan，但是单向不管如何（显示隐式）都不能转双向chan。`

#### 8.4.4 带缓存的Channels
带缓存的Channel内部持有一个元素队列。队列的最大容量是在调用make函数创建channel时通过第二个参数指定的。下面的语句创建了一个可以持有三个字符串元素的带缓存Channel。

```go
ch = make(chan string, 3)
```
因为有了缓存空间，所以在使用的时候也有了不同。

- 向缓存Channel的发送操作就是向内部缓存队列的尾部插入元素，
- 接收操作则是从队列的头部删除元素。
- 如果内部缓存队列是满的，那么发送操作将阻塞直到因另一个goroutine执行接收操作而释放了新的队列空间。
- 相反，如果channel是空的，接收操作将阻塞直到有另一个goroutine执行发送操作而向队列插入元素。

```go
func mirroredQuery() string {
    responses := make(chan string, 3)
    go func() { responses <- request("asia.gopl.io") }()
    go func() { responses <- request("europe.gopl.io") }()
    go func() { responses <- request("americas.gopl.io") }()
    return <-responses // return the quickest response
}

func request(hostname string) (response string) { /* ... */ }
```
在上例中，如果我们使用了无缓存的channel，那么两个慢的goroutines将会因为没有人接收而被永远卡住。这种情况，称为goroutines泄漏，这将是一个BUG。和垃圾变量不同，**泄漏的goroutines并不会被自动回收**，因此确保每个不再需要的goroutine能正常退出是重要的。

==无缓存和有缓存chan的使用区别==：
- 无缓存channel更强地保证了每个发送操作与相应的同步接收操作；
- 但是对于带缓存channel，这些操作是解耦的。
- 同样，即使我们知道将要发送到一个channel的信息的数量上限，创建一个对应容量大小的带缓存channel也是不现实的，因为这要求在执行任何接收操作之前缓存所有已经发送的值。如果未能分配足够的缓冲将导致程序死锁。

### 8.5. 并发的循环
本节中，我们会探索一些用来在并行时循环迭代的常见并发模型。

