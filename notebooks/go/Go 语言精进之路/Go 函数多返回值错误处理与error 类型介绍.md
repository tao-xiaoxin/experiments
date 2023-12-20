# Go 函数多返回值错误处理与error 类型介绍

[TOC]



## 一、error 类型与错误值构造

### 1.1 Error 接口介绍

在Go语言中，`error` 类型是一个接口类型，通常用于表示错误。它定义如下：

```go
type error interface {
    Error() string
}
```

`error` 接口只有一个方法，即 `Error()` 方法，该方法返回一个描述错误的字符串。这意味着任何实现了 `Error()` 方法的类型都可以被用作错误类型。通常，Go程序中的函数在遇到错误时会返回一个 `error` 类型的值，以便调用方可以处理或记录错误信息。

### 1.2 构造错误值的方法

### 1.2.1 使用errors包

Go 语言的设计者提供了两种方便 `Go` 开发者构造错误值的方法： `errors.New` 和 `fmt.Errorf` 。

+ `errors.New()` 函数是创建最简单的错误值的方法，它只包含一个错误消息字符串。这个方法适用于创建简单的错误值。
+ `fmt.Errorf()` 函数允许你构造一个格式化的错误消息，类似于 `fmt.Printf()` 函数。这对于需要构建更复杂的错误消息时非常有用。

使用这两种方法，我们可以轻松构造出一个满足 `error` 接口的错误值，就像下面代码这样：

```go
err := errors.New("your first demo error")
errWithCtx = fmt.Errorf("index %d is out of bounds", i)
```

这两种方法实际上返回的是同一个实现了 error 接口的类型的实例，这个未导出的类型就是 `errors.errorString`，它的定义是这样的：

```go
// $GOROOT/src/errors/errors.go

type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}
```

大多数情况下，使用这两种方法构建的错误值就可以满足我们的需求了。但我们也要看到，虽然这两种构建错误值的方法很方便，但**它们给错误处理者提供的错误上下文（Error Context）只限于以字符串形式呈现的信息，也就是 Error 方法返回的信息。**

### 1.2.2 自定义错误类型

在一些场景下，错误处理者需要从错误值中提取出更多信息，帮助他选择错误处理路径，显然这两种方法就不能满足了。这个时候，我们可以自定义错误类型来满足这一需求。以下是一个示例：

```go
package main

import "fmt"

// 自定义错误类型
type MyError struct {
	ErrorCode    int
	ErrorMessage string
}

// 实现 error 接口的 Error 方法
func (e MyError) Error() string {
	return fmt.Sprintf("错误 %d: %s", e.ErrorCode, e.ErrorMessage)
}

func someFunction() error {
	// 创建自定义错误值
	err := MyError{
		ErrorCode:    404,
		ErrorMessage: "未找到",
	}
	return err
}

func main() {
	// 调用 someFunction，返回自定义错误值
	err := someFunction()
	// 打印错误信息
	fmt.Println("错误:", err)
}
```

我们再来看一个例子，比如：标准库中的 `net` 包就定义了一种携带额外错误上下文的错误类型：

```go
// $GOROOT/src/net/net.go
type OpError struct {
    Op string
    Net string
    Source Addr
    Addr Addr
    Err error
}
```

这样，错误处理者就可以根据这个类型的错误值提供的额外上下文信息，比如 Op、Net、Source 等，做出错误处理路径的选择，比如下面标准库中的代码：

```go
// $GOROOT/src/net/http/server.go
func isCommonNetReadError(err error) bool {
    if err == io.EOF {
        return true
    }
    if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
        return true
    }
    if oe, ok := err.(*net.OpError); ok && oe.Op == "read" {
        return true
    }
    return false
}
```

我们看到，上面这段代码利用类型断言（`Type Assertion`），判断` error` 类型变量 err 的动态类型是否为 `*net.OpError` 或 net.Error。如果 `err` 的动态类型是 `*net.OpError`，那么类型断言就会返回这个动态类型的值（存储在 `oe` 中），代码就可以通过判断它的 `Op` 字段是否为"`read`"来判断它是否为 `CommonNetRead` 类型的错误。

