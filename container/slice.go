// go 容器 - 切片（通过范围获取切片，切片的构造，切片数据的添加/复制/删除，数组和切片的区别）

// 数组是内存中的一段连续空间，切片是对数组的某一段连续空间的引用
// 数组是固定长度的，切片是长度可变的
// 数组是指针所指的数据，切片是一个指针

package container

import (
	"fmt"
	"reflect"
)

func SliceSample() {
	slice_sample1()
	slice_sample2()
	slice_sample3()
	slice_sample4()
	slice_sample5()
	slice_sample6()
}

// 从数组或切片中获取指定范围的切片
func slice_sample1() {
	var a = [5]int{1, 2, 3, 4, 5}

	// 获取指定位置范围的切片（左闭右开原则，即包括左边，不包括右边）
	fmt.Println(a[0:3]) // [1 2 3]
	fmt.Println(a[:5])  // [1 2 3 4 5]
	fmt.Println(a[0:])  // [1 2 3 4 5]
	fmt.Println(a[0:0]) // []
	fmt.Println(a[:])   // [1 2 3 4 5]

	// a 是数组，它是固定长度的
	fmt.Println(reflect.TypeOf(a)) // [5]int
	// a[:] 是切片，它的长度是可变的
	fmt.Println(reflect.TypeOf(a[:])) // []int
}

// 构造切片
func slice_sample2() {
	// 构造一个切片
	//   第 2 个参数表示切片初始化时元素的个数
	//   第 3 个参数表示初始化时，为切片分配的初始空间的长度
	//     如果不指定此参数，则其值同第 2 个参数
	//     后续，如果切片中的元素个数发生了变化，系统会根据它自己的逻辑决定是否重新分配空间
	var a []int = make([]int, 2)
	b := make([]int, 2, 10)

	// 构造一个切片，并初始化（与数组的区别是 [] 里面啥都没有）
	c := []int{0, 0}

	// len() - 获取切片的长度
	// cap() - 获取切片的占用空间的长度
	fmt.Println(a, b, c, len(a), len(b), len(c), cap(a), cap(b), cap(c)) // [0 0] [0 0] [0 0] 2 2 2 2 10 2
}

// 在切片的头部或尾部添加数据
func slice_sample3() {
	a := []int{}

	// 在尾部添加元素
	a = append(a, 0, 1, 2)
	a = append(a, []int{3, 4, 5}...)
	fmt.Println(a) // [0 1 2 3 4 5]

	// 在开头添加元素
	a = append([]int{-3, -2, -1}, a...)
	fmt.Println(a) // [-3 -2 -1 0 1 2 3 4 5]
}

// 切片的复制
func slice_sample4() {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 10)
	c := make([]int, 10)

	// 将 a 复制给 b（返回值为复制的元素个数）
	d := copy(b, a)

	// 将 a[2:5] 复制给 c[2:5]（返回值为复制的元素个数）
	e := copy(c[2:5], a[2:5])

	fmt.Println(a, b, c, d, e) // [1 2 3 4 5] [1 2 3 4 5 0 0 0 0 0] [0 0 3 4 5 0 0 0 0 0] 5 3
}

// 切片中元素的删除
func slice_sample5() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	fmt.Println(a) // [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18]

	a = a[2:]        // 删除开头 2 个元素
	a = a[:len(a)-2] // 删除尾部 2 个元素
	fmt.Println(a)   // [3 4 5 6 7 8 9 10 11 12 13 14 15 16]

	a = append(a[:0], a[2:]...) // 删除开头 2 个元素
	a = a[:copy(a, a[2:])]      // 删除开头 2 个元素
	fmt.Println(a)              // [7 8 9 10 11 12 13 14 15 16]

	a = append(a[:2], a[5:]...)  // 删除索引位置 2 到 4 之间的 3 个元素
	a = a[:2+copy(a[2:], a[5:])] // 删除索引位置 2 到 4 之间的 3 个元素
	fmt.Println(a)               // [7 8 15 16]
}

// 数组和切片的区别
func slice_sample6() {
	// 数组，固定长度
	a := [3]int{1, 2, 3}
	// 切片，长度可变
	b := []int{1, 2, 3}

	// 数组 a 是指针所指的数据
	// 切片 b 是一个指针
	fmt.Printf("%v, %v, %T, %T, %p, %p\n", a, b, a, b, &a, b) // [1 2 3], [1 2 3], [3]int, []int, 0xc00000e168, 0xc00000e180

	// 切片是对数组的一段连续区域的引用
	c := a[:2]
	d := a[1:]
	// 所以 &a 和 c 的地址是一样的，而 &a 和 d 的地址是不一样的
	fmt.Printf("%v, %v, %v, %p, %p, %p\n", a, c, d, &a, c, d) // [1 2 3], [1 2], [2 3], 0xc00000e168, 0xc00000e168, 0xc00000e170
}
