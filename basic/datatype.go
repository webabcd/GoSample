// go 基础 - 基本数据类型（数字类型，字符串类型，字符串和数字类型的转换，字符串的格式化，枚举类型，类型别名）

package basic

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

func DataTypeSample() {
	dataType_sample1()
	dataType_sample2()
	dataType_sample3()
	dataType_sample4()
	dataType_sample5()
}

// 数字类型
func dataType_sample1() {
	// int、int8、int16、int32、int64, uint、uint8、uint16、uint32、uint64（其中的 int 和 uint 是占 32 位还是占 64 位，要看 cpu 是 32 位的还是 64 位的）
	// float32、float64
	// bool（false, true）

	// 格式化打印
	fmt.Printf("%.2f\n", math.Pi) // 3.14

	// 整型（'\x..' 以十六进制的方式定义整型）
	var a, b uint8 = 97, '\x61'
	// byte 类型（表示一个字节）其实就是 uint8 类型
	var c, d byte = 97, '\x61'
	// %d 整型，%x 十六进制, %c 字符
	fmt.Printf("%d, %d, %x, %x, %c, %c, %d, %d, %x, %x, %c, %c\n", a, b, a, b, a, b, c, d, c, d, c, d) // 97, 97, 61, 61, a, a, 97, 97, 61, 61, a, a

	var e int32 = '\u738b'
	// rune 类型（表示一个 unicode 编码）其实就是 int32 类型
	var f rune = '\u738b'
	// %d 整型，%x 十六进制, %c 字符
	fmt.Printf("%d, %x, %c，%d, %x, %c\n", e, e, e, f, f, f) // 29579, 738b, 王， 29579, 738b, 王

	// 通过 complex() 定义数学中的复数（由实数部分和虚数部分组成）
	i := complex(3.14, 0)   // 实数部分 3.14，虚数部分 0
	j := complex(3.14, 1.2) // 实数部分 3.14，虚数部分 1.2
	k := i + j              // 实数部分和实数部分相加减，虚数部分和虚数部分相加减
	fmt.Println(i, j, k)    // (3.14+0i) (3.14+1.2i) (6.28+1.2i)
	// real() - 取复数的实数部分
	// imag() - 取复数的虚数部分
	fmt.Println(real(k), imag(k)) // 6.28 1.2
}

// 字符串类型
func dataType_sample2() {
	var a = "webabcd"
	// 定义一个字符，实际上就是一个整型
	var b = 'a'
	// 通过 + 做字符串的拼接
	var c = a + "_wanglei"
	// %d 整型，%x 十六进制, %c 字符, %s 字符串
	// 字符串中的字符可以通过 [] 获取
	fmt.Printf("%s, %c, %c, %c, %d, %x, %s\n", a, a[0], a[len(a)-1], b, b, b, c) // webabcd, w, d, a, 97, 61, webabcd_wanglei

	// 通过 `` 定义的字符串会保留换行，且转义符无效
	var d = `abc
  ijk
    xyz
\r\n`
	fmt.Println(d)

	// 判断字符串是否相同（0 代表相同）
	fmt.Println(strings.Compare("abc", "abc")) // 0
	// 判断字符串是否相同（忽略大小写）
	fmt.Println(strings.EqualFold("ABC", "abc")) // true
}

// 字符串和数字类型的转换，字符串的格式化
func dataType_sample3() {
	// int32 类型转换为 int64 类型
	var a int32 = 10
	var b = int64(a)
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b)) // int32 int64

	// int 和 string 相互转换
	var c int64 = 97
	var d = strconv.FormatInt(c, 10)       // int64 按十进制转为字符串
	var e = strconv.FormatInt(c, 16)       // int64 按十六进制转换为字符串
	var f, _ = strconv.ParseInt(d, 10, 32) // 字符串按十进转换为 int32（最后一个参数表示转换后的 int 的位数）
	var g, _ = strconv.ParseInt(e, 16, 64) // 字符串按十六进制转换为 int64（最后一个参数表示转换后的 int 的位数）
	fmt.Println(c, d, e, f, g)             // 97 97 61 97 97

	// int 和 float 相互转换
	var h = math.Pi
	var i = strconv.FormatFloat(h, 'f', 4, 64) // float64 转换为字符串，保留 4 位小数（四舍五入）
	j, _ := strconv.ParseFloat(i, 32)          // 字符串转换为 float32（最后一个参数表示转换后的 float 的位数）
	fmt.Println(h, i, j)                       // 3.141592653589793 3.1416 3.1415998935699463

	// 通过 fmt.Sprintf 返回格式化后的字符串（注：fmt.Sprintf() 是不会打印的，只是返回格式化后的字符串）
	var k = fmt.Sprintf("%.2f", math.Pi)
	fmt.Println(k) // 3.14

	// %% - 输出 %
	// %d - 十进制
	// %b - 二进制
	// %o - 八进制
	// %x - 十六进制（小写）
	// %X - 十六进制（大写）
	// %c - 字符
	// %s - 字符串
	// %f - 浮点型
	// %v - 按值的方式输出
	// %T - 数据类型
	// %p - 指针地址
	var l = fmt.Sprintf("%%, %d, %b, %o, %x, %X, %c, %s, %.2f, %v, %T, %p", 'z', 'z', 'z', 'z', 'z', 'z', "z", 3.14159, [3]int{1, 2, 3}, [3]int{1, 2, 3}, &a)
	fmt.Println(l) // %, 122, 1111010, 172, 7a, 7A, z, z, 3.14, [1 2 3], [3]int, 0xc000014118
}

// 枚举类型
func dataType_sample4() {
	// go 没有枚举类型，但是可以通过 const 模拟一个枚举
	const (
		a1 = 10
		a2 = "hello"
	)

	// iota 只能在 const 中使用，下面的枚举会从 0 开始（如果想从 1 开始，那就 iota + 1）
	const (
		b1 = iota // 0
		b2        // 1
		b3        // 2
	)

	// iota 也能用来实现 flag 类型的枚举
	const (
		c1 = 1 << iota // 1 左移 0 位，即 1
		c2             // 1 左移 1 位，即 2
		c3             // 1 左移 2 位，即 4
		c4             // 1 左移 3 位，即 8
	)

	fmt.Println(a1, a2, b1, b2, b3, c1, c2, c3, c4) // 10 hello 0 1 2 1 2 4 8
}

// 类型别名
func dataType_sample5() {
	// 定义类型别名
	type MyType = int32
	// 使用类型别名
	var x MyType = 100
	fmt.Println(x, reflect.TypeOf(x)) // 100 int32
}
