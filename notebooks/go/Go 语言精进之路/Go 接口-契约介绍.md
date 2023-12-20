# Go 接口-契约介绍

[TOC]



## 一、接口基本介绍

### 1.1 接口类型介绍

接口是一种抽象类型，它定义了一组方法的契约，它规定了需要实现的所有方法。**是由 `type` 和 `interface` 关键字定义的一组方法集合，其中，方法集合唯一确定了这个接口类型所表示的接口。**

一个接口类型通常由一组方法签名组成，这些方法定义了对象必须实现的操作。接口的方法签名包括方法的名称、输入参数、返回值等信息，但不包括方法的实际实现。例如：

```go
type Writer interface {
    Write([]byte) (int, error)
}
```

上面的代码定义了一个名为 `Writer` 的接口，它有一个 `Write` 方法，该方法接受一个 `[]byte` 类型的参数并返回两个值，一个整数和一个错误。任何类型只要实现了这个 `Write` 方法的签名，就可以被认为是 `Writer` 接口的实现。

总之，Go语言提倡面向接口编程。

### 1.2 为什么要使用接口

现在假设我们的代码世界里有很多小动物，下面的代码片段定义了猫和狗，它们饿了都会叫。

```go
package main

import "fmt"

type Cat struct{}

func (c Cat) Say() {
	fmt.Println("喵喵喵~")
}

type Dog struct{}

func (d Dog) Say() {
	fmt.Println("汪汪汪~")
}

func main() {
	c := Cat{}
	c.Say()
	d := Dog{}
	d.Say()
}
```

这个时候又跑来了一只羊，羊饿了也会发出叫声。

```go
type Sheep struct{}

func (s Sheep) Say() {
	fmt.Println("咩咩咩~")
}
```

我们接下来定义一个饿肚子的场景。

```go
// MakeCatHungry 猫饿了会喵喵喵~
func MakeCatHungry(c Cat) {
	c.Say()
}

// MakeSheepHungry 羊饿了会咩咩咩~
func MakeSheepHungry(s Sheep) {
	s.Say()
}
```

接下来会有越来越多的小动物跑过来，我们的代码世界该怎么拓展呢？

在饿肚子这个场景下，我们可不可以把所有动物都当成一个“会叫的类型”来处理呢？当然可以！使用接口类型就可以实现这个目标。 我们的代码其实并不关心究竟是什么动物在叫，我们只是在代码中调用它的`Say()`方法，这就足够了。

我们可以约定一个`Sayer`类型，它必须实现一个`Say()`方法，只要饿肚子了，我们就调用`Say()`方法。

```go
type Sayer interface {
    Say()
}
```

然后我们定义一个通用的`MakeHungry`函数，接收`Sayer`类型的参数。

```go
// MakeHungry 饿肚子了...
func MakeHungry(s Sayer) {
	s.Say()
}
```

我们通过使用接口类型，把所有会叫的动物当成`Sayer`类型来处理，只要实现了`Say()`方法都能当成`Sayer`类型的变量来处理。

```go
var c cat
MakeHungry(c)
var d dog
MakeHungry(d)
```

在电商系统中我们允许用户使用多种支付方式（支付宝支付、微信支付、银联支付等），我们的交易流程中可能不太在乎用户究竟使用什么支付方式，只要它能提供一个实现支付功能的`Pay`方法让调用方调用就可以了。

再比如我们需要在某个程序中添加一个将某些指标数据向外输出的功能，根据不同的需求可能要将数据输出到终端、写入到文件或者通过网络连接发送出去。在这个场景下我们可以不关注最终输出的目的地是什么，只需要它能提供一个`Write`方法让我们把内容写入就可以了。

Go语言中为了解决类似上面的问题引入了接口的概念，接口类型区别于我们之前章节中介绍的那些具体类型，让我们专注于该类型提供的方法，而不是类型本身。使用接口类型通常能够让我们写出更加通用和灵活的代码。

### 1.3 面向接口编程

PHP、Java等语言中也有接口的概念，不过在PHP和Java语言中需要显式声明一个类实现了哪些接口，在Go语言中使用隐式声明的方式实现接口。只要一个类型实现了接口中规定的所有方法，那么它就实现了这个接口。

Go语言中的这种设计符合程序开发中抽象的一般规律，例如在下面的代码示例中，我们的电商系统最开始只设计了支付宝一种支付方式：

