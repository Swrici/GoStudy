//练习 7.11： 增加额外的handler让客服端可以创建，读取，更新和删除数据库记录。
//例如，一个形如 /update?item=socks&price=6 的请求会更新库存清单里一个货品的价格并且当这个货品不存在或价格无效时返回一个错误值。
//（注意：这个修改会引入变量同时更新的问题）
package main

import (
	//"encoding/json"
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
	mux.Handle("/update", http.HandlerFunc(db.update))
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
//一个形如 /update?item=socks&price=6 的请求会更新库存清单里一个货品的价格并且当这个货品不存在或价格无效时返回一个错误值。
//要将特殊字符转换? & 转换完才能继续。
func (db database) update(w http.ResponseWriter, req *http.Request) {
	//u,err := url.Parse((req.URL).String())
	//if err != nil {
	//	fmt.Println()
	//}
	//fmt.Println("一次全取（map格式）: %v", u.Query())
	//req.ParseForm()
	//query := req.URL.Query()
	//for _,f := range query{
	//	fmt.Println(f)
	//}


	//vars := req.URL.Query();
	//a, ok := vars["item"]
	//if !ok {
	//	fmt.Printf("param a does not exist\n");
	//} else {
	//	fmt.Printf("param a value is [%s]\n", a);
	//}

	//u ,err :=url.Parse((req.URL).String())
	//if err != nil {
	//	panic(err)
	//}

	//m, _ := url.ParseQuery(u.RawQuery)
	//fmt.Println(m["price"][0])
	item := req.URL.Query().Get("item")
	//query  := req.URL.Query()
	////req.Form
	//priceUrl := query.Get("price")
	//req.ParseForm()
	//if len(req.Form) > 0 {
	//	for k, v := range req.Form {
	//		if k=="price" {
	//
	//			priceUrl = v[0]
	//			fmt.Println("循环打印了price:",priceUrl)
	//		}
	//
	//	}
	//}

	//

	price, ok := db[item]
	//fmt.Println("打印了price:",price)
	if !ok  {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	fmt.Fprintf(w, "%s\n", price)


}




