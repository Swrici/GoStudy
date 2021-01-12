//Memo实例会记录需要缓存的函数f（类型为Func），以及缓存内容（里面是一个string到result映射的map）。
//每一个result都是简单的函数返回的值对儿——一个值和一个错误值。
package main

import (
	"fmt"
	"log"
	"time"
)

// A Memo 备忘录缓存调用Func的结果，result是重新定义的结构体
type Memo struct {
	f     Func
	cache map[string]result
}

// Func 的类型是带字符串key的func类型
type Func func(key string) (interface{}, error)


type result struct {
	value interface{}
	err   error
}
//参数New，形参是Func,Func是自定义的函数方法
func New(f Func) *Memo {
	//在传入的指针和缓存传入值
	return &Memo{f: f, cache: make(map[string]result)}
}

// 注意：不是并发安全的！
func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}
func main()  {
	m := New(httpGetBody)
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}