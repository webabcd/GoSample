// go 概述 - 基本语法，nil
// 注：go 是通过 GC 清理的

package main

import (
	"fmt"
	"math"
)

func summary() {
	summary_sample1()
	summary_sample2()
}

// 基本语法
func summary_sample1() {

	// 变量的声明和初始化
	var a int = 1
	var b = 2           // 编译器推导类型
	var c, d int = 3, 4 // c, d int 的意思是 c 和 d 都是 int 类型
	var (
		e int = 5
		f int = 6
	)

	// 声明的简短格式（必须要初始化，必须要由编译器推导类型，只能在函数内部声明）
	g := 7
	h, i := 8, 9

	// const 常量
	const pi = math.Pi
	// _ 表示匿名变量，后续是不可使用的，如果你使用这个变量则会报错
	var _ int = 10

	fmt.Println(a, b, c, d, e, f, g, h, i, pi) // 1 2 3 4 5 6 7 8 9 3.141592653589793
}

// nil
func summary_sample2() {
	// nil 和 nil 是不能做相等判断的，下面这句会报错
	// fmt.Println(nil == nil)

	var a *int
	var b *string
	// a 和 b 都是 nil（但是他们是不能做相等判断的）
	fmt.Println(a, b) // <nil> <nil>
	// a 和 b 的地址都是 0x0
	fmt.Printf("%p, %p\n", a, b) // 0x0, 0x0
	// 可以判断地址是否相等
	fmt.Println(fmt.Sprintf("%p", a) == fmt.Sprintf("%p", b)) // true
}
