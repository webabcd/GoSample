// go 容器 - 字典（字典的增删改查和遍历，以及字典的线程安全）

package container

import (
	"fmt"
	"sync"
)

func MapSample() {
	map_sample1()
	map_sample2()
}

// map 的增删改查和遍历
// map 在多线程场景下是线程不安全的
func map_sample1() {
	// 声明一个 map 并初始化
	var a map[string]string = map[string]string{"k1": "v1", "k2": "v2"}
	// 新增或修改 key/value
	a["k1"] = "v111"
	a["k3"] = "v3"
	fmt.Println(a, a["k1"]) // map[k1:v111 k2:v2 k3:v3] v111

	// 获取 map 中指定 key 的值，以及 map 中是否包含指定 key
	k1_value, k1_exists := a["k1"]
	_, k8_exists := a["k8"]
	fmt.Println(k1_value, k1_exists, k8_exists) // v111 true false

	// 删除 map 中的指定 key
	delete(a, "k1")

	// 遍历 map
	for k, v := range a {
		fmt.Println(k, v)
	}
	// k2 v2
	// k3 v3

	// 通过 make 构造一个 map
	b := make(map[int]string)
	b[0] = "a"
	fmt.Println(b) // map[0:a]

	// 需要完全清空 map 的话就给他一个空 map 即可
	b = make(map[int]string)
	fmt.Println(b) // map[]
}

// sync.Map 的增删改查和遍历
// sync.Map 在多线程场景下是线程安全的
func map_sample2() {
	// sync.Map 是 struct 类型
	// sync.Map 是线程安全的
	var a sync.Map

	// 新增或修改 key/value
	a.Store("k1", "v1")
	a.Store("k2", "v2")
	a.Store("k3", "v3")
	a.Store("k1", "v111")

	// 删除指定的 key
	a.Delete("k2")

	// 获取 sync.Map 中指定 key 的值，以及 sync.Map 中是否包含指定 key
	k1_value, k1_exists := a.Load("k1")
	fmt.Println(k1_value, k1_exists) // v111 true

	// 遍历 sync.Map
	a.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
	// k3 v3
	// k1 v111
}
