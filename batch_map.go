package main

import (
	"fmt"
	"sync"
)

type ss struct {
	Name string
}

var (
	cMaps   = make(map[string]struct{}) // 并发线程不安全
	cMaps2  = map[string]struct{}{}     // 并发线程不安全
	cSlice  []int32                     // 并发线程不安全
	cArr    [10000]int
	cStruct = ss{}
)

func main() {
	//doMap() // not ok
	doSlice() // not ok
	//doArr() // ok
}

func doArr() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		{
			go func(wg *sync.WaitGroup, i int) {
				defer wg.Done()
				cArr[i] = 1
			}(&wg, i)
		}
	}
	wg.Wait()

	for i := 0; i < 10000; i++ {
		if cArr[i] != 1 {
			fmt.Println("error:", i)
		}
	}
}

func doMap() {
	n := 0
	for {
		go func() {
			cMaps[fmt.Sprintf("%d", n)] = struct{}{}
			//fmt.Println(cMaps["a"])
		}()
		n++
		if n > 100000 {
			break
		}
	}
}

func doSlice() {
	var testSlice []int
	for i := 0; i < 1000; i++ {
		go func() {
			testSlice = append(testSlice, i)
		}()
	}

	for idx, val := range testSlice {
		fmt.Printf("idx:%d val:%d\n", idx, val)
	}

	fmt.Println("切片长度：", len(testSlice)) // 不固定
}
