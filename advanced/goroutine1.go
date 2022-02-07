// go 高级 - 多线程和线程同步 goroutine
// 并发（concurrency）- 在某个 cpu 核心上多线程
// 并行（parallelism）- 在多个 cpu 核心上多线程（go 的多线程就是并行的，它会最大效率利用多核 cpu）
// goroutine 相当于轻量级的线程（相对于传统线程来说效率要高出很多），通过 go 关键字创建 goroutine，通过 channel 做 goroutine 间的通信

package advanced

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	// 在多线程场景下设置此值
	count int

	// sync.Mutex - 互斥锁
	// sync.RWMutex - 读写互斥锁(读锁占用，则阻止写不阻止读；写锁占用，则阻止读写)
	//   Lock() - 加锁
	//   Unlock() - 解锁
	lock sync.Mutex

	// sync.WaitGroup - 等待组（维护一个计数器，用于管理阻塞）
	//   Add(delta int) - 计数器加上指定的值
	//   Done() - 计数器减 1
	//   Wait() - 阻塞，等计数器等于 0 后停止阻塞
	wg sync.WaitGroup
)

func Goroutine1Sample() {

	// 获取 cpu 的核心数
	fmt.Println("NumCPU", runtime.NumCPU())
	// 设置程序可以用到的最大 cpu 核心数（go 会最大效率利用多核 cpu，这里的默认值就是全部 cpu 核心）
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 等待组计数器 +11
	wg.Add(11)

	for i := 0; i < 10; i++ {
		// go - 让指定的函数在新的 goroutine 运行
		go running()
	}

	// go - 让指定的匿名函数在新的 goroutine 运行
	go func() {
		// 等待组计数器 -1
		defer wg.Done()

		// 当前线程睡 100 毫秒
		time.Sleep(time.Millisecond * 100) // 单位是纳秒

		for {
			plusOne()
			if count > 100 {
				break
			}
		}
	}()

	// 阻塞，直到等待组计数器等于 0 后停止阻塞
	wg.Wait()

	fmt.Println("结束")
}

func running() {
	// 等待组计数器 -1
	defer wg.Done()

	// 当前线程睡 100 毫秒
	time.Sleep(time.Millisecond * 100) // 单位是纳秒

	for {
		plusOne()
		if count > 100 {
			break
		}
	}
}

// count 变量加 1，并打印当前 count 的值
// 这个函数用于演示多线程场景下的线程同步，如果不加锁的话输出就乱七八糟，如果加锁的话则输出符合预期
func plusOne() {
	// 解锁
	defer lock.Unlock()
	// 加锁
	lock.Lock()

	count++
	fmt.Println("count", count)
}
