// go 高级 - 线程间通信 channel/select
// goroutine 相当于轻量级的线程（相对于传统线程来说效率要高出很多），通过 go 关键字创建 goroutine，通过 channel 做 goroutine 间的通信
// 说明：
// 1、一个 goroutine 向 channel 发送数据，另一个 goroutine 从 channel 接收数据（这是要阻塞的，直到接收到数据为止）
// 2、无缓冲区 channel 的意思是：写数据和读数据是一套动作，完成一套动作后才能执行下一套动作
//    向 channel 写数据时，如果 channel 中有数据未读，则是不能写的，要阻塞着
// 3、有缓冲区 channel 的意思是：向缓冲区写数据和从缓冲区读数据是分别进行的，缓冲区是先入先出的
//    只要缓冲区不满，即使没人读，也可以继续向缓冲区写数据，如果缓冲区满了，就要阻塞了

package advanced

import (
	"fmt"
	"time"
)

func Goroutine2Sample() {

	// 无缓冲区 channel 的 demo
	goroutine2_sample1()

	// 有缓冲区 channel 的 demo
	goroutine2_sample2()

	// select 的 demo
	// 本例演示了如何通过 select 为 channel 增加读取超时的判断
	goroutine2_sample3()
}

func goroutine2_sample1() {

	// 创建一个可以传输 int 类型数据的 channel（注：创建 channel 时必须指定传输的数据类型，每个 channel 只能传输你指定的数据类型）
	var ch chan int = make(chan int)

	// 创建一个只写的 channel
	// var ch_sendOnly chan<- int = make(chan int)
	// var ch_sendOnly chan<- int = make(chan<- int)

	// 创建一个只读的 channel
	// var ch_recvOnly <-chan int = make(chan int)
	// var ch_recvOnly <-chan int = make(<-chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("准备写入", i)
			// 通过 channel<- 向 channel 发送数据（即写数据）
			// 如果 channel 中有未读数据，则阻塞
			ch <- i
			fmt.Println("已经写入", i)
		}

		// 关闭 channel
		close(ch)
	}()

	// 通过 <-channel 从 channel 接收数据（即读数据）
	// 这是要阻塞的，直到读出为止
	data := <-ch
	fmt.Println("已经读出", data)

	// 遍历 channel 中的数据
	for data := range ch {
		fmt.Println("已经读出", data)
	}

	// 判断 channel 是否关闭了（返回 false 则说明 channel 已经关闭了）
	_, ok := <-ch
	fmt.Println("channel状态", ok)

	// 我这里某一次的运行结果如下：
	/*
		准备写入 0
		已经写入 0
		准备写入 1
		已经读出 0
		已经读出 1
		已经写入 1
		准备写入 2
		已经写入 2
		准备写入 3
		已经读出 2
		已经读出 3
		已经写入 3
		准备写入 4
		已经写入 4
		已经读出 4
		channel状态 false
	*/
}

func goroutine2_sample2() {
	// 创建 channel 时指定缓冲区的大小（默认是无缓冲区的）
	// 这个缓冲区的大小指的是可以保存的元素的数量
	ch := make(chan int, 5)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("准备写入", i)
			// 通过 channel<- 向缓冲区发送数据（即写数据）
			// 如果缓冲区满了，则阻塞
			ch <- i
			fmt.Println("已经写入", i)
		}

		// 关闭 channel
		close(ch)
	}()

	for i := 0; i < 5; i++ {
		// 通过 <-channel 从缓冲区接收数据（即读数据）
		// 这是要阻塞的，直到读出为止
		data := <-ch
		fmt.Println("已经读出", data)
	}

	// 判断 channel 是否关闭了（返回 false 则说明 channel 已经关闭了）
	_, ok := <-ch
	fmt.Println("channel状态", ok)

	// 我这里某一次的运行结果如下：
	/*
		准备写入 0
		已经写入 0
		准备写入 1
		已经写入 1
		准备写入 2
		已经写入 2
		准备写入 3
		已经写入 3
		准备写入 4
		已经写入 4
		已经读出 0
		已经读出 1
		已经读出 2
		已经读出 3
		已经读出 4
		channel状态 false
	*/
}

func goroutine2_sample3() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			// select 中的 case 必须是读 channel 或写 channel
			// 每次执行时，select 都会评估每个 case 语句
			//   1、如果存在可以执行的 case（即没有被阻塞），则从可以执行的 case 中随机选择一个执行
			//   2、如果不存在可以执行的 case
			//      a) 有 default 时：执行 default
			//      b) 没 default 时：阻塞，直到有 case 可以执行为止
			select {
			case data := <-ch:
				fmt.Println(time.Now(), "读出", data)
			case <-time.After(3 * time.Second):
				fmt.Println(time.Now(), "从 ch 通道中获取数据超时了，超时时间是 3 秒")
				quit <- true
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	// 阻塞，直到从 quit 通道中收到数据为止
	<-quit

	fmt.Println(time.Now(), "结束")

	// 我这里某一次的运行结果如下：
	/*
		2022-02-07 16:29:07.8439862 +0800 CST m=+0.006806901 读出 0
		2022-02-07 16:29:08.8481389 +0800 CST m=+1.010959601 读出 1
		2022-02-07 16:29:09.8511351 +0800 CST m=+2.013955801 读出 2
		2022-02-07 16:29:10.8637841 +0800 CST m=+3.026604801 读出 3
		2022-02-07 16:29:11.8757439 +0800 CST m=+4.038564601 读出 4
		2022-02-07 16:29:14.8760507 +0800 CST m=+7.038871401 从 ch 通道中获取数据超时了，超时时间是 3 秒
		2022-02-07 16:29:14.8760507 +0800 CST m=+7.038871401 结束
	*/
}
