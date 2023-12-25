# Mac 安装软件常见报错与解决方案（汇总)

## 提示：无法打开“XXX”，因为它来自身份不明的开发者

双击 .dmg 安装软件出现：打不开“XXXX”，因为它来自身份不明的开发者。

![](https://billy.taoxiaoxin.club/md/2023/12/6586973b2418f62f425b66a6.jpeg)

### 解决方案：打开允许“允许任何来源”

在 Finder 菜单栏选择 【前往】 – 【实用工具 】，找到【终端】程序，双击打开，在终端窗口中输入：

```bash
sudo spctl --master-disable
```

输入代码后，按【return 回车键】，这时候会提示输入密码：直接输入自己的电脑密码，然后按【return 回车键】即可， （提示：在输入密码的时候，终端不会有任何显示。密码为开机密码，不要错误）

![](https://billy.taoxiaoxin.club/md/2023/12/658697b09635f642f72512db.png)

关闭【终端】，重新打开 【系统偏好设置】 – 【安全性与隐私】 – 【通用】 中就会出现且选中 【允许任何来源】

![在这里插入图片描述](https://billy.taoxiaoxin.club/md/2023/12/658697b0a1b4f82bc6e4dd81.png)

## 提示：文件损坏

“XXXX”已损坏，打不开。您应该将它移到废纸篓。

![](https://billy.taoxiaoxin.club/md/2023/12/6586987a3eb4443ee8dad4e9.png)

### 解决方案一：解除隔离，绕过签名

**启动终端实用程序，使用命令如下**：

```go
sudo xattr -r -d com.apple.quarantine [键盘空格] （损坏的程序路径，通过应用程序哪里拖进来）
```

![](https://billy.taoxiaoxin.club/md/2023/12/65869c665f949f13f595ec1b.png)

记得输入命令是`sudo xattr -r -d com.apple.quarantine`

输入一个空格，最后将损坏的程序拖到终端窗口中。团队将添加它。最后回到应用程序选中`损坏的软件`右键点击打开。

![](https://billy.taoxiaoxin.club/md/2023/12/65869b4cde285abe92c513c7.png)

### 解决方案二：禁用 Gatekeeper

首先启动终端实用程序并运行命令

1. 对于 `macOS 10.12 – 10.15.7`，执行如下命令：

```
sudo spctl --master-disable
```

2. 对于 `macOS 11` 及更高版本，执行如下命令：

```
sudo spctl --global-disable
```

**无需将任何内容拖到终端中。要启用 Gatekeeper，请将“disable”替换为“enable”。**

### 解决方案三：应用签名方法

1. **先安装Command Line Tools 工具，打开终端工具输入如下命令：**

   ```bash
   xcode-select --install
   ```

2. **弹出安装窗口后选择继续安装，安装过程需要几分钟，请耐心等待。**

3. **打开终端工具输入并执行如下命令对应用签名：**

```bash
sudo codesign --force --deep --sign -[键盘空格]（损坏的程序路径，通过应用程序哪里拖进来）
```

4. 打开「访达（Finder）->应用程序」，找到应用将其拖进终端命令 `–` 的后面，然后按下回车键，输入macOS的密码然后按回车(输入过程中密码是不显示的，输入完密码直接按回车键即可！

   ![](https://billy.taoxiaoxin.club/md/2023/12/65869f2d6741ee05cdefe907.png)

   ![](https://billy.taoxiaoxin.club/md/2023/12/65869f3646b4e0d53ad797dd.png)

5. **出现 「replacing existing signature」 提示即成功！**

![](https://billy.taoxiaoxin.club/md/2023/12/65869f3ee99a90bf2ffdbe8c.png)

打开`「访达（Finder）->应用程序」`，最后选中`损坏的软件`右键点击打开。

![](https://billy.taoxiaoxin.club/md/2023/12/65869ea748c16ea93a9bf15f.png)

点击打开

![image-20230422223045718](https://billy.taoxiaoxin.club/md/2023/12/65869ea7783d13abb592a88e.png)、