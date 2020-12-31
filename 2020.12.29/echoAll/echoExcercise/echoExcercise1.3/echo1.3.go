package echoExcercise1_3

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

// BenchmarkIsMoreQuicker 对 fmt.echo 函数进行基准测试
func BenchmarkIsMoreQuicker1(b *testing.B)  {
	var s,step string
	//b.ResetTimer()
	//for i:=0;i<b.N;i++ {
	start := time.Now()
		for i:=1;i<len(os.Args);i++{
			s += step + os.Args[i]
			step = " "
		}
	end := time.Since(start)
	fmt.Println(end)
	//}

}

func BenchmarkIsMoreQuicker2(b *testing.B) {
	//b.ResetTimer()
	//start := time.Now();
	for i:=0;i<b.N;i++ {
		strings.Join(os.Args[1:], " ")
	}
}



