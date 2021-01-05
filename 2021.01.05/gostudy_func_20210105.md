## 5 函数
### 5.2 递归
函数可以是递归的，这意味着函数可以直接或间接的调用自身。
### 5.3 多返回值
- 在Go中，一个函数可以返回多个值。

- Go的垃圾回收机制会回收不被使用的内存，但是这不包括操作系统层面的资源，比如打开的文件、网络连接。因此我们必须显式的释放这些资源。

- 当调用接受多参数的函数时，可以将一个返回多参数的函数作为该函数的参数。
- 如果一个函数将所有的返回值都显示的变量名，那么该函数的return语句可以省略操作数。这称之为bare return。
### 5.4 错误

- panic是来自被调函数的信号，表示发生了某个已知的bug。 
- 错误是软件包API和应用程序用户界面的一个重要组成部分，程序运行失败仅被认为是几个预期的结果之一。

由于对于某个应该在控制流程中处理的错误而言，将这个错误以异常的形式抛出会混乱对错误的描述，这通常会导致一些糟糕的后果。当某个程序错误被当作异常处理后，这个错误会将堆栈根据信息返回给终端用户，这些信息复杂且无用，无法帮助定位错误。

但是，在Go中，函数运行失败时会返回错误信息，这些错误信息被认为是一种预期的值而非异常（exception），**Go使用控制流机制（如if和return）处理异常，这使得编码人员能更多的关注错误处理**。

#### 5.4.1. 错误处理策略
##### 5.4.1.1 传播错误
这意味着函数中某个子程序的失败，会变成该函数的失败。
**例1：**
```go
resp, err := http.Get(url)
if err != nil{
    return nil, err
}
```

**例2：**
```go
doc, err := html.Parse(resp.Body)
resp.Body.Close()
if err != nil {
    return nil, fmt.Errorf("parsing %s as HTML: %v", url,err)
}
```
**fmt.Errorf函数**使用fmt.Sprintf格式化错误信息并返回。我们使用该函数前缀添加额外的上下文信息到原始错误信息。
##### 5.4.1.2 重新尝试失败的操作
如果错误的发生是**偶然性**的，或由**不可预知**的问题导致的。一个明智的选择是重新尝试失败的操作。在重试时，我们需要限制重试的时间间隔或重试的次数，防止无限制的重试。

##### 5.4.1.3 输出错误信息并结束程序
如果错误发生后，程序无法继续运行，我们就可以采用第三种策略：输出错误信息并结束程序。需要注意的是，这种策略**只应在main中执行**。对库函数而言，应仅向上传播错误，除非该错误意味着程序内部包含不一致性，即遇到了bug，才能在库函数中结束程序。
```go
// (In function main.)
if err := WaitForServer(url); err != nil {
    fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
    os.Exit(1)
}
```

##### 5.4.1.4 只输出错误信息
第四种策略：有时，我们只需要输出错误信息就足够了，不需要中断程序的运行。我们可以通过log包提供函数。
```go
//log包中的所有函数会为没有换行符的字符串增加换行符。
if err := Ping(); err != nil {
    log.Printf("ping failed: %v; networking disabled",err)
}
//或者
if err := Ping(); err != nil {
    fmt.Fprintf(os.Stderr, "ping failed: %v; networking disabled\n", err)
}
```


##### 5.4.1.5 直接忽略掉错误

```go
dir, err := ioutil.TempDir("", "scratch")
if err != nil {
    return fmt.Errorf("failed to create temp dir: %v",err)
}
// ...use temp dir…
os.RemoveAll(dir) // ignore errors; $TMPDIR is cleaned periodically
```

#### 5.4.2. 文件结尾错误（EOF）
io包保证任何由文件结束引起的读取失败都返回同一个错误——io.EOF.
原因可以由下面这个例子中体现出来：
- 从文件中读取n个字节。如果n等于文件的长度，读取过程的任何错误都表示失败。
- 如果n小于文件的长度，调用者会重复的读取固定大小的数据直到文件结束。
- 这会导致调用者必须分别处理由文件结束引起的各种错误。

因为文件结束这种错误不需要更多的描述，所以io.EOF有固定的错误信息——“EOF”。

只需要简单的比对就能找出对应错误位置。

```go
in := bufio.NewReader(os.Stdin)
for {
    r, _, err := in.ReadRune()
    if err == io.EOF {
        break // finished reading
    }
    if err != nil {
        return fmt.Errorf("read failed:%v", err)
    }
    // ...use r…
}
```

### 5.5 函数值

在Go中，函数被看作**第一类值（first-class values）**：函数像其他值一样，拥有类型，可以被赋值给其他变量，传递给函数，从函数返回。对函数值（function value）的调用类似函数调用。

函数类型的零值是nil。调用值为nil的函数值会引起panic错误。
```go
var f func(int) int
f(3) // 此处f的值为nil, 会引起panic错误
```
函数值使得我们不仅仅可以**通过数据来参数化函数**，**亦可通过行为**。

```go
func add1(r rune) rune { return r + 1 }

fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
fmt.Println(strings.Map(add1, "Admix"))    // "Benjy"
```

