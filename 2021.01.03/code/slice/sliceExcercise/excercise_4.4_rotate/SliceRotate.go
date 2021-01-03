//按下标移置首位函数
package main

import "fmt"

func main(){
	s := []int{5, 6, 7, 8, 9}

	fmt.Println(rotate(s,2))
}

func rotate(s []int, index int) []int{
	var rotate = make([]int,len(s))
	for i:=0;i<len(s) ;i++  {
		mod := (index+i)%5
		rotate[i] = s[mod]
	}
	return rotate
}
