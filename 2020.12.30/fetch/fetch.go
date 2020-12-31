package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main()  {
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