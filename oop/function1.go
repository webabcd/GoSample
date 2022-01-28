// go 面向对象 - 函数（函数基础，带变量名的返回值，多返回值，可变参数）

package oop

import "fmt"

func Function1Sample() {
	function1_sample1()
	function1_sample2()
}

func function1_sample1() {
	// 调用函数
	a := func11(0, "a", "b")
	fmt.Println(a) // 0ab

	// 调用函数
	b := func12()
	fmt.Println(b) // 100

	// 调用函数（此函数有多个返回值）
	c1, c2, c3 := func13()
	fmt.Println(c1, c2, c3) // 0 a true
}

// 定义一个函数
// 参数 a 是 int 类型，参数 b 和 c 是 string 类型，返回值是 string 类型
func func11(a int, b, c string) string {
	return fmt.Sprint(a, b, c)
}

// 为函数定义一个带有变量名的返回值
func func12() (x int) {
	// 变量会用默认值初始化
	fmt.Println(x) // 0
	x = 100

	// 直接 return 返回的就是你上面定义的返回值的变量名
	return
	// 也可以 return 指定数据的
	// return x
}

// 函数支持多返回值
func func13() (int, string, bool) {
	return 0, "a", true
}

func function1_sample2() {
	// 调用带有可变参数的函数（可变参数为 int 类型的可变参数）
	func14("webabcd", 0, 1, 2)

	// 调用带有可变参数的函数（可变参数为任意类型的可变参数）
	func15(0, "a", true)
}

// 函数支持可变参数，可变参数要放到其他参数的后面
// 下面的函数定义了一个 int 类型的可变参数
func func14(a string, args ...int) {
	fmt.Println(a)
	// 遍历可变参数
	for _, arg := range args {
		fmt.Println(arg)
	}
}

// 下面的函数定义了一个任意类型的可变参数
func func15(args ...interface{}) {
	// 将可变参数传给其他函数时要注意参数名后面要跟 ...
	func16(args...)
}

// 下面的函数定义了一个任意类型的可变参数
func func16(args ...interface{}) {
	// 遍历可变参数
	for _, arg := range args {
		// 获取当前遍历出的参数的数据类型
		switch arg.(type) {
		case int:
			fmt.Println(arg, "int")
		case string:
			fmt.Println(arg, "string")
		case bool:
			fmt.Println(arg, "bool")
		}
	}
}
