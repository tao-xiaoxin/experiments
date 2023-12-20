哈喽,大家好!

今天教大家如何利用ChatGPT API搭建微信机器人,让它成为的小助理,帮你提高效率。

## 准备

1. 你需要有一台国外服务器
2. **个人微信账号必须实名认证并且绑定了银行卡**，不然会被微信风控

没有这些的话可以去我的微信群里面体验一下,下面的这些不用看了,直接滑到末尾加入群聊。

## 快速开始

### 1. OpenAI账号注册

注册可以看之前写的教程：

[教程 | 手把手教你注册一个ChatGPT账号](https://mp.weixin.qq.com/s/2MZgf-nYavQtouiSfnxvxg)

### 2.创建 API Key 

创建完账号则前往API管理页面地址：https://beta.openai.com/account/api-keys

然后点击 “API Keys”，就会进入一个管理 API Keys 的页面,继续点击你点击下面的 “+Create new secret key” 可以创建一个新的 API Key。

![image-20230510134651453](https://billy.taoxiaoxin.club/md/2023/05/646f61eb922ee43372173bb6.png)

输入key的名称,随便起个名字,然后点击create key

![image-20230510135206301](https://billy.taoxiaoxin.club/md/2023/05/646f61eb922ee433735a72f9.png)

 创建一个并保存下来，后面需要在项目中配置这个key。

![image-20230510135641871](https://billy.taoxiaoxin.club/md/2023/05/646f6201922ee433771f5572.png)

### 3.安装Python 和 Git 

这个自己百度搜吧,我就不写了

建议Python版本推荐3.8版本，3.10及以上版本。

### 3.运行环境

代码支持 Linux、MacOS、Windows 系统（可在Linux服务器上长期运行)，我这里使用的是同时需安装 `Python`。

**(1) 克隆项目代码：**

```bash
git clone https://github.com/zhayujie/chatgpt-on-wechat
cd chatgpt-on-wechat/
```

**(2) 安装核心依赖：**

> 能够使用`itchat`创建机器人，并具有文字交流功能所需的最小依赖集合。

```bash
pip3 install -r requirements.txt
```

**(3) 安装拓展依赖 ：**

```bash
pip3 install -r requirements-optional.txt
```

## 配置

配置文件的模板在根目录的`config-template.json`中，需复制该模板创建最终生效的 `config.json` 文件：

```bash
  cp config-template.json config.json
```

然后在`config.json`中填入配置，以下是对默认配置的说明，可根据需要进行自定义修改（请去掉注释）：

![carbon](https://billy.taoxiaoxin.club/md/2023/05/646f8744922ee4368c8f3619.png)

**配置说明：**

**1.个人聊天**

- 个人聊天中，需要以 "bot"或"@bot" 为开头的内容触发机器人，对应配置项 `single_chat_prefix` (如果不需要以前缀触发可以填写 `""single_chat_prefix": ["bot", "@bot"] `)
- 机器人回复的内容会以 "[bot] " 作为前缀， 以区分真人，对应的配置项为 `single_chat_reply_prefix` (如果不需要前缀可以填写 `"single_chat_reply_prefix": "[bot] "`)

**2.群组聊天**

- 群组聊天中，群名称需配置在 `group_name_white_list ` 中才能开启群聊自动回复。如果想对所有群聊生效，可以直接填写 `"group_name_white_list": ["ALL_GROUP"]`
- 默认只要被人 @ 就会触发机器人自动回复；另外群聊天中只要检测到以 "@bot" 开头的内容，同样会自动回复（方便自己触发），这对应配置项 `group_chat_prefix`
- 可选配置: `group_name_keyword_white_list`配置项支持模糊匹配群名称，`group_chat_keyword`配置项则支持模糊匹配群消息内容，用法与上述两个配置项相同。
- `group_chat_in_one_session`：使群聊共享一个会话上下文，配置 `["ALL_GROUP"]` 则作用于所有群聊

**3.语音识别**

- 添加 `"speech_recognition": true` 将开启语音识别，默认使用openai的whisper模型识别为文字，同时以文字回复，该参数仅支持私聊 (注意由于语音消息无法匹配前缀，一旦开启将对所有语音自动回复，支持语音触发画图)；
- 添加 `"group_speech_recognition": true` 将开启群组语音识别，默认使用openai的whisper模型识别为文字，同时以文字回复，参数仅支持群聊 (会匹配group_chat_prefix和group_chat_keyword, 支持语音触发画图)；
- 添加 `"voice_reply_voice": true` 将开启语音回复语音（同时作用于私聊和群聊），但是需要配置对应语音合成平台的key，由于itchat协议的限制，只能发送语音mp3文件，若使用wechaty则回复的是微信语音。

**4.其他配置**

- `model`: 模型名称，目前支持 `gpt-3.5-turbo`, `text-davinci-003`, `gpt-4`, `gpt-4-32k` (其中gpt-4 api暂未开放)

- `temperature`,`frequency_penalty`,`presence_penalty`: Chat API接口参数，详情参考OpenAI官方文档。

  OpenAI官方文档 : https://platform.openai.com/docs/api-reference/chat

- `proxy`：由于目前 `openai` 接口国内无法访问，需配置代理客户端的地址。

- 对于图像生成，在满足个人或群组触发条件外，还需要额外的关键词前缀来触发，对应配置 `image_create_prefix `

- 关于OpenAI对话及图片接口的参数配置（内容自由度、回复字数限制、图片大小等），可以参考对话接口和 图像接口 文档直接在 `bot/openai/open_ai_bot.py` 中进行调整。

  对话和图像接口文档地址 : https://beta.openai.com/docs/api-reference/completions

- `conversation_max_tokens`：表示能够记忆的上下文最大字数（一问一答为一组对话，如果累积的对话字数超出限制，就会优先移除最早的一组对话）

- `rate_limit_chatgpt`，`rate_limit_dalle`：每分钟最高问答速率、画图速率，超速后排队按序处理。

- `clear_memory_commands`: 对话内指令，主动清空前文记忆，字符串数组可自定义指令别名。

- `hot_reload`: 程序退出后，暂存微信扫码状态，默认关闭。

- `character_desc` 配置中保存着你对机器人说的一段话，他会记住这段话并作为他的设定，你可以为他定制任何人格 (关于会话上下文的更多内容参考该 [issue](https://github.com/zhayujie/chatgpt-on-wechat/issues/43))

- `subscribe_msg`：订阅消息，公众号和企业微信channel中请填写，当被订阅时会自动回复， 可使用特殊占位符。目前支持的占位符有{trigger_prefix}，在程序中它会自动替换成bot的触发词。

**所有可选的配置项均在该文件 `chatgpt-on-wechat/config.py` 中列出。**

## 运行

### 1.本地运行

如果是开发机 **本地运行**，直接在项目根目录下执行：

```bash
python3 app.py
```

终端输出二维码后，使用微信进行扫码，当输出 "Start auto replying" 时表示自动回复程序已经成功运行了（注意：用于登录的微信需要在支付处已完成实名认证）。扫码登录后你的账号就成为机器人了，可以在微信手机端通过配置的关键词触发自动回复 (任意好友发送消息给你，或是自己发消息给好友)。

### 2.服务器部署

使用nohup命令在后台运行程序：

```bash
touch nohup.out                                   # 首次运行需要新建日志文件  
nohup python3 app.py & tail -f nohup.out          # 在后台运行程序并通过日志输出二维码
```

扫码登录后程序即可运行于服务器后台，此时可通过 `ctrl+c` 关闭日志，不会影响后台程序的运行。

使用 `ps -ef | grep app.py | grep -v grep` 命令可查看运行于后台的进程，如果想要重新启动程序可以先 `kill` 掉对应的进程。日志关闭后如果想要再次打开只需输入 `tail -f nohup.out`。此外，`scripts` 目录下有一键运行、关闭程序的脚本供使用。

> **多账号支持：** 将项目复制多份，分别启动程序，用不同账号扫码登录即可实现同时运行。

> **特殊指令：** 用户向机器人发送 **#reset** 即可清空该用户的上下文记忆。



## 总结

不会代码的，觉得麻烦的可以去我的微信群里面体验一下，它可以帮你做任何事情。

二维码如下，感兴趣的快快扫码加入吧！

<img src="https://billy.taoxiaoxin.club/md/2023/05/646f69bb922ee4341dafff39.png" alt="image-20230525215923107" style="zoom:50%;" />

**好了,今天的教程就到这里,希望能够帮助到你,有用请点个赞!**

