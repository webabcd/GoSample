// go 面向对象 - 结构体（定义结构体，声明结构体，初始化结构体，匿名结构体）

package oop

import "fmt"

func Struct1Sample() {
	struct1_sample1()
	struct1_sample2()
}

func struct1_sample1() {
	// 声明一个指定类型的结构体（结构体中的每个成员会用默认值初始化）
	var a Struct11
	fmt.Println(a) // {0  }
	// 为结构体成员赋值
	a.a = 0
	a.b = "abc"
	a.c = "xyz"
	fmt.Println(a) // {0 abc xyz}

	// 用 new() 声明的结构体返回的是指针
	b := new(Struct11)
	// 因为 b 是指针，所以要按如下方式赋值
	(*b).a = 0
	// 因为 go 支持语法糖（syntactic sugar）技术，他会自动转换的，所以你可以按如下方式赋值
	b.a = 0
	b.b = "abc"

	// 声明一个指定类型的结构体，用 key/value 的方式对其成员做初始化，也允许不手工初始化，但是 {} 要保留
	// 前面加 & 返回的就是指针，前面不加 & 返回的就是指针指向的值
	c := &Struct11{
		a: 0,
		b: "abc",
	}
	fmt.Println(c) // &{0 abc }

	// 声明一个指定类型的结构体，按顺序对其成员做初始化（用这种方式的话，必须对每个成员都做初始化），也允许不手工初始化，但是 {} 要保留
	// 前面加 & 返回的就是指针，前面不加 & 返回的就是指针指向的值
	d := &Struct11{
		0,
		"abc",
		"xyz",
	}
	fmt.Println(d) // &{0 abc xyz}
}

// 定义一个结构体
type Struct11 struct {
	a    int
	b, c string
}

func struct1_sample2() {
	// 声明一个匿名结构体
	// 前面加 & 返回的就是指针，前面不加 & 返回的就是指针指向的值
	a := &struct {
		a int
		b string
	}{
		0,
		"abc",
	}
	fmt.Println(a) // &{0 abc}
}
