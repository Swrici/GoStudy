@[TOC](目录)
### go编写一个聊天服务
服务端代码：
```go
//一个只发送string的channel结构体
type client chan<- string

var (
	//刚进入客户端chan
	entering = make(chan client)
	//离开客户端的chan
	leaving  = make(chan client)
	//检录string
	messages = make(chan string) 
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
```
## 第九章　基于共享变量的并发

在本章中，我们会细致地了解并发机制。尤其是在多goroutine之间的**共享变量，并发问题的分析手段**，以及**解决这些问题的基本模式**。最后我们会**解释goroutine和操作系统线程之间的技术上的一些区别**。

### 9.1. 竞争条件
一般情况下我们没法去知道分别位于两个goroutine的事件x和y的执行顺序，x是在y之前还是之后还是同时发生是没法判断的。当我们没有办法自信地确认一个事件是在另一个事件的前面或者后面发生的话，就说明x和y这两个事件是并发的。

也就是并发状况下，不同goroutine的具体语句彼此前后执行的顺序。

- **并发的函数安全**：
一个函数在线性程序中可以正确地工作。

- **类型的并发安全**：
对于某个类型来说，如果其所有可访问的方法和操作都是并发安全的话，那么该类型便是并发安全的。

- **包级别的导出函数一般情况下都是并发安全的**。由于package级的变量没法被限制在单一的gorouine，所以修改这些变量“必须”使用互斥条件。

竞争条件指的是**程序在多个goroutine交叉执行操作时，没有给出正确的结果**。两个goroutine并发访问同一变量，且至少其中的一个是写操作的时候就会发生数据竞争。可重复读！

**数据竞争发生条件：**
数据竞争会在两个以上的goroutine并发访问相同的变量且至少其中一个为写操作时发生。

避免数据竞争的三种方法：
1. 不发生写操作。
2. 一个变量最多只有一个goroutine进行直接访问，其他的只能使用一个channel来发送请求给指定的goroutine来查询更新变量。也就是Go的口头禅“不要使用共享数据来通信；使用通信来共享数据”。
3. 第三种避免数据竞争的方法是允许很多goroutine去访问变量，但是在同一个时刻最多只有一个goroutine在访问。这种方式被称为“互斥”。

一个提供对一个指定的变量通过channel来请求的goroutine叫做这个变量的**monitor（监控）goroutine**。==例如==：broadcaster goroutine会监控clients map的全部访问。

如果流水线的每一个阶段都能够避免在将变量传送到下一阶段后再去访问它，那么对这个变量的所有访问就是线性的。

### 9.2. sync.Mutex互斥锁
我们可以用一个容量只有1的channel来保证最多只有一个goroutine在同一时刻访问一个共享变量。**一个只能为1和0的信号量叫做二元信号量（binary semaphore）**。
```go
var (
	//容量为1 ，阻塞后续操作
	sema = make(chan struct{}, 1) // 保护平衡的二进制信号量
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // 获得互斥锁
	balance = balance + amount
	<-sema // 释放互斥锁
}

func Balance() int {
	sema <- struct{}{} // 获得互斥锁
	b := balance
	<-sema // 释放互斥锁
	return b
}
```
这种函数、互斥锁和变量的编排叫作监控monitor

`注：goroutine在结束后释放锁是必要的，无论以哪条路径通过函数都需要释放，即使是在错误路径中，也要记得释放。`

Go的mutex**不能重入**这一点我们有很充分的理由。mutex的目的是确保共享变量在程序执行时的关键点上能够保证不变性。

### 9.3. sync.RWMutex读写锁
Go语言提供的sync.RWMutex：
其允许多个只读操作并行执行，但写操作会完全互斥。这种锁叫作“多读单写”锁。
```go
var mu sync.RWMutex
var balance int
func Balance() int {
    mu.RLock() // readers lock
    defer mu.RUnlock()
    return balance
}
```

RWMutex只有当获得锁的大部分goroutine都是读操作，而锁在竞争条件下，也就是说，goroutine们必须等待才能获取到锁的时候，RWMutex才是最能带来好处的。RWMutex需要更复杂的内部记录，所以会让它比一般的无竞争锁的mutex慢一些。

### 9.4. 内存同步
Balance方法需要用到互斥条件,这里使用mutex有两方面考虑。第一Balance不会在其它操作比如Withdraw“中间”执行。第二（更重要的）是“同步”不仅仅是一堆goroutine执行顺序的问题，同样也会涉及到内存的问题。

`前置知识：为了效率，对内存的写入一般会在每一个处理器中缓冲，并在必要时一起flush到主存。`

