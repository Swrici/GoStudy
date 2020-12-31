// Echo1 prints its command-line arguments.Demo1
package main

import (
	"fmt"
	"os"
)

func main()  {
	var s,step string
	for i:=1;i<len(os.Args);i++{
		s += step + os.Args[i]
		step = " "
	}
	fmt.Println(s)
}
