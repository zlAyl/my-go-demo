package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 1; i < 100; i++ {

			select {
			case ch <- i:
				fmt.Printf("成功发送: %d (通道长度: %d/%d)\n",
					i, len(ch), cap(ch))
			default:
				// 通道满，等待一段时间再重试
				fmt.Printf("通道已满，等待... (尝试发送: %d)\n", i)
				time.Sleep(100 * time.Millisecond)
				i-- // 重试当前值
			}
			time.Sleep(50 * time.Millisecond)
		}
		fmt.Println("生产者完成")
	}()

	go func() {
		defer wg.Done()
		for i := range ch {
			fmt.Printf("消费者接收: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println("消费者完成")
	}()

	wg.Wait()
	fmt.Println("程序结束")
}