这种前置知识情况下这些数据可能会以与当初goroutine写入顺序不同的顺序被提交到主存。像channel通信或者互斥量操作这样的原语会使处理器将其聚集的写入flush并commit，这样goroutine在某个时间点上的执行结果才能被其它处理器上运行的goroutine得到。

### 9.5. sync.Once惰性初始化
如果初始化成本比较大的话，那么将初始化延迟到需要的时候再去做就是一个比较好的选择。（懒汉模式？）

### 9.6. 竞争条件检测
Go的runtime和工具链为我们装备了一个复杂但好用的动态分析工具，**竞争检查器（the race detector）**。

只要在go build，go run或者go test命令后面加上-race的flag，就会使编译器创建一个你的应用的“修改”版或者一个附带了能够记录所有运行期对共享变量访问工具的test

竞争检查器会检查这些事件，会寻找在哪一个goroutine中出现了这样的case，例如其读或者写了一个共享变量，这个共享变量是被另一个goroutine在没有进行干预同步操作便直接写入的。

### 9.7. 示例: 并发的非阻塞缓存
### 9.8. Goroutines和线程
##### 9.8.1. 动态栈

每一个OS线程都有一个固定大小的内存块（一般会是2MB）来做栈，这个栈会用来存储当前正在被调用或挂起（指在调用其它函数时）的函数的内部变量。这个固定大小的栈同时很大又很小。
修改固定的大小可以提升空间的利用率，允许创建更多的线程，并且可以允许更深的递归调用，不过这两者是没法同时兼备的。

##### 9.8.2. Goroutine调度
OS线程会被操作系统内核调度。每几毫秒，一个硬件计时器会中断处理器，会调用一个叫作scheduler的内核函数。

这个scheduler的作用：
挂起当前执行的线程并将它的寄存器内容保存到内存中，检查线程列表并决定下一次哪个线程可以被运行，并从内存中恢复该线程的寄存器信息，然后恢复执行该线程的现场并开始执行线程。

因为操作系统线程是被内核所调度，所以从一个线程向另一个“移动”需要完整的上下文切换，也就是说，保存一个用户线程的状态到内存，恢复另一个线程的到寄存器，然后更新调度器的数据结构。这几步操作很慢，因为其局部性很差需要几次内存访问，并且会增加运行的cpu周期。

##### 9.8.3. GOMAXPROCS
Go的调度器使用了一个叫做GOMAXPROCS的变量来决定会有多少个操作系统的线程同时执行Go的代码。**其默认的值是运行机器上的CPU的核心数**，在一个有8个核心的机器上时，调度器一次会在8个OS线程上去调度GO代码。当然在休眠中的或者在通信中被阻塞的goroutine是不需要一个对应的线程来做调度的。

##### 9.8.4. Goroutine没有ID号


## 第十二章　反射
**前言：**
Go语言提供了一种机制，能够在运行时更新变量和检查它们的值、调用它们的方法和它们支持的内在操作，而不需要在编译时就知道这些变量的具体类型。这种机制被称为反射。

两个至关重要的API：
- 一个是fmt包提供的字符串格式功能，
- 另一个是类似encoding/json和encoding/xml提供的针对特定协议的编解码功能。

### 12.1 为什么需要反射
有时需要在运行时使用到变量对应的格式，但是获取会很麻烦。例如:
```go
func Sprint(x interface{}) string {
	//接口，接口方法是一个string()
	type stringer interface {
		String() string
	}
	//判断类型：根据类型返回对应的格式
	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		// array, chan, func, map, pointer, slice, struct 这些更多的类型都很麻烦
		return "???"
	}
}
```
### 12.2. reflect.Type和reflect.Value

**反射是由 reflect 包提供的。 它定义了两个重要的类型, Type 和 Value.** 

- 一个 Type 表示一个Go类型. 它是一个接口, 有许多方法来区分类型以及检查它们的组成部分, 例如一个结构体的成员或一个函数的参数等. 唯一能反映 reflect.Type 实现的是接口的类型描述信息(§7.5), 也正是这个实体标识了接口值的动态类型。

`注：reflect.TypeOf 返回的是一个动态类型的接口值`

- reflect 包中另一个重要的类型是 Value. 一个 reflect.Value 可以装载**任意类型的值**。 函数 reflect.ValueOf 接受**任意的 interface{} 类型**, 并**返回一个装载着其动态值的 reflect.Value**. 和 reflect.TypeOf 类似, reflect.ValueOf 返回的结果也是具体的类型, 但是 reflect.Value 也可以持有一个接口值。
```go
v := reflect.ValueOf(3) // a reflect.Value
fmt.Println(v)          // "3"
fmt.Printf("%v\n", v)   // "3"
fmt.Println(v.String()) // NOTE: "<int Value>"
```

