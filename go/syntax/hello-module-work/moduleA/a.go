package moduleA

import (
	"fmt"
	"gitee.com/tao-xiaoxin/study-basic-go/hello-module-work/moduleB"
)

func init() {
	result := moduleB.Hello()
	fmt.Println(result)
}
func ModuleA() {
	fmt.Println("This is module A！")
}
