package pkg1

import (
	// _ 的意思就是你无法使用这个包（但是会调用这个包中的文件的 init 方法）
	_ "GoSample/oop/pkg1/pkg2"
	"fmt"
)

func init() {
	fmt.Println("pkg1 init")
}