```go
type ZhiFuBao struct {
	// 支付宝
}

// Pay 支付宝的支付方法
func (z *ZhiFuBao) Pay(amount int64) {
  fmt.Printf("使用支付宝付款：%.2f元。\n", float64(amount/100))
}

// Checkout 结账
func Checkout(obj *ZhiFuBao) {
	// 支付100元
	obj.Pay(100)
}

func main() {
	Checkout(&ZhiFuBao{})
}
```

随着业务的发展，根据用户需求添加支持微信支付。

```go
type WeChat struct {
	// 微信
}

// Pay 微信的支付方法
func (w *WeChat) Pay(amount int64) {
	fmt.Printf("使用微信付款：%.2f元。\n", float64(amount/100))
}
```

在实际的交易流程中，我们可以根据用户选择的支付方式来决定最终调用支付宝的Pay方法还是微信支付的Pay方法。

```go
// Checkout 支付宝结账
func CheckoutWithZFB(obj *ZhiFuBao) {
	// 支付100元
	obj.Pay(100)
}

// Checkout 微信支付结账
func CheckoutWithWX(obj *WeChat) {
	// 支付100元
	obj.Pay(100)
}
```

实际上，从上面的代码示例中我们可以看出，我们其实并不怎么关心用户选择的是什么支付方式，我们只关心调用Pay方法时能否正常运行。这就是典型的“不关心它是什么，只关心它能做什么”的场景。

在这种场景下我们可以将具体的支付方式抽象为一个名为`Payer`的接口类型，即任何实现了`Pay`方法的都可以称为`Payer`类型。

```go
// Payer 包含支付方法的接口类型
type Payer interface {
	Pay(int64)
}
```

此时只需要修改下原始的`Checkout`函数，它接收一个`Payer`类型的参数。这样就能够在不修改既有函数调用的基础上，支持新的支付方式。

```go
// Checkout 结账
func Checkout(obj Payer) {
	// 支付100元
	obj.Pay(100)
}

func main() {
	Checkout(&ZhiFuBao{}) // 之前调用支付宝支付

	Checkout(&WeChat{}) // 现在支持使用微信支付
}
```

像类似的例子在我们编程过程中会经常遇到：

- 比如一个网上商城可能使用支付宝、微信、银联等方式去在线支付，我们能不能把它们当成“支付方式”来处理呢？
- 比如三角形，四边形，圆形都能计算周长和面积，我们能不能把它们当成“图形”来处理呢？
- 比如满减券、立减券、打折券都属于电商场景下常见的优惠方式，我们能不能把它们当成“优惠券”来处理呢？

接口类型是Go语言提供的一种工具，在实际的编码过程中是否使用它由你自己决定，但是通常使用接口类型可以使代码更清晰易读。

### 1.4 接口的定义

每个接口类型由任意个方法签名组成，接口的定义格式如下：

```go
type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
```

其中：

- 接口类型名：Go语言的接口在命名时，一般会在单词后面添加`er`，如有写操作的接口叫`Writer`，有关闭操作的接口叫`closer`等。接口名最好要能突出该接口的类型含义。
- 方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
- 参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。

下面是一个典型的接口类型 `MyInterface` 的定义：

```go
type MyInterface interface {
    M1(int) error
    M2(io.Writer, ...string)
}
```

通过这个定义，我们可以看到，接口类型 `MyInterface` 所表示的接口的方法集合，包含两个方法 `M1` 和 `M2`。**之所以称 `M1` 和 `M2` 为“方法”，更多是从这个接口的实现者的角度考虑的。**但从上面接口类型声明中各个“方法”的形式上来看，**这更像是不带有 `func` 关键字的函数名 + 函数签名（参数列表 + 返回值列表）的组合。**

在接口类型的方法集合中声明的方法，它的参数列表不需要写出形参名字，返回值列表也是如此。也就是说，方法的参数列表中形参名字与返回值列表中的具名返回值，都不作为区分两个方法的凭据。

比如下面的 `MyInterface` 接口类型的定义与上面的 `MyInterface` 接口类型定义都是等价的：

~~~go
type MyInterface interface {
    M1(a int) error
    M2(w io.Writer, strs ...string)
}

type MyInterface interface {
    M1(n int) error
    M2(w io.Writer, args ...string)
}
~~~

不过，**Go 语言要求接口类型声明中的方法必须是具名的，并且方法名字在这个接口类型的方法集合中是唯一的。**前面我们在学习类型嵌入时就学到过：Go 1.14 版本以后，Go 接口类型允许嵌入的不同接口类型的方法集合存在交集，但前提是交集中的方法不仅名字要一样，它的方法签名部分也要保持一致，也就是参数列表与返回值列表也要相同，否则 Go 编译器照样会报错。

