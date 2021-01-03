
#### 数组
##### 数组初始化
```go
var q [3]int = [3]int{1, 2, 3}
q := [...]int{1, 2, 3} //短声明,...省略号代表数组长度根据初始化个数来定义
var r [3]int = [3]int{1, 2} // r[2] == 0
```
`注:数组长度必须是常量，不同的数组长度是不同的数据类型。`

声明数组部分元素的方法：
```go
r := [...]int{99: -1} // 声明了r[99]的值为-1 ，其他值无声明默认为零值
```

#### Slice
 1. 大slice中的slice片段可以超过本身的长度赋给新值，
==例子如下：==
```go
func main()  {
	day :=[] string {1:"Monday",2:"Tuesday",3:"Wednesday",4:"Thursday",5:"Friday",6:"Saturday",7:"Sunday"}
	D1 := day[1:4] //在day中抽出了前三天
	//fmt.Println(D1[6:]) //报错，超过了D1的长度了
	workDay := D1[:5] //超过了D1的长度，但仍然没超过day的长度，可以赋值。
	fmt.Println(workDay) //[Monday Tuesday Wednesday Thursday Friday]
}
```

2. - slice之间不可以直接使用==比较。一是slice的元素是间接引用的，二是因为一个slice值在不同时刻可能包括着不同的底层元素。真的要比较就要遍历所有元素进行比较。
	- slice只能和nil直接用\==进行比较。

3. 一个零值的slice等于nil,一个nil的slice是没有底层数组的，它的长度和容量都为0。

4. 可以使用make函数生成一个切片。

appendInt解析。

```go
//传入要加入数值的slice X ,和需要加入slice的值 y
func appendInt(x []int, y int) []int {
    var z []int
    //新切片 z 的长度
    zlen := len(x) + 1
    //当zlen比x的容量小的话，说明y可以直接放到x里，将x赋给z
    if zlen <= cap(x) {
        // 将x赋给z
        z = x[:zlen]
    } else {
        // 当x的容量不足以添加y时
        // 先把容量定为与长度相等，此时容量刚好为在 x后添加y
        zcap := zlen
        //当x的两倍长大于z的容量时，将z容量定义为x长度的两倍
        //目的是为了不用频繁变更长度，让插入保持在一个相对稳定的频次
        if zcap < 2*len(x) {
            zcap = 2 * len(x)
        }
        //将定义好的z切片、长度、容量，使用make函数生成一个切片
        z = make([]int, zlen, zcap)
        //将x元素复制进入z中，
        copy(z, x) // a built-in function; see text
    }
    //在z切片中已存在元素后面添加y，就此新建完成。
    z[len(x)] = y
    //返回添加y后的x切片内容的新切片
    return z
}
```

#### Map

