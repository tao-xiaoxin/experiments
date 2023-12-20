# 探索 GO 项目依赖包管理与Go Module常规操作

[TOC]



## 一.Go 构建模式的演变

Go 程序由 Go 包组合而成的，**Go 程序的构建过程就是确定包版本、编译包以及将编译后得到的目标文件链接在一起的过程。**

Go 构建模式历经了三个迭代和演化过程，分别是最初期的 `GOPATH`、1.5 版本的 `Vendor` 机制，以及现在的 `Go Module`。

### 1.1 GOPATH （初版）

Go 语言在首次开源时，就内置了一种名为 **GOPATH** 的构建模式。

特点：**在这种构建模式下，Go 编译器可以在本地 GOPATH 环境变量配置的路径下，搜寻 Go 程序依赖的第三方包。如果存在，就使用这个本地包进行编译；如果不存在，就会报编译错误**

首先使用`go` 多版本管理工具`gvm` 将 Go 版本到1.10.8:

```bash
# 如果没有安装gvm,使用如下命令安装
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
gvm install go1.10.8 # 使用 GVM 安装 Go 1.10.8 
gvm use go1.10.8 # 切换到 Go 1.10.8 版本
go version # 验证是否成功设置了 Go 1.10.8
```

这里给出了一段在 GOPATH 构建模式下编写的代码：

```go
package main

import "github.com/sirupsen/logrus"
func main() {
    logrus.Println("hello, gopath mode")
}
```

然后使用Go 1.10.8编译执行如下：

```go
# go build main.go
# 直接报错如下：
main.go:3:8: cannot find package "github.com/sirupsen/logrus" in any of:
        /usr/local/go/src/github.com/sirupsen/logrus (from $GOROOT)
        /root/go/src/github.com/sirupsen/logrus (from $GOPATH)
```

**那么 Go 编译器在 GOPATH 构建模式下，究竟怎么在 GOPATH 配置的路径下搜寻第三方依赖包呢？**

为了说清楚搜寻规则，先假定 Go 程序导入了 `github.com/user/repo` 这个包，我们也同时假定当前 GOPATH 环境变量配置的值为：

```bash
export GOPATH=/usr/local/goprojects:/root/go/
```

那么在 `GOPATH` 构建模式下，`Go` 编译器在编译 `Go` 程序时，就会在下面两个路径下搜索第三方依赖包是否存在：

```go
/usr/local/goprojects/src/github.com/user/repo
/root/go/src/github.com/user/repo
```

注意：如果你没有显式设置 `GOPATH` 环境变量，`Go` 会将 `GOPATH` 设置为默认值，不同操作系统下默认值的路径不同，在 `macOS` 或 `Linux` 上，它的默认值是 `$HOME/go`。

当本地找不到第三方依赖包的情况，我们该如何解决这个问题呢？

这个时候就需要让 `go get` 登场了！

#### 1.1.1 go get

在本地没有找到程序的第三方依赖包,可以通过 go get 命令将本地缺失的第三方依赖包下载到本地，比如：

```bash
go get github.com/sirupsen/logrus
```

这里的`go get`命令会下载第三方Go包及其依赖到本地的`GOPATH`目录下。并且`go get` 下载的包只是那个时刻各个依赖包的最新主线版本，这样会给后续 Go 程序的构建带来一些问题。比如，依赖包持续演进，可能会导致不同开发者在不同时间获取和编译同一个 Go 包时，得到不同的结果，也就是不能保证可重现的构建（Reproduceable Build）。又比如，如果依赖包引入了不兼容代码，程序将无法通过编译。

最后还有一点，如果依赖包因引入新代码而无法正常通过编译，并且该依赖包的作者又没用及时修复这个问题，这种错误也会传导到你的程序，导致你的程序无法通过编译。

**在 GOPATH 构建模式下，Go 编译器实质上并没有关注 Go 项目所依赖的第三方包的版本。**但 Go 开发者希望自己的 Go 项目所依赖的第三方包版本能受到自己的控制，而不是随意变化。所以 `Go` 核心开发团队引入了 `Vendor` 机制试图解决上面的问题。

### 1.2 vendor 机制（中版）

