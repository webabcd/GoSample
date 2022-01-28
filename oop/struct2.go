// go 面向对象 - 结构体（内嵌结构体，内嵌结构体实现类似继承的效果，内嵌匿名结构体，内嵌类型）

package oop

import "fmt"

func Struct2Sample() {
	struct2_sample1()
}

func struct2_sample1() {
	a := &struct22{}
	a.c = true
	a.a = 0     // 可以直接使用内嵌结构体的成员，当然也可以像这样写全了 a.struct21.a（如果成员名字有冲突，则必须按照这种方式使用）
	a.b = "abc" // 可以直接使用内嵌结构体的成员，当然也可以像这样写全了 a.struct21.b（如果成员名字有冲突，则必须按照这种方式使用）
	a.d.x = 123
	a.d.y = "xyz"
	a.float32 = 3.14 // 使用内嵌类型
	fmt.Println(a)   // &{true {0 abc} {123 xyz} 3.14}

	// 顺序初始化（所有成员都要初始化）
	b := &struct22{
		true,
		struct21{
			0,
			"abc",
		},
		// 内嵌匿名结构体的初始化
		struct {
			x int
			y string
		}{
			123,
			"xyz",
		},
		3.14,
	}
	fmt.Println(b) // &{true {0 abc} {123 xyz} 3.14}

	// 按 key/value 初始化（不要求所有成员都初始化）
	c := &struct22{
		c: true,
		struct21: struct21{
			0,
			"abc",
		},
		// 内嵌匿名结构体的初始化
		d: struct {
			x int
			y string
		}{
			123,
			"xyz",
		},
		float32: 3.14,
	}
	fmt.Println(c) // &{true {0 abc} {123 xyz} 3.14}

}

type struct21 struct {
	a int
	b string
}

type struct22 struct {
	c bool

	// 内嵌结构体（同类型的只能有一个）
	// 内嵌结构体可以实现类似继承的效果，这里相当于 Struct21 是基类，Struct22 是子类
	struct21

	// 内嵌匿名结构体
	d struct {
		x int
		y string
	}

	// 内嵌类型（同类型的只能有一个）
	float32
}
