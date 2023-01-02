package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//debug.SetGCPercent(-1)

	dataChan := make(chan int, 1)
	call(dataChan)

	fmt.Println("读取到返回值：", <-dataChan)
}

func newGoroutine(dataChan chan int) {
	defer fmt.Println("退出后线程数量->", runtime.NumGoroutine())

	time.Sleep(time.Second * 100)
	dataChan <- 1
	return
}

func call(dataChan chan int) {
	defer func() { fmt.Println("退出线程数量->", runtime.NumGoroutine()) }()
	fmt.Println("当前线程数量->", runtime.NumGoroutine())
	go newGoroutine(dataChan)
	fmt.Println("当前线程数量->", runtime.NumGoroutine())
	go newGoroutine(dataChan)
	fmt.Println("当前线程数量->", runtime.NumGoroutine())
	go newGoroutine(dataChan)
	fmt.Println("当前线程数量->", runtime.NumGoroutine())

	//if time.Now().Hour() >= 17 {
	//	fmt.Println("提前退出：", time.Now().Hour())
	//	return
	//}
}