## 二、error 类型的好处

### 2.1 第一点：统一了错误类型

如果不同开发者的代码、不同项目中的代码，甚至标准库中的代码，都统一以 `error` 接口变量的形式呈现错误类型，就能在提升代码可读性的同时，还更容易形成统一的错误处理策略。

### 2.2 第二点：错误是值

我们构造的错误都是值，也就是说，即便赋值给 error 这个接口类型变量，我们也可以像整型值那样对错误做“==”和“!=”的逻辑比较，函数调用者检视错误时的体验保持不变。

由于 error 是一个接口类型，默认零值为`nil`。所以我们通常将调用函数返回的错误与`nil`进行比较，以此来判断函数是否返回错误。如果返回的错误为 `nil`，则表示函数执行成功，否则表示出现了错误。这种约定使得错误处理变得一致和直观。例如你会经常看到类似下面的错误判断代码。

```go
func someFunction() error {
    // 模拟一个出错的情况
    return errors.New("这是一个错误")
}

func main() {
    err := someFunction()

    if err != nil {
        fmt.Println("函数执行失败，错误信息:", err)
    } else {
        fmt.Println("函数执行成功")
    }
}

```

### 2.3 第三点：易扩展，支持自定义错误上下文

虽然错误以 error 接口变量的形式统一呈现，但我们很容易通过自定义错误类型来扩展我们的错误上下文，就像前面的 Go 标准库的 `OpError` 类型那样。

error 接口是错误值的提供者与错误值的检视者之间的契约。error 接口的实现者负责提供错误上下文，供负责错误处理的代码使用。这种错误具体上下文与作为错误值类型的 error 接口类型的解耦，也体现了 Go 组合设计哲学中“正交”的理念。

## 三、Go 错误处理的惯用策略

### 3.1 策略一：透明错误处理策略

简单来说，Go 语言中的错误处理，就是根据函数 / 方法返回的 `error` 类型变量中携带的错误值信息做决策，并选择后续代码执行路径的过程。

这样，最简单的错误策略莫过于完全不关心返回错误值携带的具体上下文信息，只要发生错误就进入唯一的错误处理执行路径，比如下面这段代码：

```go
err := doSomething()
if err != nil {
    // 不关心err变量底层错误值所携带的具体上下文信息
    // 执行简单错误处理逻辑并返回
    ... ...
    return err
}
```

这是 **Go 语言中最常见的错误处理策略**，80% 以上的 Go 错误处理情形都可以归类到这种策略下。在这种策略下，由于错误处理方并不关心错误值的上下文，所以错误值的构造方（如上面的函数 doSomething）可以直接使用 Go 标准库提供的两个基本错误值构造方法 **errors.New** 和 **fmt.Errorf** 来构造错误值，就像下面这样：

```go
func doSomething(...) error {
    ... ...
    return errors.New("some error occurred")
}
```

这样构造出的错误值代表的上下文信息，**对错误处理方是透明的，因此这种策略称为“透明错误处理策略”。**在错误处理方不关心错误值上下文的前提下，透明错误处理策略能最大程度地减少错误处理方与错误值构造方之间的耦合关系。

### 3.2 策略二：“哨兵”错误处理策略

当错误处理方不能只根据“透明的错误值”就做出错误处理路径选取的情况下，错误处理方会尝试对返回的错误值进行检视，于是就有可能出现下面代码中的**反模式**：

```go
data, err := b.Peek(1)
if err != nil {
    switch err.Error() {
    case "bufio: negative count":
        // ... ...
        return
    case "bufio: buffer full":
        // ... ...
        return
    case "bufio: invalid use of UnreadByte":
        // ... ...
        return
    default:
        // ... ...
        return
    }
}
```

