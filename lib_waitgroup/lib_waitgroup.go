package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 在函数退出时通知 WaitGroup 计数器减一

	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d completed\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)          // 每启动一个 Goroutine，增加 WaitGroup 计数器
		go worker(i, &wg)  // 启动 Goroutine
	}

	wg.Wait() // 阻塞等待，直到 WaitGroup 计数器减至零
	fmt.Println("All workers completed")
}


/*
$ go run .
Worker 5 started
Worker 1 started
Worker 2 started
Worker 3 started
Worker 4 started
Worker 1 completed
Worker 5 completed
Worker 3 completed
Worker 2 completed
Worker 4 completed
All workers completed

`sync.WaitGroup` 是 Go 语言标准库中 `sync` 包提供的一个同步原语，用于等待一组 Goroutine 完成其执行。

`WaitGroup` 主要用于以下场景：

1. **等待一组 Goroutine 完成：** 当有多个并发任务（Goroutine）时，
我们希望等待它们全部执行完毕再进行下一步操作。`WaitGroup` 提供了等待机制，确保所有的 Goroutine 执行完成后再继续执行主线程。

2. **避免主线程提前退出：** 当主线程启动了一些 Goroutine 但又不确定它们何时完成，为了避免主线程提前退出，
我们可以使用 `WaitGroup` 让主线程等待所有 Goroutine 完成后再退出。

`WaitGroup` 的基本使用方法包括三个主要方法：

- `Add(delta int)`: 用于增加 `WaitGroup` 的计数器。每次启动一个 Goroutine 时，调用 `Add(1)`；Goroutine 完成时，调用 `Add(-1)`。

- `Done()`: 当 Goroutine 完成时，调用 `Done()` 来通知 `WaitGroup` 计数器减一。

- `Wait()`: 主线程通过调用 `Wait()` 来阻塞等待，直到 `WaitGroup` 计数器减至零，表示所有 Goroutine 都已完成。

在上面的例子中，主线程启动了 5 个 Goroutine，每个 Goroutine 执行 `worker` 函数。在 `worker` 函数中，
使用 `defer wg.Done()` 来在函数退出时通知 `WaitGroup` 计数器减一。主线程最后通过 `wg.Wait()` 阻塞等待，直到所有 Goroutine 完成后继续执行。
*/