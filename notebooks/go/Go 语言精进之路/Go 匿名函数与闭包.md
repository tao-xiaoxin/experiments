# Go 匿名函数与闭包

匿名函数和闭包是一些编程语言中的重要概念，它们在Go语言中也有重要的应用。让我们来详细介绍这两个概念，并提供示例代码来帮助理解。

[TOC]



## 一、匿名函数（Anonymous Function）

匿名函数，也称为**无名函数**，是一种没有名字的函数，它通常用于一次性的、小规模的操作。匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或者作为立即执行函数：

```go
package main

import "fmt"

func main() {
    // 自执行函数：匿名函数定义完加()直接执行
    result := func(x, y int) int {
        return x + y
    }(3, 4)
    fmt.Println("Result:", result)

    // 将匿名函数分配给变量
    add := func(x, y int) int {
        return x + y
    }

    // 使用分配给变量的匿名函数
    sum := add(5, 6)
    fmt.Println("Sum:", sum)
}

```

匿名函数多用于实现回调函数和闭包。

## 二、闭包函数（Closure）

闭包是指一个函数，它包含对其外部作用域的变量的引用。这意味着闭包可以访问并操作其外部作用域中的变量，即使在外部函数已经返回后仍然可以操作这些变量。

在Go语言中，匿名函数通常用作闭包。当一个匿名函数引用外部作用域的变量时，它形成了一个闭包。闭包可以用于捕获状态、实现回调函数等。即：**`闭包=函数+引用环境`**。下面是一个示例：

```go
package main

import "fmt"

func main() {
    // 外部函数返回一个匿名函数
    funcWithClosure := func() func(int) int {
        sum := 0
        return func(x int) int {
            sum += x
            return sum
        }
    }()

    // 使用匿名函数创建闭包
    fmt.Println(funcWithClosure(1)) // 1
    fmt.Println(funcWithClosure(2)) // 3
    fmt.Println(funcWithClosure(3)) // 6
}
```

在这个示例中，我们定义了一个外部函数`funcWithClosure`，它返回一个匿名函数。这个匿名函数形成了一个闭包，它可以访问并修改外部函数中的`sum`变量。因此，每次调用匿名函数时，`sum`的值都会被累加。

闭包在Go中常用于实现函数工厂、状态管理和回调等情况，因为它们可以捕获和保持状态信息。