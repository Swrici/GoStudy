//编写类似sum的可变参数函数max和min。考虑不传参时，max和min该如何处理，再编写至少接收1个参数的版本。
package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(max(1,2,3,4,5,-1))
	fmt.Println(min(1,2,3,4,5,-1))
}

func max(vals... int) int {
	if vals==nil {
		return math.MinInt8
	}else {
		var temp = math.MinInt8
		for _,max := range vals {
			if max>temp {
				temp = max
			}
		}
		return temp
	}
}

func min(vals... int) int {
	if vals==nil {
		return math.MaxInt8
	}else {
		var temp = math.MaxInt8
		for _,min := range vals {
			if min<temp {
				temp = min
			}
		}
		return temp
	}
}

