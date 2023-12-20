##  一.GO 设置代理

### 1.1 打开模块支持

```
go env -w GO111MODULE=on
```

### 1.2 取消代理

```
go env -w GOPROXY=direct
```

### 1.3 关闭包的有效性验证

```
go env -w GOSUMDB=off
```

### 1.4 设置不走 proxy 的私有仓库或组，多个用逗号相隔（可选）

```
go env -w GOPRIVATE=git.mycompany.com,github.com/my/private
```

### 1.5 设置代理

### 1.5.1国内常用代理列表

| 提供者       | 地址                                |
| ------------ | ----------------------------------- |
| 官方全球代理 | https://proxy.golang.com.cn         |
| 七牛云       | https://goproxy.cn                  |
| 阿里云       | https://mirrors.aliyun.com/goproxy/ |
| GoCenter     | https://gocenter.io                 |
| 百度         | https://goproxy.bj.bcebos.com/      |

**“direct”** 为特殊指示符，用于指示 Go 回源到模块版本的源地址去抓取(比如 GitHub 等)，当值列表中上一个 Go module proxy 返回 404 或 410 错误时，Go 自动尝试列表中的下一个，遇见 **“direct”** 时回源，遇见 EOF 时终止并抛出类似 “invalid version: unknown revision…” 的错误。

### 1.5.2 官方全球代理

```
go env -w GOPROXY=https://proxy.golang.com.cn,direct
go env -w GOPROXY=https://goproxy.io,direct
go env -w GOSUMDB=gosum.io+ce6e7565+AY5qEHUk/qmHc5btzW45JVoENfazw8LielDsaI+lEbq6
go env -w GOSUMDB=sum.golang.google.cn
```

### 1.5.3 七牛云

```
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOSUMDB=goproxy.cn/sumdb/sum.golang.org
```

### 1.5.4 阿里云

```
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
# GOSUMDB 不支持
```

### 1.5.6 GoCenter

```
go env -w GOPROXY=https://gocenter.io,direct
# 不支持 GOSUMDB
```

### 1.5.7 百度

```
go env -w GOPROXY=https://goproxy.bj.bcebos.com/,direct
# 不支持 GOSUMDB
```

### 1.6 Goland设置代理

