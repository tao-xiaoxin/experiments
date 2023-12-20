# 潘多拉：让你呼吸顺畅的ChatGPT！

哈喽，大家好，我是小欣！

ChatGPT，一款强大的AI，然而对于国内用户而言，使用它却如履薄冰，面临诸多困扰，如网站访问受限、易遭封号、响应迟缓、输出截断等。为化解这些困境，给大家推荐一款开源项目——Pandora，它能让你畅所欲言地驾驭ChatGPT。

## Pandora-next 简介

潘多拉 (Pandora)，一个让你呼吸顺畅的 ChatGPT。

潘多拉实现了网页版 ChatGPT 的主要操作。后端优化，绕过 Cloudflare，速度喜人。

**其实就是绕过ChatGPT 种种限制，让你你自己的ChatGPT账号免费体验ChatGPT**！

项目地址：https://github.com/pandora-next/deploy

体验地址： [https://chat.zhile.io](https://chat.zhile.io/)

dockerhub镜像仓库地址：https://hub.docker.com/r/pengzhile/pandora-next

![image-20231115134755601](https://billy.taoxiaoxin.club/md/2023/11/65545b8c54ebaa06ba7747f4.png)

## 本地部署教程

首先，安装docker ，官网下载：

```bash
https://www.docker.com/get-started/
```

![image-20231115135011607](https://billy.taoxiaoxin.club/md/2023/11/65545c132d0b1f54f17979a1.png)

下载后以后启动docker,打开终端验证是否安装成功:

```bash
docker -v
```

![image-20231115135129527](https://billy.taoxiaoxin.club/md/2023/11/65545c6137a1b1aa98e868a7.png)

拉取镜像`pandora`,终端执行命令如下:

```bash
docker pull pengzhile/pandora
```

执行如下命令启动服务:

```bash
docker run  -e PANDORA_CLOUD=cloud -e PANDORA_SERVER=0.0.0.0:8888 -p 8888:8888 -d pengzhile/pandora
```

![image-20231115142335366](https://billy.taoxiaoxin.club/md/2023/11/655463e776ad0fce8cc3c6da.png)

打开浏览器在地址栏输入

```
http://127.0.0.1:8888
```

**点击输入你的ChatGPT用户名和密码,点击登录即可**

![image-20231115142445021](https://billy.taoxiaoxin.club/md/2023/11/6554642d716d3973ca413a64.png)

登录后就和ChatGPT官网一样了,你可以愉快的使用ChatGPT,也不需要什么付费的key之类的。

![image-20231115142857171](https://billy.taoxiaoxin.club/md/2023/11/65546529593d191ed1d5b2ec.png)

官网的历史记录同样会被保存记录，这种方式就是绕过了GPT的Cloud flare的IP检测，接下来你就可以畅快使用啦，再也没有网络限制的烦恼啦，而且速度极快。

更多请参考CSDN博文：https://blog.csdn.net/kingxzq/article/details/132401393

**好了，今天的分享会就到这里，散会，记得点个赞奥！**