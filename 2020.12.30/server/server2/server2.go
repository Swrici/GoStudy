// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer,"%s %s %s\n",request.Method,request.URL,request.Proto)
	for k,v := range request.Header{
		fmt.Fprintf(writer,"Header[%q]=%q\n",k,v)
	}
	fmt.Fprintf(writer,"Host = %q\n",request.Host)
	fmt.Fprintf(writer,"RemoteAddr = %q\n",request.RemoteAddr)
	if err := request.ParseForm();err!=nil {
		log.Print(err)
	}
	for k,v := range request.Form{
		fmt.Fprintf(writer,"Form[%q]=%q",k,v)
	}
}

func counter(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	fmt.Fprintf(writer, "Count的值为： %d\n", count)
	mu.Unlock()
}


