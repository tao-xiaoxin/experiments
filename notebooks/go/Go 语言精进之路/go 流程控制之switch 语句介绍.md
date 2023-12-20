# go 流程控制之switch 语句介绍

[TOC]



## 一、switch语句介绍

### 1.1 认识 switch 语句

我们先通过一个例子来直观地感受一下 `switch` 语句的优点。在一些执行分支较多的场景下，使用 `switch` 分支控制语句可以让代码更简洁，可读性更好。

比如下面例子中的 `readByExt` 函数会根据传入的文件扩展名输出不同的日志，使用 if 语句进行分支控制：

```go
func readByExt(ext string) {
    if ext == "json" {
        println("read json file")
    } else if ext == "jpg" || ext == "jpeg" || ext == "png" || ext == "gif" {
        println("read image file")
    } else if ext == "txt" || ext == "md" {
        println("read text file")
    } else if ext == "yml" || ext == "yaml" {
        println("read yaml file")
    } else if ext == "ini" {
        println("read ini file")
    } else {
        println("unsupported file extension:", ext)
    }
}
```

如果用 `switch` 改写上述例子代码，我们可以这样来写：

```go
func readByExtBySwitch(ext string) {
    switch ext {
    case "json":
        println("read json file")
    case "jpg", "jpeg", "png", "gif":
        println("read image file")
    case "txt", "md":
        println("read text file")
    case "yml", "yaml":
        println("read yaml file")
    case "ini":
        println("read ini file")
    default:
        println("unsupported file extension:", ext)
    }
}
```

从代码呈现的角度来看，针对这个例子，使用 switch 语句的实现要比 if 语句的实现更加简洁紧凑。

简单来说，`readByExtBySwitch` 函数就是将输入参数 `ext` 与每个 `case` 语句后面的表达式做比较，如果相等，就执行这个 `case` 语句后面的分支，然后函数返回。

### 1.2 基本语法

在Go编程语言中，`switch`语句的基本语法如下：

```go
switch initStmt; expr {
    case expr1:
        // 执行分支1
    case expr2:
        // 执行分支2
    case expr3_1, expr3_2, expr3_3:
        // 执行分支3
    case expr4:
        // 执行分支4
    ... ...
    case exprN:
        // 执行分支N
    default: 
        // 执行默认分支
}
```

我们按语句顺序来分析一下:

+ 首先 `switch` 语句第一行由 `switch` 关键字开始，它的后面通常接着一个表达式（`expr`），这句中的 `initStmt` 是一个可选的组成部分。和 `if`、`for` 语句一样，我们可以在 `initStmt` 中通过短变量声明定义一些在 switch 语句中使用的临时变量。
+ 接下来，switch 后面的大括号内是一个个代码执行分支，每个分支以 case 关键字开始，每个 case 后面是一个表达式或是一个逗号分隔的表达式列表。
+ 最后，还有一个以 `default` 关键字开始的特殊分支，被称为**默认分支**。`default` 子句是可选的，如果没有一个`case`子句匹配`expression`的值，将执行`default`子句中的代码块。

最后，我们再来看 switch 语句的执行流程:

+ 首先，`switch` 语句会用 `expr` 的求值结果与各个 `case` 中的表达式结果进行比较，如果发现匹配的 `case`，也就是 `case` 后面的表达式，或者表达式列表中任意一个表达式的求值结果与 `expr` 的求值结果相同，那么就会执行该 case 对应的代码分支，分支执行后，`switch` 语句也就结束了。
+ 如果所有 `case` 表达式都无法与 `expr` 匹配，那么程序就会执行 default 默认分支，并且结束 switch 语句。

## 二、Go语言switch语句中case表达式求值顺序

### 2.1 switch语句中case表达式求值次序介绍

接下来，我们再来看看，在有多个 `case` 执行分支的 `switch` 语句中，Go 是按什么次序对各个 `case` 表达式进行求值，并且与 `switch` 表达式（`expr`）进行比较的？

