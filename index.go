// go 安装
// 1、https://code.visualstudio.com/ 下载并安装 visual studio code
// 2、在 visual studio code 中安装 Go 扩展
//    如果这部分安装失败，则
//    a) 开启 vpn
//    b) 运行如下命令
//       go env -w GO111MODULE=on
//       go env -w GOPROXY=https://goproxy.cn,direct
// 3、https://golang.google.cn/dl/ 下载并安装 Go
// 4、如果运行报错 go: cannot determine module path for source directory D:\gitroot\GoSample (outside GOPATH, module path must be specified)，则
//    运行 go mod init GoSample（其中的 GoSample 是你文件夹的名字）

// http://c.biancheng.net/view/91.html
// func 大小写的说明，同包名下互相可见所以不能重名
// 泛型
// 写接口的时候别忘了 http://c.biancheng.net/view/58.html 以及 http://c.biancheng.net/view/62.html
// 要写 try/catch http://c.biancheng.net/view/64.html
// 单元测试 http://c.biancheng.net/view/5409.html
// Go语言没有像其它语言一样有public、protected、private等访问控制修饰符，它是通过字母大小写来控制可见性的，如果定义的常量、变量、类型、接口、结构、函数等的名称是大写字母开头表示能被其它包访问或调用（相当于public），非大写开头就只能在包内使用（相当于private，变量或常量也可以下划线开头）

package main

import "GoSample/basic"

func main() {

	// summary()

	// basic.DataTypeSample()
	// basic.StatementSample()
	basic.PointerSample()

	// container.ArraySample()
	// container.SliceSample()
	// container.MapSample()
	// container.ListSample()

	// oop.Function1Sample()
	// oop.Function2Sample()
	// oop.Function3Sample()
	// oop.Struct1Sample()
}