对 Value 调用 Type 方法将返回具体类型所对应的 reflect.Type:
```go
t := v.Type()           // a reflect.Type
fmt.Println(t.String()) // "int"
```
reflect.ValueOf 的逆操作是 reflect.Value.Interface 方法. 它返回一个 interface{} 类型，装载着与 reflect.Value 相同的具体值:
```go
v := reflect.ValueOf(3) // a reflect.Value
x := v.Interface()      // an interface{}
i := x.(int)            // an int
fmt.Printf("%d\n", i)   // "3"
```
reflect.Value 和 interface{} 都能装载任意的值. 所不同的是, 一个空的接口隐藏了值内部的表示方式和所有方法, 因此只有我们知道具体的动态类型才能使用类型断言来访问内部的值(就像上面那样), 内部值我们没法访问. 相比之下, 一个 Value 则有很多方法来检查其内容, 无论它的具体类型是什么。

### 12.3. Display，一个递归的值打印器
让我们看看如何改善聚合数据类型的显示。

目的：只是构建一个用于调试用的Display函数：给定任意一个复杂类型 x，打印这个值对应的完整结构，同时标记每个元素的发现路径。

==分析==：

**Slice和数组**： 两种的处理逻辑是一样的。Len方法返回slice或数组值中的元素个数，Index(i)活动索引i对应的元素，返回的也是一个reflect.Value；如果索引i超出范围的话将导致panic异常，这与数组或slice类型内建的len(a)和a[i]操作类似。display针对序列中的每个元素递归调用自身处理，我们通过在递归处理时向path附加“[i]”来表示访问路径。

`注：虽然reflect.Value类型带有很多方法，但是只有少数的方法能对任意值都安全调用。例如，Index方法只能对Slice、数组或字符串类型的值调用，如果对其它类型调用则会导致panic异常。`

**结构体**： NumField方法报告结构体中成员的数量，Field(i)以reflect.Value类型返回第i个成员的值。成员列表也包括通过匿名字段提升上来的成员。为了在path添加“.f”来表示成员路径，我们必须获得结构体对应的reflect.Type类型信息，然后访问结构体第i个成员的名字。

**Maps**: MapKeys方法返回一个reflect.Value类型的slice，每一个元素对应map的一个key。和往常一样，遍历map时顺序是随机的。**MapIndex(key)返回map中key对应的value**。我们向path添加“[key]”来表示访问路径。

**指针**： Elem方法返回指针指向的变量，依然是reflect.Value类型。即使指针是nil，这个操作也是安全的，在这种情况下指针是Invalid类型，但是我们可以用IsNil方法来显式地测试一个空指针，这样我们可以打印更合适的信息。我们在path前面添加“*”，并用括弧包含以避免歧义。

**接口**： 再一次，我们使用IsNil方法来测试接口是否是nil，如果不是，我们可以调用v.Elem()来获取接口对应的动态值，并且打印对应的类型和值。

反射能够访问到结构体中未导出的成员。

### 12.4. 示例: 编码为S表达式

Go语言的标准库支持了包括JSON、XML和ASN.1等多种编码格式。还有另一种依然被广泛使用的格式是S表达式格式，采用Lisp语言的语法。但是和其他编码格式不同的是，Go语言自带的标准库并不支持S表达式，主要是因为它没有一个公认的标准规范。

在本节中，我们将定义一个包用于将任意的Go语言对象编码为S表达式格式，它支持以下结构：
>42          integer
"hello"     string (带有Go风格的引号)
foo         symbol (未用引号括起来的名字)
(1 2 3)     list   (括号包起来的0个或多个元素)

### 12.5. 通过reflect.Value修改值
讨论如何通过反射机制来修改变量。

`注:一个变量就是一个可寻址的内存空间，里面存储了一个值，并且存储的值可以通过内存地址来更新。`

示例：
```go
x := 2                   // value   type    variable?
a := reflect.ValueOf(2)  // 2       int     no
b := reflect.ValueOf(x)  // 2       int     no
c := reflect.ValueOf(&x) // &x      *int    no
d := c.Elem()            // 2       int     yes (x)
```
分析：
- a对应的变量不可取地址。因为a中的值仅仅是整数2的拷贝副本。
- b中的值也同样不可取地址。
- c中的值还是不可取地址，它只是一个指针&x的拷贝。实际上，所有通过reflect.ValueOf(x)返回的reflect.Value都是不可取地址的。
- 但是对于d，它是c的解引用方式生成的，指向另一个变量，因此是可取地址的。
- 我们可以通过调用reflect.ValueOf(&x).Elem()，来获取任意变量x对应的可取地址的Value。

