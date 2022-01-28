// go 高级 - 通过接收器（receiver）为任意类型定义方法，即为指定类型扩展方法
// 1、本例演示了如何通过自定义类型为基本数据类型扩展方法
// 2、为结构体扩展方法请参见：/oop/struct3.go

package advanced

import "fmt"

func ReceiverSample() {
	receiver_sample1()
}

func receiver_sample1() {
	var a myint = 1

	// 调用通过接收器为 myint 实现的 add() 方法
	b := a.add(2)
	fmt.Println(b) // 3
}

// go 是不允许为基本数据类型扩展方法的
// 可以通过 type 定义一个自定义类型，然后对这个自定义类型扩展方法
type myint int // 这个是自定义类型（注意：它不是别名，别名的写法是 type myint = int）

// 为 myint 定义一个 add() 方法，这玩意叫做接收器（receiver）
func (m myint) add(other int) int {
	return int(m) + other
}
