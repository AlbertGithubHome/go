package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type SharedCounter struct {
	counter int32
}

func (sc *SharedCounter) Increment() {
	atomic.AddInt32(&sc.counter, 1)
}

func (sc *SharedCounter) GetCounter() int32 {
	return atomic.LoadInt32(&sc.counter)
}

func main() {
	var wg sync.WaitGroup
	sharedCounter := &SharedCounter{}

	// 启动多个 Goroutine 并发地增加计数器
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sharedCounter.Increment()
		}()
	}

	// 在主 Goroutine 中读取计数器的值，使用原子读取
	counterValue := sharedCounter.GetCounter()
	fmt.Println("Current counter value:", counterValue)

	// 等待所有 Goroutine 完成
	wg.Wait()

	// 再次读取计数器的值，使用原子读取
	finalCounterValue := sharedCounter.GetCounter()
	fmt.Println("Final counter value:", finalCounterValue)
}

/*
$ go run .
Current counter value: 0
Final counter value: 5

一个简单的场景，其中有一个计数器（counter）是共享的资源，多个 Goroutine 需要读取该计数器的值。
我们将使用原子读取操作来展示它的用法及好处。


在这个例子中，`SharedCounter` 类型有一个计数器字段 `counter`，
它使用了 `sync/atomic` 包中的原子操作函数 `AddInt32` 和 `LoadInt32`。
`Increment` 方法用于增加计数器，`GetCounter` 方法用于读取计数器的值。

在主 Goroutine 中，我们使用 `GetCounter` 方法原子地读取计数器的当前值，并输出。
这样我们可以确保在读取时，计数器的值是一个一致的视图。然后，
我们在多个 Goroutine 中并发地增加计数器的值，最后再次使用原子读取操作获取最终的计数器值。

这个例子中的好处包括：

1. **防止竞态条件：** 使用原子读取操作能够防止在读取过程中多个 Goroutine 同时修改计数器值，避免竞态条件。

2. **提供一致性视图：** 在主 Goroutine 中读取计数器值时，我们获取了一个一致的视图，即在读取瞬间的值。

3. **减少锁的使用：** 使用原子操作相对于互斥锁来说，减小了同步的开销。

虽然在实际开发中，有时可能需要更复杂的同步机制，但原子读取操作提供了一种轻量级的方式，
适用于某些场景，尤其是在读取共享状态时的一致性要求较低的情况。
*/