# HTML常用标签详解

## 一、head内常用标签

### 1.基本标签（非meta标签）

```html
<!--title: 定义网页标题-->
<title>Title</title>

<!--style: 定义内部样式表. 内部用来书写css代码-->
<style>
    h1 {
        color: greenyellow;
    }
</style>

<!--script: 内部用来书写js代码-->
<script>
    alert(123)
</script>
<!--script: 还可以引入外部js文件-->
<script src="myjs.js"></script>

<!--link: 引入外部css文件 或 网站图标-->
<link rel="stylesheet" href="mycss.css">
```



### 2.meta相关

#### 2.1 Meta标签介绍

+ `<meta>`元素可提供有关页面的元信息（mata-information）,针对搜索引擎和更新频度的描述和关键词。
+ `<meta>`标签位于文档的头部，不包含任何内容。
+ `<meta>`提供的信息是用户不可见的。

#### 2.2 Meta标签的组成

meta标签共有两个属性，他们分别是`http-equiv`属性和`name` 属性，不同的属性又有不同的参数值，这些不同的参数值就实现了不同的网页功能

##### http-equiv 属性: 

相当于http的文件头作用，它可以向浏览器传回一些有用的信息，以帮助正确地显示网页内容，与之对应的属性值为content，content中的内容其实就是各个参数的变量值。

> 补充: `equiv `全称 `equivalent /ɪˈkwɪvələnt/ `相等的 等值 相当于

```html
<!--指定文档的编码类型(需要知道)-->
<meta http-equiv="content-Type" charset=UTF8">
                                             
<!--2秒后跳转到对应的网址，注意引号(了解)-->
<meta http-equiv="refresh" content="2;URL=https://www.baidu.com">
                                                                
<!--3秒后刷新(了解)-->                                                                
<meta http-equiv="refresh" content="3">   
                                      
<!--告诉IE以最高级模式渲染文档(了解)-->
<meta http-equiv="x-ua-compatible" content="IE=edge">
```

##### name 属性

主要用于**描述网页**，与之对应的属性值为`content`，content中的内容主要是便于搜索引擎机器人查找信息和分类信息用的。

```html
<!--关键字：有助于搜索引擎SEO优化，再怎么优化也抵不过竞价排名-->
<meta name="keywords" content="meta总结,html meta,meta属性,meta跳转">

<!--页面描述-->
<meta name="description" content="给你骨质唱疏通">
```

> meta标签必须写在头部head标签之内，而description的meta标签务必要写在keywords的meta标签之后，像下面这样的顺序写：
>
> 1. **关键字**：description要用简短的文字描述网站或网页的主要内容，有利于各大搜索引擎的抓取收录你的网站或网页。
> 2. **页面描述**：当你设置了description网站描述文字，才会显示在搜索引擎的结果页中，而每个网页的description也是同样的道理，简短又准确的网页描述文字，可以帮助用户在搜索引擎中更方便的找到你的网站和网页！

### 3. 总结

```html
<title>网页标题</title>

<style></style>
<!--rel	Reload 加载-->
<!--href hypertext reference 超文本引用-->
<link rel="stylesheet" href="url">

<script></script>
<!--src	Source 源文件链接-->
<script scr="源文件链接"></script>

<meta name="keywords" content="关键字1, 关键字2, 关键字3">
<meta name="description" content="网页描述信息">
<meta http-equiv="content-type" charset="字符编码">
```

## 二、body内常用标签

### 1. HTML语义化

#### 1.1 body标签介绍

+ body中的标签是会显示到浏览器窗口中的
+ body内的标签只有一个作用就是用来标记语义的，语义指的是从字面意思就可以理解被标记的内容是用来做什么的

> 虽然不同的标签会有不同的显示样式，但我们一定要强制自己忘记所有标签的显示样式，只记它的语义。因为每个标签的样式各不相同，记忆不同标签的样式来控制显示样式，对前端开发来说将会是一种灾难，更何况添加样式并不是HTML擅长的事情，而且在布局的时候多使用语义化的标签，会方便搜索引擎在爬取网页文件时更好地解析文档结构，从而进行收录。

#### 1.2 基本标签（块级标签和内联标签）

对于那些只用来修改样式的标签将会被淘汰掉，比如以下标签都是没有语义的，都是用来修改样式的

```html
<!--b Bold 粗体（文本）-->
<b>加粗</b>

<!--i Italic /ɪˈtælɪk/ 斜体（文本）-->
<i>斜体</i>

<!--u Underlined 下划线（文本）-->
<u>下划线</u>

<!--s Strikethrough  /straɪk/	删除线-->
<s>删除</s>

<p>段落标签</p>

<!--h1~h6 Header 1 to Header 6 标题1到标题6-->
<h1>标题1</h1>
<h2>标题2</h2>
<h3>标题3</h3>
<h4>标题4</h4>
<h5>标题5</h5>
<h6>标题6</h6>

<!--br Break 换行-->
<br>

<!--hr Horizontal Rule  /ˌhɒrɪˈzɒntl/ 水平线，分割线-->
<hr>

<!--br Break 修改文字大小，颜色-->
<font color="red" size="10px">我是菜鸟</font>
```

HTML5中推出了一些新的标签

```html
strong == b

        ins == u

        em == i

        del == s

新的标签是有语义的，而老的只是单纯的添加样式（这是CSS干的事）
        strong的语义：定义重要性强调的文字
        ins的语义（inserted）：定义插入的文字
        em的语义（emphasized）：定义强调的文字
        del的语义（deleted）：定义被删除的文字
```

### 2. div标签和span标签

+ **div标签**（块级标签）用来定义一个块级元素，并无实际意义。主要通过CSS样式为其赋予不同的表现。

  > div用来标记一块内容，没有具体的语义

+ **span标签**用来定义内联（行内）元素，并无实际的意义。主要通过CSS样式为其赋予不同的表现。

  > span用来标记一行中的一小段内容，也没有具体的语义。

补充：全称及缩写说明

+ `div Division` 分隔
+ `span Span` 范围

### 3. 标签的分类：块级/行内

块级元素与行内元素的区别

+ 块级标签：独占一行（h1~h6	p div）

