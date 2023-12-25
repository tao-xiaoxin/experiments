# Go 泛型之类型参数

[TOC]



## 一、Go 的泛型与其他主流编程语言的泛型差异

Go泛型和其他支持泛型的主流编程语言之间的泛型设计与实现存在差异一样，Go 的泛型与其他主流编程语言的泛型也是不同的。我们先看一下 Go 泛型设计方案已经明确[不支持的若干特性](https://github.com/golang/proposal/blob/master/design/43651-type-parameters.md#omissions)，比如：

+ 不支持泛型特化（specialization），即不支持编写一个泛型函数针对某个具体类型的特殊版本；
+ 不支持元编程（metaprogramming），即不支持编写在编译时执行的代码来生成在运行时执行的代码；
+ 不支持操作符方法（operator method），即只能用普通的方法（method）操作类型实例（比如：getIndex(k)），而不能将操作符视为方法并自定义其实现，比如一个容器类型的下标访问 c[k]；
+ 不支持变长的类型参数（type parameters）；
+ ......

这些特性如今不支持，后续大概率也不会支持。在进入 Go 泛型语法学习之前，一定要先了解 Go 团队的这些设计决策。

## 二、返回切片中值最大的元素

我们先来看一个例子，实现一个函数，该函数接受一个切片作为输入参数，然后返回该切片中值最大的那个元素。题目并没有明确使用什么元素类型的切片，我们就先以最常见的整型切片为例，实现一个 `maxInt` 函数：

```go
// max_int.go
func maxInt(sl []int) int { 
    if len(sl) == 0 { 
        panic("slice is empty")
    } 

    max := sl[0]
    for _, v := range sl[1:] { 
        if v > max { 
            max = v 
        } 
    } 
    return max
}

func main() {
    fmt.Println(maxInt([]int{1, 2, -4, -6, 7, 0})) // 输出：7
}
```

`maxInt` 的逻辑十分简单。我们使用第一个元素值 (`max := sl[0]`) 作为 `max` 变量初值，然后与切片后面的元素 (`sl[1:]`) 进行逐一比较，如果后面的元素大于 `max`，则将其值赋给 `max`，这样到切片遍历结束，我们就得到了这个切片中值最大的那个元素（即变量 `max`）。

我们现在给它加一个新需求：能否针对元素为 `string` 类型的切片返回其最大（按字典序）的元素值呢？

答案肯定是能！我们来实现这个 `maxString` 函数：

```go
// max_string.go
func maxString(sl []string) string {
    if len(sl) == 0 {
        panic("slice is empty")
    }

    max := sl[0]
    for _, v := range sl[1:] {
        if v > max {
            max = v
        }
    }
    return max
}

func main() {
    fmt.Println(maxString([]string{"11", "22", "44", "66", "77", "10"})) // 输出：77
}
```

`maxString` 实现了返回 `string` 切片中值最大元素的需求。不过从实现上来看，`maxString` 与 `maxInt` 异曲同工，只是切片元素类型不同罢了。这时如果让你参考上述 `maxInt` 或 `maxString` 实现一个返回浮点类型切片中最大值的函数 `maxFloat`，你肯定“秒秒钟”就可以给出一个正确的实现：

```go
// max_float.go
func maxFloat(sl []float64) float64 {
    if len(sl) == 0 {
        panic("slice is empty")
    }

    max := sl[0]
    for _, v := range sl[1:] {
        if v > max {
            max = v
        }
    }
    return max
}

func main() {
    fmt.Println(maxFloat([]float64{1.01, 2.02, 3.03, 5.05, 7.07, 0.01})) // 输出：7.07
}
```

问题来了！你肯定在上面三个函数发现了的“糟糕味道”：**代码重复**。上面三个函数除了切片的元素类型不同，其他逻辑都一样。

那么能否实现一个“通用”的函数，可以处理上面三种元素类型的切片呢？提到“通用”，你一定想到了 Go 语言提供的 `any`（`interface{}`的别名），我们来试试：

```go
// max_any.go
func maxAny(sl []any) any {
    if len(sl) == 0 {
        panic("slice is empty")
    }

    max := sl[0]
    for _, v := range sl[1:] {
        switch v.(type) {
        case int:
            if v.(int) > max.(int) {
                max = v
            }
        case string:
            if v.(string) > max.(string) {
                max = v
            }
        case float64:
            if v.(float64) > max.(float64) {
                max = v
            }
        }
    }
    return max
}

func main() {
    i := maxAny([]any{1, 2, -4, -6, 7, 0})
    m := i.(int)
    fmt.Println(m) // 输出：7
    fmt.Println(maxAny([]any{"11", "22", "44", "66", "77", "10"})) // 输出：77
    fmt.Println(maxAny([]any{1.01, 2.02, 3.03, 5.05, 7.07, 0.01})) // 输出：7.07
}
```

我们看到，`maxAny` 利用 `any`、`type switch` 和类型断言（`type assertion`）实现了我们预期的目标。不过这个实现并不理想，它至少有如下几个问题：

1. 若要支持其他元素类型的切片，我们需对该函数进行修改；
2. `maxAny` 的返回值类型为 `any`（`interface{}`），要得到其实际类型的值还需要通过类型断言转换；
3. 使用 `any`（`interface{}`）作为输入参数的元素类型和返回值的类型，由于存在装箱和拆箱操作，其性能与 `maxInt` 等比起来要逊色不少，实测数据如下：

```go
// max_test.go
func BenchmarkMaxInt(b *testing.B) {
    sl := []int{1, 2, 3, 4, 7, 8, 9, 0}
    for i := 0; i < b.N; i++ {
        maxInt(sl)
    }
}

func BenchmarkMaxAny(b *testing.B) {
    sl := []any{1, 2, 3, 4, 7, 8, 9, 0}
    for i := 0; i < b.N; i++ {
        maxAny(sl)
    }
}
```

测试结果如下：

```go
$go test -v -bench . ./max_test.go max_any.go max_int.go
goos: darwin
goarch: amd64
... ...
BenchmarkMaxInt
BenchmarkMaxInt-8     398996863           2.982 ns/op
BenchmarkMaxAny
BenchmarkMaxAny-8     85883875          13.91 ns/op
PASS
ok    command-line-arguments  2.710s
```

我们看到，基于 `any`（`interface{}`） 实现的 `maxAny` 其执行性能要比像 `maxInt` 这样的函数慢上数倍。

在 Go 1.18 版本之前，Go 的确没有比较理想的解决类似上述“通用”问题的手段，直到 Go 1.18 版本泛型落地后，我们可以用泛型语法实现 `maxGenerics` 函数：

```go
// max_generics.go
type ordered interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
        ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
        ~float32 | ~float64 |
        ~string
}

func maxGenerics[T ordered](sl []T) T {
    if len(sl) == 0 {
        panic("slice is empty")
    }

    max := sl[0]
    for _, v := range sl[1:] {
        if v > max {
            max = v
        }
    }
    return max
}

type myString string

func main() {
    var m int = maxGenerics([]int{1, 2, -4, -6, 7, 0})
    fmt.Println(m) // 输出：7
    fmt.Println(maxGenerics([]string{"11", "22", "44", "66", "77", "10"})) // 输出：77
    fmt.Println(maxGenerics([]float64{1.01, 2.02, 3.03, 5.05, 7.07, 0.01})) // 输出：7.07
    fmt.Println(maxGenerics([]int8{1, 2, -4, -6, 7, 0})) // 输出：7
    fmt.Println(maxGenerics([]myString{"11", "22", "44", "66", "77", "10"})) // 输出：77
}
```

我们看到，从功能角度看，泛型版本的 `maxGenerics` 实现了预期的特性，对于 `ordered` 接口中声明的那些原生类型以及以这些原生类型为底层类型（underlying type）的类型（比如示例中的 `myString`），`maxGenerics` 都可以无缝支持。并且，`maxGenerics` 返回的类型与传入的切片的元素类型一致，调用者也无需通过类型断言做转换。

此外，通过下面的性能基准测试我们也可以看出，与 `maxAny` 相比，泛型版本的 `maxGenerics` 性能要好很多，但与原生版函数如 `maxInt` 等还有差距。性能测试如下：

```go
$go test -v -bench . ./max_test.go max_any.go max_int.go max_generics.go
goos: darwin
goarch: amd64
BenchmarkMaxInt
BenchmarkMaxInt-8          400910706           2.983 ns/op
BenchmarkMaxAny
BenchmarkMaxAny-8          85257433          14.04 ns/op
BenchmarkMaxGenerics
BenchmarkMaxGenerics-8     209468593           5.701 ns/op
PASS
ok    command-line-arguments  4.492s
```

通过这个例子，我们也可以看到 Go 泛型十分适合实现一些操作容器类型（比如切片、map 等）的算法，这也是 [Go 官方推荐的第一种泛型应用场景](https://go.dev/blog/when-generics)，此类容器算法的泛型实现使得容器算法与容器内元素类型彻底解耦！

## 三、类型参数（type parameters）

根据官方说法，由于“泛型”（`generic`）一词在 Go 社区中被广泛使用，所以官方也就接纳了这一说法。**但 Go 泛型方案的实质是对类型参数（`type parameter`）的支持**，包括：

+ 泛型函数（`generic function`）：带有类型参数的函数；
+ 泛型类型（`generic type`）：带有类型参数的自定义类型；
+ 泛型方法（`generic method`）：泛型类型的方法。

首先，以泛型函数为例来具体说明一下什么是类型参数。

## 四、泛型函数

### 3.1 泛型函数的结构

我们回顾一下上面的示例，`maxGenerics` 就是一个泛型函数，我们看一下 `maxGenerics` 的函数原型：

```go
func maxGenerics[T ordered](sl []T) T {
    // ... ...
}
```

我们看到，`maxGenerics` 这个函数与我们之前学过的普通 Go 函数（`ordinary function`）相比，至少有两点不同：

+ `maxGenerics` 函数在函数名称与函数参数列表之间多了一段由方括号括起的代码：`[T ordered]`；
+ `maxGenerics` 参数列表中的参数类型以及返回值列表中的返回值类型都是 `T`，而不是某个具体的类型。

`maxGenerics` 函数原型中多出的这段代码`[T ordered]`就是 **Go 泛型的类型参数列表（`type parameters list`）**，**示例中这个列表中仅有一个类型参数 `T`**，`ordered` 为类型参数的类型约束（`type constraint`）。类型约束之于类型参数，就好比常规参数列表中的类型之于常规参数。

Go 语言规范规定：**函数的类型参数列表位于函数名与函数参数列表之间，由方括号括起的固定个数的、由逗号分隔的类型参数声明组成，**其一般形式如下：

```go
func genericsFunc[T1 constraint1, T2, constraint2, ..., Tn constraintN](ordinary parameters list) (return values list)
```

函数一旦拥有类型参数，就可以用该参数作为常规参数列表和返回值列表中修饰参数和返回值的类型。我们继续 `maxGenerics` 泛型函数为例分析，它拥有一个类型参数 `T`，在常规参数列表中，`T` 被用作切片的元素类型；在返回值列表中，`T` 被用作返回值的类型。

按 Go 惯例，**类型参数名的首字母通常采用大写形式**，并且类型参数必须是具名的，即便你在后续的函数参数列表、返回值列表和函数体中没有使用该类型参数，也是这样。比如下面例子中的类型参数 `T`：

```go
func print[T any]() { // 正确
}     

func print[any]() {   // 编译错误：all type parameters must be named 
}
```

和常规参数列表中的参数名唯一一样，在同一个类型参数列表中，类型参数名字也要唯一，下面这样的代码将会导致 Go 编译器报错：

```go
func print[T1 any, T1 comparable](sl []T) { //  编译错误：T1 redeclared in this block
    //...
}
```

常规参数列表中的参数有其特定作用域，即从参数声明处开始到函数体结束。和常规参数类似，泛型函数中类型参数也有其作用域范围，这个范围从类型参数列表左侧的方括号`[`开始，一直持续到函数体结束，如下图所示：

![](https://billy.taoxiaoxin.club/md/2023/12/6583cb3b4d53a56841c06b88.png)

类型参数的作用域也决定了类型参数的声明顺序并不重要，也不会影响泛型函数的行为，于是下面的泛型函数声明与上图中的函数是等价的：

```go
func foo[M map[E]T, T any, E comparable](m M)(E, T) {
    //... ...
}
```

### 3.2 调用泛型函数

首先，我们对“类型参数”做一下细分。**和普通函数有形式参数与实际参数一样，类型参数也有类型形参（`type parameter`）和类型实参（`type argument`）之分。**其中类型形参就是泛型函数声明中的类型参数，以前面示例中的 `maxGenerics` 泛型函数为例，如下面代码，`maxGenerics` 的类型形参就是 `T`，而类型实参则是在调用 `maxGenerics` 时实际传递的类型 `int`：

```go
// 泛型函数声明：T为类型形参
func maxGenerics[T ordered](sl []T) T

// 调用泛型函数：int为类型实参
m := maxGenerics[int]([]int{1, 2, -4, -6, 7, 0})
```

从上面这段代码我们也可以看出调用泛型函数与调用普通函数的区别。**在调用泛型函数时，除了要传递普通参数列表对应的实参之外，还要显式传递类型实参，比如这里的 int。**并且，显式传递的类型实参要放在函数名和普通参数列表前的方括号中。

在反复揣摩上面代码和说明后，你可能会提出这样的一个问题：如果泛型函数的类型形参较多，那么逐一显式传入类型实参会让泛型函数的调用显得十分冗长，比如：

```go
foo[int, string, uint32, float64](1, "hello", 17, 3.14)
```

这样的写法对开发者而言显然谈不上十分友好。其实不光大家想到了这个问题，Go 团队的泛型实现者们也考虑了这个问题，并给出了解决方法：函数类型实参的自动推断（`function argument type inference`）。

顾名思义，这个机制就是通过判断传递的函数实参的类型来推断出类型实参的类型，从而允许开发者不必显式提供类型实参，下面是以 `maxGenerics` 函数为例的类型实参推断过程示意图：

![](https://billy.taoxiaoxin.club/md/2023/12/6583d36666957458f08f8c4e.png)

我们看到，当 `maxGenerics` 函数传入的实际参数为 `[]int{…}` 时，Go 编译器会将其类型 `[]int` 与泛型函数参数列表中对应参数的类型（`[]T`）作比较，并推断出 `T == int` 这一结果。当然这个例子的推断过程较为简单，那些有难度的，甚至无法肉眼可见的就交给 Go 编译器去处理吧，我们没有必要过于深入。

不过，这个类型实参自动推断有一个前提，你一定要记牢，**那就是它必须是函数的参数列表中使用了的类型形参**，否则就会像下面的示例中的代码，编译器将报无法推断类型实参的错误：

```go
func foo[T comparable, E any](a int, s E) {
}

foo(5, "hello") // 编译器错误：cannot infer T
```

在编译器无法推断出结果时，我们可以给予编译器“部分提示”，比如既然编译器无法推断出 `T` 的实参类型，那我们就显式告诉编译器 `T` 的实参类型，即在泛型函数调用时，在类型实参列表中显式传入 `T` 的实参类型，但 `E` 的实参类型依然由编译器自动推断，示例代码如下：

```go
var s = "hello"
foo[int](5, s)  //ok
foo[int,](5, s) //ok
```

那么，除了函数参数列表中的参数类型可以作为类型实参推断的依据外，函数返回值的类型是否也可以呢？我们看下面示例：

```go
func foo[T any](a int) T {
    var zero T
    return zero
}

var a int = foo(5) // 编译器错误：cannot infer T
println(a)
```

我们看到，这个函数仅在返回值中使用了类型参数，但编译器没能推断出 T 的类型，所以我们切记：**不能通过返回值类型来推断类型实参。**

有了函数类型实参推断后，在大多数情况下，我们调用泛型函数就无须显式传递类型实参了，开发者也因此获得了与普通函数调用几乎一致的体验。

其实泛型函数调用是一个不同于普通函数调用的过程，为了揭开其中的“奥秘”，接下来我们看看泛型函数调用过程究竟发生了什么。

### 3.3 泛型函数实例化（instantiation）

我们还以 `maxGenerics` 为例来演示一下这个过程：

```go
maxGenerics([]int{1, 2, -4, -6, 7, 0})
```

上面代码是对 `maxGenerics` 泛型函数的一次调用，Go 对这段泛型函数调用代码的处理分为两个阶段，如下图所示：

![](https://billy.taoxiaoxin.club/md/2023/12/6583d814da4bb83a9833d4a5.png)

我们看到，Go 首先会对泛型函数进行实例化（`instantiation`），即根据自动推断出的类型实参生成一个新函数（当然这一过程是在编译阶段完成的，不会对运行时性能产生影响），然后才会调用这个新函数对输入的函数参数进行处理。

我们也可以用一种更形象的方式来描述上述泛型函数的实例化过程。**实例化就好比一家生产“求最大值”机器的工厂**，它会根据要比较大小的对象的类型将这样的机器生产出来。以上面的例子来说，整个实例化过程如下：

+ **工厂接单**：调用 `maxGenerics([]int{…})`，工厂师傅发现要比较大小的对象类型为 `int`；
+ **模具检查与匹配**：检查 `int` 类型是否满足模具的约束要求，即 `int` 是否满足 `ordered` 约束，如满足，则将其作为类型实参替换 `maxGenerics` 函数中的类型形参 `T`，结果为 `maxGenerics[int]`；
+ **生产机器**：将泛型函数 `maxGenerics` 实例化为一个新函数，这里将其起名为 `maxGenericsInt`，其函数原型为 `func([]int) int`。本质上 `maxGenericsInt := maxGenerics[int]`。

我们实际的 Go 代码也可以真实得到这台新生产出的“机器”，如下面代码所示：

```go
maxGenericsInt := maxGenerics[int] // 实例化后得到的新“机器”：maxGenericsInt
fmt.Printf("%T\n", maxGenericsInt) // func([]int) int
```

一旦针对 `int` 对象的“求最大值”的机器被生产出来了，它就可以对目标对象进行处理了，这和普通的函数调用没有区别。这里就相当于调用如下代码：

```go
maxGenericsInt([]int{1, 2, -4, -6, 7, 0}) // 输出：7
```

整个过程只需检查传入的函数实参（`[]int{1, 2, …}`）的类型与 `maxGenericsInt` 函数原型中的形参类型（`[]int`）是否匹配即可。

另外要注意，当我们使用相同类型实参对泛型函数进行多次调用时，Go 仅会做一次实例化，并复用实例化后的函数，比如：

```go
maxGenerics([]int{1, 2, -4, -6, 7, 0})
maxGenerics([]int{11, 12, 14, -36,27, 0}) // 复用第一次调用后生成的原型为func([]int) int的函数
```

好了，接下来我们再来看 Go 对类型参数的另一类支持：带有类型参数的自定义类型，即泛型类型。

## 五、泛型类型

### 5.1 声明泛型类型

所谓泛型类型，就是在类型声明中带有类型参数的 Go 类型，比如下面代码中的 `maxableSlice`：

```go
// maxable_slice.go

type maxableSlice[T ordered] struct {
    elems []T
}
```

顾名思义，`maxableSlice` 是一个自定义切片类型，这个类型的特点是总可以获取其内部元素的最大值，其唯一的要求是其内部元素是可排序的，**它通过带有 `ordered` 约束的类型参数来明确这一要求。**像这样在定义中带有类型参数的类型就被称为泛型类型（`generic type`）。

从例子中的 `maxableSlice` 类型声明中我们可以看到，在泛型类型中，类型参数列表放在类型名字后面的方括号中。和泛型函数一样，泛型类型可以有多个类型参数，类型参数名通常是首字母大写的，这些类型参数也必须是具名的，且命名唯一。其一般形式如下：

```go
type TypeName[T1 constraint1, T2 constraint2, ..., Tn constraintN] TypeLiteral
```

和泛型函数中类型参数有其作用域一样，泛型类型中类型参数的作用域范围也是从类型参数列表左侧的方括号`[`开始，一直持续到类型定义结束的位置，如下图所示：

![](https://billy.taoxiaoxin.club/md/2023/12/6583dbaf18cae1bc0788197e.png)

这样的作用域将方便我们在各个字段中灵活使用类型参数，下面是一些自定义泛型类型的示例：

```go
type Set[T comparable] map[T]struct{}

type sliceFn[T any] struct {
  s   []T
  cmp func(T, T) bool
}

type Map[K, V any] struct {
  root    *node[K, V]
  compare func(K, K) int
}

type element[T any] struct {
  next *element[T]
  val  T
}

type Numeric interface {
  ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
    ~float32 | ~float64 |
    ~complex64 | ~complex128
}

type NumericAbs[T Numeric] interface {
  Abs() T
}
```

我们看到，泛型类型中的类型参数可以用来作为类型声明中字段的类型（比如上面的 `element` 类型）、复合类型的元素类型（比如上面的 `Set` 和 `Map` 类型）或方法的参数和返回值类型（如 `NumericAbs` 接口类型）等。

如果要在泛型类型声明的内部引用该类型名，必须要带上类型参数，如上面的 `element` 结构体中的 `next` 字段的类型：`*element[T]`。按照泛型设计方案，如果泛型类型有不止一个类型参数，那么在其声明内部引用该类型名时，不仅要带上所有类型参数，类型参数的顺序也要与声明中类型参数列表中的顺序一致，比如：

```go
type P[T1, T2 any] struct {
    F *P[T1, T2]  // ok
}
```

不过从实测结果来看，对于下面不符合技术方案的泛型类型声明也并未报错：

```go
type P[T1, T2 any] struct {
    F *P[T2, T1] // 不符合技术方案，但Go 编译器并未报错
}
```

### 5.2 使用泛型类型

和泛型函数一样，使用泛型类型时也会有一个实例化（`instantiation`）过程，比如：

```go
var sl = maxableSlice[int]{
    elems: []int{1, 2, -4, -6, 7, 0},
} 
```

Go 会根据传入的类型实参（`int`）生成一个新的类型并创建该类型的变量实例，`sl` 的类型等价于下面代码：

```go
type maxableIntSlice struct {
    elems []int
}
```

看到这里你可能会问：泛型类型是否可以像泛型函数那样实现类型实参的自动推断呢？很遗憾，目前的 Go 1.21.4 尚不支持，下面代码会遭到 Go 编译器的报错：

```go
var sl = maxableSlice {
    elems: []int{1, 2, -4, -6, 7, 0}, // 编译器错误：cannot use generic type maxableSlice[T ordered] without instantiation
} 
```

不过这一特性在 Go 的未来版本中可能会得到支持。

既然涉及到了类型，你肯定会想到诸如类型别名、类型嵌入等 Go 语言机制，那么这些语言机制对泛型类型的支持情况又是如何呢？我们逐一来看一下。

#### 5.2.1 泛型类型与类型别名

我们知道类型别名type alias）与其绑定的原类型是完全等价的，但这仅限于原类型是一个直接类型，即可直接用于声明变量的类型。那么将类型别名与泛型类型绑定是否可行呢？我们来看一个示例：

```go
type foo[T1 any, T2 comparable] struct {
    a T1
    b T2
}
  
type fooAlias = foo // 编译器错误：cannot use generic type foo[T1 any, T2 comparable] without instantiation
```

在上述代码中，我们为泛型类型 `foo` 建立了类型别名 `fooAlias`，但编译这段代码时，编译器还是报了错误！

这是因为，泛型类型只是一个生产真实类型的“工厂”，它自身在未实例化之前是不能直接用于声明变量的，因此不符合类型别名机制的要求。泛型类型只有实例化后才能得到一个真实类型，例如下面的代码就是合法的：

```go
type fooAlias = foo[int, string]
```

也就是说，我们只能为泛型类型实例化后的类型创建类型别名，实际上上述 `fooAlias` 等价于实例化后的类型 `fooInstantiation`：

```go
type fooInstantiation struct {
    a int   
    b string
}
```

#### 5.2.2 泛型类型与类型嵌入

类型嵌入是运用 Go 组合设计哲学的一个重要手段。引入泛型类型之后，我们依然可以在泛型类型定义中嵌入普通类型，比如下面示例中 `Lockable` 类型中嵌入的 `sync.Mutex`：

```go
type Lockable[T any] struct {
    t T
    sync.Mutex
}

func (l *Lockable[T]) Get() T {
    l.Lock()
    defer l.Unlock()
    return l.t
}

func (l *Lockable[T]) Set(v T) {
    l.Lock()
    defer l.Unlock()
    l.t = v
}
```

在泛型类型定义中，我们也可以将其他泛型类型实例化后的类型作为成员。现在我们改写一下上面的 `Lockable`，为其嵌入另外一个泛型类型实例化后的类型 `Slice[int]`：

```go
type Slice[T any] []T
  
func (s Slice[T]) String() string {
    if len(s) == 0 {
        return ""
    }
    var result = fmt.Sprintf("%v", s[0])
    for _, v := range s[1:] {
        result = fmt.Sprintf("%v, %v", result, v)
    }
    return result
}

type Lockable[T any] struct {
    t T
    Slice[int]
    sync.Mutex
}

func main() {
    n := Lockable[string]{
        t:     "hello",
        Slice: []int{1, 2, 3},
    }
    println(n.String()) // 输出：1, 2, 3
}
```

我们看到，代码使用泛型类型名（`Slice`）作为嵌入后的字段名，并且 `Slice[int]` 的方法 `String` 被提升为 `Lockable` 实例化后的类型的方法了。同理，在普通类型定义中，我们也可以使用实例化后的泛型类型作为成员，比如让上面的 `Slice[int]` 嵌入到一个普通类型 `Foo` 中，示例代码如下：

```go
type Foo struct {
    Slice[int]
}

func main() {
    f := Foo{
        Slice: []int{1, 2, 3},
    }
    println(f.String()) // 输出：1, 2, 3
}
```

此外，Go 泛型设计方案支持在泛型类型定义中嵌入类型参数作为成员，比如下面的泛型类型 `Lockable` 内嵌了一个类型 `T`，且 `T` 恰为其类型参数：

```go
type Lockable[T any] struct {
    T
    sync.Mutex
}
```

不过，Go 最新版`1.21.4` 编译上述代码时会针对嵌入 T 的那一行报如下错误：

```go
编译器报错：embedded field type cannot be a (pointer to a) type parameter
```

关于这个错误，[Go 官方在其 issue 中给出了临时的结论：暂不支持](https://github.com/golang/go/issues/49030)。

## 六、泛型方法

我们知道 Go 类型可以拥有自己的方法（`method`），泛型类型也不例外，为泛型类型定义的方法称为泛型方法（`generic method`），接下来我们就来看看如何定义和使用泛型方法。

我们用一个示例，给 `maxableSlice` 泛型类型定义 `max` 方法，看一下泛型方法的结构：

```go
func (sl *maxableSlice[T]) max() T {
    if len(sl.elems) == 0 {
        panic("slice is empty")
    }

    max := sl.elems[0]
    for _, v := range sl.elems[1:] {
        if v > max {
            max = v
        }
    }
    return max
}
```

我们看到，在定义泛型类型的方法时，方法的 `receiver` 部分不仅要带上类型名称，还需要带上完整的类型形参列表（如 `maxableSlice[T]`），这些类型形参后续可以用在方法的参数列表和返回值列表中。

不过在 Go 泛型目前的设计中，泛型方法自身不可以再支持类型参数了，不能像下面这样定义泛型方法：

```go
func (f *foo[T]) M1[E any](e E) T { // 编译器错误：syntax error: method must have no type parameters
    //... ...
}
```

关于泛型方法未来是否能支持类型参数，目前 Go 团队倾向于否，但最终结果 Go 团队还要根据 Go 社区在使用泛型过程中的反馈而定。

在泛型方法中，`receiver` 中某个类型参数如果没有在方法参数列表和返回值中使用，可以用“_”代替，但不能不写，比如：

```go
type foo[A comparable, B any] struct{}

func (foo[A, B]) M1() { // ok
}

或

func (foo[_, _]) M1() { // ok
}

或

func (foo[A, _]) M1() { // ok
}

但

func (foo[]) M1() { // 错误：receiver部分缺少类型参数

}
```

另外，泛型方法中的 `receiver` 中类型参数名字可以与泛型类型中的类型形参名字不同，位置和数量对上即可。我们还以上面的泛型类型 `foo` 为例，可以为它添加下面方法：

```go
type foo[A comparable, B any] struct{}

func (foo[First, Second]) M1(a First, b Second) { // First对应类型参数A，Second对应类型参数B

}
```

