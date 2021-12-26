package main

import "fmt"

var ch chan int
func main() {
	ch=make(chan int,1)

	go func() {
		fmt.Println("下山的路又堵起了")
		ch<-1
	}()
	<-ch

}