Go 在 1.5 版本中引入 **vendor 机制**。**所谓 `vendor` 机制，就是每个项目的根目录下可以有一个 `vendor` 目录，里面存放了该项目的依赖的 `package`。`go build` 的时候会先去 `vendor` 目录查找依赖，如果没有找到会再去 `GOPATH` 目录下查找。**

这样的话，**Go 编译器会优先感知和使用 `vendor` 目录下缓存的第三方包版本，而不是 `GOPATH` 环境变量所配置的路径下的第三方包版本**。这样，无论第三方依赖包自己如何变化，无论 GOPATH 环境变量所配置的路径下的第三方包是否存在、版本是什么，都不会影响到 Go 程序的构建。

**如果使用 `vendor` 机制管理第三方依赖包，最佳实践就是将 `vendor` 一并提交到代码仓库中。那么其他开发者下载你的项目后，就可以实现可重现的构建**。

下面这个目录结构就是为上面的代码示例添加 vendor 目录后的结果：

```shell
.
├── main.go
└── vendor/
    ├── github.com/
    │   └── sirupsen/
    │       └── logrus/
    └── golang.org/
        └── x/
            └── sys/
                └── unix/
```

在添加完 vendor 后，我们重新编译 main.go，这个时候 Go 编译器就会在 vendor 目录下搜索程序依赖的 logrus 包以及后者依赖的 `golang.org/x/sys/unix` 包了.

注意:**要想开启 vendor 机制，你的 Go 项目必须位于 GOPATH 环境变量配置的某个路径的 src 目录下面。如果不满足这一路径要求，那么 Go 编译器是不会理会 Go 项目目录下的 vendor 目录的**

不过 vendor 机制虽然一定程度解决了 Go 程序可重现构建的问题，但对开发者来说，它的体验却不那么好。一方面，`Go` 项目必须放在 `GOPATH` 环境变量配置的路径下，庞大的 `vendor` 目录需要提交到代码仓库，不仅占用代码仓库空间，减慢仓库下载和更新的速度，而且还会干扰代码评审，对实施代码统计等开发者效能工具也有比较大影响。另外，你还需要手工管理 `vendor` 下面的 `Go` 依赖包，包括项目依赖包的分析、版本的记录、依赖包获取和存放等等。

 为解决这个问题，`Go` 核心团队与社区将 Go 构建的重点转移到如何解决包依赖管理上。`Go` 社区先后开发了诸如 `gb`、`glide`、`dep` 等工具，来帮助 `Go` 开发者对 `vendor` 下的第三方包进行自动依赖分析和管理，但这些工具也都有自身的问题。

Go 核心团队基于社区实践的经验和教训，推出了 Go 官方的最新解决方案：`Go Module`。

### 1.3 Go Module（最新版）

Go 1.11 版本推出 `modules` 机制，简称 `mod`，**更加易于管理项目中所需要的模块。**

一个 `Go Module` 是一个 Go 包的集合。`module` 是有版本的，所以 `module `下的包也就有了版本属性。这个 `module` 与这些包会组成一个独立的版本单元，它们一起打版本、发布和分发,。

在 `Go Module` 模式下，通常一个代码仓库对应一个 `Go Module`。一个 `Go Module` 的顶层目录下会放置一个 go.mod 文件，每个 go.mod 文件会定义唯一一个 module，也就是说 `Go Module` 与 `go.mod` 是一一对应的。

并且其根目录中包含 `go.mod` 文件,`go.mod` 文件定义了模块的模块路径，它也是用于根目录的导入路径，以及它的依赖性要求。每个依赖性要求都被写为模块路径和特定语义版本。

`go.mod` 文件所在的顶层目录也被称为 `module` 的根目录，`module` 根目录以及它子目录下的所有 Go 包均归属于这个 Go Module，这个 module 也被称为 **main module**。

从 `Go 1.11` 开始，`Go` 允许在 `$GOPATH/src` 外的任何目录下使用 `go.mod` 创建项目。在 `$GOPATH/src` 中，为了兼容性，`Go` 命令仍然在旧的 `GOPATH` 模式下运行。从 `Go 1.13` 开始，`go.mod`模式将成为默认模式。

## 二.创建Go Module

