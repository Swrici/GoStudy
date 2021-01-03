﻿### 命令行参数
#### build 子命令
语法如下：
```go
go build 文件名.go
```
build命令的作用：

会生成一个可执行的二进制文件，在idea中直接在文件位置输入==文件名字==执行，如果是在windows中使用==文件名.exe== 执行。
![在这里插入图片描述](https://img-blog.csdnimg.cn/20201229154228350.png)
注意事项：
- ==导入不需要的包或者缺少必要的包==时都会导致程序无法使用。

- import必须在package之下


程序的==函数==、==变量==、==常量==、==类型==的声明语句分别由关键字==func==、==var==、==const==、==type==定义。

函数是由==func==关键字、==函数名==、==参数列表==、 ==返回值列表(可选)== 以及 ==函数体==组成。


Go不需要使用分号，除非一行上由多条语句，实际上，编译器会把换行符转换成分号。
```go
func 函数名(){
```
中的 =='{'== 要和==func==在同一行

在表达式`x+y`中，可以在`+`后换行(自动识别，不会转换分号)，但是不能在`+`前换行（识别不到导致报错）

Go的数组元素索引采用==左闭右开==的方式，
`例如：`a = [1, 2, 3, 4, 5], a[0:3] = [1, 2, 3]，不包含最后一个元素。

隐式初始化，默认赋 =="零值"==

range 类似迭代器，可以遍历数组，字符串，map等等，对象的不同，返回的结果也不同。

例如：在每次迭代中是一对值，这对值是原有的数组下标和对应的值。

>for _,arg := range os.Args[1:]{  

 去除 \_, 之后会产生报错,==为什么呢？==

range的语法要求，要处理元素，必须处理索引.

_ 空标识符可用于在任何语法需要变量名但程序逻辑不需要的时候


![在这里插入图片描述](https://img-blog.csdnimg.cn/20201229205256338.png)

```go
fmt.Println("即将输出os.Args[0]：")
fmt.Println(os.Args[0])
```
![在这里插入图片描述](https://img-blog.csdnimg.cn/20201229205311759.png)


```go
fmt.Println("即将依行输出：")
for _,args:=range os.Args{
	fmt.Println(1,args)
	fmt.Println("我是换行符")
}
```
![在这里插入图片描述](https://img-blog.csdnimg.cn/20201229205327950.png)


```go
func BenchmarkIsMoreQuicker1(b *testing.B)  {
	var s,step string
	//b.ResetTimer()
	for i:=0;i<b.N;i++ {
		for i:=1;i<len(os.Args);i++{
			s += step + os.Args[i]
			step = " "
		}
	}
}

func BenchmarkIsMoreQuicker2(b *testing.B) {
	//b.ResetTimer()
	//start := time.Now();
	for i:=0;i<b.N;i++ {
		strings.Join(os.Args[1:], " ")
	}
}
```

---

```go
counts := make(map[string]int)
```
==map==存储了键/值（key/value）的集合，对集合元素，提供常数时间的存、取或测试操作。键可以是任意类型，只要其值能用\=\=运算符比较，最常见的例子是==字符串==；==值==则可以是==任意类型==。这个例子中的键是字符串，值是整数。==内置函数make创建空map==。

bufio包的使用，读取数据输出数据的应用
```go
package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {
	//make生产一个map，存放类型为string,值类型为int
	count:=make(map[string]int)
	//将键盘输入的值保村道input中
	input := bufio.NewScanner(os.Stdin)
	//遍历input中的值
	for input.Scan() {
		count[input.Text()]++
		if input.Text() == "end" { break }
	}
	//获取遍历map，range获取的键值对中，键总为1
	for line,n:=range count{
		if n>0 {
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}
```

常见的格式：
![在这里插入图片描述](https://img-blog.csdnimg.cn/2020122920523386.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L1Nyd2ljaQ==,size_16,color_FFFFFF,t_70)
go 写一个gif

`并发和并行`
并行是同时执行
并发是交替执行

`go 函数名` 就是并发执行，

```go
//声明不带缓冲的通道
ch1 := make(chan string)
//声明2个不带缓冲的通道
ch2 := make(chan string,2)
//声明只读通道
ch3 := make(<- chan string)
//声明只写通道
ch4 := make(chan <- string)
```
`不带缓冲的通道，进和出都会阻塞，带缓冲的通道，进1长度加1，出1长度减1，当长度等于缓冲长度，再进就会产生阻塞。`

web服务，创建一个服务器

Go语言允许这样的一个简单的语句结果作为局部的变量声明出现在if语句的最前面
```go
if err := r.ParseForm(); err != nil {
        log.Print(err)
    }
```

#### 在url里的参数获取
获取url,get方式的参数：request.URL.Query().Get("参数名字")
字符串转浮点型float32:strconv.ParseFloat(string, 32/64)

可以使用label跳出特定的循环，类似C中的goto。

Go语言里没有指针运算，也就是不能像c语言里可以对指针进行加或减操作。


命名：

命名规则：一个名字必须以一个字母（Unicode字母）或下划线开头，后接字母数字下划线。大小写敏感。驼峰式命名。

一个名字在函数外部定义的整个包内可见，并且当这个名字是大写时，它对包外也是可见的，而如果是小写则对包外是不可见的。


声明：
var、const、type以及func分别对应变量、常量、类型以及函数实体对象的声明。


声明顺序：
每个源文件中以包的声明语句开始，说明该源文件是属于哪个包。包声明语句之后是import语句导入依赖的其它包，然后是包一级的类型、变量、常量、函数的声明语句，包一级的各种类型的声明语句的顺序无关紧要。
`ps:包一级声明语句意味着声明的名字可在整个包对应的每个源文件中访问`

声明---函数声明
一个函数的声明由一个==函数名字==、==参数列表==（由函数的调用者提供参数变量的具体值）、一个==可选的返回值列表==和==包含函数定义的函数体==组成。如果没有返回语句则是执行到函数末尾，然后返回到函数调用者。

声明---变量声明
完整声明语法：
>var 变量名字 类型 = 表达式

其中“类型”或“= 表达式”两个部分==可以省略其中的一个==。
- 如果省略的是类型信息，那么将根据初始化表达式来==推导==变量的类型信息。
- 如果初始化表达式被省略，那么将用==零值==初始化该变量。

例如：
```go
var i, j, k int                 // int, int, int
var b, f, s = true, 2.3, "four" // bool, float64, string
```

各类型对应的零值
-  ==数值==类型变量对应的零值是==0==。
- ==布尔==类型变量对应的零值是==false==。
- ==字符串==类型对应的零值是==空字符串==。
- ==接口==或==引用==类型（包括slice、指针、map、chan和函数）变量对应的零值是==nil==。
- ==数组==或==结构体==等==聚合类型==对应的零值是==每个元素或字段==都是对应==该类型的零值==。

短遍历声明语法：
>变量名 := 变量
>例如：
>freq := rand.Float64() * 3.0
t := 0.0
var names []string
i, j := 0, 1

`:=起的是一个声明作用，=起的是一个赋值作用`

对于上面红字还有一个比较微妙的地方：短变量声明左边的变量==可能并不是全部都是刚刚声明==的。如果有一些==已经==在相同的词法域==声明过了==（§2.7），那么简短变量声明语句对这些已经声明过的变量就==只有赋值行为==了。

```go
//神奇的交换？
i, j = j, i // 交换 i 和 j 的值
```

简短变量声明语句中==必须至少==要声明==一个新的变量==！否则会导致报错。

声明---指针声明
例子：
```go
x := 1
p := &x         // p, of type *int, points to x
fmt.Println(*p) // "1"
*p = 2          // equivalent to x = 2
fmt.Println(x)  // "2"
```
对于每一个变量必然有对应的内存地址。

对于聚合类型每个成员——比如结构体的每个字段、或者是数组的每个元素——也都是对应一个变量，因此可以被取地址。

指针零值
任何类型的指针的零值都是nil

关于带*指针自加的问题
```go
*p++ // 只是增加p指向的变量的值，并不改变p指针！！！
```

声明---new函数声明变量
表达式new(T)将创建一个T类型的匿名变量，初始化为T类型的零值，然后返回变量地址，返回的指针类型为*T。
```go
p := new(int)   // p, *int 类型, 指向匿名的 int 变量
fmt.Println(*p) // "0"
*p = 2          // 设置 int 匿名变量的值为 2
fmt.Println(*p) // "2"
```

每次调用new函数都是返回一个新的变量的地址.
```go
p := new(int)
q := new(int)
fmt.Println(p == q) // "false"
```


变量的生命周期
变量的生命周期指的是在程序运行期间变量有效存在的时间段。

包一级变量生命周期是和整个程序的运行周期是一致的。

局部变量的生命周期则是动态的：每次从创建一个新变量的声明语句开始，直到该变量不再被引用为止，然后变量的存储空间可能被回收。


`Go编辑器的特性之一： 最后插入的逗号不会导致编译错误，这是Go编译器的一个特性。`

是为了防止编译器在行尾自动插入分号而导致的编译错误，可以在末尾的参数变量后面显式插入逗号。

通过变量是否可达决定是否回收，意味着局部变量返回的同时可能仍会存在，也就是局部变量逃逸。

逃逸的变量需要额外分配内存，同时对性能的优化可能会产生细微的影响。


赋值
赋值---元组赋值
在赋值之前，赋值语句右边的所有表达式将会先进行求值，然后再统一更新左边对应变量的值。
举例：
```go
i, j, k = 2, 3, 5
x, y = y, x
a[i], a[j] = a[j], a[i]
```

特殊的元组赋值：
```go
v = m[key]                // map查找，失败时返回零值
v = x.(T)                 // type断言，失败时panic异常
v = <-ch                  // 管道接收，失败时返回零值（阻塞不算是失败）

_, ok = m[key]            // map返回2个值
_, ok = mm[""], false     // map返回1个值
_ = mm[""]                // map返回1个值
```

Golang Printf、Sprintf 、Fprintf 格式化对比

Printf : 只可以打印出格式化的字符串,可以输出字符串类型的变量，不可以输出整形变量和整形。
Println :可以打印出字符串，和变量。

Sprintf 则格式化并返回一个字符串而不带任何输出。

s := fmt.Sprintf("是字符串 %s ",“string”)

fmt.Println(s) // 是字符串 %s 对应 是字符串 string
可以使用 Fprintf 来格式化并输出
>fmt.Fprintf(os.Stderr, "格式化 %s\n", "error")

作用域：声明语句的作用域对应的是一个源代码的文本区域；它是一个==编译时==的属性。

生命周期：一个变量的生命周期是指程序运行时变量存在的有效时间段，在此时间区域内它可以被程序的其他部分引用；是一个==运行时==的概念。

词法块：这些声明在代码中并未显式地使用花括号包裹起来。
句法块：是由花括弧所包含的一系列语句

作用域---包级
任何在函数外部（也就是包级语法域）声明的名字可以在同一个包的任何源文件中访问的。
作用域---源文件级
导入的fmt包，则是对应源文件级的作用域，因此只能在当前的文件中访问导入的fmt包，当前包的其它源文件无法访问在当前源文件导入的包。

编译器通过名字引用查询时：
当编译器遇到一个名字引用时，它会对其定义进行查找，查找过程从==最内层的词法域==向==全局的作用域==进行。如果查找失败，则报告“未声明的名字”这样的错误。如果该名字在内部和外部的块==分别声明过==，则==内部块==的声明首==先被找到==。在这种情况下，内部声明屏蔽了外部同名的声明，让==外部的声明==的名字==无法被访问==。


go的基础数据类型
- 整型
- 浮点数 float32,float64
- 复数 complex64和complex128
- 布尔型
- 字符串
- 常量


==整型== 分成了有符号和无符号类型，
有符号类型：int8,int16,int32,int64
无符号类型：uint8,uint16,uint32,uint64

Unicode字符==rune类型==是和==int32==等价的类型，通常用于表示一个Unicode码点。这两个名称可以互换使用。

同样==byte==也是==uint8==类型的等价类型，byte类型一般用于强调数值是一个原始的数据而不是一个小的整数。

一种无符号的整数类型==uintptr==，没有指定具体的bit大小但是足以容纳指针。uintptr类型只有在==底层编程时==才需要，特别是Go语言和C语言函数库或操作系统接口相交互的地方。

![在这里插入图片描述](https://img-blog.csdnimg.cn/20201231112228830.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L1Nyd2ljaQ==,size_16,color_FFFFFF,t_70)
二元运算符有五种优先级。在同一个优先级，使用左优先结合规则。

`特殊的说明：`
取模运算符%仅用于整数间的运算，在Go语言中，%取模运算符的符号和被取模数的符号总是一致的，因此-5%3和-5%-3结果都是-2。

![在这里插入图片描述](https://img-blog.csdnimg.cn/20201231112740966.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L1Nyd2ljaQ==,size_16,color_FFFFFF,t_70)
位操作运算符 ==^== 作为二元运算符时是按位异或（XOR），当用作一元运算符时表示按位取反；也就是说，它返回一个每个bit位都取反的数。

位操作运算符 ==&^== 用于按位置零（AND NOT）：如果对应y中bit位为1的话，表达式z = x &\^ y结果z的对应的bit位为0，否则z对应的bit位等于x相应的bit位的值。

数值转换
浮点数到整数的转换将丢失任何小数部分，然后向数轴零方向截断。

`fmt的两个使用技巧`
通常Printf格式化字符串包含多个%参数时将会包含对应相同数量的额外操作数，但是%之后的[1]副词告诉Printf函数再次使用第一个操作数。第二，%后的#副词告诉Printf在用%o、%x或%X输出时生成0、0x或0X前缀。
```go
fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
```

复数：
内置的complex函数用于构建复数，内建的real和imag函数分别返回复数的实部和虚部：


>var x complex128 = complex(1, 2) // 1+2i
var y complex128 = complex(3, 4) // 3+4i
fmt.Println(x*y)                 // "(-5+10i)"
fmt.Println(real(x*y))           // "-5"
fmt.Println(imag(x*y))           // "10"

如果一个浮点数或一个十进制整数后有i，例如3.141592i或2i，它将构成一个复数的虚部，复数的实部是0


#### 字符串类型

可以采用s[i:j]输出字符串。

不管i还是j都可能被忽略，当它们被忽略时将采用0作为开始位置，采用len(s)作为结束的位置。

字符串不可被修改，但可以重新被分配。因为字符串是不可修改的，因此尝试修改字符串内部数据的操作也是被禁止的。

#### 转义字符
![在这里插入图片描述](https://img-blog.csdnimg.cn/20201231151316294.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L1Nyd2ljaQ==,size_16,color_FFFFFF,t_70)
Go语言的源文件采用UTF8编码。且Go语言的range循环在处理字符串的时候，会自动隐式解码UTF8字符串。

#### 字符串和数字的互转
##### 整数转字符串
- 一种方法是用fmt.Sprintf返回一个格式化的字符串；
- 另一个方法是用strconv.Itoa(“整数到ASCII”)。

##### 字符串转整形
- 可以使用strconv包的Atoi或ParseInt函数，
- 还有用于解析无符号整数的ParseUint函数。


#### 常量
常量在编译期计算，每种常量的潜在类型都是基础类型。

在常量赋值给变量过程中，无类型整数常量转换为int，它的内存大小是不确定的，但是无类型浮点数和复数常量则转换为内存大小明确的float64和complex128。

### 复合数据类型
Go中的复合数据类型主要包括：
- 数组
- Slice
- Map
- 结构体
- JSON
- 文本和HTML模板

#### 数组
数组初始化
```go
var q [3]int = [3]int{1, 2, 3}
q := [...]int{1, 2, 3} //短声明,...省略号代表数组长度根据初始化个数来定义

var r [3]int = [3]int{1, 2} // r[2] == 0
```
数组长度必须是常量，不同的数组长度是不同的数据类型。

声明数组部分元素的方法：
```go
r := [...]int{99: -1} // 声明了r[99]的值为-1 ，其他值无声明默认为零值
```