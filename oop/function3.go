// go 面向对象 - 函数（defer, panic, recover）
// panic/recover/defer 可以实现类似 try/catch/finally 的效果

package oop

import (
	"fmt"
)

func Function3Sample() {
	function3_sample1()
	function3_sample2()
}

func function3_sample1() {
	func31()
	// 输出数据如下
	// func31 start
	// func31 end
	// 3
	// 2
	// 1
}

// 函数执行完毕后，才会调用 defer 语句
// 如果有多个 defer 语句，则遵守先进后出，后进先出的原则
func func31() {
	fmt.Println("func31 start")

	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	// 如果抛出异常，则后面的代码都不执行了，然后等把前面的 defer 都执行完后，再抛异常
	// panic("throw exception")

	fmt.Println("func31 end")
}

// 实现类似 try/catch/finally 效果的示例
func function3_sample2() {
	fmt.Println("function3_sample2 start")

	defer func() {
		fmt.Println("我是 finally")
	}()

	// 要有 defer 否则捕获不到 panic 异常
	defer func() {
		// recover() - 获取到 panic 异常，同时让程序恢复正常
		if err := recover(); err != nil {
			// recover() 返回的是一个空接口，这个空接口就是你 panic 时传入的空接口
			switch err := err.(type) {
			case string:
				fmt.Println("捕获到异常：" + err)
			default:
				fmt.Println("捕获到异常", err)
			}
		}
	}()

	func32()

	fmt.Println("function3_sample2 end") //这里开始下面代码不会再执行

	// 输出数据如下
	// function3_sample2 start
	// func33 start
	// 捕获到异常：func33 throw exception
	// 我是 finally
}

func func32() {
	fmt.Println("func33 start")
	// 手动抛出一个异常（传入的参数类型是一个空接口）
	panic("func33 throw exception")
	fmt.Println("func33 end")
}