比如下面示例中 `Interface3` 嵌入了 `Interface1` 和 `Interface2`，但后两者交集中的 `M1` 方法的函数签名不同，导致了编译出错：

```go
type Interface1 interface {
    M1()
}
type Interface2 interface {
    M1(string) 
    M2()
}

type Interface3 interface{
    Interface1
    Interface2 // 编译器报错：duplicate method M1
    M3()
}
```

上面举的例子中的方法都是首字母大写的导出方法，所以在 Go 接口类型的方法集合中放入首字母小写的非导出方法也是合法的，并且我们在 Go 标准库中也找到了带有非导出方法的接口类型定义，比如 `context` 包中的 `canceler` 接口类型，它的代码如下：

```go
// $GOROOT/src/context.go

// A canceler is a context type that can be canceled directly. The
// implementations are *cancelCtx and *timerCtx.
type canceler interface {
    cancel(removeFromParent bool, err error)
    Done() <-chan struct{}
}
```

但这样的例子并不多。通过对标准库这为数不多的例子，我们可以看到，如果接口类型的方法集合中包含非导出方法，那么这个接口类型自身通常也是非导出的，它的应用范围也仅局限于包内。不过，在日常实际编码过程中，我们极少使用这种带有非导出方法的接口类型，我们简单了解一下就可以了。

## 二、空接口

除了上面这种常规情况，还有空接口（empty interface）类型这种特殊情况。

### 2.1 空接口的定义

空接口是指没有定义任何方法的接口类型。因此任何类型都可以视为实现了空接口。也正是因为空接口类型的这个特性，空接口类型的变量可以存储任意类型的值。

比如下面的 `EmptyInterface` 接口类型：

```go
type EmptyInterface interface {

}
```

这个方法集合为空的接口类型就被称为空接口类型，但通常我们不需要自己显式定义这类空接口类型，我们直接使用 `interface{}` 这个类型字面值作为所有空接口类型的代表就可以了。

### 2.2 空接口的应用

#### 2.2.1 空接口作为函数的参数

空接口（`interface{}`）作为函数的参数是一种非常灵活的方式，因为它可以接受任何类型的参数。这在处理未知类型的数据或编写通用函数时非常有用。以下是一个示例，展示了如何使用空接口作为函数参数：

```go
package main

import "fmt"

func PrintValue(value interface{}) {
    fmt.Println(value)
}

func main() {
    PrintValue(42)                 // 整数
    PrintValue("Hello, Go!")       // 字符串
    PrintValue(3.14159)            // 浮点数
    PrintValue([]int{1, 2, 3})     // 切片
}

```

在上面的示例中，`PrintValue` 函数接受一个空接口类型的参数，这意味着它可以接受任何类型的值。在 `main` 函数中，我们调用 `PrintValue` 函数并传递不同类型的参数，它们都可以被正确处理和打印。

#### 2.2.2 空接口作为map的值

空接口也可以用作`map`的值类型，这使得`map`可以存储不同类型的值。这在需要将各种类型的数据关联到特定键时非常有用。以下是一个示例：

```go
package main

import "fmt"

func main() {
    data := make(map[string]interface{})

    data["name"] = "Alice"
    data["age"] = 30
    data["isStudent"] = false

    fmt.Println(data["name"])       // 输出: Alice
    fmt.Println(data["age"])        // 输出: 30
    fmt.Println(data["isStudent"])  // 输出: false
}
```

在上面的示例中，我们创建了一个`map`，其中值的类型是`interface{}`，这意味着`map`可以存储不同类型的值。我们使用字符串键将字符串、整数和布尔值关联到`map`中，并在后续通过键来访问这些值。

### 2.3 接口类型变量

接口类型一旦被定义后，它就和其他 Go 类型一样可以用于声明变量，比如：

```go
var err error   // err是一个error接口类型的实例变量
var r io.Reader // r是一个io.Reader接口类型的实例变量
```

**这些类型为接口类型的变量被称为接口类型变量，如果没有被显式赋予初值，接口类型变量的默认值为 `nil`。**如果要为接口类型变量显式赋予初值，我们就要为接口类型变量选择合法的右值。

Go `规定`：如果一个类型 `T` 的方法集合是某接口类型 `I` 的方法集合的等价集合或超集，我们就说类型 `T` 实现了接口类型 `I`，那么类型 `T` 的变量就可以作为合法的右值赋值给接口类型 `I` 的变量。

