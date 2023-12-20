# HTML初识 与介绍

## 一、web服务的本质

Web服务的本质是通过HTTP协议进行通信的一种网络服务。下述代码片段展示了一个简单的Python程序，实现了一个基本的Web服务器。

```python
import socket

server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server.bind(('127.0.0.1', 8080))
server.listen(5)

while True:
    conn, client_addr = server.accept()
    data_bytes = conn.recv(1024)
    print(data_bytes)
    conn.send(b'http/1.1 200 OK \r\n\r\n')
    # 1. 纯文本
    # conn.send(b'hello world!')

    # 2. 标签
    # conn.send(b'<h1>hello world!</h1>')

    # 3. 使用文件方式
    """
    文件内容如下:
    <h1>hello world!</h1>
    <a href="https://www.mzitu.com/">click me! give you some color to see see!</a>
    <img src="https://dss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=2159057472,1466656787&fm=26&gp=0.jpg" />
    """
    with open('a.txt', 'rb') as f:
        conn.send(f.read())
    conn.close()


"""
=================== 请求首行 ===================
b'GET / HTTP/1.1\r\n

=================== 请求头 ===================   
Host: 127.0.0.1:8080\r\n  
Connection: keep-alive\r\n
Pragma: no-cache\r\n
Cache-Control: no-cache\r\n
Upgrade-Insecure-Requests: 1\r\n
User-Agent: Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36\r\n
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3\r\n
Accept-Encoding: gzip, deflate, br\r\n
Accept-Language: zh-CN,zh;q=0.9,en;q=0.8\r\n

=================== \r\n ===================
\r\n'  

=================== 请求体. 只有收到post请求方式才有 ===================

"""
```



首先，通过`socket`模块创建一个套接字对象`server`，并将其绑定到本地地址（`127.0.0.1`）和端口号（`8080`）。然后，使用`listen()`方法开始监听传入的连接请求。

进入主循环后，通过`accept()`方法接受客户端的连接请求，并返回一个新的套接字对象`conn`，以及客户端的地址信息`client_addr`。接下来，使用`recv()`方法接收客户端发送的数据，其中`1024`表示每次接收的最大字节数。

在这个简单的例子中，我们打印了接收到的请求数据`data_bytes`，然后通过`send()`方法发送响应给客户端。根据注释的不同选项，我们可以发送不同类型的响应内容。

1. 如果选择纯文本作为响应内容，可以发送类似`b'hello world!'`的字节串。
2. 如果选择带有HTML标签的响应内容，可以发送类似`b'<h1>hello world!</h1>'`的字节串。
3. 如果选择使用文件方式，可以通过`open()`函数打开一个文件，读取文件内容并发送给客户端。

当响应发送完成后，使用`close()`方法关闭连接。

执行流程如下：

1. 浏览器发送HTTP请求到服务器，请求包括请求首行、请求头和请求体。
2. 服务器通过套接字监听请求，接受连接请求并获取请求数据。
3. 服务器处理请求，根据请求内容生成相应的响应内容。
4. 服务器将响应内容发送给浏览器。
5. 浏览器接收到响应后，根据响应内容进行页面渲染。

## 二、HTML简介

用户使用浏览器打开网页看到结果的过程就是：

浏览器将服务端的文本文件(即网页文件)内容下载到本地，然后打开显示的过程。而文本文件的文档结构只有空格和换行两种组织方式，如果仅凭它俩，文本文件在打开显示时，显示的效果将会非常非常非常的单一。

为了让显示的效果不那么单调，我们会偏向使用word一类的文本编辑工具来编排文本内容，编排的原理就是：**在编辑文件时会选中各部分内容，然后为内容打上不同的标记，比如什么是标题，什么是段落，然后存放硬盘里**，等下次打开时，word会识别之前的标记，然后按照预先编排好的结果显示出来

### 1.什么是HTML？

站在显示文本内容的角度去看，浏览器与word的原理一样，我们可以将浏览器当成一个网页版的只读word，浏览器也必须有一套自己能识别的标记文本的规范，**该规范被称为HTML,**HTML全称是超文本标记语言（`HyperText Markup Language`），是由W3C的维护的。

**HTML的本质：**

+ 本质上是浏览器可识别的规则，我们按照规则写网页，浏览器根据规则渲染我们的网页。对于不同的浏览器，对同一个标签可能会有不同的解释。（兼容性问题）
+ 我们浏览器看到的页面，内部其实都是HTML代码(所有的网站内部都是HTML代码)
+ HTML文件是一个文本文件,包含了一些HTML元素,标签等

**作用：**

+ HTML用来制作网页的标记语言
+ HTML 是通向 WEB 技术世界的钥匙。
- 不需要编译,直接由浏览器执行

