package main

import (
	"fmt"
	"time"
)

func main() {
	// 使用标记
	fmt.Println("---- break label ----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break re
		}
	}

	fmt.Println("执行结束")
	if time.Now().Hour() >= 18 {
		goto re2
	} else {
		goto re3
	}

re2:
	fmt.Println("re2")

re3:
	fmt.Println("re3")
}