如果一个变量的类型是空接口类型，由于空接口类型的方法集合为空，这就意味着任何类型都实现了空接口的方法集合，所以我们可以将任何类型的值作为右值，赋值给空接口类型的变量，比如下面例子：

```go
var i interface{} = 15 // ok
i = "hello, golang" // ok
type T struct{}
var t T
i = t  // ok
i = &t // ok
```

空接口类型的这一可接受任意类型变量值作为右值的特性，让**它成为 Go 加入泛型语法之前唯一一种具有“泛型”能力的语法元素**，包括 Go 标准库在内的一些通用数据结构与算法的实现，都使用了`空类型` `interface{}`作为数据元素的类型，这样我们就无需为每种支持的元素类型单独做一份代码拷贝了。

### 2.4 类型断言

Go 语言还支持接口类型变量赋值的“逆操作”，也就是通过接口类型变量“还原”它的右值的类型与值信息，这个过程被称为“类型断言（`Type Assertion`）”。类型断言通常使用下面的语法形式：

```go
v, ok := i.(T) 
```

其中 `i` 是某一个接口类型变量，如果 `T` 是一个非接口类型且 `T` 是想要还原的类型，那么这句代码的含义就是**断言存储在接口类型变量 `i` 中的值的类型为 `T`。**

如果接口类型变量 `i` 之前被赋予的值确为 `T` 类型的值，那么这个语句执行后，左侧“`comma, ok`”语句中的变量 `ok` 的值将为 `true`，变量 `v` 的类型为 `T`，它的值会是之前变量 `i` 的右值。如果 `i` 之前被赋予的值不是 `T` 类型的值，那么这个语句执行后，变量 `ok` 的值为 `false`，变量 `v` 的类型还是那个要还原的类型，但它的值是类型 `T` 的零值。

类型断言也支持下面这种语法形式：

```go
v := i.(T)
```

但在这种形式下，一旦接口变量 `i` 之前被赋予的值不是 `T` 类型的值，那么这个语句将抛出 `panic`。如果变量 `i` 被赋予的值是 `T` 类型的值，那么变量 `v` 的类型为 `T`，它的值就会是之前变量 `i` 的右值。由于可能出现 `panic`，所以我们并不推荐使用这种类型断言的语法形式。

为了加深你的理解，接下来我们通过一个例子来直观看一下类型断言的语义：

```go
var a int64 = 13
var i interface{} = a
v1, ok := i.(int64) 
fmt.Printf("v1=%d, the type of v1 is %T, ok=%t\n", v1, v1, ok) // v1=13, the type of v1 is int64, ok=true
v2, ok := i.(string)
fmt.Printf("v2=%s, the type of v2 is %T, ok=%t\n", v2, v2, ok) // v2=, the type of v2 is string, ok=false
v3 := i.(int64) 
fmt.Printf("v3=%d, the type of v3 is %T\n", v3, v3) // v3=13, the type of v3 is int64
v4 := i.([]int) // panic: interface conversion: interface {} is int64, not []int
fmt.Printf("the type of v4 is %T\n", v4) 
```

你可以看到，这个例子的输出结果与我们之前讲解的是一致的。

在这段代码中，如果 `v, ok := i.(T)` 中的 `T` 是一个接口类型，那么类型断言的语义就会变成：断言 `i` 的值实现了接口类型 `T`。如果断言成功，变量 `v` 的类型为 `i` 的值的类型，而并非接口类型 `T`。如果断言失败，`v` 的类型信息为接口类型 `T`，它的值为 `nil`，下面我们再来看一个 `T` 为接口类型的示例：

```go
type MyInterface interface {
    M1()
}

type T int
               
func (T) M1() {
    println("T's M1")
}              
               
func main() {  
    var t T    
    var i interface{} = t
    v1, ok := i.(MyInterface)
    if !ok {   
        panic("the value of i is not MyInterface")
    }          
    v1.M1()    
    fmt.Printf("the type of v1 is %T\n", v1) // the type of v1 is main.T
               
    i = int64(13)
    v2, ok := i.(MyInterface)
    fmt.Printf("the type of v2 is %T\n", v2) // the type of v2 is <nil>
    // v2 = 13 //  cannot use 1 (type int) as type MyInterface in assignment: int does not implement MyInterface (missing M1   method) 
}
```

我们看到，通过`the type of v2 is <nil>`，我们其实是看不出断言失败后的变量 `v2` 的类型的，但通过最后一行代码的编译器错误提示，我们能清晰地看到 `v2` 的类型信息为 `MyInterface`。

