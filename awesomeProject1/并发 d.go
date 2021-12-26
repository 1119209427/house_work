package main

func putNum(intChan chan int){
	for i:=1;i<=1000000;i++{
		intChan<-i
	}
	close(intChan)
}
func primeNum(intChan chan int,primeChan chan int,exitChan chan bool) {
	var flag bool
	for {

		num, ok := <-intChan
		if !ok {
			break
		}

		flag = true//假设是素数
		//判断是否为素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag{
			primeChan<-num
		}
	}
	exitChan<-true
}
func main(){
	intChan :=make(chan int,1000000)
	primeChan:=make(chan int,500000)//放入结果
	extChan:=make(chan bool,8)//退出的管道

	//开启协程，向intchan放入数据
	go putNum(intChan)
	//从intchan取出，并判断是否是素数
	for i:=0;i<8;i++{
		go primeNum(intChan,primeChan,extChan)
	}
	go func(){
		for i:=0;i<8;i++{
			<-extChan
		}
		close(primeChan)
	}()
	for{
		_,ok:=<-primeChan
		if !ok{
			break
		}
		//fmt.Printf("素数=%d",)
	}
}