### 5.6 匿名函数
拥有函数名的函数只能在包级语法块中被声明。

但有时要在非包级语法块中使用声明函数，这时就要用到匿名函数。

通过函数字面量（function literal），我们可绕过这一限制，在任何表达式中表示一个函数值。函数字面量的语法和函数声明相似，区别在于func关键字后没有函数名。函数值字面量是一种表达式，它的值被称为匿名函数（anonymous function）。

```go
func add1(r rune) rune { return r + 1 }
fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
//等价于
strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")
```
通过匿名函数方式定义的函数可以访问完整的词法环境（lexical environment），意思就是在函数中定义的内部函数可以引用该函数的变量。

函数值不仅仅是一串代码，还记录了状态。
在下例squares中定义的匿名内部函数可以访问和更新squares中的局部变量，这意味着匿名函数和squares中，存在变量引用。

这就是函数值属于引用类型和函数值不可比较的原因。

Go使用闭包（closures）技术实现函数值，Go程序员也把函数值叫做闭包。

```go
// squares返回一个匿名函数。
// 该匿名函数每次被调用时都会返回下一个数的平方。
func squares() func() int {
    var x int
    return func() int {
        x++
        return x * x
    }
}
func main() {
    f := squares()
    fmt.Println(f()) // "1"
    fmt.Println(f()) // "4"
    fmt.Println(f()) // "9"
    fmt.Println(f()) // "16"
}
```

#### 5.6.1. 警告：捕获迭代变量

本节，将介绍**Go词法作用域的一个陷阱**。
举例说明：
```go
var rmdirs []func()
for _, d := range tempDirs() {
    dir := d // NOTE: necessary!
    os.MkdirAll(dir, 0755) // creates parent directories too
    rmdirs = append(rmdirs, func() {
        os.RemoveAll(dir)
    })
}
// ...do some work…
for _, rmdir := range rmdirs {
    rmdir() // clean up
}
```

```go
var rmdirs []func()
for _, dir := range tempDirs() {
    os.MkdirAll(dir, 0755)
    rmdirs = append(rmdirs, func() {
        os.RemoveAll(dir) // NOTE: incorrect!
    })
}
```
问题的原因在于循环变量的作用域。在上面的程序中，for循环语句引入了新的词法块，循环变量dir在这个词法块中被声明。在该循环中生成的所有函数值都共享相同的循环变量。需要注意，函数值中记录的是循环变量的内存地址，而不是循环变量某一时刻的值。以dir为例，后续的迭代会不断更新dir的值，当删除操作执行时，for循环已完成，dir中存储的值等于最后一次迭代的值。这意味着，每次对os.RemoveAll的调用删除的都是相同的目录。

### 5.7. 可变参数
```go
//参数个数可控
//调用者隐式的创建一个数组，并将原始参数复制到数组中，再把数组的一个切片作为参数传给被调函数。
func sum(vals...int) int {
    total := 0
    for _, val := range vals {
        total += val
    }
    return total
}

values := []int{1, 2, 3, 4}
fmt.Println(sum(values...)) // "10"
```
### 5.8. Deferred函数
Deferred函数的**作用**
随着函数变得复杂，需要处理的错误也变多，维护清理逻辑变得越来越困难。而Go语言独有的defer机制可以让事情变得简单。

作用机制
只需在调用**普通函数或方法前加上关键字defer**，就完成了defer所需要的语法。

当defer语句被执行时，跟在defer后面的函数会被延迟执行。**直到包含该defer语句的函数执行完毕时**，**defer后的函数才会被执行**，不论包含defer语句的函数是通过return正常结束，还是由于panic导致的异常结束。你可以在一个函数中执行多条defer语句，它们的执行顺序与声明顺序相反。
```go
func main() {
    f(3)
}
func f(x int) {
    fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
    defer fmt.Printf("defer %d\n", x)
    f(x - 1)
}
```
>输出如下：
>f(3)
f(2)
f(1)
defer 1
defer 2
defer 3


deferred函数的使用
- 调试复杂程序时，defer机制也常被用于记录何时进入和退出函数。
- 不论函数逻辑多复杂，都能保证在任何执行路径下，资源的释放。


被延迟执行的匿名函数甚至可以修改函数返回给调用者的返回值
```go
func triple(x int) (result int) {
    defer func() { result += x }()
    return double(x)
}
fmt.Println(triple(4)) // "12"
```

### 5.9. Panic异常
如数组访问越界、空指针引用等。这些**运行时错误**会引起painc异常。

一般而言，当panic异常发生时，程序会中断运行，并立即执行在该goroutine（可以先理解成线程，在第8章会详细介绍）中被延迟的函数（defer 机制）。

在Go的panic机制中，延迟函数的调用在释放堆栈信息之前。

### 5.10. Recover捕获异常
**如果在deferred函数中调用了内置函数recover，并且定义该defer语句的函数发生了panic异常，recover会使程序从panic中恢复，并返回panic value**。导致panic异常的函数不会继续运行，但能正常返回。在未发生panic时调用recover，recover会返回nil。