我们先来看一段示例代码，这是一个一般形式的 switch 语句，为了能呈现 switch 语句的执行次序，以多个输出特定日志的函数作为 switch 表达式以及各个 case 表达式：

```go
func case1() int {
    println("eval case1 expr")
    return 1
}

func case2_1() int {
    println("eval case2_1 expr")
    return 0 
}
func case2_2() int {
    println("eval case2_2 expr")
    return 2 
}

func case3() int {
    println("eval case3 expr")
    return 3
}

func switchexpr() int {
    println("eval switch expr")
    return 2
}

func main() {
    switch switchexpr() {
    case case1():
        println("exec case1")
    case case2_1(), case2_2():
        println("exec case2")
    case case3():
        println("exec case3")
    default:
        println("exec default")
    }
}
```

执行一下这个示例程序，我们得到如下结果：

```go
eval switch expr
eval case1 expr
eval case2_1 expr
eval case2_2 expr
exec case2
```

从输出结果中我们看到，Go 先对 switch expr 表达式进行求值，然后再按 case 语句的出现顺序，从上到下进行逐一求值。在带有表达式列表的 case 语句中，Go 会从左到右，对列表中的表达式进行求值，比如示例中的 case2_1 函数就执行于 case2_2 函数之前。

如果 switch 表达式匹配到了某个 case 表达式，那么程序就会执行这个 case 对应的代码分支，比如示例中的“exec case2”。这个分支后面的 case 表达式将不会再得到求值机会，比如示例不会执行 case3 函数。这里要注意一点，即便后面的 case 表达式求值后也能与 switch 表达式匹配上，Go 也不会继续去对这些表达式进行求值了，这是`switch`语句的工作原理。

除了这一点外，你还要注意 default 分支。**无论 default 分支出现在什么位置，它都只会在所有 case 都没有匹配上的情况下才会被执行的**。

不知道你有没有发现，这里其实有一个优化小技巧，考虑到 switch 语句是按照 case 出现的先后顺序对 case 表达式进行求值的，那么如果我们将匹配成功概率高的 case 表达式排在前面，就会有助于提升 switch 语句执行效率。这点对于 case 后面是表达式列表的语句同样有效，我们可以将匹配概率最高的表达式放在表达式列表的最左侧。

### 2.2 switch语句中case表达式的求值次序特点

Go语言switch语句中case表达式的求值次序特点:

1. switch语句首先求值`switch`表达式,然后按`case`出现顺序逐一求值`case`表达式。
2. 一旦某个`case`表达式匹配成功后,就执行对应的代码块,之后`case`不再求值。
3. 即使后续的`case`表达式匹配成功,也不会再求值。
4. 所有`case`都不匹配的情况下,会执行默认的`default`案例。
5. `default`位置灵活,可以放在开头或结尾。
6. `case`后带表达式列表时,会从左到右求值列表中的表达式。
7. 将匹配概率高的`case`排在前面,可以优化执行效率。

## 三、switch 语句的灵活性

### 3.1 switch 语句各表达式的求值结果支持各种类型值

首先，switch 语句各表达式的求值结果可以为各种类型值，只要它的类型支持比较操作就可以了。

Go 语言只要类型支持比较操作，都可以作为 `switch` 语句中的表达式类型。比如整型、布尔类型、字符串类型、复数类型、元素类型都是可比较类型的数组类型，甚至字段类型都是可比较类型的结构体类型也可以。下面就是一个使用自定义结构体类型作为` switch `表达式类型的例子：

```go
type person struct {
    name string
    age  int
}

func main() {
    p := person{"tom", 13}
    switch p {
    case person{"tony", 33}:
        println("match tony")
    case person{"tom", 13}:
        println("match tom")
    case person{"lucy", 23}:
        println("match lucy")
    default:
        println("no match")
    }
}
```

实际开发过程中，以结构体类型为 `switch`表达式类型的情况并不常见，这里举这个例子仅是为了说明 Go `switch` 语句对各种类型支持的广泛性。