### 2.1 创建步骤

将基于当前项目创建一个 Go Module，通常有如下几个步骤：

1. 通过 `go mod init [项目地址\库地址]` 创建 go.mod 文件，将当前项目变为一个 Go Module；
2. 通过` go mod tidy` 命令自动更新当前 module 的依赖信息；
3. 执行 `go build`，执行新 module 的构建。

### 2.2 简单举列

新建一个main.go文件,引入外部包 logrus

```go
package main

import "github.com/sirupsen/logrus"
func main() {
  logrus.Println("hello, go module mode")
}
```

我们通过 `go mod ini`t 命令为这个项目创建一个 `Go Modul`e（这里我们使用的是 Go 版本最新版，Go 最新版默认采用 Go Module 构建模式）

```bash
$go mod init github.com/bigwhite/module-mode
go: creating new go.mod: module github.com/bigwhite/module-mode
go: to add module requirements and sums:
  go mod tidy
```

现在，`go mod init` 在当前项目目录下创建了一个 go.mod 文件，这个 `go.mod` 文件将当前项目变为了一个 `Go Module`，项目根目录变成了 module 根目录。`go.mod `的内容是这样的.

```go
module github.com/bigwhite/module-mode
go 1.21.1
```

这个 go.mod 文件现在处于初始状态，它的**第一行内容用于声明 module 路径（module path）**,一般是指定自己项目的git地址，**最后一行是 Go 版本指示符**,表示这个 module 是在某个特定的 Go 版本的 module 语义的基础上编写的。

`go mod init` 命令日志输出提示我们可以使用 `go mod tidy` 命令，添加 module 依赖以及校验和。`go mod tidy` 命令会扫描 Go 源码，并自动找出项目依赖的外部 Go Module 以及版本，下载这些依赖并更新本地的` go.mod` 文件。我们按照这个提示执行一下 `go mod tidy `命令

```bash
$go mod tidy                                   
go: finding module for package github.com/sirupsen/logrus
go: downloading github.com/sirupsen/logrus v1.9.3
go: found github.com/sirupsen/logrus in github.com/sirupsen/logrus v1.9.3
go: downloading golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8
go: downloading github.com/stretchr/testify v1.7.0
go: downloading gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
```

我们看到，对于一个处于初始状态的 module 而言，`go mod tidy` 分析了当前 `main module` 的所有源文件，找出了当前 `main module` 的所有第三方依赖，确定第三方依赖的版本，还下载了当前 main module 的直接依赖包（比如 logrus），以及相关间接依赖包（直接依赖包的依赖，比如上面的 golang.org/x/sys 等）。

由 `go mod tidy` 下载的依赖 module 会被放置在本地的 `module` 缓存路径下，默认值为 `$GOPATH[0]/pkg/mod`，Go 1.15 及以后版本可以通过 `GOMODCACHE` 环境变量，自定义本地 module 的缓存路径。

执行 go mod tidy 后，我们示例 go.mod 的内容更新如下：

```go
module github.com/bigwhite/module-mode
go 1.21.1
require github.com/sirupsen/logrus v1.9.3
require golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
```

可以看到，当前 module 的直接依赖 logrus，还有它的版本信息都被写到了 `go.mod` 文件的 `require` 段中。而且，执行完` go mod tidy `后，当前项目除了 `go.mod `文件外，还多了一个新文件` go.sum`，内容是这样的：

```go
github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/davecgh/go-spew v1.1.1 h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=
github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/pmezard/go-difflib v1.0.0 h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=
github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
github.com/sirupsen/logrus v1.9.3 h1:dueUQJ1C2q9oE3F7wvmSGAaVtTmUizReu6fjN8uqzbQ=
github.com/sirupsen/logrus v1.9.3/go.mod h1:naHLuLoDiP4jHNo9R0sCBMtWGeIprob74mVsIT4qYEQ=
github.com/stretchr/objx v0.1.0/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
github.com/stretchr/testify v1.7.0 h1:nwc3DEeHmmLAfoZucVR881uASk0Mfjw8xYJ99tb5CcY=
github.com/stretchr/testify v1.7.0/go.mod h1:6Fq8oRcR53rry900zMqJjRRixrwX3KX962/h/Wwjteg=
golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 h1:0A+M6Uqn+Eje4kHMK80dtF3JCXC4ykBgQG4Fe06QRhQ=
golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c h1:dUUwHk2QECo/6vqA44rthZ8ie2QXMNeKRTHCNY2nXvo=
gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
```

