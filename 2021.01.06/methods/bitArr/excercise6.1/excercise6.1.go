//练习6.1: 为bit数组实现下面方法
//func (*IntSet) Len() int      // return the number of elements

package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}


// 返回s中有长度有多少
func (s *IntSet) Len() int {
	var len = 0
	for _,item := range s.words {
		if item == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if item&(1<<uint(j)) != 0 {
				len++
			}
		}
	}
	return len
}



func main()  {
	var x IntSet
	x.Add(1)
	x.Add(2)
	x.Add(214)
	x.Add(2141)
	fmt.Println(x.String())
	fmt.Println(x.Len())
	fmt.Println(x.String())
}

//下面的不是手写的代码
// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

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