//test range
package main

import "fmt"

func main() {
	x := "hello"
	for  _,x := range x {
		x := x + 'A' - 'a'
		fmt.Printf(" %c", x) // "HELLO" (one letter per iteration)
	}
	fmt.Println()
	for  i,x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%q  %c", i,x) //'\x00'  H'\x01'  E'\x02'  L'\x03'  L'\x04'  O
	}
}