package datastructure

import "fmt"

type Stack struct {
	head *LNode
	len int
}
func NewStack() *Stack{
	return &Stack{nil,0}
}

func (s *Stack) SLen() int{
	return s.len
}

func (s *Stack) IsEmpty() bool{
	return s.len == 0
}

func (s *Stack) Push(val interface{}) {
	s.head = &LNode{val,s.head}
	s.len ++
}

func (s *Stack) Pop() (interface{},bool){
	if s.len == 0 {
		fmt.Println("stack is empty!")
		return nil,false
	}
	tmp := s.head
	s.head = tmp.Next
	tmp.Next = nil
	s.len --
	return tmp.T,true
}

func (s *Stack) Top() (interface{},bool){
	if s.len == 0 {
		fmt.Println("stack is empty!")
		return nil,false
	}
	return s.head.T,true
}

