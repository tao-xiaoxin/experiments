哈喽，大家好呀！我是小欣！

**GitHub Copilot** 想必大家都是熟悉的，今天给大家详细讲解一下安装激活教程，无需账号等登陆操作。

上次分享了，出现各种问题，所以，这次我体验了一段时间再给大家发下。

## 特别声明:

+ **本教程仅供个人学习和研究使用**

## GitHub Copilot 介绍

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

## Pycharm 插件安装

打开 Pycharm -> Preferences -> Plugins

搜索插件“**GitHub Copilot**” 直接安装即可

![image-20230417213325758](https://billy.taoxiaoxin.club/md/2023/11/6551b67e5d8879d48ec4c970.png)

安装好以后,点击重启Pycharm

![image-20230417213746435](https://billy.taoxiaoxin.club/md/2023/11/6551b67ed884ac9e69ab76ca.png)



## VS Code 插件安装

首先需要在 VS Code 编辑器中安装相应的插件。在 VS Code 中，点击左侧的“扩展”选项卡，搜索“GitHub Copilot”，并安装该插件,选择第一个即可。

![image-20230417214550770](https://billy.taoxiaoxin.club/md/2023/11/6551b67e2e21b8e120fdf4ea.png)

安装完成后,重启VS Code，需要登录 GitHub 帐户进行身份验证。

## 激活插件安装

### windows 激活

`windows系统`双击执行`cocopilot.bat`

![image-20231113140854230](https://billy.taoxiaoxin.club/md/2023/11/6551bd76537f5e234ba0cf05.png)

### macOS/linux 激活

**macOS/linux** 执行`cocopilot.sh`，终端执行如下命令：

```bash
chmod +x cocopilot.sh vscode-remote.sh vscode.sh
./cocopilot.sh 
```

看到`done. please restart your ide.`表示成功。

![](https://billy.taoxiaoxin.club/md/2023/11/6551bea23722ecabcf3ef552.png)

### 温馨提示

1. 对于`VSCode`，步骤和上面相同，执行对应`vscode.sh`/`vscode.bat`
2. 如果是使用vscode远程连接Ubuntu服务器且副驾驶拓展安装在了远程服务器上，需要执行 `vscode-remote.sh`，**无需执行**`cocopilot.sh`/`cocopilot.bat`
3. `VSCode`中插件更新后需要重新执行脚本，`JetBrains`则不需要。

## 激活插件常见问题

如果你之前用过GitHub Copilot ，需要先退出你的账号，因为激活的原理是用的别人的共享token，就是你登录的是别人的GitHub Copilot！

### 对于 jetbrains 系列的问题

对于jetbrains 系列全家桶，退出你的账号登录就是点Copilot 那个小图标，选择`Logout from GitHub`，然后关闭jetbrains 系列的产品，重新运行`cocopilot.sh`/`cocopilot.bat`

<img src="https://billy.taoxiaoxin.club/md/2023/11/655d99958111b70f7185616b.png" alt="image-20231122140300495" style="zoom:50%;" />

当你执行了`cocopilot.sh`/`cocopilot.bat`， Copilot 那个小图标显示的是登录状态，没有的话加没有成功，需要自己研究修改代码。

<img src="https://billy.taoxiaoxin.club/md/2023/11/655d9ad8eec659fcb9382c9f.png" alt="image-20231122140823873" style="zoom:50%;" />

还有为什么显示是登录状态，代码还是没有提示，这个就是网络问题了，也是需要自己去研究的。

![image-20231122141039362](https://billy.taoxiaoxin.club/md/2023/11/655d9b5f78b1f62523947e6e.png)

### 对于 VSCode ：

对于退出你的账号登录

![image-20231122143814151](https://billy.taoxiaoxin.club/md/2023/11/655da1d629080310ae8ec784.png)

![image-20231122144028485](https://billy.taoxiaoxin.club/md/2023/11/655da25cf3f92c4b9a6a3c83.png)

![image-20231122144530150](https://billy.taoxiaoxin.club/md/2023/11/655da395ab745cd19f736b13.png)

退出后关闭VScode，重新执行`vscode.sh/vscode.bat`

如果依然提示如下，你还是没有退出你的账号：

![image-20231122142440485](https://billy.taoxiaoxin.club/md/2023/11/655d9ea8fae8f540140e0760.png)

![image-20231122142800962](https://billy.taoxiaoxin.club/md/2023/11/655d9f715ff33c89060a1f85.png)

那就按照上面的退出即可。

如果VSCode 没有成功的话，请参考备用方案文档。



### 说明：

**还有如果你想要其他产品中使用GitHub Copilot ，请参考`cocopilot.sh`/`cocopilot.bat` 这个两个中的源代码，自行研究激活，我只测试了jetbrains 系列产品和Vscode，按照以上操作是没有任何问题的。**

如果不会的话，请使用备用方案，备用方案支持的软件比较多，如图：

![image-20231122150137692](https://billy.taoxiaoxin.club/md/2023/11/655da7518e5c5707e257772a.png)
**综上，出现问题就是自行研究，修改对应脚本的源代码，或者使用备用方案。**

## 使用 Copilot 进行代码提示

安装并配置好后，就可以使用 GitHub Copilot 进行代码提示了。在编辑器中输入一些代码时，Copilot 会根据上下文和语法规则，自动提示一些可能的代码片段。如果需要使用 Copilot 提示的代码，只需要按下“Tab”键即可将其插入到当前光标位置。

比如，**在Pycharm中使用Copilot**，直接在编辑器里面使用。比如写一个邮箱校验函数，检查邮箱格式是否合法，直接点击右侧的Copilot机器人，就会弹出几个备选方案，双击 Accept solution 就可以把代码写入到文件中。

![图片](https://billy.taoxiaoxin.club/md/2023/11/6551b67ea8d20fdcd1077015.png)

再比如果想写一个判断是否为整数的函数，只要把函数名写上（甚至只要写一半），Copilot就会自动提示，此时只要按Tab键，就可以补全代码，注意灰色部分是它给我的建议。

![图片](https://billy.taoxiaoxin.club/md/2023/11/6551b67e217c76f71edaa619.png)

## GitHub Copilot 键盘快捷键

- 接受内联代码建议  `Tab`
- 关闭内联代码建议  `Esc`
- 显示下一个建议   `Alt + ] `
- 显示上一个建议   `Alt + [ `
- 触发建议   `Alt + \ `
- 在右侧窗口中显示十个建议   `Ctrl + Enter`

## Copilot 的优点与缺点

### Copilot 的优点

GitHub Copilot 具有许多优点，使其成为开发者喜欢使用的工具之一。以下是其中的一些优点：

- 生成代码速度快：Copilot 使用先进的自然语言处理技术和机器学习算法，可以在几乎瞬间生成高质量的代码片段，节省开发者的时间和精力。
- 提高代码质量：由于 Copilot 生成的代码是基于机器学习模型的，它可以避免一些常见的错误，从而提高代码质量。
- 适应多种编程语言：Copilot 可以适应多种编程语言和框架，包括 Python、JavaScript、Ruby 等，为开发者提供了更多的选择。
- 可定制性强：Copilot 允许开发者自定义其提示行为，例如指定要使用的语言和框架、添加自定义代码片段和快捷键等。
- 不断学习进步：Copilot 是基于机器学习技术的，可以不断学习进步，提高其生成代码的准确性和质量。

### Copilot 的缺点:

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

## 激活插件下载地址

激活插件下载地址如下：

```bash
https://txx.lanzoub.com/iy7fH1flxxod
```

好了，今天的分享就到这里，**觉得有用请点个赞！**

最后给大家安利一波我的薅羊毛群，扫码添加下方微信，自动加入群聊！

<img src="https://billy.taoxiaoxin.club/md/2023/11/655c5fb9273aeb330aa00266.jpeg" alt="图片" style="zoom:67%;" />