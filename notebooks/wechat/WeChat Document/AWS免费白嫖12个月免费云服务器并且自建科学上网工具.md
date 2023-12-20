# 背景

AWS海外区云产品免费使用页面：https://aws.amazon.com/cn/free

不知道你想不想拥有一台属于自己的云服务器呢，拥有一台自己的云服务器可以建站或者搭建科学上网节点，今天我就来教大家如何申请亚马逊 AWS 免费云服务器，这个云服务器可以长达12个月的免费。

首先我们打开 AWS海外区云产品免费使用页面：https://aws.amazon.com/cn/free

![home](https://billy.taoxiaoxin.club/md/2023/06/64959a88922ee40362aff313.png)

![free](https://billy.taoxiaoxin.club/md/2023/06/64959a8b922ee403638549cf.png)

然后可以看到AWS提供了非常多的免费云产品。

- Amazon EC2云服务器
- Amazon RDS数据库
- Amazon S3云储存

等等。

而云产品中最常使用的应该就是云服务器了，我这里教大家如何创建AWS账号并申请12个月免费的云服务器，如果你觉得它的云服务器好用，也可以在到期之后选择付费继续使用。

# 注册

要使用这些免费云产品，我们需要先创建一个AWS账号，AWS免费试用页面：https://aws.amazon.com/cn/free

### 1. 注册

点击创建免费账号：

![register](https://billy.taoxiaoxin.club/md/2023/06/64959a8e922ee403648bcdc7.png)

我们输入电子邮箱+账户名称进行注册。

### 2. 验证

然后输入邮箱返回的验证码，点击验证。

![verify](https://billy.taoxiaoxin.club/md/2023/06/64959a8f922ee403654dead7.png)

### 3. 创建密码

下一步就是创建密码，输入密码并点击继续： ![step1](https://billy.taoxiaoxin.club/md/2023/06/64959a91922ee40366f13a41.png)

### 4. 联系人信息

然后填写你的联系人信息，注意填写真实信息。

计划这里要填写`个人`。

![step2](https://billy.taoxiaoxin.club/md/2023/06/64959a92922ee4036716b674.png)

### 5. 信用卡信息

下一步需要填写信用卡信息，这里支持Visa，Mastercard，美国运通等。

它不支持银联，在国内的话我们可以使用招商银行的Visa双币信用卡是可以的,**或者拼多多淘宝之类软件买个课认证一下就可以了**。

> 对于低于 AWS 免费套餐限制的使用量，我们不会向您收取费用。我们可能会以待处理交易的方式暂时扣取最多 1 美元(或等值的当地货币) 3-5 天，以验证您的身份。

我们只要不超过AWS套餐限制，就不会收费。

![step3](https://billy.taoxiaoxin.club/md/2023/06/64959a94922ee40368787173.png)

### 6. 确认身份

在这里需要验证你的电话号码才能继续 ![step4](https://billy.taoxiaoxin.club/md/2023/06/64959a96922ee40369718b62.png)

### 7. 手机验证码

![step5](https://billy.taoxiaoxin.club/md/2023/06/64959a97922ee4036ae336cb.png)

### 8. 完成注册

离注册完成只剩最后一步了，选择支持计划-选择默认的免费就可以了。 ![finish](https://billy.taoxiaoxin.club/md/2023/06/64959a98922ee4036b176aa8.png)

然后转到AWS管理控制台来申请免费云服务器 ![congratulation](https://billy.taoxiaoxin.club/md/2023/06/64959a9a922ee4036cb7cf88.png)

# 启动实例

### 0.选择服务器地区

注册成功之后选择距离你比较近的地区,右上角点击切换地区

<img src="https://billy.taoxiaoxin.club/md/2023/06/6495b038922ee4138fd90ea9.png" alt="image-20230623224616475" style="zoom:50%;" />

你可以选择距离进的新加坡,因为我这里需要访问OPENAI,所以最好选择美国地区(美国西部 俄勒冈州)

![choose](https://billy.taoxiaoxin.club/md/2023/06/64959a9c922ee4036d9d0d7f.png)

右上角选择你想要使用的云服务器的地区，然后下拉选择启动虚拟机。

EC2，也就是虚拟云服务器。基本都是选择这种。

### 1. 选择系统

注意只有下方显示**支持免费套餐**的才会免费。所以很明显macOS是不免费的。

我们选择Ubuntu，这个是最火的Linux发行版，以后想要在上面搭建各种环境也会有更多的教程。

然后选择X86架构，因为ARM架构没有免费的套餐。

![ubuntu](https://billy.taoxiaoxin.club/md/2023/06/64959a9e922ee4036e202810.png)

### 2. 创建密钥对

然后选择密钥对，如果是第一次启动实例，是没有密钥对的，需要自己点击**Create new key pair**来创建。

密钥对名称自己随便填。

密钥对类型选择RSA加密。

密钥文件格式，如果你使用PuTTY这个SSH软件选择 `ppk` 格式，如果是其它的SSH软件或者你是macOS或者Linux系统，全部选择 `pem` 格式。

我用的是Mac系统，所以直接使用Terminal终端连接，所以选择的是**pem**格式。

如果你使用的Windows系统，下载一个免费的SSH客户端即可，使用也非常方便。 这里推荐一下XShell，把密钥导入客户端就可以直接使用。

当然如果你使用的PuTTY这个SSH客户端，就选择 ppk 格式就行。

然后点击创建密钥对，就会下载一个 **pem** 格式的密钥文件，好好保存，以后连接云服务器要用。

![rsa](https://billy.taoxiaoxin.club/md/2023/06/64959aa0922ee4036fe40482.png)

### 3. 网络设置

下一步是网络设置，也就是设置防火墙之类的，我们点击**创建安全组**，你也可以点击**选择现有的安全组**，选择Default。

这里我们创建安全组，勾选这几个选项。

SSH是你通过命令或者SSH客户端连接云服务器要用的。HTTP和HTTPS是你接受网络请求要用的，所以一般来说三个都需要勾选。

![network](https://billy.taoxiaoxin.club/md/2023/06/64959aa3922ee40370947ecd.png)

然后选择磁盘空间，最多免费30GB，然后点击启动实例。

### 4. 配置储存

> 有资格使用免费套餐的客户最多可获得 30GB 的通用型 (SSD) 或磁存储空间

一个实例最低需要8G的磁盘空间，一般来说30G可以创建1～2个示例，根据需求选择。这里我直接选择30G，只需要这一个实例就可以了。

### 5. 启动实例

所有的配置都选择好了，我们检查一下。

看到

> 第一年包括每月免费套餐 AMI 上的 750 小时 t2.micro

实际上每个月最多有744小时，所以完全足够了。如果你同时开启两个实例事时间就会翻倍，所以一般来说只启动一个。

### 6. 设置账单提醒

为了避免亚马逊有意外收费我们无法及时收到通知，建议大家都要勾选接收账单提醒。

![bill](https://billy.taoxiaoxin.club/md/2023/06/64959aa5922ee40372a05751.png)

## 搭建科学上网工具

### 1. 登录云服务器切换到root 用户

如果你是Windows 下载一个Xshell之类的SSH客户端导入密钥就可以连接了。

如果你是macOS，那么就更简单了。直接在终端使用SSH 指令连接就可以。

点击实例ID可以看到你的实例的具体信息，比如。公网IP/共有IPv4 DNS，我们可以直接使用这个DNS连接

```sh
ssh -i codingxiaoma.pem ubuntu@ec2-18-142-178-139.ap-southeast-1.compute.amazonaws.com
```

就可以连接了。

![ssh](https://billy.taoxiaoxin.club/md/2023/06/64959aa6922ee403736e112d.png)

## 搭建VPN节点

### 1.登录服务器切换到root用户

首先通过终端工具登录到你的AWS 的 VPS,如果你是通过Mac ssh登录的话,首先需要执行如下命令,将权限改为只读:

```sh
chmod 400 codingxiaoma.pem
```

不然启动会报错,提示没有权限

![image-20230623220812079](https://billy.taoxiaoxin.club/md/2023/06/6495a74c922ee40dfe5ae7aa.png)

执行上面操作后就登录成功了

![image-20230623221013664](https://billy.taoxiaoxin.club/md/2023/06/6495a7c5922ee40e1035f17f.png)

然后继续设置root 账号设置密码,aws 服务器默认root 是没有密码的,

所以你要先设置密码,输入如下命令： 

```bash
sudo passwd root
```

然后会提示你输入new password。输入一个你要设置的root的密码，需要你再输入一遍进行验证。

![image-20230623221400160](https://billy.taoxiaoxin.club/md/2023/06/6495a8a8922ee40e5b8f761d.png)

接下来，切换到root身份，输入如下命令:

```bash
su root
```

### 2.更新系统安装软件

更新系统安装软件

```bash
apt update -y
apt install -y curl socat wget
```

安装X-UI 软件,执行如下命令:

~~~sh
# 下载 X-UI 脚本并且启动
bash <(curl -Ls https://raw.githubusercontent.com/vaxilu/x-ui/master/install.sh)
~~~

### 3.设置X-UI账户密码

运行上面的脚本之后,**会提示你出于安全考虑，安装/更新完成后需要强制修改端口与账户密码,输入Y**

按照提示输入账户,密码,端口号

![image-20230623232213776](https://billy.taoxiaoxin.club/md/2023/06/6495b8a6922ee417a88d8fb0.png)

### 4.设置安全组

点击实例找到你刚刚创建的VPS,然后点击安全,设置防火墙

![image-20230623232951124](https://billy.taoxiaoxin.club/md/2023/06/6495ba75922ee417d13f5a82.png)

点击进入安全组

![image-20230623233056604](https://billy.taoxiaoxin.club/md/2023/06/6495bab0922ee417d851da50.png)

点击编辑入站规则

![image-20230623233135334](https://billy.taoxiaoxin.club/md/2023/06/6495bad7922ee417e197387a.png)

设置自定义TCP,输入你刚设置的端口,设置为源为`0.0.0.0`

![image-20230623233351129](https://billy.taoxiaoxin.club/md/2023/06/6495bb92922ee417f303601d.png)

最后点击保存即可

![image-20230623233521065](https://billy.taoxiaoxin.club/md/2023/06/6495bbb9922ee417f78bb064.png)

### 5.访问面板地址

在服务器控制台找到你的公网IP,点击复制即可

![image-20230623233834413](https://billy.taoxiaoxin.club/md/2023/06/6495bc7a922ee418089ad6a4.png)

然后打开浏览器,通过**IP加端口号**访问你的节点,输入你**设置的登录的用户名和密码**

![image-20230623233947379](https://billy.taoxiaoxin.club/md/2023/06/6495bcc3922ee4181198ce9c.png)

然后设置系统状态,切换到最新版

![image-20230623234343779](https://billy.taoxiaoxin.club/md/2023/06/6495bdb0922ee41833bbda7b.png)

我这里最新的是1.8.3,切换到此版本就可以了,点击确定即可.

<img src="https://billy.taoxiaoxin.club/md/2023/06/6495bdee922ee4183c7cbf59.png" alt="image-20230623234445894" style="zoom: 50%;" />

### 6.配置域名与网站HTTPS证书

首先在安全组开启允许ICMP通过,添加如下两条记录:

![image-20230624131823049](https://billy.taoxiaoxin.club/md/2023/06/64967c9f922ee40d54b0964f.png)

然后添加域名解析,打开域名服务商控制台,添加域名解析

![image-20230624132739935](https://billy.taoxiaoxin.club/md/2023/06/64967ecc922ee40dc0e775f6.png)

域名解析IP在AWS控制台找到

![image-20230624132931952](https://billy.taoxiaoxin.club/md/2023/06/64967f3c922ee40dd6ecf50a.png)

检测域名解析是否生效,使用[ping工具（检测解析域名是否生效）](https://tool.dnspod.cn/),点击检测会出现如下结果:

![image-20230624135751312](https://billy.taoxiaoxin.club/md/2023/06/649685df922ee41386ddd90a.png)

接下来安装Acme软件,用于申请证书

```sh
curl https://get.acme.sh | sh
```

“你的邮箱”改成你的邮箱地址

```sh
~/.acme.sh/acme.sh --register-account -m 你的邮箱
```

“你的域名”改成你前面解析好的域名

```sh
~/.acme.sh/acme.sh --issue -d 你的域名 --standalone
```

下面是签发成功后的结果:

![image-20230624133632481](https://billy.taoxiaoxin.club/md/2023/06/649680e0922ee410961dcaf9.png)

打开你的面板设置,添加你刚刚申请的域名证书,填写对应路径即可,一般证书默认路径为:

+ 面板证书公钥文件路径:`/root/.acme.sh/你的域名_ecc/你的域名.cer`

+ #### 面板证书密钥文件路径:`/root/.acme.sh/你的域名_ecc/你的域名.key`

![image-20230624142227002](https://billy.taoxiaoxin.club/md/2023/06/64968ba3922ee416ff450f6f.png)

完成之后点击保存,点击重启面板.

![image-20230624002836583](https://billy.taoxiaoxin.club/md/2023/06/6495c834922ee41a5248a5bf.png)

之后访问后台管理只能通过HTTPS访问,访问路径为:`https://你的域名:端口号/`**通过HTTP 协议访问是打不开控制面板的**

### 7.添加节点

**目前推荐使用V2Ray的WS+TLS、Xray的xtls协议 这样安全性高、抗封锁能力强**

#### 使用 WS+TLS

点击入站规则,添加节点,随便填个备注,例如:`US-01`

<img src="https://billy.taoxiaoxin.club/md/2023/06/64968eaa922ee417961161bf.png" alt="image-20230624143522049" style="zoom:50%;" />

选择协议为`vmess`:

![image-20230624143649553](https://billy.taoxiaoxin.club/md/2023/06/64968f01922ee417acbbbd01.png)

端口选择443

![image-20230624143901344](https://billy.taoxiaoxin.club/md/2023/06/64968f85922ee417c9fc466e.png)

传输协议输入选择`ws`:

![image-20230624143948634](https://billy.taoxiaoxin.club/md/2023/06/64968fb4922ee417d3df5d26.png)

路径随便填写比如abc:

![image-20230624144112381](https://billy.taoxiaoxin.club/md/2023/06/64969008922ee417e79395f9.png)

设置请求头为你的域名:

<img src="https://billy.taoxiaoxin.club/md/2023/06/649690ae922ee4180c33889b.png" alt="image-20230624144358479" style="zoom:50%;" />

开启tls,并且配置域名证书

![image-20230624144804850](https://billy.taoxiaoxin.club/md/2023/06/649691a5922ee4184006f365.png)

最后点击保存

![image-20230624144940287](https://billy.taoxiaoxin.club/md/2023/06/64969204922ee41854c2349a.png)

添加完成后台如下图:

![image-20230624145009580](https://billy.taoxiaoxin.club/md/2023/06/64969221922ee4185ca71949.png)

#### 使用 Xray的xtls协议

备注名随便输入,比如:US-02,协议选择:`vless`

![image-20230624150513914](https://billy.taoxiaoxin.club/md/2023/06/649695aa922ee41a502882df.png)

开启xtls,并且配置域名与证书

![image-20230624150702865](https://billy.taoxiaoxin.club/md/2023/06/64969617922ee41a69d4c05e.png)

完整配置如下,最后点击保存;

![image-20230624150737485](https://billy.taoxiaoxin.club/md/2023/06/64969639922ee41a70264b5c.png)

最后在服务器安全组放开刚刚设置的端口号:

![image-20230624151133794](https://billy.taoxiaoxin.club/md/2023/06/64969725922ee41a97968f9d.png)

## 客户端配置

这里客户端可以使用任何支持SSR的客户端,可以点击 [V2Ray官网对于全平台客户端的总结和一览](https://www.v2ray.com/awesome/tools.html)

这里以clash为例,按照以下步骤下载并安装 Clash for Windows：

1. 打开 https://github.com/Fndroid/clash_for_windows_pkg/releases 链接。
2. 在页面上找到最新的版本，并向下滚动到 "Assets" 部分。
3. 下载与您的操作系统相对应的适用文件。如果您的系统是 Windows，选择 `.exe` 结尾的文件；如果是 macOS，选择 `.dmg` 结尾的文件。
4. 点击所选文件下载链接。
5. 下载完成后，找到下载的文件并双击打开。
6. 如果您使用的是 Windows 操作系统，会弹出安装向导。按照指示进行安装，选择您想要安装 Clash 的位置。
7. 如果您使用的是 macOS 操作系统，双击 `.dmg` 文件会自动挂载映像，并显示一个新的窗口。将 Clash 图标拖动到“应用程序”文件夹中，完成安装。
8. 打开 Clash for Windows，应用程序会启动。

默认是英文的,不支持中文,如果你需要汉化,请参考这篇文章:https://ednovas.github.io/2021/02/01/clashlanguage/

GitHub 汉化包最新版获取地址:https://github.com/BoyceLig/Clash_Chinese_Patch/releases

在控制台生成订阅链接,点击操作

![image-20230624151640680](https://billy.taoxiaoxin.club/md/2023/06/64969858922ee41ad41c8b0a.png)

然后点击复制,

![image-20230624151731712](https://billy.taoxiaoxin.club/md/2023/06/6496988b922ee41ae1fdd062.png)

打开地址:https://v2rayse.com/clash-convert/,将节点转为clash 节点

粘贴节点地址转换,类型选择`clash`,最后点击转换

![image-20230624235857850](https://billy.taoxiaoxin.club/md/2023/06/649712c2922ee432a28df6ce.png)

继续点击**订阅**

![image-20230625000157767](https://billy.taoxiaoxin.club/md/2023/06/64971375922ee432c7d6762f.png)

复制这个链接地址

![image-20230625000251063](https://billy.taoxiaoxin.club/md/2023/06/649713ab922ee432d3287ab8.png)

将链接导入到clash ,输入生成的链接地址,然后点击download

![image-20230624153839581](https://billy.taoxiaoxin.club/md/2023/06/64969d7f922ee41d497f3059.png)

然后右键点击编辑,



![image-20230624232817163](https://billy.taoxiaoxin.club/md/2023/06/64970b91922ee42bb5216c22.png)

双击选中你的节点配置文件点击编辑,复制刚生成的内容

![image-20230624153705514](https://billy.taoxiaoxin.club/md/2023/06/64969d21922ee41d34735ec1.png)

然后双击运行节点即可

![image-20230624231733377](https://billy.taoxiaoxin.club/md/2023/06/6497090d922ee42a265aa582.png)

## X-UI 面板管理

直接在命令行输入命令:`x-ui` ,输入对应的数字指令

![image-20230624235330152](https://billy.taoxiaoxin.club/md/2023/06/6497117a922ee43110786f3d.png)

如果忘记了密码等直接可以登录VPS输入对应指令修改