而且，当 switch 表达式的类型为布尔类型时，如果求值结果始终为 true，那么我们甚至可以省略 switch 后面的表达式，比如下面例子：

```go
// 带有initStmt语句的switch语句
switch initStmt; {
    case bool_expr1:
    case bool_expr2:
    ... ...
}

// 没有initStmt语句的switch语句
switch {
    case bool_expr1:
    case bool_expr2:
    ... ...
}
```

**注意**：在带有 initStmt 的情况下，如果我们省略 switch 表达式，那么 initStmt 后面的分号不能省略，因为 initStmt 是一个语句。

### 3.2 switch 语句支持声明临时变量

在前面介绍 switch 语句的一般形式中，我们看到，和 if、for 等控制结构语句一样，switch 语句的 initStmt 可用来声明只在这个 switch 隐式代码块中使用的变量，这种就近声明的变量最大程度地缩小了变量的作用域。

示例：

```go
switch x := someFunction(); x {
case 1:
    fmt.Println("x is 1")
case 2:
    fmt.Println("x is 2")
default:
    fmt.Println("x is something else")
}

// 这里无法访问 x，因为它的作用域仅限于 switch 语句

```

在上面的示例中，`x`是一个局部变量，只在`switch`语句内部可见。这可以有效地限制变量的生存期和可见性，从而提高代码的清晰度和健壮性。这是Go语言在控制结构中的一种好实践。

### 3.3 case 语句支持表达式列表

在Go的`switch`语句中，`case`语句支持表达式列表，一个分支可以有多个值，多个case值中间使用英文逗号分隔。这意味着你可以在一个`case`子句中列出多个表达式，以匹配其中任何一个表达式。如果`switch`表达式的值与列表中的任何一个表达式匹配，相应的`case`分支将被执行。

```go
func checkWorkday(a int) {
    switch a {
    case 1, 2, 3, 4, 5:
        println("it is a work day")
    case 6, 7:
        println("it is a weekend day")
    default:
        println("are you live on earth")
    }
}
```

### 3.4 取消了默认执行下一个 case 代码逻辑的语义

在 C 语言中，如果匹配到的 case 对应的代码分支中没有显式调用 break 语句，那么代码将继续执行下一个 case 的代码分支，这种“隐式语义”并不符合日常算法的常规逻辑，这也经常被诟病为 C 语言的一个缺陷。要修复这个缺陷，我们只能在每个 case 执行语句中都显式调用 break。

**Go 语言中的 Swith 语句就修复了 C 语言的这个缺陷，取消了默认执行下一个 case 代码逻辑的“非常规”语义，每个 case 对应的分支代码执行完后就结束 switch 语句。**

如果在少数场景下，你需要执行下一个 case 的代码逻辑，你可以显式使用 Go 提供的关键字 `fallthrough` 来实现，**`fallthrough`语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的**。下面就是一个使用 fallthrough 的 switch 语句的例子，我们简单来看一下：

```go
func case1() int {
    println("eval case1 expr")
    return 1
}

func case2() int {
    println("eval case2 expr")
    return 2
}

func switchexpr() int {
    println("eval switch expr")
    return 1
}

func main() {
    switch switchexpr() {
    case case1():
        println("exec case1")
        fallthrough
    case case2():
        println("exec case2")
        fallthrough
    default:
        println("exec default")
    }
}
```

执行一下这个示例程序，我们得到这样的结果：

```go
eval switch expr
eval case1 expr
exec case1
exec case2
exec default
```

我们看到，switch expr 的求值结果与 case1 匹配成功，Go 执行了 case1 对应的代码分支。而且，由于 case1 代码分支中显式使用了 fallthrough，执行完 case1 后，代码执行流并没有离开 switch 语句，而是继续执行下一个 case，也就是 case2 的代码分支。

