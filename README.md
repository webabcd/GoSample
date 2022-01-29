# Go Sample


#### go 概述
1. 基本语法，nil
   - summary.go

#### go 基础
1. 基本数据类型（数字类型，字符串类型，字符串和数字类型的转换，字符串的格式化，枚举类型，类型别名，自定义类型）
   - basic/datatype.go
2. 语句（if/else, switch, for, continue, break, goto）
   - basic/statement.go
3. 指针
   - basic/pointer.go
4. 通过接收器（receiver）为任意类型定义方法，即为指定类型扩展方法
   - basic/receiver.go

#### go 容器
1. 数组（一维数组，多维数组，数组元素的获取和设置，数组的遍历）
   - container/array.go
2. 切片（通过范围获取切片，切片的构造，切片数据的添加/复制/删除，数组和切片的区别）
   - container/slice.go
3. 字典（字典的增删改查和遍历，以及字典的线程安全）
   - container/map.go
4. 列表（列表的增删改查和遍历）
   - container/list.go

#### go 面向对象
1. 函数（函数基础，带变量名的返回值，多返回值，可变参数，传参时指针和非指针的区别）
   - oop/function1.go
2. 函数（函数也是一种类型，匿名函数，闭包）
   - oop/function2.go
3. 函数函数（defer）
   - oop/function3.go
4. 结构体（定义结构体，声明结构体，初始化结构体，使用结构体，匿名结构体）
   - oop/struct1.go
5. 结构体（内嵌结构体，内嵌结构体实现类似继承的效果，内嵌匿名结构体，内嵌类型）
   - oop/struct2.go
6. 结构体（为结构体定义方法，使用工厂模式初始化结构体）
   - oop/struct3.go
7. 包
   - oop/package1.go
   - oop/pkg1/go.go
   - oop/pkg1/pkg2/go.go
8. 接口（接口的定义，接口的实现，接口的使用）
   - oop/interface1.go
9. 接口（空接口，类型断言，通过实现内置的 error 接口实现自定义错误类型）
   - oop/interface2.go