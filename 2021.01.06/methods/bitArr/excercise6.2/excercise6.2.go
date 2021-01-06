//练习 6.2： 定义一个变参方法(*IntSet).AddAll(...int)，这个方法可以添加一组IntSet，比如s.AddAll(1,2,3)
package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

//变参方法(*IntSet).AddAll(...int)，这个方法可以添加一组IntSet，比如s.AddAll(1,2,3)
func (s *IntSet) AddAll(x...int) {
	for _,val := range x{
		word, bit := val/64, uint(val%64)
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
	}
}

func main()  {
	var x IntSet
	x.AddAll(1,2,3)
	fmt.Println(x.String())
}

//下面的不是手写的代码
// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				//fmt.Println(word&(1<<uint(j)),s.words[word&(1<<uint(j))])
				fmt.Fprintf(&buf, "%d", 64*i+j)

			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}