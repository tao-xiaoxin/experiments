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

目前最新版本为 **V16.3.4 版本**

![](https://billy.taoxiaoxin.club/md/2023/12/65868a834cf75156aea6428b.png)

直接点击安装

![](https://billy.taoxiaoxin.club/md/2023/12/65868cd75c29eefed5c28187.png)

将左边应用拖到`Applications`即可安装

![](https://billy.taoxiaoxin.club/md/2023/12/65868daedb84f45487ad82d6.png)

![](https://billy.taoxiaoxin.club/md/2023/12/65868dc8ab88a0c3949ebcbb.png)

![](https://billy.taoxiaoxin.club/md/2023/12/65868de081e6f78d6b131cde.png)

此时直接打开软件,会提示如下:

![](https://billy.taoxiaoxin.club/md/2023/11/6557137060478b07df6a53d2.png)

接着修复已损坏打开终端，**输入以下命令，先不要回车**。

```bash
sudo codesign --force --deep --sign - 
```

打开`「访达（Finder）->应用程序」`，将Navicat 拖入到命令行后面。

![](https://billy.taoxiaoxin.club/md/2023/12/65869319d8251da164856cf6.png)

最后回车，输入密码。

![](https://billy.taoxiaoxin.club/md/2023/12/65869364c895996c0092a0fe.png)

![](https://billy.taoxiaoxin.club/md/2023/12/658693a00ce05e7eec05288e.png)

最后找到应用程序,右键点击`Navicat`打开。

![](https://billy.taoxiaoxin.club/md/2023/12/658694a4c24b605a8cb513ed.png)

点击打开

![](https://billy.taoxiaoxin.club/md/2023/11/6557137054d42f2de468087e.png)

此时再次打开软件,已经可以正常使用了

![image-20230422224051759](https://billy.taoxiaoxin.club/md/2023/11/655713701f476817ee626f84.png)

## 连接数据库

点击连接-》选择`MySql`(这里以`mysql`为例子)

![](https://billy.taoxiaoxin.club/md/2023/12/6586b9c038fb1281de719e47.png)

![](https://billy.taoxiaoxin.club/md/2023/12/6586bbac1893bb439213883b.png)

输入你的数据库信息，测试链接，最后保存。

![](https://billy.taoxiaoxin.club/md/2023/12/6586bc1a89d8689c81dc30be.png)

此版本保存密码会失败，需要关闭`SIP`才能使用，不过我给大家也准备了能正常保存密码的，在文末地址下载即可。

![](https://billy.taoxiaoxin.club/md/2023/12/6586bc73570d62eca9c0aa3a.png)

## 激活

软件**安装好后就已经帮你自动激活了**,并且软件**安装好就已经是中文版本了**

![image-20230422224518522](https://billy.taoxiaoxin.club/md/2023/11/6557137034d89c9d567acc7d.png)

## Mac 各种错误解决

文件提示损坏解决办法请参考：[Mac 安装软件常见报错与解决方案（汇总)](https://mp.weixin.qq.com/s/FiRPKI9ZMen8J3i4YUZT8g)

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

## 温馨提示

还有本软件不是不是一个数据库，它是一个数据库管理工具，有些人以为装了这个软件就能使用数据库，还有人经常加了问我，所以这里特别说明一下，Mysql 之类的数据库软件需要单独安装。

## 软件下载

下载地址：

```bash
https://www.123pan.com/s/WaG1jv-wQyQh.html
```

提取码:VyvS

好了，今天的分享就到这里，**觉得有用请点个赞+在看！**

## 小商店

+ **GitHub Copilot 激活码购买**：这个主要用于激活GitHub Copilot系列产品，远程激活码价格一年`65`元，两年`80`元。可登录自己账号，激活码一机一码，需要购买请扫描二维码添加好友备注 **GitHub 远程激活**购买。
+ 商品数量有限，卖完为止！

![](https://billy.taoxiaoxin.club/md/2023/12/6586c969c004cbaffe12bf04.png)

