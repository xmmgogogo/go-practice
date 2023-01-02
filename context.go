package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cacel := context.WithTimeout(context.Background(), time.Second*3)
	for i := 0; i < 10; i++ {
		goRun(ctx, i)
		if i >= 5 {
			cacel()
			break
		}
	}

	time.Sleep(time.Second * 30)
}

func goRun(ctx context.Context, n int) {
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("[%d]任务完成\n", n)
				return
			default:
			}
		}
	}(ctx)
}
