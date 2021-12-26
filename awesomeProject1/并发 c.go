package main

import (
	"fmt"
	"sync"
)

var ch1 chan int
var exit chan bool
var wg1 sync.WaitGroup
func main(){

	ch1=make(chan int)
	exit=make(chan bool)
	wg1.Add(10)
	go printf1()
	go printf2()
	go printf3()
	go printf4()
	go printf5()
	go printf6()
	go printf7()
	go printf8()
	go printf9()
	go printf10()

	<-exit

}
func printf1(){

		ch1<-1
		fmt.Println("明明是圣诞节")



}
func printf2(){

		<-ch1
		fmt.Println("不应该去配nyp玩吗")


}
func printf3(){

		ch1<-1
		fmt.Println("但是我还得参加后端考试")


}
func printf4(){

		<-ch1
		fmt.Println("考试还好难")



}
func printf5(){
		ch1<-1
		fmt.Println("嘤嘤嘤")

}
func printf6(){

		<-ch1
		fmt.Println("看来不能再继续摸鱼了")


}
func printf7(){
	ch1<-1
	fmt.Println("我放假得好好学学了")

}
func printf8(){
	<-ch1
	fmt.Println("算了")

}
func printf9(){
	ch1<-1
	fmt.Println("还是摸鱼香")

}
func printf10(){
	<-ch1
	fmt.Println("睡觉去了")

	exit<-true
}





