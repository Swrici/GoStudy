//map的初学习
package main

import (
	"fmt"
	"sort"
)

func main()  {
	//1. map的生成
	fmt.Println("========== map的生成==========")
	map1 := make(map[string]int) //创建了一个key为字符串类型，value为int类型的map变量
	map2 := map[string]int{
		"first":1,
		"second":2,
		"nice":666,
	}
	map3 := make(map[string]int)
	map3["first"] = 1
	map3["second"] = 2
	map4 := map[string]int{}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3) //map3和map2生成方法等效
	fmt.Println(map4)

	//2. 访问map的value
	fmt.Println("==========访问map的value==========")
	fmt.Println(map2["first"]) //打印 1

	//3.增加map
	fmt.Println("===========增加map=========")
	map2["third"] = 3
	fmt.Println(map2["third"])
	//4.删除map
	fmt.Println("==========删除map==========")
	delete(map2,"third")
	fmt.Println(map2["third"]) //查无对应的key，返回value类型的零值

	//5.map类型的运算
	fmt.Println("==========map类型的运算==========")
	map2["second"] += map2["first"]
	fmt.Println(map2["second"])
	map2["second"]--
	fmt.Println(map2["second"])

	//6.map类型的遍历
	fmt.Println("==========map类型的遍历==========")
	for name,value := range map2 {
		fmt.Printf("%s :%d\n",name,value)
	}

	//6.map类型根据key排序遍历
	fmt.Println("==========map类型根据key排序遍历==========")
	var names = []string{"first","third","nice","second"}
	map2["third"] = 3
	//sorts方法对key进行排序
	sort.Strings(names)
	for _,name := range names {
		fmt.Printf("%s :%d\n",name,map2[name])
	}

	//7.nil值map进行插入操作
	fmt.Println("==========nil值map进行插入操作==========")
	var map5 map[string]int
	fmt.Printf("map5是否为nil呢？%t\n",map5==nil)
	//map5["first"] = 1 //panic: assignment to entry in nil map
	fmt.Println(map5)

	//8.判断map中是否存在一个value和零值相等的key
	fmt.Println("==========判断map中是否存在一个value和零值相等的key==========")
	if value,ok := map2["nice"];!ok {
		fmt.Printf("%d:%t\n",value,ok)
	}else {
		fmt.Printf("%d:%t\n",value,ok)
	}
}

