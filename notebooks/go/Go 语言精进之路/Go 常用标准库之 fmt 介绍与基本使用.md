# Go 常用标准库之 fmt 介绍与基本使用

[TOC]



## 一、介绍

fmt 是 Go 语言中的一个常用标准库，它用于格式化输入和输出数据。fmt 包提供了一系列函数，可以帮助你将数据以特定的格式打印到标准输出（通常是终端）或将数据格式化为字符串以供后续处理。这个库的名称 "fmt" 来自于 "format"，因为它主要用于格式化数据。

fmt 包的主要功能包括：

1. **格式化输出**：fmt 包提供了函数如 `Print`, `Printf`, `Println`, `Fprint`, `Fprintf`, 和 `Fprintln` 用于将数据输出到标准输出或指定的 io.Writer。你可以使用这些函数将数据以不同的格式打印到屏幕上或文件中。
2. **格式化输入**：fmt 包也支持从输入源（通常是标准输入）读取数据，并根据格式规范解析数据。这是通过 `Scan`, `Scanf`, 和 `Scanln` 函数实现的。这对于从用户获取输入数据非常有用。
3. **字符串格式化**：你可以使用 `Sprintf` 函数将数据格式化为字符串而不是直接输出到标准输出，这对于构建日志消息或其他需要格式化的字符串很有用。
4. **错误格式化**：fmt 包也提供了 `Errorf` 函数，用于将格式化的错误消息作为 error 类型返回，方便错误处理。
5. **格式化占位符**：在格式化字符串中，你可以使用占位符来指定如何格式化数据。常见的占位符包括 `%d`（整数），`%f`（浮点数），`%s`（字符串）等。

## 二、向外输出

标准库 `fmt` 提供了多种用于输出的函数，每个函数都有不同的用途和输出方式。以下是一些常用的输出相关函数：

### 2.1 Print 系列

+ `Print`：用于将文本输出到标准输出。它接受任意数量的参数，并将它们串联成一个字符串输出，不会添加换行符。
+ `Printf`：用于格式化输出到标准输出。它接受一个格式化字符串和一系列参数，根据格式化字符串的占位符将参数格式化并输出。
+ `Println`：类似于 `Print`，但会在**输出后自动添加一个换行符**。

```go
func Print(a ...interface{}) (n int, err error)
func Printf(format string, a ...interface{}) (n int, err error)
func Println(a ...interface{}) (n int, err error)
```

举个简单的例子：

~~~go
func main() {
	fmt.Print("Hello, ", "world")
	name := "Alice"
	age := 30
	fmt.Printf("Hello, %s. You are %d years old.\n", name, age)
	fmt.Println("Hello, world")
}
~~~

### 2.2 Fprint 系列

`Fprint` 系列函数用于将文本输出到指定的 `io.Writer` 接口，而不仅仅是标准输出。你可以将文本输出到文件、网络连接等。这些函数的参数列表包括一个 `io.Writer` 参数，以及任意数量的参数。

- `Fprint`：将文本输出到指定的 `io.Writer`。
- `Fprintf`：将格式化文本输出到指定的 `io.Writer`。
- `Fprintln`：将带有换行符的文本输出到指定的 `io.Writer`。

```go
func Fprint(w io.Writer, a ...interface{}) (n int, err error)
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
```

举个例子：

~~~go
func main() {
	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./output.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name := "jarvis"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)
}
~~~

这个示例创建了一个名为 "output.txt" 的文件，并将数据写入文件中。

### 2.3 Sprint 系列

`Sprint` 系列函数用于将文本输出到字符串中，而不是标准输出或文件。它们将文本格式化为字符串并返回结果。

- `Sprint`：将文本输出到字符串。
- `Sprintf`：将格式化文本输出到字符串。
- `Sprintln`：将带有换行符的文本输出到字符串。

```go
func Sprint(a ...interface{}) string
func Sprintf(format string, a ...interface{}) string
func Sprintln(a ...interface{}) string
```

简单的示例代码如下：

```go
func main() {
	s1 := fmt.Sprint("jarvis")
	name := "jarvis"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("jarvis")
	fmt.Println(s1, s2, s3)
}
```

###  2.4 Errorf 系列

`Errorf` 系列函数用于创建格式化的错误消息并返回一个 `error` 类型的值。这允许你将格式化的错误消息返回给调用者，以便更好地进行错误处理。这些函数的用法类似于 `Sprintf`，但它们返回一个 `error` 值而不是字符串。

- `Errorf`：根据`format`参数生成格式化字符串并返回一个包含该字符串的错误。

```go
func Errorf(format string, a ...interface{}) error
```

通常使用这种方式来自定义错误类型，例如：

~~~go
err := fmt.Errorf("这是一个错误")
~~~

## 三、格式化占位符

`*printf`系列函数都支持format格式化参数，在这里我们按照占位符将被替换的变量类型划分，方便查询和记忆。

### 3.1 通用占位符

通用占位符用于格式化不同类型的数据：

