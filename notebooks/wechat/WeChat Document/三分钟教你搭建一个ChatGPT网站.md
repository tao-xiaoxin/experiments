哈喽,大家晚上好!

OpenAI 已经公布了 ChatGPT 的API,因为某些原因,我们是直接无法使用API的,但是可以直接通过反代服务来变相访问 ChatGPT API.

今天教大家如何使用 Laf 调用ChatGPT 的 API,并且自己拥有一个稳定的ChatGPT网站,再也不用受到官网的各种限制了.

## 准备工作

1. 一个ChatGPT账号(需要调ChatGPT 的API key)
2. 一个Laf 账号(部署使用)
3. Node.js 环境(前端页面使用)

### Laf 介绍

Laf 是一个**完全开源**的一站式云开发平台，一个开箱即用的云函数，云数据库，对象存储等能力，让你可以像写博客一样写代码。

### 使用Laf构建ChatGPT应用

首先注册一个自己的账号并且登录

新注册网址: https://login.laf.dev/signup/laf

登录成功之后点新建一个应用

![image-20230313223830780](https://billy.taoxiaoxin.club/md/2023/03/640f3567922ee40c9eeb1d5d.png)

然后新建一个应用名称为ChatGPT.

![image-20230313224037218](https://billy.taoxiaoxin.club/md/2023/03/640f35e5922ee40ca95d2ffd.png)

点击进入开发

![image-20230313224335052](https://billy.taoxiaoxin.club/md/2023/03/640f3697922ee40cbd16c673.png)

点NPM 依赖面板中点击右上角的**加号**

![image-20230313224510593](https://billy.taoxiaoxin.club/md/2023/03/640f36f6922ee40cc611c9f8.png)

然后输入 chatgpt 并回车进行搜索，选择第一个搜索结果，**保存并重启：**

![image-20230313224716304](https://billy.taoxiaoxin.club/md/2023/03/640f3774922ee40cce9adc89.png)

登录你的ChatGPT账号:

网址: https://chat.openai.com/auth/login

然后去ChatGPT官网生成一个API Key

网址:https://platform.openai.com/account/api-keys

点击页面新增一个key，并且复制保存到记事本。

![image-20230313231209780](https://billy.taoxiaoxin.club/md/2023/03/640f3d4a922ee40d4fa5420e.png)

然后新建一个云函数名字叫 **send**，

![image-20230313225451465](https://billy.taoxiaoxin.club/md/2023/03/640f393b922ee40cf4b7af46.png)

新建完成后写入以下内容：

```javascript
import cloud from '@lafjs/cloud'

export async function main(ctx: FunctionContext) {
  const { ChatGPTAPI } = await import('chatgpt')
  const data = ctx.body

  // 这里需要把 api 对象放入 cloud.shared 不然无法追踪上下文
  let api = cloud.shared.get('api')
  if (!api) {
    api = new ChatGPTAPI({ apiKey: '你的API key' })
    cloud.shared.set('api', api)
  }

  let res
  // 这里前端如果传过来 parentMessageId 则代表需要追踪上下文
  if (!data.parentMessageId) {
    res = await api.sendMessage(data.message)
  } else {
    res = await api.sendMessage(data.message, { parentMessageId: data.parentMessageId })
  }
  return res
}
```

将代码中的API key 替换为你的

![image-20230313232523414](https://billy.taoxiaoxin.club/md/2023/03/640f4063922ee40d7e07dc57.png)

继续点右上角**发布**按钮

![image-20230313232704069](https://billy.taoxiaoxin.club/md/2023/03/640f40c8922ee40d8ca9718e.png)

然后打开如下地址,下载前端页面:

地址 :https://github.com/zuoFeng59556/chatGPT

![image-20230313234204803](https://billy.taoxiaoxin.club/md/2023/03/640f444c922ee40ecaf2810f.png)

然后解压文件,继续编辑view--->index.vue文件

![image-20230313235600397](https://billy.taoxiaoxin.club/md/2023/03/640f4790922ee415e11597cb.png)

打开地址: https://laf.dev/ ,然后复制你的云函数ID

![image-20230313235718906](https://billy.taoxiaoxin.club/md/2023/03/640f47df922ee415e80471d0.png)

将下面的index.vue文件里面的ID替换为你的ID

![image-20230313235459089](https://billy.taoxiaoxin.club/md/2023/03/640f4753922ee415d3f4ed31.png)

然后运行命令如下:

```bash
npm i
npm run dev
```

**注意**:这里需要node环境的支持, 没有node 环境的可以去bing.com**搜下 node 安装教程**

执行上面的命令后,打开访问地址:

```bash
http://127.0.0.1:5173/
```

![image-20230314000431525](https://billy.taoxiaoxin.club/md/2023/03/640f498f922ee417293563f0.png)

对话框中测试是否可以正常使用

![image-20230314000548347](https://billy.taoxiaoxin.club/md/2023/03/640f49dc922ee41731ff862d.png)

最后我们我们把页面打包一下并且部署上去,执行如下命令:

```bash
npm run build
```

然后继续打开你的 Laf，点击存储界面 --> 点击上方加号 --> 创建一个权限为 readonly 的存储桶（名字随意）。

![image-20230314001429325](https://billy.taoxiaoxin.club/md/2023/03/640f4be5922ee41dd63f0651.png)

我这里创建了一个ChatGPT-Web 的桶,**将权限一定要设置为公共读**

![image-20230314001729896](https://billy.taoxiaoxin.club/md/2023/03/640f4c9a922ee41ded1c495b.png)

继续上传刚刚打包生成的文件夹`ChatGPT-main/dist ,将文件和文件夹挨个上传.

![image-20230314002631848](https://billy.taoxiaoxin.club/md/2023/03/640f4eb8922ee41f310f8230.png)

![image-20230314002758563](https://billy.taoxiaoxin.club/md/2023/03/640f4f0e922ee41f3990ec82.png)

上传完毕之后，发现右上角有一个 “开启网站托管”，点一下它！

![图片](https://billy.taoxiaoxin.club/md/2023/03/640f4dd9922ee41f210b123d.jpeg)

然后打开右上角域名就好了,

![image-20230314003152154](https://billy.taoxiaoxin.club/md/2023/03/640f4ff8922ee41f5bc8abff.png)

能访问后成功!

![image-20230314003257656](https://billy.taoxiaoxin.club/md/2023/03/640f5039922ee41f625d567c.png)

### 总结

+ 前端代码地址**https://github.com/zuoFeng59556/chatGPT**
+ 在线体验网站：**https://jyf6wk-chat-gpt.site.laf.dev/**

