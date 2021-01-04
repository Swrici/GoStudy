// 修改issues程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年。
//使用 Issues repo:golang/go is:open json decoder 在终端进行搜索
package main

import (
	"com/swrici/dataType/complexType/json/excercise/excercise4.10/github"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	//调用SearchIssues 获得搜索结果集
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("总共有 %d 个issues:\n", result.TotalCount)
	//三个切片用于分类分别是 一月以内的 一年以内的 超过一年的
	var monthSlice []*github.Issue
	var yearSlice []*github.Issue
	var moreYearSlice []*github.Issue
	//对搜索结果集进行按时间分类
	for _,item :=range result.Items{
		//times = append(times, item.CreatedAt)
		if year,month,day :=item.CreatedAt.Date();year==2020&&month==time.Month(12)&&day>4{
			monthSlice = append(monthSlice,item)
		}else if year>2019&&month>time.Month(1)&&day>4{
			yearSlice = append(yearSlice,item)
		}else {
			moreYearSlice = append(moreYearSlice,item)
		}
	}
	//分别展示三个分类的结果
	//一月内的
	fmt.Println()
	fmt.Printf("一个月内的有 %d issues:\n", len(monthSlice))
	for _,item := range monthSlice{
		//fmt.Printf(*item.)
		fmt.Printf("#%-5d %9.9s %.55s  %v\n",
			item.Number, item.User.Login, item.Title,item.CreatedAt)
	}
	//一年内的
	fmt.Println()
	fmt.Printf("一年内的有 %d issues:\n", len(yearSlice))
	for _,item := range yearSlice{
		//fmt.Printf(*item.)
		fmt.Printf("#%-5d %9.9s %.55s  %v\n",
			item.Number, item.User.Login, item.Title,item.CreatedAt)
	}
	//超过一年的
	fmt.Println()
	fmt.Printf("超过一年的有 %d issues:\n", len(moreYearSlice))
	for _,item := range moreYearSlice{
		//fmt.Printf(*item.)
		fmt.Printf("#%-5d %9.9s %.55s  %v\n",
			item.Number, item.User.Login, item.Title,item.CreatedAt)
	}
}