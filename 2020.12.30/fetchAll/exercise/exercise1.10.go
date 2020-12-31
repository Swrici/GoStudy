package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func fetch(url string,ch chan<-string)  {
	//for i := 0;i<2;i++ {
	start := time.Now();
	resp,err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	//rbyte,err := io.Copy(ioutil.Discard,resp.Body)
	rbyte,err := io.Copy(os.Stderr,resp.Body)

	if err != nil {
		ch <- fmt.Sprint(err)
	}
	sec := time.Since(start).Seconds()
	ch <- fmt.Sprintln(sec,rbyte,url)
	//}

}

func main()  {
	start := time.Now()
	ch := make(chan string)
	ch1 := make(chan string)
	for _,url := range os.Args[1:] {
		go fetch(url,ch)
		go fetch(url,ch1)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
		fmt.Println(<-ch1)
	}
	fmt.Printf("%.2fs elapsed\n",time.Since(start).Seconds())
}