简单来说，反模式就是，错误处理方以透明错误值所能提供的唯一上下文信息（描述错误的字符串），作为错误处理路径选择的依据。但这种“反模式”会造成严重的**隐式耦合**。这也就意味着，错误值构造方不经意间的一次错误描述字符串的改动，都会造成错误处理方处理行为的变化，并且这种通过字符串比较的方式，对错误值进行检视的性能也很差。

那这有什么办法吗？Go 标准库采用了定义导出的（Exported）“哨兵”错误值的方式，来辅助错误处理方检视（inspect）错误值并做出错误处理分支的决策，比如下面的 bufio 包中定义的“哨兵错误”：

```go
// $GOROOT/src/bufio/bufio.go
var (
    ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
    ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
    ErrBufferFull        = errors.New("bufio: buffer full")
    ErrNegativeCount     = errors.New("bufio: negative count")
)
```

下面的代码片段利用了上面的哨兵错误，进行错误处理分支的决策：

```go
data, err := b.Peek(1)
if err != nil {
    switch err {
    case bufio.ErrNegativeCount:
        // ... ...
        return
    case bufio.ErrBufferFull:
        // ... ...
        return
    case bufio.ErrInvalidUnreadByte:
        // ... ...
        return
    default:
        // ... ...
        return
    }
}
```

你可以看到，**一般“哨兵”错误值变量以 ErrXXX 格式命名。和透明错误策略相比，“哨兵”策略让错误处理方在有检视错误值的需求时候，**可以“有的放矢”。

不过，对于 API 的开发者而言，暴露“哨兵”错误值也意味着这些错误值和包的公共函数 / 方法一起成为了 API 的一部分。一旦发布出去，开发者就要对它进行很好的维护。而“哨兵”错误值也让使用这些值的错误处理方对它产生了依赖。

从 **Go 1.13 版本开始**，**标准库 errors 包提供了 Is 函数用于错误处理方对错误值的检视。Is 函数类似于把一个 error 类型变量与“哨兵”错误值进行比较**，比如下面代码：

```go
// 类似 if err == ErrOutOfBounds{ … }
if errors.Is(err, ErrOutOfBounds) {
    // 越界的错误处理
}
```

不同的是，如果 error 类型变量的底层错误值是一个包装错误（Wrapped Error），errors.Is 方法会沿着该包装错误所在错误链（Error Chain)，与链上所有被包装的错误（Wrapped Error）进行比较，直至找到一个匹配的错误为止。下面是 Is 函数应用的一个例子：

```go
var ErrSentinel = errors.New("the underlying sentinel error")

func main() {
  err1 := fmt.Errorf("wrap sentinel: %w", ErrSentinel)
  err2 := fmt.Errorf("wrap err1: %w", err1)
    println(err2 == ErrSentinel) //false
  if errors.Is(err2, ErrSentinel) {
    println("err2 is ErrSentinel")
    return
  }

  println("err2 is not ErrSentinel")
}
```

在这个例子中，我们通过 `fmt.Errorf` 函数，并且使用` %w `创建包装错误变量 err1 和 err2，其中 err1 实现了对 ErrSentinel 这个“哨兵错误值”的包装，而 err2 又对 err1 进行了包装，这样就形成了一条错误链。位于错误链最上层的是 err2，位于最底层的是 ErrSentinel。之后，我们再分别通过值比较和 errors.Is 这两种方法，判断 err2 与 ErrSentinel 的关系。运行上述代码，我们会看到如下结果：

```go
false
err2 is ErrSentinel
```

我们看到，通过比较操作符对 err2 与 ErrSentinel 进行比较后，我们发现这二者并不相同。而 errors.Is 函数则会沿着 err2 所在错误链，向下找到被包装到最底层的“哨兵”错误值`ErrSentinel`。

如果你使用的是 **Go 1.13 及后续版本，建议你尽量使用`errors.Is`方法去检视某个错误值是否就是某个预期错误值，或者包装了某个特定的“哨兵”错误值。**

### 3.3 策略三：错误值类型检视策略

