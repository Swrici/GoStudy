//将结构体输出成json
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type stu struct {
	Id int `json:"jsonId"` //别名修改
	Name string  `json:"jsonName,omitempty"`
	books
}

type books struct {
	Math string
	Chinese string
}


func main() {
	var stus = []stu{
		stu{1,"林1",books{"高树","大禹"}},
		stu{2,"林2",books{"中树","中禹"}},
		stu{3,"",books{"低树","低禹"}},
	}
	fmt.Println(stus)
	// 编组
	//""是前缀是方括号（不包含第一个，保括最后一个方括号）内每行第一个字符，indent是缩进字符换行使用
	str,err := json.MarshalIndent(stus,"","	")//MarshalIndent美观
	//str,err := json.Marshal(stus)//Marshal不美观
	if err!=nil {
		log.Fatalf("错误信息是：%s\n",err)
	}
	fmt.Printf("%s\n",str)
	// 解码 ，可以接受自己想接受的数据
	var book []struct{Math string}
	if err := json.Unmarshal(str,&book);err!=nil {
		log.Fatalf("%s/n",err)
	}
	fmt.Println(book)
}