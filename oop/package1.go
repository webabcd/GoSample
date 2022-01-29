// 设置当前文件的包名，只需要指定当前文件夹的名称即可，不用写全名
package oop

// go 可以导入内置包（GOROOT 路径）和第三方包（GOPATH 路径）和自定义包（你自己的路径）
// go 的可见性控制：如果定义的常量、变量、类型、接口、接口中的方法，结构体、结构体中的成员、函数等的名称是大写字母开头，则表示能被其它包访问或调用，非大写字母开头就只能在包内使用

/*
// 单行导入
import "fmt"
// 然后就可以通过导入的包名使用这个包下的东西了
fmt.Println("xxx")
*/

/*
// 多行导入
import (
	"fmt"
	"math"
)
*/

/*
// 为导入的包指定别名
import F "fmt"
// 然后就可以通过别名使用这个包下的东西了
F.Println("xxx")
*/

/*
// 以合并的方式导入包
import . "fmt"
// 然后就可以直接使用这个包下的东西了
Println("xxx")
*/

// 导入包时要写全路径
// 导入包时，会先调用包中的文件的 init 方法
// 如果 a 导入了 b，b 导入了 c，则先调用 c 中文件的 init 方法，再调用 b 中文件的 init 方法，最后调用 c 中文件的 init 方法
// 如果同包有多个文件都有 init 方法，则他们的调用顺序不定
import (
	// _ 的意思就是你无法使用这个包（但是会调用这个包中的文件的 init 方法）
	_ "GoSample/oop/pkg1"
	// 注：因为当前包引用了 GoSample/oop/pkg1，而 GoSample/oop/pkg1 引用了 GoSample/oop/pkg1/pkg2，所以当前包如果被引用后则会打印
	// pkg2 init
	// pkg1 init
)

func Package1Sample() {

}
