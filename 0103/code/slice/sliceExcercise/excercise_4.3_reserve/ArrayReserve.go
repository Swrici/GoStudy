//反转函数
package main

import "fmt"

func main() {
	s := []int{5, 6, 7, 8, 9}
	point := &s
	fmt.Println(reserve(point,5)) // "[5 6 9 8]
}

func reserve(s *[]int, lens int) *[]int {
	ints := make([]int, lens)
	for i,j :=range *s  {
		if j == 0 {
			return &ints
		}
		ints[len(ints)-i-1] = j
	}
	return &ints;
}
