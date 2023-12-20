哈喽,大家好!

最近收到一部分小伙伴反馈 Mac 更新到MacOs 14之后打开Navicat闪退等问题

所以本教程将介绍**Navicat v16.2.9 最新版在Mac电脑上的安装与永久激活教程**。

## 特别声明:

+ **本教程仅供个人学习和研究使用**
+ Navicat 官网地址: https://www.navicat.com/en/

## 软件介绍

Navicat Premium for Mac是一款功能强大的数据库管理工具，具有多重连接的特点。无论是专业开发人员还是数据库服务器的新手，都能够轻松学习和使用它。

Navicat Premium支持连接到目前主流数据库的所有版本,包括MySQL、SQL Server、SQLite、Oracle和PostgreSQL，Redis等。这意味着你可以使用单一程序连接并管理不同类型的数据库，从而使数据库管理更加方便。

此外，Navicat Premium还提供了许多有用的功能，如数据可视化和操作的便捷性，以及对数据安全性的重视。它不仅可以帮助你快速执行查询和数据操作，还可以保护数据的安全性，从而避免数据泄露和不良后果。

总之，Navicat Premium for Mac是一个强大且易于学习和使用的数据库管理工具，提供了多重连接和支持多种数据库的特点，以及许多有用的功能，使数据库管理变得更加轻松和高效。

![image-20231012161103721](https://billy.taoxiaoxin.club/md/2023/10/6527aa17ef2d66adda3d0275.png)

## 版本 16.2.9 更新日志

错误修正：

+ 将数据导入 SQL Server 时，出现 "INSERT 语句中的行值表达式数量超过允许的最大 1000 行值数量 "错误。
+ 在 macOS 14 上崩溃。
+ 数据传输期间多次传输 MariaDB 触发器。

**温馨提示:**如果你之前**装过Navicat,请备份你的Navicat 数据库设置**，这个自行百度。

## 安装

首先打开终端，开启“任何来源”，执行如下命令,

```go
sudo spctl  --master-disable
```

然后回车，继续输入密码（密码输入时是不可见的），然后回车。

![image-20231012151358747](https://billy.taoxiaoxin.club/md/2023/10/65279cb60e9bdda1b9c77f26.png)

接着打开【**系统偏好设置**】，选择【**安全性与隐私**】，选择【**通用**】，可以看到【**任何来源**】已经选定。

![image-20231012152846076](https://billy.taoxiaoxin.club/md/2023/10/6527a02e477f727456e782fc.png)

继续下载我给大家提供的最新办的安装包(**下载地址文末**)，直接点击安装。

![image-20230422222252627](https://billy.taoxiaoxin.club/md/2023/10/65278f2c6561d3ddccee5687.png)

将左边应用拖到**Applications**即可安装

![image-20231012150448119](https://billy.taoxiaoxin.club/md/2023/10/65279a901106b531cc59031c.png)

然后先不要打开软件,打开终端执行如下命令:

```sh
sudo xattr -r -d com.apple.quarantine "/Applications/Navicat Premium.app" 
```

按回车键回车，输入开机密码，如图：

![image-20231012152241275](https://billy.taoxiaoxin.club/md/2023/10/65279ec19d8d6ec5179467cf.png)

接下来打开在 Mac 上的“访达”，或者使用快捷键**Command + N** 键，打开一个新的“访达”窗口

![image-20231012145704695](https://billy.taoxiaoxin.club/md/2023/10/652798c0efa463944330de8e.png)

找到应用程序，按住 **Control** 键鼠标点**Navicat Premium App** 图标，然后从**快捷键菜单中选取“打开”*，如图所示：

![image-20231012150809733](https://billy.taoxiaoxin.club/md/2023/10/65279b59157c25cf59a20d84.png)

然后选择打开，如图：

![image-20231012150949417](https://billy.taoxiaoxin.club/md/2023/10/65279bbd0fd4dc957eec6aba.png)

选择打开即可

![image-20231012155522310](https://billy.taoxiaoxin.club/md/2023/10/6527a66abe3df1692bd6a341.png)

如果你以前装过Navicat，并且对数据库设置做了备份，你可以选择导入数据库设置，首次进入Navicat 就会提示你导入：

![image-20231012160859247](https://billy.taoxiaoxin.club/md/2023/10/6527a99b5d0a2511b5aa1104.png)

进入软件后，界面如下：

![image-20231012161603471](https://billy.taoxiaoxin.club/md/2023/10/6527ab430a43f155f412aa0d.png)

## 软件使用帮助

如果打开依旧提示**已损坏，无法打开，移动到废纸篓**，是因为刚刚签名没有下发成功导致的

![image-20230422222731886](https://billy.taoxiaoxin.club/md/2023/10/65278f2c7a297820e61ec005.png)

看到这个不要慌,咱们上绝活，在终端粘贴复制下面的命令（**注意最后有一个空格**）**先不要按回车！先不要按回车！先不要按回车！先不要按回车！**：

```sh
sudo xattr -r -d com.apple.quarantine 
```

![image-20231012153603582](https://billy.taoxiaoxin.club/md/2023/10/6527a1e350e8a574bf0d1ef6.png)

然后打开 **“访达”（Finder）**进入 **“应用程序”** 目录，找到**Navicat Premium App** 图标，将图标拖到刚才的终端窗口里面

![image-20231012154231391](https://billy.taoxiaoxin.club/md/2023/10/6527a367cf4376a79365c445.png)

会得到如下组合(如图所示),最后回到终端窗口按回车，输入系统密码回车即可：

```sh
sudo xattr -r -d com.apple.quarantine "/Applications/Navicat Premium.app"
```

![image-20231012154341228](https://billy.taoxiaoxin.club/md/2023/10/6527a3ad10d9a1b1858ce908.png)

接着我们**Navicat Premium** APP，选择打开即可

![image-20231012155522310](https://billy.taoxiaoxin.club/md/2023/10/6527a66abe3df1692bd6a341.png)

**温馨提示**:如果使用默认系统自带的终端,会有空格换行符等问题导致签名不成功，所以**我这里推荐你使用的命令行终端工具是Tabby**，我用的是这个工具，签名一次就下发成功了，使用默认的终端拖入APP图表会有点问题，依旧提示软件打不开，我换了个命令行终端工具就好了。

## 激活

软件**安装好后就已经帮你自动激活了**,需要自己再去手动激活

并且软件**安装好就已经是中文版本了**,对英文不好的人特别友好

![image-20230422224518522](https://billy.taoxiaoxin.club/md/2023/10/65278f2cd4102490146c4508.png)

## 兼容性

+ **同时兼容Mac M2和M1 芯片！**

## 软件下载地址

**最后，希望本教程能够帮助到你，有用请点个赞！**

下载地址链接: https://pan.baidu.com/s/1UuSy52hNls4IhjKalfuSHA?pwd=3342 

提取码: 3342

![image-20231012163613744](https://billy.taoxiaoxin.club/md/2023/10/6527affe08334fd0c960ad2f.png)

