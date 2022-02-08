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

package main

import (
	"GoSample/advanced"
	"GoSample/basic"
	"GoSample/container"
	"GoSample/oop"
)

func main() {

	summary()

	basic.DataTypeSample()
	basic.StatementSample()
	basic.PointerSample()
	basic.ReceiverSample()

	container.ArraySample()
	container.SliceSample()
	container.MapSample()
	container.ListSample()

	oop.Function1Sample()
	oop.Function2Sample()
	oop.Function3Sample()
	oop.Struct1Sample()
	oop.Struct2Sample()
	oop.Struct3Sample()
	oop.Struct4Sample()
	oop.Package1Sample()
	oop.Interface1Sample()
	oop.Interface2Sample()

	advanced.Goroutine1Sample()
	advanced.Goroutine2Sample()
	advanced.ReflectSample()
}
