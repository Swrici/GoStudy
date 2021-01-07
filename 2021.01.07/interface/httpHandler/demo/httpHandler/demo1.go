//初步使用http.Handler接口
//使用了一个将库存商品价格映射成美元的demo
package main

import (
	"fmt"
	"log"
	"net/http"
)
//定义了一个dollars结构体 类型是float32
type dollars float32

//重写了dollars的String(),输出更方便查看
func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

//定义了一个database结构体 类型是key为String，value为dollars结构体的map
type database map[string]dollars
func main() {
	//db := database{"shoes": 50, "socks": 5}
	////ListenAndServe需要例如“localhost:8000”的服务器地址，和一个所有请求都可以分派的Handler接口实例
	//log.Fatal(http.ListenAndServe("localhost:8000", db))

	//但在一个实际的应用中，将每个case中的逻辑定义到一个分开的方法或函数中会很实用
	db := database{"shoes": 50, "socks": 5}

	//ServeMux来简化URL和handlers的联系。一个ServeMux将一批http.Handler聚集到一个单一的http.Handler中
	mux := http.NewServeMux()
	//
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}



//实现了ServerHTTP方法，也就是实现了http.Handler接口
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	//输出给访问的客户端
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	//进入price路径
	case "/price":
		//item是通过get方法获得的参数名称，将获得的参数值在db这个map里进行查询
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	//差不到返回 no such page 提示信息
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}
