@[TOC](目录)
## 7 接口
### 7.6 sort.Interface接口
Go语言的sort.Sort函数不会对具体的序列和它的元素做任何假设。

实现机制：

- 它使用了一个**接口类型sort.Interface**来指定通用的排序算法和可能被排序到的序列类型之间的约定。

- 这个接口的实现由序列的具体表示和它希望排序的元素决定，序列的表示经常是一个切片。

实例演示：
一个内置的排序算法需要知道三个东西：序列的长度，表示两个元素比较的结果，一种交换两个元素的方式；这就是sort.Interface的三个方法：
```go
package sort

type Interface interface {
    Len() int //序列的长度
    Less(i, j int) bool // 表示两个元素比较的结果 i, j are indices of sequence elements
    Swap(i, j int) //交换两个元素的方式
}
```

### 7.7. http.Handler接口
这是一个分发器接口，通过这个接口里的ServeHTTP方法将分发信息给访问的客户端。
```go
package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}

func ListenAndServe(address string, h Handler) error
```

实例：
```go
//初步使用http.Handler接口
//使用了一个将库存商品价格映射成美元的demo
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	//ListenAndServe需要例如“localhost:8000”的服务器地址，和一个所有请求都可以分派的Handler接口实例
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
//定义了一个dollars结构体 类型是float32
type dollars float32

//重写了dollars的String(),输出更方便查看
func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }
//定义了一个database结构体 类型是key为String，value为dollars结构体的map
type database map[string]dollars
//实现了ServerHTTP方法，也就是实现了http.Handler接口
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	//输出给访问的客户端
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	//进入price路径
	case "/price":
		//item是通过get方法获得的参数名称，将获得的参数值在db这个map里进行查询
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	//差不到返回 no such page 提示信息
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

```
但在一个实际的应用中，将每个case中的逻辑定义到一个分开的方法或函数中会很实用。

### 7.8 error接口

```go
type error interface {
    Error() string
}
```

创建一个error最简单的方法就是调用errors.New函数，它会根据传入的错误信息返回一个新的error。

### 7.9 示例: 表达式求值
在本节中，我们会构建一个简单算术表达式的求值器。

先复习表达式是什么：

表达式语言由浮点数符号（小数点）；二元操作符+，-，*， 和/；一元操作符-x和+x；调用pow(x,y)，sin(x)，和sqrt(x)的函数；当然也有括号和标准的优先级运算符

### 7.10 类型断言
类型断言是一个使用在接口值上的操作。语法上它看起来像x.(T)被称为断言类型。

一个类型断言检查**它操作对象的动态类型**是否和**断言的类型匹配。**

这个类型断言这里有两种可能。

1. 第一种，如果**断言的类型T是一个具体类型**，然后**类型断言检查x的动态类型是否和T相同**。如果这个检查**成功**了，类型断言的结果是**x的动态值，当然它的类型是T**。如果检查失败，接下来这个操作会抛出panic。换句话说，==具体类型的类型断言从它的操作对象中获得具体的值==。

实例：
```go
var w io.Writer
w = os.Stdout
f := w.(*os.File)      // success: f == os.Stdout
c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
```
2. 第二种，如果相反断言的**类型T是一个接口类型**，然后**类型断言检查是否x的动态类型满足T**。如果这个检查**成功**了，**动态值没有获取到**；这个结果仍然是**一个有相同类型和值部分的接口值**，但是结果有类型T。换句话说，对**一个接口类型的类型断言改变了类型的表述方式**，改变了可以获取的方法集合（通常更大），但是它**保护了接口值内部的动态类型和值的部分**。

实例：
```go
var w io.Writer
w = os.Stdout
rw := w.(io.ReadWriter) // success: *os.File has both Read and Write
w = new(ByteCounter)
rw = w.(io.ReadWriter) // panic: *ByteCounter has no Read method
```

**如果断言操作的对象是一个nil接口值，那么不论被断言的类型是什么这个类型断言都会失败**。我们几乎不需要对一个更少限制性的接口类型（更少的方法集合）做断言，因为它表现的就像赋值操作一样，除了对于nil接口值的情况。

```go
w = rw             // io.ReadWriter 可以分配给 io.Writer
w = rw.(io.Writer) // fails only if rw == nil 当rw等于Nil时失败
```

经常地我们对一个接口值的动态类型是不确定的，并且我们更愿意去检验它是否是一些特定的类型。

