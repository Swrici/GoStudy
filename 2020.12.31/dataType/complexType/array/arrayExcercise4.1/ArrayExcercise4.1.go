package main


import (
	"crypto/sha256"
	"fmt"
)

func main() {

	b1 :=sha256.Sum256([]byte("a"))
	b2 :=sha256.Sum256([]byte("A"))
	fmt.Printf("%x\n%x\n%t\n%T\n", b1, b2, b1 == b2, b1)
	fmt.Println(PopCount(b1,b2))
}

func PopCount(b1 [32]uint8, b2 [32]uint8) int{
	num := 0
	//for i := 0;i<32;i++ {
	//	fmt.Printf("%x---%x %t\n", b1[i], b2[i], b1 == b2)
	//	//b3 := b1[i]&b2[i]
	//	//fmt.Printf("%d--%x\n",i, uint8(b3))
	//	if b1[i]==b2[i] {
	//		continue
	//	}
	//	num++
	//}
	//for i:=31;i>0;i++{
	//	return int(byte(b1>>(1*8)))
	//}

	return num
}