| 占位符 |                说明                |
| :----: | :--------------------------------: |
|   %v   |          值的默认格式表示          |
|  %+v   | 类似%v，但输出结构体时会添加字段名 |
|  %#v   |           值的Go语法表示           |
|   %T   |            打印值的类型            |
|   %%   |               百分号               |

代码示例：

```go
func main() {
	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	o := struct{ name string }{"jarvis"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%\n")
}
```

### 3.2 布尔型

| 占位符 |    说明     |
| :----: | :---------: |
|   %t   | true或false |

### 3.3 整型

| 占位符 |                             说明                             |
| :----: | :----------------------------------------------------------: |
|   %b   |                         表示为二进制                         |
|   %c   |                    该值对应的unicode码值                     |
|   %d   |                         表示为十进制                         |
|   %o   |                         表示为八进制                         |
|   %x   |                   表示为十六进制，使用a-f                    |
|   %X   |                   表示为十六进制，使用A-F                    |
|   %U   |          表示为Unicode格式：U+1234，等价于"U+%04X"           |
|   %q   | 该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示 |

示例代码如下：

```go
n := 65
fmt.Printf("%b\n", n)
fmt.Printf("%c\n", n)
fmt.Printf("%d\n", n)
fmt.Printf("%o\n", n)
fmt.Printf("%x\n", n)
fmt.Printf("%X\n", n)
```

### 3.4 浮点数与复数

| 占位符 | 说明                                                   |
| ------ | ------------------------------------------------------ |
| %b     | 无小数部分、二进制指数的科学计数法，如-123456p-78      |
| %e     | 科学计数法，如-1234.456e+78                            |
| %E     | 科学计数法，如-1234.456E+78                            |
| %f     | 有小数部分但无指数部分，如123.456                      |
| %F     | 等价于%f                                               |
| %g     | 根据实际情况采用%e或%f格式（以获得更简洁、准确的输出） |
| %G     | 根据实际情况采用%E或%F格式（以获得更简洁、准确的输出） |

示例代码如下：

```go
f := 12.34
fmt.Printf("%b\n", f)
fmt.Printf("%e\n", f)
fmt.Printf("%E\n", f)
fmt.Printf("%f\n", f)
fmt.Printf("%g\n", f)
fmt.Printf("%G\n", f)
```

### 3.5  字符串和[]byte

| 占位符 | 说明                                                         |
| ------ | ------------------------------------------------------------ |
| %s     | 直接输出字符串或者[]byte                                     |
| %q     | 该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示 |
| %x     | 每个字节用两字符十六进制数表示（使用a-f                      |
| %X     | 每个字节用两字符十六进制数表示（使用A-F）                    |

示例代码如下：

```go
    s := "jarvis"
    fmt.Printf("%s\n", s)
    fmt.Printf("%q\n", s)
    fmt.Printf("%x\n", s)
    fmt.Printf("%X\n", s)
```

### 3.6 指针

| 占位符 | 说明                           |
| ------ | ------------------------------ |
| %p     | 表示为十六进制，并加上前导的0x |

示例代码如下：

```go
a := 18
fmt.Printf("%p\n", &a)
fmt.Printf("%#p\n", &a)
```

### 3.7 宽度标识符

宽度通过一个紧跟在百分号后面的十进制数指定，如果未指定宽度，则表示值时除必需之外不作填充。精度通过（可选的）宽度后跟点号后跟的十进制数指定。如果未指定精度，会使用默认精度；如果点号后没有跟数字，表示精度为0。举例如下

| 占位符 | 说明               |
| ------ | ------------------ |
| %f     | 默认宽度，默认精度 |
| %9f    | 宽度9，默认精度    |
| %.2f   | 默认宽度，精度2    |
| %9.2f  | 宽度9，精度2       |
| %9.f   | 宽度9，精度0       |

示例代码如下：

```go
n := 88.88
fmt.Printf("%f\n", n)
fmt.Printf("%9f\n", n)
fmt.Printf("%.2f\n", n)
fmt.Printf("%9.2f\n", n)
fmt.Printf("%9.f\n", n)
```

### 3.8 其他flag

| 占位符 |                             说明                             |
| :----: | :----------------------------------------------------------: |
|  ‘+’   | 总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）； |
|  ’ '   | 对数值，正数前加空格而负数前加负号；对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格 |
|  ‘-’   | 在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）； |
|  ‘#’   | 八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）对%q（%#q），对%U（%#U）会输出空格和单引号括起来的go字面值； |
|  ‘0’   | 使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面； |

举个例子：

```go
s := "jarvis"
fmt.Printf("%s\n", s)
fmt.Printf("%5s\n", s)
fmt.Printf("%-5s\n", s)
fmt.Printf("%5.7s\n", s)
fmt.Printf("%-5.7s\n", s)
fmt.Printf("%5.2s\n", s)
fmt.Printf("%05s\n", s)
```

## 四、获取输入

Go 语言的 `fmt` 包提供了 `fmt.Scan`、`fmt.Scanf` 和 `fmt.Scanln` 这三个函数，用于从标准输入获取用户的输入。这些函数允许你与用户交互，从标准输入流中读取不同类型的数据并将其存储在相应的变量中。

### 4.1 fmt.Scan 函数

