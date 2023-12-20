哈喽,大家好!

Navicat是一款常用的数据库管理软件，

本教程将介绍**Navicat 16最新版在Mac电脑上的安装与永久激活教程**。

## 特别声明:

+ **本教程仅供个人学习和研究使用**
+ Navicat 官网地址: https://www.navicat.com/en/

## 软件介绍

Navicat Premium for Mac是一款功能强大的数据库管理工具，具有多重连接的特点。无论是专业开发人员还是数据库服务器的新手，都能够轻松学习和使用它。

Navicat Premium支持连接到目前主流数据库的所有版本,包括MySQL、SQL Server、SQLite、Oracle和PostgreSQL等。这意味着你可以使用单一程序连接并管理不同类型的数据库，从而使数据库管理更加方便。

此外，Navicat Premium还提供了许多有用的功能，如数据可视化和操作的便捷性，以及对数据安全性的重视。它不仅可以帮助你快速执行查询和数据操作，还可以保护数据的安全性，从而避免数据泄露和不良后果。

总之，Navicat Premium for Mac是一个强大且易于学习和使用的数据库管理工具，提供了多重连接和支持多种数据库的特点，以及许多有用的功能，使数据库管理变得更加轻松和高效。

## 温馨提示

开始之前备份数据库配置文件等文件，终端执行如下命令：

```bash
# 创建Navicat 目录
mkdir -p ~/Desktop/Navicat 
# 复制公共工作目录
cd ~/Library/Application\ Support/PremiumSoft\ CyberTech/Navicat\ CC
cp -r Common/ ~/Desktop/Navicat
```

## 安装

首先下载我给大家提供的最新办的安装包(**下载地址文末**)

目前最新版本为 **V16.3.2 版本**

