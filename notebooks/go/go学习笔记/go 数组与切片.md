# Go 数组与切片

## 一.数组

### 1.1 数组定义与基本语法

**数组** 是一个由 **固定长度** 的特定类型元素组成的序列，一个数组可以由零个或多个元素组成。在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化。 基本语法：

```go
var arr [n]type
```

其中:

+ `arr` 为数组变量名称
+ `n`为数组的长度,一旦定义，长度不能变
+ `type` 为数组的类型

举个例子:

```go
// 定义一个长度为3元素类型为int的数组a
var a [3]int
```

在go 中,**如果两个数组类型的元素类型 type 与数组长度 n 都是一样的，那么这两个数组类型是等价的，如果有一个属性不同，它们就是两个不同的数组类型**,下面这个示例很好地诠释了这一点：

```go
package main

func foo(arr [5]int) {}
func main() {
	var arr1 [5]int
	var arr2 [6]int
	var arr3 [5]string

	foo(arr1) // ok
	foo(arr2) // 错误：[6]int与函数foo参数的类型[5]int不是同一数组类型
	foo(arr3) // 错误：[5]string与函数foo参数的类型[5]int不是同一数组类型
}
```

### 1.2 数组的初始化

方式一:

+ 初始化数组时可以使用初始化列表来设置数组元素的值。

```go
func main() {
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      //[0 0 0]
	fmt.Println(numArray)                       //[1 2 0]
	fmt.Println(cityArray)                      //[北京 上海 深圳]
}
```

方式二:

+ 一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度,**忽略掉初始化表达式中数组的长度，用“…”替代**,比如：

```go
func main() {
	var testArray [3]int
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray)                          //[0 0 0]
	fmt.Println(numArray)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray)   //type of numArray:[2]int
	fmt.Println(cityArray)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray) //type of cityArray:[3]string
}
```

方式三:

+ 使用指定索引值的方式来初始化数组

```go
func main() {
	var arr4 = [...]int{
		99: 39, // 将第100个元素(下标值为99)的值赋值为39，其余元素值均为0
	}
	fmt.Printf("%T\n", arr4) // [100]int
	fmt.Println(arr4[99])    //39
}
```

我们声明一个数组类型变量的同时，也可以显式地对它进行初始化。如果不进行显式初始化，那么数组中的元素值就是它类型的零值。比如下面的数组类型变量 arr1 的各个元素值都为 0：

```go
var arr1 [6]int // [0 0 0 0 0 0]
```



### 1.3 数组长度与大小

**Go 提供了预定义函数 len 可以用于获取一个数组类型变量的长度，`cap`获取数组的容量,通过`unsafe` 包提供的 `Sizeof` 函数，我们可以获得一个数组变量的总大小**,如下面代码：

```go
	var arr = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("数组长度：", len(arr))           // 6
	fmt.Println("数组大小：", unsafe.Sizeof(arr)) // 48
	fmt.Println("数组容量：", cap(arr))           // 6
```

**数组大小就是所有元素的大小之和**，这里数组元素的类型为 int。在 64 位平台上，int 类型的大小为 8，数组 arr 一共有 6 个元素，因此它的总大小为 6x8=48 个字节



### 1.4 数组的下标

**数组的下标值是从 0 开始的**。如果下标值超出数组长度范畴，或者是负数，那么 Go 编译器会给出错误提示，防止访问溢出:

```go
func main() {
	var arr = [6]int{11, 12, 13, 14, 15, 16}
	fmt.Println(arr[0], arr[5]) // 11 16
	fmt.Println(arr[-1])        // 错误：下标值不能为负数
	fmt.Println(arr[8])         // 错误：小标值超出了arr的长度范围
}
```

### 1.5 数组遍历

遍历数组a有以下两种方法：

```go
func main() {
	var a = [...]string{"甘肃", "上海", "北京"}
	// 方法1：for循环遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 方法2：for range遍历
	for index, value := range a {
		fmt.Println(index, value)
	}
}
```

### 1.6 多维数组