其实，接口类型的类型断言还有一个变种，那就是 type switch ，这个你可以去看看[【go 流程控制之switch 语句介绍】](https://blog.csdn.net/weixin_44621343/article/details/133801611)。

## 三、尽量定义“小接口”

### 3.1 “小接口”介绍

**接口类型的背后，是通过把类型的行为抽象成契约**，建立双方共同遵守的约定，这种契约将双方的耦合降到了最低的程度。和生活工作中的契约有繁有简，签署方式多样一样，代码间的契约也有多有少，有大有小，而且达成契约的方式也有所不同。 而 Go 选择了去繁就简的形式，这主要体现在以下两点上：

+ **隐式契约，无需签署，自动生效**：Go 语言中接口类型与它的实现者之间的关系是隐式的，不需要像其他语言（比如 Java）那样要求实现者显式放置“implements”进行修饰，实现者只需要实现接口方法集合中的全部方法便算是遵守了契约，并立即生效了。
+ **更倾向于“小契约”**：这点也不难理解。你想，如果契约太繁杂了就会束缚了手脚，缺少了灵活性，抑制了表现力。所以 **Go 选择了使用“小契约”，表现在代码上就是尽量定义小接口，即方法个数在 1~3 个之间的接口。**Go 语言之父 Rob Pike 曾说过的“**接口越大，抽象程度越弱**”，这也是 Go 社区倾向定义小接口的另外一种表述。

Go 对小接口的青睐在它的标准库中体现得淋漓尽致，这里我给出了标准库中一些我们日常开发中常用的接口的定义：

```go
// $GOROOT/src/builtin/builtin.go
type error interface {
    Error() string
}

// $GOROOT/src/io/io.go
type Reader interface {
    Read(p []byte) (n int, err error)
}

// $GOROOT/src/net/http/server.go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

type ResponseWriter interface {
    Header() Header
    Write([]byte) (int, error)
    WriteHeader(int)
}
```

我们看到，上述这些接口的方法数量在 1~3 个之间，这种“小接口”的 Go 惯例也已经被 Go 社区项目广泛采用。我统计了早期版本的 Go 标准库（Go 1.13 版本）、Docker 项目（Docker 19.03 版本）以及 Kubernetes 项目（Kubernetes 1.17 版本）中定义的接口类型方法集合中方法数量，你可以看下：

![img](https://billy.taoxiaoxin.club/md/2023/11/65490213cd1648c051659338.png)

从图中我们可以看到，无论是 Go 标准库，还是 Go 社区知名项目，它们基本都遵循了“尽量定义小接口”的惯例，接口方法数量在 1~3 范围内的接口占了绝大多数。那么在编码层面，小接口究竟有哪些优势呢？

### 3.2 小接口优势

#### 3.2.1 第一点：接口越小，抽象程度越高

计算机程序本身就是对真实世界的抽象与再建构。抽象就是对同类事物去除它具体的、次要的方面，抽取它相同的、主要的方面。不同的抽象程度，会导致抽象出的概念对应的事物的集合不同。抽象程度越高，对应的集合空间就越大；抽象程度越低，也就是越具像化，更接近事物真实面貌，对应的集合空间越小。

我们举一个生活中的简单例子。你可以看下这张示意图，它是对生活中不同抽象程度的形象诠释：

![img](https://billy.taoxiaoxin.club/md/2023/11/654902b6b06562ffaf88a5ea.jpg)

这张图中我们分别建立了三个抽象：

+ 会飞的。这个抽象对应的事物集合包括：蝴蝶、蜜蜂、麻雀、天鹅、鸳鸯、海鸥和信天翁；
+ 会游泳的。它对应的事物集合包括：鸭子、海豚、人类、天鹅、鸳鸯、海鸥和信天翁；
+ 会飞且会游泳的。这个抽象对应的事物集合包括：天鹅、鸳鸯、海鸥和信天翁。

我们看到，“会飞的”、“会游泳的”这两个抽象对应的事物集合，要大于“会飞且会游泳的”所对应的事物集合空间，也就是说“会飞的”、“会游泳的”这两个抽象程度更高。

我们将上面的抽象转换为 Go 代码看看：

```go
// 会飞的
type Flyable interface {
  Fly()
}

// 会游泳的
type Swimable interface {
  Swim()
}

// 会飞且会游泳的
type FlySwimable interface {
  Flyable
  Swimable
}
```

我们用上述定义的接口替换上图中的抽象，再得到这张示意图：

![img](https://billy.taoxiaoxin.club/md/2023/11/654902fce0860f4c51178b64.jpg)

我们可以直观地看到，这张图中的 `Flyable` 只有一个 `Fly` 方法，`FlySwimable` 则包含两个方法 `Fly` 和 `Swim`。我们看到，具有更少方法的 `Flyable` 的抽象程度相对于 `FlySwimable` 要高，包含的事物集合（7 种动物）也要比 `FlySwimable` 的事物集合（4 种动物）大。也就是说，接口越小（接口方法少)，抽象程度越高，对应的事物集合越大。

而这种情况的极限恰恰就是无方法的空接口 `interface{}`，空接口的这个抽象对应的事物集合空间包含了 Go 语言世界的所有事物。

#### 3.2.2 第二点：小接口易于实现和测试

**Go 推崇通过组合的方式构建程序**。Go 开发人员一般会尝试通过嵌入其他已有接口类型的方式来构建新接口类型，就像通过嵌入 io.Reader 和 io.Writer 构建 io.ReadWriter 那样。

那构建时，如果有众多候选接口类型供我们选择，我们会怎么选择呢？

显然，我们会选择那些新接口类型需要的契约职责，同时也要求不要引入我们不需要的契约职责。在这样的情况下，拥有单一或少数方法的小接口便更有可能成为我们的目标，而那些拥有较多方法的大接口，可能会因引入了诸多不需要的契约职责而被放弃。由此可见，小接口更契合 Go 的组合思想，也更容易发挥出组合的威力。

## 四、定义小接口，可以遵循的几点

保持简单有时候比复杂更难。小接口虽好，但如何定义出小接口是摆在所有 Gopher 面前的一道难题。这道题没有标准答案，但有一些点可供我们在实践中考量遵循。

### 4.1 首先，别管接口大小，先抽象出接口

要设计和定义出小接口，前提是需要先有接口。

Go 语言还比较年轻，它的设计哲学和推崇的编程理念可能还没被广大 Gopher 100% 理解、接纳和应用于实践当中，尤其是 Go 所推崇的基于接口的组合思想。

尽管接口不是 Go 独有的，**但专注于接口是编写强大而灵活的 Go 代码的关键。**因此，在定义小接口之前，我们需要先针对问题领域进行深入理解，聚焦抽象并发现接口，就像下图所展示的那样，先针对领域对象的行为进行抽象，形成一个接口集合：

![WechatIMG267](https://billy.taoxiaoxin.club/md/2023/11/654903ee79b48b09aa6c02d5.jpg)

**初期，我们先不要介意这个接口集合中方法的数量，**因为对问题域的理解是循序渐进的，在第一版代码中直接定义出小接口可能并不现实。而且，标准库中的 `io.Reader` 和 `io.Writer` 也不是在 Go 刚诞生时就有的，而是在发现对网络、文件、其他字节数据处理的实现十分相似之后才抽象出来的。并且越偏向业务层，抽象难度就越高，这或许也是前面图中 Go 标准库小接口（1~3 个方法）占比略高于 Docker 和 Kubernetes 的原因。

### 4.2 第二，将大接口拆分为小接口

有了接口后，我们就会看到接口被用在了代码的各个地方。一段时间后，我们就来分析哪些场合使用了接口的哪些方法，是否可以将这些场合使用的接口的方法提取出来，放入一个新的小接口中，就像下面图示中的那样：

![WechatIMG268](https://billy.taoxiaoxin.club/md/2023/11/6549045dc8008f9b8063e8f0.jpg)

这张图中的大接口 1 定义了多个方法，一段时间后，我们发现方法 1 和方法 2 经常用在场合 1 中，方法 3 和方法 4 经常用在场合 2 中，方法 5 和方法 6 经常用在场合 3 中，大接口 1 的方法呈现出一种按业务逻辑自然分组的状态。

这个时候我们可以将这三组方法分别提取出来放入三个小接口中，也就是将大接口 1 拆分为三个小接口 A、B 和 C。拆分后，原应用场合 1~3 使用接口 1 的地方就可以无缝替换为接口 A、B、C 了。

### 4.3 最后，我们要注意接口的单一契约职责

那么，上面已经被拆分成的小接口是否需要进一步拆分，直至每个接口都只有一个方法呢？这个依然没有标准答案，不过你依然可以考量一下现有小接口是否需要满足单一契约职责，就像 `io.Reader` 那样。如果需要，就可以进一步拆分，提升抽象程度。