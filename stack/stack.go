package stack

import "errors"

type StackType interface {
	Pop()
	Push()
}

type Stack struct {
	stk []string
}

func (s *Stack) Pop() (string, error) {
	if len(s.stk) == 0 {
		return "", errors.New("Stack is empty")
	}
	head := s.stk[len(s.stk)-1]
	s.stk = s.stk[:len(s.stk)-1]
	return head, nil
}
func (s *Stack) Push(elem string) {
	s.stk = append(s.stk, elem)
}

func (s *Stack) Head() string {
	if len(s.stk) == 0 {
		return ""
	}
	return s.stk[len(s.stk)-1]
}

func reverse(slice []string) {
	for i := len(slice)/2 - 1; i >= 0; i-- {
		opp := len(slice) - 1 - i
		slice[i], slice[opp] = slice[opp], slice[i]
	}
}
func (s *Stack) Reverse() {
	reverse(s.stk)
}

func New() *Stack {
	return new(Stack)
}