多维数组:数组类型自身也可以作为数组元素的类型，这样就会产生多维数组，比如下面的变量 mArr 的类型就是一个多维数组[2] [3][4]int:

```go
var mArr [2][3][4]int
```

举个栗子:

```go
func main() {
	// 定义多维数组
	arr := [3][2]string{
		{"1", "Go语言极简一本通"},
		{"2", "Go语言微服务架构核心22讲"},
		{"3", "从0到Go语言微服务架构师"}}
	fmt.Println(arr) // [[1 Go语言极简一本通] [2 Go语言微服务架构核心22讲] [3 从0到Go语言微服务架构师]]
}
```

### 1.7 多维数组循环

```go
func main() {
	// 定义多维数组
	arr := [3][2]string{
		{"1", "Go语言极简一本通"},
		{"2", "Go语言微服务架构核心22讲"},
		{"3", "从0到Go语言微服务架构师"}}
	for _, v1 := range arr {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
}
```

### 1.8 数组是值类型

Go 中的数组是值类型而不是引用类型。当数组赋值给一个新的变量时，该变量会得到一个原始数组的一个副本。如果对新变量进行更改，不会影响原始数组。

```go
func main() {
	arr := [...]string{"Go语言极简一本通", "Go语言微服务架构核心22讲", "从0到Go语言微服务架构师"}
	copy := arr
	copy[0] = "Golang"
	fmt.Println(arr) //[Go语言极简一本通 Go语言微服务架构核心22讲 从0到Go语言微服务架构师]
	fmt.Println(copy) //[Golang Go语言微服务架构核心22讲 从0到Go语言微服务架构师]
}
```

## 2.切片

### 2.1 切片定义与基本语法

**定义**:切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。

切片是一个引用类型，它的内部结构包含`地址`、`长度`和`容量`。切片一般用于快速地操作一块数据集合。

声明切片类型的基本语法如下:

```go
var name []type
```

其中，

- name:表示变量名
- type:表示切片中的元素类型

举个列子:

```go
// 声明整型切片
var numList []int

// 声明一个空切片
var numListEmpty = []int{}
```

### 2.2 切片的初始化

方式一:

直接声明:

```go
var name []type
```

方式二:

**通过 make 函数来创建切片，并指定底层数组的长度**。我们直接看下面这行代码：

```go
var a:= make([]type, length, capacity)
```

其中:

+ type为数据类型
+ length 为长度
+ capacity 为容量

举个例子:

```go
sl := make([]byte, 6, 10) // 其中10为cap值，即底层数组长度，6为切片的初始长度
```

如果没有在 make 中指定 cap 参数，那么底层数组长度 cap 就等于 len，比如：

```go
sl := make([]byte, 6) // cap = len = 6
```

方式三:

**采用` array[low : high : max]`语法基于一个已存在的数组创建切片。这种方式被称为数组的切片化，比如下面代码：**

```go
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sl := arr[3:7:9]
	fmt.Println("arr:", arr) //arr: [1 2 3 4 5 6 7 8 9 10]
	fmt.Println("s1:", sl)   //s1: [4 5 6 7]
```

### 2.3 .实现切片的实现原理与长度,容量

Go 切片在运行时其实是一个三元组结构，它在 Go 运行时中的表示如下：

```go
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}
```

我们可以看到，每个切片包含三个字段：

+ array: 是指向底层数组的指针；

+ len: 是切片的长度，即切片中当前元素的个数；
+ cap: 容量就是从创建切片索引开始的底层数组中的元素个数,也是切片的最大容量，**cap 值永远大于等于 len 值。**

内置的 `len` 和 `cap` 函数分别返回 slice 的长度和容量。

```go
s := make([]string, 3, 5)
fmt.Println(len(s)) // 3
fmt.Println(cap(s)) // 5
```

如果切片操作超出上限将导致一个 `panic` 异常。

```go
s := make([]int, 3, 5)
fmt.Println(s[10]) //panic: runtime error: index out of range [10] with length 3
```

