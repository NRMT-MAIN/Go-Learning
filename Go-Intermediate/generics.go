package main

import "fmt"



type Stack[T any] struct {
    elements []T
    size int32 
}

func (s *Stack[T]) Push(element T) T {
    s.elements = append(s.elements, element)
    s.size = int32(len(s.elements)) - 1
    return element
}

func (s *Stack[T]) Peek() (T, bool) {
    if len(s.elements) == 0 {
        var zero T
        return zero, false
    }
    return s.elements[len(s.elements)-1], true
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.elements) == 0 {
        var zero T
        return zero, false
    }
    
    index := len(s.elements) - 1
    element := s.elements[index]
    
    s.elements = s.elements[:index]
    s.size = int32(len(s.elements)) - 1
    
    return element, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func swap[T any](a, b T) (T, T) {
	return b, a
}

func genericsExample() {
	x, y := 1, 2
	x, y = swap(x, y)

	fmt.Println(x, y)

	x1, y1 := "Nirmit", "Sahu"
	x1, y1 = swap(x1, y1)

	fmt.Println(x1, y1)

	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(4)
	intStack.Push(5)

	fmt.Println(intStack.Peek())
	fmt.Println(intStack.Pop())
	fmt.Println(intStack.Pop())
	fmt.Println(intStack.Pop())
	fmt.Println(intStack.Pop())

	fmt.Println(intStack.isEmpty())
	fmt.Println(intStack.Peek())

}
