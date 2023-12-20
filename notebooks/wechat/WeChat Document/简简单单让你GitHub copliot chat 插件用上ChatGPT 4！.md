# 简简单单让你GitHub copliot chat 插件用上ChatGPT 4！

哈喽，大家好！

今天教你的GitHub copliot chat 插件配置gpt 4！**再也不用订阅每月昂贵的GPT4 了！**

GitHub copliot chat 插件默认是自动3.5/4模型切换的。

**目前这种方式只对 VsCode 有效！** 

## 前提条件

+ 你开通了 GitHub copliot  系列产品
+ 你的VsCode 安装了GitHub copliot chat  和GitHub copliot 插件
+ 需要GitHub 账号学生认证的可以文末联系我

## 修改chat 使用GPT 4！

直接打开如下地址：

```go
https://txx.lanzoub.com/idEea1hko9vi
```

将文件解压，解压后得到cocopilot-0.0.6.vsix

![](https://billy.taoxiaoxin.club/md/2023/12/65770c5dfda16b7185d7f1f6.png)

接着打开VScode 安装vsix扩展

![](https://billy.taoxiaoxin.club/md/2023/12/65770e5e98e6be3a1cf96fd5.png)

![](https://billy.taoxiaoxin.club/md/2023/12/65770e89f28739bb976b0f93.png)

接下来重载。

![](https://billy.taoxiaoxin.club/md/2023/12/65770eaee79cb398ef8a4017.png)

接着打开选择菜单：`查看` -> `命令面板` 后输入 `Enable Copilot Chat (GPT-4)` 回车确认，按提示重新加载VSCode即可。

![](https://billy.taoxiaoxin.club/md/2023/12/6577100652dedd72f0fc8442.png)

![](https://billy.taoxiaoxin.club/md/2023/12/6577106cc77748c62fd73401.png)

![](https://billy.taoxiaoxin.club/md/2023/12/657710c9808c697276ad2917.png)

回到VSCode 聊天，效果如下：

![](https://billy.taoxiaoxin.club/md/2023/12/6577115ec7cfc7b80c97cdfc.png)

接下来，再教大家一个小技巧，将你的GitHub copliot 分享给你的朋友使用，无需你的账号密码。

## GitHub copliot 拼车

首先，打开如下地址：

```g
https://cocopilot.org/dash
```

会提示你需要授权GitHub 账号，点击授权。

![](https://billy.taoxiaoxin.club/md/2023/12/657720e9d89144c1841908c2.png)

授权后点击管理我开的车：

![](https://billy.taoxiaoxin.club/md/2023/12/65771264e6701d06c3bfa36f.png)

随便起个车队名称，点击获取GitHub token。

![](https://billy.taoxiaoxin.club/md/2023/12/65771304fcc050beff886ee2.png)

点击登录链接，会自动帮你复制设备码

![](https://billy.taoxiaoxin.club/md/2023/12/657713929df0e688cd84a283.png)

粘贴设备码，点击继续。

![](https://billy.taoxiaoxin.club/md/2023/12/657713d2eccd7615a0b6e52a.png)

点击授权。

![](https://billy.taoxiaoxin.club/md/2023/12/65771405efd3317c5405ad04.png)

然后回到刚刚点击Open login url的那个网页，复制你的token。

![](https://billy.taoxiaoxin.club/md/2023/12/657714769a1725b89811d218.png)

回到车队管理网页，粘贴输入你的token。

![](https://billy.taoxiaoxin.club/md/2023/12/657714dcf4006d87c0209664.png)

滑到最底部，点击保存。

![](https://billy.taoxiaoxin.club/md/2023/12/6577151b7c21d8667c07a08d.png)

接下来就是就是让你的朋友加入你的车队，按照以上方式登录授权一遍。

让你的朋友把上面的用户ID发给你，你填写到对应的成员框中。

![](https://billy.taoxiaoxin.club/md/2023/12/6577168dfb4cb86f65d99da6.png)

滑到最底部，点击保存。

![](https://billy.taoxiaoxin.club/md/2023/12/6577151b7c21d8667c07a08d.png)

回到如下地址

```go
https://cocopilot.org/dash
```

下载GitHub Copliot激活插件并且解压。

![](https://billy.taoxiaoxin.club/md/2023/12/65771800603227bd56cef864.png)

使用VSCode 查找替换，windows 打开编辑`cocopilot.bat`脚本，Linux 和Mac则编辑`cocopilot.sh` 脚本

![](https://billy.taoxiaoxin.club/md/2023/12/657719c43598dd4ca3479425.png)

查找输入ghu_ThisIsARealFreeCopilotKeyByCoCopilot

然后替换为你这里的token，这里的token 用各自的就好，就是自己自己的。

![](https://billy.taoxiaoxin.club/md/2023/12/657718cafdf3f21d25a6091a.png)

![](https://billy.taoxiaoxin.club/md/2023/12/65771a6a5e8ce22209700a5a.png)

最后保存文件，执行对应平台的脚本，windows 右键以管理员身份运行`cocopilot.bat`脚本，Mac则使用终端打开`cocopilot.sh` 脚本，接下来重启你的 IDEA 即可。

![](https://billy.taoxiaoxin.club/md/2023/12/65771b560464f19076a437a2.png)

好了，最后来总结一下：

1. 目前拼车不支持激活VScode 插件，这个我试过好几遍都不行。
2. 拼车的话车主必须具有GitHub Copliot 权限，没有的话可以找我帮你GitHub 学生认证，费用你可以几个人平摊一下即可，或者你可以学学那种网上卖远程激活码的，也是和上面拼车一个原理。
3. 所谓拼车，就是平摊copilot的成本（目前10人拼一车，很稳），如果你有copilot想跟人均摊成本，那就快去找人开车吧。

## GitHub 学生代认证

+ 不管你是否毕业，都可以帮你认证，100%秒通过，
+ 充值到自己账号，认证通过后质保一年，稳定使用2年，有官方售后群
+ 现在咨询可能还有点优惠，后面会涨价。
+ Github 学生认证答疑：https://docs.qq.com/doc/DVFFXeU5nc2ZOWGhI

**扫码以下二维码咨询，备注 GitHub 学生包优先通过。**

![](https://billy.taoxiaoxin.club/md/2023/12/65771d788bba5f693a400a60.png)