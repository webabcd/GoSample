// go 基础 - 语句（if/else, switch, for, continue, break, goto）
// 注：需要用到 {} 的语句，{ 必须要与语句在同一行，否则会报错

package basic

import (
	"fmt"
	"math/rand"
	"time"
)

func StatementSample() {
	statement_sample1()
	statement_sample2()
	statement_sample3()
	statement_sample4()

	var a int = 123
	// {} 是一个独立的块，可以访问父级的变量，但是无法访问同级或子级 {} 中的变量
	{
		var b int = 456
		fmt.Println(a, b) // 123 456
	}
	{
		var b int = 456
		fmt.Println(a, b) // 123 456
	}
	// 下面这句会报错，因为无法访问子级 {} 中的变量
	// fmt.Println(b)
}

// if/else
func statement_sample1() {
	// 0 到 2 之间的随机数
	a := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(3)
	fmt.Printf("a: %d\n", a)

	if a == 0 { // { 必须和关键字在同一行，否则会报错。另外，即使 if 后只有一条语句也不能省略 {}
		fmt.Println("a == 0")
	} else if a == 1 { // } 必须和 else if 在同一行，否则会报错
		fmt.Println("a == 1")
	} else {
		fmt.Println("a == 2")
	}

	// if 语句也支持先赋值，再判断
	if b := a; b > 0 {
		fmt.Println("b := a; b > 0")
	}
}

// switch
func statement_sample2() {
	// 0 到 9 之间的随机数
	a := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10)
	fmt.Printf("a: %d\n", a)

	// 注：switch 的 case 是不需要 break 的
	switch a {
	case 0, 1: // case 支持多个值
		fmt.Println("case 0, 1")
	case 2:
		fmt.Println("case 2")
	default:
		fmt.Println("default")
	}

	// switch 后面不跟变量，则可以在 case 中使用条件语句
	switch {
	case a < 5:
		fmt.Println("a < 5")
	case a < 10:
		fmt.Println("a < 10")
		fallthrough // switch 的 case 不需要 break 了，那么我就想继续走下一句怎么办呢？用 fallthrough 即可
	case a < 8:
		fmt.Println("a < 8")
	}
}

// for
// for 语句支持 continue 和 break
func statement_sample3() {
	// 经典 for 语句
	for a := 0; a < 3; a++ {
		fmt.Println(a)
	}

	b := 0
	// for 后的语句不需要的话可以不写
	for ; ; b++ {
		if b >= 3 {
			break
		}
		fmt.Println(b)
	}

	c := 0
	// for 后的语句全省略（无限循环）
	for {
		fmt.Println(c)
		c++
		if c >= 3 {
			break
		}
	}

	var d int
	// for 后只跟条件语句（相当于 while 循环）
	for d < 3 {
		fmt.Println(d)
		d++
	}

	e := []int{1, 2, 3}
	// for...range 用于遍历数组或切片或字典或字符串
	for index, value := range e {
		fmt.Printf("index:%d, value:%d\n", index, value)
	}
}

// goto
func statement_sample4() {
	goto onPrint

onPrint:
	fmt.Println("goto")
}
