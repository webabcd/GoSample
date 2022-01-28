// go 基础 - 指针

package basic

import (
	"fmt"
)

func PointerSample() {
	pointer_sample1()
	pointer_sample2()
	pointer_sample3()
	pointer_sample4()
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
	// 声明一个指针，其默认值为 nil
	var a *int
	fmt.Printf("%p\n", a) // 0x0
	// nil 的地址是 0x0，所以像下面这样赋值是不可以的，因为无法对地址 0x0 取值
	// *a = 100

	// 可以通过 new() 的方式声明一个指向指定类型的指针，其会自动用默认值初始化
	b := new(int)
	fmt.Printf("%d, %p\n", *b, b) // 0, 0xc000186060
	*b = 1
	fmt.Printf("%d, %p\n", *b, b) // 1, 0xc000186060
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

// 本例用于演示指针的指针
// 指针会指向一个值，也就是说这个值的地址就是这个指针，而保存这个地址的变量的地址，就是指针的指针
func pointer_sample4() {
	a := 0   // 变量 a
	b := &a  // 变量 a 的指针
	c := &b  // 变量 b 的指针，即变量 a 的指针的指针（注意：这里不支持 &&a 的写法）
	d := *b  // 对变量 b 取值，返回变量 a 的值
	e := *c  // 对变量 c 取值，返回变量 a 的指针
	f := **c // 对变量 c 取值再取值，返回变量 a 的值

	fmt.Printf("%v, %v, %v, %v, %v, %v\n", a, b, c, d, e, f) // 0, 0xc0000140e8, 0xc000006030, 0, 0xc0000140e8, 0

	// 声明一个指针的指针
	var g **int
	// 上面说过了，像下面这样赋值是不允许的，因为 nil 的地址是 0x0，无法对 0x0 取值
	// **g = 100
	g = c
	fmt.Printf("%v, %v, %v\n", g, *g, **g) // 0xc000006030, 0xc0000140e8, 0
}
