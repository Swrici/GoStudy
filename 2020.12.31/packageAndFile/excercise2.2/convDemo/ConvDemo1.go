//This demo is to test excercise2.2
package main

import (
	"com/swrici/packageAndFile/excercise2.2/lengthConversion"
	"fmt"
	"os"
	"strconv"
)

func main()  {
	for _,args := range os.Args[1:]{
		num,err :=strconv.ParseFloat(args,64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		cm := lengthConversion.Centimeter(num)
		m := lengthConversion.Metre(num)
		fmt.Printf("%s=%s   %s==%s",cm,lengthConversion.CToM(cm),m,lengthConversion.MToC(m))
	}
}