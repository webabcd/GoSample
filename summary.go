// go 概述 - 基本语法

package main

import (
	"fmt"
	"math"
)

func summary() {
	summary_sample1()
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
