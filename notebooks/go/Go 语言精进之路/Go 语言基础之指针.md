# Go语言中的指针

[TOC]



## 一、Go语言中的指针介绍

### 1.1 指针介绍

指针是一个存储变量内存地址的变量。它们允许程序直接访问和操作内存中的数据，而不是对数据的副本进行操作。以下是指针的一些关键概念：

- **内存地址：** 每个变量在计算机内存中都有一个唯一的地址，指针存储了这个地址。
- **指针变量：** 用于存储其他变量地址的变量称为指针变量。
- **取地址操作符（&）：** 可以使用取地址操作符`&`来获取变量的地址。
- **解引用操作符（\*）：** 可以使用解引用操作符`*`来访问指针所指向的变量的值。

Go语言中的值类型（int、float、bool、string、array、struct）都有对应的指针类型，如：`*int`、`*int64`、`*string`等。

### 1.2 基本语法

- `var ptr *int:` 声明指针变量ptr,用于指向一个int类型变量的地址。
- `&a`: **获取变量a的内存地址**,返回一个指向该地址的指针。
- `*ptr`: 读取ptr指针指向地址的值,这个操作称为“解引用”。
- `*ptr = 100`: 将100赋值给ptr指向的变量。

### 1.3 声明和初始化

在 Go 语言中，可以使用指针来引用任何类型的变量。指针的声明和初始化可以通过如下语法完成：

```csharp
var p *int  // 声明一个指向 int 类型的指针 p
var str *string  // 声明一个指向 string 类型的指针 str
```

初始化指针可以通过 new 函数来分配内存并返回指针的地址：

```go
p := new(int)  // 分配一个 int 类型的内存，并将指针 p 指向该内存
```

示例代码：

```go
package main
 
 import "fmt"
 
 func main() {
     var p *int
     var str *string
 
     fmt.Printf("p: %v, str: %v\n", p, str) // 输出 p: <nil>, str: <nil>
 
     x := 10
     p = &x // 将指针p指向变量x的地址
 
     fmt.Printf("p: %v\n", p)   // 输出 p: 0xc0000100e0
     fmt.Printf("*p: %d\n", *p) // 输出 *p: 10
 
     str = new(string) // 分配一个string类型的内存，并将指针str指向该内存
 
     fmt.Printf("str: %v\n", str)   // 输出 str: 0xc000010120
     fmt.Printf("*str: %s\n", *str) // 输出 *str: ""
 
     *str = "Hello, Go!" // 通过指针修改字符串的值
 
     fmt.Printf("*str: %s\n", *str) // 输出 *str: Hello, Go!
 }
```

### 1.4 Go 指针的3个重要概念

#### 1.4.1 指针地址（Pointer Address）

+ 在Go语言中，指针地址表示**指针所指向的变量或数据在内存中的位置**。
+ 在Go语言中，与C/C++等语言不同，您不能直接获取指针的具体地址值，因为Go语言为了安全性和内存管理而采用了更抽象的设计。但是，您可以通过获取变量的地址来创建和使用指针，而这个地址由Go语言自动管理。

#### 1.4.2 指针类型（Pointer Type）

+ Go语言的**指针类型表示指针可以指向的数据类型**。

#### 1.4.3 指针取值（Pointer Dereferencing）

+ **指针取值是指通过指针来访问其所指向的内存位置上的数据**。在Go语言中，要获取指针所指向的数据的值，您需要使用解引用操作符 `*`。

### 1.5 获取指针的地址和解引用

通过 & 操作符可以获取变量的地址，例如：

```go
func main() {
	a := 10
	b := &a  // 将指针 b 指向变量 a 的地址
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
	fmt.Println(&b)                    // 0xc00000e018
}
```

