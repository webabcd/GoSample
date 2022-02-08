// go 面向对象 - 结构体（结构体标签，结构体和 json 互相转换）

package oop

import (
	"encoding/json"
	"fmt"
)

func Struct4Sample() {
	// 结构体转 json
	struct4_sample1()
	// json 转结构体
	struct4_sample2()
}

// 结构体成员的后面加 `` 用于定义结构体标签
// 结构体标签的应用之一就是：通过反射实现结构体与 json 的互相转换
type struct41 struct {
	// 因为要通过反射做结构体与 json 的互相转换，所以成员小写开头的话是做不到的（因为小写开头的不可导出，所以反射不出来）
	id int
	// 因为要通过反射做结构体与 json 的互相转换，所以成员小写开头的话是做不到的（因为小写开头的不可导出，所以反射不出来）
	name string `json:"my_name"`
	// 如果用 encoding/json 库，则默认的话，结构体的字段名称会映射到 json 的同名字段
	Age int
	// 如果用 encoding/json 库，则通过标签，可以指定结构体的字段名称与 json 的字段名称的映射关系
	Mobile string `json:"my_mobile"`
}

// 结构体转 json
func struct4_sample1() {
	a := struct41{
		id:     1001,
		name:   "webabcd",
		Age:    40,
		Mobile: "12345678900",
	}
	jsonBytes, _ := json.Marshal(a)
	jsonString := string(jsonBytes)
	fmt.Println(jsonString) // {"Age":40,"my_mobile":"12345678900"}
}

// json 转结构体
func struct4_sample2() {
	a := `{"id":1001,"my_name":"webabcd","Age":40,"my_mobile":"12345678900"}`
	var s struct41
	err := json.Unmarshal([]byte(a), &s)
	if err == nil {
		fmt.Printf("%+v\n", s) // {id:0 name: Age:40 Mobile:12345678900}
	}
}