![img](https://billy.taoxiaoxin.club/md/2023/09/650c5797ab14b22c5324ec21.png)



### 1.7 查看Go的配置

```
$ go env
//以JSON格式输出
$ go env -json
```

## 二.设置 GO111MODULE

要启用go module支持首先要设置环境变量GO111MODULE，通过它可以开启或关闭模块支持，它有三个可选值：off、on、auto，默认值是auto。

- GO111MODULE=off禁用模块支持，编译时会从GOPATH和vendor文件夹中查找包。
- GO111MODULE=on启用模块支持，编译时会忽略GOPATH和vendor文件夹，只根据 go.mod下载依赖。
- GO111MODULE=auto，当项目在$GOPATH/src外且项目根目录有go.mod文件时，开启模块支持。

设置Go Model

```
# 临时开启 Go modules 功能
export GO111MODULE=on
# 永久开启 Go modules 功能
go env -w GO111MODULE=on
```

## 三.godep 基本使用

Go语言从v1.5开始开始引入vendor模式，如果项目目录下有vendor目录，那么go工具链会优先使用vendor内的包进行编译、测试等。

godep是一个通过vender模式实现的Go语言的第三方依赖管理工具，类似的还有由社区维护准官方包管理工具dep。

### 3.1 安装

执行以下命令安装godep工具。

```
go get github.com/tools/godep
```

### 3.2 基本命令

安装好godep之后，在终端输入godep查看支持的所有命令。

```
godep save     将依赖项输出并复制到Godeps.json文件中
godep go       使用保存的依赖项运行go工具
godep get      下载并安装具有指定依赖项的包
godep path     打印依赖的GOPATH路径
godep restore  在GOPATH中拉取依赖的版本
godep update   更新选定的包或go版本
godep diff     显示当前和以前保存的依赖项集之间的差异
godep version  查看版本信息
```

使用godep help [command]可以看看具体命令的帮助信息。

### 3.3 使用godep

在项目目录下执行godep save命令，会在当前项目中创建Godeps和vender两个文件夹。

其中Godeps文件夹下有一个Godeps.json的文件，里面记录了项目所依赖的包信息。 vender文件夹下是项目依赖的包的源代码文件。

### 3.4 vender机制

Go1.5版本之后开始支持，能够控制Go语言程序编译时依赖包搜索路径的优先级。

例如查找项目的某个依赖包，首先会在项目根目录下的vender文件夹中查找，如果没有找到就会去$GOAPTH/src目录下查找。

### 3.5 godep开发流程

1. 保证程序能够正常编译
2. 执行godep save保存当前项目的所有第三方依赖的版本信息和代码
3. 提交Godeps目录和vender目录到代码库。
4. 如果要更新依赖的版本，可以直接修改Godeps.json文件中的对应项

## 四.Go module 常用操作

### 4.1初始化项目

基于当前项目创建一个 Go Module，通常有如下几个步骤：

1. 通过 go mod init 项目名 创建 go.mod 文件，将当前项目变为一个 Go Module；
2. 通过 go mod tidy 命令自动更新当前 module 的依赖信息；
3. 执行 go build，执行新 module 的构建。

然后会生成两个文件go.mod和go.sum.

### go.mod

go.mod文件记录了项目所有的依赖信息，其结构大致如下：

```
module dome
go 1.18
require (
	github.com/google/uuid v1.3.0
	github.com/sirupsen/logrus v1.9.0
)

require (
	github.com/kr/fs v0.1.0 // indirect 
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/tools/godep v0.0.0-20180126220526-ce0bfadeb516 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/tools v0.4.0 // indirect
)
```

其中，

- module用来模块名称
- require用来定义依赖包及版本
- exclude 禁止依赖包列表，不下载和引用哪些包(仅在当前模块为主模块时生效)
- replace 替换依赖包列表和引用路径(仅在当前模块为主模块时生效)
- indirect 表示这个库是间接引用进来的。

### replace

在国内访问golang.org/x的各个包都需要翻墙，你可以在go.mod中使用replace替换成github上对应的库。

```
replace (
	golang.org/x/net => github.com/golang/net latest
	golang.org/x/tools => github.com/golang/tools latest
	golang.org/x/crypto => github.com/golang/crypto latest
	golang.org/x/sys => github.com/golang/sys latest
	golang.org/x/text => github.com/golang/text latest
	golang.org/x/sync => github.com/golang/sync latest
)
```

### go 查看当前项目所有包依赖

- 使用 go list -m all 可以查看到所有依赖列表，也可以使用 go list -json -m all 输出 json格式的打印结果。

### 4.2 升级 / 降级版本

首先,查看依赖历史版本

```
$go list -m -versions github.com/sirupsen/logrus
github.com/sirupsen/logrus v0.1.0 v0.1.1 v0.2.0 v0.3.0 v0.4.0 v0.4.1 v0.5.0 v0.5.1 v0.6.0 v0.6.1 v0.6.2 v0.6.3 v0.6.4 v0.6.5 v0.6.6 v0.7.0 v0.7.1 v0.7.2 v0.7.3 v0.8.0 v0.8.1 v0.8.2 v0.8.3 v0.8.4 v0.8.5 v0.8.6 v0.8.7 v0.9.0 v0.10.0 v0.11.0 v0.11.1 v0.11.2 v0.11.3 v0.11.4 v0.11.5 v1.0.0 v1.0.1 v1.0.3 v1.0.4 v1.0.5 v1.0.6 v1.1.0 v1.1.1 v1.2.0 v1.3.0 v1.4.0 v1.4.1 v1.4.2 v1.5.0 v1.6.0 v1.7.0 v1.7.1 v1.8.0 v1.8.1
```

我们可以在项目的 module 根目录下，执行带有版本号的 go get 命令：

```
$go get github.com/sirupsen/logrus@v1.7.0
go: downloading github.com/sirupsen/logrus v1.7.0
go get: downgraded github.com/sirupsen/logrus v1.8.1 => v1.7.0
```

当然我们也可以使用万能命令 go mod tidy来帮助我们降级，但前提是首先要用 go mod edit 命令，明确告知我们要依赖 v1.7.0 版本，而不是 v1.8.1，这个执行步骤是这样的：

```
$go mod edit -require=github.com/sirupsen/logrus@v1.7.0
$go mod tidy       
go: downloading github.com/sirupsen/logrus v1.7.0
```

升级版本和降级版本依赖一样,参照上面的操作即可,

### 4.3 删除未使用的依赖

可以用 go mod tidy 命令来清除这些没用到的依赖项：

```
go mod tidy
```

go mod tidy会自动分析源码依赖，而且将不再使用的依赖从 go.mod 和 go.sum 中移除。

### 4.4 引入主版本号大于 1 的三方库

语义导入版本机制有一个原则：**如果新旧版本的包使用相同的导入路径，那么新包与旧包是兼容的**。也就是说，如果新旧两个包不兼容，那么我们就应该采用不同的导入路径。

按照语义版本规范，如果我们要为项目引入主版本号大于 1 的依赖，比如 v2.0.0，那么由于这个版本与 v1、v0 开头的包版本都不兼容，我们在导入 v2.0.0 包时，不能再直接使用 github.com/user/repo，而要使用像下面代码中那样不同的包导入路径：

```
import github.com/user/repo/v2/xxx
```

也就是说，如果我们要为 Go 项目添加主版本号大于 1 的依赖，我们就需要使用“语义导入版本”机制，在声明它的导入路径的基础上，加上版本号信息。我们以“向 module-mode 项目添加 github.com/go-redis/redis 依赖包的 v7 版本”为例，看看添加步骤。

首先，我们在源码中，以空导入的方式导入 v7 版本的 github.com/go-redis/redis 包：

```
package main

import (
  _ "github.com/go-redis/redis/v7" // “_”为空导入
  "github.com/google/uuid"
  "github.com/sirupsen/logrus"
)

func main() {
  logrus.Println("hello, go module mode")
  logrus.Println(uuid.NewString())
}
```

接下来我们通过 go get 获取redis的v7版本：

```
$go get github.com/go-redis/redis/v7
go: downloading github.com/go-redis/redis/v7 v7.4.1
go: downloading github.com/go-redis/redis v6.15.9+incompatible
go get: added github.com/go-redis/redis/v7 v7.4.1
```

### 4.5 升级依赖版本到一个不兼容版本

我们前面说了，按照语义导入版本的原则，不同主版本的包的导入路径是不同的。所以，同样地，我们这里也需要先将代码中 redis 包导入路径中的版本号改为 v8：

```
import (
  _ "github.com/go-redis/redis/v8"
  "github.com/google/uuid"
  "github.com/sirupsen/logrus"
)
```

接下来，我们再通过 go get 来获取 v8 版本的依赖包：

```
$go get github.com/go-redis/redis/v8
go: downloading github.com/go-redis/redis/v8 v8.11.1
go: downloading github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f
go: downloading github.com/cespare/xxhash/v2 v2.1.1
go get: added github.com/go-redis/redis/v8 v8.11.1
```

### 4.6 特殊情况：使用 vendor

vendor 机制虽然诞生于 GOPATH 构建模式主导的年代，但在 Go Module 构建模式下，它依旧被保留了下来，并且成为了 Go Module 构建机制的一个很好的补充。特别是在一些不方便访问外部网络，并且对 Go 应用构建性能敏感的环境。

和GOPATH 构建模式不同，Go Module 构建模式下，我们再也无需手动维护 vendor 目录下的依赖包了，Go 提供了可以快速建立和更新 vendor 的命令，我们还是以前面的module-mode 项目为例，通过下面命令为该项目建立 vendor：

```
$go mod vendor
$tree -LF 2 vendor
vendor
├── github.com/
│   ├── google/
│   ├── magefile/
│   └── sirupsen/
├── golang.org/
│   └── x/
└── modules.txt
```

我们看到，go mod vendor 命令在 vendor 目录下，创建了一份这个项目的依赖包的副本，并且通过 vendor/modules.txt 记录了 vendor 下的 module 以及版本。

如果我们要基于 vendor 构建，而不是基于本地缓存的 Go Module 构建，我们需要在 go build 后面加上 -mod=vendor 参数。在 Go 1.14 及以后版本中，如果 Go 项目的顶层目录下存在 vendor 目录，那么 go build 默认也会优先基于 vendor构建，除非你给 go build 传入-mod=mod的参数。

通常我们直接使用 go module (非vendor) 模式即可满足大部分需求。如果是那种开发环境受限，因无法访问外部代理而无法通过 go 命令自动解决依赖和下载依赖的环境下，我们通过 vendor 来辅助解决。

## 五.Go module 常用命令总结

### 5.1 go mod命令

常用的go mod命令如下：

```
go mod download    下载依赖的module到本地cache（默认为$GOPATH/pkg/mod目录）
go mod edit        编辑go.mod文件
go mod graph       打印模块依赖图
go mod init        初始化当前文件夹, 创建go.mod文件
go mod tidy        增加缺少的module，删除无用的module
go mod vendor      将依赖复制到vendor下
go mod verify      校验依赖
go mod why         解释为什么需要依赖
```

### 5.2 go get命令

在项目中执行go get命令可以下载依赖包，并且还可以指定下载的版本。

1. 运行go get -u将会升级到最新的次要版本或者修订版本(x.y.z, z是修订版本号， y是次要版本号)
2. 运行go get -u=patch将会升级到最新的修订版本
3. 运行go get package@version将会升级到指定的版本号version

如果下载所有依赖可以使用go mod download命令。

### 5.3 go mod edit

### 格式化

因为我们可以手动修改go.mod文件，所以有些时候需要格式化该文件。Go提供了一下命令：

```
go mod edit -fmt
```

### 添加依赖项

```
go mod edit -require=golang.org/x/text
```

### 移除依赖项

如果只是想修改go.mod文件中的内容，那么可以运行go mod edit -droprequire=package path，比如要在go.mod中移除golang.org/x/text包，可以使用如下命令：

```
go mod edit -droprequire=golang.org/x/text
```

关于go mod edit的更多用法可以通过go help mod edit查看。

## 六.项目中使用Go module

### 既有项目

如果需要对一个已经存在的项目启用go module，可以按照以下步骤操作：

1. 在项目目录下执行go mod init，生成一个go.mod文件。
2. 执行go get，查找并记录当前项目的依赖，同时生成一个go.sum记录每个依赖库的版本和哈希值。

### 新项目

对于一个新创建的项目，我们可以在项目文件夹下按照以下步骤操作：

1. 执行go mod init 项目名命令，在当前项目文件夹下创建一个go.mod文件。
2. 手动编辑go.mod中的require依赖项或执行go get自动发现、维护依赖。

### 参考资料

- [Go Modules 使用详解（设置GO111MODULE、设置代理、初始化项目、创建依赖项、回退依赖项版本、删除未使用的依赖项](https://blog.csdn.net/wohu1104/article/details/110505489)