哈喽,大家好!

今天给大家推荐一款数据库管理工具--**DataGrip**

我发现很多人在管理数据库时都喜欢使用`navicat`,目前市面上比较常用的数据库管理工具主要要三种:

1. navicat 官网:https://www.navicat.com/
2. dbeaver 官网:https://dbeaver.io/
3. DataGrip 官网:https://www.jetbrains.com/datagrip/

## 简介

当然,我也之前使用过navicat，它是一款非常流行的数据库管理工具。

然而，我发现DataGrip比navicat更适合我的需求。

首先，相比于`navicat`和`dbeaver`，`DataGrip`更加智能化和高效。它能够识别表中的主键和外键，**代码补全功能也很智能，能够自动识别数据库中的表和列，并提供有用的代码提示**。此外，它还具有更好的代码重构功能和更好的结果集过滤器。

其次，**DataGrip的可视化数据编辑器非常好用**，可以让我**直接在界面上进行数据修改**，而不是在命令行中输入SQL语句。最后，DataGrip的代码导航功能非常强大，可以让我轻松地跳转到想要查看的代码位置。

`DataGrip`支持多种数据库，包括MySQL、PostgreSQL、Oracle和SQL Server等。它还支持在多个数据库之间进行快速导航和搜索，方便您在不同的数据库之间切换。

最后，`DataGrip`的用户界面非常友好，易于使用，适合各种技能水平的用户。

## 安装

### 1.下载程序

首先，我们需要从JetBrains官网（https://www.jetbrains.com/datagrip/）下载DataGrip的安装程序。根据你的操作系统选择合适的版本。

