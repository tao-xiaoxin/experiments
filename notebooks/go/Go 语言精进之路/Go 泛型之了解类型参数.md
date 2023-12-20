# Go 泛型之了解类型参数

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