上面我们看到，基于 Go 标准库提供的错误值构造方法构造的“哨兵”错误值，除了让错误处理方可以“有的放矢”的进行值比较之外，并没有提供其他有效的错误上下文信息。那如果遇到错误处理方需要错误值提供更多的“错误上下文”的情况，上面这些错误处理策略和错误值构造方式都无法满足。

这种情况下，我们需要通过自定义错误类型的构造错误值的方式，来提供更多的“错误上下文”信息。并且，由于错误值都通过 error 接口变量统一呈现，要得到底层错误类型携带的错误上下文信息，错误处理方需要使用 Go 提供的**类型断言机制（Type Assertion）或类型选择机制（Type Switch）**，这种错误处理方式，我称之为**错误值类型检视策略**。

我们来看一个标准库中的例子加深下理解，这个 json 包中自定义了一个 `UnmarshalTypeError` 的错误类型：

```go
// $GOROOT/src/encoding/json/decode.go
type UnmarshalTypeError struct {
    Value  string       
    Type   reflect.Type 
    Offset int64        
    Struct string       
    Field  string      
}
```

错误处理方可以通过错误类型检视策略，获得更多错误值的错误上下文信息，下面就是利用这一策略的` json `包的一个方法的实现：

```go
// $GOROOT/src/encoding/json/decode.go
func (d *decodeState) addErrorContext(err error) error {
    if d.errorContext.Struct != nil || len(d.errorContext.FieldStack) > 0 {
        switch err := err.(type) {
        case *UnmarshalTypeError:
            err.Struct = d.errorContext.Struct.Name()
            err.Field = strings.Join(d.errorContext.FieldStack, ".")
            return err
        }
    }
    return err
}
```

我们看到，这段代码通过类型 switch 语句得到了 err 变量代表的动态类型和值，然后在匹配的 case 分支中利用错误上下文信息进行处理。

这里，一般自定义导出的错误类型以 `XXXError` 的形式命名。和“哨兵”错误处理策略一样，错误值类型检视策略，由于暴露了自定义的错误类型给错误处理方，因此这些错误类型也和包的公共函数 / 方法一起，成为了 API 的一部分。一旦发布出去，开发者就要对它们进行很好的维护。而它们也让使用这些类型进行检视的错误处理方对其产生了依赖。

从 **Go 1.13 版本开始，标准库 errors 包提供了As函数给错误处理方检视错误值。As函数类似于通过类型断言判断一个 error 类型变量是否为特定的自定义错误类型**，如下面代码所示：

```go
// 类似 if e, ok := err.(*MyError); ok { … }
var e *MyError
if errors.As(err, &e) {
    // 如果err类型为*MyError，变量e将被设置为对应的错误值
}
```

不同的是，如果 error 类型变量的动态错误值是一个包装错误，`errors.As`函数会沿着该包装错误所在错误链，与链上所有被包装的错误的类型进行比较，直至找到一个匹配的错误类型，就像 `errors.Is` 函数那样。下面是`As`函数应用的一个例子：

```go
type MyError struct {
    e string
}

func (e *MyError) Error() string {
    return e.e
}

func main() {
    var err = &MyError{"MyError error demo"}
    err1 := fmt.Errorf("wrap err: %w", err)
    err2 := fmt.Errorf("wrap err1: %w", err1)
    var e *MyError
    if errors.As(err2, &e) {
        println("MyError is on the chain of err2")
        println(e == err)                  
        return                             
    }                                      
    println("MyError is not on the chain of err2")
} 
```

运行上述代码会得到：

```go
MyError is on the chain of err2
true
```

我们看到，`errors.As` 函数沿着 `err2` 所在错误链向下找到了被包装到最深处的错误值，并将 `err2` 与其类型 `* MyError` 成功匹配。匹配成功后，errors.As 会将匹配到的错误值存储到 As 函数的第二个参数中，这也是为什么 `println(e == err`)输出 `true` 的原因。

如果你使用的是 **Go 1.13 及后续版本，请尽量使用 errors.As方法去检视某个错误值是否是某自定义错误类型的实例**。

### 3.4 策略四：错误行为特征检视策略

