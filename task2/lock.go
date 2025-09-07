package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

type UnsafeCounter struct {
	value int
}

// Increment 方法安全地递增计数器
func (c *Counter) Increment() {
	c.mu.Lock()         // 获取锁
	defer c.mu.Unlock() // 确保在函数返回时释放锁
	c.value++
}

func (c *UnsafeCounter) Increment() {
	c.value++ // 这里会有竞态条件
}

func main() {
	counter := Counter{}
	unsafeCounter := UnsafeCounter{}
	atomicCounter := int64(0)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
				unsafeCounter.Increment()
				atomic.AddInt64(&atomicCounter, 1)
			}
			fmt.Printf("协程 %d 完成\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Println("使用Mutex锁操作的结果为", counter.value)
	fmt.Println("不使用任何原子操作的结果为", unsafeCounter.value)
	fmt.Println("使用原子操作atomic的结果为：", atomicCounter)
}
