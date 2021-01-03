//sliceDemo to practice knowledge learning
package main

import (
	"fmt"
)

func main()  {
	day :=[] string {1:"Monday",2:"Tuesday",3:"Wednesday",4:"Thursday",5:"Friday",6:"Saturday",7:"Sunday"}
	D1 := day[1:4]
	//fmt.Println(D1[6:]) //报错
	fmt.Println(D1[1:3]) //报错
	workDay := D1[:5]
	fmt.Println(workDay)
	//make创建slice
	fmt.Println("=======make创建slice=======")
	intSlice1 := make([]int, 5)//make([]T, len)
	intSlice2 := make([]int,5,10)//make([]T, len, cap)
	intSlice3 := make([]int,10)[1:5]//make([]T, cap)[:len]
	fmt.Println(intSlice1)
	fmt.Println(intSlice2)
	fmt.Println(intSlice3)
	//append函数的使用
	fmt.Println("=======append函数的使用=======")
	var slogan []rune
	for _,r:=range "Hello,append()~"{
		slogan = append(slogan,r)
	}
	fmt.Println(string(slogan))
}
