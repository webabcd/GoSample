// go 容器 - 列表（列表的增删改查和遍历）
// 注：
// 1、List 中可以包含不同类型的数据
// 2、添加进 List 中的数据都会被封装为 Element 对象
// 3、通过 Element.Value 获取或修改元素的值

package container

import (
	"container/list"
	"fmt"
)

func ListSample() {
	list_sample1()
	list_sample2()
}

func list_sample1() {
	// 声明一个 List（此处的 a 是一个指针）
	a := list.New()
	// 声明一个 List（此处的 b 是指针所指的数据）
	var b list.List
	fmt.Printf("%p, %p\n", a, &b) // 0xc000110480, 0xc0001104b0

	// 在尾部添加一个数据
	a.PushBack("a") // a
	// 在开头添加一个数据
	a.PushFront(0) // 0 a
	// 在尾部添加一个 List
	a.PushBackList(a) // 0 a 0 a

	b.PushBack(true)
	// 在开头添加一个 List
	a.PushFrontList(&b) // true 0 a 0 a

	// 返回值就是你添加数据对应的 Element 对象
	el := a.PushBack("x") // true 0 a 0 a x
	// 在指定的 Element 对象之后添加数据
	a.InsertAfter("x_after", el) // true 0 a 0 a x x_after
	// 在指定的 Element 对象之前添加数据
	a.InsertBefore("x_before", el) // true 0 a 0 a x_before x x_after

	// 删除指定的 Element 对象
	a.Remove(el) // true 0 a 0 a x_before x_after

	// 遍历 List
	for i := a.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

func list_sample2() {
	a := list.New()
	a.PushBack(0)
	a.PushBack(1)
	a.PushBack(2)
	a.PushBack(3)
	a.PushBack(4)
	a.PushBack(5)

	// List 有如下方法
	//   Font() - 取第一个数据的 Element 对象
	//   Back() - 取最后一个数据的 Element 对象
	//   MoveToFront(), MoveToBack(), MoveBefore(), MoveAfter() - 移动 Element 对象的相关操作
	// Element 有如下方法和属性
	//   Next() - 获取当前 Element 对象的在 List 中的下一个 Element 对象（没有则会返回 nil）
	//   Prev() - 获取当前 Element 对象的在 List 中的上一个 Element 对象（没有则会返回 nil）
	//   Value - 表示 Element 对象的值
	b := a.Front().Next()        // 1
	c := a.Front().Next().Next() // 2
	// 将 c 移动到 List 的开头
	a.MoveToFront(c) // 2 0 1 3 4 5
	// 将 b 移动到 c 的后面
	a.MoveAfter(b, c) // 2 1 0 3 4 5

	// 修改指定 Element 对象的值
	a.Back().Value = 100 // // 2 1 0 3 4 5

	// 取指定 Element 对象的值
	fmt.Println(b.Value, c.Value) // 1 2

	// 遍历 List
	for i := a.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
