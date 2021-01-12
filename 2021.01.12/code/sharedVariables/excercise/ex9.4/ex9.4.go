package main

import "fmt"

func main() {
	var ch = make(chan int)
	var num = 0
	for{
		select {
			case ch<-num :
			 go func() {
				for {
					if num%2==0 {
						num ++
						ch <- num
					}
					num++
					fmt.Println(num)
				}
			}()
			case <- ch :
			go func() {
				for {
					if num%2==0 {
						num ++
						ch <- num
					}
					num++
					fmt.Println(num)
				}
			}()
		}
	}


}
