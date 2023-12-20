# GitHub学生包权益领取教程（超级详细）

哈喽，大家好，昨天给大家介绍了认证GitHub 学生！

今天教大家如何领取GitHub 学生包的权益。

## 往期推荐

+ [GitHub学生认证手把手教学，终于可以使用Github Copilot了(保姆级教程)](https://mp.weixin.qq.com/s?__biz=Mzg3ODA5ODY3MQ==&mid=2247499956&idx=1&sn=94d0f766756b789f7ac13711b53fdb42&chksm=cf1a6134f86de82272949bba126b43b79557b25f424906921aaceba1dd9f98a877c5631685c0&token=2096597808&lang=zh_CN#rd)

> GitHub学生包其实就是GitHub Pro，这是GitHub的付费认证用户身份，享有许多权益。接下来，我将为您详细介绍如何绑定各项权益。（**需要代认证的可以文末联系我，不是学生也可以**）

## GitHub Copilot 和 GitHub Copilot  chat 

### 简介

GitHub Copilot 简单一句话，`AI帮你写代码`，目前最强的AI写代码插件，只需要一行注释，轻松帮你给出代码提示。

GitHub Copilot Chat 旨在以自然语言专门回答你询问的编码相关问题。

### Github开启copilot权限

Github开启copilot权限，如图：

![](https://billy.taoxiaoxin.club/md/2023/12/65769bb260f92b5ac2d4ee62.png)

点击授权

![](https://billy.taoxiaoxin.club/md/2023/12/65769bb3af108beae6664478.png)

确认权限

![](https://billy.taoxiaoxin.club/md/2023/12/65769bb49179f7da34d8a741.png)

成功撒花

![](https://billy.taoxiaoxin.club/md/2023/12/65769bb56f469ccf3f828265.png)

### Jetbrains使用Copilot

1. 在 JetBrains IDE 中，在 Windows 的“文件”菜单或 IDE 名称 (Mac) 下（例如 PyCharm 或 IntelliJ），单击“设置”(Windows) 或“首选项”(Mac)。
2. 在“设置/首选项”对话框的左侧菜单中，单击“插件”。
3. 在“设置/首选项”对话框顶部，单击“市场”。 在搜索栏中，搜索“GitHub Copilot”，然后单击“安装”。

![](https://billy.taoxiaoxin.club/md/2023/12/65769bb50e2a0ab74d16898c.png)

安装 GitHub Copilot 后，单击“重启 IDE”。

JetBrains IDE 重启后，单击“工具”菜单。 单击“GitHub Copilot”，然后单击“登录到 GitHub”。

![](https://billy.taoxiaoxin.club/md/2023/12/6576a1230636ec873a1a83b7.png)

弹出下面这个框，点击 `Copy and Open`

![](https://billy.taoxiaoxin.club/md/2023/12/65769bb7d3cee0f682fad830.png)

输入设备码，直接粘贴

![](https://billy.taoxiaoxin.club/md/2023/12/65769bb7360c94f783552ec4.png)

通过认证

![](https://billy.taoxiaoxin.club/md/2023/12/65769bb836db5e56e09e1b0f.png)

认证成功

![](https://billy.taoxiaoxin.club/md/2023/12/65769bb86213dd7bbe537748.png)

重启idea查看

![](https://billy.taoxiaoxin.club/md/2023/12/65769bbaebd7df6e27903aa0.png)

接下来测试，只打了冒泡两个字，好家伙，tab、tab、tab，全都出来了

![](https://billy.taoxiaoxin.club/md/2023/12/65769bbab118599d0cc8652c.png)

### VSCode 安装 Copilot

首先需要在 VS Code 编辑器中安装相应的插件。在 VS Code 中，点击左侧的“扩展”选项卡，搜索“GitHub Copilot”，并安装该插件,选择第一个即可

![](https://billy.taoxiaoxin.club/md/2023/12/6576a7fd9528e03b3d954642.png)

### VsCode安装 GitHub Copilot chat 

搜索“GitHub Copilot chat”，并安装该插件,选择第一个即可

![](https://billy.taoxiaoxin.club/md/2023/12/6576a8cc0f86a81cc2416fb7.png)

### VsCode安装 GitHub Copilot Lab

搜索“GitHub Copilot Lab”，并安装该插件,选择第一个即可

![](https://billy.taoxiaoxin.club/md/2023/12/6576a90c89eb9a838ad64b7b.png)

安装完成后,重启VS Code，需要登录 GitHub 帐户进行身份验证。

### 设置防止代码泄漏

通常我们在公司写的代码属于机密，如果要在公司使用copilot，推荐关闭使用自己的代码进行机器学习，打开如下链接地址设置：

```go
https://github.com/settings/copilot
```

点击取消勾选`Allow GitHub to use my code snippets for product improvements *`

![](https://billy.taoxiaoxin.club/md/2023/12/6576a3526b3fb7f2486c9c99.png)

### GitHub Copilot Chat 提问问题

1. 在 JetBrains IDE 窗口右侧，单击 **GitHub Copilot Chat** 图标以将 GitHub Copilot Chat 窗口打开。

   ![](https://billy.taoxiaoxin.club/md/2023/12/6576a525f5163e06e01caa2a.png)

2. 在 GitHub Copilot Chat 窗口底部的“向 Copilot 询问问题或键入 `/` 作为命令”**** 文本框中，键入与编码相关的问题，然后按 **Enter**。 例如，键入“如何编写返回两个数字之和的函数？”。

3. GitHub Copilot Chat 将在聊天窗口中处理你的问题、提供答案，并在适当时提供代码建议。

   如果你的问题超出了 GitHub Copilot Chat 的范围，它会告诉你，并可能建议询问的替代问题。

4. （可选）如果 GitHub Copilot Chat 在“向 Copilot 询问问题或键入 `/` 作为命令”**** 文本框上方建议后续问题，请单击后续问题进行提问。

## 领取 Jetbrains 全家桶权益

### 简介

JetBrains全家桶是一组针对程序员的开发工具集合，包括集成开发环境（IDE）、代码编辑器和其他相关工具。以下是JetBrains全家桶的主要软件：

- IntelliJ IDEA：这是一个强大的Java开发工具，也支持Kotlin、Groovy、Scala等其他编程语言的开发。
- AppCode：主要用于Objective-C和Swift的开发。
- CLion：这是为C和C++开发者提供的IDE。
- DataGrip：这个工具主要用于数据库管理和SQL查询。
- GoLand：专门用于Golang的开发。
- PhpStorm：适用于PHP应用程序的开发。
- PyCharm：一款专为Python开发的IDE。
- Rider：这是一个跨平台的.NET开发工具。
- RubyMine：用于开发Ruby和Rails应用程序的工具。
- WebStorm：WebStorm是一款功能强大的JavaScript IDE，也支持TypeScript、HTML5和CSS3等前端技术的开发。

### 绑定Jetbrains账号

进入Github官网，点如下连接：

```go
https://education.github.com/pack
```

![](https://billy.taoxiaoxin.club/md/2023/12/65769ba88eeba704dd5b3ecf.png)

往下滑，找到Jetbrains，点击`Get access by connecting your GitHub account on JetBrains`

![](https://billy.taoxiaoxin.club/md/2023/12/65769ba9e6f9008bd2ae357c.png)

往下滑，点击Apply now

![](https://billy.taoxiaoxin.club/md/2023/12/65769ba9c56e2b2baeb72af6.png)

通过GTIHUB授权

![](https://billy.taoxiaoxin.club/md/2023/12/65769baa4767c15ad956b6f7.png)

授权

[![](https://billy.taoxiaoxin.club/md/2023/12/65769baaedc7b7c3ecfd7dfa.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/5.png)

**信息按照下图填**

![](https://billy.taoxiaoxin.club/md/2023/12/65769bab707146959d1dd8a9.png)

打开你刚刚输入的邮箱，邮箱会收到邮件

![](https://billy.taoxiaoxin.club/md/2023/12/65769bab492d56550b8ecec1.png)

点击邮箱中的链接，点击`Get started to use`

![](https://billy.taoxiaoxin.club/md/2023/12/65769bace8c5e76d75c292cb.png)

滑到底，点击I accept，这个页面我加载比较慢，等加载完成后才可以点击那个`I accept`按钮

![](https://billy.taoxiaoxin.club/md/2023/12/65769bad5348fec723f11a1a.png)

登录自己的Jetbrains账号

![](https://billy.taoxiaoxin.club/md/2023/12/65769bad899bfd8e49674bda.png)

出现以下画面就好了

![](https://billy.taoxiaoxin.club/md/2023/12/65769baec51d73c0012047a9.png)

登录idea，点击login

![](https://billy.taoxiaoxin.club/md/2023/12/65769baf06d0bbc82fab0464.png)

使用Github登录

![](https://billy.taoxiaoxin.club/md/2023/12/65769baf8e023ed650b3bdc8.png)

登录成功

![](https://billy.taoxiaoxin.club/md/2023/12/65769bb0d7cacdb605239526.png)

打开idea确认 ，点击Active

![](https://billy.taoxiaoxin.club/md/2023/12/65769bb1979feb17b4cdd890.png)

操作成功

![](https://billy.taoxiaoxin.club/md/2023/12/65769bb1b6570eb6ba6de9f6.png)

## 领取Termius权益

### 简介

👋`Termius`是一个终端工具，它是收费的，Github Pro用户可以免费使用

访问Termius官网（下面链接）：

```go
https://termius.com/education
```

点击 LINK YOUR GITHUB ACCOUNT

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bbb1653e45c9b390207.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/29.png)

点击 Authorize termius

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bbc123dea65f89bd0ef.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/30.png)

登录Termius

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bbc80c07b9d266d6b5a.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/31.png)

完成（不成功的话左下角会有过期时间显示）

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bbd7fc8e4ef0ac7ddb9.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/32.png)

也可以登录客户端查看（试用的话左下角会有过期时间显示）

👋客户端下载地址：https://termius.com/

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bbe607588949d31995c.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/33.png)

## 领取NameCheap域名权益

### 简介

`Namecheap`是互联网名称与数字地址分配机构认可的一家域名注册商，他们提供了免费的域名，只要你是GitHub Pro用户，你就可以免费使用.me后缀的域名。

进入Github官网（下面链接）:

```go
https://education.github.com/pack/offers
```

如果显示`502`，则进入下面这个

```go
https://education.github.com/pack
```

往下滑，找到namecheap，点击`Get access by connecting your GitHub account on Namecheap`

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bbe8bc66751ca4078a8.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/34.png)

填写想要的域名

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bc03b571e8cd6c869da.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/35.png)

选择.me域名，下单，免费的

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bc09d20623bbb3d73a4.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/36.png)

填写Gtihub邮箱

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bc1559c8a414fc62f6a.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/37.png)

注册Namecheap账号，填写账号信息，没有的话点击图中注册，记得记住账号和密码，后面登录需要

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bc1e202082e86265425.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/38.png)

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bc224da8cfffbbff1eb.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/39.png)

确认订单

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bc278eb678642ec05d6.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/40.png)

注册成功，这里可以点击一下设置Github账户，他就是添加一个Github的DNS而已，没什么，后续有需要可以再修改。

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bc3c4c5793071ab0d40.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/41.png)

确认Github信息

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bc3a8f1253476162b0f.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/42.png)

注册成功

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bc4e7139e9fc0dd4f6f.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/43.png)

 查看域名，登录NameCheap官网：

```go
https://www.namecheap.com/myaccount/login
```

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bc46d05d6b7db8fd44a.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/44.png)

填写邮箱验证码

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bc53415addd9e8b6d9b.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/45.png)

登录成功

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bc5af1ebf823bcc1ade.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/46.png)

查看域名列表

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bc6366816cb787f16ea.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/47.png)

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bc621d04f42fd77bae4.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/48.png)

查看域名信息，可修改DNS，点击上图域名后面的`MANAGE`，进入域名管理页面

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bc7053bfc6d2f2e1a0e.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/49.png)

## 领取.tech域名权益

### 简介

> 使用 .tech 域名，专注于技术领域。.tech 是广受散户投资者、风险投资家、未来人才和终端消费者推崇的域名后缀。由于 tech 是“技术”的缩写，因此专注于各个技术领域的众多组织均可以使用该域名来突出它们的技术专长。从缩短您的品牌名称到巩固您作为软件即服务（SaaS）行业领导者的品牌声誉，.tech 域名具有众多优势，可以成为您的理想选择。

注册域名，进入Github官网（下面链接）

```go
https://education.github.com/pack/offers
```

如果显示`502`，则进入下面这个

```go
https://education.github.com/pack
```

往下滑，找到**.tech**，点击`Get access by connecting your GitHub account on Namecheap`

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bc78c10d2fcbeb20022.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/50.png)

选择域名

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bc94f5fe8f1f5e449c7.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/51.png)

加入购物车，并去支付

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bc9083e81cc770d9389.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/52.png)

登录Github并授权，授权之后就是0元

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bca04f2d2f6733e08f3.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/53.png)

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bca40ee0e8e1d33a26b.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/54.png)

注册.tech账号，如果你有.tech账号，直接登录就好了，如果没有，注册一个

![](https://billy.taoxiaoxin.club/md/2023/12/65769bcb1a3e102340232957.png)

![](https://billy.taoxiaoxin.club/md/2023/12/65769bcbf98a56f97ef2aa5e.png)

注册完成后确认订单

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bcc36beebe807aa1ea8.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/57.png)

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bcc2d27ebdb0cc12acf.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/58.png)

管理域名，点击右上角`MY ACCOUNT`，进入后台界面

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bcd74cf0ba9aa733d92.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/59.png)

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bcd5878d4e2f74bcfc3.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/60.png)

可以修改默认的DNS服务器

[![](https://billy.taoxiaoxin.club/md/2023/12/65769bced22e06d5b25e11e1.png)](https://blog-1305951218.cos.ap-shanghai.myqcloud.com/blog/image/articleContent/GitHub_Pro/61.png)

## 其他权益领取

进入Github官网（下面链接）

```go
https://education.github.com/pack/offers
```

如果显示`502`，则进入下面这个

```go
https://education.github.com/pack
```

往下滑，找到**你想要产品**，点击`Get access by connecting your GitHub account on Namecheap`，接着使用GitHub 邮箱认证或者使用GitHub 账号登录授权即可。基本上都是这个套路，我写不动了，剩下的权益领取自己研究吧。

## GitHub 学生代认证

+ 不管你是否毕业，都可以帮你认证，100%秒通过，
+ 充值到自己账号，认证通过后质保一年，稳定使用2年，有官方售后群
+ 现在咨询可能还有点优惠，后面会涨价。

**扫码以下二维码咨询，备注 GitHub 学生包优先通过。**

![](https://billy.taoxiaoxin.club/md/2023/12/65769c1c429704b121bf5458.png)