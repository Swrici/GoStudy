//很多图形界面提供了一个有状态的多重排序表格插件：主要的排序键是最近一次点击过列头的列，第二个排序键是第二最近点击过列头的列，等等。
//定义一个sort.Interface的实现用在这样的表格中。比较这个实现方式和重复使用sort.Stable来排序的方式。
package main

import (
	"fmt"
	"sort"
	"time"
)
type t time.Time
type byTime []t
func (x byTime) Len() int           { return len(x) }
func (x byTime) Less(i, j int) bool {
	timeString1 := time.Time(x[i]).Format("2006-01-02 15:04:05")
	timeString2 := time.Time(x[j]).Format("2006-01-02 15:04:05")

	t1, err := time.Parse("2006-01-02 15:04:05",timeString1 )
	t2, err := time.Parse("2006-01-02 15:04:05",timeString2 )
	//为了展示，特别修改了让他倒序显示
	return	err==nil &&t2.Before(t1)
}

func (x byTime) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {
	t1 := t(time.Now())

	time.Sleep(1000000000)
	t2 := t(time.Now())
	time.Sleep(1000000000)
	t3 := t(time.Now())
	var times = []t{t1,t2,t3}
	for _,item := range times{
		fmt.Println(time.Time(item))
	}
	fmt.Println("================即将进行sort.Interface显示排序后的================")
	t12 := time.Now()
	sort.Sort(byTime(times))
	t23 := time.Since(time.Time(t12))
	fmt.Println(t23)
	for _,item := range times{
		fmt.Println(time.Time(item))
	}
	fmt.Println("==================即将进行sort.Stable显示排序后的======================")
	sort.Stable(byTime(times))
	for _,item := range times{
		fmt.Println(time.Time(item))
	}
}