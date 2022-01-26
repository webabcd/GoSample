// go 容器 - 数组（一维数组，多维数组，数组元素的获取和设置，数组的遍历）

package container

import "fmt"

func ArraySample() {
	array_sample1()
	array_sample2()
}

// 一维数组
func array_sample1() {
	// 声明一个长度为 5 个元素的整型数组，并初始化
	var a [5]int = [5]int{1, 2, 3, 4, 5}

	// 声明一个长度为 5 个元素的整型数组，并初始化
	// 初始化时可以按位置指定值，未指定值则初始化为默认值
	b := [5]int{0: 1, 3: 4}
	// 通过 [] 获取或设置指定位置的值
	b[len(b)-1] = 5

	// ... 表示数组的长度根据初始化时元素的个数决定
	c := [...]int{1, 2, 3, 4, 5}

	// 如果两个数组的元素个数，元素类型，每个位置的值都相同，则这两个数组是相等的
	d := a == c

	fmt.Println(a, b, c, c[len(c)-1], d) // [1 2 3 4 5] [1 0 0 4 5] [1 2 3 4 5] 5 true

	// 经典的数组遍历的方式
	for i := 0; i < len(a); i++ {
		fmt.Printf("index:%d, value:%d\n", i, a[i])
	}

	// 通过 range 遍历数组（获取索引位置和对应的值）
	for i, v := range a {
		fmt.Printf("%d:%d\n", i, v)
	}

	// 通过 range 遍历数组（仅获取值）
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}

// 二维数组
func array_sample2() {
	// 声明一个二维整型数组（2 行 3 列），并初始化
	var a [2][3]int = [2][3]int{{11, 12, 13}, {21, 22, 23}}

	// 声明一个二维整型数组（2 行 3 列），并初始化
	// 初始化时可以按位置指定值，未指定值则初始化为默认值
	b := [2][3]int{1: {21, 22, 23}}

	// ... 表示数组的长度根据初始化时元素的个数决定（对于二维数组来说，只能在第一维使用 ...）
	c := [...][3]int{{1, 1, 1}, {1, 1, 1}}

	// 通过 [] 获取或设置指定位置的值
	c[0] = [...]int{11, 12, 13}
	c[1][0] = 21
	c[1][1] = 22
	c[1][2] = 23

	fmt.Println(a, b, c) // [[11 12 13] [21 22 23]] [[0 0 0] [21 22 23]] [[11 12 13] [21 22 23]]

	// 通过 range 遍历二维数组
	for i, v := range a {
		for i2, v2 := range v {
			fmt.Printf("%d - %d:%d\n", i, i2, v2)
		}
	}
}