![image-20230503225126177](https://billy.taoxiaoxin.club/md/2023/05/645274ee922ee48dff12d64a.png)

### 2.Mac 安装

下载好之后,在下载里面点击`datagrip-xxx.dmg`安装

![image-20230503225704859](https://billy.taoxiaoxin.club/md/2023/05/64527641922ee48f2f03e87f.png)

直接拖进去`Applications`

![image-20230503225828710](https://billy.taoxiaoxin.club/md/2023/05/64527694922ee48f89796969.png)

### 2. Window 安装

1. 双击运行安装包安装,选择是

   ![image-20230503230625842](https://billy.taoxiaoxin.club/md/2023/05/64527872922ee493e677ed9e.png)

2. 点击Next

![image-20230503230700454](https://billy.taoxiaoxin.club/md/2023/05/64527894922ee4940a820a4d.png)

2. 下载完成选择合适的路径即可。

![image-20230503230752185](https://billy.taoxiaoxin.club/md/2023/05/645278c8922ee4943f00eca4.png)

3. 选择创建快捷方式,点下一步。

![image-20230503230835701](https://billy.taoxiaoxin.club/md/2023/05/645278f3922ee4946ac6539a.png)

4. 点击下一步`install`安装即可

![image-20230423214000007](https://billy.taoxiaoxin.club/md/2023/05/645277a5922ee492c62b3e5e.png)

5. 等待安装完成

![image-20230503230902878](https://billy.taoxiaoxin.club/md/2023/05/6452790f922ee4948321cda5.png)

6. 等待安装完成之后点击Finish。

![image-20230503231005972](https://billy.taoxiaoxin.club/md/2023/05/6452794e922ee494bd3a3329.png)

## 激活

### 1.Mac 激活插件安装

首先下载激活插件并解压**(文末获取)**

解压后打开目录jetbra/scripts

![image-20230423215531087](https://billy.taoxiaoxin.club/md/2023/05/6452743e922ee48d5b0ce18a.png)

打开Mac 终端

![image-20230423220113133](https://billy.taoxiaoxin.club/md/2023/05/6452743e922ee48d5c0655d4.png)

切换到你的目录 `jetbra/scripts`, 执行如下命令:

```bash
cd Downloads/jetbra/scripts # 切换你的脚本目录
sh install.sh # 执行脚本
```

![image-20230423221037303](https://billy.taoxiaoxin.club/md/2023/05/6452743e922ee48d5d9e4e47.png)

### 2. Window 安装

软件安装完成后无需打开软件,就算之前试用期过了也没事,

点击打开`jetbra/scripts`目录

![image-20230423221806949](https://billy.taoxiaoxin.club/md/2023/05/6452743e922ee48d60d7905d.png)

Windows选择任意一个插件安装脚本运行即可,然后点击确定,确定即可.

![image-20220721215756894](https://billy.taoxiaoxin.club/md/2023/05/6452743e922ee48d6135d199.png)

执行过程中，需要大概10-30秒时间，这个根据个人安装的Jetbrains下IDE的数量决定，会在参数文件中，添加激活工具的路径进去。

![image-20220415134709202](https://billy.taoxiaoxin.club/md/2023/05/6452743e922ee48d620c5b40.png)

看到提示“Done”时，表示成功并完成。

![激活完成](https://billy.taoxiaoxin.club/md/2023/05/6452743e922ee48d631e4f05.jpeg)

### 3.输入激活码

安装好激活插件后,打开软件,复制我给大家准备的激活码(**链接放在文末**)

输入激活码激活软件即可

![image-20230503232042048](https://billy.taoxiaoxin.club/md/2023/05/64527bca922ee49fd1906519.png)

最后打开`help`--->`Register`

![image-20220721220611855](https://billy.taoxiaoxin.club/md/2023/05/6452743f922ee48d65b72b00.png)

打开如图所示,成功激活,激活到2025年!

![image-20230503232112401](https://billy.taoxiaoxin.club/md/2023/05/64527be8922ee49fed7cb5af.png)

## 连接数据库

首先打开 DataGrip IDE,点击“新建项目”按钮。

![image-20230503234046524](https://billy.taoxiaoxin.club/md/2023/05/6452807f922ee4a94cbdffd0.png)

意思就是新建了一个文件夹,这里我新建一个test 文件夹,点击确定

![image-20230503234243954](https://billy.taoxiaoxin.club/md/2023/05/645280f4922ee4a9b4057a91.png)

选择在当前窗口打开.

![image-20230503234332817](https://billy.taoxiaoxin.club/md/2023/05/64528125922ee4a9e1ede031.png)

点击如图所示,选择你的数据库类型,这里我以MySQL 举例子

![image-20230503234752386](https://billy.taoxiaoxin.club/md/2023/05/64528228922ee4aad069ec08.png)

首先,选择数据库对应驱动,点击`Driver`选择驱动

![image-20230504000121678](https://billy.taoxiaoxin.club/md/2023/05/64528551922ee4ada9cf5585.png)

选择好驱动之后,点击下载数据库驱动,点击`如图下载`

![image-20230504000446504](https://billy.taoxiaoxin.club/md/2023/05/6452861e922ee4ae6376dd14.png)

等待下载完成

![image-20230504000638737](https://billy.taoxiaoxin.club/md/2023/05/6452868e922ee4aecd828b13.png)

填写您的 MySQL 数据库连接信息。这些信息包括数据库主机名、数据库地址,端口号、用户名和密码。

![image-20230503235527127](https://billy.taoxiaoxin.club/md/2023/05/645287f3922ee4b01124fb9b.png)

选择是否保存数据库密码,

![image-20230503235914426](https://billy.taoxiaoxin.club/md/2023/05/645287f3922ee4b012e9eb13.png)

点击“测试连接”按钮验证连接是否成功。

![image-20230504001048327](https://billy.taoxiaoxin.club/md/2023/05/6452882b922ee4b045210321.png)

如果出现如图所示表示链接成功

![image-20230504001437810](https://billy.taoxiaoxin.club/md/2023/05/6452886e922ee4b08220da8d.png)

### **温馨提示(很重要):**

+ **如果你的链接地址,密码,用户名,数据库,端口号等信息都填写的是对的,显示链接失败,请更换其他驱动试试,可能是当前你选择的驱动不支持**
+ **切换好驱动后,再次下载驱动重试,执行上面的操作**

## 指定数据库连接

如果要链接指定的数据库,请在`Database `中填入

![image-20230504002012974](https://billy.taoxiaoxin.club/md/2023/05/645289bd922ee4b1a877ca53.png)

我这里不需要指定数据库,所以默认为空,不填写即可,

**最后点击OK即可**

![image-20230504002137348](https://billy.taoxiaoxin.club/md/2023/05/64528a11922ee4b1f63a2e41.png)

现在您已经成功连接到 MySQL 数据库。

![image-20230504002505882](https://billy.taoxiaoxin.club/md/2023/05/64528ae2922ee4b2cade1c8b.png)

## 新增数据库

选中数据库连接名称,右键点击数据库连接名称，点击New->Schema,按照如图所示

![image-20230504002753901](https://billy.taoxiaoxin.club/md/2023/05/64528b8a922ee4b3605e6c0b.png)

输入数据库的名字，如`db`

![image-20230504003150632](https://billy.taoxiaoxin.club/md/2023/05/64528c76922ee4b441695e08.png)

最后点击OK

![image-20230504003225045](https://billy.taoxiaoxin.club/md/2023/05/64528c99922ee4b4622da9ac.png)

然后在控制台可以看到你刚刚创建的数据库啦

![image-20230504003434812](https://billy.taoxiaoxin.club/md/2023/05/64528d1b922ee4b4dc457d33.png)

## 添加显示或隐藏已有数据库

右键点击主机名，点击Properties

![image-20230504004536489](https://billy.taoxiaoxin.club/md/2023/05/64528fb0922ee4b736b57385.png)

点击Schemas，然后勾选想要显示的数据库，最后点击Apply->Ok
如果想要隐藏某个数据库，取消勾选，然后点击Apply->Ok即可
如果想要操作所有数据库，则勾选`All shemas`

![image-20230504004758589](https://billy.taoxiaoxin.club/md/2023/05/64529075922ee4b7f68734bd.png)

在`控制台`里可以看到目前所要显示的数据库

![image-20230504004920487](https://billy.taoxiaoxin.club/md/2023/05/64529090922ee4b811f8129c.png)

## 删除数据库

右键点击所要删除的数据库，然点击Drop

![img](https://billy.taoxiaoxin.club/md/2023/05/6453c7be922ee4cb4241b56d.png)

## 新建SQL编辑器

DataGrip 的 SQL 编辑器可以让你轻松地编写和运行 SQL 查询。它支持多种数据库类型，包括 MySQL、PostgreSQL、Oracle 等。你可以使用 DataGrip 的 SQL 编辑器编写复杂的 SQL 查询，并通过数据可视化工具来更好地理解查询结果。

通过鼠标右键,点击新增--->点击信息SQL编辑器,如图:

![image-20230504005459983](https://billy.taoxiaoxin.club/md/2023/05/645291e4922ee4b93c973657.png)

## 运行SQL

选中要运行的SQL,点击运行按钮

![image-20230504005934170](https://billy.taoxiaoxin.club/md/2023/05/645292f6922ee4ba2ca98b3d.png)

控制台会看到执行SQL日志

![image-20230504010011036](https://billy.taoxiaoxin.club/md/2023/05/6452931b922ee4ba510dd880.png)

这样就创建好一张学生表了,在控制台点击数据库名,tables就看到了

![image-20230504010118045](https://billy.taoxiaoxin.club/md/2023/05/6452935e922ee4ba8f527f79.png)

## 新建数据表

右键点击数据库,选择新建-->选择Table

![image-20230504010754876](https://billy.taoxiaoxin.club/md/2023/05/645294eb922ee4bbffc48be1.png)

比如我们新增一个teacher表,字段包括 编号id,姓名name,：

输入表名为`teacher`

![image-20230504011451549](https://billy.taoxiaoxin.club/md/2023/05/6452968b922ee4bd7040a9f9.png)

创建字段名,点击`columns`,点击➕号

![image-20230504012054182](https://billy.taoxiaoxin.club/md/2023/05/645297f6922ee4beb0c333c0.png)

选择字段类型,并且设置为不可以为空

![image-20230504012348137](https://billy.taoxiaoxin.club/md/2023/05/645298a4922ee4bf4aa44056.png)

继续新增name字段,如图:

![image-20230504012821787](https://billy.taoxiaoxin.club/md/2023/05/645299b5922ee4c048cca8a5.png)

设置不可以为空,字符长度为100

![image-20230504012931051](https://billy.taoxiaoxin.club/md/2023/05/645299fb922ee4c08dbe6d28.png)

如果不小心写错了,可以选中字段名,点击`-`号删除

![image-20230504013106013](https://billy.taoxiaoxin.club/md/2023/05/64529a5a922ee4c0e4620a8a.png)



## 修改表字段

### 1.新增表字段

右键点击表名--->columns---->Column

![image-20230504010517899](https://billy.taoxiaoxin.club/md/2023/05/6452944e922ee4bb66bf4022.png)

点击`columns`,点击➕号

![image-20230504012054182](https://billy.taoxiaoxin.club/md/2023/05/6453c460922ee4c8323eeb3b.png)

### 2.修改表字段名称

直接选中表字段name 编辑后保存就好了

![image-20230504224548516](https://billy.taoxiaoxin.club/md/2023/05/6453c51c922ee4c8dc0246a8.png)

### 3.增加索引

选中index--->右键--->index

![image-20230504224929029](https://billy.taoxiaoxin.club/md/2023/05/6453c5f9922ee4c9aec3bd13.png)

然后输入相关信息就好了,点击OK就好了

![image-20230504225016714](https://billy.taoxiaoxin.club/md/2023/05/6453c628922ee4c9d97cd4a1.png)

## 完整SQL 日志

所有查询均记录在文本文件中。 要打开该文件，请转到 *Help | Show SQL log*。

![img](https://billy.taoxiaoxin.club/md/2023/05/6453c937922ee4cc93a173a2.png)

## 数据编辑器

DataGrip 的内置数据编辑器可以让你轻松地查看和编辑数据。

而 SQL 编辑器可以让我运行复杂的 SQL 查询。此外，DataGrip 的可视化工具可以帮助我更好地理解和分析数据。

![img](https://billy.taoxiaoxin.club/md/2023/05/6453ca4d922ee4cd88655e10.gif)

## 新增数据

**可以直接使用内置的编辑器进行编辑操作。**

运行SQL--->点击➕号

![image-20230504231331367](https://billy.taoxiaoxin.club/md/2023/05/6453cb9b922ee4ceb9a7499a.png)

输入内容后点击提交

![image-20230504231947760](https://billy.taoxiaoxin.club/md/2023/05/6453cd13922ee4d01b227547.png)

如果你要查看SQL,可以选择点击第二种提交

![image-20230504232104056](https://billy.taoxiaoxin.club/md/2023/05/6453cd60922ee4d05d267810.png)

最后点击提交就好了

![image-20230504232125895](https://billy.taoxiaoxin.club/md/2023/05/6453cd76922ee4d07a5d2e6d.png)

## 删除数据

选中要删除的某一行,直接点击`-`号,然后提交

![image-20230504232528423](https://billy.taoxiaoxin.club/md/2023/05/6453ce68922ee4d155bc0945.png)

## 编辑数据

直接在数据编辑器里面编辑就好了,最后点击提交

## 导出数据

此外，DataGrip 还支持导出多种数据格式，包括 JSON、XML、CSV 等。

直接点击右上角下载

![image-20230504232817500](https://billy.taoxiaoxin.club/md/2023/05/6453cf11922ee4d1f351f8cf.png)'

选择路径,直接导出就好了

![image-20230504232922316](https://billy.taoxiaoxin.club/md/2023/05/6453cf52922ee4d22d421633.png)

## SQL 自动补全

在使用 DataGrip 时，我最喜欢的功能是它的 SQL 自动完成功能。它不仅可以自动补全 SQL 查询的关键字和语法，还可以补全表名、列名等。这可以大大提高编写 SQL 查询的效率和准确性。

当你输入 SQL 查询时，DataGrip 会自动显示与你输入的内容相关的选项，并根据你的输入进行过滤和排序。这使得编写 SQL 查询变得更加容易和准确。

![image-20230504230705381](https://billy.taoxiaoxin.club/md/2023/05/6453ca19922ee4cd56e5568e.png)

## 导出表数据关系图

![image-20230504233407458](https://billy.taoxiaoxin.club/md/2023/05/6453d06f922ee4d331ae5462.png)

右键选中图,然后导出来就好了

![image-20230504233734356](https://billy.taoxiaoxin.club/md/2023/05/6453d13e922ee4d3eadf4370.png)

## 激活码与激活插件下载地址

激活码与激活插件下载地址

```bash
https://txx.lanzoub.com/iIdDA0upu7wf
```

## 总结

DataGrip 是一款功能强大、易于使用的数据库 IDE。它提供了许多方便的工具和功能，可以大大提高你的工作效率。如果你正在寻找一个方便、易于使用的 SQL 工具，我强烈推荐你试试 DataGrip。

我相信，当你第一眼看到DataGrip以后，会有一种惊艳的感觉，就好比你第一眼看到一个姑娘，就是那么一瞥，你对自己说，就是她了！

希望这篇 DataGrip 保姆级教程能够帮助你入门并充分利用 DataGrip 的各种功能和工具。如果你有任何问题或反馈，请随时通过评论区留言，我会尽力为你解答。

**好了,今天的分享就到这里,觉得有用请点个赞和在看!**

