//练习5.18：不修改fetch的行为，重写fetch函数，要求使用defer机制关闭文件。
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)
// Fetch downloads the URL and returns the
// name and length of the local file.
func main() {
	for _,url := range os.Args[1:]{
		resp,err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr,"getErr:%v\n",err)
			os.Exit(1)
		}
		//b,err := ioutil.ReadAll(resp.Body)
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr,"ReadAllErr:%v\n",err)
			os.Exit(1)
		}
		fmt.Printf("%v",os.Stdout)
	}
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr :=  f.Close(); err == nil{
		err = closeErr
	}
	defer f.Close()

	return local, n, err
}