+ 块儿级标签可以修改长宽. 行内标签不可以, 就算修改了也不会变化.

+ 块级标签内部可以嵌套任意的块级标签和行内标签. 

  > 特例: 但是p标签虽然是块级标签 但是它只能嵌套行内标签 不能嵌套块级标签. 如果你套了 问题也不大 因为浏览器会自动帮你解除嵌套关系(浏览器是直接面向用户的 不会轻易的报错 哪怕有报错用户也基本感觉不出来)

+ **行内标签**: 自身文本多大就占多大.·`i u s b span`

+ 行内标签不能嵌套块级标签, 只能嵌套行内标签.

+ 补充：上述的规定只是HTML书写规范 如果你不遵循 不会报错

> 总结：
>
> 1. 只要是块儿级标签都可以嵌套任意的块儿级标签和行内标签
> 2. 但是p标签只能嵌套行内标签，不能包含块级标签，p标签也不能包含p标签（HTML书写规范）
> 3. 内联元素不能包含块级元素，它只能包含其它内联元素

### 4. 实体字符（特殊符号）

**注意**：**在HTML中对空格／回车／tab不敏感，会把多个空格／回车／tab当作一个空格来处理**

#### 4.1 什么是实体字符？

在HTML中，有的字符是被保留的比如大于号小于号。

有的字符是被HTML保留的比如大于号小于号
有的HTML字符，在HTML中是有特殊含义的，是不能在浏览器中直接显示出来的，那么这些东西想显示出来就必须通过字符实体，如下:

> **注释**：实体名称对大小写敏感！！

| 内容 |   代码   |
| :--: | :------: |
| 空格 | `&nbsp;` |
|  >   |  `&gt;`  |
|  <   |  `&lt;`  |
|  &   | `&amp;`  |
|  ¥   | `&yen;`  |
| 版权 | `&copy;` |
| 注册 | `&reg;`  |