我们来看一下`b := &a`的图示：![取变量地址图示](https://billy.taoxiaoxin.club/md/2023/10/651ffbe0d0b2dddea54256f8.png)

使用` * `操作符可以解引用指针，获取指针指向的值：

```go
 fmt.Println(*b)  // 输出指针 b 指向的值，即变量 a 的值
```

示例代码：

```go
func main() {
	//指针取值
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
}
```

输出如下：

```go
type of b:*int
type of c:int
value of c:10
```

**总结：** 取地址操作符`&`和取值操作符`*`是一对互补操作符，`&`取出地址，`*`根据地址取出地址指向的值。

变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：

- 对变量进行取地址（&）操作，可以获得这个变量的指针变量。
- 指针变量的值是指针地址。
- 对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。

### 1.6 传递指针给函数

您可以将指针作为参数传递给函数，从而可以在函数内部修改原始变量的值，而不是复制。这可以用于实现函数的副作用。

```go
func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func main() {
	a := 10
	modify1(a)
	fmt.Println(a) // 10
	modify2(&a)
	fmt.Println(a) // 100
}
```

### 1.7 指针的比较

您可以使用`==`和`!=`运算符来比较指针。它们将比较指针是否引用相同的内存地址。

```go
var x int = 42
var p *int  // 声明一个整数指针
p = &x      // 将变量x的地址分配给指针p
fmt.Println(p == &x) // true，p和&x都指向相同的内存地址
```

### 1.8 指针的使用注意事项

- 谨慎使用指针，以避免悬挂指针（dangling pointers）和内存泄漏等问题。
- 在Go中，指针通常用于传递大型数据结构，以避免复制数据。
- Go没有指针运算（如C/C++中的指针算术运算），因此您不能像C/C++那样执行指针加法和减法操作。

## 二、空指针和指针的零值

- **指针的零值：**如果您声明了一个指针但没有初始化它，它将具有零值，即`nil`。
- **空指针：**如果指针没有指向任何有效的内存地址，它将具有`nil`值，表示空指针。在使用指针之前，通常会检查指针是否为`nil`。

```go
package main

import "fmt"

func main() {
    var p *string
    fmt.Println(p)
    fmt.Printf("p的值是%s/n", p)
    if p != nil {
        fmt.Println("非空")
    } else {
        fmt.Println("空值")
    }
}
```

## 三、指针的应用场景

### 3.1 传递大对象

在函数参数传递时，如果直接传递大对象的副本，会产生额外的内存开销。通过传递指针，可以避免复制整个对象，提高程序的性能。

示例代码：

```go
 package main
 
 import "fmt"
 
 type BigObject struct {
     // 大对象的定义...
 }
 
 func processObject(obj *BigObject) {
     // 对大对象进行处理...
 }
 
 func main() {
     obj := BigObject{}
     processObject(&obj) // 传递大对象的指针
 }
```

### 3.2 指针作为函数参数和修改函数外部变量

在 Go 语言中，函数的参数传递默认是值传递。通过指针传递，函数可以修改函数外部的变量。这在需要修改外部变量的值时非常有用，特别是在处理复杂数据结构或需要对全局状态进行修改的情况下。

示例代码：

```go
 package main
 
 import "fmt"
 
 func modifyValue(ptr *int) {
     *ptr = 30 // 修改指针指向的值
 }
 
 func main() {
     x := 10
     modifyValue(&x) // 传递x的地址给modifyValue函数
     fmt.Println(x) // 输出修改后的x的值，即30
 }
```

### 3.3 动态分配内存

指针的另一个重要应用是动态分配内存。通过 new 函数可以在堆上动态分配内存，避免了在栈上分配固定大小的内存空间的限制。这对于需要返回动态分配的数据或创建复杂数据结构非常有用。

示例代码：

```go
 package main
 
 import "fmt"
 
 type ComplexStruct struct {
     // 复杂数据结构的定义...
 }
 
 func createComplexStruct() *ComplexStruct {
     cs := new(ComplexStruct) // 动态分配内存并返回指针
     // 初始化复杂数据结构...
     return cs
 }
 
 func main() {
     obj := createComplexStruct()
     // 对动态分配的数据结构进行操作...
 }
```

### 3.4 函数返回指针

在函数中返回指针可以将函数内部创建的变量的地址传递给调用者。这样做可以避免复制整个变量，并允许调用者直接访问和修改函数内部的数据。

示例代码：

```go
 package main
 
 import "fmt"
 
 func createValue() *int {
     x := 10 // 在函数内部创建变量
     return &x // 返回变量的地址
 }
 
 func main() {
     p := createValue()
     fmt.Println(*p) // 输出通过指针访问的函数内部变量的值，即10
 }
```

## 四、new和make

我们先来看一个例子：

```go
func main() {
	var a *int
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b["测试"] = 100
	fmt.Println(b)
}
```

执行上面的代码会引发panic，为什么呢？ 在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。要分配内存，就引出来今天的new和make。 Go语言中new和make是内建的两个函数，主要用来分配内存。

### 4.1 new

new是一个内置的函数，它的函数签名如下：

```go
func new(Type) *Type
```

其中，

- Type表示类型，new函数只接受一个参数，这个参数是一个类型
- *Type表示类型指针，new函数返回一个指向该类型内存地址的指针。

new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值。举个例子：

```go
func main() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)       // false
}	
```

本节开始的示例代码中`var a *int`只是声明了一个指针变量a但是没有初始化，指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值。应该按照如下方式使用内置的new函数对a进行初始化之后就可以正常对其赋值了：

```go
func main() {
	var a *int
	a = new(int)
	*a = 10
	fmt.Println(*a)
}
```

### 4.2 make

make也是用于内存分配的，区别于new，它只用于slice、map以及channel的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。make函数的函数签名如下：

```go
func make(t Type, size ...IntegerType) Type
```

make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作。这个我们在上一章中都有说明，关于channel我们会在后续的章节详细说明。

本节开始的示例中`var b map[string]int`只是声明变量b是一个map类型的变量，需要像下面的示例代码一样使用make函数进行初始化操作之后，才能对其进行键值对赋值：

```go
func main() {
    var b map[string]int
    b = make(map[string]int, 10)
    b["测试"] = 100
    fmt.Println(b)
}
```

### 4.3 new与make的区别

1. 二者都是用来做内存分配的。
2. make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
3. 而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。