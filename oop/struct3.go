// go 面向对象 - 结构体（为结构体定义方法，使用工厂模式初始化结构体）

package oop

import "fmt"

func Struct3Sample() {
	// 为结构体指针添加方法
	struct3_sample1()
	// 为结构体指针指向的值添加方法
	struct3_sample2()
	// 使用工厂模式初始化结构体
	struct3_sample3()
}

func struct3_sample1() {
	a := &struct31{}

	// 调用结构体 struct31 的 add() 方法
	a.add(0)
	a.add(1)
	a.add(2)

	fmt.Println(a.items) // [0 1 2]
}

type struct31 struct {
	items []int
}

// 为结构体 struct31 定义一个 add() 方法，这玩意叫做接收器（receiver）
// 这里是为结构体指针添加方法
func (s *struct31) add(item int) {
	s.items = append(s.items, item)
}

func struct3_sample2() {
	a := struct32{1, 2}
	b := struct32{3, 4}

	// 调用结构体 struct32 的 add() 方法
	c := a.add(b)
	fmt.Println(c) // {4 6}
}

type struct32 struct {
	x int
	y int
}

// 为结构体 struct32 定义一个 add() 方法，这玩意叫做接收器（receiver）
// 这里是为结构体指针指向的值添加方法
func (s struct32) add(other struct32) struct32 {
	return struct32{s.x + other.x, s.y + other.y}
}

// 演示如何使用工厂模式初始化结构体
func struct3_sample3() {
	var a *struct33 = create_struct33("webabcd")
	a.age = 40
	fmt.Println(a) // &{webabcd 40}
}

type struct33 struct {
	name string
	age  int
}

// 工厂模式
func create_struct33(name string) *struct33 {
	return &struct33{
		name: name,
	}
}
