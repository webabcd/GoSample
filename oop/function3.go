// go 面向对象 - 函数（defer）

package oop

import "fmt"

func Function3Sample() {
	function3_sample1()
}

func function3_sample1() {
	func31()
	// 输出数据如下
	// func31 start
	// func31 end
	// 3
	// 2
	// 1

	func32()
	// 输出数据如下
	// func31 start
	// func31 running
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

	fmt.Println("func31 end")
}

// 如果函数遇到异常了，则先会把之前的 defer 语句执行完后，再抛出异常
func func32() {
	fmt.Println("func31 start")

	defer fmt.Println(1)
	defer fmt.Println(2)

	fmt.Println("func31 running")
	// 抛出异常
	panic("throw exception")

	defer fmt.Println(3)

	fmt.Println("func31 end")
}