不知道你注意到没有，在前面我们已经讲过的三种策略中，其实只有第一种策略，也就是“透明错误处理策略”，有效降低了错误的构造方与错误处理方两者之间的耦合。虽然前面的策略二和策略三，都是我们实际编码中有效的错误处理策略，但其实使用这两种策略的代码，依然在错误的构造方与错误处理方两者之间建立了耦合。

那么除了“透明错误处理策略”外，我们是否还有手段可以降低错误处理方与错误值构造方的耦合呢？

在 Go 标准库中，我们发现了这样一种错误处理方式：**将某个包中的错误类型归类，统一提取出一些公共的错误行为特征，并将这些错误行为特征放入一个公开的接口类型中。这种方式也被叫做错误行为特征检视策略。**

以标准库中的`net`包为例，它将包内的所有错误类型的公共行为特征抽象并放入 `net.Error` 这个接口中，如下面代码：

```go
// $GOROOT/src/net/net.go
type Error interface {
    error
    Timeout() bool  
    Temporary() bool
}
```

我们看到，`net.Error` 接口包含两个用于判断错误行为特征的方法：`Timeout` 用来判断是否是超时（`Timeout`）错误，`Temporary` 用于判断是否是临时（`Temporary`）错误。

而错误处理方只需要依赖这个公共接口，就可以检视具体错误值的错误行为特征信息，并根据这些信息做出后续错误处理分支选择的决策。

这里，我们再看一个 http 包使用错误行为特征检视策略进行错误处理的例子，加深下理解：

```go
// $GOROOT/src/net/http/server.go
func (srv *Server) Serve(l net.Listener) error {
    ... ...
    for {
        rw, e := l.Accept()
        if e != nil {
            select {
            case <-srv.getDoneChan():
                return ErrServerClosed
            default:
            }
            if ne, ok := e.(net.Error); ok && ne.Temporary() {
                // 注：这里对临时性(temporary)错误进行处理
                ... ...
                time.Sleep(tempDelay)
                continue
            }
            return e
        }
        ...
    }
    ... ...
}
```

在上面代码中，`Accept` 方法实际上返回的错误类型为 *OpError，它是 `net` 包中的一个自定义错误类型，它实现了错误公共特征接口 `net.Error`，如下代码所示：

```go
// $GOROOT/src/net/net.go
type OpError struct {
    ... ...
    // Err is the error that occurred during the operation.
    Err error
}

type temporary interface {
    Temporary() bool
}

func (e *OpError) Temporary() bool {
  if ne, ok := e.Err.(*os.SyscallError); ok {
      t, ok := ne.Err.(temporary)
      return ok && t.Temporary()
  }
  t, ok := e.Err.(temporary)
  return ok && t.Temporary()
}
```

因此，OpError 实例可以被错误处理方通过 net.Error 接口的方法，判断它的行为是否满足 Temporary 或 Timeout 特征。

## 四、总结

Go 语言统一错误类型为 error 接口类型，并提供了多种快速构建可赋值给 error 类型的错误值的函数，包括 errors.New、fmt.Errorf 等，我们还讲解了使用统一 error 作为错误类型的优点，你要深刻理解这一点。

基于 Go 错误处理机制、统一的错误值类型以及错误值构造方法的基础上，Go 语言形成了多种错误处理的惯用策略，包括透明错误处理策略、“哨兵”错误处理策略、错误值类型检视策略以及错误行为特征检视策略等。这些策略都有适用的场合，但没有某种单一的错误处理策略可以适合所有项目或所有场合。

在错误处理策略选择上，你可以参考以下：

+ 请尽量使用“透明错误”处理策略，降低错误处理方与错误值构造方之间的耦合；
+ 如果可以从众多错误类型中提取公共的错误行为特征，那么请尽量使用“错误行为特征检视策略”;
+ 在上述两种策略无法实施的情况下，再使用“哨兵”策略和“错误值类型检视”策略；
+ Go 1.13 及后续版本中，尽量用 `errors.Is` 和 `errors.As `函数替换原先的错误检视比较语句。