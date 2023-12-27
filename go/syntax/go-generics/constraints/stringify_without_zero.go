package main

// stringify_without_zero.go

//func StringifyWithoutZero[T fmt.Stringer](s []T) (ret []string) {
//	var zero T
//	for _, v := range s {
//		if v == zero { // 编译器报错：invalid operation: v == zero (incomparable types in type set)
//			continue
//		}
//		ret = append(ret, v.String())
//	}
//	return ret
//}