通过调用reflect.Value的==CanAddr方法==来判断其是否可以被取地址：
```go
fmt.Println(a.CanAddr()) // "false"
fmt.Println(b.CanAddr()) // "false"
fmt.Println(c.CanAddr()) // "false"
fmt.Println(d.CanAddr()) // "true"
```
规则：

每当我们通过指针间接地获取的reflect.Value都是可取地址的，即使开始的是一个不可取地址的Value。

要**从变量对应的可取地址的reflect.Value来访问变量**需要三个步骤：
- 第一步是调用Addr()方法，它返回一个Value，里面保存了指向变量的指针。
- 然后是在Value上调用Interface()方法，也就是返回一个interface{}，里面包含指向变量的指针。
- 最后，如果我们知道变量的类型，我们可以使用类型的断言机制将得到的interface{}类型的接口强制转为普通的类型指针。


不使用指针，也可以通过调用可取地址的**reflect.Value**的**reflect.Value.Set方法**来更新对于的值：
```go
//尝试指针更新
x := 66
//获取任意变量x对应的可取地址的Value。
y := reflect.ValueOf(&x).Elem()
//进行调用可取地址的reflect.Value的reflect.Value.Set方法
y.Set(reflect.ValueOf(88))
fmt.Println(x)
```

Set方法的panic异常:
- Set方法将在运行时执行和编译时进行类似的可赋值性约束的检查，如果类型不同则会导致panic。
- 对一个不可取地址的reflect.Value调用Set方法也会导致panic异常
- 对于一个引用interface{}类型的reflect.Value调用SetInt会导致panic异常，即使那个interface{}变量对于整数类型也不行。

```go
x := 1
rx := reflect.ValueOf(&x).Elem()
rx.SetInt(2)                     // OK, x = 2
rx.Set(reflect.ValueOf(3))       // OK, x = 3
rx.SetString("hello")            // panic: string is not assignable to int
rx.Set(reflect.ValueOf("hello")) // panic: string is not assignable to int

var y interface{}
ry := reflect.ValueOf(&y).Elem()
ry.SetInt(2)                     // panic: SetInt called on interface Value
ry.Set(reflect.ValueOf(3))       // OK, y = int(3)
ry.SetString("hello")            // panic: SetString called on interface Value
ry.Set(reflect.ValueOf("hello")) // OK, y = "hello"
```

利用反射机制并不能修改结构体中未导出的成员。
```go
stdout := reflect.ValueOf(os.Stdout).Elem() // *os.Stdout, an os.File var
fmt.Println(stdout.Type())                  // "os.File"
fd := stdout.FieldByName("fd")
fmt.Println(fd.Int()) // "1"
fd.SetInt(2)          // panic: unexported field
```
一个可取地址的reflect.Value会记录一个结构体成员是否是未导出成员，如果是的话则拒绝修改操作。因此，CanAddr方法并不能正确反映一个变量是否是可以被修改的。另一个相关的方法CanSet是用于检查对应的reflect.Value是否是可取地址并可被修改的：
```go
fmt.Println(fd.CanAddr(), fd.CanSet()) // "true false"
```

### 12.6. 示例: 解码S表达式
### 12.7. 获取结构体字段标识
### 12.8. 显示一个类型的方法集
```go
// Print方法打印x的方法集
func Print(x interface{}) {
	//获得x的值
    v := reflect.ValueOf(x)
    //拿到x的类型
    t := v.Type()
    //打印类型
    fmt.Printf("type %s\n", t)
	//遍历该类型的方法 v.NumMethod()
    for i := 0; i < v.NumMethod(); i++ {
        methType := v.Method(i).Type()
        fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name,
            strings.TrimPrefix(methType.String(), "func"))
    }
}
```

### 12.9. 几点忠告
反射应该被小心地使用，原因有三：
- 第一个原因是，基于反射的代码是比较脆弱的。而反射则是在真正运行到的时候才会抛出panic异常，可能是写完代码很久之后了，而且程序也可能运行了很长的时间。
- 避免使用反射的第二个原因是，即使对应类型提供了相同文档，但是反射的操作不能做静态类型检查，而且大量反射的代码通常难以理解。
- 第三个原因，基于反射的代码通常比正常的代码运行速度慢一到两个数量级。对于一个典型的项目，大部分函数的性能和程序的整体性能关系不大，所以使用反射可能会使程序更加清晰。
