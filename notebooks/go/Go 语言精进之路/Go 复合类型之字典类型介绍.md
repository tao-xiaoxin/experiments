#  Go 复合类型之字典类型介绍

[TOC]



## 一、map类型介绍

### 1.1 什么是 map 类型？

**map 是 Go 语言提供的一种抽象数据类型，它表示一组无序的键值对。**用 key 和 value 分别代表 map 的键和值。而且，map 集合中每个 key 都是唯一的：

![WechatIMG200](https://billy.taoxiaoxin.club/md/2023/10/652404b280b0d66193b891db.jpg)

和切片类似，作为复合类型的 `map`，它在 `Go` 中的类型表示也是由 `key` 类型与 value 类型组成的，就像下面代码：

```go
map[key_type]value_type
```

key 与 value 的类型可以相同，也可以不同：

```go
map[string]string // key与value元素的类型相同
map[int]string    // key与value元素的类型不同
```

如果两个 map 类型的 key 元素类型相同，value 元素类型也相同，那么我们可以说它们是同一个 map 类型，否则就是不同的 map 类型。

这里，我们要注意，map 类型对 value 的类型没有限制，但是对 **key 的类型却有严格要求，因为 map 类型要保证 key 的唯一性**。**因此在这里，你一定要注意：函数类型、map 类型自身，以及切片类型是不能作为 map 的 key 类型的。**比如下面这段代码：

```go
// 函数类型不能作为key，因为函数类型是不可比较的
func keyFunc() {}

m := make(map[string]int)
m[keyFunc] = 1 // 编译错误

// map类型不能作为key
m1 := make(map[string]int)
m[m1] = 1 // 编译错误

// 切片类型不能作为key，因为切片是可变长度的，它们的内容可能会在运行时更改
s1 := []int{1,2,3}  
m[s1] = 1 // 编译错误
```

上面代码中,试图使用函数类型、map类型和切片类型作为key都会导致编译错误。

这是因为Go语言在实现map时,需要比较key是否相等,因此key需要支持==比较。但函数、map和切片类型的相等性比较涉及内存地址,无法简单判断,所以不能作为key。**所以，key 的类型必须支持“==”和“!=”两种比较操作符**。

还需要注意的是，**在 Go 语言中，函数类型、map 类型自身，以及切片只支持与 nil 的比较，而不支持同类型两个变量的比较**。如果像下面代码这样，进行这些类型的比较，Go 编译器将会报错：

```go
s1 := make([]int, 1)
s2 := make([]int, 2)
f1 := func() {}
f2 := func() {}
m1 := make(map[int]string)
m2 := make(map[int]string)
println(s1 == s2) // 错误：invalid operation: s1 == s2 (slice can only be compared to nil)
println(f1 == f2) // 错误：invalid operation: f1 == f2 (func can only be compared to nil)
println(m1 == m2) // 错误：invalid operation: m1 == m2 (map can only be compared to nil)
```

### 1.2 map 类型特性

在Go中，`map`具有以下特性：

- **无序性**: `map`中的键值对没有固定的顺序，遍历时可能不按照添加的顺序返回键值对。
- **动态增长**: `map`是动态的，它会根据需要自动增长以容纳更多的键值对，不需要预先指定大小。
- **零值**: 如果未初始化一个`map`，它将是`nil`，并且不能存储键值对。需要使用`make`函数来初始化一个`map`。
- **键的唯一性**: 在同一个`map`中，每个键只能出现一次。如果尝试使用相同的键插入多次，新值将覆盖旧值。
- **查询效率高**: `map`的查询操作通常非常快，因为它使用哈希表来存储数据，这使得通过键查找值的时间复杂度接近常数。
- **引用类型**: `map`是一种引用类型，多个变量可以引用并共享同一个`map`实例。

## 二.map 变量的声明和初始化

和切片一样，为 **map 类型变量显式赋值有两种方式：一种是使用复合字面值；另外一种是使用 make 这个预声明的内置函数。**

### 2.1 方法一：使用 `make` 函数声明和初始化（推荐）

这是最常见和推荐的方式，特别是在需要在`map`中添加键值对之前初始化`map`的情况下。使用`make`函数可以为`map`分配内存并进行初始化。

```go
// 使用 make 函数声明和初始化 map
myMap := make(map[keyType]valueType,capacity)
```

其中：

+ `keyType` 是键的类型。

+ `valueType` 是值的类型。

+ capacity表示`map`的初始容量,它是可选的，可以省略不写。

例如：和切片通过 `make` 进行初始化一样，通过 `make` 的初始化方式，我们可以为 `map` 类型变量指定键值对的初始容量，但无法进行具体的键值对赋值，就像下面代码这样：

```go
	// 创建一个存储整数到字符串的映射
	m1 := make(map[int]string) // 未指定初始容量
	m1[1] = "key"
	fmt.Println(m1)
```

**map 类型的容量不会受限于它的初始容量值，当其中的键值对数量超过初始容量后，Go 运行时会自动增加 `map` 类型的容量**，保证后续键值对的正常插入，比如下面这段代码：

```go
	m2 := make(map[int]string, 2) // 指定初始容量为2
	m2[1] = "One"
	m2[2] = "Two"
	m2[3] = "Three"
	fmt.Println(m2) // 输出：map[1:One 2:Two 3:Three] ，并不会报错
	fmt.Println(len(m2)) // 此时，map容量已经变为3
```

**总结：使用`make`函数初始化的`map`是空的，需要在后续代码中添加键值对。**

```go
	mm := make(map[int]string)
	fmt.Println(mm) // 输出 map[]
```

### 2.2 方法二：使用复合字面值声明初始化 map 类型变量

和切片类型变量一样，如果我们没有显式地赋予 map 变量初值，**map 类型变量的默认值为 `nil`**，比如，我们来看下面这段代码：

```go
var m map[string]int

if m == nil {
    fmt.Println("Map is nil")
} else {
    fmt.Println("Map is not nil")
}
```

不过切片变量和 map 变量在这里也有些不同。初值为零值 nil 的切片类型变量，可以借助内置的 append 的函数进行操作，这种在 Go 语言中被称为“**零值可用**”。定义“零值可用”的类型，可以提升我们开发者的使用体验，我们不用再担心变量的初始状态是否有效。比如，创建一个存储字符串到整数的映射，**但 map 类型，因为它内部实现的复杂性，无法“零值可用”**。所以，如果我们对处于零值状态的 map 变量直接进行操作，就会导致运行时异常（panic），从而导致程序进程异常退出：

```go
var m map[string]int // m = nil
m["key"] = 1         // 发生运行时异常：panic: assignment to entry in nil map
```

所以，我们必须对 map 类型变量进行显式初始化后才能使用。我们先来看这句代码：

```go
m := map[int]string{}
```

这里，我们显式初始化了 map 类型变量 m。不过，你要注意，虽然此时 map 类型变量 m 中没有任何键值对，但变量 m 也不等同于初值为 nil 的 map 变量。这个时候，我们对 m 进行键值对的插入操作，不会引发运行时异常。

这里我们再看看怎么通过稍微复杂一些的复合字面值，对 map 类型变量进行初始化：

```go
m1 := map[int][]string{
    1: []string{"val1_1", "val1_2"},
    3: []string{"val3_1", "val3_2", "val3_3"},
    7: []string{"val7_1"},
}

type Position struct { 
    x float64 
    y float64
}

m2 := map[Position]string{
    Position{29.935523, 52.568915}: "school",
    Position{25.352594, 113.304361}: "shopping-mall",
    Position{73.224455, 111.804306}: "hospital",
}
```

我们看到，上面代码虽然完成了对两个 map 类型变量 m1 和 m2 的显式初始化，但不知道你有没有发现一个问题，作为初值的字面值似乎有些“臃肿”。你看，作为初值的字面值采用了复合类型的元素类型，而且在编写字面值时还带上了各自的元素类型，比如作为 `map[int] []string` 值类型的`[]string`，以及作为 `map[Position]string` 的 key 类型的 Position。

别急！针对这种情况，Go 提供了“语法糖”。这种情况下，**Go 允许省略字面值中的元素类型。**因为 map 类型表示中包含了 key 和 value 的元素类型，Go 编译器已经有足够的信息，来推导出字面值中各个值的类型了。我们以 m2 为例，这里的显式初始化代码和上面变量 m2 的初始化代码是等价的：

```go
m2 := map[Position]string{
    {29.935523, 52.568915}: "school",
    {25.352594, 113.304361}: "shopping-mall",
    {73.224455, 111.804306}: "hospital",
}
```

综上，这种方式通常用于创建具有初始值的`map`。在这种情况下，不需要使用`make`函数。`map`的声明方式如下：

```go
// 使用字面量声明和初始化 map
myMap := map[keyType]valueType{
    key1: value1,
    key2: value2,
    // ...
}

```

其中：

+ `keyType` 是键的类型
+ `valueType` 是值的类型
+ 然后使用大括号 `{}` 包围键值对

## 三.map 变量的传递开销（map是引用传递）

和切片类型一样，map 也是**引用类型**。这就意味着 map 类型变量作为参数被传递给函数或方法的时候，实质上传递的只是一个“**描述符**”，而不是整个 map 的数据拷贝，所以这个传递的开销是固定的，而且也很小。

并且，当 map 变量被传递到函数或方法内部后，我们在函数内部对 map 类型参数的修改在函数外部也是可见的。比如你从这个示例中就可以看到，函数 foo 中对 map 类型变量 m 进行了修改，而这些修改在 foo 函数外也可见。

```go
package main
  
import "fmt"

func foo(m map[string]int) {
    m["key1"] = 11
    m["key2"] = 12
}

func main() {
    m := map[string]int{
        "key1": 1,
        "key2": 2,
    }

    fmt.Println(m) // map[key1:1 key2:2]  
    foo(m)
    fmt.Println(m) // map[key1:11 key2:12] 
}
```

所以，**map 引用类型**。**当 map 被赋值为一个新变量的时候，它们指向同一个内部数据结构**。因此，**当改变其中一个变量，就会影响到另一变量。**

## 四.map 的内部实现

### 4.1 map 类型在 Go 运行时层实现的示意图

和切片相比，map 类型的内部实现要更加复杂。Go 运行时使用一张哈希表来实现抽象的 map 类型。运行时实现了 map 类型操作的所有功能，包括查找、插入、删除等。在编译阶段，Go 编译器会将 Go 语法层面的 map 操作，重写成运行时对应的函数调用。大致的对应关系是这样的：

```go
// 创建map类型变量实例
m := make(map[keyType]valType, capacityhint) → m := runtime.makemap(maptype, capacityhint, m)

// 插入新键值对或给键重新赋值
m["key"] = "value" → v := runtime.mapassign(maptype, m, "key") v是用于后续存储value的空间的地址

// 获取某键的值 
v := m["key"]      → v := runtime.mapaccess1(maptype, m, "key")
v, ok := m["key"]  → v, ok := runtime.mapaccess2(maptype, m, "key")

// 删除某键
delete(m, "key")   → runtime.mapdelete(maptype, m, “key”)
```

这是 map 类型在 Go 运行时层实现的示意图：

![img](https://billy.taoxiaoxin.club/md/2023/10/65242b79193ee81a0ba77ede.jpg)

我们可以看到，和切片的运行时表示图相比，map 的实现示意图显然要复杂得多。接下来，我们结合这张图来简要描述一下 map 在运行时层的实现原理。接下来我们来看一下一个 map 变量在初始状态、进行键值对操作后，以及在并发场景下的 Go 运行时层的实现原理。

### 4.2 初始状态

从图中我们可以看到，与语法层面 map 类型变量（m）一一对应的是 `*runtime.hmap` 的实例，即 `runtime.hmap` 类型的指针，也就是我们前面在讲解 map 类型变量传递开销时提到的 **map 类型的描述符**。hmap 类型是 map 类型的头部结构（`header`），它存储了后续 `map` 类型操作所需的所有信息，包括：

![WechatIMG17](https://billy.taoxiaoxin.club/md/2023/10/6524ee3f708d0a718537c1b7.jpg)

真正用来存储键值对数据的是桶，也就是 bucket，每个 bucket 中存储的是 Hash 值低 bit 位数值相同的元素，默认的元素个数为 BUCKETSIZE（值为 8，Go 1.17 版本中在 `$GOROOT/src/cmd/compile/internal/reflectdata/reflect.go` 中定义，与 `runtime/map.go` 中常量 `bucketCnt` 保持一致）。

当某个 bucket（比如 buckets[0]) 的 8 个空槽 slot）都填满了，且 map 尚未达到扩容的条件的情况下，运行时会建立 overflow bucket，并将这个 overflow bucket 挂在上面 bucket（如 buckets[0]）末尾的 overflow 指针上，这样两个 buckets 形成了一个链表结构，直到下一次 map 扩容之前，这个结构都会一直存在。

从图中我们可以看到，每个 bucket 由三部分组成，从上到下分别是 tophash 区域、key 存储区域和 value 存储区域。

### 4.3 tophash 区域

当我们向 `map` 插入一条数据，或者是从 `map` 按 `key` 查询数据的时候，运行时都会使用哈希函数对 `key` 做哈希运算，并获得一个哈希值`（hashcode）`。这个 `hashcode` 非常关键，运行时会把 `hashcode`“一分为二”来看待，其中低位区的值用于选定 `bucket`，高位区的值用于在某个 `bucket` 中确定 `key` 的位置。我把这一过程整理成了下面这张示意图，你理解起来可以更直观：

![WechatIMG18](https://billy.taoxiaoxin.club/md/2023/10/6524ef01ad7080989961b003.jpg)

因此，每个 bucket 的 tophash 区域其实是用来快速定位 key 位置的，这样就避免了逐个 key 进行比较这种代价较大的操作。尤其是当 key 是 size 较大的字符串类型时，好处就更突出了。这是一种以空间换时间的思路。

### 4.4 key 存储区域

接着，我们看 tophash 区域下面是一块连续的内存区域，存储的是这个 bucket 承载的所有 key 数据。运行时在分配 bucket 的时候需要知道 key 的 Size。那么运行时是如何知道 key 的 size 的呢？

当我们声明一个 map 类型变量，比如 `var m map[string]int` 时，Go 运行时就会为这个变量对应的特定 map 类型，生成一个 `runtime.maptype` 实例。如果这个实例已经存在，就会直接复用。maptype 实例的结构是这样的：

```go
type maptype struct {
    typ        _type
    key        *_type
    elem       *_type
    bucket     *_type // internal type representing a hash bucket
    keysize    uint8  // size of key slot
    elemsize   uint8  // size of elem slot
    bucketsize uint16 // size of bucket
    flags      uint32
} 
```

我们可以看到，这个实例包含了我们需要的 map 类型中的所有"元信息"。我们前面提到过，编译器会把语法层面的 map 操作重写成运行时对应的函数调用，这些运行时函数都有一个共同的特点，那就是第一个参数都是 maptype 指针类型的参数。

**Go 运行时就是利用 maptype 参数中的信息确定 key 的类型和大小的**。`map` 所用的 hash 函数也存放在 `maptype.key.alg.hash(key, hmap.hash0)` 中。同时 maptype 的存在也让 Go 中所有 map 类型都共享一套运行时 map 操作函数，而不是像 `C++` 那样为每种 `map` 类型创建一套 `map` 操作函数，这样就节省了对最终二进制文件空间的占用。

### 4.5 value 存储区域

我们再接着看 key 存储区域下方的另外一块连续的内存区域，这个区域存储的是 key 对应的 `value`。和 `key` 一样，这个区域的创建也是得到了 `maptype` 中信息的帮助。Go 运行时采用了把 `key` 和 `value` 分开存储的方式，而不是采用一个 `kv` 接着一个 `kv` 的 `kv` 紧邻方式存储，这带来的其实是算法上的复杂性，但却减少了因内存对齐带来的内存浪费。

我们以 `map[int8]int64` 为例，看看下面的存储空间利用率对比图：

![img](https://billy.taoxiaoxin.club/md/2023/10/6524f039b0105468f91ec380.jpg)

你会看到，当前 Go 运行时使用的方案内存利用效率很高，而 kv 紧邻存储的方案在 `map[int8]int64` 这样的例子中内存浪费十分严重，它的内存利用率是 72/128=56.25%，有近一半的空间都浪费掉了。

另外，还有一点我要跟你强调一下，如果 key 或 value 的数据长度大于一定数值，那么运行时不会在 bucket 中直接存储数据，而是会存储 key 或 value 数据的指针。目前 Go 运行时定义的最大 key 和 value 的长度是这样的：

```go
// $GOROOT/src/runtime/map.go
const (
    maxKeySize  = 128
    maxElemSize = 128
)
```

## 五.map 扩容

我们前面提到过，map 会对底层使用的内存进行自动管理。因此，在使用过程中，当插入元素个数超出一定数值后，map 一定会存在自动扩容的问题，也就是怎么扩充 bucket 的数量，并重新在 bucket 间均衡分配数据的问题。

那么 map 在什么情况下会进行扩容呢？Go 运行时的 map 实现中引入了一个 `LoadFactor`（负载因子），当 count > **LoadFactor * 2^B** 或 `overflow bucket` 过多时，运行时会自动对 map 进行扩容。目前 Go 1.17 版本 `LoadFactor` 设置为 6.5`（loadFactorNum/loadFactorDen）`。这里是 Go 中与 map 扩容相关的部分源码：

```go
// $GOROOT/src/runtime/map.go
const (
  ... ...

  loadFactorNum = 13
  loadFactorDen = 2
  ... ...
)

func mapassign(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer {
  ... ...
  if !h.growing() && (overLoadFactor(h.count+1, h.B) || tooManyOverflowBuckets(h.noverflow, h.B)) {
    hashGrow(t, h)
    goto again // Growing the table invalidates everything, so try again
  }
  ... ...
}
```

这两方面原因导致的扩容，在运行时的操作其实是不一样的。如果是因为 overflow bucket 过多导致的“扩容”，实际上运行时会新建一个和现有规模一样的 bucket 数组，然后在 assign 和 delete 时做排空和迁移。

如果是因为当前数据数量超出 LoadFactor 指定水位而进行的扩容，那么运行时会建立一个**两倍于现有规模的 bucket 数组**，但真正的排空和迁移工作也是在 assign 和 delete 时逐步进行的。原 bucket 数组会挂在 hmap 的 oldbuckets 指针下面，直到原 buckets 数组中所有数据都迁移到新数组后，原 buckets 数组才会被释放。你可以结合下面的 map 扩容示意图来理解这个过程，这会让你理解得更深刻一些：

![WechatIMG20](https://billy.taoxiaoxin.club/md/2023/10/6524f16721fed608e6cd69d8.jpg)

## 六.map 与并发

接着我们来看一下 map 和并发。从上面的实现原理来看，充当 map 描述符角色的 hmap 实例自身是有状态的（hmap.flags），而且对状态的读写是没有并发保护的。所以说 **map 实例不是并发写安全的，也不支持并发读写**。如果我们对 map 实例进行并发读写，程序运行时就会抛出异常。你可以看看下面这个并发读写 map 的例子：

```go
package main

import (
    "fmt"
    "time"
)

func doIteration(m map[int]int) {
    for k, v := range m {
        _ = fmt.Sprintf("[%d, %d] ", k, v)
    }
}

func doWrite(m map[int]int) {
    for k, v := range m {
        m[k] = v + 1
    }
}

func main() {
    m := map[int]int{
        1: 11,
        2: 12,
        3: 13,
    }

    go func() {
        for i := 0; i < 1000; i++ {
            doIteration(m)
        }
    }()

    go func() {
        for i := 0; i < 1000; i++ {
            doWrite(m)
        }
    }()

    time.Sleep(5 * time.Second)
}
```

运行这个示例程序，我们会得到下面的执行错误结果：

```go
fatal error: concurrent map iteration and map write
```

不过，如果我们仅仅是进行并发读，map 是没有问题的。而且，Go 1.9 版本中引入了支持并发写安全的 sync.Map 类型，可以在并发读写的场景下替换掉 map。如果你有这方面的需求，可以查看一下[sync.Map 的手册](https://pkg.go.dev/sync#Map)。

另外，你要注意，考虑到 map 可以自动扩容，map 中数据元素的 value 位置可能在这一过程中发生变化，所以 **Go 不允许获取 map 中 value 的地址，这个约束是在编译期间就生效的**。下面这段代码就展示了 Go 编译器识别出获取 map 中 value 地址的语句后，给出的编译错误：

```go
p := &m[key]  // cannot take the address of m[key]
fmt.Println(p)
```

## 七、`map` 的基本操作

### 7.1 修改和更新键值对

首先 nil 的 map 类型变量，我们可以在其中插入符合 map 类型定义的任意新键值对。插入新键值对只需要把 value 赋值给 map 中对应的 key 就可以了：

```go
// 创建并初始化一个 map
myMap := make(map[string]int)
myMap["apple"] = 1
myMap["banana"] = 2
```

不需要自己判断数据有没有插入成功，因为 Go 会保证插入总是成功的。不过，如果我们插入新键值对的时候，某个 key 已经存在于 map 中了，那我们的插入操作就会用新值覆盖旧值：

```go
// 修改键 "apple" 对应的值
myMap["apple"] = 3

// 更新键 "cherry" 对应的值，如果键不存在则创建新键值对
myMap["cherry"] = 4

// 打印修改后的 map
fmt.Println(myMap) // 输出: map[apple:3 banana:2 cherry:4]
```

从这段代码中，您可以看到如何执行以下操作：

1. **修改键 "apple" 对应的值**：使用`myMap["apple"] = 3`这行代码，将键 "apple" 对应的值从原来的 1 修改为 3。
2. **更新键 "cherry" 对应的值**：使用`myMap["cherry"] = 4`这行代码，更新了键 "cherry" 对应的值为 4。如果键 "cherry" 不存在于`map`中，这行代码会创建一个新的键值对。
3. **打印修改后的 map**：最后使用`fmt.Println(myMap)`打印整个修改后的`map`，以显示更新后的键值对。

### 7.2 批量更新和修改(合并同类型map)

在Go中，可以使用循环遍历另一个`map`，然后使用遍历的键值对来批量更新或修改目标`map`的键值对。以下是一个实现类似于Python字典的`update()`方法的步骤：

1. 创建一个目标`map`，它将被更新或修改。
2. 创建一个源`map`，其中包含要合并到目标`map`的键值对。
3. 遍历源`map`的键值对。
4. 对于每个键值对，检查它是否存在于目标`map`中。
   - 如果存在，将目标`map`中的值更新为源`map`中的值。
   - 如果不存在，将源`map`中的键值对添加到目标`map`中。
5. 最终，目标`map`将包含源`map`中的所有键值对以及更新后的值。

以下是具体的Go代码示例：

```go
package main

import (
	"fmt"
)
func updateMap(target map[string]int, source map[string]int) {
	for key, value := range source {
		target[key] = value
	}
}
func main() {
	// 创建目标 map
	targetMap := map[string]int{
		"apple":  1,
		"banana": 2,
	}

	// 创建源 map，包含要更新或修改的键值对
	sourceMap := map[string]int{
		"apple":  3, // 更新 "apple" 的值为 3
		"cherry": 4, // 添加新的键值对 "cherry": 4
	}

	// 调用 updateMap 函数，将源 map 合并到目标 map 中
	updateMap(targetMap, sourceMap)

	// 打印更新后的目标 map
	fmt.Println(targetMap) // 输出：map[apple:3 banana:2 cherry:4]
}
```

### 7.3 获取键值对数量

要获取一个`map`中键值对的数量（也称为长度），可以使用Go语言的`len`函数。`len`函数返回`map`中键值对的数量。以下是获取`map`中键值对数量的示例：

```go
	// 创建并初始化一个 map
	myMap := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	// 使用 len 函数获取 map 的键值对数量
	count := len(myMap)

	// 打印键值对数量
	fmt.Println("键值对数量:", count)
```

不过，这里要注意的是**我们不能对 map 类型变量调用 cap，来获取当前容量**，这是 map 类型与切片类型的一个不同点。

### 7.4 查找和数据读取（判断某个键是否存在）

#### 7.4.1 查找和数据读取 map 语法格式

Go语言中有个判断map中键是否存在的特殊写法，格式如下:

```go
value, ok := map[key]
```

其中：

- `myMap` 是目标`map`，您希望在其中查找键。
- `key` 是您要查找的键。
- `value` 是一个变量，如果键存在，它将存储键对应的值，如果键不存在，则会获得值类型的零值。
- `ok` 是一个布尔值，用于指示键是否存在。如果键存在，`ok`为`true`；如果键不存在，`ok`为`false`。

map 类型更多用在查找和数据读取场合。所谓查找，就是判断某个 key 是否存在于某个 map 中。Go 语言的 map 类型支持通过用一种名为“**comma ok**”的惯用法，进行对某个 key 的查询。接下来我们就用“comma ok”惯用法改造一下上面的代码：

```go
m := make(map[string]int)
v, ok := m["key1"]
if !ok {
    // "key1"不在map中
}

// "key1"在map中，v将被赋予"key1"键对应的value
```

我们看到，这里我们通过了一个布尔类型变量 ok，来判断键“key1”是否存在于 map 中。如果存在，变量 v 就会被正确地赋值为键“key1”对应的 value。

不过，如果我们并不关心某个键对应的 value，而只关心某个键是否在于 map 中，我们可以使用空标识符替代变量 v，忽略可能返回的 value：

```go
m := make(map[string]int)
_, ok := m["key1"]
... ...
```

因此，你一定要记住：**在 Go 语言中，请使用“comma ok”惯用法对 map 进行键查找和键值读取操作。**

#### 7.4.2 实现get 方法查找map 对应的key

在Go中，要实现类似Python字典的`get()`方法，可以编写一个函数，该函数接受一个`map`、一个键以及一个默认值作为参数。函数将尝试从`map`中获取指定键的值，如果键不存在，则返回默认值。以下是实现类似`get()`方法的步骤：

1. 创建一个函数，命名为`get`，该函数接受三个参数：`map`、键和默认值。
2. 在函数中，使用键来尝试从`map`中获取对应的值。
3. 如果值存在，返回该值；如果不存在，则返回默认值空字符串。

```go
package main

import (
	"fmt"
)

// 实现类似 Python 字典的 get() 方法
func get(myMap map[string]string, key string) string {
	value, ok := myMap[key]
	if !ok {
		return ""
	}
	return value
}

func main() {
	// 创建并初始化一个 map
	myMap := map[string]string{
		"apple":  "red",
		"banana": "yellow",
		"cherry": "red",
	}

	// 使用 get() 方法获取键 "apple" 的值，如果不存在返回空字符串
	appleValue := get(myMap, "apple")
	fmt.Println("Color of 'apple':", appleValue)

	// 使用 get() 方法获取键 "tangerine" 的值，如果不存在返回空字符串
	grapeValue := get(myMap, "tangerine")
	if grapeValue == "" {
		fmt.Println("没有获取到tangerine的对应的值！")
	} else {
		fmt.Println("Color of 'tangerine':", grapeValue)
	}
}

```

运行此代码将输出：

```go
Color of 'apple': red
没有获取到tangerine的对应的值！
```

### 7.5 使用delete()函数删除键值对

使用`delete()`内建函数从map中删除一组键值对，`delete()`函数的格式如下：

```go
delete(map, key)
```

其中：

- map:表示要删除键值对的map
- key:表示要删除的键值对的键

使用 delete 函数的情况下，传入的第一个参数是我们的 map 类型变量，第二个参数就是我们想要删除的键。我们可以看看这个代码示例：

```go
m := map[string]int {
  "key1" : 1,
  "key2" : 2,
}

fmt.Println(m) // map[key1:1 key2:2]
delete(m, "key2") // 删除"key2"
fmt.Println(m) // map[key1:1]
```

### 7.6 遍历 map 中的键值数据

最后，我们来说一下如何遍历 map 中的键值数据。这一点虽然不像查询和读取操作那么常见，但日常开发中我们还是有这个需求的。在 Go 中，遍历 map 的键值对只有一种方法，那就是像**对待切片那样通过 for range 语句对 map 数据进行遍历**。我们看一个例子：

```go
package main
  
import "fmt"

func main() {
    m := map[int]int{
        1: 11,
        2: 12,
        3: 13,
    }

    fmt.Printf("{ ")
    for k, v := range m {
        fmt.Printf("[%d, %d] ", k, v)
    }
    fmt.Printf("}\n")
}
```

你看，通过 for range 遍历 map 变量 m，每次迭代都会返回一个键值对，其中键存在于变量 k 中，它对应的值存储在变量 v 中。我们可以运行一下这段代码，可以得到符合我们预期的结果：

```go
{ [1, 11] [2, 12] [3, 13] }
```

如果我们只关心每次迭代的键，我们可以使用下面的方式对 map 进行遍历：

```go
for k, _ := range m { 
  // 使用k
}
```

当然更地道的方式是这样的：

```go
for k := range m {
  // 使用k
}
```

如果我们只关心每次迭代返回的键所对应的 value，我们同样可以通过空标识符替代变量 k，就像下面这样：

```go
for _, v := range m {
  // 使用v
}
```

不过，前面 map 遍历的输出结果都非常理想，给我们的表象好像是迭代器按照 map 中元素的插入次序逐一遍历。那事实是不是这样呢？我们再来试试，多遍历几次这个 map 看看。

我们先来改造一下代码：

```go
package main
  
import "fmt"

func doIteration(m map[int]int) {
    fmt.Printf("{ ")
    for k, v := range m {
        fmt.Printf("[%d, %d] ", k, v)
    }
    fmt.Printf("}\n")
}

func main() {
    m := map[int]int{
        1: 11,
        2: 12,
        3: 13,
    }

    for i := 0; i < 3; i++ {
        doIteration(m)
    }
}
```

运行一下上述代码，我们可以得到这样结果：

```go
{ [1, 11] [2, 12] [3, 13] }
{ [2, 12] [3, 13] [1, 11] }
{ [1, 11] [2, 12] [3, 13] }
```

我们可以看到，**对同一 map 做多次遍历的时候，每次遍历元素的次序都不相同**。这是 Go 语言 map 类型的一个重要特点，也是很容易让 Go 初学者掉入坑中的一个地方。所以这里你一定要记住：**程序逻辑千万不要依赖遍历 map 所得到的的元素次序。**

## 八、Map的相等性

map 之间不能使用 `==` 操作符判断，`==` 只能用来检查 map 是否为 `nil`。

```go
func main() {
	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	map2 := map1
    if map1 ==nil{
    	fmt.Println("map1为空")
	}else {
		fmt.Println("map1不为空")
	}
	if map1 == map2 { // 直接报错，不能直接比较
	}
	
}
```