### 2.3 判断切片是否为空

- 由于 slice 是引用类型，所以你不对它进行赋值的话，它的默认值是 `nil`

  ```go
  var numList []int
  fmt.Println(numList == nil) // true
  ```

- 切片之间不能比较，因此我们不能使用 `==` 操作符来判断两个 slice 是否含有全部相等元素。特别注意，如果你需要测试一个 slice 是否是空的，使用 `len(s) == 0` 来判断，而不应该用 `s == nil` 来判断。

### 2.4 切片的元素的修改

**切片自己不拥有任何数据。它只是底层数组的一种表示**。对切片所做的任何修改都会反映在底层数组中。

下面的代码中演示了拷贝前后两个变量共享底层数组，**对一个切片的修改会影响另一个切片的内容，这点需要特别注意**。

```go
func main() {
	s1 := make([]int, 3) //[0 0 0]
	s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	s2[0] = 100
	fmt.Println("s1:", s1) //s1: [100 0 0]
	fmt.Println("s2:", s2) //s2: [100 0 0]
}
```

### 2.5 动态追加元素

使用 `append` 可以将新元素追加到切片上。

`append` 函数的定义是 `func append(slice []Type, elems ...Type) []Type` 。其中 `elems ...Type` 在函数定义中表示该函数接受参数 `elems` 的个数是可变的。这些类型的函数被称为可变函数。

```go
  func appendSliceData() {
	s := []string{"Go语言极简一本通"}
	fmt.Println(s)
	fmt.Println(cap(s))

	s = append(s, "Go语言微服务架构核心22讲")
	fmt.Println(s)
	fmt.Println(cap(s))

	s = append(s, "从0到Go语言微服务架构师", "分布式")
	fmt.Println(s)
	fmt.Println(cap(s))

	s = append(s, []string{"微服务", "分布式锁"}...)
	fmt.Println(s)
	fmt.Println(cap(s))
}
```

**当新的元素被添加到切片时，如果容量不足，会创建一个新的数组。现有数组的元素被复制到这个新数组中，并返回新的引用。现在新切片的容量是旧切片的两倍。**

### 2.6 多维切片

类似于数组，切片也可以有多个维度。

```go
func mSlice() {
	numList := [][]string{
		{"1", "Go语言极简一本通"},
		{"2", "Go语言微服务架构核心22讲"},
		{"3", "从0到Go语言微服务架构师"},
	}
	fmt.Println(numList)
}
```

### 2.7 切片的遍历

切片的遍历方式和数组是一致的，支持索引遍历和`for range`遍历。

```go
func main() {
	s := []int{1, 3, 5}

	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	for index, value := range s {
		fmt.Println(index, value)
	}
}
```

### 2.8 使用copy()函数复制切片

首先我们来看一个问题：

```go
func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(b) //[1 2 3 4 5]
	b[0] = 1000
	fmt.Println(a) //[1000 2 3 4 5]
	fmt.Println(b) //[1000 2 3 4 5]
}
```

由于切片是引用类型，所以a和b其实都指向了同一块内存地址。修改b的同时a的值也会发生变化。

Go语言内建的`copy()`函数可以迅速地将一个切片的数据复制到另外一个切片空间中，`copy()`函数的使用格式如下：

```go
copy(destSlice, srcSlice []T)
```

其中：

- srcSlice: 数据来源切片
- destSlice: 目标切片

举个例子：

```go
func main() {
	// copy()复制切片
	a := []int{1, 2, 3, 4, 5}
	c := make([]int, 5, 5)
	copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1 2 3 4 5]
	c[0] = 1000
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1000 2 3 4 5]
}
```

### 2.9 从切片中删除元素

Go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素。 代码如下：

```go
func main() {
	// 从切片中删除元素
	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	a = append(a[:2], a[3:]...)
	fmt.Println(a) //[30 31 33 34 35 36 37]
}
```

总结一下就是：要从切片a中删除索引为`index`的元素，操作方法是`a = append(a[:index], a[index+1:]...)`