// go 面向对象 - 函数（函数也是一种类型，匿名函数，闭包）

package oop

import "fmt"

func Function2Sample() {
	function2_sample1()
	function2_sample2()
	function2_sample3()
}

func function2_sample1() {
	// 函数也是一种类型（可以实现类似函数指针的效果）
	// 下面声明了一个名为 myFunc 的变量，其类型为 func(name string) string
	var myFunc func(name string) string

	// 为函数类型的变量赋值时，要求函数的参数和返回值定义要与变量的类型相同
	myFunc = func21
	// 调用指针指向的函数
	fmt.Println(myFunc("webabcd")) // func21: webabcd

	// 为函数类型的变量赋值时，要求函数的参数和返回值定义要与变量的类型相同
	myFunc = func22
	// 调用指针指向的函数
	fmt.Println(myFunc("webabcd")) // func22: webabcd
}
func func21(name string) string {
	return "func21: " + name
}
func func22(name string) string {
	return "func22: " + name
}

func function2_sample2() {
	// 定义匿名函数
	a := func(name string) {
		fmt.Println("hello: " + name)
	}
	// 调用匿名函数
	a("webabcd") // hello: webabcd

	// 匿名函数做字典键值的示例
	var b = map[string]func(){
		"x": func() {
			fmt.Println("xxx")
		},
		"y": func() {
			fmt.Println("yyy")
		},
	}
	b["x"]() // xxx
	b["y"]() // yyy

	// 匿名函数实现回调的示例
	func23("webabcd", func(result string) {
		fmt.Println(result) // hello: webabcd
	})

}

// 此函数的第 2 个参数是一个函数类型，其用于实现回调逻辑
func func23(name string, callback func(string)) {
	// 回调
	callback("hello: " + name)
}

func function2_sample3() {
	// 获取一个闭包
	a := func24()

	// 闭包引用的闭包外的变量的生命周期会拉长到与闭包一致
	fmt.Println(a()) // 1
	fmt.Println(a()) // 2
	fmt.Println(a()) // 3
}

// 函数返回的函数就是一个闭包（closure）
func func24() func() int {
	// 闭包外的变量
	a := 0
	// 返回一个闭包
	return func() int {
		// 闭包的特性：闭包引用的闭包外的变量的生命周期会拉长到与闭包一致
		a++
		return a
	}
}
