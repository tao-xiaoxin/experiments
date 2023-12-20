# 全自动薅羊毛神器，一键部署青龙面板，全自动领取京东京豆、阿里云盘会员...

## 缘起

哈喽，大家好！我是小欣。

相信大家在平时会为了参加一些活动，在各个APP进行签到等活动，每天都要自己手动操作，任务搞的人很烦，每天都要花费很多时间和精力。

所以，我们为何不将重复的事情交给机器去做呢？

于是，搭建一个青龙面板自动化处理这些事情就觉得非常有必要了。

## 青龙面板介绍

青龙面板是一个功能强大的 web 可视化任务管理系统，具有以下基本功能：

1. **Docker 安装支持：** 提供通过 Docker 方式进行安装，方便部署和管理。
2. **多语言脚本支持：** 可执行多种脚本语言，包括 Python 3、JavaScript、Shell、TypeScript 等，为用户提供灵活的任务执行环境。
3. **在线管理功能：** 提供在线管理脚本、环境变量和配置文件的功能，使用户可以方便地进行任务配置和管理。
4. **任务日志查看：** 支持在线查看任务日志，帮助用户监控任务执行过程和结果。
5. **秒级任务设置：** 提供秒级任务设置，使用户可以精确控制任务的执行时间和频率。
6. **系统级通知：** 支持系统级通知，方便管理员或用户及时了解系统状态和任务执行情况。
7. **手机端操作：** 提供手机端操作支持，使用户可以通过移动设备方便地管理和监控任务。

## 前提条件

首先你需要**一台服务器**，或者24小时不关机的电脑也可以。

## 服务器购买

有服务器的可以跳过这一步，**没有服务器的可以跟着我操作起来。**

注册一个腾讯云/阿里云账号，复制到浏览器打开如下地址：

```go
腾讯云购买地址:https://curl.qcloud.com/V31JrAAF
阿里云购买地址:https://www.aliyun.com/minisite/goods?userCode=izmoik21
```

阿里云新用户一年是目前是87元,老用户是99一年

腾讯云新用户是88元，看自己需求选择。

这里**以腾讯云为例**，选择第一个88元点击购买。

