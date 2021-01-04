package main

import "fmt"

func main() {
	//1.1 结构体字面值声明结构体
	type pointer struct {
		X,Y int

	}
	type stu struct {
		id int
		name string
		pointer
	}
	newPointer1 := pointer{6,6}
	fmt.Println(newPointer1)
	//1.2 以成员名字和相应的值来初始化声明结构体
	newPointer2 := pointer{X:1} //这种情况下，无提及到的变量默认为零值
	fmt.Println(newPointer2)

	//1.3 指针方式创建及初始化一个结构体变量
	newPointer3 := &pointer{6,6}
	//等价于 newPointer3 :=new(pointer)    *newPointer3 = Point{1,2}
	fmt.Println(*newPointer3)

	//2 结构体的比较
	//  当结构体内的属性都是可以比较的，那么结构体都是可以比较的,
	//  且可以比较也意味着可以用于map中的key
	fmt.Printf("newPointer1==newPointer3?%t\n",newPointer1==*newPointer3)

	//3 匿名成员和嵌入机制
	var stu1 stu
	stu1.name = "张白"
	stu1.id = 1
	stu1.X = 2
	stu2 := stu{1,"李千",pointer{2,1}}
	fmt.Println(stu1,stu2)
	fmt.Printf("%#v\n%#v\n",stu1,stu2) //在里面的#使得打印的时候也输出了对应的属性名
}
