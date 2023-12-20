

哈喽,大家好!

继上次更[Mac 最新版Navicat v16.2.9 安装与永久激活教程](https://mp.weixin.qq.com/s?__biz=Mzg3ODA5ODY3MQ==&mid=2247498853&idx=1&sn=e638f7282e61c94d1f6069160d3712cb&chksm=cf1a5de5f86dd4f362cd5ad378d1624f19efb3d092c00f3d80926a5b3d59c38e15df200364d0&token=608821632&lang=zh_CN#rd)

之后，收到一部分小伙伴反馈`Navicat`各种问题，安装提示什么试用期了，还有什么打不开的问题啦一堆。

![image-20231101192849108](https://billy.taoxiaoxin.club/md/2023/11/654236721b551528b7728bd2.png)

![image-20231101192950461](https://billy.taoxiaoxin.club/md/2023/11/654236ae73d2b74b9fea516c.png)

所以，今天就再给大家更一波，主要解决以上问题。 

## 特别声明

+ **本教程仅供个人学习和研究使用**
+ Navicat 官网地址: https://www.navicat.com/en/

## 软件介绍

Navicat Premium for Mac是一款功能强大的数据库管理工具，具有多重连接的特点。无论是专业开发人员还是数据库服务器的新手，都能够轻松学习和使用它。

Navicat Premium支持连接到目前主流数据库的所有版本,包括MySQL、SQL Server、SQLite、Oracle和`PostgreSQL`，Redis等。这意味着你可以使用单一程序连接并管理不同类型的数据库，从而使数据库管理更加方便。

此外，Navicat Premium还提供了许多有用的功能，如数据可视化和操作的便捷性，以及对数据安全性的重视。它不仅可以帮助你快速执行查询和数据操作，还可以保护数据的安全性，从而避免数据泄露和不良后果。

总之，Navicat Premium for Mac是一个强大且易于学习和使用的数据库管理工具，提供了多重连接和支持多种数据库的特点，以及许多有用的功能，使数据库管理变得更加轻松和高效。

![image-20231012161103721](https://billy.taoxiaoxin.club/md/2023/11/6542135fa3527b431a149add.png)

## 解决方案

针对以上问题有两种解决方案：

1. 方案一：从软件商店下载，然后激活程序，需要App Store 下载软件，**系统要求必须是 macOS 10.14及以上**，最后使用我提供的激活工具激活。
2. 方案二：使用14天无限白嫖大法，就是无限重置14天试用期，适合企业内部或者考虑安全性的同学使用。

## 方案优缺点

方案一：

+ 优点：直接从应用商店下载，不会提示什么打不开的问题，即使软件更新了，激活依然有效。
+ 缺点：**激活软件系统要求必须是 macOS 10.14及以上，否则软件崩溃报错，不会保留你之前的数据库连接记录，需要手动备份并且导入数据库链接，这个自行百度**。

方案二：

+ 优点：安全性之类的没有任何问题，都是从官网下载的软件，激活软件只是清除了试用期时长，安装直接替换会保留你的连接记录。
+ 缺点：一直会有14天试用期的弹窗。

鉴于以上问题优缺点对比，看你自己选择。

## 温馨提示

如果是商用和公司的电脑，不要使用方案一，这软件有后门，会定期检测你是不是正常激活，不然公司会因为你使用了这个被罚n多钱...并且，软件厂家直接就有你公司的地址甚至哪台电脑在用盗版，综上，个人学习使用没有问题，公司商用请支持正版或者换软件，也可以考虑使用方案二。

## 方案一

打开`App Store` 搜索 Navicat,点击获取下载即可。

![image-20231101200520043](https://billy.taoxiaoxin.club/md/2023/11/65423f00e2c311eca981b228.png)

先不要着急打开 Navicat Premium！！！

接着，继续打开终端，开启“任何来源”（已经开启请跳过），执行如下命令：

```bash
sudo spctl  --master-disable
```

然后回车，继续输入密码（密码输入时是不可见的），然后回车。

接着打开【**系统偏好设置**】，选择【**安全性与隐私**】，选择【**通用**】，可以看到【**任何来源**】已经选定。

![image-20231012152846076](https://billy.taoxiaoxin.club/md/2023/11/6542135f84b5b67a86c0c345.png)

接下来安装激活工具，**点击文章末尾百度云下载地址下载**

![image-20231101201110458](https://billy.taoxiaoxin.club/md/2023/11/6542405e185fa17924a352d1.png)

**运行激活工具前务必关闭Navicat Premium**，然后双击打开激活软件如下：

![image-20231101201259715](https://billy.taoxiaoxin.club/md/2023/11/654240cb331dd25265eef066.png)

**最后点击Navicat Premium Crack 激活**

![image-20231101201848828](https://billy.taoxiaoxin.club/md/2023/11/65424229c91f555a469106f6.png)

最后点击一键激活就好了。

![image-20231101202024541](https://billy.taoxiaoxin.club/md/2023/11/65424288a2ebdf1109b7cb33.png)

## 方案二

先从官网下载地址如下:

```bash
https://www.navicat.com.cn/download/navicat-premium
```

点击下载即可

![image-20231101210305091](https://billy.taoxiaoxin.club/md/2023/11/65424c896a5fa84479dbaf11.png)

这里官网最新版本是16.3.1,贴上我下载的时候官网下载地址,希望对你们有用:

```bash
https://download.navicat.com.cn/download/navicat163_premium_cs.dmg
```

将左边应用拖到**Applications**即可安装

![image-20231101210353628](https://billy.taoxiaoxin.club/md/2023/11/65424cb97363ed20ce75ebd2.png)

然后打开软件点击试用

![image-20231101214830049](https://billy.taoxiaoxin.club/md/2023/11/6542572e53fd8fd01e8523ea.png)

如果提示试用期已过期

![image-20231101214656007](https://billy.taoxiaoxin.club/md/2023/11/654256d0b3603c16e1f72f14.png)

首先下载`reset_navicat.sh`或者 `reset_navicat_by_52pojie.sh`,打开下载目录，鼠标右击，选择“在终端中打开”，这样打开终端后直接就进入了当前目录

![输入图片说明](https://billy.taoxiaoxin.club/md/2023/11/65425b4c868b22b4def32bf6.png)

为文件授予可执行权限：

```bash
# 为reset_navicat.sh文件授予可执行权限
chmod u+x reset_navicat.sh
# 或者给reset_navicat_by_52pojie.sh文件授予可执行权限
chmod u+x reset_navicat_by_52pojie.sh
```

接着执行如下命令：

```bash
# 执行reset_navicat.sh文件
sh reset_navicat.sh
# 或者执行reset_navicat_by_52pojie.sh文件
sh reset_navicat_by_52pojie.sh
```

如果你觉得这样太过麻烦，下面教你设置定时任务自动执行脚本。**我们假定让自己的电脑在每天上午9:10自动执行脚本重置Navicat Premium 16试用期，下面是操作步骤。**

1. 下载`reset.navicat.premium.trial.period.plist`、`reset_navicat.sh`或`reset_navicat_by_52pojie.sh`
2. 此时只要使用命令`chmod u+x xxx.sh`给`reset_navicat.sh`或`reset_navicat_by_52pojie.sh`文件赋予可执行权限，然后双击执行该脚本即可重置试用期。
3. 按照注释修改`reset.navicat.premium.trial.period.plist`文件，你可以按照我1，2，3，4排序进行修改，其他的你自己看着改吧！

```html
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <!-- 1.此处定义的是定时任务的名称，之后可用于搜索或停止该任务，建议与文件名一致即可 -->
    <string>reset.navicat.premium.trial.period</string>
    <!-- 以下两个<string>标签修改为你的reset_navicat.sh脚本的绝对路径,或者reset_navicat_by_52pojie.sh脚本路径,获取当前路径使用命令pwd查看 -->
    <key>Program</key>
    <!-- 2.修改为你的reset_navicat.sh脚本的绝对路径,或者reset_navicat_by_52pojie.sh脚本路径 -->
    <string>/Users/xxx/Public/MyShell/reset_navicat.sh</string>
    <key>ProgramArguments</key>
    <array>
    <!-- 3.修改为你的reset_navicat.sh脚本的绝对路径,或者reset_navicat_by_52pojie.sh脚本路径 -->
        <string>/Users/xxx/Public/MyShell/reset_navicat.sh</string>
    </array>
  	<!-- 在加载该文件时就执行任务，如果不需要可以删掉或改为false，调试阶段建议打开，以便查看脚本执行结果 -->
    <key>RunAtLoad</key>
    <true/>
    <!-- 在指定时间执行任务 -->
    <key>StartCalendarInterval</key>
    <dict>
        <!-- 3.下面表示每天9点10分执行任务,可以修改为你想要的时间段执行任务 -->
        <key>Hour</key>
        <integer>9</integer>
        <key>Minute</key>
        <integer>10</integer>
    </dict>
    <!-- 运行日志，请以实际为准，调试阶段建议打开，以便查看脚本执行结果 -->
    <key>StandardOutPath</key>
      <!-- 4.运行日志，修改为你的实际想要输出的日志路径 -->
    <string>/Users/xxx/Public/MyShell/reset_navicat.log</string>
    <!-- 错误日志，请以实际为准，调试阶段建议打开，以便查看脚本执行结果 -->
    <key>StandardErrorPath</key>
    <!-- 5.错误日志，修改为你的实际想要输出的日志路径 -->
    <string>/Users/xxx/Public/MyShell/reset_navicat_error.log</string>
</dict>
</plist>
```

4. 最后打开终端，切换到当前目录，依次执行下面的命令加载定时任务

```bash
# 为reset_navicat.sh文件授予可执行权限
chmod u+x reset_navicat.sh
# 将com.chaofan.reset.navicat.premium.trial.period.plist复制到~/Library/LaunchAgents文件夹中，当前用户登录后便会自动加载该定时任务
cp reset.navicat.premium.trial.period.plist ~/Library/LaunchAgents/reset.navicat.premium.trial.period.plist
# 加载定时任务，如果没有报错则任务就加载成功了，会按照计划执行重置脚本，如果上面开启了加载即执行任务和任务日志输出，此时可以去查看日志文件，获取脚本执行情况
launchctl load -w ~/Library/LaunchAgents/reset.navicat.premium.trial.period.plist
# 如果要调整plist文件或是停止任务，请执行以下命令后再进行调整，更多launchctl使用技巧请看文末的参考链接
launchctl unload -w ~/Library/LaunchAgents/reset.navicat.premium.trial.period.plist
```

## 方案三

如果你觉得麻烦,请直接使用之前已经激活版本,有可能会遇到各种奇奇怪怪的问题,地址:[Mac 最新版Navicat v16.2.9 安装与永久激活教程](https://mp.weixin.qq.com/s?__biz=Mzg3ODA5ODY3MQ==&mid=2247498853&idx=1&sn=e638f7282e61c94d1f6069160d3712cb&chksm=cf1a5de5f86dd4f362cd5ad378d1624f19efb3d092c00f3d80926a5b3d59c38e15df200364d0&token=608821632&lang=zh_CN#rd)

## 软件使用帮助

打开软件提示**已损坏，无法打开，移动到废纸篓**，这是因为百度网盘的问题。

![image-20230422222731886](https://billy.taoxiaoxin.club/md/2023/11/6542135fabbcbe03e179f570.png)

看到这个在终端粘贴复制下面的命令（**注意最后有一个空格**）**先不要按回车！先不要按回车！先不要按回车！先不要按回车！**：

```sh
sudo xattr -r -d com.apple.quarantine 
```

![image-20231012153603582](https://billy.taoxiaoxin.club/md/2023/11/6542135f32b52e490b444908.png)

然后打开 **“访达”（Finder）**进入 **“应用程序”** 目录，找到**Navicat Premium App** 图标，将图标拖到刚才的终端窗口里面

![image-20231012154231391](https://billy.taoxiaoxin.club/md/2023/11/6542135f4754429f87eb0e8c.png)

会得到如下组合(如图所示),最后回到终端窗口按回车，输入系统密码回车即可：

```sh
sudo xattr -r -d com.apple.quarantine "/Applications/Navicat Premium.app"
```

![image-20231012154341228](https://billy.taoxiaoxin.club/md/2023/11/6542135f2f96e99cb4dfa1d7.png)

接着我们**Navicat Premium** APP，选择打开即可

![image-20231012155522310](https://billy.taoxiaoxin.club/md/2023/11/6542135f8962d227a452ba61.png)

如果还是不行,试试禁用Gatekeeper，对于 macOS 10.12 - 10.15.7 执行如下命令:

```bash
sudo spctl --master-disable
```


对于 macOS 11 及更高版本

```bash
sudo spctl --global-disable
```

还是不行就自行百度吧！

## 软件下载地址

**最后，希望本教程能够帮助到你，有用请点个赞！**

文中提到的软件，脚本下载地址如下：

```go
链接: https://pan.baidu.com/s/1kUhTtoN3lX-K1L85efiN0A?pwd=3342 提取码: 3342
```

<img src="https://billy.taoxiaoxin.club/md/2023/11/6542740cc8e322da3e8e0f45.png" alt="image-20231101235140389" style="zoom:50%;" />
