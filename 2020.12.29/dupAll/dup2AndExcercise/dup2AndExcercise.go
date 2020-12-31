package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func main()  {
	counts := make(map[string]int)
	file := os.Args[1:]
	if len(file) == 0 {
		countLines(os.Stdin,counts)
	}else{
		for _,args := range file{
			f,err := os.Open(args)
			if err != nil {
				fmt.Fprintf(os.Stderr,"dup2:%v\n",err)
				continue
			}
			countLines(f,counts)
			f.Close()
		}
	}
	for line,n := range counts {
		if n>1 {
			funcName,file,lines,ok := runtime.Caller(0)
			if ok {
				//打印出文件名
				fmt.Println("func name: " + runtime.FuncForPC(funcName).Name())
				fmt.Printf("file: %s, line: %d\n",file,lines)
			}

			fmt.Printf("%d\t%s\n",n,line)
		}
	}

}
func countLines(f *os.File,counts map[string]int)  {
	input := bufio.NewScanner(f)
	for  input.Scan() {
		counts[input.Text()]++
		if input.Text()=="end" {
			break
		}
	}
}
