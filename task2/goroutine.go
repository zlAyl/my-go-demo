package main

import (
	"fmt"
	"sync"
	"time"
)

type Task func()

type Scheduler struct {
	Id    int
	tasks []Task
}

func newScheduler() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) AddScheduler(task Task) {
	s.tasks = append(s.tasks, task)

}

// runTasks 执行所有任务并等待完成
func (s *Scheduler) runTasks() {
	var wg sync.WaitGroup
	wg.Add(len(s.tasks))

	for k, task := range s.tasks {
		go func() {
			defer wg.Done()
			start := time.Now()
			task()
			duration := time.Since(start)
			fmt.Printf("任务【%d】完成，用时: %v \n", k+1, duration)
		}()
	}

	wg.Wait()
	fmt.Println("所有任务执行完成!")
}

func Print() {
	// 使用 WaitGroup 等待所有协程完成
	var wg sync.WaitGroup
	wg.Add(2) // 等待两个协程

	//奇数协程
	go func() {
		defer wg.Done() // 协程结束时通知 WaitGroup
		for i := 0; i <= 10; i += 2 {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 100)
		}
	}()

	//偶数协程
	go func() {
		defer wg.Done() // 协程结束时通知 WaitGroup
		for i := 1; i <= 10; i += 2 {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 100)
		}
	}()

	// 等待所有协程完成
	wg.Wait()
	fmt.Println("打印完成!")

}

func exampleTask1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("任务1执行完成")
}

func exampleTask2() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("任务2执行完成")
}

func Schedulers() {

	scheduler := newScheduler()
	scheduler.AddScheduler(exampleTask1)
	scheduler.AddScheduler(exampleTask2)
	scheduler.runTasks()
}

func main() {
	//Print()
	Schedulers()
}
