package main

import (
	"fmt"
	"sync"
)
var mu sync.Mutex
var wg sync.WaitGroup
func main() {
	wg.Add(1)


	go func() {
		fmt.Println("有点强人锁男")
		mu.Lock()
		wg.Done()
		mu.Unlock()
	}()//出现了死锁的问题,并且主协程不知道子协程是否运行完了，并且主协程里的语句执行速度快过其他gorouting，主协程先退出，其他gorouting不会继续执行

	wg.Wait()

}

