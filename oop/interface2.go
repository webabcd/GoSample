// go 面向对象 - 接口（空接口，类型断言，通过实现内置的 error 接口实现自定义错误类型）

package oop

import (
	"errors"
	"fmt"
)

func Interface2Sample() {
	interface2_sample1()
	interface2_sample2()
}

func interface2_sample1() {
	// 声明一个空接口类型的变量（类似于 c# 的 object，也就是说空接口可以保存任何类型的值）
	var a interface{}
	a = 1
	fmt.Println(a)
	a = "abc"
	fmt.Println(a)

	var b interface{}
	b = "abc"
	// 如果两个空接口保存的值是相等的，则这两个空接口是相等的
	fmt.Println(a == b)

	// 如果需要从空接口中取出原类型的值，是不能像下面这样直接写的
	// var c string = b

	// 如果需要从空接口中取出原类型的值，则要使用类型断言（type assertion），其语法为 x.(T)
	var c string = b.(string)
	fmt.Println(c)
	// 也可以这么写，第一个返回值是转换后的值，第二个返回值用于表示转换是否成功
	e, f := b.(string)
	fmt.Println(e, f) // abc true

	// 通过类型断言判断接口的原数据类型是什么
	// 注：只能在 switch 中使用 a.(type) 这种写法
	switch a.(type) {
	case string:
		fmt.Println("a 是 string: " + a.(string)) // 要将空接口转为字符串，记得要像这样 a.(string) 断言
	default:
		fmt.Println("不知道 a 是啥类型")
	}

	// 类似这样 g := a.(type) 将断言的结果给 g，可以避免后续使用的时候再断言
	switch g := a.(type) {
	case string:
		fmt.Println("a 是 string: " + g) // g 就是 a 断言后的结果，所以不用再对 a 断言了
	default:
		fmt.Println("不知道 a 是啥类型")
	}

}

// 本例用于演示如何通过实现内置的 error 接口实现自定义错误类型
func interface2_sample2() {
	// 调用函数，拿到函数返回的自定义 error 类型
	_, err := interface2_func1(5)
	if err != nil {
		fmt.Println(err.Error()) // 错误信息：参数不能小于 10
	}

	// 调用函数，拿到函数返回的内置 error 类型
	_, err2 := interface2_func1(100)
	if err2 != nil {
		fmt.Println(err2.Error()) // 参数不能是 100
	}
}

// 这个函数的返回值包含一个 error 接口
func interface2_func1(a int) (int, error) {
	if a < 10 {
		// 返回自定义 error 类型
		return -1, interface2_struct1{"参数不能小于 10"}
	} else if a == 100 {
		// 返回内置 error 类型
		return -1, errors.New("参数不能是 100")
	} else {
		return a, nil
	}
}

// 定义一个结构体
type interface2_struct1 struct {
	errorMessage string
}

/*
type error interface {
    Error() string
}
*/
// 为结构体实现 error 接口，error 接口是一个内置接口，其定义就是上面注释的代码
func (e interface2_struct1) Error() string {
	return "错误信息：" + e.errorMessage
}