`Scan` 函数用于从标准输入中获取用户的输入，并将输入的数据存储在变量中。它根据空格分隔输入，适合获取多个输入值。

函数定义如下：

```go
func Scan(a ...interface{}) (n int, err error)
```

- Scan从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符。
- 本函数返回成功扫描的数据个数和遇到的任何错误。如果读取的数据个数比提供的参数少，会返回一个错误报告原因。

具体代码示例如下：

```go
package main

import "fmt"

func main() {
    var name string
    var age int

    fmt.Print("Enter your name: ")
    fmt.Scan(&name)
    fmt.Print("Enter your age: ")
    fmt.Scan(&age)

    fmt.Printf("Name: %s, Age: %d\n", name, age)
}
```

在这个示例中，`fmt.Scanf` 使用格式字符串 `%s %d` 来解析输入的姓名和年龄。

### 4.2 fmt.Scanln 函数

`Scanln` 函数用于从标准输入中获取用户的输入，并将输入的数据存储在变量中，每行一个变量。它通常用于获取多个输入值，每个值在单独的行中输入。

函数定义如下：

```go
func Scanln(a ...interface{}) (n int, err error)
```

- Scanln类似Scan，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。
- 本函数返回成功扫描的数据个数和遇到的任何错误。

代码示例：

```go
package main

import "fmt"

func main() {
    var name string
    var age int

    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)
    fmt.Print("Enter your age: ")
    fmt.Scanln(&age)

    fmt.Printf("Name: %s, Age: %d\n", name, age)
}
```

在上面的示例中，`fmt.Scanln` 用于获取用户输入的姓名和年龄，并将它们存储在相应的变量中。输入的每一行都对应一个变量。

### 4.3 fmt.Scanf 函数

`Scanf` 函数用于根据格式规范解析输入，并将数据存储在变量中。它允许你指定输入的格式，并可以处理不同类型的数据。

函数签名如下：

```go
func Scanf(format string, a ...interface{}) (n int, err error)
```

- Scanf从标准输入扫描文本，根据format参数指定的格式去读取由空白符分隔的值保存到传递给本函数的参数中。
- 本函数返回成功扫描的数据个数和遇到的任何错误。

代码示例如下：

```go
package main

import "fmt"

func main() {
    var name string
    var age int

    fmt.Print("Enter your name and age: ")
    fmt.Scanf("%s %d", &name, &age)

    fmt.Printf("Name: %s, Age: %d\n", name, age)
}
```

在这个示例中，`fmt.Scanf` 使用格式字符串 `%s %d` 来解析输入的姓名和年龄。

### 4.4 使用 `bufio` 包获取输入

`bufio` 包提供了一种更灵活的方式来处理输入，特别是在需要完整读取一行或多行输入的情况下。你可以使用 `bufio.NewReader` 创建一个输入缓冲区，然后使用 `ReadString` 函数来读取输入，直到指定的分隔符（例如换行符 `\n`）。这允许你获取包含空格在内的完整输入内容。

```go
func bufioDemo() {
    reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
    fmt.Print("请输入内容：")
    text, _ := reader.ReadString('\n') // 读取直到换行符
    text = strings.TrimSpace(text)
    fmt.Printf("%#v\n", text)
}

```

### 4.5 使用 `Fscan` 系列函数

`Fscan` 系列函数允许你从 `io.Reader` 接口中读取数据，而不仅仅是标准输入。这些函数与 `fmt.Scan`、`fmt.Scanf` 和 `fmt.Scanln` 类似，但允许你从任何实现 `io.Reader` 接口的地方读取数据。

- `Fscan`：从 `io.Reader` 中读取数据。
- `Fscanln`：从 `io.Reader` 中读取一行数据。
- `Fscanf`：根据指定的格式从 `io.Reader` 中读取数据。

```go
func Fscan(r io.Reader, a ...interface{}) (n int, err error)
func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)

```

代码示例：

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    input := "42 John"
    reader := strings.NewReader(input) // 从字符串生成读对象

    var age int
    var name string

    n, err := fmt.Fscanf(reader, "%d %s", &age, &name)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Read %d values: Age: %d, Name: %s\n", n, age, name)
}
```

### 4.6 使用 `Sscan` 系列函数

`Sscan` 系列函数允许你从字符串中读取数据，而不仅仅是从标准输入。这些函数与 `fmt.Scan`、`fmt.Scanf` 和 `fmt.Scanln` 类似，但允许你从字符串中读取数据。

- `Sscan`：从字符串中读取数据。
- `Sscanln`：从字符串中读取一行数据。
- `Sscanf`：根据指定的格式从字符串中读取数据。

```go
func Sscan(str string, a ...interface{}) (n int, err error)
func Sscanln(str string, a ...interface{}) (n int, err error)
func Sscanf(str string, format string, a ...interface{}) (n int, err error)

```

代码示例：

```go
package main

import (
    "fmt"
)

func main() {
    input := "Alice 30"
    var name string
    var age int

    n, err := fmt.Sscanf(input, "%s %d", &name, &age)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Read %d values: Name: %s, Age: %d\n", n, name, age)
}
```

