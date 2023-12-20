# Go 方法介绍，理解“方法”的本质

[TOC]



## 一、认识 Go 方法

### 1.1 基本介绍

我们知道，Go 语言从设计伊始，就不支持经典的面向对象语法元素，比如类、对象、继承，等等，但 Go 语言仍保留了名为“方法（`method`）”的语法元素。当然，Go 语言中的方法和面向对象中的方法并不是一样的。Go 引入方法这一元素，并不是要支持面向对象编程范式，而是 Go 践行组合设计哲学的一种实现层面的需要。

在 Go 编程语言中，方法是与特定类型相关联的函数。它们允许您在自定义类型上定义行为，这个自定义类型可以是结构体（struct）或任何用户定义的类型。方法本质上是一种函数，但它们具有一个特定的接收者（receiver），也就是方法所附加到的类型。这个接收者可以是指针类型或值类型。方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。

### 1.2 声明

### 1.2.1 引入

首先我们这里以 Go 标准库 `net/http` 包中 `*Server` 类型的方法 `ListenAndServeTLS` 为例，讲解一下 Go 方法的一般形式：

![img](https://billy.taoxiaoxin.club/md/2023/11/654632378fd9545ed33eaebc.jpg)

和 Go 函数一样，Go 的方法也是以 `func` 关键字修饰的，并且和函数一样，也包含方法名（对应函数名）、参数列表、返回值列表与方法体（对应函数体）。

而且，方法中的这几个部分和函数声明中对应的部分，在形式与语义方面都是一致的，比如：方法名字首字母大小写决定该方法是否是导出方法；方法参数列表支持变长参数；方法的返回值列表也支持具名返回值等。

不过，它们也有不同的地方。从上面这张图我们可以看到，和由五个部分组成的函数声明不同，**Go 方法的声明有六个组成部分**，多的一个就是图中的 `receiver` 部分。在 `receiver` 部分声明的参数，Go 称之为 `receiver` 参数，**这个 `receiver` 参数也是方法与类型之间的纽带，也是方法与函数的最大不同。**

Go 中的方法必须是归属于一个类型的，而 `receiver` 参数的类型就是这个方法归属的类型，或者说这个方法就是这个类型的一个方法。以图中的 `ListenAndServeTLS` 为例，这里的 `receiver` 参数 `srv` 的类型为 `*Server`，那么我们可以说，这个方法就是 `*Server` 类型的方法。

注意！这里说的是 `ListenAndServeTLS` 是 `*Server` 类型的方法，而不是 `Server` 类型的方法。

### 1.2.2 一般声明形式

方法的声明形式如下：

```go
func (t *T或T) MethodName(参数列表) (返回值列表) {
    // 方法体
}
```

其中各部分的含义如下：

- `(t *T或T)`：括号中的部分是方法的接收者，用于指定方法将附加到的类型。`t` 是接收者的名称，`T` 是接收者的类型。接收者可以是值类型（`T`）或指针类型（`*T`）。如果使用值类型作为接收者，方法操作的是接收者的副本，而指针类型允许方法修改接收者的原始值。无论 `receiver` 参数的类型为 `*T` 还是 `T`，我们都把一般声明形式中的 `T` 叫做 `receiver` 参数 `t` 的基类型。如果 `t` 的类型为 `T`，那么说这个方法是类型 `T` 的一个方法；如果 `t` 的类型为 `*T`，那么就说这个方法是类型 `*T` 的一个方法。而且，要注意的是，每个方法只能有一个 `receiver` 参数，Go 不支持在方法的 `receiver` 部分放置包含多个 `receiver` 参数的参数列表，或者变长 `receiver` 参数。
- `MethodName`：这是方法的名称，用于在调用方法时引用它。
- `(参数列表)`：这是方法的参数列表，定义了方法可以接受的参数。如果方法不需要参数，此部分为空。
- `(返回值列表)`：这是方法的返回值列表，定义了方法返回的结果。如果方法不返回任何值，此部分为空。
- 方法体：方法体包含了方法的具体实现，这里可以编写方法的功能代码。

### 1.2.3 receiver 参数作用域

**方法接收器（receiver）参数、函数 / 方法参数，以及返回值变量对应的作用域范围，都是函数 / 方法体对应的显式代码块。**

这就意味着，`receiver` 部分的参数名不能与方法参数列表中的形参名，以及具名返回值中的变量名存在冲突，必须在这个方法的作用域中具有唯一性。如果不唯一，比如下面的例子中那样，Go 编译器就会报错：

```go
type T struct{}

func (t T) M(t string) { // 编译器报错：duplicate argument t (重复声明参数t)
    ... ...
}
```

不过，如果在方法体中没有使用 receiver 参数，我们也可以省略 receiver 的参数名，就像下面这样：

```go
type T struct{}

func (T) M(t string) { 
    ... ...
}
```

仅当方法体中的实现不需要 receiver 参数参与时，我们才会省略 receiver 参数名，不过这一情况很少使用，了解一下即可。

### 1.2.4 receiver 参数的基类型约束

Go 语言对 receiver 参数的基类型也有约束，那就是 **receiver 参数的基类型本身不能为指针类型或接口类型。**

下面的例子分别演示了基类型为指针类型和接口类型时，Go 编译器报错的情况：

```go
type MyInt *int
func (r MyInt) String() string { // r的基类型为MyInt，编译器报错：invalid receiver type MyInt (MyInt is a pointer type)
    return fmt.Sprintf("%d", *(*int)(r))
}

type MyReader io.Reader
func (r MyReader) Read(p []byte) (int, error) { // r的基类型为MyReader，编译器报错：invalid receiver type MyReader (MyReader is an interface type)
    return r.Read(p)
}
```

### 1.2.5 方法声明的位置约束

Go 要求，**方法声明要与 receiver 参数的基类型声明放在同一个包内**。基于这个约束，我们还可以得到两个推论。

+ 第一个推论：**我们不能为原生类型（例如 int、float64、map 等）添加方法**。例如，下面的代码试图为 Go 原生类型 `int` 增加新方法 `Foo`，这是不允许的，Go 编译器会报错：

```go
func (i int) Foo() string { // 编译器报错：cannot define new methods on non-local type int
    return fmt.Sprintf("%d", i) 
}
```

+ 第二个推论：**不能跨越 Go 包为其他包的类型声明新方法**。例如，下面的代码试图跨越包边界，为 Go 标准库中的 `http.Server` 类型添加新方法 `Foo`，这是不允许的，Go 编译器同样会报错：

```go
import "net/http"

func (s http.Server) Foo() { // 编译器报错：cannot define new methods on non-local type http.Server
}
```

### 1.2.6 如何使用方法

我们直接还是通过一个例子理解一下。如果 receiver 参数的基类型为 T，那么我们说 receiver 参数绑定在 T 上，我们可以通过 *T 或 T 的变量实例调用该方法：

```go
type T struct{}

func (t T) M(n int) {
}

func main() {
    var t T
    t.M(1) // 通过类型T的变量实例调用方法M

    p := &T{}
    p.M(2) // 通过类型*T的变量实例调用方法M
}
```

这段代码中，方法 M 是类型 T 的方法，通过 *T 类型变量也可以调用 M 方法。

## 二、方法的本质

通过以上，我们知道了 Go 的方法与 Go 中的类型是通过 receiver 联系在一起，我们可以为任何非内置原生类型定义方法，比如下面的类型 T：

```go
type T struct { 
    a int
}

func (t T) Get() int {  
    return t.a 
}

func (t *T) Set(a int) int { 
    t.a = a 
    return t.a 
}
```

在Go 中，Go 方法中的原理是将 `receiver` 参数以第一个参数的身份并入到方法的参数列表中。按照这个原理，我们示例中的类型 `T` 和 `*T` 的方法，就可以分别等价转换为下面的普通函数：

```go
// 类型T的方法Get的等价函数
func Get(t T) int {  
    return t.a 
}

// 类型*T的方法Set的等价函数
func Set(t *T, a int) int { 
    t.a = a 
    return t.a 
}
```

**这种等价转换后的函数的类型就是`方法的类型`**。只不过在 Go 语言中，这种等价转换是由 Go 编译器在编译和生成代码时自动完成的。**Go 语言规范中还提供了`方法表达式`（Method Expression）的概念**，可以让我们更充分地理解上面的等价转换。

以上面类型 T 以及它的方法为例，结合前面说过的 Go 方法的调用方式，我们可以得到下面代码：

```go
var t T
t.Get()
(&t).Set(1)
```

我们可以用另一种方式，把上面的方法调用做一个等价替换：

```go
var t T
T.Get(t)
(*T).Set(&t, 1)
```

这种直接以类型名 `T` 调用方法的表达方式，被称为`Method Expression`。通过`Method Expression`这种形式，类型 `T` 只能调用 `T` 的方法集合（Method Set）中的方法，同理类型 `*T` 也只能调用 `*T` 的方法集合中的方法。

我们看到，`Method Expression` 有些类似于 C++ 中的静态方法（Static Method）。在 C++ 中的静态方法使用时，以该 C++ 类的某个对象实例作为第一个参数。而 Go 语言的 `Method Expression` 在使用时，同样以 `receiver` 参数所代表的类型实例作为第一个参数。

这种通过 `Method Expression` 对方法进行调用的方式，与我们之前所做的方法到函数的等价转换是如出一辙的。所以，**Go 语言中的方法的本质就是，一个以方法的 `receiver` 参数作为第一个参数的普通函数。**

而且，`Method Expression` 就是 Go 方法本质的最好体现，因为方法自身的类型就是一个普通函数的类型，我们甚至可以将它作为右值，赋值给一个函数类型的变量，比如下面示例：

```go
func main() {
    var t T
    f1 := (*T).Set // f1的类型，也是*T类型Set方法的类型：func (t *T, int)int
    f2 := T.Get    // f2的类型，也是T类型Get方法的类型：func(t T)int
    fmt.Printf("the type of f1 is %T\n", f1) // the type of f1 is func(*main.T, int) int
    fmt.Printf("the type of f2 is %T\n", f2) // the type of f2 is func(main.T) int
    f1(&t, 3)
    fmt.Println(f2(t)) // 3
}
```

## 三、巧解难题

我们来看一段代码：

```go
package main

import (
    "fmt"
    "time"
)

type field struct {
    name string
}

func (p *field) print() {
    fmt.Println(p.name)
}

func main() {
    data1 := []*field{{"one"}, {"two"}, {"three"}}
    for _, v := range data1 {
        go v.print()
    }

    data2 := []field{{"four"}, {"five"}, {"six"}}
    for _, v := range data2 {
        go v.print()
    }

    time.Sleep(3 * time.Second)
}
```

这段代码在我的多核 macOS 上的运行结果是这样（由于 Goroutine 调度顺序不同，你自己的运行结果中的行序可能与下面的有差异）：

```go
one
two
three
six
six
six
```

**为什么对 data2 迭代输出的结果是三个“six”，而不是 four、five、six？**

我们来分析一下。首先，我们根据 **Go 方法的本质**，也就是一个以方法的 **`receiver` 参数作为第一个参数的普通函数**，对这个程序做个**等价变换**。这里我们利用 `Method Expression` 方式，等价变换后的源码如下：

```go
type field struct {
    name string
}

func (p *field) print() {
    fmt.Println(p.name)
}

func main() {
    data1 := []*field{{"one"}, {"two"}, {"three"}}
    for _, v := range data1 {
        go (*field).print(v)
    }

    data2 := []field{{"four"}, {"five"}, {"six"}}
    for _, v := range data2 {
        go (*field).print(&v)
    }

    time.Sleep(3 * time.Second)
}
```

这段代码中，我们把对 `field` 的方法 `print` 的调用，替换为 `Method Expression` 形式，替换前后的程序输出结果是一致的。但变换后，问题是不是豁然开朗了！我们可以很清楚地看到使用 `go` 关键字启动一个新 Goroutine 时，`Method Expression` 形式的 `print` 函数是如何绑定参数的：

+ 迭代 `data1` 时，由于 `data1` 中的元素类型是 `field` 指针 (`*field`)，因此赋值后 `v` 就是元素地址，与 `print` 的 `receiver` 参数类型相同，每次调用 `(*field).print` 函数时直接传入的 `v` 即可，实际上传入的也是各个 `field` 元素的地址。
+ 迭代 `data2` 时，由于 `data2` 中的元素类型是 `field`（非指针），与 `print` 的 `receiver` 参数类型不同，因此需要将其取地址后再传入 `(*field).print` 函数。这样每次传入的 `&v` 实际上是变量 `v` 的地址，而不是切片 `data2` 中各元素的地址。

在[《Go 的 for 循环，仅此一种》](https://blog.csdn.net/weixin_44621343/article/details/133782073)中，我们学习过 `for range` 使用时应注意的几个问题，其中循环变量复用是关键的一个。这里的 `v` 在整个 `for range` 过程中只有一个，因此 `data2` 迭代完成之后，**`v` 是元素 "six" 的拷贝**。

这样，一旦启动的各个子 goroutine 在 main goroutine 执行到 `Sleep` 时才被调度执行，那么最后的三个 goroutine 在打印 `&v` 时，实际打印的也就是在 `v` 中存放的值 "six"。而前三个子 goroutine 各自传入的是元素 "one"、"two" 和 "three" 的地址，所以打印的就是 "one"、"two" 和 "three" 了。

那么原程序要如何修改，才能让它按我们期望，输出“one”、“two”、“three”、“four”、 “five”、“six”呢？

其实，我们只需要将 field 类型 print 方法的 receiver 类型由 `*field` 改为 `field` 就可以了。我们直接来看一下修改后的代码：

```go
type field struct {
    name string
}

func (p field) print() {
    fmt.Println(p.name)
}

func main() {
    data1 := []*field{{"one"}, {"two"}, {"three"}}
    for _, v := range data1 {
        go v.print()
    }

    data2 := []field{{"four"}, {"five"}, {"six"}}
    for _, v := range data2 {
        go v.print()
    }

    time.Sleep(3 * time.Second)
}
```

修改后的程序的输出结果是这样的（因 Goroutine 调度顺序不同，在你的机器上的结果输出顺序可能会有不同）：

```go
one
two
three
four
five
six
```