这同样是由 `go mod` 相关命令维护的一个文件，它存放了特定版本 `module` 内容的哈希值。这是 `Go Module` 的一个安全措施。当将来这里的某个 module 的特定版本被再次下载的时候，go 命令会使用 go.sum 文件中对应的哈希值，和新下载的内容的哈希值进行比对，只有哈希值比对一致才是合法的，这样可以确保你的项目所依赖的 module 内容，不会被恶意或意外篡改。**因此，我推荐你把 go.mod 和 go.sum 两个文件与源码，一并提交到代码版本控制服务器上。**

接下来，我们**只需在当前 module 的根路径下，执行 go build 就可以完成 module 的构建了**！

```bash
$go build
$$ls
go.mod    go.sum    main.go    module-mode
$./module-mode 
INFO[0000] hello, go module mode
```

整个过程的执行步骤是这样：`go build` 命令会读取 `go.mod` 中的依赖及版本信息，并在本地 module 缓存路径下找到对应版本的依赖 `module`，执行编译和链接。如果顺利的话，我们会在当前目录下看到一个新生成的可执行文件 `module-mode`，执行这个文件我们就能得到正确结果了。

## 三.深入理解 Go Module 构建模式

Go 语言设计者在设计 Go Module 构建模式，来解决“包依赖管理”的问题时，进行了几项创新，这其中就包括**语义导入版本 (Semantic Import Versioning)**，以及和其他主流语言不同的**最小版本选择 (Minimal Version Selection)** 等机制。

### 3.1 Go Module 的语义导入版本机制

