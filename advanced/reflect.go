// go 高级 - 反射（通过反射获取变量的类型，通过反射获取变量的值，修改变量的值，获取结构体成员的值，修改结构体成员的值，调用结构体的方法，通过反射实例化，通过反射调用函数）

package advanced

import (
	"fmt"
	"reflect"
)

func ReflectSample() {
	// 通过反射获取变量的类型
	reflect_Sample1()
	// 通过反射获取变量的值，修改变量的值，获取结构体成员的值，修改结构体成员的值，调用结构体的方法
	reflect_Sample2()
	// 通过反射实例化，通过反射调用函数
	reflect_Sample3()
}

func reflect_Sample1() {
	a := myStruc2{salary: 3500, myStruct1: myStruct1{name: "webabcd", age: 40}}

	// reflect.TypeOf() - 通过反射获取指定变量的类型
	//   Name() - 类型的名称
	//   Kind() - 类型的种类
	b := reflect.TypeOf(a)
	fmt.Println(b, b.Name(), b.Kind()) // advanced.myStruc2 myStruc2 struct

	c := &a
	d := reflect.TypeOf(c)
	// 如果是指针类型，则可以通过 Elem() 获取指针指向的值的类型
	fmt.Println(d.Kind(), d.Elem(), d.Elem().Name(), d.Elem().Kind()) // ptr advanced.myStruc2 myStruc2 struct

	// 遍历结构体的所有成员
	for i := 0; i < b.NumField(); i++ {
		// 获取结构体的指定索引的成员
		field := b.Field(i)
		// 获取成员的类型，名称，标签
		fmt.Printf("type:%v, name:%v, tag:%v\n", field.Type, field.Name, field.Tag)
	}
	// type:float32, name:salary, tag:
	// type:advanced.myStruct1, name:myStruct1, tag:
	// type:float32, name:SalaryExpected, tag:
	// type:*int, name:ptr, tag

	// 根据名称查找结构体的指定成员
	field_salary, ok_salary := b.FieldByName("salary")
	fmt.Println(field_salary.Type, ok_salary) // float32 true
	// 根据索引查找结构体的指定成员（这种方式支持多级查找，本例先找 myStruc2 的索引为 1 的成员，即 myStruct1，然后再找 myStruct1 的索引为 0 的成员，即 name）
	field_name := b.FieldByIndex([]int{1, 0})
	fmt.Println(field_name.Type) // string
}

func reflect_Sample2() {
	var a int = 123

	// reflect.ValueOf() - 通过反射获取指定变量的值
	//   Int(), Float(), String() 等 - 按指定的类型获取值
	//   Interface() - 获取接口，后续可以通过类型断言的方式获取值
	b := reflect.ValueOf(a)
	// 通过反射获取变量的值
	c := b.Int()
	// 通过反射获取变量的值（用断言的方式）
	d, ok := b.Interface().(int)
	fmt.Println(c, d, ok) // 123 123 true

	// 通过反射修改变量的值（下面这种方式会报错，因为传入的是 a 的值，它是不能被寻址的，即不知道 a 的值保存在哪里）
	// reflect.ValueOf(a).SetInt(456)
	// 通过反射修改变量的值（下面这种方式是正确的，因为传入的是 a 的指针，它是可以被寻址的，即知道 a 的值保存在哪里）
	reflect.ValueOf(&a).Elem().SetInt(456)
	fmt.Println(a) // 456

	e := myStruc2{salary: 3500, myStruct1: myStruct1{name: "webabcd", age: 40}, SalaryExpected: 4000}
	f := reflect.ValueOf(e)
	// 通过反射获取结构体的指定成员的值（获取结构体成员的其他方式可以参见上面的 reflect_Sample1() 中的代码，方法是类似的）
	field_salaryExpected := f.FieldByName("SalaryExpected")
	fmt.Println(field_salaryExpected.Float()) // 4000

	// 如果找不到指定的成员则 IsValid() 为 false
	// 如果值是 nil 则 IsNil() 为 true
	fmt.Println(f.FieldByName("xxx").IsValid(), f.FieldByName("ptr").IsNil()) // false true

	// 通过反射修改结构体的指定成员的值
	// 下面这种方式会报错，虽然传入的是 e 的指针，是可被寻址的，但是需要修改的成员是不可导出的，即小写开头的
	// reflect.ValueOf(&e).Elem().FieldByName("salary").SetFloat(5000)
	// 下面这种方式是正确的，也就是说传入的是 e 的指针，是可被寻址的，并且需要修改的成员是可导出的，即大写开头的
	reflect.ValueOf(&e).Elem().FieldByName("SalaryExpected").SetFloat(5000)
	fmt.Println(e.SalaryExpected) // 5000

	// 通过反射调用结构体的方法
	// 拿到方法
	method := reflect.ValueOf(&e).MethodByName("Show")
	// 构造参数
	args := []reflect.Value{reflect.ValueOf("v1"), reflect.ValueOf("v2")}
	// 调用方法并传参，然后获取结果
	g := method.Call(args)[0]
	fmt.Println(g) // name:webabcd, age:40, p1:v1, p2:v2
}

func reflect_Sample3() {
	a := 123
	// 通过反射获取指定变量的类型
	b := reflect.TypeOf(a)
	// 根据反射类型创建类型实例
	c := reflect.New(b)
	// 通过反射修改变量的值
	c.Elem().SetInt(456)
	// 通过反射获取变量的值
	fmt.Println(a, c.Elem().Int()) // 123 456

	// 通过反射调用函数
	// 拿到函数
	method := reflect.ValueOf(myFunc1)
	// 构造参数
	args := []reflect.Value{reflect.ValueOf("v1"), reflect.ValueOf("v2")}
	// 调用函数并传参，然后获取结果
	d := method.Call(args)[0]
	fmt.Println(d) // p1:v1, p2:v2
}

type myStruct1 struct {
	name string
	age  int
}

type myStruc2 struct {
	salary float32
	myStruct1
	SalaryExpected float32
	ptr            *int
}

func (s *myStruc2) Show(p1, p2 string) string {
	return fmt.Sprintf("name:%s, age:%d, p1:%s, p2:%s", s.name, s.age, p1, p2)
}

func myFunc1(p1, p2 string) string {
	return fmt.Sprintf("p1:%s, p2:%s", p1, p2)
}
