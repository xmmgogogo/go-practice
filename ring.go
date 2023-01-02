package main

import (
	"container/ring"
	"fmt"
)

// 练习golang里面的一个数据结构，环结构
func main() {
	ringCount()
}

func ringCount() {
	r := ring.New(5)
	for i := 0; i < 10; i++ {
		r.Value = i
		r = r.Next()
	}

	fmt.Printf("Length of ring: %d\n", r.Len())

	var num int
	r.Do(func(i interface{}) {
		fmt.Println(i)
		num += i.(int)
	})

	fmt.Println("总计：", num)
}
