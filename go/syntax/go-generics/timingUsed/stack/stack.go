// stack.go
package stack

type Stack[T any] []T

func (s *Stack[T]) Top() (t T) {
	l := len(*s)
	if l == 0 {
		return t
	}
	return (*s)[l-1]
}

func (s *Stack[T]) Push(v T) {
	(*s) = append((*s), v)
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(*s) < 1 {
		return zero, false
	}

	// Get the last element from the stack.
	result := (*s)[len(*s)-1]

	// Remove the last element from the stack.
	*s = (*s)[:len(*s)-1]

	return result, true
}
