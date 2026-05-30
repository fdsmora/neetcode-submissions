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
    var stack = NewStack()
    chars := make(map[rune]rune, 3)
    chars['{'] = '}'
    chars['['] = ']'
    chars['('] = ')'

    for _, c := range s {
        switch c {
            case '{', '[', '(':
                stack.Push(chars[c])
            case '}', ']', ')':
                cls, err := stack.Pop()
                if errors.Is(err, errStackIsEmpty) {
                    fmt.Println("stack is empty")
                    return false
                }
                if c != cls {
                    return false
                }
        }
    }

    if stack.Size()>0{
        return false
    }
    
    return true 
}

type Stack struct {
    list []rune
    top int
}

func NewStack() *Stack {
    list := []rune{}
    return &Stack{list, -1}
}

func (s *Stack) Pop() (rune, error) {
    if s.top<0 {
        return 0, errStackIsEmpty
    }
    value := s.list[s.top]
    s.top--
    return value, nil 
}

func (s *Stack) Push(value rune) {
    if s.top >= len(s.list)-1{
        s.list = append(s.list, value)
        s.top=len(s.list)-1
    }else{
        s.list[s.top+1]=value
        s.top++
    }

}

func (s *Stack) Size() int {
    return s.top+1
}
