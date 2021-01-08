package main

import "fmt"

func main() {
	ints := make(chan int)
	var num = 6
	ints <- num

	 chanNum := <-ints
	fmt.Println(chanNum) //fatal error: all goroutines are asleep - deadlock!
}