![image-20231117163224092](https://billy.taoxiaoxin.club/md/2023/11/6557251890f9f284b751b3ab.png)

直接点击安装

![image-20230422222252627](https://billy.taoxiaoxin.club/md/2023/11/65571370338868b2290ce10d.png)

将左边应用拖到Applications即可安装![image-20230422222436164](https://billy.taoxiaoxin.club/md/2023/11/65571370eb991d97f87d2c2e.png)

然后打开软件,可能会提示如下:

![image-20230422222731886](https://billy.taoxiaoxin.club/md/2023/11/6557137060478b07df6a53d2.png)

看到这个不要慌,点击**已损坏修复按钮即可**

![image-20230422223004527](https://billy.taoxiaoxin.club/md/2023/11/65571370335d80036b4ee266.png)

然后点击打开

![image-20230422223045718](https://billy.taoxiaoxin.club/md/2023/11/6557137054d42f2de468087e.png)

此时需要输入你的Mac 密码,等待修复完成

![image-20230422223206415](https://billy.taoxiaoxin.club/md/2023/11/655713702c0ae3409c0276f3.png)

修复完成后关闭窗口

![image-20230422223351882](https://billy.taoxiaoxin.club/md/2023/11/65571370fb175b19f3ff755a.png)

此时再次打开软件,已经可以正常使用了

![image-20230422224051759](https://billy.taoxiaoxin.club/md/2023/11/655713701f476817ee626f84.png)



## 激活

软件**安装好后就已经帮你自动激活了**,需要自己再去手动激活

并且软件**安装好就已经是中文版本了**,对中文不好的人特别友好

![image-20230422224518522](https://billy.taoxiaoxin.club/md/2023/11/6557137034d89c9d567acc7d.png)

## Mac 各种错误解决

### 打开允许“允许任何来源”

如何打开允许任何来源？在 Finder 菜单栏选择 【前往】 – 【实用工具 】，找到【终端】程序，双击打开，在终端窗口中输入：

```bash
sudo spctl --master-disable
```

![在这里插入图片描述](https://billy.taoxiaoxin.club/md/2023/11/65571d06e9d5376bfff85045.png)

输入代码后，按【return 回车键】，这时候会提示输入密码：直接输入自己的电脑密码，然后按【return 回车键】即可， （提示：在输入密码的时候，终端不会有任何显示。密码为开机密码，不要错误）

![在这里插入图片描述](https://billy.taoxiaoxin.club/md/2023/11/65571d0d7a204ffed05d4462.png)

关闭【终端】，重新打开 【系统偏好设置】 – 【安全性与隐私】 – 【通用】 中就会出现且选中 【允许任何来源】

### 文件提示损坏解决办法

方案一(**解除隔离**，绕过签名)：

+ **启动终端实用程序，使用命令：`sudo xattr -r -d com.apple.quarantine`输入一个空格，然后将损坏的程序拖到终端窗口中。团队将添加它。**

方案二（**禁用 Gatekeeper**）：

启动终端实用程序并运行命令

> 对于 macOS 10.12 – 10.15.7，执行如下命令：

```
sudo spctl --master-disable
```

> 对于 macOS 11 及更高版本，执行如下命令：

```
sudo spctl --global-disable
```

**无需将任何内容拖到终端中。要启用 Gatekeeper，请将“disable”替换为“enable”。**

### 应用签名方法

1. **先安装Command Line Tools 工具，打开终端工具输入如下命令：**

   ```bash
   xcode-select --install
   ```

2. **弹出安装窗口后选择继续安装，安装过程需要几分钟，请耐心等待。**
3. **打开终端工具输入并执行如下命令对应用签名：**

```bash
sudo codesign --force --deep --sign - (应用路径)
```

> **注意：应用路径是「访达（Finder）->应用程序」找到应用将其拖进终端命令 – 的后面，然后按下回车键，输入macOS的密码然后按回车(输入过程中密码是不显示的，输入完密码直接按回车键即可！)**

4. **出现 「replacing existing signature」 提示即成功！**

## Navicat 各种问题解决

### 数据库工作区文件存储位置

参考官网文档：https://help.navicat.com/hc/zh-cn/articles/360036036792

### 数据库配置文件路径

路径如下：

```bash
# 数据库配置文件
~/Library/Application Support/PremiumSoft CyberTech/Navicat CC/Common/Settings/0/0

# 公共目录文件
~/Library/Application Support/PremiumSoft CyberTech/Navicat CC/Common/
```

### 保存密码失败

1. 切换到英文版本，去官网下载英文版本，保存密码后重新安装激活版，然后导入配置。
2. 切换到官网版本，保存密码后重新安装激活版，然后导入配置。
3. 参考知乎链接：https://zhuanlan.zhihu.com/p/637861715

## 往期内容

如果没有成功，请看看之前的版本，毕竟是个软件，能用就行。

*往期内容：*

1. [Navicat Mac App Store 版本和无限试用大法](https://mp.weixin.qq.com/s?__biz=Mzg3ODA5ODY3MQ==&mid=2247498910&idx=1&sn=29d896a81dd97b49710fb67ec9700c23&chksm=cf1a5d1ef86dd40825a8d7760841f288fee9a70d9e06d80c5bd9afa2700384473484d5e270be&token=187072541&lang=zh_CN#rd)
2. [Mac Navicat v16.2.9 安装与永久激活教程](https://mp.weixin.qq.com/s?__biz=Mzg3ODA5ODY3MQ==&mid=2247498853&idx=1&sn=e638f7282e61c94d1f6069160d3712cb&chksm=cf1a5de5f86dd4f362cd5ad378d1624f19efb3d092c00f3d80926a5b3d59c38e15df200364d0&token=187072541&lang=zh_CN#rd)

推荐使用无限重置大法，稳定，安全，可靠！！！

## 软件下载地址

下载地址链接: https://www.123pan.com/s/WaG1jv-xl6Qh.html

提取码:3342

<img src="https://billy.taoxiaoxin.club/md/2023/11/655714c353e5570f95760c78.png" alt="image-20231117152243344" style="zoom: 67%;" />



**好了今天的分享会就到这里,有用请点个赞，散会 !**

最后再推荐下我的羊毛群。

<img src="https://billy.taoxiaoxin.club/md/2023/11/6557256e2ee283066b199533.png" alt="image-20231117163350536" style="zoom:25%;" />

<img src="https://billy.taoxiaoxin.club/md/2023/11/6557257637be459bc0de4656.png" alt="image-20231117163358475" style="zoom:25%;" />

<img src="https://billy.taoxiaoxin.club/md/2023/11/6557257e9482336642ffae51.png" alt="image-20231117163406710" style="zoom: 25%;" />

二维码失效了可以扫码添加下方机器人自动邀请入群

<img src="https://billy.taoxiaoxin.club/md/2023/11/655725f173cd472578238393.jpeg" alt="img" style="zoom:50%;" />