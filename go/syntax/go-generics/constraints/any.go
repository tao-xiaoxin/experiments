package main

// any.go
//func doSomething[T1, T2 any](t1 T1, t2 T2) T1 {
//	var a T1 // 声明变量
//	var b T2
//	a, b = t1, t2 // 同类型赋值
//	_ = b
//
//	f := func(t T1) {
//	}
//	f(a) // 传给其他函数
//
//	p := &a // 取变量地址
//	_ = p
//
//	var i interface{} = a // 转换或赋值给interface{}类型变量
//	_ = i
//
//	c := new(T1) // 传递给预定义函数
//	_ = c
//
//	f(a) // 将变量传给其他函数
//
//	sl := make([]T1, 0, 10) // 作为复合类型中的元素类型
//	_ = sl
//
//	j, ok := i.(T1) // 用在类型断言中
//	_ = ok
//	_ = j
//
//	switch i.(type) { // 作为type switch中的case类型
//	case T1:
//	case T2:
//	}
//	return a // 从函数返回
//}

// any.go

//func doSomething[T1, T2 any](t1 T1, t2 T2) T1 {
//	var a T1
//	if a == t1 { // 编译器报错：invalid operation: a == t1 (incomparable types in type set)
//	}
//
//	if a != t1 { // 编译器报错：invalid operation: a != t1 (incomparable types in type set)
//	}
//	... ...
//}