![image-20231110173204316](https://billy.taoxiaoxin.club/md/2023/11/654df894e43c8b7ca0623321.png)

会提示你登录,使用**微信或者QQ扫码**登录即可.

登录成功后,再次点击立即购买,**会提示你实名认证**,先去认证即可.

![图片](https://billy.taoxiaoxin.club/md/2023/11/654df9ae3c74766e985f4003.png)



![image-20231110173956744](https://billy.taoxiaoxin.club/md/2023/11/654dfa6c6d8a09a9088d1aa2.png)

点击同意,最后点击支付即可.

![图片](https://billy.taoxiaoxin.club/md/2023/11/654dfa8a79fe82094e580e26.png)

购买成功之后,在搜索框输入轻量应用服务器.

![图片](https://billy.taoxiaoxin.club/md/2023/11/654dfab73fac996ff490be3e.png)

进入控制台，就看到你刚刚购买的服务器了,点击进入你的实例。

![图片](https://billy.taoxiaoxin.club/md/2023/11/654dfab72a66d1034f626fb3.png)

点击进入管理服务器

![image-20231110174535942](https://billy.taoxiaoxin.club/md/2023/11/654dfbc081097a5897463a7d.png)

添加防火墙规则

![image-20231110174738291](https://billy.taoxiaoxin.club/md/2023/11/654dfc3a08887476bd6b71af.png)

放开22 端口号

![image-20231110175059683](https://billy.taoxiaoxin.club/md/2023/11/654dfd03cacbba7bdf016b7b.png)

## 连接服务器

下载**FinalShell SSH**工具，支持**Windows、macOS、Linux**,并且**完全免费**，官网地址如下：

```go
https://www.hostbuf.com/t/988.html
```

安装好软件以后，复制服务器公网`IP`

![image-20231110181954478](https://billy.taoxiaoxin.club/md/2023/11/654e03ca41565e7e5eceaa31.png)

接着打开**FinalShell SSH**工具，如图新建SSH 链接。

![image-20231110182150676](https://billy.taoxiaoxin.club/md/2023/11/654e043eb1b5b2f61d863ad0.png)

填写ssh信息如下：

![image-20231110182842892](https://billy.taoxiaoxin.club/md/2023/11/654e05db48976dbe709b94c1.png)

不知道服务器密码到服务器控制台重置密码。

![image-20231110183001366](https://billy.taoxiaoxin.club/md/2023/11/654e0629a58dec70f2648f01.png)

继续连接到服务器，**点击刚刚新建的服务器信息**即可连接到服务器，连接后**点击接受并且保存**。

![image-20231110183219134](https://billy.taoxiaoxin.club/md/2023/11/654e06b3f17e445ced8db4b8.png)

## 安装docker 

如果你之前安装过 docker，请先删掉：

```bash
yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-engine
```

安装依赖，下载 repo 文件，并把软件仓库地址替换为镜像站：

添加清华镜像源：

```bash
 yum install -y yum-utils
yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
sed -i 's+https://download.docker.com+https://mirrors.tuna.tsinghua.edu.cn/docker-ce+' /etc/yum.repos.d/docker-ce.repo
```

最后安装docker：

```bash
yum install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

验证是否安装成功：

```go
docker -v
```

![image-20231110183734380](https://billy.taoxiaoxin.club/md/2023/11/654e07eeb8aa68b382444dc7.png)

Docker配置开机启动

```javascript
sudo systemctl start docker  #启动Docker
sudo systemctl enable docker #配置开机自启
```

## 安装青龙面板

首先，在服务器内部执行如下命令,永久关闭服务器内部防火墙：

```bash
# 关闭防火墙
systemctl stop firewalld
# 永久关闭防火墙
systemctl disable firewalld
```

使用以下命令从Docker Hub上拉取青龙面板的Docker镜像。

```bash
docker pull whyour/qinglong:latest
```

创建青龙面板容器，并设置必要的环境变量和端口映射，执行如下命令：

```bash
docker run -dit \
  -v /home/docker/ql/config:/ql/config \
  -v /home/docker/ql/log:/ql/log \
  -v /home/docker/ql/db:/ql/db \
  -p 5700:5700 \
  --name qinglong \
  --hostname qinglong \
  whyour/qinglong:latest
```

这个命令解释如下：

- `docker run`: 运行一个新的容器。
- `-dit`: 分别表示以交互模式运行容器，并在后台运行。
- `-v /home/docker/ql/config:/ql/config`: 将本地主机的 `/home/docker/ql/config` 目录映射到容器内的 `/ql/config` 目录，这样配置文件可以在主机和容器之间共享。
- `-v /home/docker/ql/log:/ql/log`: 将本地主机的 `/home/docker/ql/log` 目录映射到容器内的 `/ql/log` 目录，用于存储日志文件。
- `-v /home/docker/ql/db:/ql/db`: 将本地主机的 `/home/docker/ql/db` 目录映射到容器内的 `/ql/db` 目录，用于存储数据库文件。
- `-p 5700:5700`: 将主机的5700端口映射到容器的5700端口，允许通过主机的5700端口访问容器内的服务。
- `--name qinglong`: 为容器指定一个名称为 "qinglong"。
- `--hostname qinglong`: 设置容器的主机名为 "qinglong"。
- `whyour/qinglong:latest`: 使用指定的Docker镜像 `whyour/qinglong`，并选择其最新版本（`latest`标签）。

看到如下提示则证明容器启动成功：

![image-20231110185013155](https://billy.taoxiaoxin.club/md/2023/11/654e0ae5b3cd1de7bd926dc1.png)

设置开放青龙面板端口，到服务器防火墙控制台界面：

![image-20231110185536710](https://billy.taoxiaoxin.club/md/2023/11/654e0c29066c0d52fed6e996.png)

现在可以通过浏览器访问青龙面板了。在浏览器中输入`http://<服务器IP地址>:5700`，即可访问青龙面板的Web界面。

![image-20231110185756692](https://billy.taoxiaoxin.club/md/2023/11/654e0cb50c01e2b1934a476c.png)

按照引导设置用户名和密码：

![image-20231110224112982](https://billy.taoxiaoxin.club/md/2023/11/654e4109c5ea3407e0ec2b65.png)

通知方式暂时跳过，最后我们就安装好啦！

![image-20231110224236871](https://billy.taoxiaoxin.club/md/2023/11/654e415d6b26a5ee60e40e45.png)

## 配置域名解析

如果你有备案过的域名，可以将服务器IP和域名绑定。

在服务器控制台复制服务器`IP`

![image-20231110181954478](https://billy.taoxiaoxin.club/md/2023/11/654e444081e625fd78ca2692.png)

打开腾讯云搜索，搜索**云解析 DNS**，进入控制台

![image-20231110225702318](https://billy.taoxiaoxin.club/md/2023/11/654e44be1cbc34617b78251d.png)

接着点击解析

![image-20231110230034461](https://billy.taoxiaoxin.club/md/2023/11/654e4592b61428324563c89d.png)

添加记录，输入记录值等信息。

![image-20231110230331023](https://billy.taoxiaoxin.club/md/2023/11/654e46439079b2789c346414.png)

![image-20231110230417838](https://billy.taoxiaoxin.club/md/2023/11/654e46729046d7f9b296e2a5.png)

执行如下命令一键安装 `1Panel`:

```go
curl -sSL https://resource.fit2cloud.com/1panel/package/quick_start.sh -o quick_start.sh && sudo bash quick_start.sh
```

按照指引提示操作即可。最后在浏览器中打开`1Panel `控制台。这个和宝塔面板差不多的，看起来页面比宝塔面板更加美观。

![image-20231110230914777](https://billy.taoxiaoxin.club/md/2023/11/654e479bdcc8b54812c5712b.png)

点击添加网站

![image-20231110231129185](https://billy.taoxiaoxin.club/md/2023/11/654e48217e59ed2ada85db63.png)

配置反向代理

![image-20231110232520676](https://billy.taoxiaoxin.club/md/2023/11/654e4b61117b86c873fc9ae9.png)

记得配置的时候去掉`HTTP://`，图中错误。

最后通过域名访问你的青龙面板就好啦。

## 配置HTTPS证书

打开腾讯云如下地址,申请ssl 证书:

```bash
https://console.cloud.tencent.com/ssl
```

打开后点击申请免费证书

![image-20231111002057481](https://billy.taoxiaoxin.club/md/2023/11/654e58690a5e50a404a9c2db.png)

![image-20231111002136639](https://billy.taoxiaoxin.club/md/2023/11/654e58903c7852f7173d205d.png)

输入你的域名,点击申请验证

![image-20231111002519507](https://billy.taoxiaoxin.club/md/2023/11/654e596f4758f05f205064d8.png)

通过后点击下载,选择Nginx 格式

![image-20231111002728968](https://billy.taoxiaoxin.club/md/2023/11/654e59f1812dbd7d7ddcad3e.png)

![image-20231111002855508](https://billy.taoxiaoxin.club/md/2023/11/654e5a47a9a969b9bc3b5921.png)

解压后用文本编辑器打开.pem和.key 结尾的文件,点击上传证书,复制粘贴到对应的文本框中。

![image-20231111003419834](https://billy.taoxiaoxin.club/md/2023/11/654e5b8c3d4cc2d0ec3418f2.png)

最后回到网站开启HTTPS。

![image-20231111003712471](https://billy.taoxiaoxin.club/md/2023/11/654e5c38b070070966986a5c.png)

