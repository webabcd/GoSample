// go 面向对象 - 接口（接口的定义，接口的实现，接口的使用）

package oop

import "fmt"

func Interface1Sample() {
	// 结构体实现接口的示例
	interface1_sample1()
	// 函数实现接口的示例（这个不常用）
	interface1_sample2()
}

func interface1_sample1() {

	// 声明一个 interface11 类型的变量，初始化的结构体必须要实现 interface11 接口
	var a interface11 = new(interface1_struct1)
	// 声明一个 interface12 类型的变量，初始化的结构体必须要实现 interface12 接口
	var b interface12 = new(interface1_struct1)
	// 声明一个 interface13 类型的变量，初始化的结构体必须要实现 interface13 接口（即必须同时实现 interface11 接口和 interface12 接口）
	var c interface13 = new(interface1_struct1)

	// 调用你实现的接口的方法
	a.hello("webabcd") // hello: webabcd
	b.hi("webabcd")    // hi: webabcd
	c.hello("webabcd") // hello: webabcd
	c.hi("webabcd")    // hi: webabcd

	// 接口作为参数传递
	interface1_func1(nil)
	interface1_func1(new(interface1_struct1))
}

func interface1_func1(a interface12) {
	if a == nil {
		fmt.Println("参数 a 为 nil") // 参数 a 为 nil
	} else {
		a.hi("webabcd") // hi: webabcd
	}
}

// 定义一个接口
type interface11 interface {
	hello(name string)
}

type interface12 interface {
	hi(name string)
}

type interface13 interface {
	// 内嵌接口
	// 内嵌接口可以实现类似继承的效果，这里相当于 interface13 继承了 interface11 和 interface12
	interface11
	interface12
}

// 定义一个结构体（后续会为这个结构体实现 interface11 接口，interface12 接口，interface13 接口）
type interface1_struct1 struct {
}

// 实现 interface11 接口的方法
func (s *interface1_struct1) hello(name string) {
	fmt.Println("hello: " + name)
}

// 实现 interface12 接口的方法
func (s *interface1_struct1) hi(name string) {
	fmt.Println("hi: " + name)
}

func interface1_sample2() {
	// 函数实现接口的示例
	var b interface11 = myfunc(interface1_func2)
	b.hello("webabcd") // hello: webabcd

	// 匿名函数实现接口的示例
	var a interface11 = myfunc(func(name string) {
		fmt.Println("hello: " + name)
	})
	a.hello("webabcd") // hello: webabcd
}

func interface1_func2(name string) {
	fmt.Println("hello: " + name) // hello: webabcd
}

// 为函数定义一个自定义类型（如果需要让函数实现接口，那么就需要通过自定义类型来实现）
type myfunc func(string)

// 为自定义类型实现 interface11 接口的方法
func (f myfunc) hello(name string) {
	// 调用函数
	f(name)
}
