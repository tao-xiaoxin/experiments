# Go 语言开发环境搭建

[TOC]



## 一. GO 环境安装

### 1.1 下载

Go官网：https://golang.org/dl/

Go镜像站（推荐）：https://golang.google.cn/dl/

### 1.2 Go 版本的选择

默认下载最新自己对应的平台即可。

![image-20230920222329866](https://billy.taoxiaoxin.club/md/2023/09/650b0062bd7d4b03481ce2f7.png)

### 1.3 安装

#### 1.3.1 Windows安装

打开下载的安装程序（.msi文件），然后按照安装向导中的说明进行安装

![image-20230920223507929](https://billy.taoxiaoxin.club/md/2023/09/650b031c259a79736d9f39ae.png)

![image-20230920223737470](https://billy.taoxiaoxin.club/md/2023/09/650b03b162ed5f71b4c07c07.png)

#### 1.3.2 Linux下安装

```bash
wget https://golang.google.cn/dl/go1.21.1.linux-amd64.tar.gz
```

将下载好的文件解压到`/usr/local`目录下：

```bash
tar -zxvf go1.21.1.linux-amd64.tar.gz -C /usr/local  # 解压
```

如果提示没有权限，加上`sudo`以root用户的身份再运行。执行完就可以在`/usr/local/`下看到`go`目录了。

配置环境变量,Linux下有两个文件可以配置环境变量，

+ 其中`/etc/profile`是对所有用户生效的；

+ `$HOME/.profile`是对当前用户生效的，根据自己的情况自行选择一个文件打开，添加如下两行代码，保存退出。

```bash
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
```

修改`/etc/profile`后要重启生效，修改`$HOME/.profile`后使用source命令加载`$HOME/.profile`文件即可生效。

#### 1.3.2 Mac下安装

打开下载的安装程序（.pkg文件），然后按照安装向导中的说明进行安装，默认会将go安装到`/usr/local/go`目录下。

![image-20230920224348700](https://billy.taoxiaoxin.club/md/2023/09/650b05241ac89bcbf7b2dd7b.png)

### 1.4 验证Go安装

验证Go安装： 打开终端（Terminal）应用程序，并运行以下命令来验证Go是否正确安装：

```go
go version
```

## 二. Go 语言环境变量

### 2,1 查看Go 环境变量

打开终端（Terminal）应用程序，并运行以下命令来查看Go 环境变量：

```bash
go env
```

终端会返回如下命令：

```bash
GO111MODULE='on'
GOARCH='amd64'
GOBIN=''
GOCACHE='/Users/thinkook/Library/Caches/go-build'
GOENV='/Users/thinkook/Library/Application Support/go/env'
GOEXE=''
GOEXPERIMENT=''
GOFLAGS=''
GOHOSTARCH='amd64'
GOHOSTOS='darwin'
GOINSECURE=''
GOMODCACHE='/Users/thinkook/go/pkg/mod'
GONOPROXY=''
GONOSUMDB=''
GOOS='darwin'
GOPATH='/Users/thinkook/go'
GOPRIVATE=''
GOPROXY='https://goproxy.cn,direct'
GOROOT='/usr/local/Cellar/go/1.21.1/libexec'
GOSUMDB='sum.golang.org'
GOTMPDIR=''
GOTOOLCHAIN='auto'
GOTOOLDIR='/usr/local/Cellar/go/1.21.1/libexec/pkg/tool/darwin_amd64'
GOVCS=''
GOVERSION='go1.21.1'
GCCGO='gccgo'
GOAMD64='v1'
AR='ar'
CC='cc'
CXX='c++'
CGO_ENABLED='1'
GOMOD='/dev/null'
GOWORK=''
CGO_CFLAGS='-O2 -g'
CGO_CPPFLAGS=''
CGO_CXXFLAGS='-O2 -g'
CGO_FFLAGS='-O2 -g'
CGO_LDFLAGS='-O2 -g'
PKG_CONFIG='pkg-config'
GOGCCFLAGS='-fPIC -arch x86_64 -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -ffile-prefix-map=/var/folders/fs/wzn8gx9n3cq_gqvy502mn17r0000gn/T/go-build3919122639=/tmp/go-build -gno-record-gcc-switches -fno-common'
```

### 2.2 常用Go 配置项

| 环境变量    | 作用                                                         | 默认值                                                       |
| ----------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| GOARCH      | 用于指示Go编译器生成代码所针对的平台CPU架构                  | 本机的CPU架构（例如，AMD64、Arm等）                          |
| GOOS        | 用于指示Go编译器生成代码所针对的操作系统                     | 本机的操作系统（例如，Linux、Darwin、Windows等）             |
| GO111MODULE | 它的值决定了当前使用的构建模式是传统的GOPATH模式还是新引入的Go Module模式 | Go 1.16版本后，默认为"on"，开启Go Module构建模式             |
| GOCACHE     | 用于指示存储构建结果缓存的路径，这些缓存可能会被后续的构建所使用 | 不同操作系统上有不同的默认值。在Linux上默认值为 "$HOME/.cache/go-build" |
| GOMODCACHE  | 用于指示存放Go Module的路径                                  | 不同操作系统上有不同的默认值。在Linux上默认值为 "$HOME/go/pkg/mod" |
| GOPROXY     | 用来配置Go Module proxy服务                                  | 默认值为 "[https://proxy.golang.org,direct"。在中国大陆地区通常设置为](https://proxy.golang.org%2Cdirect".xn--fiq0a913absedgxbz7usxnil9etzzay9jw5r/) "[https://goproxy.cn,direct](https://goproxy.cn%2Cdirect/)" 以加速下载速度 |
| GOPRIVATE   | 指向自己的私有库，比如说自己公司的私有库                     | 无（默认值由用户配置）                                       |
| GOPATH      | 在传统的GOPATH构建模式下，用于指示Go包搜索路径的环境变量，在Go module机制启用之前是Go核心配置项。Go 1.8版本之前需要手工配置，Go 1.8版本引入了默认的GOPATH（$HOME/go）。在Go Module模式正式上位后，可能会被移除 | 无（默认值由用户配置）                                       |
| GOROOT      | 指示Go安装路径。Go 1.10版本引入了默认GOROOT，开发者无需显式设置，Go程序会自动根据自己所在路径推导出GOROOT的路径 | 无（由Go安装程序设置）                                       |

### **2.3 GOPATH**

- **src：存放源代码**：在GOPATH中的`src`目录用于存放你的Go项目的源代码文件。每个Go项目都应该有自己的目录，通常按照导入路径的结构来组织。例如，如果你的项目的导入路径是`github.com/yourusername/yourproject`，那么你的项目应该位于`$GOPATH/src/github.com/yourusername/yourproject`。
- **pkg：存放依赖包**：`pkg`目录用于存放已经编译好的Go包，这些包是你的项目依赖的其他包。这些包通常是由Go编译器自动构建并存储在`pkg`目录中，以便提高构建速度。
- **bin：存放可执行文件**：`bin`目录用于存放通过`go install`命令构建的可执行文件。当你使用`go install`安装一个Go程序时，可执行文件将被放置在`$GOPATH/bin`目录下，可以直接运行。

### **2.4 GOROOT**

- **go的安装目录**：GOROOT是指示Go语言安装目录的环境变量。它告诉Go编译器在哪里找到标准库以及其他与Go相关的工具和资源。GOROOT的默认值由Go安装程序设置，通常位于系统的特定目录中，例如`/usr/local/go`或`C:\Go`。开发者不需要手动设置GOROOT，Go会自动识别它。

### **2.5 GOPROXY**

#### 2.5.1 介绍说明:

- **用来配置Go Module proxy服务**：GOPROXY是一个环境变量，用于指定用于下载Go Module的代理服务。默认值为"[https://proxy.golang.org,direct"，它允许直接从官方代理下载Go](https://proxy.golang.xn--org%2Cdirect"%2Cgo-z512ao5ivmk85bwp1dqoa657njkli87g3erxu1lnwxa/) Module。
- 在中国大陆等地区，由于网络限制，通常会将GOPROXY设置为本地的Go Module代理服务，例如"[https://goproxy.cn,direct"，以加速下载速度和提高可靠性。设置合适的GOPROXY可以改善Go](https://goproxy.xn--cn%2Cdirect"%2C-gk3x49hf3w50fyodu54fc6gyvw5s3vxkdca7521cmku.xn--goproxygo-h37nn75bmiat4ony0e9g2cf23a7z7bxbq/) Module的下载体验。

#### 2.5.2 设置代理

对于Go版本1.13及以上：

```bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://proxy.golang.com.cn/,direct
```

**“direct”** 是一个特殊指示符，用于指示Go在需要时回源到模块版本的源地址（例如GitHub等）。如果在GOPROXY值列表中的前一个Go模块代理返回404或410错误，Go会自动尝试列表中的下一个代理，遇到 **“direct”** 时会回源，遇到EOF时会终止并抛出类似“invalid version: unknown revision…”的错误。

对于其他版本，可以使用以下方式设置代理：

```bash
export GO111MODULE=on
export GOPROXY=https://proxy.golang.com.cn/
```

#### 2.5.3 常用 GO 代理

以下是一些常用的Go Module代理服务的地址：

| 提供者       | 地址                                                        |
| :----------- | :---------------------------------------------------------- |
| 官方全球代理 | [https://proxy.golang.com.cn](https://proxy.golang.com.cn/) |
| 官方         | [https://goproxy.io,direct](https://goproxy.io%2Cdirect/)   |
| 七牛云       | [https://goproxy.cn](https://goproxy.cn/)                   |
| 阿里云       | https://mirrors.aliyun.com/goproxy/                         |
| GoCenter     | [https://gocenter.io](https://gocenter.io/)                 |
| 百度         | https://goproxy.bj.bcebos.com/                              |

这些代理服务可以帮助加速Go Module的下载，特别是在网络受限的情况下。您可以根据自己的需求选择合适的代理服务，并将GOPROXY设置为相应的地址。

## 三. 常用的IDE 设置

### 3.1 VsCode

**特点**：免费开源，强大的社区支持，扩展丰富，支持多种编程语言，包括Go。

**官网地址**:https://code.visualstudio.com/download

![image-20230921140250374](https://billy.taoxiaoxin.club/md/2023/09/650bdc8a11327e212447b72a.png)

**Go相关扩展**：在VsCode中使用Go开发，你可以安装以下一些常用的Go相关扩展来增强开发体验：

- **Go** - 微软官方的Go扩展,提供自动补全、格式化、重构等语言特性。
- **gopls** - 另一个语言服务器,提供智能提示、代码跳转、诊断等。有些人会觉得它比默认的Go扩展更好。
- **Go Test Explorer** - 在VS Code内运行测试并在侧边栏显示结果。
- **Go Modules** - 支持Go modules,如导入语句的模块路径自动补全。
- **gocode-gomod** - 使用Go Modules时自动补全模块。
- **dlv-dap** - 提供调试Delve调试器的配置。
- **gotests** - 自动为你的代码生成测试文件。
- **guru** - 集成guru工具,提供引用、定义等显示。
- **goreturns** - 添加快捷键运行goreturns来格式化和组织导入语句。

### 3.2 Goland

**特点**：Goland是JetBrains公司开发的商业IDE，专门用于Go语言开发。它提供了丰富的功能和工具，使Go开发更加高效和愉快。

**官网**:https://www.jetbrains.com/go/

![image-20230921140954264](https://billy.taoxiaoxin.club/md/2023/09/650bde3260ac7ceb07210307.png)

**Go相关功能**：Goland针对Go开发提供了许多功能，包括：

- **智能代码补全**：Goland具有高度智能的代码补全功能，可以准确地预测你要输入的代码，提高编码速度。
- **深度集成**：Goland深度集成了Go工具链，包括调试器、测试工具、代码导航等，使开发流程更加顺畅。
- **代码重构**：Goland支持各种代码重构操作，可以帮助你改进代码质量和结构。
- **实时错误检查**：Goland会在你编写代码时实时检查错误和代码质量问题，并提供即时反馈。
- **版本控制**：集成了常见的版本控制工具，如Git，方便团队协作和版本管理。

**配置**：

打开后点击新建项目:

![image-20230921145047247](https://billy.taoxiaoxin.club/md/2023/09/650be7c71c3dce4b19eafb13.png)

选择一个项目路径,和Go编译器创建项目

![image-20230921145324551](https://billy.taoxiaoxin.club/md/2023/09/650be864fe2cb5d39a81b0d3.png)

继续打开settings 设置Go Module 和 GOPROXY

<img src="https://billy.taoxiaoxin.club/md/2023/09/650bea55f168e91a2b91005a.png" alt="image-20230921150140801" style="zoom: 67%;" />

`Environment`正确填写内容：`GOPROXY=https://goproxy.cn,direct`

![image-20230921150913851](https://billy.taoxiaoxin.club/md/2023/09/650bec1aceba690b172c9959.png)

这些是常用的Go开发IDE中的两个主要选择。你可以根据自己的需求和偏好选择其中一个，它们都提供了强大的工具和功能来支持Go开发。