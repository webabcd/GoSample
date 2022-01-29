// go 面向对象 - 函数（函数基础，带变量名的返回值，多返回值，可变参数，传参时指针和非指针的区别）

package oop

import "fmt"

func Function1Sample() {
	function1_sample1()
	function1_sample2()
	function1_sample3()
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
		switch arg.(type) { // 这个是接口的类型断言
		case int:
			fmt.Println(arg, "int")
		case string:
			fmt.Println(arg, "string")
		case bool:
			fmt.Println(arg, "bool")
		}
	}
}

// 本例由于演示传参时指针和非指针的区别
func function1_sample3() {
	a, b := 0, 0
	c := &b
	fmt.Printf("%d, %d, %p, %p, %p\n", a, b, &a, &b, &c) // 0, 0, 0xc0000140c0, 0xc0000140c8, 0xc000006030
	func17(a, &b)
	fmt.Printf("%d, %d, %p, %p, %p\n", a, b, &a, &b, &c) // 0, 1, 0xc0000140c0, 0xc0000140c8, 0xc000006030
}

// 无论实参是非指针还是指针，形参都会复制实参，然后在函数中使用
// 实参是非指针，则形参会复制实参，然后在函数中使用，所以函数中对形参的操作不会影响到实参
// 实参是指针，则形参会复制实参（也就是说形参指向的值和实参指向的值是同一个），然后在函数中使用，所以函数中对形参指向的值做操作就是对实参指向的值做操作
func func17(a int, b *int) {
	// 注意：这里实参 b 和形参 b 本身保存的值（这个保存的值是一个指针）是相同的，但是实参 b 的指针和形参 b 的指针是不同的（也就是说不管是不是指针，形参都会将实参复制出一份并使用）
	fmt.Printf("%d, %d, %p, %p, %p\n", a, *b, &a, b, &b) // 0, 0, 0xc0000140d0, 0xc0000140c8, 0xc000006038
	a = 1
	*b = 1
	fmt.Printf("%d, %d, %p, %p, %p\n", a, *b, &a, b, &b) // 1, 1, 0xc0000140d0, 0xc0000140c8, 0xc000006038
}
