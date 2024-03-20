package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// 主 Goroutine 在等待任务完成
	wg.Add(1)

	go func() {
		// 模拟执行任务
		fmt.Println("Doing some work...")
		// 假设在主 Goroutine 调用 Wait() 后执行 Done()
		wg.Done()
	}()

	// 在主 Goroutine 调用 Wait() 后立即添加新的任务
	// 但是 Wait() 已经在等待，这会导致 panic
	wg.Add(1)

	// 主 Goroutine 等待任务完成
	wg.Wait()
	fmt.Println("All tasks completed")
}

/*
$ go run .
Doing some work...
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0xc000014250?)
        /home/shz/sdk/go1.19/src/runtime/sema.go:62 +0x25
sync.(*WaitGroup).Wait(0xc000104f70?)
        /home/shz/sdk/go1.19/src/sync/waitgroup.go:139 +0x52
main.main()
        /home/shz/WorkSpace/gospace/gopractice/lib_sync/lib_sync.go:26 +0x8f
exit status 2

起初没看懂，后来根据代码和报错的提示可知，死锁了，因为所有的goroutines都结束了
而wg还差一个值没有减为0，所以程序肯定就卡死在这了

在使用时要保证所有的goroutine都调用了wg.Done()，否则就会出现死锁了
也就是函数Add()和Done()要是一样的
*/