package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 1; i <= 10; i++ {
			fmt.Printf("生产者发送: %d\n", i)
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println("生产者完成")
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case i, ok := <-ch:
				if !ok {
					fmt.Println("通道已关闭，消费者退出")
					return
				}
				fmt.Printf("消费者接收: %d\n", i)
			case <-time.After(500 * time.Millisecond):
				fmt.Println("接收超时，继续等待...")
			}
		}

	}()

	wg.Wait()
}
