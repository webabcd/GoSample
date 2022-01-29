// go 概述 - 基本语法，nil
// 注：
// 1、go 是通过 GC 清理的
// 2、go 是通过开头字母大小写来控制可见性的
//    如果定义的常量、变量、类型、接口、接口中的方法，结构体、结构体中的成员、函数等的名称是大写字母开头，则表示能被其它包访问或调用，非大写字母开头就只能在包内使用

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

	// 运算符 =, +=, -=, *=, /=, %=, >>=, <<=, &=, ^=, |=, ||, &&, |, ^, &, ==, !=, <, <=, >, >=, <<, >>, +, -, *, /, %, !, *, &, ++, --

	// 变量的声明和初始化
	var a int = 1
	var b = 2           // 编译器推导类型
	var c, d int = 3, 4 // c, d int 的意思是 c 和 d 都是 int 类型
	var (
		e int = 5
		f     = 6
	)

	// 声明的简短格式（必须要初始化，必须要由编译器推导类型，只能在函数内部声明）
	g := 7
	h, i := 8, 9

	// const 常量
	const pi = math.Pi
	// _ 表示匿名变量，后续是不可使用的，如果你使用这个变量则会报错
	var _ int = 10

	fmt.Println(a, b, c, d, e, f, g, h, i, pi) // 1 2 3 4 5 6 7 8 9 3.141592653589793

	// 通过 {} 构造数组或切片或字典时，需要注意
	// } 与最后一个元素在同一行时，最后一个元素后面不用加逗号
	j := [3]int{1, 2, 3}
	// } 与最后一个元素不在同一行时，最后一个元素后面要加逗号
	k := [3]int{
		1, 2, 3,
	}
	fmt.Println(j, k) // [1 2 3] [1 2 3]

	l := &struct {
		a int
	}{}
	// 虽然 l 是指针，需要这么 (*l).a = 123 使用，但是 go 是支持语法糖（syntactic sugar）技术的，他会自动转换的，所以你可以这么 l.a = 456 使用
	(*l).a = 123
	l.a = 456
	fmt.Println(l) // &{456}
}

// nil
func summary_sample2() {
	// nil 和 nil 是不能做相等判断的，下面这句会编译时报错
	// fmt.Println(nil == nil)

	// 指针的默认值是 nil
	var a *int
	var b *int
	var c *string
	// 指针 a, b, c 本身的值都是 0x0
	fmt.Printf("%p, %p, %p\n", a, b, c) // 0x0, 0x0, 0x0
	// 判断指针是否为 nil 就用 == 即可
	fmt.Println(a == nil, b == nil, c == nil) // true true true
	// 指向相同类型的指针，如果他们都是 nil 则可以做相等判断（切片类型除外）
	fmt.Println(a == b) // true

	// 指向不同类型的指针，即使他们都是 nil 也不能做相等判断，下面这句会编译时报错
	// fmt.Println(a == c)

	var d []int
	var e []int
	// 可以通过 == 判断某一个切片是否为 nil
	fmt.Println(d == nil, e == nil) // true true
	// 但是因为切片类型是不能通过 == 做相等判断的，所以即使两个切片类型的指针都是 nil 也不能做相等判断，下面这句会编译时报错
	// fmt.Println(d == e)
}
