# Vue 组件化开发

## 一、组件化开发介绍

### 1.组件是什么？组件是什么？

组件，或者说组件化，是一种设计思想，广泛应用于现代前端开发中，尤其是在 Vue.js、React.js、Angular.js 等前端框架中，这种设计思想起到了关键的作用。

**组件的定义**:组件是一个独立的、可重用的代码模块，它可以包含 HTML、CSS、JavaScript，以及这些语言的各种组合。**组件的主要目标是提高应用的可重用性，逻辑清晰和代码整洁。**它可以看作是一种**封装**，将一部分相对独立的功能代码封装起来，使其能够以模块化的形式重用。

![img](https://billy.taoxiaoxin.club/md/2023/07/64b537f9922ee4a6a3fb6eb7.png)

**组件的作用与优点 :**

- **重用性**：将公共代码或经常使用的功能抽象为组件，可以在不同的地方重复使用，节省开发时间。
- **解耦**：每个组件都有其独立的功能和责任，便于开发者对单个组件进行开发和维护，减少代码间的耦合度。
- **可维护性**：通过组件化，代码结构更清晰，模块化程度更高，增强了代码的可读性和可维护性。
- **协同开发**：组件化可以更好地支持团队协作开发，各成员负责不同组件的开发，提高开发效率。

在 Vue.js 中，组件化开发是其核心理念之一，它让开发者能够通过构建可重用的组件来构建大型应用，每个组件都可以拥有自己的状态、标记和行为。

例如：有一个轮播图，可以在很多页面中使用，一个轮播有js，css，html,
组件可以把共有代码的js，css，html放到一起，有逻辑，有样式，有html,提高代码的可维护性,解耦合性,从而提高代码的重用性。

![img](https://billy.taoxiaoxin.club/md/2023/07/64b537f9922ee4a6a4775716.png)

#### 工程化开发之后：

1个组件 就是1个`xx.vue`

### 2.什么是组件化？

#### 人面对复杂问题的处理方式：

+ 任何一个人处理信息的逻辑能力都是有限的
+ 所以，当面对一个非常复杂的问题时，我们不太可能一次性搞定一大堆的内容。
+ 但是，我们人有一种天生的能力，就是将问题进行拆解。
+ 如果将一个复杂的问题，拆分成很多个可以处理的小问题，再将其放在整体当中，你会发现大的问题也会迎刃而解。

![image-20210429101551431](https://billy.taoxiaoxin.club/md/2023/07/64b537f9922ee4a6a52cef2e.png)

#### 组件化也是类似的思想：

+ 如果我们将一个页面中所有的处理逻辑全部放在一起，处理起来就会变得非常复杂，而且不利于后续的管理以及扩展。

  

+ 但如果，我们讲一个页面拆分成一个个小的功能块，每个功能块完成属于自己这部分独立的功能，那么之后整个页面的管理和维护就变得非常容易了。

![image-20210429101822546](https://billy.taoxiaoxin.club/md/2023/07/64b537f9922ee4a6a692cf9e.png)

#### 组件化定义：

+ 我们将一个完整的页面分成很多个组件。
+ 每个组件都用于实现页面的一个功能块。
+ 而每一个组件又可以进行细分。

### 3.Vue组件化思想

#### 组件化是Vue.js中的重要思想

+ 它提供了一种抽象，让我们可以开发出一个个独立可复用的小组件来构造我们的应用。
+ **任何的应用都会被抽象成一颗组件树**。

![image-20210429102232079](https://billy.taoxiaoxin.club/md/2023/07/64b537f9922ee4a6a77342f1.png)

#### 组件化思想的应用：

+ 有了组件化的思想，我们在之后的开发中就要充分的利用它,目的是复用。
+ 尽可能的将页面拆分成一个个小的、可复用的组件。
+ 这样让我们的代码更加方便组织和管理，并且扩展性也更强。
  

所以，组件是Vue开发中，非常重要的一个篇章，要认真学习。

## 二、组件编写方式

### 2.1 组件注册与使用的三个步骤：

#### 2.1.1 创建组件构造器

在 Vue 中，我们可以直接使用一个对象来定义组件：

```js
let MyComponent = {
  template: `<div>Hello, Vue!</div>`,
  data() {
    return {
      message: 'Hello, Vue!'
    };
  },
  methods: {
    sayHello() {
      alert(this.message);
    }
  }
};

```

#### 2.1.2 注册组件

在 Vue 3 中，我们可以选择全局或局部注册组件。

+ 全局注册：

  ```js
  const app = Vue.createApp({});
  
  app.component('my-component', MyComponent);
  
  ```

  在这里，我们使用 `app.component()` 方法注册了一个全局组件。第一个参数是组件的名字（在 HTML 中使用），第二个参数是我们之前创建的组件对象。

+ 局部注册：

  在创建另一个组件时，我们可以在这个新组件的 `components` 选项中注册之前创建的组件：

  ```vue
  let AnotherComponent = {
    components: {
      'my-component': MyComponent
    },
    template: `<div><my-component></my-component></div>`
  };
  
  ```

  在这里，我们创建了一个新的组件 `AnotherComponent`，并在这个组件中局部注册了 `MyComponent`。

#### 2.1.3 使用组件

注册组件之后，就可以在模板中像使用普通 HTML 元素一样使用组件了。使用的时候，组件名就是之前注册时用的那个名字：

```vue
<template>
  <my-component></my-component>
</template>
```

### 2.2 基本使用

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Vue Component Example</title>
  <script src="https://unpkg.com/vue@next"></script>
</head>
<body>
  <div id="app">
    <my-component></my-component>
  </div>

  <script>
    // 创建组件
    const MyComponent = {
      data() {
        return {
          message: 'Hello, Vue!'
        }
      },
      template: `
        <div>
          <h2>{{ message }}</h2>
          <button @click="sayHello">Say Hello</button>
        </div>
      `,
      methods: {
        sayHello() {
          alert(this.message);
        }
      }
    }

    // 创建 Vue 实例并注册组件
    const app = Vue.createApp({});
    app.component('my-component', MyComponent);
    app.mount('#app');
  </script>
</body>
</html>

```

### 2.3 全局组件和局部组件

#### 2.3.1 全局组件

全局组件在整个 Vue 应用程序中是可用的，一旦它被注册，我们就可以在任何地方使用它，包括在其他组件的模板中。在 Vue 3 中，我们使用 `app.component()` 方法在创建的 Vue 应用实例 `app` 中注册全局组件。

```js
const app = Vue.createApp({});

app.component('global-component', {
  template: `<div>I am a global component!</div>`,
});

app.mount('#app');
```

然后，你可以在你的 HTML 代码中直接使用全局组件：

```vue
<div id="app">
  <global-component></global-component>
</div>
```

#### 2.3.2 局部组件

局部组件只能在它被注册的组件中使用，也就是说，它的作用范围被限定在注册它的组件中。在 Vue 3 中，我们可以在一个组件的 `components` 选项中注册局部组件：

```js
const ChildComponent = {
  template: `<div>I am a local component!</div>`
};

const app = Vue.createApp({
  components: {
    'child-component': ChildComponent
  },
  template: `
    <div>
      <child-component></child-component>
    </div>
  `,
});
app.mount('#app');
```

在这个例子中，`ChildComponent` 只能在我们刚刚创建的 `app` 应用中使用，因为它是在这个应用中注册为局部组件的。如果你在其他地方使用这个组件，Vue 将不会识别它。

完整代码如下:

```vue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Vue Component Example</title>
  <script src="https://unpkg.com/vue@next"></script>
</head>
<body>
  <div id="app">
    <global-component></global-component>
    <parent-component></parent-component>
    <local-component></local-component>
  </div>

  <script>
    const app = Vue.createApp({});
    
    // 创建全局组件
    app.component('global-component', {
      template: `<div>I am a global component!</div>`
    });

    // 创建局部组件
    const LocalComponent = {
      template: `<div>I am a local component!</div>`
    };

    // 创建包含局部组件的组件
    app.component('parent-component', {
      components: {
        'local-component': LocalComponent
      },
      template: `
        <div>
          I am a parent component!
          <local-component></local-component>
        </div>
      `,
    });

    // 挂载 Vue 应用
    app.mount('#app');
  </script>
</body>
</html>
```

在这个例子中，我们在 `parent-component` 外部试图使用 `local-component`，这将导致一个错误，因为 `local-component` 只在 `parent-component` 中被注册了。打开你的开发者工具，你会在控制台中看到类似这样的错误消息：“[Vue warn]: Failed to resolve component: local-component”，表示 Vue 无法解析 `local-component` 组件。这就是试图在其注册范围外使用局部组件时会发生的情况。

![image-20230720000551656](https://billy.taoxiaoxin.club/md/2023/07/64b809e0922ee4cb6121413e.png)

### 5.父组件与子组件

在前面我们看到了组件树：

+ 组件和组件之间存在层级关系
+ 而其中一种非常重要的关系就是父子组件的关系

我们来看通过代码如何组成的这种层级关系：

![image-20210430223221103](https://billy.taoxiaoxin.club/md/2023/07/64b537fa922ee4a6abe17b84.png)

父子组件错误用法：以子标签的形式在Vue实例中使用

+ 因为当子组件注册到父组件的components时，Vue会编译好父组件的模块
+ 该模板的内容已经决定了父组件将要渲染的HTML（相当于父组件中已经有了子组件中的内容了）
+ `<child-cpn></child-cpn>`是只能在父组件中被识别的。
+ 类似这种用法，`<child-cpn></child-cpn>`是会被浏览器忽略的。

```python
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Title</title>
</head>
<body>

<div id="app">
  <cpn2></cpn2>
  <!--<cpn1></cpn1>-->
</div>

<script src="../js/vue.js"></script>
<script>
  // 1.创建第一个组件构造器(子组件)
  const cpnC1 = Vue.extend({
    template: `
      <div>
        <h2>我是标题1</h2>
        <p>我是内容, 哈哈哈哈</p>
      </div>
    `
  })


  // 2.创建第二个组件构造器(父组件)
  const cpnC2 = Vue.extend({
    template: `
      <div>
        <h2>我是标题2</h2>
        <p>我是内容, 呵呵呵呵</p>
        <cpn1></cpn1>
      </div>
    `,
    components: {
      cpn1: cpnC1
    }
  })

  // 根组件
  const app = new Vue({
    el: '#app',
    data: {
      message: '你好啊'
    },
    components: {
      cpn2: cpnC2
    }
  })
</script>

</body>
</html>
```

### 6.注册组件语法糖

在上面注册组件的方式，可能会有些繁琐。

+ Vue为了简化这个过程，提供了注册的语法糖。
+ 主要是省去了调用Vue.extend()的步骤，而是可以直接使用一个对象来代替。

语法糖注册全局组件和局部组件：

![image-20210430224502244](https://billy.taoxiaoxin.club/md/2023/07/64b537fa922ee4a6aca7d210.png)

```python
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Title</title>
</head>
<body>

<div id="app">
  <cpn1></cpn1>
  <cpn2></cpn2>
</div>

<script src="../js/vue.js"></script>
<script>
  // 1.全局组件注册的语法糖
  // 1.创建组件构造器
  // const cpn1 = Vue.extend()

  // 2.注册组件
  Vue.component('cpn1', {
    template: `
      <div>
        <h2>我是标题1</h2>
        <p>我是内容, 哈哈哈哈</p>
      </div>
    `
  })

  // 2.注册局部组件的语法糖
  const app = new Vue({
    el: '#app',
    data: {
      message: '你好啊'
    },
    components: {
      'cpn2': {
        template: `
          <div>
            <h2>我是标题2</h2>
            <p>我是内容, 呵呵呵</p>
          </div>
    `
      }
    }
  })
</script>

</body>
</html>
```

### 7.模板的分离写法

刚才，我们通过语法糖简化了Vue组件的注册过程，另外还有一个地方的写法比较麻烦，就是template模块写法。

如果我们能将其中的HTML分离出来写，然后挂载到对应的组件上，必然结构会变得非常清晰。

Vue提供了两种方案来定义HTML模块内容：

+ 使用`<script>`标签

+ 使用`<template>`标签

![image-20210430225437516](https://billy.taoxiaoxin.club/md/2023/07/64b537fa922ee4a6ae33cfad.png)

![image-20210430225445236](https://billy.taoxiaoxin.club/md/2023/07/64b537fa922ee4a6afeece4a.png)

```python
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Title</title>
</head>
<body>

<div id="app">
  <cpn></cpn>
  <cpn></cpn>
  <cpn></cpn>
</div>

<!--1.script标签, 注意:类型必须是text/x-template-->
<!--<script type="text/x-template" id="cpn">-->
<!--<div>-->
  <!--<h2>我是标题</h2>-->
  <!--<p>我是内容,哈哈哈</p>-->
<!--</div>-->
<!--</script>-->

<!--2.template标签-->
<template id="cpn">
  <div>
    <h2>我是标题</h2>
    <p>我是内容,呵呵呵</p>
  </div>
</template>

<script src="../js/vue.js"></script>
<script>

  // 1.注册一个全局组件
  Vue.component('cpn', {
    template: '#cpn'
  })

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

### 8.父子组件的`data`是无法共享的

- 这一点就像**Docker的容器**一样，是**相互隔离**的
- 就算父子的data中数据相同，拥有相同的方法，也是**互不影响**的

组件是一个单独功能模块的封装：

这个模块有属于自己的HTML模板，也应该有属性自己的数据data。

组件中的数据是保存在哪里呢？顶层的Vue实例中吗？

我们先来测试一下，组件中能不能直接访问Vue实例中的data

![image-20210430230137932](https://billy.taoxiaoxin.club/md/2023/07/64b537fa922ee4a6b00666e4.png)

我们发现不能访问，而且即使可以访问，如果将所有的数据都放在Vue实例中，Vue实例就会变的非常臃肿。

+ 结论：Vue组件应该有自己保存数据的地方。

### 9组件可以有data、methods、computed....，但是 `data` 必须是一个`函数`

Vue实例：data是1个键值对，用来存放属性的

```js
var vm = new Vue({
    el: '#box',
    data: {
        isShow: true
    }
})
```

组件：data是1个函数，需要有返回值(`return`)

```js
Vue.component('global', {
    template: `
        <div>
            <div style="background: rgba(255,104,104,0.7); padding: 5px;" @click="handleClick">我是头部组件</div>
            <div v-if="isShow">显示消失</div>
        </div>
`,
    methods: {
        handleClick() {
            console.log('我被点击了')
            this.isShow = !this.isShow
        }
    },
    data() {
        return {
            isShow: true
        }
    }
})
```



## 三、组件通信

### 1.父传子

- 在全局组件中自定义属性：`<global :myname="name" :myage="19"></global>`
- 在组件中获取：`{{myname}}`

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>组件</title>
    <script src="https://cdn.bootcdn.net/ajax/libs/vue/2.6.12/vue.min.js"></script>
</head>
<body>

<div id="box">
    <!-- myName是自定义属性 -->
    <global myname="name" myage="18"></global>
    <global :myname="name" :myage="19"></global>
    <global :myname="'Ben'" :myage="20"></global>
</div>

</body>
<script>
    // 创建1个组件对象（全局组件/子组件）
    Vue.component('global', {
        template: `
            <div>
                <div style="background: rgba(255,104,104,0.7); padding: 5px;">全局组件/子组件</div>
                {{myname}}
                {{myage}}
            </div>
        `,
        props: ['myname', 'myage']
    })
    // 父组件
    let vm = new Vue({
        el: '#box',
        data: {
            name: 'darker'
        },
    })
</script>
</html>
```

[![img](https://billy.taoxiaoxin.club/md/2023/07/64b537fa922ee4a6b1a6deb6.png)](https://gitee.com/xuexianqi/img/raw/master/img/image-20201217104106154.png)

#### 属性验证

- 限制父传子的变量类型

```js
props: {
    myname: String,
    isshow: Boolean
}
```

- 父传子时候注意以下区别

```html
<global :myname="name" :is_show="'false'"></global>
<global :myname="name" :is_show="false"></global>
<global :myname="name" :is_show="is_show"></global>
```

- 实例

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>组件</title>
    <script src="https://cdn.bootcdn.net/ajax/libs/vue/2.6.12/vue.min.js"></script>
</head>
<body>

<div id="box">
    <!-- myName是自定义属性 -->
    <!--    <global :myname="name" :myage="19" :isshow="'false'"></global>-->
    <global :my_name="name" :is_show="is_show"></global>
    <global :my_name="name" :is_show="false"></global>
</div>

</body>
<script>
    // 创建1个组件对象（全局组件/子组件）
    Vue.component('global', {
        template: `
            <div>
                <div style="background: rgba(255,104,104,0.7); padding: 5px;">我是子组件:{{is_show}}</div>
                <span>{{my_name}}</span>
            </div>
        `,
        props: {
            my_name: String,
            is_show: Boolean
        }
    })
    // 父组件
    let vm = new Vue({
        el: '#box',
        data: {
            name: 'darker',
            is_show: true
        },
    })
</script>
</html>
```

### 2.子传父（通过事件）

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>子传父</title>
    <script src="https://cdn.bootcdn.net/ajax/libs/vue/2.6.12/vue.min.js"></script>
</head>
<body>

<div id="box">
    <global @my_event="handleClick($event)"></global>
</div>

</body>
<script>
    // 创建1个组件对象（全局组件/子组件）
    Vue.component('global', {
        template: `
            <div>
                <div style="background: rgba(255,104,104,0.7); padding: 5px;">全局组件/子组件</div>
                <button @click="handleNav">点我</button>
            </div>
        `,
        data() {
            return {
                name: 'Darker'
            }
        },
        methods: {
            handleNav() {
                console.log('我是子组件的函数')
                this.$emit('my_event', 666, 777, this.name)
            }
        }
    })
    // 父组件
    let vm = new Vue({
        el: '#box',
        data: {},
        methods: {
            handleClick(a,b,c) {
                console.log('我是父组件的函数')
                console.log(a)
                console.log(b)
                console.log(c)
            }
        }
    })
</script>
</html>
```

[![img](https://billy.taoxiaoxin.club/md/2023/07/64b537fa922ee4a6b2cb8973.png)](https://gitee.com/xuexianqi/img/raw/master/img/image-20201217161446678.png)

### 3.子传父（控制子组件的显示和隐藏）

点击子组件，就会触发父组件的某个函数执行

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>子传父</title>
    <script src="https://cdn.bootcdn.net/ajax/libs/vue/2.6.12/vue.min.js"></script>
</head>
<body>

<div id="box">
    <global @my_event="handleClick($event)"></global>
</div>

</body>
<script>
    // 创建1个组件对象（全局组件/子组件）
    Vue.component('global', {
        template: `
            <div>
                <div style="background: rgba(255,104,104,0.7); padding: 5px;">全局组件/子组件</div>
                <button @click="handleNav">点我</button>
            </div>
        `,
        data() {
            return {
                name: 'Darker'
            }
        },
        methods: {
            handleNav() {
                console.log('我是子组件的函数')
                this.$emit('my_event', 666, 777, this.name)
            }
        }
    })
    // 父组件
    let vm = new Vue({
        el: '#box',
        data: {},
        methods: {
            handleClick(a,b,c) {
                console.log('我是父组件的函数')
                console.log(a)
                console.log(b)
                console.log(c)
            }
        }
    })
</script>
</html>
```

[![img](https://billy.taoxiaoxin.club/md/2023/07/64b537fa922ee4a6b3f9c9f4.png)](https://gitee.com/xuexianqi/img/raw/master/img/image-20201218085953397.png)

#### 小案例

- 子组件有1个按钮 和 1个输入框，子组件输入完内容后，数据在父组件中展示

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>子传父 小案例</title>
    <script src="https://cdn.bootcdn.net/ajax/libs/vue/2.6.12/vue.min.js"></script>
</head>
<body>

<div id="box">
    <global @my_event="handleShow($event)"></global>
    <br>
    <div>父组件接收到的数据：{{name}}</div>
</div>

</body>
<script>
    // 创建1个组件对象（全局组件/子组件）
    Vue.component('global', {
        template: `
            <div>
                <input type="text" v-model="myText">
                <button @click="handleClick">点我传数据</button>
            </div>
        `,
        data() {
            return {
                myText: ''
            }
        },
        methods: {
            handleClick() {
                this.$emit('my_event', this.myText)
            }
        }
    })
    // 父组件
    let vm = new Vue({
        el: '#box',
        data: {
            name: ''
        },
        methods: {
            handleShow(a) {
                this.name = a
            }
        }
    })
</script>
</html>
```

[![img](https://billy.taoxiaoxin.club/md/2023/07/64b537fa922ee4a6b445b9b2.gif)](https://gitee.com/xuexianqi/img/raw/master/img/11 szf01.gif)

### 4.ref属性（也可以实现组件间通信：子和父都可以实现通信）

- ref放在`标签`上，拿到的是`原生的DOM节点`
- ref放在`组件`上，拿到的是`组件对象` ，对象中的数据、函数 都可以直接使用
- 通过这种方式实现子传父（this.$refs.mychild.text）
- 通过这种方式实现父传子（调用子组件方法传参数）

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>子传父</title>
    <script src="https://cdn.bootcdn.net/ajax/libs/vue/2.6.12/vue.min.js"></script>
</head>
<body>

<div id="box">
    <input type="text" ref="myRef">
    <button @click="handleButton">点我</button>
</div>

</body>
<script>
    // 创建1个组件对象（全局组件/子组件）
    Vue.component('global', {
        template: `
            <div>
                <input type="text" v-model="myText">
            </div>
        `,
        data() {
            return {
                myText: ''
            }
        },
        methods: {
            handleClick() {
                this.$emit('my_event', this.myText)
                this.$emit('my_event', this.innerHTML)
            }
        }
    })
    // 父组件
    let vm = new Vue({
        el: '#box',
        data: {
            name: ''
        },
        methods: {
            handleShow(a) {
                this.name = a
            },
            handleButton() {
                console.log(this.$refs)
                console.log(this.$refs.myRef)
                console.log(this.$refs.myRef.value)
            }
        }
    })
</script>
</html>
```

[![img](https://billy.taoxiaoxin.club/md/2023/07/64b537fa922ee4a6b5bde7e4.gif)](https://gitee.com/xuexianqi/img/raw/master/img/12 ref-1608175598467.gif)

### 5.事件总线（不同层级的不同组件通信）

#### 原本的通信方式

[![img](https://billy.taoxiaoxin.club/md/2023/07/64b537fa922ee4a6b6b648ee.png)](https://gitee.com/xuexianqi/img/raw/master/img/image-20201218090309658.png)

#### 事件总线的通信方式

[![img](https://billy.taoxiaoxin.club/md/2023/07/64b537fa922ee4a6b73d577d.png)](https://gitee.com/xuexianqi/img/raw/master/img/image-20201217151252545.png)

#### 实例

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>子传父</title>
    <script src="https://cdn.bootcdn.net/ajax/libs/vue/2.6.12/vue.min.js"></script>
</head>
<body>

<div id="box">
    <global1></global1>
    <hr>
    <global2></global2>
</div>

</body>
<script>
    // 定义1个事件总线
    let bus = new Vue({})

    // 组件1
    Vue.component('global1', {
        template: `
            <div>
                <h3>组件1</h3>
                <input type="text" v-model="myText">
                <button @click="handleClick1">点我传递数据到另一个组件</button>
            </div>
        `,
        data() {
            return {
                myText: ''
            }
        },
        methods: {
            handleClick1() {
                console.log(this.myText)
                bus.$emit('any', this.myText)  // 通过事件总线发送
            }
        }
    })
    // 组件2
    Vue.component('global2', {
        template: `
            <div>
                <h3>组件2</h3>
                收到的消息是：{{recvText}}
            </div>
        `,
        data() {
            return {
                recvText: ''
            }
        },
        mounted() { // 组件的挂载（生命周期钩子函数中的1个），开始监听时间总线上的：any
            bus.$on('any', (item) => {
                console.log('收到了', item,)
                this.recvText = item
            })
        },
        methods: {}
    })
    // 父组件
    let vm = new Vue({
        el: '#box',
        data: {},
    })
</script>
</html>
```

[![img](https://billy.taoxiaoxin.club/md/2023/07/64b537fa922ee4a6b8ae2cc8.gif)](https://gitee.com/xuexianqi/img/raw/master/img/13 bus.gif)

## 四：动态组件

### 1.基本使用

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>动态组件</title>
    <script src="https://cdn.bootcdn.net/ajax/libs/vue/2.6.12/vue.min.js"></script>
</head>
<body>

<div id="box">
    <ul>
        <li>
            <button @click="who='child1'">首页</button>
        </li>
        <li>
            <button @click="who='child2'">订单</button>
        </li>
        <li>
            <button @click="who='child3'">商品</button>
        </li>
    </ul>
    <component :is="who"></component>
</div>

</body>
<script>
    let vm = new Vue({
        el: '#box',
        data: {
            who: 'child1'
        },
        components: {
            child1: {
                template: `
                    <div>
                        <span style="border-bottom: 5px solid rgba(255,104,104,0.7)">我是首页</span>
                    </div>
                `,
            },
            child2: {
                template: `
                    <div>
                        <span style="border-bottom: 5px solid rgba(255,104,255,0.7)">我是订单</span>
                    </div>
                `,
            },
            child3: {
                template: `
                    <div>
                        <span style="border-bottom: 5px solid rgba(104,255,104,0.7)">我是商品</span>
                    </div>
                `,
            }
        }
    })
</script>
</html>
```

[![img](https://billy.taoxiaoxin.club/md/2023/07/64b537fa922ee4a6b9ee87c5.gif)](https://gitee.com/xuexianqi/img/raw/master/img/16 dynamic01.gif)

### 2.keep-alive的使用

`keep-alive`可以让输入框内有的内容一致保持，不会因为切换而重置

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <script src="js/vue.js"></script>
</head>
<body>

<div id="box">
    <ul>
        <li>
            <button @click="who='child1'">首页</button>
        </li>
        <li>
            <button @click="who='child2'">订单</button>
        </li>
        <li>
            <button @click="who='child3'">商品</button>
        </li>
    </ul>
    <keep-alive>
        <component :is="who"></component>
    </keep-alive>
</div>

</body>
<script>
    let vm = new Vue({
        el: '#box',
        data: {
            who: 'child1'
        },
        components: {
            child1: {
                template: `
                    <div>
                        <span style="border-bottom: 5px solid rgba(255,104,104,0.7)">我是首页</span>
                        <input type="text">
                    </div>
                `,
            },
            child2: {
                template: `
                    <div>
                        <span style="border-bottom: 5px solid rgba(255,104,255,0.7)">我是订单</span>
                        <input type="text">
                    </div>
                `,
            },
            child3: {
                template: `
                    <div>
                        <span style="border-bottom: 5px solid rgba(104,255,104,0.7)">我是商品</span>
                        <input type="text">
                    </div>
                `,
            }
        }
    })
</script>
</html>
```

[![img](https://billy.taoxiaoxin.club/md/2023/07/64b537fa922ee4a6ba9ca0b1.gif)](https://gitee.com/xuexianqi/img/raw/master/img/17 dynamic02.gif)

## 五：slot 插槽

- 一般情况下，编写完1个组件之后，组件的内容都是写死的，需要加数据 只能去组件中修改，扩展性很差
- 然后就出现了**插槽**这个概念，只需在组件中添加`<slot></slot>`，就可以在body的组件标签中添加内容

### 1.基本使用

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>slot 插槽</title>
    <script src="https://cdn.bootcdn.net/ajax/libs/vue/2.6.12/vue.min.js"></script>
</head>
<body>

<div id="box">
    <child>
        <h6>Hello World</h6>
    </child>
</div>

</body>
<script>
    let vm = new Vue({
        el: '#box',
        data: {
            who: 'child1'
        },
        components: {
            child: {
                template: `
                    <div>
                        <slot></slot>
                        <span style="border-bottom: 5px solid rgba(255,104,104,0.7)">我是组件的原内容</span>
                        <slot></slot>
                    </div>
                `,
            },
        }
    })
</script>
</html>
```

[![img](https://billy.taoxiaoxin.club/md/2023/07/64b537fa922ee4a6bb23c81b.png)](https://gitee.com/xuexianqi/img/raw/master/img/image-20201217121749833.png)

### 2.小案例（通过插槽实现在1个组件中控制另1个组件的显示隐藏）

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>slot 插槽</title>
    <script src="https://cdn.bootcdn.net/ajax/libs/vue/2.6.12/vue.min.js"></script>
</head>
<body>

<div id="box">
    <!--通过插槽实现在一个组件中控制另一个组件的显示隐藏-->
    <child1>
        <button @click="isShow=!isShow">显示/隐藏组件2</button>
    </child1>

    <child2 v-if="isShow"></child2>
</div>
</body>
<script>
    Vue.component('child1', {
        template: `<div>
          组件1
          <slot></slot>
        </div>`,

    })
    Vue.component('child2', {
        template: `<div>
          <h3>组件2</h3>
        </div>`,

    })
    var vm = new Vue({
        el: '#box',
        data: {
            isShow: true
        }

    })
</script>
</html>
```

[![img](https://billy.taoxiaoxin.club/md/2023/07/64b537fa922ee4a6bcc75cf1.gif)](https://gitee.com/xuexianqi/img/raw/master/img/15 slot1.gif)

### 3.具名插槽

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>具名插槽</title>
    <script src="https://cdn.bootcdn.net/ajax/libs/vue/2.6.12/vue.min.js"></script>
</head>
<body>

<div id="box">
    <!-- 具名插槽，把p标签给a插槽，div标签给b插槽-->
    <child>
        <p slot="a">我是具名插槽a插入的内容</p>
        <div slot="b">我是具名插槽b插入的内容</div>
    </child>
</div>
</body>
<script>
    Vue.component('child', {
        template: `<div>
            <slot name="a"></slot>
            <hr>
            <span style="border-bottom: 5px solid rgba(255,104,104,0.7)">我是组件的原内容</span>
            <hr>
            <slot name="b"></slot>
        </div>`,

    })
    var vm = new Vue({
        el: '#box',
        data: {}

    })
</script>
</html>
```

可以指定标签放在某个插槽的位置

[![img](https://billy.taoxiaoxin.club/md/2023/07/64b537fa922ee4a6bd7551ed.png)](https://gitee.com/xuexianqi/img/raw/master/img/image-20201217122536239.png)

## 六：自定义组件的封装

#### 详情见：http://www.xuexianqi.top/archives/732.html

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <link rel="stylesheet" href="https://unpkg.com/swiper/swiper-bundle.css">
    <script src="https://unpkg.com/swiper/swiper-bundle.js"></script>
    <script src="https://cdn.bootcdn.net/ajax/libs/vue/2.6.12/vue.min.js"></script>
    <style>
        .swiper-container {
            width: 600px;
            height: 200px;
        }
    </style>
</head>
<body>

<div id="box">

    <swipper>
        <div class="swiper-wrapper">
            <div class="swiper-slide" v-for="data in dataList1"><h1 style="text-align: center">{{data}}</h1></div>
        </div>
    </swipper>

    <swipper :key="dataList2.length">
        <div class="swiper-wrapper">
            <div class="swiper-slide" v-for="data in dataList2"><h1 style="text-align: center">{{data}}</h1></div>
        </div>
    </swipper>

</div>

</body>
<script>
    Vue.component('swipper', {
        template: `
        <div>
            <div class="swiper-container">
                <slot></slot>
                <div class="swiper-pagination"></div>
            </div>
        </div>
        `,
        mounted() {
            // 每次更新都会执行该代码，会耗费资源
            let mySwiper = new Swiper('.swiper-container', {
                direction: 'horizontal', // 垂直切换选项
                loop: true, // 循环模式选项
                // 如果需要分页器
                pagination: {
                    el: '.swiper-pagination',
                },
            })
        }
    })

    let vm = new Vue({
        el: '#box',
        data: {
            dataList1: [],
            dataList2: []
        },
        mounted() {
            setTimeout(() => {
                this.dataList1 = ['11111', '22222', '33333']
                this.dataList2 = ['66666', '77777', '88888']
            }, 3000)
        },
    })
</script>
</html>
```

## 七：自定义指令

### 1.基本使用

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>自定义指令 基本使用</title>
    <script src="https://cdn.bootcdn.net/ajax/libs/vue/2.6.12/vue.min.js"></script>
</head>
<body>

<div id="box">
    <div v-mystyle>我是1个DIV</div>
</div>

</body>
<script>
    // 自定义指令，使用的时候 v-自定义指令名
    Vue.directive('mystyle', {
        inserted(ev) {    // 在标签上使用这个指令，就会触发 inserted
            console.log('我执行了')
        }
    })


    let vm = new Vue({
        el: '#box'
    })
</script>
</html>
```

[![img](https://billy.taoxiaoxin.club/md/2023/07/64b537fa922ee4a6be6a6a5d.gif)](https://gitee.com/xuexianqi/img/raw/master/img/23 zdyzhiling.gif)

### 2.让所有使用自定义指令的标签背景都变红色

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>自定义指令 基本使用</title>
    <script src="https://cdn.bootcdn.net/ajax/libs/vue/2.6.12/vue.min.js"></script>
</head>
<body>

<div id="box">
    <div v-mystyle>我是1个DIV</div>
    <br>
    <div v-mystyle>我也是1个DIV</div>
</div>

</body>
<script>
    // 自定义指令，使用的时候 v-自定义指令名
    Vue.directive('mystyle', {
        inserted(ev) {    // 在标签上使用这个指令，就会触发 inserted
            ev.style.background='red'
        }
    })

    let vm = new Vue({
        el: '#box'
    })
</script>
</html>
```

[![img](https://billy.taoxiaoxin.club/md/2023/07/64b537fa922ee4a6bfda3877.png)](https://gitee.com/xuexianqi/img/raw/master/img/image-20201218113850277.png)

### 3.用户指定自定义指令的背景色,修改变量，背景变化

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>自定义指令</title>
    <script src="https://cdn.bootcdn.net/ajax/libs/vue/2.6.12/vue.min.js"></script>
</head>
<body>

<div id="box">
    <!--    <div v-mystyle>我是1个DIV</div>-->
    <div v-mystyle>我是1个DIV</div>
    <div v-mystyle="'red'">我是1个DIV</div>
    <div v-mystyle="'green'">我是1个DIV</div>
    <div v-mystyle="'blue'">我是1个DIV</div>
    <div v-mystyle="myColor">我是1个DIV</div>
</div>

</body>
<script>
    Vue.directive('mystyle', {
        inserted(ev, color) {    // 这里的ev就是DOM对象
            console.log(ev)
            console.log(color)
            ev.style.backgrond = color.value
        },
        updated(el, color) {
            el.style.background = color.value
        }
    })

    let vm = new Vue({
        el: '#box',
        data: {
            myColor: 'purple'
        }
    })
</script>
</html>
```

## 八：过滤器

#### json数据：film.json

```json
{
  "coming": [
    {
      "id": 1240838,
      "haspromotionTag": false,
      "img": "http://p1.meituan.net/w.h/movie/38dd31a0e1b18e1b00aeb2170c5a65b13885486.jpg",
      "version": "",
      "nm": "除暴",
      "preShow": false,
      "sc": 8.6,
      "globalReleased": true,
      "wish": 76513,
      "star": "王千源,吴彦祖,春夏",
      "rt": "2020-11-20",
      "showInfo": "今天50家影院放映79场",
      "showst": 3,
      "wishst": 0,
      "comingTitle": "11月20日 周五"
    },
    {
      "id": 1228788,
      "haspromotionTag": false,
      "img": "http://p0.meituan.net/w.h/movie/b16c1c0d5ac9e743c6ffbbf7eba900522725807.jpg",
      "version": "",
      "nm": "一秒钟",
      "preShow": false,
      "sc": 8.6,
      "globalReleased": true,
      "wish": 54493,
      "star": "张译,刘浩存,范伟",
      "rt": "2020-11-27",
      "showInfo": "今天11家影院放映12场",
      "showst": 3,
      "wishst": 0,
      "comingTitle": "11月27日 周五"
    },
    {
      "id": 1358968,
      "haspromotionTag": false,
      "img": "http://p0.meituan.net/w.h/movie/d33858dbfc207da3b36c0dc7fff7a8bb2028677.jpg",
      "version": "",
      "nm": "汪汪队立大功之超能救援",
      "preShow": false,
      "sc": 8.3,
      "globalReleased": true,
      "wish": 24833,
      "star": "杨鸥,韩娇娇,李敏妍",
      "rt": "2020-11-13",
      "showInfo": "今天5家影院放映7场",
      "showst": 3,
      "wishst": 0,
      "comingTitle": "11月13日 周五"
    },
    {
      "id": 345809,
      "haspromotionTag": false,
      "img": "http://p1.meituan.net/w.h/moviemachine/7c4ba9633635503044a8f8fb6426aa8d416264.jpg",
      "version": "v2d imax",
      "nm": "隐形人",
      "preShow": false,
      "sc": 8.4,
      "globalReleased": true,
      "wish": 9894,
      "star": "伊丽莎白·莫斯,奥利弗·杰森-科恩,阿尔迪斯·霍吉",
      "rt": "2020-12-04",
      "showInfo": "今天21家影院放映30场",
      "showst": 3,
      "wishst": 0,
      "comingTitle": "12月4日 周五"
    },
    {
      "id": 1330790,
      "haspromotionTag": false,
      "img": "http://p0.meituan.net/w.h/movie/88e54f3e670789ba1f08e48a5f1170c1188102.jpg",
      "version": "",
      "nm": "明天你是否依然爱我",
      "preShow": false,
      "sc": 0,
      "globalReleased": false,
      "wish": 217699,
      "star": "杨颖,李鸿其,黄柏钧",
      "rt": "2020-12-24",
      "showInfo": "2020-12-24 下周四上映",
      "showst": 4,
      "wishst": 0,
      "comingTitle": "12月24日 周四"
    },
    {
      "id": 1277751,
      "haspromotionTag": false,
      "img": "http://p0.meituan.net/w.h/movie/303c2e671cc4df875c151d688ecbd8962085989.jpg",
      "version": "v2d imax",
      "nm": "赤狐书生",
      "preShow": false,
      "sc": 7.7,
      "globalReleased": true,
      "wish": 177525,
      "star": "陈立农,李现,哈妮克孜",
      "rt": "2020-12-04",
      "showInfo": "今天26家影院放映43场",
      "showst": 3,
      "wishst": 0,
      "comingTitle": "12月4日 周五"
    },
    {
      "id": 1225578,
      "haspromotionTag": false,
      "img": "http://p0.meituan.net/w.h/moviemachine/cf7d6942f2aa9189cce20519b490b6b1879487.jpg",
      "version": "",
      "nm": "野性的呼唤",
      "preShow": false,
      "sc": 9.2,
      "globalReleased": true,
      "wish": 14703,
      "star": "哈里森·福特,丹·史蒂文斯,凯伦·吉兰",
      "rt": "2020-11-13",
      "showInfo": "今天暂无场次",
      "showst": 3,
      "wishst": 0,
      "comingTitle": "11月13日 周五"
    },
    {
      "id": 1302281,
      "haspromotionTag": false,
      "img": "http://p0.meituan.net/w.h/moviemachine/1d2b4985d0187b437d41a73994ba2e191607376.jpg",
      "version": "",
      "nm": "奇妙王国之魔法奇缘",
      "preShow": true,
      "sc": 0,
      "globalReleased": false,
      "wish": 20309,
      "star": "卢瑶,张洋,陈新玥",
      "rt": "2020-12-26",
      "showInfo": "2020-12-26 下周六上映",
      "showst": 4,
      "wishst": 0,
      "comingTitle": "12月26日 周六"
    },
    {
      "id": 1301902,
      "haspromotionTag": false,
      "img": "http://p0.meituan.net/w.h/movie/f686425a1ad1f502254abef593d508bf428685.jpg",
      "version": "",
      "nm": "沉默东京",
      "preShow": false,
      "sc": 5.8,
      "globalReleased": true,
      "wish": 52,
      "star": "佐藤浩市,石田百合子,西岛秀俊",
      "rt": "2020-12-04",
      "showInfo": "今天暂无场次",
      "showst": 3,
      "wishst": 0,
      "comingTitle": ""
    },
    {
      "id": 1286015,
      "haspromotionTag": false,
      "img": "http://p0.meituan.net/w.h/moviemachine/a0c6d6e130abe399e4cba58be2b1f871840268.jpg",
      "version": "",
      "nm": "宝可梦：超梦的逆袭 进化",
      "preShow": false,
      "sc": 8.2,
      "globalReleased": true,
      "wish": 53255,
      "star": "松本梨香,大谷育江,市村正亲",
      "rt": "2020-12-04",
      "showInfo": "今天10家影院放映10场",
      "showst": 3,
      "wishst": 0,
      "comingTitle": "12月4日 周五"
    }
  ]
}
```

#### 前端：index.html

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>过滤器</title>
    <script src="https://cdn.bootcdn.net/ajax/libs/vue/2.6.12/vue.min.js"></script>
    <script src="https://cdn.bootcdn.net/ajax/libs/axios/0.21.0/axios.min.js"></script>
</head>
<body>

<div id="box">
    <ul>
        <li v-for="item in dataList">
            <h2>{{item.nm}}</h2>
            <p>主演：{{item.star}}</p>
            <img :src="item.img | repUrl" alt="">
        </li>
    </ul>
</div>

</body>
<script>
    // 过滤器
    Vue.filter('repUrl', function (url) {
        return url.replace('w.h','128.180')
    })
    let vm = new Vue({
        el: '#box',
        data: {
            dataList: ''
        },
        mounted() {
            axios.get("http://127.0.0.1:5000/").then(res => {
                console.log(res.data.coming)
                this.dataList = res.data.coming
            }).catch(err => {
                console.log(err);
            })
        }
    })
</script>
</html>
```

#### 后端：main.py

```python
import json

from flask import Flask, jsonify

app = Flask(__name__)


@app.route('/')
def index():
    print('请求来了')
    with open('film.json', mode='rt', encoding='utf-8') as f:
        dic = json.load(f)
    res = jsonify(dic)
    res.headers['Access-Control-Allow-Origin'] = '*'
    return res


if __name__ == '__main__':
    app.run()
```

#### 组件的分类：

- **全局组件**：全局组件，顾名思义，就是在整个 Vue.js 应用中都能被访问和使用的组件。
- **局部组件：** 

[[外链图片转存失败,源站可能有防盗链机制,建议将图片保存下来直接上传(img-7MZ4xCmy-1619795199183)(https://gitee.com/xuexianqi/img/raw/master/img/22%20filter.gif)]](https://gitee.com/xuexianqi/img/raw/master/img/22 filter.gif)