# Go 包操作之如何拉取私有的Go Module

在前面，我们已经了解了[GO 项目依赖包管理与Go Module常规操作](https://blog.csdn.net/weixin_44621343/article/details/133256292)，Go Module 构建模式已经成为了 Go 语言的依赖管理与构建的标准。

在平时使用Go Module 时候，可能会遇到以下问题：

+ 在某 module 尚未发布到类似GitHub 或 Gitee  这样的网站前，如何 import 这个本地的 module？
+ 如何拉取私有 module？

[TOC]



## 一、导入本地 module

### 1.1 依赖本地尚未发布的 module

**如果我们的项目依赖的是本地正在开发、尚未发布到公共站点上的 Go Module，那么我们应该如何做呢？**

例如：假设有个hello-module的项目，你的main包中依赖了moduleA,代码如下:

```go
package main

import "gitee.com/tao-xiaoxin/study-basic-go/hello-module/moduleA"

func main() {
	moduleA.ModuleA()
}
```

并且，这个项目中的moduleA 依赖 moduleB,此时此刻，module A 和 moduleB 还没有发布到`gitee`公共托管站点上，它的源码还在你的开发机器上。也就是说，Go 命令无法在`gitee.com/user/`上找到并拉取 module A 和 module B，这时，使用`go mod tidy`命令，就会收到类似下面这样的报错信息：

```go
$go mod tidy
go: finding module for package gitee.com/user/moduleB
go: finding module for package gitee.com/user/moduleA
go: gitee.com/tao-xiaoxin/study-basic-go imports
        gitee.com/user/moduleA: module gitee.com/user: git ls-remote -q origin in /Users/thinkook/go/pkg/mod/cache/vcs/ff424152e6f6be73e07b96e5d8e06c6cd9f86dc9903058919a7b8737718a8418: exit status 128:
        致命错误：仓库 'https://gitee.com/user/' 未找到
go: gitee.com/tao-xiaoxin/study-basic-go/moduleA imports
        gitee.com/user/moduleB: module gitee.com/user: git ls-remote -q origin in /Users/thinkook/go/pkg/mod/cache/vcs/ff424152e6f6be73e07b96e5d8e06c6cd9f86dc9903058919a7b8737718a8418: exit status 128:
        致命错误：仓库 'https://gitee.com/user/' 未找到
```

所以，Go提供了两种方式可以导入本地正在开发的 Go Module

### 1.2 Go Module 开发中本地导入两种方式

#### 1.2.1 使用 replace 指令

**介绍：** 使用`replace`指令可以替代远程依赖模块的路径，将其指向本地的模块路径，便于本地开发和测试。

**基本使用：** 下面是一个示例replace指令的使用方式：

```go
replace example.com/module@版本号 => 你的本地Module路径（可以使用相对路径或者绝对路径）
```

接着，我们继续回到上面的举例中，首先，我们需要在 module a 的 go.mod 中的 `require` 块中手工加上这一条并且替换为本地路径上的`module A`和`moduleB`:

```go
replace (
	gitee.com/user/moduleA v1.0.0 => ../moduleA
	gitee.com/user/moduleB v1.0.0 => ../moduleB
)
```

这里的`v1.0.0`版本号是一个“假版本号”，目的是满足`go.mod`中`require`块的语法要求。

或者使用`go mod edit` 命令编辑 `go.mod` 文件:

```go
go mod edit -replace=gitee.com/user/moduleA@v1.0.0=../moduleA -replace=gitee.com/user/moduleB@v1.0.0=../moduleB
```

这样修改之后，Go 命令就会让`module A`依赖你本地正在开发、尚未发布到代码托管网站的`module B`的源码了，并且`main`函数依赖你本地正在开发、尚未发布到代码托管网站的`module B`的源码了。

虽然虽然这个方案可以解决上述问题，但是在平时开发过程中，`go.mod` 文件通常需要上传到代码服务器上，这意味着，另一个开发人员下载了这份代码后，很可能无法成功编译。在这个方法中，`require`指示符将`gitee.com/user/moduleA v1.0.0`替换为一个本地路径下的`module A`的源码版本，但这个**本地路径因开发者环境而异**。为了成功编译`module A`和主程序，该开发人员必须将`replace`后面的本地路径更改为适应自己的环境路径。

于是，每当开发人员 pull 代码后，第一件事就是要修改`go.mod`中的`replace`块。每次上传代码前，可能还要将`replace`路径还原，这是一个很繁琐的事情。于是，Go开发团队在Go 1.18 版本中加入了 Go 工作区（Go workspace，也译作 Go 工作空间）辅助构建机制。

上述举例代码仓库地址:[点我进入](https://gitee.com/tao-xiaoxin/study-basic-go/tree/master/syntax/hello-module)

#### 1.2.2 使用工作区模式

**介绍：**Go 工作区模式是 Go 语言 1.18 版本引入的新功能，允许开发者将多个本地路径放入同一个工作区中，这样，在这个工作区下各个模块的构建将优先使用工作区下的模块的源码。工作区模式具有以下优势：

- 可以将多个本地模块放入同一个工作区中，方便开发者管理。
- 可以解决“伪造 go.mod”方案带来的那些问题。
- 可以提高模块构建的性能。

**常用命令：**

Go 工具提供了以下命令来帮助开发者使用工作区模式：

+ `go work edit`：提供了用于修改`go.work`的命令行接口，主要是给工具或脚本使用。
+ `go work init`：初始化工作区文件 go.work
+ `go work use`：将模块添加到工作区文件
+ `go work sync`：把`go.work`文件里的依赖同步到workspace包含的Module的`go.mod`文件中。

**基本使用：**

1. 首先，我们初始化 Go workspace 使用命令`go work init`命令如下：

```go
go work init [moddirs]
```

`moddirs`是Go Module所在的本地目录。如果有多个Go Module，就用空格分开。如果`go work init`后面没有参数，会创建一个空的workspace。

执行`go work init`后会生成一个`go.work`文件，`go.work`里列出了该workspace需要用到的Go Module所在的目录，workspace目录不需要包含你当前正在开发的Go Module代码。

2. 如果要给workspace新增Go Module，可以使用如下命令：

```go
go work use [-r] moddir
```

如果带有`-r`参数，会递归查找`-r`后面的路径参数下的所有子目录，把所有包含`go.mod`文件的子目录都添加到`go work`文件中。

3. 如果要同步依赖到workspace包含的Module的`go.mod`文件中，可以使用如下命令：

   ```go
   go work sync
   ```

介绍完之后,我们回到上面的例子中，现在我们进入 gowork下面，然后通过下面命令初始化一个go.work:

```go
go work init .
```

我们看到`go work init`命令创建了一个`go.work`文件，使用`go env GOWORK`命令查看`go.work`所在位置

```go
$go env GOWORK
~/workspace/GolandProjects/study-basic-go/syntax/gowork/go.work
```

接着，我们在 `module a` 的`go.work` 中的 `use` 块中替换为本地路径上的`module A`和`moduleB`:

```go
go 1.21.1

use (
	.
	./moduleA
	./moduleB
)
```

**支持replace指示符**：go.work还支持replace指示符,使用方法和上面一样

上面的代码地址:[点我](https://gitee.com/tao-xiaoxin/study-basic-go/tree/master/syntax/gowork)

## 二、拉取私有 module 的需求与参考方案

自从 Go 1.11 版本引入 Go Module 构建模式后，通过 Go 命令拉取项目依赖的公共 Go Module，已不再是一个“痛点”。现在，我们只需要在每个开发机上设置环境变量 GOPROXY，配置一个高效且可靠的公共 GOPROXY 服务，就可以轻松地拉取所有公共 Go Module 了。

![img](https://billy.taoxiaoxin.club/md/2023/10/6537f723a99a80040e666175.png)

但随着公司内 Go 使用者和 Go 项目的增多，“重造轮子”的问题就出现了。抽取公共代码放入一个独立的、可被复用的内部私有仓库成为了必然，这样我们就有了拉取私有 Go Module 的需求。

一些公司或组织的所有代码，都放在公共 vcs 托管服务商那里（比如 github.com），私有 Go Module 则直接放在对应的公共 vcs 服务的 private repository（私有仓库）中。如果你的公司也是这样，那么拉取托管在公共 vcs 私有仓库中的私有 Go Module，也很容易，见下图：

![img](https://billy.taoxiaoxin.club/md/2023/10/6537f755616be05b709cceba.png)

也就是说，只要我们在每个开发机上，配置公共 GOPROXY 服务拉取公共 Go Module，同时将私有仓库配置到 GOPRIVATE 环境变量，就可以了。这样，所有私有模块的拉取都将直接连接到代码托管服务器，不会通过 GOPROXY 代理服务，并且不会向 GOSUMDB 服务器发出 Go 包的哈希值校验请求。

当然，这个方案有一个前提，那就是每个开发人员都需要具有访问公共 vcs 服务上的私有 Go Module 仓库的权限，凭证的形式不限，可以是 basic auth 的 user 和 password，也可以是 personal access token（类似 GitHub 那种），只要按照公共 vcs 的身份认证要求提供就可以了。

不过，更多的公司 / 组织，可能会将私有 Go Module 放在公司 / 组织内部的 vcs（代码版本控制）服务器上，就像下面图中所示：

![img](https://billy.taoxiaoxin.club/md/2023/10/6537f79b41006b8e6962d8d6.png)

那么这种情况，我们该如何让 Go 命令，自动拉取内部服务器上的私有 Go Module 呢？这里给出两个参考方案。

### 2.1 方案一：通过直连组织公司内部的私有 Go Module 服务器拉取

![img](https://billy.taoxiaoxin.club/md/2023/10/6537f7c27fc3c39afa8ab20f.png)

在这个方案中，我们看到，公司内部会搭建一个内部 goproxy 服务（也就是上图中的 in-house goproxy）。这样做有两个目的，一是为那些无法直接访问外网的开发机器，以及 ci 机器提供拉取外部 Go Module 的途径，二来，由于 in-house goproxy 的 cache 的存在，这样做还可以加速公共 Go Module 的拉取效率。

另外，对于私有 Go Module，开发机只需要将它配置到 GOPRIVATE 环境变量中就可以了，这样，Go 命令在拉取私有 Go Module 时，就不会再走 GOPROXY，而会采用直接访问 vcs（如上图中的 git.yourcompany.com）的方式拉取私有 Go Module。

这个方案十分适合内部有完备 IT 基础设施的公司。这类型的公司内部的 vcs 服务器都可以通过域名访问（比如 git.yourcompany.com/user/repo），因此，公司内部员工可以像访问公共 vcs 服务那样，访问内部 vcs 服务器上的私有 Go Module。

### 2.2 方案二：将外部 Go Module 与私有 Go Module 都交给内部统一的 GOPROXY 服务去处理：

![img](https://billy.taoxiaoxin.club/md/2023/10/6537f7ff02246dfbfedbc96f.png)

在这种方案中，开发者只需要把 GOPROXY 配置为 in-house goproxy，就可以统一拉取外部 Go Module 与私有 Go Module。

但由于 go 命令默认会对所有通过 goproxy 拉取的 Go Module，进行 sum 校验（默认到 sum.golang.org)，而我们的私有 Go Module 在公共 sum 验证 server 中又没有数据记录。因此，开发者需要将私有 Go Module 填到 GONOSUMDB 环境变量中，这样，go 命令就不会对其进行 sum 校验了。

不过这种方案有一处要注意：in-house goproxy 需要拥有对所有 private module 所在 repo 的访问权限，才能保证每个私有 Go Module 都拉取成功。

在平时开发中，更推荐第二个方案。在第二个方案中，我们可以将所有复杂性都交给 in-house goproxy 这个节点，开发人员可以无差别地拉取公共 module 与私有 module，心智负担降到最低。

## 三、统一 Goproxy 方案的实现思路与步骤

### 3.1  goproxy 服务搭建

[Go module proxy](https://pkg.go.dev/cmd/go@master#hdr-Module_proxy_protocol) 协议规范发布后，Go 社区出现了很多成熟的 Goproxy 开源实现，比如最初的 [Athens](https://github.com/gomods/athens)，还有国内的两个优秀的开源实现：[goproxy.cn](https://github.com/goproxy/goproxy) 和 [goproxy.io](https://github.com/goproxyio/goproxy) 等。其中，goproxy.io 在官方站点给出了[企业内部部署的方法](https://goproxy.io/zh/docs/enterprise.html)，所以今天我们将基于 goproxy.io 来实现我们的方案。

我们在上图中的 in-house goproxy 节点上执行这几个步骤安装 goproxy：

```go
$mkdir ~/.bin/goproxy
$cd ~/.bin/goproxy
$git clone https://github.com/goproxyio/goproxy.git
$cd goproxy
$make
```

编译后，我们会在当前的 bin 目录（~/.bin/goproxy/goproxy/bin）下看到名为 goproxy 的可执行文件。

然后，我们建立 goproxy cache 目录：

```go
$mkdir /root/.bin/goproxy/goproxy/bin/cache
```

再启动 goproxy：

```go
$./goproxy -listen=0.0.0.0:8081 -cacheDir=/root/.bin/goproxy/goproxy/bin/cache -proxy https://goproxy.io
goproxy.io: ProxyHost https://goproxy.io
```

启动后，goproxy 会在 8081 端口上监听（即便不指定，goproxy 的默认端口也是 8081），指定的上游 goproxy 服务为 goproxy.io。

不过要注意下：goproxy 的这个启动参数并不是最终版本的，这里我仅仅想验证一下 goproxy 是否能按预期工作。我们现在就来实际验证一下。

首先，我们在开发机上配置 GOPROXY 环境变量指向 10.10.20.20:8081：

~~~go
// .bashrc
export GOPROXY=http://10.10.20.20:8081
~~~

生效环境变量后，执行下面命令：

```go
$go get github.com/pkg/errors
```

结果和我们预期的一致，开发机顺利下载了 github.com/pkg/errors 包。我们可以在 goproxy 侧，看到了相应的日志：

```go
goproxy.io: ------ --- /github.com/pkg/@v/list [proxy]
goproxy.io: ------ --- /github.com/pkg/errors/@v/list [proxy]
goproxy.io: ------ --- /github.com/@v/list [proxy]
goproxy.io: 0.146s 404 /github.com/@v/list
goproxy.io: 0.156s 404 /github.com/pkg/@v/list
goproxy.io: 0.157s 200 /github.com/pkg/errors/@v/list
```

在 goproxy 的 cache 目录下，我们也看到了下载并缓存的 github.com/pkg/errors 包：

```go
$cd /root/.bin/goproxy/goproxy/bin/cache
$tree
.
└── pkg
    └── mod
        └── cache
            └── download
                └── github.com
                    └── pkg
                        └── errors
                            └── @v
                                └── list

8 directories, 1 file
```

这就标志着我们的 goproxy 服务搭建成功，并可以正常运作了。

### 3.2 自定义包导入路径并将其映射到内部的 vcs 仓库

一般公司可能没有为 VCS 服务器分配域名，我们也不能在 Go 私有包的导入路径中放入 IP 地址，因此我们需要给我们的私有 Go Module 自定义一个路径，比如：`mycompany.com/go/module1`。我们统一将私有 Go Module 放在 `mycompany.com/go` 下面的代码仓库中。

那么，接下来的问题就是，当 goproxy 去拉取 `mycompany.com/go/module1` 时，应该得到 `mycompany.com/go/module1` 对应的内部 VCS 上 `module1` 仓库的地址，这样，goproxy 才能从内部 VCS 代码服务器上下载 `module1` 对应的代码，具体的过程如下：

![WechatIMG240](https://billy.taoxiaoxin.club/md/2023/10/6538853aea8fb65c255a98d1.jpg)

那么我们如何实现为私有 module 自定义包导入路径，并将它映射到内部的 vcs 仓库呢？

其实方案不止一种，这里我使用了 Google 云开源的一个名为 govanityurls 的工具，来为私有 module 自定义包导入路径。然后，结合 [govanityurls](https://github.com/GoogleCloudPlatform/govanityurls) 和 Nginx，我们就可以将私有 Go Module 的导入路径映射为其在 VCS 上的代码仓库的真实地址。具体原理你可以看一下这张图：

![WechatIMG241](https://billy.taoxiaoxin.club/md/2023/10/653885c6bc34147f00f164e3.jpg)

首先，goproxy 要想不把收到的拉取私有 Go Module（mycompany.com/go/module1）的请求转发给公共代理，需要在其启动参数上做一些手脚，比如下面这个就是修改后的 goproxy 启动命令：

```go
$./goproxy -listen=0.0.0.0:8081 -cacheDir=/root/.bin/goproxy/goproxy/bin/cache -proxy https://goproxy.io -exclude "mycompany.com/go"
```

这样，凡是与 `-exclude` 后面的值匹配的 Go Module 拉取请求，goproxy 都不会将其转发给 goproxy.io，而是直接请求 Go Module 的“源站”。

而上面这张图中要做的，就是将这个“源站”的地址，转换为企业内部 VCS 服务中的一个仓库地址。然后我们假设 `mycompany.com` 这个域名并不存在（很多小公司没有内部域名解析能力），从图中我们可以看到，我们会在 `goproxy` 所在节点的 `/etc/hosts` 中添加这样一条记录：

~~~go
127.0.0.1 mycompany.com
~~~

这样做了后，goproxy 发出的到 `mycompany.com` 的请求实际上是发向了本机。而上面这图中显示，监听本机 `80` 端口的正是 `nginx`，`nginx` 关于 `mycompany.com` 这一主机的配置如下：

```go
// /etc/nginx/conf.d/gomodule.conf

server {
        listen 80;
        server_name mycompany.com;

        location /go {
                proxy_pass http://127.0.0.1:8080;
                proxy_redirect off;
                proxy_set_header Host $host;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;

                proxy_http_version 1.1;
                proxy_set_header Upgrade $http_upgrade;
                proxy_set_header Connection "upgrade";
        }
}
```

我们看到，对于路径为 `mycompany.com/go/xxx` 的请求，`nginx` 将请求转发给了 `127.0.0.1:8080`，而这个服务地址恰恰就是 `govanityurls` 工具监听的地址。

`govanityurls` 这个工具，是前 Go 核心开发团队成员 [Jaana B. Dogan](https://rakyll.org/) 开源的一个工具，这个工具可以帮助 Gopher 快速实现自定义 Go 包的 `go get` 导入路径。

govanityurls 本身，就好比一个“导航”服务器。当 go 命令向自定义包地址发起请求时，实际上是将请求发送给了 govanityurls 服务，之后，govanityurls 会将请求中的包所在仓库的真实地址（从 vanity.yaml 配置文件中读取）返回给 go 命令，后续 go 命令再从真实的仓库地址获取包数据。

>注：govanityurls 的安装方法很简单，直接 go install/go get github.com/GoogleCloudPlatform/govanityurls 就可以了。在我们的示例中，vanity.yaml 的配置如下：

```go
host: mycompany.com

paths:
  /go/module1:
      repo: ssh://admin@10.10.30.30/module1
      vcs: git
```

也就是说，当 `govanityurls` 收到 `nginx` 转发的请求后，会将请求与 `vanity.yaml` 中配置的 `module` 路径相匹配，如果匹配 OK，就会将该 `module` 的真实 repo 地址，通过 `go` 命令期望的应答格式返回。在这里我们看到，`module1` 对应的真实 VCS 上的仓库地址为：`ssh://admin@10.10.30.30/module1`。

所以，`goproxy` 会收到这个地址，并再次向这个真实地址发起请求，并最终将 `module1` 缓存到本地 `cache` 并返回给客户端。

### 3.3 开发机 (客户端) 的设置

前面示例中，我们已经将开发机的 `GOPROXY` 环境变量，设置为 `goproxy` 的服务地址。但我们说过，凡是通过 `GOPROXY` 拉取的 Go Module，`go` 命令都会默认把它的 `sum` 值放到公共 `GOSUM` 服务器上去校验。

但我们实质上拉取的是私有 Go Module，`GOSUM` 服务器上并没有我们的 Go Module 的 `sum` 数据。这样就会导致 `go build` 命令报错，无法继续构建过程。

因此，开发机客户端还需要将 `mycompany.com/go`，作为一个值设置到 `GONOSUMDB` 环境变量中：

~~~go
export GONOSUMDB=mycompany.com/go
~~~

这个环境变量配置一旦生效，就相当于告诉 `go` 命令，凡是与 `mycompany.com/go` 匹配的 Go Module，都不需要再做 `sum` 校验了。

到这里，我们就实现了拉取私有 Go Module 的方案。

### 3.4 方案的“不足”

#### 3.4.1 第一点：开发者还是需要额外配置 GONOSUMDB 变量

由于 Go 命令默认会对从 `GOPROXY` 拉取的 Go Module 进行 `sum` 校验，因此我们需要将私有 Go Module 配置到 `GONOSUMDB` 环境变量中，这就给开发者带来了一个小小的“负担”。

对于这个问题，我的解决建议是：公司内部可以将私有 Go 项目都放在一个特定域名下，这样就不需要为每个 Go 私有项目单独增加 `GONOSUMDB` 配置了，只需要配置一次就可以了。

#### 3.4.2 第二点：新增私有 Go Module，vanity.yaml 需要手工同步更新

这是这个方案最不灵活的地方了，由于目前 `govanityurls` 功能有限，针对每个私有 Go Module，我们可能都需要单独配置它对应的 VCS 仓库地址，以及获取方式（git、svn 或 hg）。

关于这一点，我的建议是：在一个 VCS 仓库中管理多个私有 Go Module。相比于最初 Go 官方建议的[一个 repo 只管理一个 module](https://go.dev/doc/modules/managing-source#multiple-module-source)，新版本的 Go 在一个 repo 下管理多个 Go Module 方面，已经有了长足的进步，我们已经可以通过 repo 的 tag 来区别同一个 repo 下的不同 Go Module。

不过对于一个公司或组织来说，这点额外工作与得到的收益相比，应该也不算什么！

#### 3.4.3 第三点：无法划分权限

在讲解上面的方案的时候，我们也提到过，`goproxy` 所在节点需要具备访问所有私有 Go Module 所在 VCS repo 的权限，但又无法对 Go 开发者端做出有差别授权。这样，只要是 `goproxy` 能拉取到的私有 Go Module，Go 开发者都能拉取到。

不过对于多数公司而言，内部所有源码原则上都是企业内部公开的，这个问题似乎也不大。如果觉得这是个问题，那么只能使用前面提到的第一个方案，也就是直连私有 Go Module 的源码服务器的方案了。

参考链接：

+ [小厂内部私有Go module拉取方案3](https://tonybai.com/?s=%E5%B0%8F%E5%8E%82%E5%86%85%E9%83%A8%E7%A7%81%E6%9C%89Go+module%E6%8B%89%E5%8F%96%E6%96%B9%E6%A1%883)
+ [Go 1.18新特性前瞻：Go工作区模式](https://tonybai.com/2021/11/12/go-workspace-mode-in-go-1-18/)
