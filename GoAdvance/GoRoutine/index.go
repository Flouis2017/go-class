package GoRoutine

import (
	"fmt"
	"sync"
	"time"
)

// 什么是goroutine —— 在方法调用之前使用“go”关键字开启协程执行程序

func ShowTime()  {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}

func Run(wg *sync.WaitGroup) {
	fmt.Println("我跑起来了..")
	wg.Done()
}

func SyncTest() {
	ShowTime()
	for i := 1; i <= 1000; i++ {
		fmt.Print(i, ",")
	}
}

func AsyncTest() {
	go ShowTime()
	for i := 1; i <= 1000; i++ {
		fmt.Print(i, ",")
	}
}

// WaitGroupTest 协程管理器测试
func WaitGroupTest() {
	var wg sync.WaitGroup
	wg.Add(2)
	go Run(&wg)
	go Run(&wg)
	wg.Wait()
}

