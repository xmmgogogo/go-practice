package main

import (
	"fmt"
	"sync"
)

var (
	catChan  = make(chan struct{})
	dogChan  = make(chan struct{})
	fishChan = make(chan struct{})
	num      = 2
)

func main() {
	run()
}

func run() {
	go func() {
		catChan <- struct{}{}
	}()

	wg := sync.WaitGroup{}
	{
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			var c int
			for {
				<-catChan
				fmt.Println("cat")

				dogChan <- struct{}{}
				c++
				if c >= num {
					close(catChan)
					return
				}
			}

		}(&wg)
	}
	{
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			var c int
			for {
				<-dogChan
				fmt.Println("dog")

				fishChan <- struct{}{}
				c++
				if c >= num {
					close(dogChan)
					return
				}
			}
		}(&wg)
	}
	{
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			var c int
			for {
				<-fishChan
				fmt.Println("fish")

				c++
				if c >= num {
					close(fishChan)
					return
				} else {
					catChan <- struct{}{}
				}
			}
		}(&wg)
	}
	wg.Wait()
}
