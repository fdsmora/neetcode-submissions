import (
    "errors"
)

var (
    errStackIsEmpty = errors.New("stack is empty")
)

func isValid(s string) bool {
    if len(s)==1 {
        return false
    }
    var stack = Stack[rune]{}
    closeToOpen := map[rune]rune{'}':'{', ']':'[', ')':'('}

    for _, c := range s {
        if open, exists := closeToOpen[c]; exists{
            if !stack.IsEmpty(){
                top, ok := stack.Pop()
                if !ok || top != open {
                    return false
                }
            } else {
                return false
            }
        } else {
            stack.Push(c)
        }
    }
    return stack.Size()==0
}

type Stack[T any] []T // type Stack[T any] []T

// O(1)
func (s *Stack[T]) Peek() T {
	if len(*s) == 0 {
		var zero T
		return zero
	}
	return (*s)[len(*s)-1]
}

// O(1) amortized
func (s *Stack[T]) Pop() (T, bool) {
	if len(*s) == 0 {
		var zero T
		return zero, false
	}
	idx := len(*s) - 1
	val := (*s)[idx]
	*s = (*s)[:idx] // O(1) amortized
	return val, true
}

// O(1) amortized
func (s *Stack[T]) Push(val T) {
	*s = append(*s, val)
}

func (s *Stack[T]) Size() int {
	return len(*s)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}