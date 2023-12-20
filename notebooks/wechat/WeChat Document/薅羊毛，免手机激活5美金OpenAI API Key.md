# 薅羊毛，免手机激活5美金OpenAI API Key（全网首发，建议收藏）

现在chatgpt注册都是免手机号的，还没有注册的可以去看看之前写的教程。

[ChatGPT 最新注册教程，无需验证手机号，100%成功！](https://mp.weixin.qq.com/s?__biz=Mzg3ODA5ODY3MQ==&mid=2247499106&idx=1&sn=2b779e153fd509638ce0ab6274572ee9&chksm=cf1a5ce2f86dd5f45126ba1397e1dc750694b4c18c478366be865d0211ed3fa3d34ec1582227&token=1101936020&lang=zh_CN#rd)

虽然现在注册不需要手机号了，但是创建API Key需要验证手机号。

![](https://billy.taoxiaoxin.club/md/2023/12/6571686ae5269c945f15ccf0.png)

下面就教大家如何跳过手机号验证获取API Key。

首先打开如下地址：

```go
https://platform.openai.com/apps
```

选择API

![](https://billy.taoxiaoxin.club/md/2023/12/657168c7f9acc41fda60a1b3.png)

找到左侧的API Key。

![](https://billy.taoxiaoxin.club/md/2023/12/657168f50e090e187c1724b9.png)

打开这个页面会提示需要验证手机号。

![](https://billy.taoxiaoxin.club/md/2023/12/6571697eb15a318bcc577d8d.png)

接下来我们打开谷歌浏览器，右键检查

![](https://billy.taoxiaoxin.club/md/2023/12/65716a37b07824d3890f53d6.png)

打开浏览器控制台，**勾选保留日志，停用缓存，Fetch/XHR** 如图：

![](https://billy.taoxiaoxin.club/md/2023/12/65716ae4fe59856074f8e362.png)

接着，刷新网页。

![](https://billy.taoxiaoxin.club/md/2023/12/65716b2323a5386df8fc379f.png)

在浏览器控制台找到login，确保响应状态码是200。

![](https://billy.taoxiaoxin.club/md/2023/12/65716c765dce530fcb9d73f5.png)

接着打开响应，找到sensitive_id并复制它。

![](https://billy.taoxiaoxin.club/md/2023/12/65716db53346db1ebe27303c.png)

打开如下地址，测试API Key 是否可用：

```go
https://huggingface.co/spaces/JohnSmith9982/ChuanhuChatGPT
```

![](https://billy.taoxiaoxin.club/md/2023/12/65716ead6f140953945ad641.png)

在API Key 输入框粘贴

![](https://billy.taoxiaoxin.club/md/2023/12/65716f1a335d98d79f357e2d.png)

如果你在API Key 随便输入一个key ，会得到如下提示：

![](https://billy.taoxiaoxin.club/md/2023/12/65716fac7abd8520377357ad.png)

但是如果你使用的是刚刚粘贴sensitive_id，会话是正常的，并且会显示你消耗了多少token。

![](https://billy.taoxiaoxin.club/md/2023/12/657170bc84af26647d886dcb.png)

接着在官网控制台也可以看到你刚刚消耗了多少美元。

```bash
https://platform.openai.com/usage
```

![](https://billy.taoxiaoxin.club/md/2023/12/65717320ea06308f32320fbd.png)

通过这种方式申请的API Key 可以用5天左右，过期了你需要按照上面的方式重新获取一遍。如果用邮箱激活，每一个邮箱都会送这个5美元，用完了，你就可以换一个邮箱，不停的薅羊毛，无穷无尽。当然你也可以通过某些手段，**注册100个账号，获得100个API Key，然后轮流使用，那这样是不是就薅到了500美金的余额了**。

**建议你收藏本文，以免API Key 过期了找不到教程。**

**今天的分享会就到这里，希望对你有用，有用记得点赞+在看支持一下奥！**

**最后，安利一波我的薅羊毛群，感兴趣的可以加入一下。**

## 神车群

+ **每天高频率更新：陶宝，京东等平台的神车** 
+ **主要是大牌、拆单、平行叠加 bug**

![](https://billy.taoxiaoxin.club/md/2023/12/657176397c9c0c2edfaedaa3.jpg)

## 纸巾

+ **每天高频率更新：拼夕夕的纸巾。**
+ **一提纸巾价格也就1～7元左右。**

![](https://billy.taoxiaoxin.club/md/2023/12/65717646c4ec0d3d93697ce3.png)

## 京东酒水

+ **每天高频率更新：陶宝，京东的酒水神车**

+ **当然也可以有撸茅台**

![](https://billy.taoxiaoxin.club/md/2023/12/65717655ee3f769e41dac5e1.png)

## 拼西西零食水果

+ **每天高频率更新：拼西西水果零食纸巾券**

![](https://billy.taoxiaoxin.club/md/2023/12/65717660c101676ca21ba1ee.png)
