

MVVM -`Model View ViewModel`，最核心的就是 `ViewModel` 。

## 一.MVVM介绍：

MVVM（Model-View-ViewModel）是一种用于前端开发的架构模式，它是一种事件驱动的编程方式。

在MVVM模式中，有以下几个核心组件：

1. Model（模型）：代表应用程序中的数据和业务逻辑。在Vue.js中，Model是指Vue实例中的`data`属性中的数据。这些数据需要在页面中显示。
2. View（视图）：是用户界面的可视部分，通常是由HTML和CSS构成。在Vue.js中，View就是模板（Template），用于显示数据。
3. ViewModel（视图模型）：是MVVM模式的核心部分，负责连接View和Model。它包含了DOM监听器（DOM Listener）和数据绑定（Data Binding）两个重要部分。
   - 数据绑定（Data Bindings）：MVVM模式通过数据绑定实现了从Model到View的映射关系。当Model中的数据发生变化时，ViewModel会自动更新相关的View，省去了手动更新视图的代码和时间。这种双向绑定机制大大简化了开发流程，减少了开发者的工作量。
   - DOM监听器（DOM Listeners）：MVVM模式还通过DOM监听器实现了从View到Model的事件监听。当用户在View中进行交互操作时，例如点击按钮或输入内容，DOM监听器会捕获这些事件并触发对应的处理函数，从而更新Model中的数据。这种机制实现了View对Model的响应，并且将数据请求和视图请求进行了解耦，彼此之间相互独立，提高了代码的可维护性和灵活性。

综上所述，MVVM模式通过数据绑定和DOM监听器实现了Model（数据）与View（视图）之间的自动同步。开发者只需要关注数据的变化和用户交互事件的处理逻辑，而不需要手动更新视图或处理视图事件。这种架构模式使得开发更加高效、简洁，并且提高了代码的可读性和可维护性。

![image-20210425214215993](https://billy.taoxiaoxin.club/md/2023/06/647de673922ee46f3ecd4a47.png)

[更多参考简书：聊一聊基本的MVVM设计思想](https://www.jianshu.com/p/cf224ae5deb6)

## 二.MVVM特性：

- **低耦合**：`视图`（View）可以`独立于Model变化和修改`，1个ViewModel可以绑定到不同的View上，当View变化的时候 Model可以不变，当Model变化的时候 View也可以不变
- **可复用**：可以把一些视图逻辑放在1个ViewModel中，让很多View`重用这端视图的逻辑`（以此减少代码冗余）
- **独立开发**：`开发`人员可以专注于`业务逻辑`和`数据的开发`（ViewModel），`设计`人员可以专注于`页面设计`
- **可测试**：界面元素是比较难以测试的，而现在的测试可以`针对ViewModel`来编写

## 三.MVVM逻辑

![img](https://billy.taoxiaoxin.club/md/2023/06/647de673922ee46f3f18823a.jpeg)

------

## 四.两个方向：

MVVM思想有两个方向。

- 一是将模型转换成视图，即将后端传递的数据转换成看到的页面。 实现方式是：**数据绑定**。
- 二是将视图转换成模型，即将看到的页面转换成后端的数据。实现的方式是：**DOM 事件监听**。

这两个方向都实现的，就称为**数据的双向绑定**。

## 五.MVC 和 MVVM 的区别(关系)

MVC - `Model View Controller`( controller: 控制器 )，**M** 和 **V** 和 MVVM 中的 M 和 V 意思一样，**C** 指**页面业务逻辑**。使用 MVC 的目的就是将 M 和 V 的代码分离，但 MVC 是单向通信，也就是将 Model 渲染到 View 上，必须通过 Controller 来承上启下。

MVC 和 MVVM 的区别(关系)并不是 ViewModel 完全取代了 Controller 。

- ViewModel 目的在于抽离 Controller 中的数据渲染功能，而不是替代。
- 其他操作业务等还是应该放在 Controller 中实现，这样就实现了业务逻辑组件的复用。

## 六. 常见关于Vue的面试题

- 什么是MVVM思想？

> MVVM -`Model View ViewModel`，它包括 DOM Listenters 和 Data bindings，前者实现了页面与数据的绑定，当数据发生变化的时候会自动渲染页面。后者实现了数据与页面的绑定，当页面操作数据的时候 DOM 和 Model 也会发生相应的变化。

- MVVM相对于MVC的优势？

> 1. MVVM 实现了数据与页面的双向绑定，MVC 只实现了 Model 和 View 的单向绑定。
> 2. MVVM 实现了页面业务逻辑和渲染之间的解耦，也实现了数据与视图的解耦，并且可以组件化开发。

- VUE是如何体现MVVM思想的？

> 1. 数据绑定：Vue.js 使用了胡子语法（Mustache syntax）来实现数据与视图的绑定。通过在模板中使用双大括号（{{ }}）将数据绑定到视图中，Vue.js能够自动将数据的变化反映到视图中。这使得开发者无需手动更新视图，只需关注数据的变化，实现了数据驱动的视图更新。
> 2. 双向数据绑定：Vue.js 还通过 v-model 指令实现了双向数据绑定。v-model 可以将表单元素的值与 Vue 实例中的数据进行双向绑定。当用户在表单元素中输入内容时，数据会自动更新；反过来，当数据发生变化时，表单元素的值也会相应地更新。这种双向数据绑定让开发者能够轻松实现表单的数据收集与处理，减少了手动处理输入事件的代码量。
> 3. 事件绑定：Vue.js 提供了 v-on 指令来实现事件绑定。通过在模板中使用 v-on 指令，可以将 DOM 事件与 Vue 实例中的方法进行绑定。当事件触发时，Vue.js 能够自动调用对应的方法，进行相应的业务处理。这样，开发者可以通过事件操作数据，实现了从视图到数据的响应，体现了 MVVM 模式中的 ViewModel。