这里有一个注意点，由于 fallthrough 的存在，Go 不会对 case2 的表达式做求值操作，而会直接执行 case2 对应的代码分支。而且，在这里 case2 中的代码分支也显式使用了 fallthrough，于是最后一个代码分支，也就是 default 分支对应的代码也被执行了。

另外，还有一点要注意的是，如果某个 case 语句已经是 switch 语句中的最后一个 case 了，并且它的后面也没有 default 分支了，那么这个 case 中就不能再使用 fallthrough，否则编译器就会报错。

到这里，我们看到 Go 的 switch 语句不仅修复了 C 语言 switch 的缺陷，还为 Go 开发人员提供了更大的灵活性，我们可以使用更多类型表达式作为 switch 表达式类型，也可以使用 case 表达式列表简化实现逻辑，还可以自行根据需要，确定是否使用 fallthrough 关键字继续向下执行下一个 case 的代码分支。

## 四、type switch

“type switch”这是一种特殊的 switch 语句用法，我们通过一个例子来看一下它具体的使用形式：

```go
func main() {
    var x interface{} = 13
    switch x.(type) {
    case nil:
        println("x is nil")
    case int:
        println("the type of x is int")
    case string:
        println("the type of x is string")
    case bool:
        println("the type of x is string")
    default:
        println("don't support the type")
    }
}
```

我们看到，这个例子中 switch 语句的形式与前面是一致的，不同的是 switch 与 case 两个关键字后面跟着的表达式。

switch 关键字后面跟着的表达式为` x.(type) `，这种表达式形式是 switch 语句专有的，而且也只能在 switch 语句中使用。**这个表达式中的 x 必须是一个接口类型变量**，表达式的求值结果是这个接口类型变量对应的动态类型。

什么是一个接口类型的动态类型呢？我们简单解释一下。以上面的代码 `var x interface{} = 13` 为例，x 是一个接口类型变量，它的静态类型为` interface{} `，如果我们将整型值 13 赋值给 x，x 这个接口变量的动态类型就为 int。关于接口类型变量的动态类型，我们后面还会详细讲，这里先简单了解一下就可以了。

接着，case 关键字后面接的就不是普通意义上的表达式了，而是一个个具体的类型。这样，Go 就能使用变量 x 的动态类型与各个 case 中的类型进行匹配，之后的逻辑就都是一样的了。

现在我们运行上面示例程序，输出了 x 的动态变量类型：

```go
the type of x is int
```

不过，通过 `x.(type) `，我们除了可以获得变量 x 的动态类型信息之外，也能获得其动态类型对应的值信息，现在我们把上面的例子改造一下：

```go
func main() {
    var x interface{} = 13
    switch v := x.(type) {
    case nil:
        println("v is nil")
    case int:
        println("the type of v is int, v =", v)
    case string:
        println("the type of v is string, v =", v)
    case bool:
        println("the type of v is bool, v =", v)
    default:
        println("don't support the type")
    }
}
```

这里我们将 switch 后面的表达式由` x.(type)` 换成了 `v := x.(type) `。对于后者，你千万不要认为变量 v 存储的是类型信息，**其实 v 存储的是变量 x 的动态类型对应的值信息**，这样我们在接下来的 case 执行路径中就可以使用变量 v 中的值信息了。

然后，我们运行上面示例，可以得到 v 的动态类型和值：

```go
the type of v is int, v = 13
```

另外，你可以发现，在前面的 type switch 演示示例中，我们一直使用 interface{}这种接口类型的变量，Go 中所有类型都实现了 interface{}类型，所以 case 后面可以是任意类型信息。

但如果在 switch 后面使用了某个特定的接口类型 I，那么 case 后面就只能使用实现了接口类型 I 的类型了，否则 Go 编译器会报错。你可以看看这个例子：

```go
  type I interface {
      M()
  }
  
  type T struct {
  }
  
 func (T) M() {
 }
 
 func main() {
     var t T
     var i I = t
     switch i.(type) {
     case T:
         println("it is type T")
     case int:
         println("it is type int")
     case string:
         println("it is type string")
     }
 }
```