[HTML特殊符号对照表](http://tool.chinaz.com/Tools/HtmlChar.aspx)

### 5.h系列标签

**语义：标记内容为一个标题，全称headline**

**h系列标签从h1-h6共6个，没有h7标签**，标记内容为1~6级标题，h1用作主标题（代表最重要的），其次是h2，依次往下排序，直到H6。

虽然h1-h6标签的显示样式是从大到小，但再次强调：记忆HTML标签的显示样式是没有意义的

```html
<!DOCTYPE HTML>
<html>
    <head lang='en'>
        <meta charset="utf-8">
        <title>Egon才华无敌</title>
    </head>
    <body>
        <h1>一级标题</h1><h2>二级标题</h2>
        <h3>三级标题</h3>
        <h4>四级标题</h4>
        <h5>五级标题</h5>
        <h6>六级标题</h6>
        <h7>没有七级标题</h7>
        没有七级标题
    </body>
</html> 
```

> **注意**：
>
> 在企业开发中一定要慎用h系列标签，特别是h1标签，在企业开发中一般一个界面中只能出现一个h1标签（出于SEO考虑，搜索引擎会使用标题将网页的结构和内容编制索引）,比如www.163.com。

### 6.p标签

语义：标记内容为一个段落，全称`paragraph`

```html
<!DOCTYPE HTML>
<html>
    <head lang='en'>
        <meta charset="utf-8">
        <title>Egon无敌</title>
    </head>
    <body>
        <h1>Egon</h1>
        <p>论颜值，鹤立鸡群</p>
        <p>论才华，天下无敌</p>
    </body>
</html>
```

### 7. img标签

语义：标记一个图片，全称image

##### 1. 用法

```html
<img src="图片地址" alt="图片加载失败时显示的内容" title = "鼠标悬停到图片上时显示的内容" />
```

##### 2. 注意

+ **src**指定的图片地址可以是网络地址，也可以是一个本地地址，本地地址可以用绝对或相对路径，但通常用相对路径，相对路径是以html文件当前所在路径为基准进行的

+ 图片的格式可以是`png`、`jpg`和`gif`

+ `alt="图片加载失败时显示的内容" ` ：为`img`标签加上该属性可用于支持搜索引擎和盲人读屏软件。

+ `title = "鼠标悬停到图片上时显示的内容"`

+ `height="800px" ` `width=""`

  > 如果没有指定图片的width和height则按照图片默认的宽高显示，如果指定图片的width和height则可能让图片变形
  > 那如果又想指定宽度和高度，又不想让图片变形，我们可以只指定宽度和高度的一个值即可
  > 只要指定了一个值，系统会根据该值计算另外一个值，并且都是等比拉伸的，图片将不会变形

全称及缩写说明：

+ `alt alter` 替用(一般是图片显示不出的提示）
+ `src Source `源文件链接

### 8.a标签 

#### 8.1 a标签介绍

语义：标记一个内容为超链接，全称`anchor`，锚

> 所谓的超链接是指从一个网页指向一个目标的连接关系，这个目标可以是另一个网页，也可以是相同网页上的不同位置，还可以是一个图片，一个电子邮件地址，一个文件，甚至是一个应用程序

超链接标签是超文本文件的精髓，可以控制页面与页面之间的跳转，语法如下:

```html
<a href="跳转到的目标页面地址" target="是否在新页面中打开" title="鼠标悬浮显示的内容">需要展现给用户查看的内容/也可以是图片</a>
```

`href`属性指定目标网页地址。该地址可以有几种类型：

- 绝对URL - 指向另一个站点`（比如 href=”[http://www.jd.com）](http://www.jd.com)/)`
- 相对URL - 指当前站点中确切的路径`（href=”index.htm”）`
- 锚URL - 指向页面中的锚`（href=”#top”）`

`target：`

- `_blank`表示在新标签页中打开目标网页
- `_self`表示在当前标签页中打开目标网页

#### 8.2 a标签注意点

1. a标签不仅可以标记文字，也可以标记图片

   ```html
   <a href="https://www.baidu.com"><img src="mv.png" />百度一下，你就知道</a>
   ```

   

2. a标签必须有href属性，href的值必须是`http://`或`https://`开头

3. a标签还可以跳转到自己的页面

   ```html
   <a href="template/aaa.html">锤你胸口</a>
   ```

4. target="_blank"代表在新页面中打开，其余的值均无需记忆。

   ```html
   如果页面中大量的a标签都需要设置target="_blank",那么我们可以在head标签内新增一个base标签进行统一设置
   <base target="_blank">
   如果a标签自己设置了target，那么就以自己的为准，否则就会参照base的设置
   ```

5. title="鼠标悬浮显示的

补充：

+ 当a标签指定的网址从来没有被点击过 那么a标签的字体颜色是蓝色 
+ 如果点击过了就会是紫色（浏览器给你记忆了）

#### 8.3 假链接

什么假链接？

+ 就是点击之后不会跳转的链接，我们称之为假链接

假链接存在的意义

+ 在企业开发前期，其他界面都还没有写出来。
+ 那么我们就不知道应该跳转到什么地方，所以只能使用假链接来代替。

假链接的定义格式

```html
1、href="#"   :会自动回到网页的顶部
2、href="javascript:" ：不会返回顶部
```

#### 8.4 页面内锚点

1. 要想通过a标签跳转到指定的位置，那么必须告诉a标签一个独一无二的身份证号码，这样a标签才能在当前界面中找到需要跳转到的目标位置

2. 如何为html中的标签绑定一个独一无二的身份证号码呢？

   > 在html中，每一个标签都有一个名称叫做id的属性,这个属性就是用来给标签指定一个独一无二的身份证号码的

3. 所以要想实现通过a标签跳转到指定的位置，分为两步

   1. 给目标位置的标签添加一个id属性，然后指定一个独一无二的值
   2. 告诉a标签你需要跳转到的目标标签对应的独一无二的身份证号码是多少

4. a标签除了可以跳转当前页面，还可以跳转到其他页面的指定位置 

**a标签的锚点功能**

```html
<!--1.点击一个文本标题 页面自动跳转到标题对应的内容区域-->
<a href="" id="d1">顶部</a>
<h1 id="d4">hello world</h1>
<div style="height: 1800px; background: aliceblue"></div>

<a href="" id="d2">中间</a>
<div style="height: 1800px; background: aqua"></div>

<a href="" id="d3">底部</a>
<div style="height: 1800px; background: cadetblue"></div>

<a href="#d1">回到顶部</a>
<a href="#d2">回到中间</a>
<a href="#d4">回到hello world</a>

<!--2.跳到首页-->
<a href="">刷新页面，回到顶部，人类感觉不出来区别</a> 
<a href="#">回到顶部</a>

<!--3.注意点-->
    通过我们的a标签跳转到指定的位置，是没有过度动画的
    是直接一下子就跳转到了指定位置,比如京东主页
    如果跳到首页需要过渡动画，则不用a标签做，比如天猫主页
```

#### 8.5 全称缩写说明

```python
Anchor /ˈæŋkə(r)/
Division  /dɪˈvɪʒn/
a Anchor 锚(定义超链接，用于从一张页面链接到另一张页面)
href hypertext reference 超文本引用
div  Division 分隔
```

#### 8.6 页面锚点调回首页练习

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>标签锚点练习</title>
</head>
<body>
<!--
    <a href="">刷新页面，回到顶部</a> 刷新页面，人类感觉不出来区别
    <a href="#">回到顶部</a>


-->


<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p id="cn1">我是个大菜鸟</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p id="xd">小弟乐意为您效劳</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <p>我上山是虎，我下海是龙</p>
    <a href="#cn1">找第一个菜鸟</a>
    <br>
    <a href="#xd">找第一个小弟</a>
    <br>
    <a href="#">不刷新，跳回首页</a>
    <br>
    <a href="">刷新，跳回首页</a>
</body>
</html>
```



### 9.列表标签

  语义：标记一堆数据是一个整体/列表

`html`中列表标签分为三种：

#### 9.1 无序列表

**无序列表（列表标签中使用最多的一种，非常重要）：unordered list**

+ 作用：制作导航条，商品列表，新闻列表等

+ 语法如下，组合使用`ul>li`

```html
<!--基本语法-->
<ul type="disc">
  <li>第一项</li>
  <li>第二项</li>
</ul>

<!--示例：京东商城导航条-->
    <ul>
        <li>秒杀</li>
        <li>优惠券</li>
        <li>PLUS会员</li>
        <li>闪购</li>
        <li>拍卖</li>
        <li>京东服饰</li>
        <li>京东超市</li>
        <li>生鲜</li>
        <li>全球购</li>
        <li>京东金融</li>
    </ul>

<!--全称及缩写说明
ul Unordered List 不排序列表
li List Item    列表项目
-->
```

ul标签的type属性（这属于列表的样式，所以了解即可）：

`type`：列表识别的类型

+ disc（实心圆点，默认值）

+ circle（空心圆圈）

+ square（实心方块）

+ none（无样式，不显示标识，可以通过css去掉小圆点，如下：）

  ~~~html
  <style type="text/css">
              ul {
                  list-style: none;
              }
  </style>
  ~~~

**注意**

+ ul与li是组合标签应该一起出现，并且ul的子标签只应该是li，而li的子标签则可以是任意其他标签

无序列表示例

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>无序列表示例</title>
</head>
<body>
    <h1>物品清单</h1>
    <ul>
        <li>
            <h2>蔬菜</h2>
            <ul>
                <li>西红柿</li>
                <li>花瓜</li>
                <li>芹菜</li>
            </ul>
        </li>
        <li>
            <h2>水果</h2>
            <ul>
                <li>香蕉</li>
                <li>菠萝</li>
                <li>火龙果</li>
            </ul>
        </li>
    </ul>
</body>
</html>
```

#### 9.2 有序列表（极少使用）

基本语法如下:

~~~html
<ol type="1" start="2">
  <li>第一项</li>
  <li>第二项</li>
</ol>
~~~

type属性：

```python
1  数字列表，默认值
A  大写字母
a  小写字母
Ⅰ 大写罗马
ⅰ 小写罗马

#全称及缩写说明
'''
ol Ordered List 排序列表
li List Item    列表项目
'''
```

有序列表示例：

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>智商排名</title>
</head>
<body>
<h1>智商排名</h1>
<ol>
    <li>刘思远</li>
    <li>刘清政</li>
    <li>刘sir</li>
    <li>刘铭洋</li>

</ol>
    <!--有序列表能干的事，完全可以用无序列表取代-->
    <h1>智商排名</h1>
    <ul style="list-style: none">
    <li>1.刘思远</li>
    <li>2.刘清政</li>
    <li>3.刘sir</li>
    <li>4.刘铭洋</li>
    </ul>
</body>
</html>
```

#### 9.3 自定义列表（也会经常使用）

作用分析：

选择用什么标签的唯一标准，是看文本的实际语义，而不是看长什么样子

1. 无序列表：内容是并列的,没有先后顺序
2. 有序列表：内容是有先后顺序的
3. 自定义列表：对一个题目进行解释说明的时候，用自定义列表,可以做网站尾部相关信息，网易注册界面的输入框

基本语法如下：

```html
自定义列表也是一个组合标签：dl>dt+dd
dl:defination list，
dt：defination title，自定义标题
dd：defination description，自定义描述
<dl>
  <dt>自定义标题1</dt>
  <dd>内容1</dd>
  <dt>自定义标题2</dt>
  <dd>内容1</dd>
  <dd>内容2</dd>
</dl>

# 全称及缩写说明
"""
Definition /ˌdefɪˈnɪʃn/
dl Definition List 自定义列表
dt Definition title 自定义标题
dd Definition Description 自定义描述
"""
```

注意：

```python
1.dl>dt+dd应该组合出现，dl中只应该存放dt和dd，而可以在dt和dd中添加任意其他标签 
2.一个dt可以可以没有对应的dd，也可以有多个，但建议一个dt对应一个dd
```

示例：

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>自定义列表标签</title>
</head>
<body>
<dl>
    <dt><h5>购物流程</h5></dt>
    <dd>购物流程</dd>
    <dd>会员介绍</dd>
    <dd>生活旅行</dd>

    <dt><h5>配送方式</h5></dt>
    <dd>上门自提</dd>
    <dd>211限时达</dd>
    <dd>配送服务查询</dd>

    <dt><h5>支付方式</h5></dt>
    <dd>货到付款</dd>
    <dd>在线支付</dd>
    <dd>分期付款</dd>
</dl>
</body>
</html>
```

### 10. 表格标签

语义：标记一段数据为表格

**作用**

+ 表格标签是一种数据的展现形式，当数据量非常大的时候，使用表格的形式来展示被认为是最清晰的

  > 表格是一个二维数据空间，一个表格由若干行组成，一个行又有若干单元格组成，单元格里可以包含文字、列表、图案、表单、数字符号、预置文本和其它的表格等内容。
  > 表格最重要的目的是显示表格类数据。表格类数据是指最适合组织为表格格式（即按行和列组织）的数据。

基本结构格式：

```html
<table>
    <tr>
        <td></td>
    </tr>
</table>

tr代表表格的一行数据
td表一行中的一个普通单元格
th表示表头单元格

注意点：
表格标签有一个边框属性，这个属性决定了边框的宽度。默认情况下这个属性的值为0，所以看不到边框 
```

属性：

- border: 表格边框.
- cellpadding: 内边距
- cellspacing: 外边距.
- width: 像素 百分比.（最好通过css来设置长宽）
- rowspan: 单元格竖跨多少行
- colspan: 单元格横跨多少列（即合并单元格）

其他表格属性了解性知识点：

```python
#1、宽度和高度
    可以给table和td设置width和height属性
    
    1.1 默认情况下表格的宽高是按照内容的尺寸来调整的，也可以通过给table标签设置widht和height来手动指定表格的宽高

    1.2 如果给td标签设置width和height属性，会修改当前单元格的宽度和高度，只要不超过table的宽高，则不会影响整个表格的宽度和高度


#2、水平对齐和垂直对齐
    水平对齐align可以给table、tr、td标签设置
    垂直对齐valign只能给tr、td标签设置

    ========水平对齐===========
    取值
    align=“left”
    align=“center”
    align=“right”

    2.1 给table标签设置水平对齐，可以让表格在水平方向上对齐
          强调：table只能设置水平方向

    2.2 给tr设置水平对齐，可以控制当前行所有单元格内容都水平对齐

    2.3 给td设置水平对齐，可以控制当前单元格内容水平对齐，tr与td冲突的情况下，以td为准

    ========垂直对齐===========
    取值
    valign=“top”
    valign=“center”
    valign=“bottom”
    
    2.4 给tr设置垂直对齐可以让当前行所有单元格内容都垂直对齐
    2.5 给td设置垂直对齐可以让当前单元格内容垂直对齐


#3、外边距和内边距
    只能给table设置

    3.1 外边距：单元格与单元格之间的间隔，cellspacing="3px"，默认值为2px
    3.2 内边距：单元格边框与文字之间的距离:cellpadding="200px"
```

三种方式细线表格:

```html
#1、方式一
    在标签中，想通过指定外边距为0来实现细线表格是不靠谱的，其实他是将2条线合成了一条线.所以看上去很不舒服,如下实现
<table width="200px" height="200px" bgcolor="black" border="1" cellspacing="0px">
    <tr bgcolor="white">
        <td>姓名</td>
        <td>性别</td>
        <td>年龄</td>
    </tr>

    <tr bgcolor="white" >
        <td>Egon</td>
        <td>male</td>
        <td>18</td>
    </tr>

    <tr bgcolor="white">
        <td>ALex</td>
        <td>male</td>
        <td>73</td>
    </tr>

    <tr bgcolor="white">
        <td>Wxx</td>
        <td>female</td>
        <td>84</td>
    </tr>
</table>
#2、方式二
 细线表格的制作方式：
        1、给table标签设置bgcolor
        2、给tr标签设置bgcolor
        3、给table标签设置cellspacing="1px"


      注意：
      table、tr、td标签都支持bgcolor属性

<table width="200px" height="200px" bgcolor="black" cellspacing="1px">
    <tr bgcolor="white">
        <td>姓名</td>
        <td>性别</td>
        <td>年龄</td>
    </tr>

    <tr bgcolor="white" >
        <td>Egon</td>
        <td>male</td>
        <td>18</td>
    </tr>

    <tr bgcolor="white">
        <td>ALex</td>
        <td>male</td>
        <td>73</td>
    </tr>

    <tr bgcolor="white">
        <td>Wxx</td>
        <td>female</td>
        <td>84</td>
    </tr>
</table>

#3、方式三（style="border-collapse: collapse;border: 1px solid red"）
<table border="1px" style="border-collapse: collapse;border: 1px solid red">
    <tr>
        <td>姓名</td>
        <td>性别</td>
        <td>年龄</td>
    </tr>
    <tr>
        <td>egon</td>
        <td>male</td>
        <td>18</td>
    </tr>
    <tr>
        <td>alex</td>
        <td>female</td>
        <td>19</td>
    </tr>
</table>
```

#### 表格的结构：

![image-20210304200030757](https://billy.taoxiaoxin.club/md/2023/05/647616a1922ee45c3defa737.png)

表格结构详解：

```html
为了方便管理维护以及提升语义，我们将表格中存储的数据分为四类：
#1、表格的标题:caption
    特点:相对于表格宽度自动居中对齐
    注意:
        1.1 该标签一定要写在table标签里，否则无效
        1.2 caption一定要紧跟在table标签内的第一个

#2、表格的表头信息:thead
    特点：专门用来存储每一列的标题，只要将当前列的标题存储在这个标签中就会自动居中+加粗文字


#3、表格的主体信息:tbody
    注意：
        3.1 如果没有添加tbody，浏览器会自动添加
        3.2 如果指定了thread和tfoot，那么在修改整个表格的高度时，thead和tfoot有自己默认的高度，不会随着
            表格的高度变化而变化

#4、表尾信息:tfoot


<html>
<head>
    <meta charset="utf-8"/>
</head>
<body>
    <table bgcolor="black" border="1" width="300px" height="300px" cellspacing="1px">

        <caption>学员信息统计</caption>
        <thead>
            <tr bgcolor="white">
                <th>姓名</th>
                <th>性别</th>
                <th>年龄</th>
            </tr>
        </thead>

        <tbody>
            <tr bgcolor="white">
                <td>egon</td>
                <td>male</td>
                <td>18</td>
            </tr>

            <tr bgcolor="white">
                <td>egon</td>
                <td>male</td>
                <td>18</td>
            </tr>

            <tr bgcolor="white">
                <td>egon</td>
                <td>male</td>
                <td>18</td>
            </tr>
        </tbody>

        <tfoot>
            <tr bgcolor="white">
                <td>3</td>
                <td>3</td>
                <td>3</td>
            </tr>
        </tfoot>
    </table>

</body>
</html>
```

单元格合并：

```html
#1、水平向上的单元格colspan
    可以给td标签添加一个colspan属性，来把水平方向的单元格当做多个单元格来看待
    <td colspan="2"></td>

#2、垂直向上的单元格rowspan
    可以给td标签设置一个rowspan属性，来把垂直方向的的单元格当成多个去看待

#注意注意注意:
1、由于把某一个单元格当作了多个单元格来看待，所以就会多出一些单元格，所以需要删掉一些单元格
2、一定要记住，单元格合并永远是向后或者向下合并，而不能向前或向上合并
```

传统布局：

```python
传统的布局方式就是使用table来做整体页面的布局，布局的技巧归纳为如下几点：

#1、定义表格宽高，将border、cellpadding、cellspacing全部设置为0

#2、单元格里面嵌套表格

#3、单元格中的元素和嵌套的表格用align和valign设置对齐方式

#4、通过属性或者css样式设置单元格中元素的样式

传统布局目前应用：
#1、快速制作用于演示的html页面

#2、商业推广EDM制作(广告邮件)
```

#### 示例：

|          |          |        |        |
| :------: | :------: | :----: | :----: |
| username | password | hobby  | others |
|  jsaon   |          |  read  |        |
|   egon   |   DBJ    | 吃生蚝 |        |
|   tank   |   摸鱼   | 弹棉花 |        |

```html
<!--<table border="10" cellpadding="5" cellspacing="10">-->
<table border="10">   <!-- 加外边宽-->
    <!--表头: 存放字段信息-->
    <thead>
        <tr>  <!--一个tr就表示一行-->
            <th>username</th>   <!--加粗文本-->
            <td>password</td>   <!--正常文本-->
            <th>hobby</th>
            <th>others</th>
        </tr>
    </thead>

    <!--表单: 存放数据信息-->
    <tbody>
        <tr>
            <td>jason</td>
            <!--rowspan 合并行属性-->
            <td rowspan="3">123</td>
            <!--colspan 合并列属性(合并当前行的列)-->
            <td colspan="2">read</td>
        </tr>
        <tr>
            <td>egon</td>
            <!--<td>123</td>-->
            <td>DBJ</td>
            <td>吃生蚝</td>
        </tr>
        <tr>
            <td>tank</td>
            <!--<td>123</td>-->
            <td>摸鱼</td>
            <td>弹棉花</td>
        </tr>
    </tbody>
</table>
```

+ 创建流程

  ~~~html
  """
  colspan 当前行中的列合并(水平方向占多行)
  rowspan 合并多行属性(垂直方向占多行)
  """
  table[border=1]>(thead>th)+(tbody>tr>td+td)
  ~~~

#### 练习：

![image-20210304200757893](https://billy.taoxiaoxin.club/md/2023/05/647616a2922ee45c3e45cdc7.png)

```html
<!DOCTYPE HTML>
<html>
    <head lang='en'>
        <meta charset="utf-8">
        <title>超级美味无敌菜谱</title>
        <base target="_blank">
        <style>

        </style>
    </head>
    <body>

    <table border="0" cellspacing="1" bgcolor="#D7A7EE"  width="500px" height="200px">
        <tr bgcolor="white">
            <td colspan="3" align="center">星期一菜谱</td>
        </tr>
        <tr bgcolor="white">
            <td rowspan="2">素菜</td>
            <td>情操茄子</td>
            <td>花椒扁豆</td>
        </tr>

        <tr bgcolor="white">
            <td>小葱豆腐</td>
            <td>炒白菜</td>
        </tr>

        <tr bgcolor="white">
            <td rowspan="2">荤菜</td>
            <td>油焖大虾</td>
            <td>海参鱼翅</td>
        </tr>

        <tr bgcolor="white">
            <td>
                红烧肉<img src="https://dss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=1032699337,4070192957&fm=26&gp=0.jpg" alt="" height="220">
            </td>
            <td>烤全羊</td>
        </tr>
    </table>


    </body>
</html>
```

![image-20210304200923713](https://billy.taoxiaoxin.club/md/2023/05/647616a2922ee45c3fe61427.png)

```html
<!DOCTYPE HTML>
<html>
    <head lang='en'>
        <meta charset="utf-8">
        <title>Egon无敌</title>
        <base target="_blank">
    </head>
    <body>

    <table border="0" cellspacing="1" bgcolor="blue"  width="500px" height="200px">
        <caption>课程表</caption>
        <tr bgcolor="white" align="center">
            <td>项目</td>
            <td colspan="6">上课</td>
            <td align="center">休息</td>
        </tr>
        <tr bgcolor="white" align="center">
            <td>星期</td>
            <td>星期一</td>
            <td>星期二</td>
            <td>星期三</td>
            <td>星期四</td>
            <td>星期五</td>
            <td>星期六</td>
            <td>星期日</td>
        </tr>
        <tr bgcolor="white" align="center">
            <td rowspan="4">上午</td>
            <td>语文</td>
            <td>数学</td>
            <td>英语</td>
            <td>英语</td>
            <td>物理</td>
            <td>计算机</td>
            <td rowspan="4">休息</td>
        </tr>
        <tr bgcolor="white" align="center">
            <td>数学</td>
            <td>数学</td>
            <td>地理</td>
            <td>历史</td>
            <td>化学</td>
            <td>计算机</td>
        </tr>
        <tr bgcolor="white" align="center">
            <td>化学</td>
            <td>语文</td>
            <td>体育</td>
            <td>计算机</td>
            <td>英语</td>
            <td>计算机</td>
        </tr>
        <tr bgcolor="white" align="center">
            <td>语文</td>
            <td>数学</td>
            <td>英语</td>
            <td>英语</td>
            <td>物理</td>
            <td>计算机</td>
        </tr>
        <tr bgcolor="white" align="center">
            <td rowspan="2">下午</td>
            <td>数学</td>
            <td>数学</td>
            <td>地理</td>
            <td>历史</td>
            <td>化学</td>
            <td>计算机</td>
            <td rowspan="2">休息</td>
        </tr>
        <tr bgcolor="white" align="center">
            <td>数学</td>
            <td>数学</td>
            <td>地理</td>
            <td>历史</td>
            <td>化学</td>
            <td>计算机</td>
        </tr>
    </table>

    </body>
</html>
```

### 11. form表单标签

语义：标记表单

什么是表单？

+ 表单就是专门用来接收输入或采集用户信息的

表单格式：

```html
   <form>
        <表单元素>
    </form>
       
       
补充：
在form内还可以添加一种标签
<fieldset>添加边框
    <legend>注册页面</legend>
    表单控件......
</fieldset>
```

#### 表单作用：

+ 表单用于向服务器传输数据，从而实现用户与Web服务器的交互

+ 表单能够包含input系列标签，比如文本字段、复选框、单选框、提交按钮等等。

+ 表单还可以包含textarea、select、fieldset和 label标签。

#### 表单属性：

|      属性      |                            描述                            |
| :------------: | :--------------------------------------------------------: |
| accept-charset |    规定在被提交表单中使用的字符集（默认：页面字符集）。    |
|     action     |       规定向何处提交表单的地址（URL）（提交页面）。        |
|  autocomplete  |         规定浏览器应该自动完成表单（默认：开启）。         |
|    enctype     |        规定被提交数据的编码（默认：url-encoded）。         |
|     method     |      规定在提交表单时所用的 HTTP 方法（默认：GET）。       |
|      name      | 规定识别表单的名称（对于 DOM 使用：document.forms.name）。 |
|   novalidate   |                   规定浏览器不验证表单。                   |
|     target     |       规定 action 属性中地址的目标（默认：_self）。        |

#### 表单标签图解

![form表单标签](https://billy.taoxiaoxin.club/md/2023/05/647616a9922ee45c40e47c96.png)

[链接：https://www.processon.com/view/link/5aeea789e4b084d6e4bf6911 ](https://www.processon.com/view/link/5aeea789e4b084d6e4bf6911)

#### 表单元素

基本概念：

+ HTML表单是HTML元素中较为复杂的部分，表单往往和脚本、动态页面、数据处理等功能相结合，因此它是制作动态网站很重要的内容。
+ 表单一般用来收集用户的输入信息

表单的工作原理：

+ 访问者在浏览有表单的网页时，可填写必需的信息，然后按某个按钮提交。这些信息通过Internet传送到服务器上。

+ 服务器上专门的程序对这些数据进行处理，如果有错误会返回错误信息，并要求纠正错误。当数据完整无误后，服务器反馈一个输入完成的信息。

```html
<form action="http://127.0.0.1:5000/index/" method="post" enctype="multipart/form-data">
    <p>
        <!--label第一种写法: 直接将input框写在label内-->
        <label for="d1">
            <!--username: <input type="text" id="d1" disabled>-->
            <!--username: <input type="text" id="d1" value="默认值">-->
            <!--username: <input type="text" id="d1" value="默认值" readonly>-->
            username: <input type="text" id="d1" placeholder="请输入用户名" name="username" value="root">
            username: <input type="text" id="d1" placeholder="请输入用户名" name="username" value="root" disabled>
        </label>
    </p>
    <p>
        <!--label第二种写法: 通过id链接即可 无序嵌套(补充: input不和label关联也可以)-->
        <label for="d2">password:</label>
        <!--<input type="text" id="d2">-->
        <input type="password" id="d2" name="password" placeholder="请输入用户密码" value="123">
    </p>
    <p>
        <!--你看不见我: <input type="hidden">-->
        <input type="hidden" name="usernamename" value="骗子账户">
    </p>
    <p>
        <label for="d3">
            birthday: <input type="date" id="d3" name="birthday">
        </label>
    </p>
    <p>
        gender:
        <label for="d4">
            <input id="d4" type="radio" name="gender" value="male" checked> 男
        </label>
        <input type="radio" name="gender" value="female"> 女
        <input type="radio" name="gender" value="other"> 其他
    </p>
    <p>
        hobby:
        <input type="checkbox" name="hobby" value="read"> read
        <input type="checkbox" name="hobby" value="DBJ" checked> DBJ
        <input type="checkbox" name="hobby" value="JBD" checked> JBD
        <input type="checkbox" name="hobby" value="HeCha"> HeCha
    </p>
    <p>
        province:
        <select name="province" id="">
            <option value="Shanghai">上海</option>
            <option value="Beijing" selected>北京</option>
            <option value="Shenzhen">深圳</option>
        </select>
    </p>        
    <p>
        前女友:
        <!--<select name="ex-girlfriend" id="">-->
        <select name="ex-girlfriend" id="" multiple>
            <option value="xxx">xxx</option>
            <option value="yyy" selected>yyy</option>
            <option value="uuu" selected>uuu</option>
        </select>
    </p>
    <p>
        province1:
        <select name="province1" id="">
            <optgroup label="上海">
                <option value="Pudong">浦东</option>
                <option value="Puxi" selected>浦西</option>
                <option value="PuNa" selected>浦南</option>
            </optgroup>
            <optgroup label="北京">
                <option value="Tian An Men">天安门</option>
                <option value="The gates">地安门</option>
                <option value="When the door">中安门</option>
            </optgroup>
            <optgroup label="深圳">
                <option value="Foxconn">富士康</option>
                <option value="Fuji bad">富士坏</option>
                <option value="Fuji in">富士中</option>
            </optgroup>
        </select>
    </p>
    <p>
        文件: <input type="file" multiple name="myfile">
        <!--文件: <input type="file" name="myfile">-->
    </p>
    <p>
        <!--12345678901234567890-->
        <!--自我介绍: <textarea name="self-introduction" id="" cols="30" rows="10"></textarea>-->
        自我介绍: <textarea name="info" id="" cols="30" rows="10" maxlength="20" placeholder="请简单的介绍自己不多余20个字!"></textarea>
    </p>
    <p>
        <!--当你没有指定按钮的value属性的值, 不同的浏览器打开之后可能宣染的按钮展示的文本内容不一致-->
        <input type="submit" value="注 册">
        <input type="button" value="按 钮">
        <input type="reset" value="重 置">
        <button>按 钮</button>
    </p>

</form>
```

使用Flask验证表单提交数据

安装Flask：

```python
# 添加了环境变量以后在命令行中下载(这里使用清华源下载)
pip3 install https://pypi.tuna.tsinghua.edu.cn/simple Flask
```

~~~python
from flask import Flask
from flask import request

app = Flask(__name__)

# 当前url既可以支持get请求也可以支持post请求  如果不写默认只能支持get请求
@app.route('/index/', methods=['GET', 'POST'])
def index():
    print(request.form)   # 获取form表单提交过来的非文件数据
    '''
    ImmutableMultiDict([('username', 'root'), ('password', '123'), ...])
    '''

    print(request.files)  # 获取文件数据
    '''
    ImmutableMultiDict([('myfile', <FileStorage: '图片1.png' ('image/png')>),])
    '''

    file_obj = request.files.get("myfile")  # 通过get表单标签中的file标签中定义的name属性的值myfile获取到文件对象
    # print('file_obj.name:', file_obj.name)  # 获取`文件名`
    print('file_obj.filename:', file_obj.filename)  # 获取`文件名.后缀`
    if file_obj:
        file_obj.save(file_obj.filename)  # 保存文件对象到当前目录下(.filename以当前`文件名.后缀`保存)

    return 'OK'


app.run()
~~~

#### 总结：

```html
# 标签分类:
    块级标签: form
    行内标签: label, input, textarea, select, option, optgroup


# form种所有标签共用属性:
        name   提交到后端的key.
        value  提交到后端的value. 文本框标签可以不指定, 通过获取用户输入的内容就是value. 选选框类型等都需要自己指定value值.(特例: 如果是按钮类型就仅仅是显示按钮的展示文本)
        disabled 禁用.

        
# 文本框类型共用属性: (文本框类型包括: 普通文本,密文文本, 文本域)
    readonly 只读
    placeholder 提示占位符
    maxlength 最大输入长度

    
# 选框类型共用属性: (选框类型包括: 单选框, 多选框, 文件选框, 下拉选框单选框, 下拉多选框, 下拉选项组选框)
    文件选框多选, 下拉多选框:  multiple
    单/多选框默认选中: checked
    下拉系列默认选中:  selected(需先为select标签指定multiple属性)

# 提示: 除了按钮没有必要指定label其他标签都可以被label包裹.
    label第一种写法: 直接将input框写在label内
    label第二种写法: 通过id链接即可 无序嵌套(补充: input不和label关联也可以)

        
# input标签:
    text 普通文本
        input[type=text][name]
    password 密文文本
        input[type=password][name]
    date 提供时间日期选择
        input[type=date][name][value]

        
    提示: 当你没有指定按钮的value属性的值, 不同的浏览器打开之后可能宣染的按钮展示的文本内容不一致
    submit 提交按钮.  用来触发form表单提交数据的动作
        input[type=submit][value]
    button 普通按钮.  本身没有任何的功能, 但是它是最有用的, 学完js之后就可以给它自定义各种功能
        input[type=button][value]
    reset  重置按钮.  重置
        input[type=reset][value]

    radio 单选框  默认选中要加checked='checked', 当标签的属性名和属性值一样的时候可以简写成checked(注意: 每个单选框都需要为指定相同的name属性的值, 才能达到单选的目的)
        input[type=radio][name][value]
    checkbox 多选框
        input[type=checkbox][name][value]
    file 获取文件.
        input[type=file][name][value]

    hidden 隐藏当前input框. 只是不显示在页面中, 后台还存在着.(钓鱼网站)
        input[hidden][name][value]

        
# textarea文本域标签
    textarea[name]

    
# button 按钮标签
    button[name][value]

    
# select标签:
    下拉单选
        select[name]>option[value]
    下拉多选:
        select[name][multiple]>option[value]
    下拉选项组单选
        select[name]>(optgroup[label]>option[value])+(optgroup[label]>option[value])
    下拉选项组多选
        select[name][multiple]>(optgroup[label]>option[value])+(optgroup[label]>option[value])

!!!注意!!!: 触发form表单提交功能的按钮有2种: button input[type=submit]
```

#### input

![image-20210304210807922](https://billy.taoxiaoxin.club/md/2023/05/647616aa922ee45c4138b1a3.png)

属性说明:

- name：表单提交时的“键”，注意和id的区别
- value：表单提交时对应项的值
  - type=”button”, “reset”, “submit”时，为按钮上显示的文本年内容
  - type=”text”,”password”,”hidden”时，为输入框的初始值
  - type=”checkbox”, “radio”, “file”，为输入相关联的值
- checked：radio和checkbox默认被选中的项
- readonly：text和password设置只读
- disabled：所有input均适用

### 12. select 标签

```html
<form action="" method="post">
  <select name="city" id="city">
    <option value="1">北京</option>
    <option selected="selected" value="2">上海</option>
    <option value="3">广州</option>
    <option value="4">深圳</option>
  </select>
</form>
```

属性说明：

- multiple：布尔属性，设置后为多选，否则默认单选
- disabled：禁用
- selected：默认选中该项
- value：定义提交时的选项值

### 13.  label标签

定义： 标签为 input 元素定义标注（标记）。
说明：

1. label 元素不会向用户呈现任何特殊效果。
2. 标签的 for 属性值应当与相关元素的 id 属性值相同。

```html
<form action="">
  <label for="username">用户名</label>
  <input type="text" id="username" name="username">
</form>
```

### 14 textarea多行文本

```html
<textarea name="memo" id="memo" cols="30" rows="10">
  默认内容
</textarea>
```

属性说明：

- name：名称
- rows：行数
- cols：列数
- disabled：禁用

## 三、其他了解

###  1. 格式排版标签

- `<br/>` 换行标签，完成文字的紧凑显示。可以使用连续多个`<br/>`标签来换行
- `<hr/>` 水平分割线标签，用于段落与段落之间的分割
- `<p></p>`段落标签,里面可以加入文字,列表,表格等.
- `<pre></pre>`按原文显示标签，可以把原文件中的空格,回车,换行,tab键表现出来
- `<hn></hn>` 标题字标签，n为1-6，定义六级标题，而且会自动换行插入一个空行
- `<div></div>` 没有任何语义的标签

### 2. 文本标签

- `<em></em>` 表示强调，通常为斜体字
- `<strong></strong>` 表示强调(语气更强)，通常为粗体字
- `<del></del>` 标签定义文档中已删除的文本。
- `<ins></ins>` 标签定义已经被插入文档中的文本
- `<sub></sub>` 文字下标字体标签
- `<sup></sup>` 文字上标字体标签
- `<mark></mark>` **H5新增** 标签定义带有记号的文本 请在需要突出显示文本时使用,如搜索引擎搜索页面
- `<ruby></ruby>` **H5新增** 标签定义 ruby 注释（中文注音或字符） 在东亚使用，显示的是东亚字符的发音。
- `<rt></rt>` **H5新增** 标签定义字符（中文注音或字符）的解释或发音

```html
<!--一下文本标签  作为了解-->
<cite>    用于引证、举例、(标签定义作品（比如书籍、歌曲、电影、电视节目、绘画、雕塑等等）的标题)通常为斜体字
<dfn> 定义一个定义项目
<code> 定义计算机代码文本
<samp> 定义样式文本 标签并不经常使用。只有在要从正常的上下文中将某些短字符序列提取出来，对它们加以强调的极少情况下，才使用这个标签。
<kbd> 定义键盘文本。它表示文本是从键盘上键入的。它经常用在与计算机相关的文档或手册中。
<abbr> 定义缩写 配合title属性  (IE6以上)
<bdo>  来覆盖默认的文本方向 dir属性 值: lrt  rtl
<var> 定义变量。您可以将此标签与 <pre> 及 <code> 标签配合使用。
<small> 标签定义小型文本（和旁注）
<b>    粗体字标签 根据 HTML 5 的规范，<b> 标签应该做为最后的选择，只有在没有其他标记比较合适时才使用它。
<i>    斜体字标签 标签被用来表示科技术语、其他语种的成语俗语、想法、宇宙飞船的名字等等。
<u>    下划线字体标签 标签定义与常规文本风格不同的文本，像拼写错误的单词或者汉语中的专有名词。 请尽量避免使用 <u> 为文本加下划线，用户会把它混淆为一个超链接。
<q>  签定义一个短的引用。浏览器经常会在这种引用的周围插入引号。(小段文字)
<blockquote> 标签定义摘自另一个源的块引用。浏览器通常会对 <blockquote> 元素进行缩进。(大段文字) (块状元素)
<address>  定义地址 通常为斜体 (注意非通讯地址)  块状元素
<font>       H5已删除 字体标签，可以通过标签的属性指定文字的大小、颜色及字体等信息
<tt>       H5已删除 打字机文字
<big>       H5已删除 大型字体标签
<strike>   H5已删除 添加删除线
<acronym>  H5已删除 首字母缩写 请使用<abbr>代替
<bdi>      H5新增 标签允许您设置一段文本，使其脱离其父元素的文本方向设置。(经测试,各大浏览器都不起作用)
<mark>     H5新增 标签定义带有记号的文本 请在需要突出显示文本时使用,如搜索引擎搜索页面
<meter>    H5新增 定义预定义范围的度量
<progress> H5新增 标签标示任务的进度（进程）
<time>     H5新增 定义时间和日期 
<wbr>        H5新增    规定在文本中的何处适合添加换行符。Word Break Opportunity
```



## 参考资料

+ [百度百科：HTML标签语义化](https://baike.baidu.com/item/html%E6%A0%87%E7%AD%BE%E8%AF%AD%E4%B9%89%E5%8C%96/10959478?fr=aladdin)