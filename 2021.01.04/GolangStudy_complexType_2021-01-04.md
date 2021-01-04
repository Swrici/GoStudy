
@[TOC](目录)
## 复杂类型
### Map
- 无序，key各不相同，能在常数时间复杂度里检索、更新或删除对应的value。
- 所有key的类型相同、所有value的类型相同。
- **key必须是支持\=\=的数据类型**（例如slice就不支持\==比较），而map可以通过key是否相等来判定是否已经存在。
- 不能对map中的元素进行取址。一是由于Map中的元素不是一个变量，二是因为map可能随着元素的增多重新分配更大的内存空间，导致之前的地址无效。
- map的直接遍历是无序的。
- map类型的零值是nil，也就是没有引用任何哈希表。
- nil值的map可以安全进行大多数的操作，包括查找，删除、len和range，但是当插入的时候会报panic异常。

### 结构体
- 结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体。
- 结构体变量的成员可以通过点操作符访问。
- 一个结构体可能同时包含导出和未导出的成员。

语法：
```go
type typename struct{
	attributeName:attributeType
	...
}

```

**调用函数返回的是值，并不是一个可取地址的变量**。

如果要**在函数内部修改结构体成员的话，用指针传入是必须的**；因为在Go语言中，所有的**函数参数都是值拷贝传入**的，函数参数将不再是函数调用时的原始变量。
```go
type student struct {
	stuId int
	name string
	age int
}
var stus = make([]student,3)
func main()  {
	var stu1 student
	stu1.stuId = 1
	stu1.name = "李四"
	stu1.age = 12
	stus[0] = stu1

	stuNew := getStuById(1)
	fmt.Println(stuNew)
	//getStuById(1).name = "李五"  //报错 ：Cannot assign to getStuById(1).name

}

func getStuById(id int) student  {
	for _,stu :=range  stus {
		if stu.stuId == id {
			return stu
		}
	}
	return student{}
}
```

使用结构体声明变量
```go
func main() {
	//1.1 结构体字面值声明结构体
	type pointer struct {
		X,Y int
	}
	newPointer1 := pointer{6,6}
	fmt.Println(newPointer1)
	//1.2 以成员名字和相应的值来初始化声明结构体
	newPointer2 := pointer{X:1} //这种情况下，无提及到的变量默认为零值
	fmt.Println(newPointer2)
}

```

Go的==匿名成员==
在结构体中，我们**只声明一个成员对应的数据类型而不指名成员的名字**；这类成员就叫**匿名成员**。

匿名成员的==特点==
数据类型必须是**命名的类型或指向一个命名的类型的指针**。

而当有别的结构体使用了这些匿名成员的时候，就叫**匿名成员类型嵌入了使用匿名成员类型的结构体**中。
而使用这种嵌入机制带来的好处就是，**在原先需要通过多次点操作才能到达的变量可以直接一次点操作获得**。

例子：
```go
type pointer struct{
	X,Y int
}
type circle struct{
	pointer
	radius int
}
type wheels struct{
	circle
	spokes int
}
var w Wheel
w.X = 8            // 等价于 w.Circle.Point.X = 8
w.Y = 8            // 等价于 w.Circle.Point.Y = 8
w.Radius = 5       // 等价于 w.Circle.Radius = 5
w.Spokes = 20
```
但是匿名成员的简化声明**不能在结构体的字面值声明中生效**。
```go
w = Wheel{8, 8, 5, 20}                       // 编译错误: 未知的属性
w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} // 编译错误: 未知的属性
//字面值声明应该如下
w = Wheel{Circle{Point{8, 8}, 5}, 20}
//或者
w = Wheel{
    Circle: Circle{
        Point:  Point{X: 8, Y: 8},
        Radius: 5,
    },
    Spokes: 20, // 此处（和半径处）需要加上逗号，最后插入的逗号不会导致编译错误
    //加逗号是为了防止编译器在行尾自动插入分号导致的编译错误，所有在末尾的参数变量后面显式加逗号。

}
```
在使用匿名成员时的==注意点==
- 不能同时包含两个类型相同的匿名成员。
- 任何命名的类型都可以作为结构体的匿名成员。目的不止使用该类型，还有该类型导出的方法集。

### JSON
JSON(JavaScript对象表示法)是一种**用于发送和接收结构化信息的标准协议**。

#### 编组
将一个Go语言中的**结构体slice转为JSON的过程叫编组**（marshaling）。通过调用**json.Marshal**函数完成。该函数带的参数是需要转化的结构体变量。返还值是一个编码后的字节slice，包含很长的字符串，并且没有空白缩进；

又因为Marshal函数生成的不便于阅读。所以使用**json.MarshalIndent**函数来生成美观便于阅读的json。

`在编码时，默认使用Go语言结构体的成员名字作为JSON的对象（通过reflect反射技术）。只有导出的结构体成员才会被编码`

结构体中的==Tag==
结构体的field属性后可以加上对应的**Tag**,它的作用就是**起别名**，将原名（fieldName）改成想要的名字。而Tag里面还有一个额外的omitempty选项，表示当Go语言结构体成员为**空或零值**时不生成json对象。

#### 解码
解码是编组的逆操作，也就是将JSON数据解码为Go语言的数据结构，通过json.Unmarshal函数完成。

在接受json数据的时候可以选择性的接受。通过**json.Unmarshal**函数完成。

### 文本和HTML模板
有时我们会需要复杂的打印格式，这些功能是由**text/template和html/template等模板包**提供的，它们**提供了一个将变量值填充到一个文本或HTML格式的模板的机制**。

一个模板是一个字符串或一个文件，里面包含了一个或多个由双花括号包含的{{action}}对象。

在一个action中，|操作符表示将前一个表达式的结果作为后一个函数的输入，

>action中的函数:printf函数，是一个基于fmt.Sprintf实现的内置函数。

#### 如何生成模板？
生成模板的输出需要两个处理步骤。
- 第一步是要分析模板并转为内部表示，分析模板部分一般只需要执行一次。
- 第二步是基于指定的输入执行模板。

## 5 函数

### 5.1 函数声明
函数声明包括**函数名、形式参数列表、返回值列表（可省略）以及函数体**。

`注：如果一组形参或返回值有相同的类型，我们不必为每个形参都写出参数类型`

如果**两个函数形式参数列表和返回值列表中的变量类型一一对应**，那么这**两个函数被认为有相同的类型和标识符**。形参和返回值的变量名不影响函数标识符也不影响它们是否可以以省略参数类型的形式表示。

`注：函数的类型被称为函数的标识符。`
>func add(x int, y int) int   {return x + y}
func sub(x, y int) (z int)   { z = x - y; return}
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

>fmt.Printf("%T\n", add)   // "func(int, int) int"
fmt.Printf("%T\n", sub)   // "func(int, int) int"
fmt.Printf("%T\n", first) // "func(int, int) int"
fmt.Printf("%T\n", zero)  // "func(int, int) int"

实参通过**值的方式传递**，因此函数的形参是实参的拷贝。对形参进行修改不会影响实参。但是，如果实参包括引用类型，如指针，slice(切片)、map、function、channel等类型，实参可能会由于函数的间接引用被修改。

### 5.2 递归
