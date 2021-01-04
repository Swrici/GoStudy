//编写一个程序 word freq 程序，报告输入文本中每个单词出现的频率。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	wordCountFunc()
}

func wordCountFunc()  {
	wordCount := make(map[string]int)
	//开启输入流
	input := bufio.NewScanner(os.Stdin)
	//保证输入的时候是按单词计数不是按行计数
	input.Split(bufio.ScanWords)
	for input.Scan(){
		line := input.Text()
		wordCount[line]++
		if line=="end" {
			break
		}
	}
	delete(wordCount,"end")
	fmt.Println(wordCount)
}