**文件的扩展名：**

+ HTML对大小写不敏感的,HTML与html是一样的
+ HTML文件必须使用.html或.htm为文件名后缀

**强调**：**它不是一种编程语言**，而是一种标记语言`（markup language）`。

**注意**：如果你想要让浏览器能够渲染出你写的页面。你就必须遵循HTML语法

**“超文本”指的是用超链接的方法，将各种不同空间的文字信息组织在一起的网状文本**

**“标记”指的是在编辑文本时用特殊的记号标记一下各部分内容的意义，该记号称之为标签**，比如用标签h1标记标题，用标签p标签段落，如此我们标记一首唐诗就成了如下格式：

```html
<h1>剑客 / 述剑</h1>
<p>【作者】贾岛 【朝代】唐</p>
<p>十年磨一剑，霜刃未曾试。</p>
<p>今日把示君，谁有不平事？</p>
```

![image-20210302224309731](https://billy.taoxiaoxin.club/md/2023/05/6474bddf922ee44da7eeecf1.png)

**所以我们学习HTML就是在学习一系列的标签**

### 2.HTML不是编程语言

HTML是一种标记语言（markup language），它不是一种编程语言。

HTML使用标签来描述网页。

![img](https://billy.taoxiaoxin.club/md/2023/05/6474bddf922ee44da8044988.png)

## 三、发展历史

### 1. 发展历史介绍

![HTML发展史](https://billy.taoxiaoxin.club/md/2023/05/6474bddf922ee44da955c029.png)



```python
超文本标记语言（第一版）：在1993年6月作为互联网工程工作小组（IETF）工作草案发布（并非标准），后来陆续由w3c制定标准

#IETF简介
IETF是英文Internet Engineering Task Force的缩写, 翻译过来就是"互联网工程任务组"
IETF负责定义并管理因特网技术的所有方面。包括用于数据传输的IP协议、让域名与IP地址匹配的域名系统（DNS）、用于发送邮件的简单邮件传输协议（SMTP）等

#W3C简介
W3C是英文World Wide Web Consortium的缩写， 翻译过来就是W3C理事会或万维网联盟, W3C是全球互联网最具权威的技术标准化组织.
W3C于1994年10月在麻省理工学院计算机科学实验室成立。创建者是万维网的发明者Tim Berners-Lee
W3C负责web方面标准的制定，像HTML、XHTML、CSS、XML的标准就是由W3C来定制的。 
Tim Berners-Lee（蒂姆·伯纳斯-李），万维网之父、html设计者、w3c创始人
```

目前常用的两种文档类型是xhtml 1.0和html5，pc端两种都可以，而且html5是向下兼容的

### 2. HTML5的由来

- HTML5草案的前身名为 Web Applications 1.0，于2004年被WHATWG提出，于2007年被W3C接纳，并成立了新的 HTML 工作团队。
- HTML 5 的第一份正式草案已于2008年1月22日公布。HTML5 仍处于完善之中。然而，大部分现代浏览器已经具备了某些 HTML5 支持。
- 2012年12月17日，万维网联盟（W3C）正式宣布凝结了大量网络工作者心血的HTML5规范已经正式定稿。根据W3C的发言稿称：“HTML5是开放的Web网络平台的奠基石。”
- 2013年5月6日， HTML 5.1正式草案公布。该规范定义了第五次重大版本，第一次要修订万维网的核心语言：超文本标记语言（HTML）。在这个版本中，新功能不断推出，以帮助Web应用程序的作者，努力提高新元素互操作性。
- 2014年10月29日，万维网联盟宣布，经过接近8年的艰苦努力，该标准规范终于制定完成。

### 4.HTML5的兼容性

- Internet Explorer 9 以及 以上版本
- chrome、Safari、opera、Firefox和各种以wekkit为内核的国产浏览器

### 5. **HTML5与XTML的区别**

```python
#1、XHTML更为严格,它要求标签必须小写、必须严格闭合、标签中的属性必须使用引号引起，img必须要加alt属性(对图片的描述)
等等;

#2、HTML5是HTML的下一个版本所以除了非常宽松容错性强（可以选择性遵守xhtml制定的文档编写规范）以外,还增加许多新的特性
```

**xhtml1.0文档类型创建的快捷方式 html:xt + tab**

**html5文档类型创建的快捷方式： html:5 + tab 或者 ! + tab**

## 四、两种打开HTML文件的方式

+ 方式一：找到文件所在的位置右键选择浏览器打开
+ 方式二：在pycharm内部，集成了自动调用浏览器的功能，直接点击即可(前提是你的电脑上安装了对应的浏览器) 直接全部使用谷歌浏览器