在这个例子中，我们在 `type switch` 中使用了自定义的接口类型 I。那么，理论上所有 case 后面的类型都只能是实现了接口 I 的类型。但在这段代码中，只有类型 T 实现了接口类型 I，Go 原生类型 int 与 string 都没有实现接口 I，于是在编译上述代码时，编译器会报出如下错误信息：

```go
19:2: impossible type switch case: i (type I) cannot have dynamic type int (missing M method)
21:2: impossible type switch case: i (type I) cannot have dynamic type string (missing M method)
```

## 五、跳不出循环的 break

这里，我们来看一个找出整型切片中第一个偶数的例子，使用 switch 分支结构：

```go
func main() {
    var sl = []int{5, 19, 6, 3, 8, 12}
    var firstEven int = -1

    // find first even number of the interger slice
    for i := 0; i < len(sl); i++ {
        switch sl[i] % 2 {
        case 0:
            firstEven = sl[i]
            break
        case 1:
            // do nothing
        }        
    }         
    println(firstEven) 
}
```

我们运行一下这个修改后的程序，得到结果为 12。

奇怪，这个输出的值与我们的预期的好像不太一样。这段代码中，切片中的第一个偶数是 6，而输出的结果却成了切片的最后一个偶数 12。为什么会出现这种结果呢？

这就是 Go 中 break 语句与 switch 分支结合使用会出现一个“小坑”。和我们习惯的 C 家族语言中的 break 不同，Go 语言规范中明确规定，**不带 label 的 break 语句中断执行并跳出的，是同一函数内 break 语句所在的最内层的 for、switch 或 select**。所以，上面这个例子的 break 语句实际上只跳出了 switch 语句，并没有跳出外层的 for 循环，这也就是程序未按我们预期执行的原因。

要修正这一问题，我们可以利用 `label` 的 `break` 语句试试。这里我们也直接看看改进后的代码:

```go
func main() {
    var sl = []int{5, 19, 6, 3, 8, 12}
    var firstEven int = -1

    // find first even number of the interger slice
loop:
    for i := 0; i < len(sl); i++ {
        switch sl[i] % 2 {
        case 0:
            firstEven = sl[i]
            break loop
        case 1:
            // do nothing
        }
    }
    println(firstEven) // 6
}
```

在改进后的例子中，我们定义了一个 label：loop，这个 label 附在 for 循环的外面，指代 for 循环的执行。当代码执行到“break loop”时，程序将停止 label loop 所指代的 for 循环的执行。

## 六、switch与if 比较

Go编程语言中的`switch`语句和`if`语句是用于控制程序流程的两个不同工具，它们可以用来执行条件性代码块，但它们在使用方式和适用场景上有所不同。

### 相似之处：

- `if`语句和`switch`语句都用于根据某个条件执行不同的代码块。
- 两者都可以用于处理多个条件或值的情况。

### 不同之处：

- `if`语句通常用于处理更复杂的条件逻辑，可以检查任何布尔表达式。它是通用的条件控制工具。
- `switch`语句专门用于根据一个表达式的值选择执行不同的代码块。它通常用于在多个值之间进行精确的比较。

在`if`语句中，你可以编写任意复杂的条件，例如：

```go
if condition1 {
    // 当condition1为真时执行这里的代码
} else if condition2 {
    // 当condition2为真时执行这里的代码
} else {
    // 如果以上条件都不为真，执行这里的代码
}
```

而在`switch`语句中，你主要是根据某个表达式的值进行选择，比较简洁：

```go
switch expression {
case value1:
    // 当expression等于value1时执行这里的代码
case value2:
    // 当expression等于value2时执行这里的代码
default:
    // 如果expression不等于任何一个value，执行这里的代码
}
```

使用`if`语句更适合处理复杂的条件逻辑，而`switch`语句更适合在多个值之间进行简单的比较。

