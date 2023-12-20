# Vue 入门介绍

<img src="https://billy.taoxiaoxin.club/md/2023/06/647ddcc1922ee46e5cd04c5e.png" alt="image-20230605210152885" style="zoom:50%;" />

## 一、前端的发展史

**1**.`HTML`(5)、`CSS`(3)、`JavaScript`(ES5、ES6)：编写一个个的页面 -> 给后端(PHP、Python、Go、Java) -> 后端嵌入模板语法 -> 后端渲染完数据 -> 返回数据给前端 -> 在浏览器中查看

**2**.Ajax的出现 -> 后台发送异步请求，`Render`+`Ajax`混合

**3**.单用Ajax（加载数据，DOM渲染页面）：`前后端分离的雏形`

**4**.[Angular框架](https://angular.cn/)的出现（1个JS框架）：出现了“`前端工程化`”的概念（前端也是1个工程、1个项目）

**5**.[React](https://react.docschina.org/)、[Vue框架](https://cn.vuejs.org/)：当下最火的2个前端框架（`Vue`：国人喜欢用，`React`：外国人喜欢用）

**6**.移动开发（Android+IOS） + Web（Web+微信小程序+支付宝小程序） + 桌面开发（Windows桌面）：前端 -> `大前端`

**7**.一套代码在各个平台运行（**大前端**）：`谷歌Flutter（Dart语言：和Java很像）`可以运行在IOS、Android、PC端

**8**.在Vue框架的基础性上 [uni-app](https://uniapp.dcloud.io/)：**一套编码 编到10个平台**

**9**.在不久的将来 ，前端框架可能会一统天下

[详细了解,点击进入前端开发 20 年变迁史](https://zhuanlan.zhihu.com/p/337276087?utm_source=wechat_session&utm_medium=social&utm_oi=41967790587904)



## 二、Vue介绍

### 1.Vue介绍

[Vue](https://cn.vuejs.org/) (读音 `/vjuː/`，类似于 **view**) 是一套用于构建用户界面的**JavaScript 渐进式框架**。 框架。 它基于标准 HTML、 CSS和 Javascript JavaScript 构建， 并提供了⼀套声明式的、 并提供了一套声明式的、 组件化的编程模型， 组件化的编程模型， 帮助你⾼效地开发⽤户界⾯。 帮助你高效地开发用户界面。无论是简单还 ⽆论是简单还 是复杂的界面，Vue 都可以胜任。不仅易于上手，还便于与第三方库或既有项目整合,可以在某个文件单独使用，只用一部分，也可以整个项目中都使用。

官网文档：https://cn.vuejs.org/

### 2.jQuery和vue的区别

+ `jQuery`的定位是获取元素和完成特效
+ `vue`的定位是方便操作和控制数据和完成特效。

### 3.Vue 特性

#### 易用

+ 通过 HTML、CSS、JavaScript构建应用
+ HTML 模板 + JSON 数据，即可创建一个 Vue 实例，就这么简单

#### 灵活

+ 不断繁荣的生态系统，可以在一个库和一套完整框架之间自如伸缩。
+ **数据驱动：** 自动追踪依赖的模板表达式和计算属性。

#### 高效

- **轻量：** ~24kb min+gzip，无依赖。
- **快速：** 精确有效的异步批量 DOM 更新。
- 最省心的优化

## 三、Vue组件化开发与单页面开发

#### 组件化开发

+ 页面上小到一个按钮都可以是一个单独的文件.vue，这些小组件直接可以像乐高积木一样通过互相引用而组装起来
+ 类似于DTL中的`include`，**每一个组件**的内容都可以被**替换**和**复用**

![img](https://billy.taoxiaoxin.club/md/2023/06/647ddab7922ee46e2c8dce28.jpeg)

#### 单页面开发

只需要1个页面，结合组件化开发来替换页面中的内容

页面的切换只是组件的替换，页面还是只有1个`index.html`

## 四、安装

#### 1.版本介绍

1. Vue.js 1.x(**几乎很少使用**):
   - Vue.js 1.x 是Vue.js框架最初的版本。
   - 它引入了Vue.js的核心特性，如Vue实例、指令、模板语法和基本的数据绑定。
   - Vue.js 1.x 使用 Object.defineProperty 方法实现响应式，但在性能和深层对象响应性方面存在一些限制。
   - 尽管被许多开发者使用，但Vue.js 1.x已被认为过时，并且不再进行积极维护，不建议在新项目中使用。
2. Vue.js 2.x(**广泛使用**):
   - Vue.js 2.x 是目前最广泛使用且稳定的版本。
   - 它解决了1.x版本的限制，并带来了重大改进和新功能。
   - 响应式系统被重写以提高性能，使用了虚拟DOM diffing算法。
   - Vue.js 2.x 引入了单文件组件（SFC）格式，允许开发者在单个文件中编写模板、脚本和样式。
   - 它添加了计算属性、观察者、自定义指令以及对组件和生命周期钩子的更好支持。
   - Vue.js的官方路由库Vue Router和状态管理库Vuex也在2.x版本中引入，并被广泛采用。
   - Vue.js 2.x 目前仍然提供关键错误修复的支持，但不再为新功能进行积极维护。
3. Vue.js 3.x(**Vue2 在向Vue 3迁移** ):
   - Vue.js 3.x 是Vue.js的最新主要版本，带来了重大改进和变化。
   - 它提供了更好的性能、更小的打包体积、增强的TypeScript支持和更好的开发体验。
   - Vue.js 3.x 引入了组合式API（Composition API），这是一种更灵活、可重用的组件逻辑编写方式。
   - 响应式系统完全使用了JavaScript Proxy进行重写，提供了更好的性能和更强大的响应式能力。
   - 打包体积缩小，并且支持Tree-shaking，使其在生产环境中更加高效。
   - Vue.js 3.x 建议在新项目中使用，但由于与2.x版本存在一些不兼容的变化，迁移现有项目可能需要进行调整。

需要注意的是，虽然Vue.js 1.x和2.x仍然在一些老旧项目中使用，但社区的重点和未来的开发工作都集中在Vue.js 3.x上。

#### 2.CDN方式引入

对于学习，你可以这样使用最新版本：

```js
<script src="https://unpkg.com/vue@3/dist/vue.global.js"></script>
```

通过 CDN 使⽤vue 时， 不涉及“构建步骤”。这使得设置更加简单， 并且可以⽤于增强静态的 HTML 或与后端框架 集成。但是，你将无法使用单文件组件 (SFC) 语法。

#### 3.本地导入

打开地址:https://unpkg.com/vue@3/dist/vue.global.js

然后直接鼠标右键,存储为将文件名命名为vue.js即可.

<img src="https://billy.taoxiaoxin.club/md/2023/06/647deed7922ee46ffab43f46.png" alt="image-20230605221903233" style="zoom: 50%;" />

#### 4.NPM安装

```js
# 最新稳定版本
$ npm install vue
# 最新稳定 CSP 兼容版本
$ npm install vue@csp
# 开发版本（直接从 GitHub 安装）
$ npm install vuejs/vue#dev
```

#### 5.Bower 安装

```js
# 最新稳定版本
$ bower install vue
```

## 五、nodejs介绍

#### 解释型的语言是需要解释器的

js就是一门解释型语言，只不过js解释器被集成到了浏览器中

所以，在浏览器的Console中输入命令，就和在cmd中输入python后，进入交互式环境一样

#### nodejs：一门后端语言

把chrome的v8引擎（解释器），安装到操作系统之上，写javascript的代码

优势：前端工程师，不用学后端语言，只会js，就可以写后端了

```python
npm：类似于python的pip3，用来安装第三方包

## 前端开发的ide
	-webstorm
    -vscode
    -hbuilder
    -sublinetext
## 咱们用pycharm
	-webstorm和pycharm是一家，只需要装vue插件
```

## 六、简单使用

#### 使用步骤：

1. 先引入`vue`核心文件
2. 对`vue`的核心对象`vm`进行实例化:` vue.js`的代码开始于一个`Vue`对象。所以每次操作数据都要声明`Vue`对象开始
3. 在`el`属性对象的标签中，填写正确的`vue`语法展示或者控制数据

#### 简单使用

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <script src="vue.js"></script>
</head>
  
<body>
<div id="box">
    {{ 10+20 }}
    <span>{{20+30}}</span>
</div>
{{10+20}}
<div id="box2">
    {{10+20}}
</div>
<div id="app">
  <!-- {{ message }} 表示把vue对象里面data属性中的对应数据输出到页面中 -->
  <!-- 在双标签中显示数据要通过{{  }}来完成 -->
    {{message}}
</div>
<script>
    Vue.createApp().mount("#box")  
    Vue.createApp().mount("#box2")
    var app = Vue.createApp({
        data(){   // 保存一些前端使用的数据, 这里的data是将要展示到HTML标签元素中的数据
            return {
                message:"我的第一个Vue应用" //状态
            }
        }
    }).mount("#app") // 通过mount("#app")将Vue应用实例与"app"元素关联起来
</script>
</body>
</html>
```

1. 首先，通过`<script src="vue.js"></script>`引入了Vue框架的JavaScript文件。
2. 在HTML结构中，有三个`<div>`标签，分别具有`id`属性为"box"、"box2"和"app"，用于将Vue应用绑定到不同的DOM元素上。
3. 在JavaScript代码块中，首先使用`Vue.createApp()`方法创建了两个Vue应用实例，分别与"box"和"box2"元素相关联，并通过`mount()`方法将Vue应用挂载到对应的DOM元素上。
4. 接下来，使用`var app = Vue.createApp({})`创建了一个Vue应用实例，并传入了一个配置对象作为参数。在配置对象中，使用`data()`方法定义了一个`data`函数，返回一个包含应用需要展示的数据的对象。这里的`data`函数中定义了一个属性`message`，其初始值为"我的第一个Vue应用"。
5. 最后，通过`mount("#app")`将Vue应用实例与"app"元素关联起来，并将数据渲染到对应的DOM元素中。`{{ message }}`是Vue的模板语法，用于在HTML中插入数据。

#### vue的基本使用三个注意事项：

1. 每个Vue实例应该具有唯一的变量名： 虽然一个HTML页面可以包含多个Vue实例对象，但强烈建议为每个实例选择唯一的变量名。这样可以避免变量名冲突和混淆，使代码更易于理解和维护。确保为每个Vue实例选择一个有意义且描述性的变量名，以便清晰地区分它们的功能和职责。
2. 注意JavaScript的大小写敏感性： JavaScript是大小写敏感的语言，这意味着变量名、函数名和关键字的大小写必须与其在其他地方的使用保持一致。在编写Vue代码时，确保所有变量和函数名的大小写一致，并注意正确的语法。遵循一致的命名约定有助于减少错误和调试困难。
3. 将实例化Vue对象的代码放在HTML文件的末尾： 在编写Vue应用时，建议将实例化Vue对象的代码放在HTML文件的末尾。这样做的原因是，Vue实例化之前，HTML文件中的元素可能尚未完全加载和渲染。通过将Vue对象实例化的代码放在文件的末尾，可以确保在Vue实例化之前，所有需要的HTML元素都已加载完毕，避免因为元素未找到而引发的错误。
