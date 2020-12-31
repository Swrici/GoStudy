package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {
	count:=make(map[string]int)
	
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		count[input.Text()]++
		fmt.Println(input.Text())
		if input.Text() == "end" { break }
	}

	for line,n:=range count{
		if n>1 {
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}