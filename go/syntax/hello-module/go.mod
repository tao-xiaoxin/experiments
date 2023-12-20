module gitee.com/tao-xiaoxin/study-basic-go/hello-module

go 1.21.1

replace (
	gitee.com/user/moduleA v1.0.0 => ../moduleA
	gitee.com/user/moduleB v1.0.0 => ../moduleB
)
