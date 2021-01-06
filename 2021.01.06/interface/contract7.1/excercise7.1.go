//练习 7.1： 使用来自ByteCounter的思路，实现一个针对对单词和行数的计数器。你会发现bufio.ScanWords非常的有用。
package main

import (
	"bufio"
	"fmt"
	"os"
)

type WordCounter int

func (c *WordCounter) Write(input *bufio.Scanner)  {
	input.Split(bufio.ScanWords)
	for input.Scan(){
		line := input.Text()
		*c++
		if line=="end" {
			*c--
			break
		}
	}
	fmt.Println(*c)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	var c WordCounter
	c.Write(input)
}