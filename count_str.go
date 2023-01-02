package main

import (
	"fmt"
	"time"
)

var (
	inputL = make(chan string, 3)
	//outPutL = make(chan string, 2)
)

func main() {
	// 计算一组字符里面，某个字母的数量
	list := []string{"app", "apq", "恭喜a", "恭喜b"}

	//fmt.Println(count1(list))
	//sliceChar()
	fmt.Println(count2(list))

	//for {
	//	time.Sleep(time.Millisecond * 500)
	//}
}

func count2(items []string) (res map[string]int) {
	res = make(map[string]int)

	//maxCount := len(items)
	go func() {
		for _, v := range items {
			inputL <- v
			fmt.Println("写入：", v)
		}
		close(inputL)
	}()

	// 异步等待的程序
	for {
		// 这一行很重要，如果通道关闭了，则直接退出for循环即可
		if getOne, isOk := <-inputL; isOk {
			//go func() {
			fmt.Println("读取到一个值", getOne)
			for _, v := range getOne {
				res[string(v)]++
			}
			time.Sleep(time.Millisecond * 1000)
			//}()
		} else {
			break
		}
	}

	return
}

func count1(items []string) map[string]int {
	var countMap = make(map[string]int)

	for _, v := range items {
		for _, v1 := range v {
			countMap[string(v1)]++
		}
	}

	return countMap
}

func sliceChar() {
	word := "abc"
	fmt.Println(word[:1]) // a

	word2 := "中国abc"
	fmt.Println(word2[:1]) // a
	fmt.Println(word2[:3]) // a
	fmt.Println(word2[:5]) // a
	fmt.Println(word2[:6]) // a

	word3 := []rune(word2)
	fmt.Println(string(word3[:1])) // a
	fmt.Println(string(word3[:2])) // a
	fmt.Println(string(word3[:3])) // a
	fmt.Println(string(word3[:4])) // a
	fmt.Println(string(word3[:5])) // a

}
