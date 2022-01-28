// go 基础 - 指针

package basic

import (
	"fmt"
)

func PointerSample() {
	pointer_sample1()
	pointer_sample2()
	pointer_sample3()
}

func pointer_sample1() {
	// 作为数据类型时 * 表示指针，比如 *int 就是指向 int 的指针
	// 作为运算符时 * 表示取指针指向的值
	// 运算符 & 表示取值的指针

	var a int = 1   // 为变量 a 赋值 1
	var b *int = &a // 取变量 a 的地址给指针 b
	var c int = *b  // 取指针 b 指向的值给 c

	// b 指向的值是 a
	// a 和 c 的值是相同的，但是他们的指针是不同的
	fmt.Printf("%d, %d, %d\n", a, *b, c)  // 1, 1, 1
	fmt.Printf("%p, %p, %p\n", &a, b, &c) // 0xc000128058, 0xc000128058, 0xc00012807
	fmt.Printf("%T, %T, %T\n", a, b, c)   // int, *int, int

	// 注：虽然切片是一个指针，但是不能对其做取值操作，否则会编译时报错 invalid operation: cannot indirect d (variable of type []int)
	// d := []int{1, 2, 3, 4, 5}
	// fmt.Println(*d)
}

func pointer_sample2() {
	// 声明一个指向 int 的指针，但是不初始化，那么运行时是会报错的 "runtime error: invalid memory address or nil pointer dereference"
	// var a *int

	// 可以通过 new() 的方式声明一个指向指定类型的指针，其会自动用默认值初始化
	b := new(int)
	fmt.Println(*b) // 0
	*b = 1
	fmt.Println(*b) // 1
}

func pointer_sample3() {
	a := &struct {
		a int
	}{}

	// 虽然 a 是指针，需要这么 (*a).a = 123 使用，但是 go 是支持语法糖（syntactic sugar）技术的，他会自动转换的，所以你可以这么 a.a = 456 使用
	(*a).a = 123
	a.a = 456
	fmt.Println(a) // &{456}
}
