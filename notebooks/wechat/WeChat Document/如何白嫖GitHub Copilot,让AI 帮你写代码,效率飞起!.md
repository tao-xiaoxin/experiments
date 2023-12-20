哈喽,大家好呀!

最近呢,我在用**GitHub Copilot** 写代码,感觉真的挺好用的

所以今天给大家推荐下!

### **GitHub Copilot** 介绍

GitHub Copilot 是由 GitHub 和 OpenAI 共同开发的人工智能代码辅助工具，可以自动地生成高质量代码片段、上下文信息等。通过自然语言处理和机器学习技术，能够**通过分析程序员编写的代码、注释和上下文信息，自动生成代码**，减轻程序员的工作量,节省开发者的时间和精力。

GitHub Copilot 支持的代码编辑其如下:

+ **Visual Studio Code**
+ **Visual Studio**
+ **Pycharm** ...

支持目前主流的多种编程语言:

+ Python
+ JavaScript
+ Go
+ TypeScript ...

### 前提条件

首先,你要有一个**GitHub账号**,没有的话先去百度注册一下吧!

有一个稳定的网络环境,可以访问到GitHub,就是需要**准备个科学的上网 工具 !

### 注册 Copilot

打开浏览器,登录Github账号后

如下链接申请注册Copilot：

```bash
 https://github.com/github-copilot/free_signup 
```

Copilot 需要付费使用，**普通用户有30天试用期**，但针对开源作者、学生、老师免费开放使用，比如我就是用淘宝买的学生认证,

也就25块钱一年,也不是很贵,里面包含office 正版激活,IDEA 全家桶正版激活等等,总体来说挺划算的.

好了,回到正文,如果满足免费使用条件,界面应该是如下:

![image-20230417212707405](https://billy.taoxiaoxin.club/md/2023/04/643d492b922ee48c5c0edce8.png)

如果你已经用过了,界面应该是如下:

![image-20230417212550806](https://billy.taoxiaoxin.club/md/2023/04/643d4916922ee48c588c0734.png)

### 安装插件和配置

#### Pycharm

打开Pycharm -> Preferences -> Plugins

搜索插件“**GitHub Copilot**” 直接安装即可

![image-20230417213325758](https://billy.taoxiaoxin.club/md/2023/04/643d4aa5922ee48c7d51eb2c.png)

安装好以后,点击重启Pycharm

![image-20230417213746435](https://billy.taoxiaoxin.club/md/2023/04/643d4baa922ee48c95bfa6cc.png)

右下角会多一个copilot 的logo，需要你登录GitHub,登录即可

![image-20230417214022842](https://billy.taoxiaoxin.club/md/2023/04/643d4c47922ee48db80f273d.png)

访问GitHub 注意你的网络环境奥,可能需要开科学的工具.

### VS Code 

首先需要在 VS Code 编辑器中安装相应的插件。在 VS Code 中，点击左侧的“扩展”选项卡，搜索“GitHub Copilot”，并安装该插件,选择第一个即可

![image-20230417214550770](https://billy.taoxiaoxin.club/md/2023/04/643d4d8e922ee48dfb58dae2.png)

安装完成后,重启VS Code，需要登录 GitHub 帐户进行身份验证。

### 使用 Copilot 进行代码提示

安装并配置好后，就可以使用 GitHub Copilot 进行代码提示了。在编辑器中输入一些代码时，Copilot 会根据上下文和语法规则，自动提示一些可能的代码片段。如果需要使用 Copilot 提示的代码，只需要按下“Tab”键即可将其插入到当前光标位置。

比如，**在Pycharm中使用Copilot**

登录成功后就可以直接在编辑器里面使用。比如写一个邮箱校验函数，检查邮箱格式是否合法，直接点击右侧的Copilot机器人，就会弹出几个备选方案，双击 Accept solution 就可以把代码写入到文件中。

![图片](https://billy.taoxiaoxin.club/md/2023/04/643d4f2f922ee48ef6306fd8.png)

再比如果想写一个判断是否为整数的函数，只要把函数名写上（甚至只要写一半），Copilot就会自动提示，此时只要按Tab键，就可以补全代码，注意灰色部分是它给我的建议。

![图片](https://billy.taoxiaoxin.club/md/2023/04/643d4f2f922ee48ef776289a.png)

### GitHub Copilot 键盘快捷键

- 接受内联代码建议  `Tab`
- 关闭内联代码建议  `Esc`
- 显示下一个建议   `Alt + ] `
- 显示上一个建议   `Alt + [ `
- 触发建议   `Alt + \ `
- 在右侧窗口中显示十个建议   `Ctrl + Enter`

### Copilot 的优点与缺点

#### Copilot 的优点

GitHub Copilot 具有许多优点，使其成为开发者喜欢使用的工具之一。以下是其中的一些优点：

- 生成代码速度快：Copilot 使用先进的自然语言处理技术和机器学习算法，可以在几乎瞬间生成高质量的代码片段，节省开发者的时间和精力。
- 提高代码质量：由于 Copilot 生成的代码是基于机器学习模型的，它可以避免一些常见的错误，从而提高代码质量。
- 适应多种编程语言：Copilot 可以适应多种编程语言和框架，包括 Python、JavaScript、Ruby 等，为开发者提供了更多的选择。
- 可定制性强：Copilot 允许开发者自定义其提示行为，例如指定要使用的语言和框架、添加自定义代码片段和快捷键等。
- 不断学习进步：Copilot 是基于机器学习技术的，可以不断学习进步，提高其生成代码的准确性和质量。

#### Copilot 的缺点:

虽然 GitHub Copilot 是一个非常有用的工具，但它仍然存在一些限制。以下是其中的一些限制：

+ **对国内用户不是很友好,因为GitHub有时候需要通过代理才能访问**

- 有时会生成错误的代码：Copilot 生成的代码并不总是完全正确，有时需要开发者自己对其进行修改和调整。
- 安全性问题：由于 Copilot 是基于机器学习技术的，因此可能会存在一些安全性问题，例如泄漏敏感信息等。
- 不支持所有编程语言和框架：尽管 Copilot 可以适应多种编程语言和框架，但仍有一些不支持的编程语言和框架。

### 总结

GitHub Copilot 是一款非常有用的代码提示工具，可以帮助开发者更快速、更高效地编写代码。它具有许多优点，例如生成代码速度快、提高代码质量、适应多种编程语言和框架等，但仍存在一些限制，**例如有时会生成错误的代码**。因此，在使用 Copilot 时，**开发者需要根据实际情况权衡其优缺点**，以便更好地使用这个工具。

GitHub Copilot 官网文档：

```bash
https://docs.github.com/zh/copilot
```

当然，GitHub Copilot 不仅仅能够帮你生成代码片段，也可帮你生成测试等等

对GitHub Copilot 感兴趣的具体大家可以去看看极客时间最近新出的公开课，免费不要钱的那种。

具体课程大纲如下：

![image-20230417221226388](https://billy.taoxiaoxin.club/md/2023/04/643d53ca922ee49187580767.png)



首先送大家14 天会员，免费畅看极客时间会员专栏。

领取地址如下（**数量有限，先到先的**）

```go
http://gk.link/a/121UM
```

GitHub Copilot 公开课地址：

```go
http://gk.link/a/11ZDr
```



好了，今天的分享就到这里，**觉得有用请点个赞！**