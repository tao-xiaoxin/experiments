package main

//func foo[T comparable, E any](a int, s E) {
//}

//func main() {
//	//foo(5, "hello") // 编译器错误：cannot infer T
//	var s = "hello"
//	foo[int](5, s) //ok
//	foo[int](5, s) //ok
//}

//	func foo[T any](a int) T {
//		var zero T
//		return zero
//	}
//
//	func main() {
//		var a int = foo(5) // 编译器错误：cannot infer T
//		println(a)
//	}

type foo[T1 any, T2 comparable] struct {
	a T1
	b T2
}

func main() {
	//type fooAlias = foo // 编译器错误：cannot use generic type foo[T1 any, T2 comparable] without instantiation
}
