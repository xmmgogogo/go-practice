package main

import (
	"fmt"
	"time"
)

func main() {
	// 统计总和

	var (
		ch    = make(chan string, 3)
		items = []string{"ap", "ad", "dd", "d2"}
	)
	go producer(items, ch)
	consumer(ch)
}

func producer(items []string, ch chan string) {
	for _, v := range items {
		ch <- v
		time.Sleep(time.Second * 1)
	}
	close(ch)
}

func consumer(ch chan string) {
	// 为什么读完不会阻塞？？？
	//for _, v := range <-ch {
	//	fmt.Println("获取的内容是:", string(v))
	//}

	for {
		if word, isOk := <-ch; isOk {
			fmt.Println("获取的内容是:", word)
			//time.Sleep(time.Millisecond * 100)
		} else {
			fmt.Println("读取结束")
			break
		}
	}
}
