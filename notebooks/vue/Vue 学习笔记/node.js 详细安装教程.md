

[TOC]



## 一.下载

### 1.1 官网下载

官网下载地址: [Node.js (nodejs.org)](https://nodejs.org/en) 

![image-20230725222345789](http://billy.taoxiaoxin.club/md/2023/07/64bfdaf2922ee414630350f4.png)

可以看到当前的版本 `LTS`是大多用户使用的稳定版本, `Current`是最新版本, 这里选择的是稳定版本(18.17.0)

### 1.2 国内下载

由于国内下载比较慢,建议使用国内地址:[https://nodejs.cn](https://nodejs.cn/download/),选择你对应的平台和系统安装

![image-20230726193507125](https://billy.taoxiaoxin.club/md/2023/07/64c104eb922ee429b5cf9869.png)

## 二. Windows 安装

### 2.1 安装

1.打开下载安装的文件夹下的安装包, 双击进行安装

![image-20230726193913315](https://billy.taoxiaoxin.club/md/2023/07/64c105e1922ee42a1214a561.png)

2.点击next下一步

![image-20230726194022115](https://billy.taoxiaoxin.club/md/2023/07/64c10626922ee42a8daf4db5.png)

3.勾选同意协议,点击下一步next

![image-20230726194147907](https://billy.taoxiaoxin.club/md/2023/07/64c1067c922ee42aa9b794d5.png)

4.选择安装的位置 ,然后点击下一步next

![image-20230726194255108](https://billy.taoxiaoxin.club/md/2023/07/64c106bf922ee42ac4a0a612.png)

5.点击下一步next

![image-20230726194352423](https://billy.taoxiaoxin.club/md/2023/07/64c106f8922ee42ad7624b46.png)

6.看到如下图所示, 不要勾选!不要勾选! 不要勾选! 直接next

![image-20230726194507915](https://billy.taoxiaoxin.club/md/2023/07/64c10744922ee42af24f31be.png)

7.点击安装install

![image-20230726194548447](https://billy.taoxiaoxin.club/md/2023/07/64c1076c922ee42b059a12ea.png)

最后点击Finish

![image-20230726194810518](https://billy.taoxiaoxin.club/md/2023/07/64c107fa922ee42b374977a7.png)

### 2.2 配置环境变量

点击此电脑--->点击高级--->找到电脑的环境变量

![image-20230727185257737](https://billy.taoxiaoxin.club/md/2023/07/64c24c8a922ee45bc3bf147a.png)

双击Path

![image-20230727185710873](https://billy.taoxiaoxin.club/md/2023/07/64c24d87922ee45c2310b7b8.png)

然后点击新建

![image-20230727185758863](https://billy.taoxiaoxin.club/md/2023/07/64c24db7922ee45c3462a5c8.png)

输入Node 安装路径,最后点击确定

![image-20230727190218636](https://billy.taoxiaoxin.club/md/2023/07/64c24eba922ee45c93f91212.png)

## 三.macOS 安装

点击下载安装包即可

![image-20230727200921540](https://billy.taoxiaoxin.club/md/2023/07/64c25f5d922ee461cf733d52.png)

点击下载好的安装包进行安装

![image-20230727201018814](https://billy.taoxiaoxin.club/md/2023/07/64c25f59922ee461cc136674.png)

接下来一路继续即可:

![image-20230727201136158](https://billy.taoxiaoxin.club/md/2023/07/64c25f55922ee461c971fd91.png)

![image-20230727201157784](https://billy.taoxiaoxin.club/md/2023/07/64c25f51922ee461c6c36517.png)

![image-20230727201229270](https://billy.taoxiaoxin.club/md/2023/07/64c25f4d922ee461c49bf091.png)

![image-20230727201349043](https://billy.taoxiaoxin.club/md/2023/07/64c25f7d922ee461d845fca9.png)

## 四.Linux 安装

首先安装

选中Linux 版本,然后右键复制链接地址

![image-20230727212956524](https://billy.taoxiaoxin.club/md/2023/07/64c27154922ee4693707efb8.png)

执行如下命令:

```bash
cd /usr/local/bin
wget https://npmmirror.com/mirrors/node/v18.17.0/node-v18.17.0-linux-x64.tar.xz
```

![image-20230727203604413](https://billy.taoxiaoxin.club/md/2023/07/64c264b4922ee4635a5305af.png)

继续执行如下命令:

```bash
tar xvf node-v18.17.0-linux-x64.tar.xz
```

![image-20230727213138328](https://billy.taoxiaoxin.club/md/2023/07/64c271ba922ee46959069441.png)

重命名文件夹:

```bash
mv node-v18.17.0-linux-x64 node-v18.17.0
```

删除压缩包:

```bash
rm -rf node-v18.17.0-linux-x64.tar.xz
```

将Node 加入环境变量:

```bash
export PATH="export PATH=$PATH:/usr/local/bin/node-v18.17.0/bin:$PATH" >> ~/.bashrc
source ~/.bashrc  
```

现在，你可以在任何位置运行 `node`、`npm` 和 `npx` 命令，它们将会生效。=

## 五.验证安装

安装完成后，你可以验证 Node.js 和 npm 是否正确安装：

```bash
node -v
npm -v
```

## 六. 修改下载位置

### 1.查看npm默认存放位置

查看npm全局模块的存放路径

```bash
npm get prefix
```

查看npm缓存默认存放路径

```bash
npm get cache
```

### 2.修改默认文件夹

在 nodejs 安装目录下，创建 `node_global` 和 `node_cache` 两个文件夹

设置全局模块的安装路径到 “node_global” 文件夹 `npm config set prefix `命令

```bash
npm config set prefix  "nodejs安装目录/node_global"
```

设置缓存到 “node_cache” 文件夹 `npm config set cache `命令

```bash
npm config set cache "nodejs安装目录/node_cache"
```

### 3.查看是否修改成功

使用命令

```bash
npm install express -g
```

## 七.设置淘宝镜像

将npm默认的registry修改为淘宝registry, 查看当前的镜像

```java
npm config get registry
```

设置淘宝镜像

```java
npm config set  registry  https://registry.npmmirror.com/
```

安装淘宝源cnpm

```bash
npm install -g cnpm --registry=https://registry.npmmirror.com 
```

执行命令查看cnpm模块
```bash
cnpm -v
```

