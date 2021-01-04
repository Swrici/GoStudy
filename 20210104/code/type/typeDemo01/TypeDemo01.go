//测试
//如果将studentByID函数的返回值从*student指针类型改为Employee值类型，那么更新语句将不能编译通过，因为在赋值语句的左边并不确定是一个变量
package main

import "fmt"

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