在上面的例子中，我们看到 go.mod 的 require 段中依赖的版本号，都符合 vX.Y.Z 的格式。在 Go Module 构建模式下，一个符合 Go Module 要求的版本号，由前缀 v 和一个满足语义[版本规范](https://semver.org/)的版本号组成。例如,上面的 logrus module 的版本号是 v1.9.3，这就表示它的主版本号为 1，次版本号为 9，补丁版本号为 3.

语义版本号分成 3 部分：

1. 主版本号 (major)
2. 次版本号 (minor)
3. 补丁版本号 (patch)

![img](https://billy.taoxiaoxin.club/md/2023/09/650c52fc7822fe9180ab038b.png)

Go 命令和 go.mod 文件都使用上面这种符合语义版本规范的版本号，作为描述 Go Module 版本的标准形式。借助于语义版本规范，**Go 命令可以确定同一 module 的两个版本发布的先后次序，而且可以确定它们是否兼容**。

**按照语义版本规范，主版本号不同的两个版本是相互不兼容的。**而且，在主版本号相同的情况下，次版本号大都是向后兼容次版本号小的版本。补丁版本号也不影响兼容性。

而且，Go Module 规定：**如果同一个包的新旧版本是兼容的，那么它们的包导入路径应该是相同的**。

怎么理解呢？我们来举个简单示例。我们就以 logrus 为例，它有很多发布版本，我们从中选出两个版本 v1.7.0 和 v1.8.1.。按照上面的语义版本规则，这两个版本的主版本号相同，新版本 v1.8.1 是兼容老版本 v1.7.0 的。那么，我们就可以知道，如果一个项目依赖 logrus，无论它使用的是 v1.7.0 版本还是 v1.8.1 版本，它都可以使用下面的包导入语句导入 logrus 包：

```bash
import "github.com/sirupsen/logrus"
```

Go Module 创新性地给出了一个方法：将包主版本号引入到包导入路径中，我们可以像下面这样导入 logrus v2.0.0 版本依赖包：

```go
import "github.com/sirupsen/logrus/v2"
```

这就是 Go 的“语义导入版本”机制，也就是说通过在包导入路径中引入主版本号的方式，来区别同一个包的不兼容版本，这样一来我们甚至可以同时依赖一个包的两个不兼容版本：

```go
import (
    "github.com/sirupsen/logrus"
    logv2 "github.com/sirupsen/logrus/v2"
)
```

不过到这里，你可能会问，v0.y.z 版本应该使用哪种导入路径呢？

按照语义版本规范的说法，v0.y.z 这样的版本号是用于项目初始开发阶段的版本号。在这个阶段任何事情都有可能发生，其 API 也不应该被认为是稳定的。**Go Module 将这样的版本 (v0) 与主版本号 v1 做同等对待**，也就是**采用不带主版本号的包导入路径**，这样一定程度降低了 Go 开发人员使用这样版本号包时的心智负担。

Go 语义导入版本机制是 Go Module 机制的基础规则，同样它也是 Go Module 其他规则的基础。

### 3.2 Go Module 的最小版本选择原则

在前面的例子中，Go 命令都是在项目初始状态分析项目的依赖，并且项目中两个依赖包之间没有共同的依赖，这样的包依赖关系解决起来还是比较容易的。但依赖关系一旦复杂起来，比如像下图中展示的这样，Go 又是如何确定使用依赖包 C 的哪个版本的呢？

![img](https://billy.taoxiaoxin.club/md/2023/09/650c52fc03f1f9e05ea6ff6b.jpeg)

在这张图中，myproject 有两个直接依赖 A 和 B，A 和 B 有一个共同的依赖包 C，但 A 依赖 C 的 v1.1.0 版本，而 B 依赖的是 C 的 v1.3.0 版本，并且此时 C 包的最新发布版为 C v1.7.0。这个时候，Go 命令是如何为 myproject 选出间接依赖包 C 的版本呢？

其实，当前存在的主流编程语言，以及 Go Module 出现之前的很多 Go 包依赖管理工具都会选择依赖项的“**最新最大 (Latest Greatest) 版本**”，对应到图中的例子，这个版本就是 v1.7.0。

当然了，理想状态下，如果语义版本控制被正确应用，并且这种“社会契约”也得到了很好的遵守，那么这种选择算法是有道理的，而且也可以正常工作。在这样的情况下，依赖项的“最新最大版本”应该是最稳定和安全的版本，并且应该有向后兼容性。至少在相同的主版本 (Major Verion) 依赖树中是这样的

但我们这个问题的答案并不是这样的。Go 设计者另辟蹊径，在诸多兼容性版本间，他们不光要考虑最新最大的稳定与安全，还要尊重各个 module 的述求：A 明明说只要求 C v1.1.0，B 明明说只要求 C v1.3.0。**所以 Go 会在该项目依赖项的所有版本中，选出符合项目整体要求的“最小版本”**.

这个例子中，C v1.3.0 是符合项目整体要求的版本集合中的版本最小的那个，于是 Go 命令选择了 C v1.3.0，而不是最新最大的 C v1.7.0。并且，Go 团队认为“最小版本选择”为 Go 程序实现持久的和可重现的构建提供了最佳的方案。

即:对于导入路径不同的包，则两个包是同时被依赖的，导入路径相同，则只能选择依赖一个，并且会选择所有直接间接依赖这个包的版本的最高版本，而不是该包本身的最高版本，这是所有依赖这个包的其他包都能接受的最小版本，这样可以保证服务整体的稳定性。

### 3.3 Go 各版本构建模式机制和切换

在 `Go 1.11` 版本中，Go 开发团队引入 `Go Modules` 构建模式。这个时候，`GOPATH` 构建模式与 `Go Modules `构建模式各自独立工作，我们可以通过设置环境变量 `GO111MODULE` 的值在两种构建模式间切换。

然后，随着 Go 语言的逐步演进，从 `Go 1.11` 到 `Go 1.16` 版本，不同的 Go 版本在 `GO111MODULE` 为不同值的情况下，开启的构建模式几经变化，直到 `Go 1.16` 版本，`Go Module` 构建模式成为了默认模式。

所以，要分析 Go 各版本的具体构建模式的机制和切换，我们只需要找到这几个代表性的版本就好了。

我这里将 `Go 1.13` 版本之前、`Go 1.13` 版本以及 `Go 1.16` 版本，在 `GO111MODULE` 为不同值的情况下的行为做了一下对比，这样我们可以更好地理解不同版本下、不同构建模式下的行为特性，下面我们就来用表格形式做一下比对：

![img](https://billy.taoxiaoxin.club/md/2023/09/650c52fc88761110c83d0c3b.jpeg)

## 四.设置 GO111MODULE

要启用go module支持首先要设置环境变量GO111MODULE，通过它可以开启或关闭模块支持，它有三个可选值：off、on、auto，默认值是auto。

- GO111MODULE=off禁用模块支持，编译时会从GOPATH和vendor文件夹中查找包。
- GO111MODULE=on启用模块支持，编译时会忽略GOPATH和vendor文件夹，只根据 go.mod下载依赖。
- GO111MODULE=auto，当项目在$GOPATH/src外且项目根目录有go.mod文件时，开启模块支持。

设置Go Model

```bash
# 临时开启 Go modules 功能
export GO111MODULE=on
# 永久开启 Go modules 功能
go env -w GO111MODULE=on
```

## 五.Go module 常用操作

### 5.1初始化项目

基于当前项目创建一个 Go Module，通常有如下几个步骤：

1. 通过 `go mod init` 项目名 创建 `go.mod` 文件，将当前项目变为一个` Go Module`；
2. 通过 `go mod tidy` 命令自动更新当前 `module` 的依赖信息；
3. 执行 `go build`，执行新 `module` 的构建。

然后会生成两个文件`go.mod`和`go.sum`.

#### 5.1.1 go.mod

go.mod文件记录了项目所有的依赖信息，其结构大致如下：

```bash
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

- `module`用来模块名称
- `require`用来定义依赖包及版本
- `exclude` 禁止依赖包列表，不下载和引用哪些包(仅在当前模块为主模块时生效)
- `replace` 替换依赖包列表和引用路径(仅在当前模块为主模块时生效)
- `indirect` 表示这个库是间接引用进来的。

#### 5.1.2 replace

在国内访问golang.org/x的各个包都需要翻墙，你可以在go.mod中使用replace替换成github上对应的库。

```bash
replace (
	golang.org/x/net => github.com/golang/net latest
	golang.org/x/tools => github.com/golang/tools latest
	golang.org/x/crypto => github.com/golang/crypto latest
	golang.org/x/sys => github.com/golang/sys latest
	golang.org/x/text => github.com/golang/text latest
	golang.org/x/sync => github.com/golang/sync latest
)
```

#### 5.1.3 go 查看当前项目所有包依赖

- 使用 `go list -m all` 可以查看到所有依赖列表，也可以使用 `go list -json -m all `输出 json格式的打印结果。

### 5.2 升级 / 降级版本

首先,查看依赖历史版本

```
$go list -m -versions github.com/sirupsen/logrus
github.com/sirupsen/logrus v0.1.0 v0.1.1 v0.2.0 v0.3.0 v0.4.0 v0.4.1 v0.5.0 v0.5.1 v0.6.0 v0.6.1 v0.6.2 v0.6.3 v0.6.4 v0.6.5 v0.6.6 v0.7.0 v0.7.1 v0.7.2 v0.7.3 v0.8.0 v0.8.1 v0.8.2 v0.8.3 v0.8.4 v0.8.5 v0.8.6 v0.8.7 v0.9.0 v0.10.0 v0.11.0 v0.11.1 v0.11.2 v0.11.3 v0.11.4 v0.11.5 v1.0.0 v1.0.1 v1.0.3 v1.0.4 v1.0.5 v1.0.6 v1.1.0 v1.1.1 v1.2.0 v1.3.0 v1.4.0 v1.4.1 v1.4.2 v1.5.0 v1.6.0 v1.7.0 v1.7.1 v1.8.0 v1.8.1
```

我们可以在项目的 module 根目录下，执行带有版本号的 go get 命令：

```bash
$go get github.com/sirupsen/logrus@v1.7.0
go: downloading github.com/sirupsen/logrus v1.7.0
go get: downgraded github.com/sirupsen/logrus v1.8.1 => v1.7.0
```

当然我们也可以使用万能命令 go mod tidy来帮助我们降级，但前提是首先要用 go mod edit 命令，明确告知我们要依赖 v1.7.0 版本，而不是 v1.8.1，这个执行步骤是这样的：

```bash
$go mod edit -require=github.com/sirupsen/logrus@v1.7.0
$go mod tidy       
go: downloading github.com/sirupsen/logrus v1.7.0
```

升级版本和降级版本依赖一样,参照上面的操作即可,

### 5.3 删除未使用的依赖

可以用 go mod tidy 命令来清除这些没用到的依赖项：

```bash
go mod tidy
```

go mod tidy会自动分析源码依赖，而且将不再使用的依赖从 go.mod 和 go.sum 中移除。

### 5.4 引入主版本号大于 1 的三方库

语义导入版本机制有一个原则：**如果新旧版本的包使用相同的导入路径，那么新包与旧包是兼容的**。也就是说，如果新旧两个包不兼容，那么我们就应该采用不同的导入路径。

按照语义版本规范，如果我们要为项目引入主版本号大于 1 的依赖，比如 v2.0.0，那么由于这个版本与 v1、v0 开头的包版本都不兼容，我们在导入 v2.0.0 包时，不能再直接使用 github.com/user/repo，而要使用像下面代码中那样不同的包导入路径：

```bash
import github.com/user/repo/v2/xxx
```

也就是说，如果我们要为 Go 项目添加主版本号大于 1 的依赖，我们就需要使用“语义导入版本”机制，在声明它的导入路径的基础上，加上版本号信息。我们以“向 module-mode 项目添加 github.com/go-redis/redis 依赖包的 v7 版本”为例，看看添加步骤。

首先，我们在源码中，以空导入的方式导入 v7 版本的 github.com/go-redis/redis 包：

```bash
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

```bash
$go get github.com/go-redis/redis/v7
go: downloading github.com/go-redis/redis/v7 v7.4.1
go: downloading github.com/go-redis/redis v6.15.9+incompatible
go get: added github.com/go-redis/redis/v7 v7.4.1
```

### 5.5 升级依赖版本到一个不兼容版本

我们前面说了，按照语义导入版本的原则，不同主版本的包的导入路径是不同的。所以，同样地，我们这里也需要先将代码中 redis 包导入路径中的版本号改为 v8：

```bash
import (
  _ "github.com/go-redis/redis/v8"
  "github.com/google/uuid"
  "github.com/sirupsen/logrus"
)
```

接下来，我们再通过 go get 来获取 v8 版本的依赖包：

```bash
$go get github.com/go-redis/redis/v8
go: downloading github.com/go-redis/redis/v8 v8.11.1
go: downloading github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f
go: downloading github.com/cespare/xxhash/v2 v2.1.1
go get: added github.com/go-redis/redis/v8 v8.11.1
```

### 5.6 特殊情况：使用 vendor

vendor 机制虽然诞生于 GOPATH 构建模式主导的年代，但在 Go Module 构建模式下，它依旧被保留了下来，并且成为了 Go Module 构建机制的一个很好的补充。特别是在一些不方便访问外部网络，并且对 Go 应用构建性能敏感的环境。

和GOPATH 构建模式不同，Go Module 构建模式下，我们再也无需手动维护 vendor 目录下的依赖包了，Go 提供了可以快速建立和更新 vendor 的命令，我们还是以前面的module-mode 项目为例，通过下面命令为该项目建立 vendor：

```bash
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

## 六、Go module 常用命令总结

### 6.1 常用 go mod命令

常用的go mod命令如下：

```bash
go mod download    下载依赖的module到本地cache（默认为$GOPATH/pkg/mod目录）
go mod edit        编辑go.mod文件
go mod graph       打印模块依赖图
go mod init        初始化当前文件夹, 创建go.mod文件
go mod tidy        增加缺少的module，删除无用的module
go mod vendor      将依赖复制到vendor下
go mod verify      校验依赖
go mod why         解释为什么需要依赖
```

### 6.2 go get命令

在项目中执行go get命令可以下载依赖包，并且还可以指定下载的版本。

1. 运行go get -u将会升级到最新的次要版本或者修订版本(x.y.z, z是修订版本号， y是次要版本号)
2. 运行go get -u=patch将会升级到最新的修订版本
3. 运行go get package@version将会升级到指定的版本号version

如果下载所有依赖可以使用go mod download命令。

### 6.3 go mod edit

#### 6.3.1 格式化

因为我们可以手动修改go.mod文件，所以有些时候需要格式化该文件。Go提供了一下命令：

```bash
go mod edit -fmt
```

#### 6.3.2 添加依赖项

```bash
go mod edit -require=golang.org/x/text
```

#### 6.3.3 移除依赖项

如果只是想修改go.mod文件中的内容，那么可以运行`go mod edit -droprequire=package path`，比如要在go.mod中移除golang.org/x/text包，可以使用如下命令：

```bash
go mod edit -droprequire=golang.org/x/text
```

关于go mod edit的更多用法可以通过go help mod edit查看。

## 七、Go Module 代理

###  7.1 GO 设置代理

### 7.1.1 打开模块支持

```bash
go env -w GO111MODULE=on
```

### 7.1.2 取消代理

```bash
go env -w GOPROXY=direct
```

### 7.1.3 关闭包的有效性验证

```bash
go env -w GOSUMDB=off
```

### 7.1.4 设置不走 proxy 的私有仓库或组，多个用逗号相隔（可选）

```
go env -w GOPRIVATE=git.mycompany.com,github.com/my/private
```

### 7.1.5 国内常用代理列表

| 提供者       | 地址                                |
| ------------ | ----------------------------------- |
| 官方全球代理 | https://proxy.golang.com.cn         |
| 七牛云       | https://goproxy.cn                  |
| 阿里云       | https://mirrors.aliyun.com/goproxy/ |
| GoCenter     | https://gocenter.io                 |
| 百度         | https://goproxy.bj.bcebos.com/      |

**“direct”** 为特殊指示符，用于指示 Go 回源到模块版本的源地址去抓取(比如 GitHub 等)，当值列表中上一个 Go module proxy 返回 404 或 410 错误时，Go 自动尝试列表中的下一个，遇见 **“direct”** 时回源，遇见 EOF 时终止并抛出类似 “invalid version: unknown revision…” 的错误。

### 7.1.6 官方全球代理

```bash
go env -w GOPROXY=https://proxy.golang.com.cn,direct
go env -w GOPROXY=https://goproxy.io,direct
go env -w GOSUMDB=gosum.io+ce6e7565+AY5qEHUk/qmHc5btzW45JVoENfazw8LielDsaI+lEbq6
go env -w GOSUMDB=sum.golang.google.cn
```

### 7.5.3 七牛云

```bash
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOSUMDB=goproxy.cn/sumdb/sum.golang.org
```

### 7.5.4 阿里云

```bash
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
# GOSUMDB 不支持
```

### 7.5.6 GoCenter

```bash
go env -w GOPROXY=https://gocenter.io,direct
# 不支持 GOSUMDB
```

### 7.5.7 百度

```bash
go env -w GOPROXY=https://goproxy.bj.bcebos.com/,direct
# 不支持 GOSUMDB
```

### 7.6 Goland设置代理

![img](https://billy.taoxiaoxin.club/md/2023/09/651055e21d5b9695b0f669d3.png)



## 八.项目中使用Go module

### 8.1 既有项目

如果需要对一个已经存在的项目启用go module，可以按照以下步骤操作：

1. 在项目目录下执行go mod init，生成一个go.mod文件。
2. 执行go get，查找并记录当前项目的依赖，同时生成一个go.sum记录每个依赖库的版本和哈希值。

### 8.2 新项目

对于一个新创建的项目，我们可以在项目文件夹下按照以下步骤操作：

1. 执行go mod init 项目名命令，在当前项目文件夹下创建一个go.mod文件。
2. 手动编辑go.mod中的require依赖项或执行go get自动发现、维护依赖。

## 九、查看Go的配置

```bash
$ go env
//以JSON格式输出
$ go env -json
```

## 十、参考资料

- [Go Modules 使用详解（设置GO111MODULE、设置代理、初始化项目、创建依赖项、回退依赖项版本、删除未使用的依赖项](https://blog.csdn.net/wohu1104/article/details/110505489)