如果类型断言出现在一个预期有两个结果的赋值操作中，例如如下的定义，这个操作不会在失败的时候发生panic但是代替地返回一个额外的第二个结果，这个结果是一个标识成功的布尔值。
```go
var w io.Writer = os.Stdout
f, ok := w.(*os.File)      // success:  ok, f == os.Stdout
b, ok := w.(*bytes.Buffer) // failure: !ok, b == nil
```
第二个结果常规地赋值给一个命名为ok的变量。如果这个操作失败了，那么ok就是false值，第一个结果等于被断言类型的零值，在这个例子中就是一个nil的*bytes.Buffer类型。

这个ok结果经常立即用于决定程序下面做什么。if语句的扩展格式让这个变的很简洁。

### 7.11 基于类型断言区别错误类型
思考在os包中文件操作返回的错误集合。I/O可以因为任何数量的原因失败，但是有三种经常的错误必须进行不同的处理：
1. 文件已经存在（对于创建操作），
2. 找不到文件（对于读取操作），
3. 权限拒绝。
```go
package os

func IsExist(err error) bool //1 文件已经存在（对于创建操作），
func IsNotExist(err error) bool //2.  找不到文件（对于读取操作），
func IsPermission(err error) bool //权限拒绝。
```

### 7.12 通过类型断言询问行为
下面这段逻辑和net/http包中web服务器负责写入HTTP头字段（例如："Content-type:text/html）的部分相似。io.Writer接口类型的变量w代表HTTP响应；写入它的字节最终被发送到某个人的web浏览器上。
```go
func writeHeader(w io.Writer, contentType string) error {
    if _, err := w.Write([]byte("Content-Type: ")); err != nil {
        return err
    }
    if _, err := w.Write([]byte(contentType)); err != nil {
        return err
    }
}
```
[]byte(str)中包含着一个分配内存再拷贝的过程。

io.Writer接口告诉我们关于w持有的具体类型的唯一东西：就是可以向它写入字节切片。如果我们回顾net/http包中的内幕，我们知道**在这个程序中的w变量持有的动态类型也有一个允许字符串高效写入的WriteString方法**；这个方法会避免去分配一个临时的拷贝

我们可以定义一个**只有这个方法的新接口并且使用类型断言来检测是否w的动态类型满足这个新接口**


### 7.13. 类型开关

接口被以两种不同的方式使用。
1. 一个接口的方法表达了实现这个接口的具体类型间的相似性，但是隐藏了代表的细节和这些具体类型本身的操作。重点在于方法上，而不是具体的类型上。以io.Reader，io.Writer，fmt.Stringer，sort.Interface，http.Handler，和error为典型。
2. 第二个方式利用**一个接口值可以持有各种具体类型值的能力并且将这个接口认为是这些类型的union（联合）**。类型断言用来动态地区别这些类型并且对每一种情况都不一样。在这个方式中，重点在于具体的类型满足这个接口，而不是在于接口的方法（如果它确实有一些的话），并且没有任何的信息隐藏。我们将以这种方式使用的接口描述为discriminated unions（可辨识联合）。

### 注意
- 当设计一个新的包时，新手Go程序员总是先创建一套接口，然后再定义一些满足它们的具体类型。这种方式的结果就是有很多的接口，它们中的每一个仅只有一个实现。不要再这么做了。这种接口是不必要的抽象；
- 当一个接口只被一个单一的具体类型实现时有一个例外，就是由于它的依赖，这个具体类型不能和这个接口存在在一个相同的包中。这种情况下，一个接口是解耦这两个包的一个好方式。
- 在Go语言中只有当两个或更多的类型实现一个接口时才使用接口，它们必定会从任意特定的实现细节中抽象出来。结果就是有更少和更简单方法（经常和io.Writer或 fmt.Stringer一样只有一个）的更小的接口。当新的类型出现时，小的接口更容易满足。对于接口设计的一个好的标准就是 ask only for what you need（只考虑你需要的东西）
## 第八章　Goroutines和Channels
Go语言中的并发程序可以用两种手段来实现。本章讲解**goroutine**和**channel**，其支持“**顺序通信进程**”(communicating sequential processes)或被简称为**CSP**。

### 8.1. Goroutines
每一个并发的执行单元叫作一个goroutine。使用go加函数开启goroutines

### 8.2. 示例: 并发的Clock服务
